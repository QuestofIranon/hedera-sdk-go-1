// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CryptoGetStakers.proto

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
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xd1, 0x4a, 0xfb, 0x30,
	0x14, 0xc6, 0xe9, 0xca, 0xbf, 0x7f, 0x3c, 0xbd, 0x70, 0x06, 0x57, 0x42, 0x2f, 0x64, 0xf4, 0x4a,
	0x04, 0x83, 0xa8, 0x2f, 0xb0, 0x29, 0xa8, 0x77, 0x33, 0xea, 0x03, 0xc4, 0xee, 0x60, 0xc4, 0x6d,
	0x09, 0x49, 0x86, 0x16, 0x5f, 0x5e, 0x4c, 0x53, 0x97, 0xf6, 0x46, 0xbc, 0xea, 0xe9, 0xc9, 0xef,
	0xfb, 0x4e, 0xce, 0x17, 0x28, 0xae, 0x4c, 0xa3, 0x9d, 0xba, 0x41, 0xf7, 0xe0, 0xc4, 0x1b, 0x1a,
	0xcb, 0xb4, 0x51, 0x4e, 0x91, 0x7f, 0xfe, 0x53, 0x8e, 0xe7, 0xc2, 0xbe, 0xd6, 0x8f, 0x8d, 0xc6,
	0x70, 0x50, 0x1e, 0xdc, 0x6f, 0xd1, 0x34, 0xb7, 0x28, 0x96, 0x68, 0x42, 0xeb, 0x90, 0xa3, 0xd5,
	0x6a, 0x63, 0x31, 0xee, 0x56, 0x16, 0x26, 0x43, 0x6f, 0x2f, 0x25, 0x27, 0x90, 0x49, 0x0f, 0xd2,
	0x64, 0x9a, 0x1c, 0xe7, 0xe7, 0xa4, 0x15, 0xb0, 0xc8, 0x98, 0x07, 0x82, 0x30, 0xd8, 0x13, 0x75,
	0xad, 0xb6, 0x1b, 0x77, 0x77, 0x4d, 0x47, 0x1e, 0x1f, 0x07, 0x7c, 0xd6, 0xf5, 0xf9, 0x0e, 0xa9,
	0x9e, 0x20, 0x5f, 0x18, 0xf5, 0xd1, 0xb4, 0x03, 0xfb, 0xf2, 0xe4, 0x57, 0x39, 0x29, 0x20, 0x13,
	0xeb, 0xef, 0xda, 0xcf, 0x4a, 0x79, 0xf8, 0xab, 0xde, 0x61, 0x7f, 0xb6, 0x5a, 0x45, 0xce, 0xf6,
	0xcf, 0xd6, 0x97, 0x90, 0xeb, 0x9d, 0x9e, 0x8e, 0xa6, 0x69, 0xb4, 0x7a, 0xe4, 0xcc, 0x63, 0xac,
	0xfa, 0x04, 0x3a, 0x0c, 0xb1, 0x0b, 0x9b, 0x9c, 0x0e, 0x72, 0x9c, 0x04, 0xb3, 0xfe, 0x6b, 0xfc,
	0x44, 0x79, 0x06, 0xff, 0x6d, 0xeb, 0x40, 0x53, 0xcf, 0x17, 0xdd, 0x75, 0xfb, 0x9b, 0xf1, 0x0e,
	0x9b, 0x1f, 0x41, 0x59, 0xab, 0x35, 0x93, 0xb8, 0x44, 0x23, 0x98, 0x14, 0x56, 0xbe, 0x18, 0xa1,
	0x65, 0x2b, 0x5b, 0x24, 0xcf, 0x99, 0x2f, 0x2e, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x35, 0xd5,
	0x22, 0x05, 0x44, 0x02, 0x00, 0x00,
}
