// Code generated by protoc-gen-go. DO NOT EDIT.
// source: FileGetContents.proto

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

// Get the contents of a file. The content field is empty (no bytes) if the file is empty.
type FileGetContentsQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	FileID               *FileID      `protobuf:"bytes,2,opt,name=fileID,proto3" json:"fileID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FileGetContentsQuery) Reset()         { *m = FileGetContentsQuery{} }
func (m *FileGetContentsQuery) String() string { return proto.CompactTextString(m) }
func (*FileGetContentsQuery) ProtoMessage()    {}
func (*FileGetContentsQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_13ec89ff2179f06a, []int{0}
}

func (m *FileGetContentsQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileGetContentsQuery.Unmarshal(m, b)
}
func (m *FileGetContentsQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileGetContentsQuery.Marshal(b, m, deterministic)
}
func (m *FileGetContentsQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileGetContentsQuery.Merge(m, src)
}
func (m *FileGetContentsQuery) XXX_Size() int {
	return xxx_messageInfo_FileGetContentsQuery.Size(m)
}
func (m *FileGetContentsQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FileGetContentsQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FileGetContentsQuery proto.InternalMessageInfo

func (m *FileGetContentsQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FileGetContentsQuery) GetFileID() *FileID {
	if m != nil {
		return m.FileID
	}
	return nil
}

// Response when the client sends the node FileGetContentsQuery
type FileGetContentsResponse struct {
	Header               *ResponseHeader                       `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	FileContents         *FileGetContentsResponse_FileContents `protobuf:"bytes,2,opt,name=fileContents,proto3" json:"fileContents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *FileGetContentsResponse) Reset()         { *m = FileGetContentsResponse{} }
func (m *FileGetContentsResponse) String() string { return proto.CompactTextString(m) }
func (*FileGetContentsResponse) ProtoMessage()    {}
func (*FileGetContentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13ec89ff2179f06a, []int{1}
}

func (m *FileGetContentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileGetContentsResponse.Unmarshal(m, b)
}
func (m *FileGetContentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileGetContentsResponse.Marshal(b, m, deterministic)
}
func (m *FileGetContentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileGetContentsResponse.Merge(m, src)
}
func (m *FileGetContentsResponse) XXX_Size() int {
	return xxx_messageInfo_FileGetContentsResponse.Size(m)
}
func (m *FileGetContentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileGetContentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileGetContentsResponse proto.InternalMessageInfo

func (m *FileGetContentsResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FileGetContentsResponse) GetFileContents() *FileGetContentsResponse_FileContents {
	if m != nil {
		return m.FileContents
	}
	return nil
}

type FileGetContentsResponse_FileContents struct {
	FileID               *FileID  `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
	Contents             []byte   `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileGetContentsResponse_FileContents) Reset()         { *m = FileGetContentsResponse_FileContents{} }
func (m *FileGetContentsResponse_FileContents) String() string { return proto.CompactTextString(m) }
func (*FileGetContentsResponse_FileContents) ProtoMessage()    {}
func (*FileGetContentsResponse_FileContents) Descriptor() ([]byte, []int) {
	return fileDescriptor_13ec89ff2179f06a, []int{1, 0}
}

func (m *FileGetContentsResponse_FileContents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileGetContentsResponse_FileContents.Unmarshal(m, b)
}
func (m *FileGetContentsResponse_FileContents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileGetContentsResponse_FileContents.Marshal(b, m, deterministic)
}
func (m *FileGetContentsResponse_FileContents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileGetContentsResponse_FileContents.Merge(m, src)
}
func (m *FileGetContentsResponse_FileContents) XXX_Size() int {
	return xxx_messageInfo_FileGetContentsResponse_FileContents.Size(m)
}
func (m *FileGetContentsResponse_FileContents) XXX_DiscardUnknown() {
	xxx_messageInfo_FileGetContentsResponse_FileContents.DiscardUnknown(m)
}

var xxx_messageInfo_FileGetContentsResponse_FileContents proto.InternalMessageInfo

func (m *FileGetContentsResponse_FileContents) GetFileID() *FileID {
	if m != nil {
		return m.FileID
	}
	return nil
}

func (m *FileGetContentsResponse_FileContents) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*FileGetContentsQuery)(nil), "proto.FileGetContentsQuery")
	proto.RegisterType((*FileGetContentsResponse)(nil), "proto.FileGetContentsResponse")
	proto.RegisterType((*FileGetContentsResponse_FileContents)(nil), "proto.FileGetContentsResponse.FileContents")
}

func init() { proto.RegisterFile("FileGetContents.proto", fileDescriptor_13ec89ff2179f06a) }

var fileDescriptor_13ec89ff2179f06a = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x75, 0xcb, 0xcc, 0x49,
	0x75, 0x4f, 0x2d, 0x71, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x02, 0x4e, 0x89, 0xc5, 0x99, 0xc9, 0x21, 0x95, 0x05, 0xa9,
	0x50, 0x09, 0x29, 0xc1, 0xc0, 0xd2, 0xd4, 0xa2, 0x4a, 0x8f, 0xd4, 0xc4, 0x94, 0xd4, 0x22, 0xa8,
	0x90, 0x48, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0xb2, 0xa8, 0x52, 0x26, 0x97, 0x08,
	0x9a, 0xd1, 0x60, 0x9d, 0x42, 0x5a, 0x5c, 0x6c, 0x19, 0x60, 0x75, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x42, 0x10, 0xf5, 0x7a, 0x48, 0xe6, 0x06, 0x41, 0x55, 0x08, 0xa9, 0x72, 0xb1, 0xa5,
	0x65, 0xe6, 0xa4, 0x7a, 0xba, 0x48, 0x30, 0x81, 0xd5, 0xf2, 0x42, 0xd5, 0xba, 0x81, 0x05, 0x83,
	0xa0, 0x92, 0x4a, 0x6f, 0x18, 0xb9, 0xc4, 0xd1, 0xec, 0x82, 0x39, 0x49, 0x48, 0x17, 0xcd, 0x3a,
	0x51, 0xa8, 0x11, 0xa8, 0x6e, 0x86, 0xdb, 0xe8, 0xcf, 0xc5, 0x03, 0x32, 0x14, 0x66, 0x0c, 0xd4,
	0x5e, 0x6d, 0x24, 0x7b, 0xb1, 0x58, 0x02, 0x16, 0x87, 0x0b, 0xa2, 0x18, 0x20, 0x15, 0xc8, 0xc5,
	0x83, 0x2c, 0x8b, 0xe4, 0x25, 0x46, 0x3c, 0x5e, 0x12, 0x92, 0xe2, 0xe2, 0x48, 0x46, 0x76, 0x03,
	0x4f, 0x10, 0x9c, 0xef, 0x24, 0xc7, 0x25, 0x95, 0x9c, 0x9f, 0xab, 0x97, 0x91, 0x9a, 0x92, 0x5a,
	0x94, 0xa8, 0x97, 0x91, 0x58, 0x9c, 0x91, 0x5e, 0x94, 0x58, 0x90, 0x01, 0x31, 0x28, 0x80, 0x31,
	0x89, 0x0d, 0xcc, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xac, 0xc8, 0x9a, 0x81, 0xdb, 0x01,
	0x00, 0x00,
}