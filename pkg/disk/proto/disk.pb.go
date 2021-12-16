// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: disk.proto

package proto

import (
	context "context"
	fmt "fmt"
	github_com_containerd_ttrpc "github.com/containerd/ttrpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// 文件系统扩容请求消息
type ExpandVolumeRequest struct {
	Volume               string   `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpandVolumeRequest) Reset()      { *m = ExpandVolumeRequest{} }
func (*ExpandVolumeRequest) ProtoMessage() {}
func (*ExpandVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96b80c5532b4167, []int{0}
}
func (m *ExpandVolumeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExpandVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExpandVolumeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpandVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpandVolumeRequest.Merge(m, src)
}
func (m *ExpandVolumeRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExpandVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpandVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExpandVolumeRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExpandVolumeRequest)(nil), "katacontainers.extends.v1.ExpandVolumeRequest")
}

func init() { proto.RegisterFile("disk.proto", fileDescriptor_f96b80c5532b4167) }

var fileDescriptor_f96b80c5532b4167 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xc9, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xcc, 0x4e, 0x2c, 0x49, 0x4c, 0xce, 0xcf, 0x2b,
	0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0x2a, 0xd6, 0x4b, 0xad, 0x28, 0x49, 0xcd, 0x4b, 0x29, 0xd6, 0x2b,
	0x33, 0x94, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x2b, 0x4c, 0x2a, 0x4d, 0xd3,
	0x4f, 0xcd, 0x2d, 0x28, 0xa9, 0x84, 0xe8, 0x53, 0xd2, 0xe5, 0x12, 0x76, 0xad, 0x28, 0x48, 0xcc,
	0x4b, 0x09, 0xcb, 0xcf, 0x29, 0xcd, 0x4d, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe3, 0x62, 0x2b, 0x03, 0x0b, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x46, 0x59,
	0x5c, 0x7c, 0xae, 0x60, 0x93, 0x53, 0x53, 0x82, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x85, 0x22, 0xb8,
	0x78, 0x90, 0x0d, 0x10, 0xd2, 0xd3, 0xc3, 0xe9, 0x12, 0x3d, 0x2c, 0x36, 0x49, 0x89, 0xe9, 0x41,
	0x9c, 0xa7, 0x07, 0x73, 0x9e, 0x9e, 0x2b, 0xc8, 0x79, 0x4a, 0x0c, 0x4e, 0x0a, 0x27, 0x1e, 0xca,
	0x31, 0xdc, 0x78, 0x28, 0xc7, 0xd0, 0xf0, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x8c, 0x62, 0x83, 0xfa, 0x85, 0x0d, 0x4c, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x98, 0x8a, 0x1e, 0xce, 0x09, 0x01, 0x00, 0x00,
}

func (m *ExpandVolumeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpandVolumeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExpandVolumeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Volume) > 0 {
		i -= len(m.Volume)
		copy(dAtA[i:], m.Volume)
		i = encodeVarintDisk(dAtA, i, uint64(len(m.Volume)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDisk(dAtA []byte, offset int, v uint64) int {
	offset -= sovDisk(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExpandVolumeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Volume)
	if l > 0 {
		n += 1 + l + sovDisk(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDisk(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDisk(x uint64) (n int) {
	return sovDisk(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ExpandVolumeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExpandVolumeRequest{`,
		`Volume:` + fmt.Sprintf("%v", this.Volume) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDisk(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

type ExtendedStatusService interface {
	ExpandVolume(ctx context.Context, req *ExpandVolumeRequest) (*types.Empty, error)
}

func RegisterExtendedStatusService(srv *github_com_containerd_ttrpc.Server, svc ExtendedStatusService) {
	srv.Register("katacontainers.extends.v1.ExtendedStatus", map[string]github_com_containerd_ttrpc.Method{
		"ExpandVolume": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req ExpandVolumeRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.ExpandVolume(ctx, &req)
		},
	})
}

type extendedStatusClient struct {
	client *github_com_containerd_ttrpc.Client
}

func NewExtendedStatusClient(client *github_com_containerd_ttrpc.Client) ExtendedStatusService {
	return &extendedStatusClient{
		client: client,
	}
}

func (c *extendedStatusClient) ExpandVolume(ctx context.Context, req *ExpandVolumeRequest) (*types.Empty, error) {
	var resp types.Empty
	if err := c.client.Call(ctx, "katacontainers.extends.v1.ExtendedStatus", "ExpandVolume", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
func (m *ExpandVolumeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDisk
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExpandVolumeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExpandVolumeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volume", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDisk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Volume = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDisk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDisk
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDisk(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDisk
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDisk
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDisk
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDisk
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDisk
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDisk
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDisk        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDisk          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDisk = fmt.Errorf("proto: unexpected end of group")
)
