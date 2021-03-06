// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ContractDelete.proto

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

// Modify a smart contract instance to have the given parameter values. Any null field is ignored (left unchanged). If only the contractInstanceExpirationTime is being modified, then no signature is needed on this transaction other than for the account paying for the transaction itself. But if any of the other fields are being modified, then it must be signed by the adminKey. The use of adminKey is not currently supported in this API, but in the future will be implemented to allow these fields to be modified, and also to make modifications to the state of the instance. If the contract is created with no admin key, then none of the fields can be changed that need an admin signature, and therefore no admin key can ever be added. So if there is no admin key, then things like the bytecode are immutable. But if there is an admin key, then they can be changed. For example, the admin key might be a threshold key, which requires 3 of 5 binding arbitration judges to agree before the bytecode can be changed. This can be used to add flexibility to the mangement of smart contract behavior. But this is optional. If the smart contract is created without an admin key, then such a key can never be added, and its bytecode will be immutable.
type ContractDeleteTransactionBody struct {
	ContractID *ContractID `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	// Types that are valid to be assigned to Obtainers:
	//	*ContractDeleteTransactionBody_TransferAccountID
	//	*ContractDeleteTransactionBody_TransferContractID
	Obtainers            isContractDeleteTransactionBody_Obtainers `protobuf_oneof:"obtainers"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *ContractDeleteTransactionBody) Reset()         { *m = ContractDeleteTransactionBody{} }
func (m *ContractDeleteTransactionBody) String() string { return proto.CompactTextString(m) }
func (*ContractDeleteTransactionBody) ProtoMessage()    {}
func (*ContractDeleteTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ef03593b4c31413, []int{0}
}

func (m *ContractDeleteTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractDeleteTransactionBody.Unmarshal(m, b)
}
func (m *ContractDeleteTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractDeleteTransactionBody.Marshal(b, m, deterministic)
}
func (m *ContractDeleteTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractDeleteTransactionBody.Merge(m, src)
}
func (m *ContractDeleteTransactionBody) XXX_Size() int {
	return xxx_messageInfo_ContractDeleteTransactionBody.Size(m)
}
func (m *ContractDeleteTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractDeleteTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_ContractDeleteTransactionBody proto.InternalMessageInfo

func (m *ContractDeleteTransactionBody) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

type isContractDeleteTransactionBody_Obtainers interface {
	isContractDeleteTransactionBody_Obtainers()
}

type ContractDeleteTransactionBody_TransferAccountID struct {
	TransferAccountID *AccountID `protobuf:"bytes,2,opt,name=transferAccountID,proto3,oneof"`
}

type ContractDeleteTransactionBody_TransferContractID struct {
	TransferContractID *ContractID `protobuf:"bytes,3,opt,name=transferContractID,proto3,oneof"`
}

func (*ContractDeleteTransactionBody_TransferAccountID) isContractDeleteTransactionBody_Obtainers() {}

func (*ContractDeleteTransactionBody_TransferContractID) isContractDeleteTransactionBody_Obtainers() {}

func (m *ContractDeleteTransactionBody) GetObtainers() isContractDeleteTransactionBody_Obtainers {
	if m != nil {
		return m.Obtainers
	}
	return nil
}

func (m *ContractDeleteTransactionBody) GetTransferAccountID() *AccountID {
	if x, ok := m.GetObtainers().(*ContractDeleteTransactionBody_TransferAccountID); ok {
		return x.TransferAccountID
	}
	return nil
}

func (m *ContractDeleteTransactionBody) GetTransferContractID() *ContractID {
	if x, ok := m.GetObtainers().(*ContractDeleteTransactionBody_TransferContractID); ok {
		return x.TransferContractID
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ContractDeleteTransactionBody) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ContractDeleteTransactionBody_TransferAccountID)(nil),
		(*ContractDeleteTransactionBody_TransferContractID)(nil),
	}
}

func init() {
	proto.RegisterType((*ContractDeleteTransactionBody)(nil), "proto.ContractDeleteTransactionBody")
}

func init() { proto.RegisterFile("ContractDelete.proto", fileDescriptor_4ef03593b4c31413) }

var fileDescriptor_4ef03593b4c31413 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x71, 0xce, 0xcf, 0x2b,
	0x29, 0x4a, 0x4c, 0x2e, 0x71, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x53, 0x52, 0x02, 0x4e, 0x89, 0xc5, 0x99, 0xc9, 0x21, 0x95, 0x05, 0xa9, 0xc5,
	0x10, 0x09, 0xa5, 0x67, 0x8c, 0x5c, 0xb2, 0xa8, 0x3a, 0x42, 0x8a, 0x12, 0xf3, 0x8a, 0x13, 0x93,
	0x4b, 0x32, 0xf3, 0xf3, 0x9c, 0xf2, 0x53, 0x2a, 0x85, 0x0c, 0xb9, 0xb8, 0x92, 0xa1, 0x0a, 0x3c,
	0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x04, 0x21, 0xba, 0xf5, 0x9c, 0xe1, 0x12, 0x41,
	0x48, 0x8a, 0x84, 0x1c, 0xb8, 0x04, 0x4b, 0x40, 0xa6, 0xa4, 0xa5, 0x16, 0x39, 0x26, 0x27, 0xe7,
	0x97, 0xe6, 0x81, 0x74, 0x32, 0x81, 0x75, 0x0a, 0x40, 0x75, 0xc2, 0xc5, 0x3d, 0x18, 0x82, 0x30,
	0x15, 0x0b, 0x39, 0x73, 0x09, 0xc1, 0x04, 0x11, 0x76, 0x48, 0x30, 0xe3, 0xb0, 0xdc, 0x83, 0x21,
	0x08, 0x8b, 0x72, 0x27, 0x6e, 0x2e, 0xce, 0xfc, 0xa4, 0x92, 0xc4, 0xcc, 0xbc, 0xd4, 0xa2, 0x62,
	0x27, 0x39, 0x2e, 0xa9, 0xe4, 0xfc, 0x5c, 0xbd, 0x8c, 0xd4, 0x94, 0xd4, 0xa2, 0x44, 0xbd, 0x8c,
	0xc4, 0xe2, 0x8c, 0xf4, 0xa2, 0xc4, 0x82, 0x0c, 0x88, 0x59, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0x86,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x29, 0x15, 0x7f, 0x38, 0x40, 0x01, 0x00, 0x00,
}
