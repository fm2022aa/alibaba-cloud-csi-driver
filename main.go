/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cpfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/dbfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/ens"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/om"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/pov"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	flag.Set("logtostderr", "true")
}

const (
	// TypePluginSuffix is the suffix of all storage plugins.
	TypePluginSuffix = "plugin.csi.alibabacloud.com"
	// TypePluginVar is the yaml variable that needs to be replaced.
	TypePluginVar = "driverplugin.csi.alibabacloud.com-replace"
	// PluginServicePort default port is 11260.
	PluginServicePort = "11260"
	// ProvisionerServicePort default port is 11270.
	ProvisionerServicePort = "11270"
	// TypePluginDBFS local type plugin
	TypePluginDBFS = "dbfsplugin.csi.alibabacloud.com"
	// TypePluginDISK DISK type plugin
	TypePluginDISK = "diskplugin.csi.alibabacloud.com"
	// TypePluginNAS NAS type plugin
	TypePluginNAS = "nasplugin.csi.alibabacloud.com"
	// TypePluginOSS OSS type plugin
	TypePluginOSS = "ossplugin.csi.alibabacloud.com"
	// TypePluginCPFS CPFS type plugin
	TypePluginCPFS = "cpfsplugin.csi.alibabacloud.com"
	// TypePluginENS ENS type plugins
	TypePluginENS = "ensplugin.csi.alibabacloud.com"
	// TypePluginPOV POV type plugins
	TypePluginPOV = "povplugin.csi.alibabacloud.com"
	// ExtenderAgent agent component
	ExtenderAgent = "agent"
)

var (
	endpoint        = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	nodeID          = flag.String("nodeid", "", "node id")
	runAsController = flag.Bool("run-as-controller", false, "Only run as controller service")
	driver          = flag.String("driver", TypePluginDISK, "CSI Driver")
	// Deprecated: rootDir is instead by KUBELET_ROOT_DIR env.
	rootDir = flag.String("rootdir", "/var/lib/kubelet/csi-plugins", "Kubernetes root directory")
)

type globalMetricConfig struct {
	enableMetric bool
	serviceType  string
}

// Nas CSI Plugin
func main() {
	flag.Parse()
	serviceType := os.Getenv(utils.ServiceType)

	if len(serviceType) == 0 || serviceType == "" {
		serviceType = utils.PluginService
	}

	// When serviceType is neither plugin nor provisioner, the program will exits.
	if serviceType != utils.PluginService && serviceType != utils.ProvisionerService {
		log.Fatalf("Service type is unknown:%s", serviceType)
	}
	// enable pprof analyse
	pprofPort := os.Getenv("PPROF_PORT")
	if pprofPort != "" {
		if _, err := strconv.Atoi(pprofPort); err == nil {
			log.Infof("enable pprof & start port at %v", pprofPort)
			go func() {
				err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%v", pprofPort), nil)
				log.Errorf("start server err: %v", err)
			}()
		}
	}

	// setLogAttribute(logAttribute)
	// log.AddHook(rotateHook(logAttribute))

	log.Infof("Multi CSI Driver Name: %s, nodeID: %s, endPoints: %s", *driver, *nodeID, *endpoint)
	log.Infof("CSI Driver, Version: %s, Build time: %s", version.VERSION, version.BUILDTIME)

	multiDriverNames := *driver
	driverNames := strings.Split(multiDriverNames, ",")
	var wg sync.WaitGroup

	// Storage devops
	go om.StorageOM()

	// initialize node metadata
	meta := metadata.NewMetadata()
	meta.EnableEcs(http.DefaultTransport)

	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Warnf("newGlobalConfig: build kubeconfig failed: %v", err)
	} else {
		kubeClient, err := kubernetes.NewForConfig(cfg)
		if err != nil {
			log.Warnf("Error building kubernetes clientset: %v", err)
		} else {
			meta.EnableKubernetes(kubeClient)
		}
	}

	ac := utils.GetAccessControl()
	ecsClient := utils.NewEcsClient(ac)
	stsClient := utils.NewStsClient(ac)
	meta.EnableOpenAPI(ecsClient, stsClient)

	for i, driverName := range driverNames {
		if !strings.Contains(driverName, TypePluginSuffix) && driverName != ExtenderAgent {
			driverNames[i] = joinCsiPluginSuffix(driverName)
		}
	}

	for _, driverName := range driverNames {
		wg.Add(1)
		endPointName := replaceCsiEndpoint(driverName, *endpoint)
		log.Infof("CSI endpoint for driver %s: %s", driverName, endPointName)

		if err := createPersistentStorage(path.Join(utils.KubeletRootDir, "/csi-plugins", driverName, "controller")); err != nil {
			log.Errorf("failed to create persistent storage for controller: %v", err)
			os.Exit(1)
		}
		if err := createPersistentStorage(path.Join(utils.KubeletRootDir, "/csi-plugins", driverName, "node")); err != nil {
			log.Errorf("failed to create persistent storage for node: %v", err)
			os.Exit(1)
		}
		switch driverName {
		case TypePluginNAS:
			go func(endPoint string) {
				defer wg.Done()
				driver := nas.NewDriver(*nodeID, endPoint, serviceType)
				driver.Run()
			}(endPointName)
		case TypePluginOSS:
			go func(endPoint string) {
				defer wg.Done()
				driver := oss.NewDriver(*nodeID, endPoint, meta, *runAsController)
				driver.Run()
			}(endPointName)
		case TypePluginDISK:
			go func(endPoint string) {
				defer wg.Done()
				driver := disk.NewDriver(meta, endPoint, *runAsController)
				driver.Run()
			}(endPointName)

		case TypePluginCPFS:
			go func(endPoint string) {
				defer wg.Done()
				driver := cpfs.NewDriver(*nodeID, endPoint)
				driver.Run()
			}(endPointName)
		case TypePluginDBFS:
			go func(endPoint string) {
				defer wg.Done()
				driver := dbfs.NewDriver(*nodeID, endPoint)
				driver.Run()
			}(endPointName)
		case TypePluginENS:
			go func(endpoint string) {
				defer wg.Done()
				driver := ens.NewDriver(*nodeID, endpoint)
				driver.Run()
			}(endPointName)
		case ExtenderAgent:
			go func() {
				defer wg.Done()
				queryServer := agent.NewAgent()
				queryServer.RunAgent()
			}()
		case TypePluginPOV:
			go func(endPoint string) {
				defer wg.Done()
				driver := pov.NewDriver(*nodeID, endPoint, *runAsController)
				driver.Run()
			}(endPointName)
		default:
			log.Fatalf("CSI start failed, not support driver: %s", driverName)
		}
	}
	servicePort := os.Getenv("SERVICE_PORT")

	if len(servicePort) == 0 || servicePort == "" {
		switch serviceType {
		case utils.PluginService:
			servicePort = PluginServicePort
		case utils.ProvisionerService:
			servicePort = ProvisionerServicePort
		default:
		}
	}

	metricConfig := &globalMetricConfig{
		true,
		"plugin",
	}

	enableMetric := os.Getenv("ENABLE_METRIC")
	version.SetPrometheusVersion()
	if enableMetric == "false" {
		metricConfig.enableMetric = false
	}
	metricConfig.serviceType = serviceType

	log.Info("CSI is running status.")
	csiMux := http.NewServeMux()
	csiMux.HandleFunc("/healthz", healthHandler)
	log.Infof("Metric listening on address: /healthz")
	if metricConfig.enableMetric {
		metricHandler := metric.NewMetricHandler(metricConfig.serviceType, driverNames)
		csiMux.Handle("/metrics", metricHandler)
		log.Infof("Metric listening on address: /metrics")
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", servicePort), csiMux); err != nil {
		log.Fatalf("Service port listen and serve err:%s", err.Error())
	}

	wg.Wait()
	os.Exit(0)
}

func createPersistentStorage(persistentStoragePath string) error {
	log.Infof("Create Stroage Path: %s", persistentStoragePath)
	return os.MkdirAll(persistentStoragePath, os.FileMode(0755))
}

func joinCsiPluginSuffix(storageType string) string {
	return storageType + TypePluginSuffix
}

func replaceCsiEndpoint(pluginType string, endPointName string) string {
	return strings.Replace(endPointName, TypePluginVar, pluginType, -1)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	time := time.Now()
	message := "Liveness probe is OK, time:" + time.String()
	w.Write([]byte(message))
}
