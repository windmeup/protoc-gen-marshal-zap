// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/types/types.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/kei2100/protoc-gen-marshal-zap"
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

type Types struct {
	Secret1              string   `protobuf:"bytes,1,opt,name=secret1,proto3" json:"secret1,omitempty"`
	Double1              float64  `protobuf:"fixed64,2,opt,name=double1,proto3" json:"double1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Types) Reset()         { *m = Types{} }
func (m *Types) String() string { return proto.CompactTextString(m) }
func (*Types) ProtoMessage()    {}
func (*Types) Descriptor() ([]byte, []int) {
	return fileDescriptor_61f12dd0de6ed4ab, []int{0}
}

func (m *Types) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Types.Unmarshal(m, b)
}
func (m *Types) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Types.Marshal(b, m, deterministic)
}
func (m *Types) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Types.Merge(m, src)
}
func (m *Types) XXX_Size() int {
	return xxx_messageInfo_Types.Size(m)
}
func (m *Types) XXX_DiscardUnknown() {
	xxx_messageInfo_Types.DiscardUnknown(m)
}

var xxx_messageInfo_Types proto.InternalMessageInfo

func (m *Types) GetSecret1() string {
	if m != nil {
		return m.Secret1
	}
	return ""
}

func (m *Types) GetDouble1() float64 {
	if m != nil {
		return m.Double1
	}
	return 0
}

func init() {
	proto.RegisterType((*Types)(nil), "types.Types")
}

func init() {
	proto.RegisterFile("test/types/types.proto", fileDescriptor_61f12dd0de6ed4ab)
}

var fileDescriptor_61f12dd0de6ed4ab = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x49, 0x2d, 0x2e,
	0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac,
	0x60, 0x8e, 0x94, 0x60, 0x6e, 0x62, 0x51, 0x71, 0x46, 0x62, 0x8e, 0x6e, 0x55, 0x62, 0x01, 0x44,
	0x46, 0xc9, 0x91, 0x8b, 0x35, 0x04, 0x24, 0x27, 0x24, 0xc7, 0xc5, 0x5e, 0x9c, 0x9a, 0x5c, 0x94,
	0x5a, 0x62, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xd2, 0xb1, 0x55, 0x82, 0x31, 0x08,
	0x26, 0x28, 0x24, 0xc1, 0xc5, 0x9e, 0x92, 0x5f, 0x9a, 0x94, 0x93, 0x6a, 0x28, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x18, 0x04, 0xe3, 0x3a, 0xd9, 0x44, 0x59, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0x67, 0xa7, 0x66, 0x1a, 0x19, 0x1a, 0x18, 0xe8, 0x83, 0x8d, 0x4f, 0xd6,
	0x4d, 0x4f, 0xcd, 0xd3, 0x45, 0xb2, 0x55, 0x1f, 0xe1, 0x40, 0x6b, 0x30, 0x99, 0xc4, 0x06, 0x56,
	0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x56, 0x92, 0xfa, 0x88, 0xbb, 0x00, 0x00, 0x00,
}
