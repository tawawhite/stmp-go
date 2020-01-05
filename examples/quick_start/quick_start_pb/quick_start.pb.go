// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/quick_start/quick_start_pb/quick_start.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	io "io"
	math "math"
	math_bits "math/bits"
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

type RoomModel struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Users                []string `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomModel) Reset()         { *m = RoomModel{} }
func (m *RoomModel) String() string { return proto.CompactTextString(m) }
func (*RoomModel) ProtoMessage()    {}
func (*RoomModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bfa63495ffd955, []int{0}
}
func (m *RoomModel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoomModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoomModel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoomModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomModel.Merge(m, src)
}
func (m *RoomModel) XXX_Size() int {
	return m.Size()
}
func (m *RoomModel) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomModel.DiscardUnknown(m)
}

var xxx_messageInfo_RoomModel proto.InternalMessageInfo

type JoinInput struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinInput) Reset()         { *m = JoinInput{} }
func (m *JoinInput) String() string { return proto.CompactTextString(m) }
func (*JoinInput) ProtoMessage()    {}
func (*JoinInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bfa63495ffd955, []int{1}
}
func (m *JoinInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinInput.Merge(m, src)
}
func (m *JoinInput) XXX_Size() int {
	return m.Size()
}
func (m *JoinInput) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinInput.DiscardUnknown(m)
}

var xxx_messageInfo_JoinInput proto.InternalMessageInfo

type ExitInput struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitInput) Reset()         { *m = ExitInput{} }
func (m *ExitInput) String() string { return proto.CompactTextString(m) }
func (*ExitInput) ProtoMessage()    {}
func (*ExitInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bfa63495ffd955, []int{2}
}
func (m *ExitInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExitInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExitInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExitInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitInput.Merge(m, src)
}
func (m *ExitInput) XXX_Size() int {
	return m.Size()
}
func (m *ExitInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitInput.DiscardUnknown(m)
}

var xxx_messageInfo_ExitInput proto.InternalMessageInfo

type JoinEvent struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinEvent) Reset()         { *m = JoinEvent{} }
func (m *JoinEvent) String() string { return proto.CompactTextString(m) }
func (*JoinEvent) ProtoMessage()    {}
func (*JoinEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bfa63495ffd955, []int{3}
}
func (m *JoinEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinEvent.Merge(m, src)
}
func (m *JoinEvent) XXX_Size() int {
	return m.Size()
}
func (m *JoinEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinEvent.DiscardUnknown(m)
}

var xxx_messageInfo_JoinEvent proto.InternalMessageInfo

type ExitEvent struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitEvent) Reset()         { *m = ExitEvent{} }
func (m *ExitEvent) String() string { return proto.CompactTextString(m) }
func (*ExitEvent) ProtoMessage()    {}
func (*ExitEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bfa63495ffd955, []int{4}
}
func (m *ExitEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExitEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExitEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExitEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitEvent.Merge(m, src)
}
func (m *ExitEvent) XXX_Size() int {
	return m.Size()
}
func (m *ExitEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ExitEvent proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RoomModel)(nil), "stmp.examples.quick_start.RoomModel")
	proto.RegisterType((*JoinInput)(nil), "stmp.examples.quick_start.JoinInput")
	proto.RegisterType((*ExitInput)(nil), "stmp.examples.quick_start.ExitInput")
	proto.RegisterType((*JoinEvent)(nil), "stmp.examples.quick_start.JoinEvent")
	proto.RegisterType((*ExitEvent)(nil), "stmp.examples.quick_start.ExitEvent")
}

func init() {
	proto.RegisterFile("examples/quick_start/quick_start_pb/quick_start.proto", fileDescriptor_86bfa63495ffd955)
}

var fileDescriptor_86bfa63495ffd955 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x4e, 0xf3, 0x30,
	0x14, 0xc7, 0xeb, 0xef, 0x2b, 0x48, 0x31, 0x9b, 0x85, 0x50, 0x09, 0x52, 0xa8, 0x22, 0x86, 0x2e,
	0x75, 0xa4, 0xa2, 0x4e, 0x6c, 0x95, 0x8a, 0x04, 0x12, 0x4b, 0xe9, 0xc4, 0x52, 0x25, 0xc1, 0x18,
	0x8b, 0x3a, 0xcf, 0xc4, 0x4e, 0x55, 0x38, 0x0b, 0x37, 0xe0, 0x22, 0x1d, 0x39, 0x02, 0xf4, 0x24,
	0xc8, 0x8e, 0x52, 0x32, 0xd0, 0x92, 0xed, 0x3d, 0xf9, 0xf7, 0x7e, 0x7e, 0xfe, 0xcb, 0x78, 0xc8,
	0x96, 0xb1, 0x54, 0x73, 0xa6, 0xa3, 0xe7, 0x42, 0xa4, 0x4f, 0x33, 0x6d, 0xe2, 0xdc, 0xd4, 0xeb,
	0x99, 0x4a, 0xea, 0x2d, 0x55, 0x39, 0x18, 0x20, 0xc7, 0xda, 0x48, 0x45, 0xab, 0x59, 0x5a, 0x03,
	0xfc, 0x3e, 0x17, 0xe6, 0xb1, 0x48, 0x68, 0x0a, 0x32, 0xe2, 0xc0, 0x21, 0x72, 0x13, 0x49, 0xf1,
	0xe0, 0x3a, 0xd7, 0xb8, 0xaa, 0x34, 0xf9, 0x27, 0x1c, 0x80, 0xcf, 0xd9, 0x0f, 0xc5, 0xa4, 0x32,
	0x2f, 0xe5, 0x61, 0x38, 0xc4, 0xde, 0x04, 0x40, 0xde, 0xc0, 0x3d, 0x9b, 0x13, 0x82, 0xdb, 0x59,
	0x2c, 0x59, 0x07, 0x75, 0x51, 0xcf, 0x9b, 0xb8, 0x9a, 0x1c, 0xe2, 0xbd, 0x42, 0xb3, 0x5c, 0x77,
	0xfe, 0x75, 0xff, 0xf7, 0xbc, 0x49, 0xd9, 0x84, 0xa7, 0xd8, 0xbb, 0x06, 0x91, 0x5d, 0x65, 0xaa,
	0x30, 0xbf, 0x8d, 0x59, 0x60, 0xbc, 0x14, 0x66, 0x27, 0x60, 0x0d, 0xe3, 0x05, 0xcb, 0x1c, 0x60,
	0xbd, 0x15, 0x60, 0xeb, 0xca, 0xb0, 0x15, 0x18, 0xbc, 0x23, 0x7c, 0x60, 0x77, 0xbf, 0x65, 0xf9,
	0x42, 0xa4, 0x8c, 0x4c, 0x71, 0xdb, 0x1a, 0xc9, 0x19, 0xdd, 0x1a, 0x1d, 0xdd, 0x2c, 0xed, 0xef,
	0xa2, 0x36, 0x89, 0x84, 0x2d, 0x72, 0x89, 0xdb, 0x76, 0x8d, 0x9d, 0xd6, 0xcd, 0x4b, 0xfd, 0x23,
	0x5a, 0x86, 0x4d, 0xab, 0xb0, 0xe9, 0xd8, 0x86, 0x1d, 0xb6, 0x06, 0x6f, 0x08, 0x63, 0xeb, 0x75,
	0xef, 0xd1, 0x56, 0xdb, 0x68, 0x59, 0x87, 0x6f, 0xd7, 0x36, 0x5e, 0xef, 0x0f, 0xcf, 0x68, 0xba,
	0xfa, 0x0a, 0x5a, 0xab, 0x75, 0x80, 0x3e, 0xd6, 0x01, 0xfa, 0x5c, 0x07, 0xe8, 0x6e, 0x54, 0xfb,
	0x65, 0x71, 0x9a, 0xc7, 0xaf, 0x22, 0xe3, 0x91, 0xbd, 0xa4, 0xcf, 0x21, 0x6a, 0xf0, 0xa7, 0x2f,
	0x54, 0x92, 0xec, 0xbb, 0x7b, 0xce, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x4b, 0x42, 0x27,
	0x04, 0x03, 0x00, 0x00,
}

func (m *RoomModel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomModel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoomModel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Users) > 0 {
		for iNdEx := len(m.Users) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Users[iNdEx])
			copy(dAtA[i:], m.Users[iNdEx])
			i = encodeVarintQuickStart(dAtA, i, uint64(len(m.Users[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuickStart(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JoinInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuickStart(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExitInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExitInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExitInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuickStart(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JoinEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQuickStart(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExitEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExitEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExitEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQuickStart(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuickStart(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuickStart(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RoomModel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuickStart(uint64(l))
	}
	if len(m.Users) > 0 {
		for _, s := range m.Users {
			l = len(s)
			n += 1 + l + sovQuickStart(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *JoinInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuickStart(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExitInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuickStart(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *JoinEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQuickStart(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExitEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQuickStart(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovQuickStart(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuickStart(x uint64) (n int) {
	return sovQuickStart(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoomModel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuickStart
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
			return fmt.Errorf("proto: RoomModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoomModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuickStart
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
				return ErrInvalidLengthQuickStart
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuickStart
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuickStart
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
				return ErrInvalidLengthQuickStart
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuickStart
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuickStart(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuickStart
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuickStart
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
func (m *JoinInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuickStart
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
			return fmt.Errorf("proto: JoinInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuickStart
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
				return ErrInvalidLengthQuickStart
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuickStart
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuickStart(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuickStart
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuickStart
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
func (m *ExitInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuickStart
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
			return fmt.Errorf("proto: ExitInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExitInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuickStart
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
				return ErrInvalidLengthQuickStart
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuickStart
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuickStart(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuickStart
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuickStart
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
func (m *JoinEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuickStart
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
			return fmt.Errorf("proto: JoinEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuickStart
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
				return ErrInvalidLengthQuickStart
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuickStart
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuickStart(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuickStart
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuickStart
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
func (m *ExitEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuickStart
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
			return fmt.Errorf("proto: ExitEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExitEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuickStart
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
				return ErrInvalidLengthQuickStart
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuickStart
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuickStart(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuickStart
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuickStart
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
func skipQuickStart(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuickStart
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
					return 0, ErrIntOverflowQuickStart
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
					return 0, ErrIntOverflowQuickStart
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
				return 0, ErrInvalidLengthQuickStart
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuickStart
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuickStart
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuickStart        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuickStart          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuickStart = fmt.Errorf("proto: unexpected end of group")
)