// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CryptoGetStakers.proto

package hedera_proto

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

// Get all the accounts that are proxy staking to this account. For each of them, give the amount currently staked. This is not yet implemented, but will be in a future version of the API.
type CryptoGetStakersQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountID            *AccountID   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CryptoGetStakersQuery) Reset()         { *m = CryptoGetStakersQuery{} }
func (m *CryptoGetStakersQuery) String() string { return proto.CompactTextString(m) }
func (*CryptoGetStakersQuery) ProtoMessage()    {}
func (*CryptoGetStakersQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af165a1b4c34d1f, []int{0}
}

func (m *CryptoGetStakersQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetStakersQuery.Unmarshal(m, b)
}
func (m *CryptoGetStakersQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetStakersQuery.Marshal(b, m, deterministic)
}
func (m *CryptoGetStakersQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetStakersQuery.Merge(m, src)
}
func (m *CryptoGetStakersQuery) XXX_Size() int {
	return xxx_messageInfo_CryptoGetStakersQuery.Size(m)
}
func (m *CryptoGetStakersQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetStakersQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetStakersQuery proto.InternalMessageInfo

func (m *CryptoGetStakersQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetStakersQuery) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

// information about a single account that is proxy staking
type ProxyStaker struct {
	AccountID            *AccountID `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Amount               int64      `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProxyStaker) Reset()         { *m = ProxyStaker{} }
func (m *ProxyStaker) String() string { return proto.CompactTextString(m) }
func (*ProxyStaker) ProtoMessage()    {}
func (*ProxyStaker) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af165a1b4c34d1f, []int{1}
}

func (m *ProxyStaker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyStaker.Unmarshal(m, b)
}
func (m *ProxyStaker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyStaker.Marshal(b, m, deterministic)
}
func (m *ProxyStaker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyStaker.Merge(m, src)
}
func (m *ProxyStaker) XXX_Size() int {
	return xxx_messageInfo_ProxyStaker.Size(m)
}
func (m *ProxyStaker) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyStaker.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyStaker proto.InternalMessageInfo

func (m *ProxyStaker) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *ProxyStaker) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// all of the accounts proxy staking to a given account, and the amounts proxy staked
type AllProxyStakers struct {
	AccountID            *AccountID     `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	ProxyStaker          []*ProxyStaker `protobuf:"bytes,2,rep,name=proxyStaker,proto3" json:"proxyStaker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AllProxyStakers) Reset()         { *m = AllProxyStakers{} }
func (m *AllProxyStakers) String() string { return proto.CompactTextString(m) }
func (*AllProxyStakers) ProtoMessage()    {}
func (*AllProxyStakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af165a1b4c34d1f, []int{2}
}

func (m *AllProxyStakers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllProxyStakers.Unmarshal(m, b)
}
func (m *AllProxyStakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllProxyStakers.Marshal(b, m, deterministic)
}
func (m *AllProxyStakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllProxyStakers.Merge(m, src)
}
func (m *AllProxyStakers) XXX_Size() int {
	return xxx_messageInfo_AllProxyStakers.Size(m)
}
func (m *AllProxyStakers) XXX_DiscardUnknown() {
	xxx_messageInfo_AllProxyStakers.DiscardUnknown(m)
}

var xxx_messageInfo_AllProxyStakers proto.InternalMessageInfo

func (m *AllProxyStakers) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *AllProxyStakers) GetProxyStaker() []*ProxyStaker {
	if m != nil {
		return m.ProxyStaker
	}
	return nil
}

// Response when the client sends the node CryptoGetStakersQuery
type CryptoGetStakersResponse struct {
	Header               *ResponseHeader  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Stakers              *AllProxyStakers `protobuf:"bytes,3,opt,name=stakers,proto3" json:"stakers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CryptoGetStakersResponse) Reset()         { *m = CryptoGetStakersResponse{} }
func (m *CryptoGetStakersResponse) String() string { return proto.CompactTextString(m) }
func (*CryptoGetStakersResponse) ProtoMessage()    {}
func (*CryptoGetStakersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af165a1b4c34d1f, []int{3}
}

func (m *CryptoGetStakersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetStakersResponse.Unmarshal(m, b)
}
func (m *CryptoGetStakersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetStakersResponse.Marshal(b, m, deterministic)
}
func (m *CryptoGetStakersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetStakersResponse.Merge(m, src)
}
func (m *CryptoGetStakersResponse) XXX_Size() int {
	return xxx_messageInfo_CryptoGetStakersResponse.Size(m)
}
func (m *CryptoGetStakersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetStakersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetStakersResponse proto.InternalMessageInfo

func (m *CryptoGetStakersResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetStakersResponse) GetStakers() *AllProxyStakers {
	if m != nil {
		return m.Stakers
	}
	return nil
}

func init() {
	proto.RegisterType((*CryptoGetStakersQuery)(nil), "proto.CryptoGetStakersQuery")
	proto.RegisterType((*ProxyStaker)(nil), "proto.ProxyStaker")
	proto.RegisterType((*AllProxyStakers)(nil), "proto.AllProxyStakers")
	proto.RegisterType((*CryptoGetStakersResponse)(nil), "proto.CryptoGetStakersResponse")
}

func init() { proto.RegisterFile("CryptoGetStakers.proto", fileDescriptor_4af165a1b4c34d1f) }

var fileDescriptor_4af165a1b4c34d1f = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xe9, 0xca, 0xbf, 0x7f, 0xbc, 0x15, 0x9c, 0xc1, 0x95, 0xb2, 0xa7, 0xd1, 0x27, 0x11,
	0x2c, 0xa2, 0xfb, 0x02, 0x9b, 0x82, 0xfa, 0x36, 0xa3, 0xbe, 0xf8, 0x22, 0xd7, 0xee, 0x62, 0xa7,
	0xdb, 0x12, 0x92, 0x4c, 0x2d, 0x7e, 0x79, 0x31, 0x4d, 0x5d, 0xda, 0x17, 0xf1, 0x29, 0xed, 0xe9,
	0xef, 0x9c, 0xdb, 0x7b, 0x02, 0xc9, 0xb9, 0xaa, 0xa4, 0x11, 0x97, 0x64, 0x6e, 0x0d, 0xbe, 0x92,
	0xd2, 0xb9, 0x54, 0xc2, 0x08, 0xf6, 0xcf, 0x1e, 0xc3, 0xfe, 0x14, 0xf5, 0xa2, 0xb8, 0xab, 0x24,
	0xb9, 0x0f, 0xc3, 0xfd, 0x9b, 0x0d, 0xa9, 0xea, 0x8a, 0x70, 0x4e, 0xca, 0x49, 0x07, 0x9c, 0xb4,
	0x14, 0x6b, 0x4d, 0xbe, 0x9a, 0x69, 0x18, 0x74, 0xb3, 0xad, 0x95, 0x1d, 0x41, 0x54, 0x5a, 0x30,
	0x0d, 0x46, 0xc1, 0x61, 0x7c, 0xca, 0x6a, 0x43, 0xee, 0x05, 0x73, 0x47, 0xb0, 0x1c, 0x76, 0xb0,
	0x28, 0xc4, 0x66, 0x6d, 0xae, 0x2f, 0xd2, 0x9e, 0xc5, 0xfb, 0x0e, 0x9f, 0x34, 0x3a, 0xdf, 0x22,
	0xd9, 0x3d, 0xc4, 0x33, 0x25, 0x3e, 0xaa, 0x7a, 0x60, 0xdb, 0x1e, 0xfc, 0x6a, 0x67, 0x09, 0x44,
	0xb8, 0xfa, 0x7e, 0xb6, 0xb3, 0x42, 0xee, 0xde, 0xb2, 0x77, 0xd8, 0x9b, 0x2c, 0x97, 0x5e, 0xb2,
	0xfe, 0x73, 0xf4, 0x18, 0x62, 0xb9, 0xf5, 0xa7, 0xbd, 0x51, 0xe8, 0xad, 0xee, 0x25, 0x73, 0x1f,
	0xcb, 0x3e, 0x21, 0xed, 0x96, 0xd8, 0x94, 0xcd, 0x8e, 0x3b, 0x3d, 0x0e, 0x5c, 0x58, 0xfb, 0x36,
	0x7e, 0xaa, 0x3c, 0x81, 0xff, 0xba, 0x4e, 0x48, 0x43, 0xcb, 0x27, 0xcd, 0xef, 0xb6, 0x37, 0xe3,
	0x0d, 0x36, 0x1d, 0x43, 0x56, 0x88, 0x55, 0x5e, 0xd2, 0x9c, 0x14, 0x96, 0xa8, 0xcb, 0x67, 0x85,
	0xb2, 0xcc, 0x51, 0x2e, 0x9c, 0xf3, 0x05, 0xdf, 0x70, 0x16, 0x3c, 0xec, 0xd6, 0xc4, 0xa3, 0x15,
	0x9f, 0x22, 0x7b, 0x9c, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xa0, 0x58, 0x7c, 0x5a, 0x02,
	0x00, 0x00,
}