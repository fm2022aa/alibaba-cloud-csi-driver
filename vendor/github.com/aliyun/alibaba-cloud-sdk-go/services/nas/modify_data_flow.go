package nas

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyDataFlow invokes the nas.ModifyDataFlow API synchronously
func (client *Client) ModifyDataFlow(request *ModifyDataFlowRequest) (response *ModifyDataFlowResponse, err error) {
	response = CreateModifyDataFlowResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDataFlowWithChan invokes the nas.ModifyDataFlow API asynchronously
func (client *Client) ModifyDataFlowWithChan(request *ModifyDataFlowRequest) (<-chan *ModifyDataFlowResponse, <-chan error) {
	responseChan := make(chan *ModifyDataFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDataFlow(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyDataFlowWithCallback invokes the nas.ModifyDataFlow API asynchronously
func (client *Client) ModifyDataFlowWithCallback(request *ModifyDataFlowRequest, callback func(response *ModifyDataFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDataFlowResponse
		var err error
		defer close(result)
		response, err = client.ModifyDataFlow(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyDataFlowRequest is the request struct for api ModifyDataFlow
type ModifyDataFlowRequest struct {
	*requests.RpcRequest
	ClientToken  string           `position:"Query" name:"ClientToken"`
	Description  string           `position:"Query" name:"Description"`
	Throughput   requests.Integer `position:"Query" name:"Throughput"`
	FileSystemId string           `position:"Query" name:"FileSystemId"`
	DryRun       requests.Boolean `position:"Query" name:"DryRun"`
	DataFlowId   string           `position:"Query" name:"DataFlowId"`
}

// ModifyDataFlowResponse is the response struct for api ModifyDataFlow
type ModifyDataFlowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDataFlowRequest creates a request to invoke ModifyDataFlow API
func CreateModifyDataFlowRequest() (request *ModifyDataFlowRequest) {
	request = &ModifyDataFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "ModifyDataFlow", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDataFlowResponse creates a response to parse from ModifyDataFlow response
func CreateModifyDataFlowResponse() (response *ModifyDataFlowResponse) {
	response = &ModifyDataFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}