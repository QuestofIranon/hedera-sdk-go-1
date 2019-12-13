// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ConsensusTopicInfo.proto

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

// Current state of a topic.
type ConsensusTopicInfo struct {
	Memo string `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
	// SHA-384 running hash <previousRunningHash, topicId, consensusTimestamp, message>
	RunningHash []byte `protobuf:"bytes,2,opt,name=runningHash,proto3" json:"runningHash,omitempty"`
	// Sequence number (starting at 1 for the first submitMessage) of messages on the topic.
	SequenceNumber uint64 `protobuf:"varint,3,opt,name=sequenceNumber,proto3" json:"sequenceNumber,omitempty"`
	// Effective consensus timestamp at (and after) which submitMessage calls will no longer succeed on the topic
	// and the topic will expire and be marked as deleted.
	ExpirationTime       *Timestamp `protobuf:"bytes,4,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	AdminKey             *Key       `protobuf:"bytes,5,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	SubmitKey            *Key       `protobuf:"bytes,6,opt,name=submitKey,proto3" json:"submitKey,omitempty"`
	AutoRenewPeriod      *Duration  `protobuf:"bytes,7,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	AutoRenewAccount     *AccountID `protobuf:"bytes,8,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConsensusTopicInfo) Reset()         { *m = ConsensusTopicInfo{} }
func (m *ConsensusTopicInfo) String() string { return proto.CompactTextString(m) }
func (*ConsensusTopicInfo) ProtoMessage()    {}
func (*ConsensusTopicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c77abb18a291246, []int{0}
}

func (m *ConsensusTopicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusTopicInfo.Unmarshal(m, b)
}
func (m *ConsensusTopicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusTopicInfo.Marshal(b, m, deterministic)
}
func (m *ConsensusTopicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusTopicInfo.Merge(m, src)
}
func (m *ConsensusTopicInfo) XXX_Size() int {
	return xxx_messageInfo_ConsensusTopicInfo.Size(m)
}
func (m *ConsensusTopicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusTopicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusTopicInfo proto.InternalMessageInfo

func (m *ConsensusTopicInfo) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *ConsensusTopicInfo) GetRunningHash() []byte {
	if m != nil {
		return m.RunningHash
	}
	return nil
}

func (m *ConsensusTopicInfo) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *ConsensusTopicInfo) GetExpirationTime() *Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *ConsensusTopicInfo) GetAdminKey() *Key {
	if m != nil {
		return m.AdminKey
	}
	return nil
}

func (m *ConsensusTopicInfo) GetSubmitKey() *Key {
	if m != nil {
		return m.SubmitKey
	}
	return nil
}

func (m *ConsensusTopicInfo) GetAutoRenewPeriod() *Duration {
	if m != nil {
		return m.AutoRenewPeriod
	}
	return nil
}

func (m *ConsensusTopicInfo) GetAutoRenewAccount() *AccountID {
	if m != nil {
		return m.AutoRenewAccount
	}
	return nil
}

func init() {
	proto.RegisterType((*ConsensusTopicInfo)(nil), "proto.ConsensusTopicInfo")
}

func init() { proto.RegisterFile("ConsensusTopicInfo.proto", fileDescriptor_2c77abb18a291246) }

var fileDescriptor_2c77abb18a291246 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4e, 0x02, 0x31,
	0x14, 0x85, 0x53, 0x04, 0x84, 0x8b, 0x01, 0xd2, 0x55, 0xc3, 0xc2, 0x4c, 0x5c, 0x90, 0x59, 0xcd,
	0x42, 0x37, 0x9a, 0xb8, 0x11, 0x59, 0x48, 0x48, 0x0c, 0x69, 0x78, 0x81, 0x52, 0xae, 0x4c, 0x17,
	0xfd, 0xb1, 0x3f, 0x51, 0x9e, 0xdc, 0xad, 0x61, 0x66, 0x44, 0x85, 0x55, 0x6f, 0xcf, 0xf9, 0x4e,
	0x73, 0x7a, 0x81, 0x3d, 0x5b, 0x13, 0xd0, 0x84, 0x14, 0xd6, 0xd6, 0x29, 0xb9, 0x30, 0x6f, 0xb6,
	0x70, 0xde, 0x46, 0x4b, 0x3b, 0xd5, 0x31, 0x19, 0xcf, 0x44, 0x50, 0x72, 0xbd, 0x77, 0x18, 0x6a,
	0x63, 0x32, 0x9c, 0x27, 0x2f, 0xa2, 0xb2, 0xa6, 0xb9, 0x8f, 0xd6, 0x4a, 0x63, 0x88, 0x42, 0xbb,
	0x5a, 0xb8, 0xf9, 0x6a, 0x01, 0x3d, 0x7f, 0x96, 0x52, 0x68, 0x6b, 0xd4, 0x96, 0x91, 0x8c, 0xe4,
	0x7d, 0x5e, 0xcd, 0x34, 0x83, 0x81, 0x4f, 0xc6, 0x28, 0xb3, 0x7b, 0x11, 0xa1, 0x64, 0xad, 0x8c,
	0xe4, 0x57, 0xfc, 0xaf, 0x44, 0xa7, 0x30, 0x0c, 0xf8, 0x9e, 0xd0, 0x48, 0x7c, 0x4d, 0x7a, 0x83,
	0x9e, 0x5d, 0x64, 0x24, 0x6f, 0xf3, 0x13, 0x95, 0xde, 0xc3, 0x10, 0x3f, 0x9d, 0xaa, 0x9b, 0x1d,
	0x1a, 0xb1, 0x76, 0x46, 0xf2, 0xc1, 0xed, 0xb8, 0x2e, 0x55, 0x1c, 0x4b, 0xf2, 0x13, 0x8e, 0x4e,
	0xa1, 0x27, 0xb6, 0x5a, 0x99, 0x25, 0xee, 0x59, 0xa7, 0xca, 0x40, 0x93, 0x59, 0xe2, 0x9e, 0x1f,
	0x3d, 0x9a, 0x43, 0x3f, 0xa4, 0x8d, 0x56, 0xf1, 0x00, 0x76, 0xcf, 0xc0, 0x5f, 0x93, 0x3e, 0xc0,
	0x48, 0xa4, 0x68, 0x39, 0x1a, 0xfc, 0x58, 0xa1, 0x57, 0x76, 0xcb, 0x2e, 0x2b, 0x7e, 0xd4, 0xf0,
	0x3f, 0x1b, 0xe4, 0xa7, 0x1c, 0x7d, 0x84, 0xf1, 0x51, 0x7a, 0x92, 0xd2, 0x26, 0x13, 0x59, 0xef,
	0xdf, 0x47, 0x1a, 0x75, 0x31, 0xe7, 0x67, 0xe4, 0xec, 0x1a, 0x26, 0xd2, 0xea, 0xa2, 0xc4, 0x2d,
	0x7a, 0x51, 0x94, 0x22, 0x94, 0x3b, 0x2f, 0x5c, 0x59, 0x27, 0x57, 0x64, 0xd3, 0xad, 0x86, 0xbb,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x89, 0xee, 0x5b, 0xf6, 0x01, 0x00, 0x00,
}