// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CryptoGetAccountRecords.proto

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

// Get all the records for an account for any transfers into it and out of it, that were above the threshold, during the last 25 hours.
type CryptoGetAccountRecordsQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountID            *AccountID   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CryptoGetAccountRecordsQuery) Reset()         { *m = CryptoGetAccountRecordsQuery{} }
func (m *CryptoGetAccountRecordsQuery) String() string { return proto.CompactTextString(m) }
func (*CryptoGetAccountRecordsQuery) ProtoMessage()    {}
func (*CryptoGetAccountRecordsQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d0010e1e19b708a, []int{0}
}

func (m *CryptoGetAccountRecordsQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetAccountRecordsQuery.Unmarshal(m, b)
}
func (m *CryptoGetAccountRecordsQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetAccountRecordsQuery.Marshal(b, m, deterministic)
}
func (m *CryptoGetAccountRecordsQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetAccountRecordsQuery.Merge(m, src)
}
func (m *CryptoGetAccountRecordsQuery) XXX_Size() int {
	return xxx_messageInfo_CryptoGetAccountRecordsQuery.Size(m)
}
func (m *CryptoGetAccountRecordsQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetAccountRecordsQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetAccountRecordsQuery proto.InternalMessageInfo

func (m *CryptoGetAccountRecordsQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetAccountRecordsQuery) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

// Response when the client sends the node CryptoGetAccountRecordsQuery
type CryptoGetAccountRecordsResponse struct {
	Header               *ResponseHeader      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountID            *AccountID           `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Records              []*TransactionRecord `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CryptoGetAccountRecordsResponse) Reset()         { *m = CryptoGetAccountRecordsResponse{} }
func (m *CryptoGetAccountRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*CryptoGetAccountRecordsResponse) ProtoMessage()    {}
func (*CryptoGetAccountRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d0010e1e19b708a, []int{1}
}

func (m *CryptoGetAccountRecordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetAccountRecordsResponse.Unmarshal(m, b)
}
func (m *CryptoGetAccountRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetAccountRecordsResponse.Marshal(b, m, deterministic)
}
func (m *CryptoGetAccountRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetAccountRecordsResponse.Merge(m, src)
}
func (m *CryptoGetAccountRecordsResponse) XXX_Size() int {
	return xxx_messageInfo_CryptoGetAccountRecordsResponse.Size(m)
}
func (m *CryptoGetAccountRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetAccountRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetAccountRecordsResponse proto.InternalMessageInfo

func (m *CryptoGetAccountRecordsResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetAccountRecordsResponse) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *CryptoGetAccountRecordsResponse) GetRecords() []*TransactionRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*CryptoGetAccountRecordsQuery)(nil), "proto.CryptoGetAccountRecordsQuery")
	proto.RegisterType((*CryptoGetAccountRecordsResponse)(nil), "proto.CryptoGetAccountRecordsResponse")
}

func init() { proto.RegisterFile("CryptoGetAccountRecords.proto", fileDescriptor_2d0010e1e19b708a) }

var fileDescriptor_2d0010e1e19b708a = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x75, 0x2e, 0xaa, 0x2c,
	0x28, 0xc9, 0x77, 0x4f, 0x2d, 0x71, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x09, 0x4a, 0x4d, 0xce,
	0x2f, 0x4a, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x02, 0x4e,
	0x89, 0xc5, 0x99, 0xc9, 0x21, 0x95, 0x05, 0xa9, 0x50, 0x09, 0x29, 0xf1, 0x90, 0xa2, 0xc4, 0xbc,
	0xe2, 0xc4, 0xe4, 0x92, 0xcc, 0xfc, 0x3c, 0x88, 0x16, 0xa8, 0x84, 0x60, 0x60, 0x69, 0x6a, 0x51,
	0xa5, 0x47, 0x6a, 0x62, 0x4a, 0x6a, 0x11, 0x54, 0x48, 0x24, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x15, 0x59, 0x54, 0xa9, 0x8a, 0x4b, 0x06, 0x87, 0xdd, 0x60, 0x13, 0x84, 0xb4, 0xb8, 0xd8,
	0x32, 0xc0, 0xea, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x84, 0x20, 0xfa, 0xf4, 0x90, 0xcc,
	0x0f, 0x82, 0xaa, 0x10, 0xd2, 0xe3, 0xe2, 0x4c, 0x84, 0x18, 0xe1, 0xe9, 0x22, 0xc1, 0x04, 0x56,
	0x2e, 0x00, 0x55, 0xee, 0x08, 0x13, 0x0f, 0x42, 0x28, 0x51, 0xda, 0xc2, 0xc8, 0x25, 0x8f, 0xc3,
	0x72, 0x98, 0x5b, 0x85, 0x74, 0xd1, 0xec, 0x17, 0x85, 0x1a, 0x88, 0xea, 0x19, 0x72, 0x9d, 0x20,
	0x64, 0xc4, 0xc5, 0x5e, 0x04, 0xb1, 0x51, 0x82, 0x59, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x02, 0xaa,
	0x1a, 0x23, 0x60, 0x83, 0x60, 0x0a, 0x9d, 0xe4, 0xb8, 0xa4, 0x92, 0xf3, 0x73, 0xf5, 0x32, 0x52,
	0x53, 0x52, 0x8b, 0x12, 0xf5, 0x32, 0x12, 0x8b, 0x33, 0xd2, 0x8b, 0x12, 0x0b, 0x32, 0x20, 0x1a,
	0x03, 0x18, 0x93, 0xd8, 0xc0, 0x0c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xb3, 0xdc,
	0x34, 0xd5, 0x01, 0x00, 0x00,
}
