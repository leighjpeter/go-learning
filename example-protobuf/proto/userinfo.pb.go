// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/userinfo.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FOO int32

const (
	FOO_X FOO = 0
)

var FOO_name = map[int32]string{
	0: "X",
}

var FOO_value = map[string]int32{
	"X": 0,
}

func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}

func (FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6d6ce1f3572b1fb9, []int{0}
}

//message是固定的。UserInfo是类名，可以随意指定，符合规范即可
type UserInfo struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Length               int32    `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Cnt                  int32    `protobuf:"varint,3,opt,name=cnt,proto3" json:"cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d6ce1f3572b1fb9, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserInfo) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *UserInfo) GetCnt() int32 {
	if m != nil {
		return m.Cnt
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.FOO", FOO_name, FOO_value)
	proto.RegisterType((*UserInfo)(nil), "proto.UserInfo")
}

func init() { proto.RegisterFile("proto/userinfo.proto", fileDescriptor_6d6ce1f3572b1fb9) }

var fileDescriptor_6d6ce1f3572b1fb9 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2d, 0x4e, 0x2d, 0xca, 0xcc, 0x4b, 0xcb, 0xd7, 0x03, 0x73, 0x85, 0x58, 0xc1,
	0x94, 0x92, 0x1f, 0x17, 0x47, 0x68, 0x71, 0x6a, 0x91, 0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x04, 0x17,
	0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c,
	0x2b, 0x24, 0xc6, 0xc5, 0x96, 0x93, 0x9a, 0x97, 0x5e, 0x92, 0x21, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x1a, 0x04, 0xe5, 0x09, 0x09, 0x70, 0x31, 0x27, 0xe7, 0x95, 0x48, 0x30, 0x83, 0x05, 0x41, 0x4c,
	0x2d, 0x1e, 0x2e, 0x66, 0x37, 0x7f, 0x7f, 0x21, 0x56, 0x2e, 0xc6, 0x08, 0x01, 0x86, 0x24, 0x36,
	0xb0, 0x25, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x91, 0x9a, 0xde, 0x83, 0x00, 0x00,
	0x00,
}
