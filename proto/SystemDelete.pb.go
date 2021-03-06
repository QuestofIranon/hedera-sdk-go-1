// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SystemDelete.proto

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

// Delete a file or smart contract - can only be done with a Hedera admin multisig. When it is deleted, it immediately disappears from the system as seen by the user, but is still stored internally until the expiration time, at which time it is truly and permanently deleted. Until that time, it can be undeleted by the Hedera admin multisig. When a smart contract is deleted, the cryptocurrency account within it continues to exist, and is not affected by the expiration time here.
type SystemDeleteTransactionBody struct {
	// Types that are valid to be assigned to Id:
	//	*SystemDeleteTransactionBody_FileID
	//	*SystemDeleteTransactionBody_ContractID
	Id                   isSystemDeleteTransactionBody_Id `protobuf_oneof:"id"`
	ExpirationTime       *TimestampSeconds                `protobuf:"bytes,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SystemDeleteTransactionBody) Reset()         { *m = SystemDeleteTransactionBody{} }
func (m *SystemDeleteTransactionBody) String() string { return proto.CompactTextString(m) }
func (*SystemDeleteTransactionBody) ProtoMessage()    {}
func (*SystemDeleteTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_860c88b39ffc5dcc, []int{0}
}

func (m *SystemDeleteTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemDeleteTransactionBody.Unmarshal(m, b)
}
func (m *SystemDeleteTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemDeleteTransactionBody.Marshal(b, m, deterministic)
}
func (m *SystemDeleteTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemDeleteTransactionBody.Merge(m, src)
}
func (m *SystemDeleteTransactionBody) XXX_Size() int {
	return xxx_messageInfo_SystemDeleteTransactionBody.Size(m)
}
func (m *SystemDeleteTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemDeleteTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_SystemDeleteTransactionBody proto.InternalMessageInfo

type isSystemDeleteTransactionBody_Id interface {
	isSystemDeleteTransactionBody_Id()
}

type SystemDeleteTransactionBody_FileID struct {
	FileID *FileID `protobuf:"bytes,1,opt,name=fileID,proto3,oneof"`
}

type SystemDeleteTransactionBody_ContractID struct {
	ContractID *ContractID `protobuf:"bytes,2,opt,name=contractID,proto3,oneof"`
}

func (*SystemDeleteTransactionBody_FileID) isSystemDeleteTransactionBody_Id() {}

func (*SystemDeleteTransactionBody_ContractID) isSystemDeleteTransactionBody_Id() {}

func (m *SystemDeleteTransactionBody) GetId() isSystemDeleteTransactionBody_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SystemDeleteTransactionBody) GetFileID() *FileID {
	if x, ok := m.GetId().(*SystemDeleteTransactionBody_FileID); ok {
		return x.FileID
	}
	return nil
}

func (m *SystemDeleteTransactionBody) GetContractID() *ContractID {
	if x, ok := m.GetId().(*SystemDeleteTransactionBody_ContractID); ok {
		return x.ContractID
	}
	return nil
}

func (m *SystemDeleteTransactionBody) GetExpirationTime() *TimestampSeconds {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SystemDeleteTransactionBody) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SystemDeleteTransactionBody_FileID)(nil),
		(*SystemDeleteTransactionBody_ContractID)(nil),
	}
}

func init() {
	proto.RegisterType((*SystemDeleteTransactionBody)(nil), "proto.SystemDeleteTransactionBody")
}

func init() { proto.RegisterFile("SystemDelete.proto", fileDescriptor_860c88b39ffc5dcc) }

var fileDescriptor_860c88b39ffc5dcc = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x3b, 0x55, 0xbb, 0xb8, 0xe2, 0x5f, 0x36, 0x0e, 0x23, 0x88, 0xb8, 0xd1, 0xd5, 0x2c,
	0xec, 0x03, 0x08, 0x71, 0x90, 0x76, 0x27, 0xd3, 0x79, 0x81, 0x6b, 0x72, 0x35, 0x81, 0xe6, 0x87,
	0x24, 0x0b, 0xe7, 0xf5, 0x7c, 0x32, 0x69, 0x32, 0x94, 0xc1, 0x55, 0xc2, 0x39, 0xdf, 0x39, 0xf7,
	0x00, 0xdb, 0x8d, 0x31, 0x91, 0xe9, 0x68, 0x4f, 0x89, 0x5a, 0x1f, 0x5c, 0x72, 0xec, 0x2c, 0x3f,
	0xcd, 0x35, 0xc7, 0xa8, 0xc5, 0x30, 0x7a, 0x8a, 0xc5, 0x68, 0xae, 0x06, 0x6d, 0x28, 0x26, 0x34,
	0xbe, 0x08, 0x8f, 0xbf, 0x15, 0xdc, 0xcd, 0x0b, 0x86, 0x80, 0x36, 0xa2, 0x48, 0xda, 0x59, 0xee,
	0xe4, 0xc8, 0x9e, 0x60, 0xf5, 0xa5, 0xf7, 0xb4, 0xed, 0xea, 0xea, 0xa1, 0x7a, 0x3e, 0x7f, 0xb9,
	0x28, 0xb9, 0xf6, 0x3d, 0x8b, 0x9b, 0x45, 0x3f, 0xd9, 0x6c, 0x0d, 0x20, 0x9c, 0x4d, 0x01, 0x45,
	0xda, 0x76, 0xf5, 0x32, 0xc3, 0x37, 0x13, 0xfc, 0x76, 0x34, 0x36, 0x8b, 0x7e, 0x86, 0xb1, 0x57,
	0xb8, 0xa4, 0x1f, 0xaf, 0x03, 0x1e, 0xee, 0x1d, 0xa6, 0xd5, 0x27, 0x39, 0x78, 0x3b, 0x05, 0x8f,
	0x6b, 0x77, 0x24, 0x9c, 0x95, 0xb1, 0xff, 0x87, 0xf3, 0x53, 0x58, 0x6a, 0xc9, 0xef, 0xa1, 0x11,
	0xce, 0xb4, 0x8a, 0x24, 0x05, 0x6c, 0x15, 0x46, 0xf5, 0x1d, 0xd0, 0xab, 0x52, 0xf2, 0x51, 0x7d,
	0xae, 0xf2, 0x67, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x15, 0x44, 0xad, 0x96, 0x2b, 0x01, 0x00,
	0x00,
}
