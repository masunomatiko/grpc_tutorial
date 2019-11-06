// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package file

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

type FileRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FileResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileResponse) Reset()         { *m = FileResponse{} }
func (m *FileResponse) String() string { return proto.CompactTextString(m) }
func (*FileResponse) ProtoMessage()    {}
func (*FileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{1}
}

func (m *FileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResponse.Unmarshal(m, b)
}
func (m *FileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResponse.Marshal(b, m, deterministic)
}
func (m *FileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResponse.Merge(m, src)
}
func (m *FileResponse) XXX_Size() int {
	return xxx_messageInfo_FileResponse.Size(m)
}
func (m *FileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileResponse proto.InternalMessageInfo

func (m *FileResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*FileRequest)(nil), "file.FileRequest")
	proto.RegisterType((*FileResponse)(nil), "file.FileResponse")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162) }

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0xcc, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x14, 0xb9, 0xb8, 0xdd, 0x32,
	0x73, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x25, 0x2e, 0x1e, 0x88, 0x92,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90, 0x9a, 0x94, 0xc4, 0x92, 0x44, 0xb0, 0x1a, 0x9e, 0x20,
	0x30, 0xdb, 0xc8, 0x09, 0x62, 0x4c, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x31, 0x17,
	0x87, 0x4b, 0x7e, 0x79, 0x5e, 0x4e, 0x7e, 0x62, 0x8a, 0x90, 0xa0, 0x1e, 0xd8, 0x52, 0x24, 0x5b,
	0xa4, 0x84, 0x90, 0x85, 0x20, 0xa6, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0xdd, 0x65, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x7f, 0x12, 0xe7, 0xec, 0xa5, 0x00, 0x00, 0x00,
}