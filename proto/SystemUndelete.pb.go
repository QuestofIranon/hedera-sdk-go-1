// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SystemUndelete.proto

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

// Undelete a file or smart contract that was deleted by AdminDelete - can only be done with a Hedera admin multisig. When it is deleted, it immediately disappears from the system as seen by the user, but is still stored internally until the expiration time, at which time it is truly and permanently deleted. Until that time, it can be undeleted by the Hedera admin multisig. When a smart contract is deleted, the cryptocurrency account within it continues to exist, and is not affected by the expiration time here.
type SystemUndeleteTransactionBody struct {
	// Types that are valid to be assigned to Id:
	//	*SystemUndeleteTransactionBody_FileID
	//	*SystemUndeleteTransactionBody_ContractID
	Id                   isSystemUndeleteTransactionBody_Id `protobuf_oneof:"id"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *SystemUndeleteTransactionBody) Reset()         { *m = SystemUndeleteTransactionBody{} }
func (m *SystemUndeleteTransactionBody) String() string { return proto.CompactTextString(m) }
func (*SystemUndeleteTransactionBody) ProtoMessage()    {}
func (*SystemUndeleteTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba44bcfbfa1b6abb, []int{0}
}

func (m *SystemUndeleteTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemUndeleteTransactionBody.Unmarshal(m, b)
}
func (m *SystemUndeleteTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemUndeleteTransactionBody.Marshal(b, m, deterministic)
}
func (m *SystemUndeleteTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemUndeleteTransactionBody.Merge(m, src)
}
func (m *SystemUndeleteTransactionBody) XXX_Size() int {
	return xxx_messageInfo_SystemUndeleteTransactionBody.Size(m)
}
func (m *SystemUndeleteTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemUndeleteTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_SystemUndeleteTransactionBody proto.InternalMessageInfo

type isSystemUndeleteTransactionBody_Id interface {
	isSystemUndeleteTransactionBody_Id()
}

type SystemUndeleteTransactionBody_FileID struct {
	FileID *FileID `protobuf:"bytes,1,opt,name=fileID,proto3,oneof"`
}

type SystemUndeleteTransactionBody_ContractID struct {
	ContractID *ContractID `protobuf:"bytes,2,opt,name=contractID,proto3,oneof"`
}

func (*SystemUndeleteTransactionBody_FileID) isSystemUndeleteTransactionBody_Id() {}

func (*SystemUndeleteTransactionBody_ContractID) isSystemUndeleteTransactionBody_Id() {}

func (m *SystemUndeleteTransactionBody) GetId() isSystemUndeleteTransactionBody_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SystemUndeleteTransactionBody) GetFileID() *FileID {
	if x, ok := m.GetId().(*SystemUndeleteTransactionBody_FileID); ok {
		return x.FileID
	}
	return nil
}

func (m *SystemUndeleteTransactionBody) GetContractID() *ContractID {
	if x, ok := m.GetId().(*SystemUndeleteTransactionBody_ContractID); ok {
		return x.ContractID
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SystemUndeleteTransactionBody) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SystemUndeleteTransactionBody_FileID)(nil),
		(*SystemUndeleteTransactionBody_ContractID)(nil),
	}
}

func init() {
	proto.RegisterType((*SystemUndeleteTransactionBody)(nil), "proto.SystemUndeleteTransactionBody")
}

func init() { proto.RegisterFile("SystemUndelete.proto", fileDescriptor_ba44bcfbfa1b6abb) }

var fileDescriptor_ba44bcfbfa1b6abb = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x09, 0xae, 0x2c, 0x2e,
	0x49, 0xcd, 0x0d, 0xcd, 0x4b, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x53, 0x52, 0x02, 0x4e, 0x89, 0xc5, 0x99, 0xc9, 0x21, 0x95, 0x05, 0xa9, 0xc5,
	0x10, 0x09, 0xa5, 0x66, 0x46, 0x2e, 0x59, 0x54, 0x1d, 0x21, 0x45, 0x89, 0x79, 0xc5, 0x89, 0xc9,
	0x25, 0x99, 0xf9, 0x79, 0x4e, 0xf9, 0x29, 0x95, 0x42, 0xea, 0x5c, 0x6c, 0x69, 0x99, 0x39, 0xa9,
	0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xbc, 0x10, 0x9d, 0x7a, 0x6e, 0x60, 0x41,
	0x0f, 0x86, 0x20, 0xa8, 0xb4, 0x90, 0x31, 0x17, 0x57, 0x72, 0x7e, 0x5e, 0x49, 0x51, 0x62, 0x72,
	0x89, 0xa7, 0x8b, 0x04, 0x13, 0x58, 0xb1, 0x20, 0x54, 0xb1, 0x33, 0x5c, 0xc2, 0x83, 0x21, 0x08,
	0x49, 0x99, 0x13, 0x0b, 0x17, 0x53, 0x66, 0x8a, 0x93, 0x1c, 0x97, 0x54, 0x72, 0x7e, 0xae, 0x5e,
	0x46, 0x6a, 0x4a, 0x6a, 0x51, 0xa2, 0x5e, 0x46, 0x62, 0x71, 0x46, 0x7a, 0x51, 0x62, 0x41, 0x06,
	0x44, 0x73, 0x00, 0x63, 0x12, 0x1b, 0x98, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x91, 0x05,
	0xdd, 0x12, 0xdd, 0x00, 0x00, 0x00,
}
