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

type JoinRoomInput struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomInput) Reset()         { *m = JoinRoomInput{} }
func (m *JoinRoomInput) String() string { return proto.CompactTextString(m) }
func (*JoinRoomInput) ProtoMessage()    {}
func (*JoinRoomInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bfa63495ffd955, []int{1}
}
func (m *JoinRoomInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinRoomInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinRoomInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinRoomInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomInput.Merge(m, src)
}
func (m *JoinRoomInput) XXX_Size() int {
	return m.Size()
}
func (m *JoinRoomInput) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomInput.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomInput proto.InternalMessageInfo

type ExitRoomInput struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitRoomInput) Reset()         { *m = ExitRoomInput{} }
func (m *ExitRoomInput) String() string { return proto.CompactTextString(m) }
func (*ExitRoomInput) ProtoMessage()    {}
func (*ExitRoomInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bfa63495ffd955, []int{2}
}
func (m *ExitRoomInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExitRoomInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExitRoomInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExitRoomInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitRoomInput.Merge(m, src)
}
func (m *ExitRoomInput) XXX_Size() int {
	return m.Size()
}
func (m *ExitRoomInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitRoomInput.DiscardUnknown(m)
}

var xxx_messageInfo_ExitRoomInput proto.InternalMessageInfo

type UserJoinEvent struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserJoinEvent) Reset()         { *m = UserJoinEvent{} }
func (m *UserJoinEvent) String() string { return proto.CompactTextString(m) }
func (*UserJoinEvent) ProtoMessage()    {}
func (*UserJoinEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bfa63495ffd955, []int{3}
}
func (m *UserJoinEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserJoinEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserJoinEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserJoinEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserJoinEvent.Merge(m, src)
}
func (m *UserJoinEvent) XXX_Size() int {
	return m.Size()
}
func (m *UserJoinEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UserJoinEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UserJoinEvent proto.InternalMessageInfo

type UserExitEvent struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserExitEvent) Reset()         { *m = UserExitEvent{} }
func (m *UserExitEvent) String() string { return proto.CompactTextString(m) }
func (*UserExitEvent) ProtoMessage()    {}
func (*UserExitEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bfa63495ffd955, []int{4}
}
func (m *UserExitEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserExitEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserExitEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserExitEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserExitEvent.Merge(m, src)
}
func (m *UserExitEvent) XXX_Size() int {
	return m.Size()
}
func (m *UserExitEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UserExitEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UserExitEvent proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RoomModel)(nil), "stmp.examples.quick_start.RoomModel")
	proto.RegisterType((*JoinRoomInput)(nil), "stmp.examples.quick_start.JoinRoomInput")
	proto.RegisterType((*ExitRoomInput)(nil), "stmp.examples.quick_start.ExitRoomInput")
	proto.RegisterType((*UserJoinEvent)(nil), "stmp.examples.quick_start.UserJoinEvent")
	proto.RegisterType((*UserExitEvent)(nil), "stmp.examples.quick_start.UserExitEvent")
}

func init() {
	proto.RegisterFile("examples/quick_start/quick_start_pb/quick_start.proto", fileDescriptor_86bfa63495ffd955)
}

var fileDescriptor_86bfa63495ffd955 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0xfe, 0xa3, 0x5d, 0xf1, 0xb2, 0x88, 0xd4, 0x08, 0xa1, 0x44, 0x0f, 0xbd, 0x74,
	0x03, 0x95, 0x9e, 0xbc, 0x15, 0x7a, 0x50, 0xd0, 0x43, 0xd5, 0x8b, 0x08, 0x25, 0x89, 0xeb, 0x1a,
	0x6c, 0xb2, 0x6b, 0x76, 0x53, 0xaa, 0xcf, 0xe4, 0x1b, 0xf8, 0x02, 0x3d, 0xfa, 0x08, 0xda, 0x27,
	0x91, 0xd9, 0xb0, 0x35, 0x05, 0xd3, 0x7a, 0x9b, 0x21, 0xdf, 0xfc, 0x66, 0xe6, 0x9b, 0x2c, 0xee,
	0xb1, 0x69, 0x90, 0xc8, 0x31, 0x53, 0xfe, 0x4b, 0x1e, 0x47, 0xcf, 0x23, 0xa5, 0x83, 0x4c, 0x97,
	0xe3, 0x91, 0x0c, 0xcb, 0x29, 0x95, 0x99, 0xd0, 0x82, 0x1c, 0x2a, 0x9d, 0x48, 0x6a, 0x6b, 0x69,
	0x49, 0xe0, 0x74, 0x78, 0xac, 0x9f, 0xf2, 0x90, 0x46, 0x22, 0xf1, 0xb9, 0xe0, 0xc2, 0x37, 0x15,
	0x61, 0xfe, 0x68, 0x32, 0x93, 0x98, 0xa8, 0x20, 0x39, 0x47, 0x5c, 0x08, 0x3e, 0x66, 0xbf, 0x2a,
	0x96, 0x48, 0xfd, 0x5a, 0x7c, 0xf4, 0x7a, 0xb8, 0x31, 0x14, 0x22, 0xb9, 0x14, 0x0f, 0x6c, 0x4c,
	0x08, 0xde, 0x4a, 0x83, 0x84, 0x35, 0x51, 0x0b, 0xb5, 0x1b, 0x43, 0x13, 0x93, 0x7d, 0xbc, 0x9d,
	0x2b, 0x96, 0xa9, 0xe6, 0x46, 0x6b, 0xb3, 0xdd, 0x18, 0x16, 0x89, 0x77, 0x8c, 0xf7, 0x2e, 0x44,
	0x9c, 0x42, 0xe9, 0x79, 0x2a, 0x73, 0xfd, 0x57, 0x29, 0x88, 0x06, 0xd3, 0x58, 0xaf, 0x15, 0xdd,
	0x2a, 0x96, 0x01, 0x6d, 0x30, 0x61, 0xa9, 0x11, 0x41, 0x0f, 0x2b, 0x82, 0xd8, 0x8a, 0x80, 0x56,
	0x29, 0xea, 0x7e, 0x20, 0xbc, 0x0b, 0xbd, 0xae, 0x59, 0x36, 0x89, 0x23, 0x46, 0xee, 0x71, 0xdd,
	0xce, 0x48, 0xda, 0xb4, 0xd2, 0x4e, 0xba, 0xb4, 0x88, 0x73, 0xb2, 0x42, 0xb9, 0x70, 0xca, 0xab,
	0x91, 0x2b, 0x5c, 0xb7, 0xcb, 0xad, 0xa4, 0x2f, 0x39, 0xe0, 0x1c, 0xd0, 0xe2, 0x18, 0xd4, 0x1e,
	0x83, 0x0e, 0xe0, 0x18, 0x5e, 0xad, 0xfb, 0x8e, 0x30, 0x06, 0x9d, 0xd9, 0x4f, 0x01, 0xde, 0xda,
	0xb2, 0x12, 0xbf, 0xe4, 0x5d, 0x35, 0xde, 0xf2, 0x60, 0x9a, 0xb5, 0xbc, 0x85, 0xcd, 0xd5, 0xbc,
	0xfe, 0xcd, 0xec, 0xdb, 0xad, 0xcd, 0xe6, 0x2e, 0xfa, 0x9c, 0xbb, 0xe8, 0x6b, 0xee, 0xa2, 0xbb,
	0x7e, 0xe9, 0xaf, 0x0c, 0xa2, 0x2c, 0x78, 0x8b, 0x53, 0xee, 0x43, 0xa3, 0x0e, 0x17, 0xfe, 0x3f,
	0xde, 0xc0, 0x99, 0x0c, 0xc3, 0x1d, 0xd3, 0xe7, 0xf4, 0x27, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x28,
	0x3b, 0xa2, 0x34, 0x03, 0x00, 0x00,
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

func (m *JoinRoomInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinRoomInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinRoomInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *ExitRoomInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExitRoomInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExitRoomInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *UserJoinEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserJoinEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserJoinEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *UserExitEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserExitEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserExitEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *JoinRoomInput) Size() (n int) {
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

func (m *ExitRoomInput) Size() (n int) {
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

func (m *UserJoinEvent) Size() (n int) {
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

func (m *UserExitEvent) Size() (n int) {
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
func (m *JoinRoomInput) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: JoinRoomInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinRoomInput: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *ExitRoomInput) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ExitRoomInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExitRoomInput: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *UserJoinEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UserJoinEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserJoinEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *UserExitEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UserExitEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserExitEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
