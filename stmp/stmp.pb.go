// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stmp/stmp.proto

package stmp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// the service kind
// if service is true, will generate server api, such as STMPRegisterXxxServer, STMPXxxClient
// if events is true, will generate events api, such as STMPRegisterXxxListener, STMPXxxBroadcaster
// if both not set, will detect by service name suffix
// if ends with Service, will treat as service
// else if ends with Events, will treat as events
// else both will be true default
type ServiceKind struct {
	Service              bool     `protobuf:"varint,1,opt,name=service,proto3" json:"service,omitempty"`
	Events               bool     `protobuf:"varint,2,opt,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceKind) Reset()         { *m = ServiceKind{} }
func (m *ServiceKind) String() string { return proto.CompactTextString(m) }
func (*ServiceKind) ProtoMessage()    {}
func (*ServiceKind) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8e5dfa1586309a1, []int{0}
}

func (m *ServiceKind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceKind.Unmarshal(m, b)
}
func (m *ServiceKind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceKind.Marshal(b, m, deterministic)
}
func (m *ServiceKind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceKind.Merge(m, src)
}
func (m *ServiceKind) XXX_Size() int {
	return xxx_messageInfo_ServiceKind.Size(m)
}
func (m *ServiceKind) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceKind.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceKind proto.InternalMessageInfo

func (m *ServiceKind) GetService() bool {
	if m != nil {
		return m.Service
	}
	return false
}

func (m *ServiceKind) GetEvents() bool {
	if m != nil {
		return m.Events
	}
	return false
}

var E_Service = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         5588225,
	Name:          "stmp.service",
	Tag:           "varint,5588225,opt,name=service",
	Filename:      "stmp/stmp.proto",
}

var E_Kind = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*ServiceKind)(nil),
	Field:         5588226,
	Name:          "stmp.kind",
	Tag:           "bytes,5588226,opt,name=kind",
	Filename:      "stmp/stmp.proto",
}

var E_Method = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         5588226,
	Name:          "stmp.method",
	Tag:           "varint,5588226,opt,name=method",
	Filename:      "stmp/stmp.proto",
}

func init() {
	proto.RegisterType((*ServiceKind)(nil), "stmp.ServiceKind")
	proto.RegisterExtension(E_Service)
	proto.RegisterExtension(E_Kind)
	proto.RegisterExtension(E_Method)
}

func init() { proto.RegisterFile("stmp/stmp.proto", fileDescriptor_e8e5dfa1586309a1) }

var fileDescriptor_e8e5dfa1586309a1 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2e, 0xc9, 0x2d,
	0xd0, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x94, 0x42, 0x7a,
	0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x2c, 0xa9, 0x34, 0x4d, 0x3f, 0x25, 0xb5, 0x38, 0xb9,
	0x28, 0xb3, 0xa0, 0x24, 0xbf, 0x08, 0xa2, 0x4e, 0xc9, 0x9e, 0x8b, 0x3b, 0x38, 0xb5, 0xa8, 0x2c,
	0x33, 0x39, 0xd5, 0x3b, 0x33, 0x2f, 0x45, 0x48, 0x82, 0x8b, 0xbd, 0x18, 0xc2, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x08, 0x82, 0x71, 0x85, 0xc4, 0xb8, 0xd8, 0x52, 0xcb, 0x52, 0xf3, 0x4a, 0x8a,
	0x25, 0x98, 0xc0, 0x12, 0x50, 0x9e, 0x95, 0x0d, 0x5c, 0x87, 0x90, 0xbc, 0x1e, 0xc4, 0x3a, 0x3d,
	0x98, 0x75, 0x7a, 0x50, 0xa3, 0xfd, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x8a, 0x25, 0x1a, 0xbb, 0xae,
	0x82, 0x34, 0xb3, 0xc0, 0x4d, 0xb5, 0xf2, 0xe4, 0x62, 0xc9, 0x06, 0xd9, 0x4b, 0x50, 0x6b, 0x13,
	0x44, 0x2b, 0xb7, 0x91, 0xa0, 0x1e, 0xd8, 0x97, 0x48, 0x8e, 0x0e, 0x02, 0x1b, 0x61, 0x65, 0xc9,
	0xc5, 0x96, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0x22, 0x24, 0x87, 0x61, 0x98, 0x2f, 0x58, 0x02, 0xcd,
	0x2c, 0x96, 0x20, 0xa8, 0x06, 0x27, 0xa5, 0x28, 0x85, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0xfd, 0xc4, 0xe4, 0xa2, 0xc4, 0xaa, 0xcc, 0xbc, 0x74, 0x70, 0x70, 0xea, 0xa6,
	0xe7, 0x83, 0xe9, 0x24, 0x36, 0xb0, 0x61, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x69,
	0x7b, 0xd9, 0x6a, 0x01, 0x00, 0x00,
}
