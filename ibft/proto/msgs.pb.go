// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bloxapp/ssv/ibft/proto/msgs.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type RoundState int32

const (
	RoundState_NotStarted  RoundState = 0
	RoundState_PrePrepare  RoundState = 1
	RoundState_Prepare     RoundState = 2
	RoundState_Commit      RoundState = 3
	RoundState_ChangeRound RoundState = 4
	RoundState_Decided     RoundState = 5
	RoundState_Stopped     RoundState = 6
)

var RoundState_name = map[int32]string{
	0: "NotStarted",
	1: "PrePrepare",
	2: "Prepare",
	3: "Commit",
	4: "ChangeRound",
	5: "Decided",
	6: "Stopped",
}

var RoundState_value = map[string]int32{
	"NotStarted":  0,
	"PrePrepare":  1,
	"Prepare":     2,
	"Commit":      3,
	"ChangeRound": 4,
	"Decided":     5,
	"Stopped":     6,
}

func (x RoundState) String() string {
	return proto.EnumName(RoundState_name, int32(x))
}

func (RoundState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dc1187fce89a5e11, []int{0}
}

type Message struct {
	Type   RoundState `protobuf:"varint,1,opt,name=type,proto3,enum=proto.RoundState" json:"type,omitempty"`
	Round  uint64     `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Lambda []byte     `protobuf:"bytes,3,opt,name=lambda,proto3" json:"lambda,omitempty"`
	// sequence number is an incremental number for each instance, much like a block number would be in a blockchain
	SeqNumber            uint64   `protobuf:"varint,4,opt,name=seq_number,json=seqNumber,proto3" json:"seq_number,omitempty"`
	PreviousLambda       []byte   `protobuf:"bytes,5,opt,name=previous_lambda,json=previousLambda,proto3" json:"previous_lambda,omitempty"`
	Value                []byte   `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	ValidatorPk          []byte   `protobuf:"bytes,7,opt,name=validator_pk,json=validatorPk,proto3" json:"validator_pk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1187fce89a5e11, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() RoundState {
	if m != nil {
		return m.Type
	}
	return RoundState_NotStarted
}

func (m *Message) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *Message) GetLambda() []byte {
	if m != nil {
		return m.Lambda
	}
	return nil
}

func (m *Message) GetSeqNumber() uint64 {
	if m != nil {
		return m.SeqNumber
	}
	return 0
}

func (m *Message) GetPreviousLambda() []byte {
	if m != nil {
		return m.PreviousLambda
	}
	return nil
}

func (m *Message) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Message) GetValidatorPk() []byte {
	if m != nil {
		return m.ValidatorPk
	}
	return nil
}

type SignedMessage struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	SignerIds            []uint64 `protobuf:"varint,3,rep,packed,name=signer_ids,json=signerIds,proto3" json:"signer_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedMessage) Reset()         { *m = SignedMessage{} }
func (m *SignedMessage) String() string { return proto.CompactTextString(m) }
func (*SignedMessage) ProtoMessage()    {}
func (*SignedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1187fce89a5e11, []int{1}
}

func (m *SignedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedMessage.Unmarshal(m, b)
}
func (m *SignedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedMessage.Marshal(b, m, deterministic)
}
func (m *SignedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedMessage.Merge(m, src)
}
func (m *SignedMessage) XXX_Size() int {
	return xxx_messageInfo_SignedMessage.Size(m)
}
func (m *SignedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SignedMessage proto.InternalMessageInfo

func (m *SignedMessage) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SignedMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedMessage) GetSignerIds() []uint64 {
	if m != nil {
		return m.SignerIds
	}
	return nil
}

type ChangeRoundData struct {
	PreparedRound        uint64   `protobuf:"varint,1,opt,name=prepared_round,json=preparedRound,proto3" json:"prepared_round,omitempty"`
	PreparedValue        []byte   `protobuf:"bytes,2,opt,name=prepared_value,json=preparedValue,proto3" json:"prepared_value,omitempty"`
	JustificationMsg     *Message `protobuf:"bytes,3,opt,name=justification_msg,json=justificationMsg,proto3" json:"justification_msg,omitempty"`
	JustificationSig     []byte   `protobuf:"bytes,4,opt,name=justification_sig,json=justificationSig,proto3" json:"justification_sig,omitempty"`
	SignerIds            []uint64 `protobuf:"varint,5,rep,packed,name=signer_ids,json=signerIds,proto3" json:"signer_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeRoundData) Reset()         { *m = ChangeRoundData{} }
func (m *ChangeRoundData) String() string { return proto.CompactTextString(m) }
func (*ChangeRoundData) ProtoMessage()    {}
func (*ChangeRoundData) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1187fce89a5e11, []int{2}
}

func (m *ChangeRoundData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeRoundData.Unmarshal(m, b)
}
func (m *ChangeRoundData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeRoundData.Marshal(b, m, deterministic)
}
func (m *ChangeRoundData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeRoundData.Merge(m, src)
}
func (m *ChangeRoundData) XXX_Size() int {
	return xxx_messageInfo_ChangeRoundData.Size(m)
}
func (m *ChangeRoundData) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeRoundData.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeRoundData proto.InternalMessageInfo

func (m *ChangeRoundData) GetPreparedRound() uint64 {
	if m != nil {
		return m.PreparedRound
	}
	return 0
}

func (m *ChangeRoundData) GetPreparedValue() []byte {
	if m != nil {
		return m.PreparedValue
	}
	return nil
}

func (m *ChangeRoundData) GetJustificationMsg() *Message {
	if m != nil {
		return m.JustificationMsg
	}
	return nil
}

func (m *ChangeRoundData) GetJustificationSig() []byte {
	if m != nil {
		return m.JustificationSig
	}
	return nil
}

func (m *ChangeRoundData) GetSignerIds() []uint64 {
	if m != nil {
		return m.SignerIds
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.RoundState", RoundState_name, RoundState_value)
	proto.RegisterType((*Message)(nil), "proto.Message")
	proto.RegisterType((*SignedMessage)(nil), "proto.SignedMessage")
	proto.RegisterType((*ChangeRoundData)(nil), "proto.ChangeRoundData")
}

func init() {
	proto.RegisterFile("github.com/bloxapp/ssv/ibft/proto/msgs.proto", fileDescriptor_dc1187fce89a5e11)
}

var fileDescriptor_dc1187fce89a5e11 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x5d, 0xd6, 0xa6, 0xd5, 0xbe, 0x76, 0x6d, 0x66, 0x4d, 0x28, 0x42, 0x42, 0x74, 0x45, 0x13,
	0x15, 0x3f, 0xad, 0x18, 0x4f, 0xc0, 0xb6, 0x1b, 0x24, 0x36, 0x55, 0x8d, 0xc4, 0x05, 0x37, 0x91,
	0x53, 0xbb, 0x9e, 0x59, 0x13, 0x7b, 0xb6, 0x53, 0xc1, 0x2d, 0x4f, 0xc7, 0x53, 0xc0, 0x7b, 0x70,
	0x85, 0xf2, 0x39, 0xd9, 0x4a, 0x05, 0xe2, 0xca, 0x3e, 0xe7, 0x3b, 0xc7, 0xc9, 0x39, 0x4e, 0xe0,
	0x95, 0x90, 0xee, 0xa6, 0xcc, 0xa6, 0x4b, 0x95, 0xcf, 0xb2, 0xb5, 0xfa, 0x42, 0xb5, 0x9e, 0x59,
	0xbb, 0x99, 0xc9, 0x6c, 0xe5, 0x66, 0xda, 0x28, 0xa7, 0x66, 0xb9, 0x15, 0x76, 0x8a, 0x5b, 0x12,
	0xe2, 0xf2, 0xf8, 0xf5, 0x96, 0x49, 0x28, 0xa1, 0xbc, 0x30, 0x2b, 0x57, 0x88, 0xbc, 0xab, 0xda,
	0x79, 0xd7, 0xf8, 0x67, 0x00, 0xdd, 0x2b, 0x6e, 0x2d, 0x15, 0x9c, 0x9c, 0x42, 0xdb, 0x7d, 0xd5,
	0x3c, 0x0e, 0x46, 0xc1, 0x64, 0x70, 0x76, 0xe4, 0x15, 0xd3, 0x85, 0x2a, 0x0b, 0x96, 0x38, 0xea,
	0xf8, 0x02, 0xc7, 0xe4, 0x18, 0x42, 0x53, 0x71, 0xf1, 0xfe, 0x28, 0x98, 0xb4, 0x17, 0x1e, 0x90,
	0x47, 0xd0, 0x59, 0xd3, 0x3c, 0x63, 0x34, 0x6e, 0x8d, 0x82, 0x49, 0x7f, 0x51, 0x23, 0xf2, 0x04,
	0xc0, 0xf2, 0xbb, 0xb4, 0x28, 0xf3, 0x8c, 0x9b, 0xb8, 0x8d, 0x96, 0x03, 0xcb, 0xef, 0xae, 0x91,
	0x20, 0xcf, 0x61, 0xa8, 0x0d, 0xdf, 0x48, 0x55, 0xda, 0xb4, 0xf6, 0x87, 0xe8, 0x1f, 0x34, 0xf4,
	0x07, 0x7f, 0xce, 0x31, 0x84, 0x1b, 0xba, 0x2e, 0x79, 0xdc, 0xc1, 0xb1, 0x07, 0xe4, 0x04, 0xfa,
	0x1b, 0xba, 0x96, 0x8c, 0x3a, 0x65, 0x52, 0x7d, 0x1b, 0x77, 0x71, 0xd8, 0xbb, 0xe7, 0xe6, 0xb7,
	0xe3, 0x6f, 0x01, 0x1c, 0x26, 0x52, 0x14, 0x9c, 0x35, 0x39, 0xa7, 0xd0, 0xcd, 0xfd, 0x16, 0xa3,
	0xf6, 0xce, 0x06, 0x75, 0xd4, 0x5a, 0x70, 0xde, 0xfe, 0xfe, 0xe3, 0xe9, 0xde, 0xa2, 0x11, 0x91,
	0x31, 0x1c, 0x58, 0x29, 0x0a, 0xea, 0x4a, 0xc3, 0x31, 0x74, 0xbf, 0x56, 0x3c, 0xd0, 0x18, 0xb3,
	0x7a, 0x88, 0x49, 0x25, 0xb3, 0x71, 0x6b, 0xd4, 0xc2, 0x98, 0xc8, 0xbc, 0x67, 0x76, 0xfc, 0x2b,
	0x80, 0xe1, 0xc5, 0x0d, 0x2d, 0x04, 0xc7, 0x3a, 0x2f, 0xa9, 0xa3, 0xe4, 0x14, 0xaa, 0x8c, 0x9a,
	0x1a, 0xce, 0x52, 0x5f, 0x68, 0x80, 0xed, 0x1c, 0x36, 0x2c, 0x4a, 0xc9, 0xcb, 0x2d, 0x99, 0x6f,
	0x60, 0xfb, 0x15, 0xee, 0xc5, 0x1f, 0xb1, 0x8f, 0x77, 0x70, 0xf4, 0xb9, 0xb4, 0x4e, 0xae, 0xe4,
	0x92, 0x3a, 0xa9, 0x8a, 0x34, 0xb7, 0x02, 0x2f, 0xe4, 0x5f, 0x21, 0xa3, 0x3f, 0xe4, 0x57, 0x56,
	0x90, 0x37, 0xbb, 0x47, 0x58, 0x29, 0xf0, 0xde, 0xfa, 0x7f, 0xb5, 0x24, 0x52, 0xec, 0x84, 0x0f,
	0x77, 0xc2, 0xbf, 0xd0, 0x00, 0x0f, 0x1f, 0x11, 0x19, 0x00, 0x5c, 0x2b, 0x97, 0x38, 0x6a, 0x1c,
	0x67, 0xd1, 0x5e, 0x85, 0xe7, 0x86, 0xcf, 0x7d, 0x8c, 0x28, 0x20, 0x3d, 0xe8, 0x36, 0x60, 0x9f,
	0x00, 0x74, 0x2e, 0x54, 0x9e, 0x4b, 0x17, 0xb5, 0xc8, 0x10, 0x7a, 0x5b, 0x15, 0x46, 0xed, 0x4a,
	0x79, 0xc9, 0x97, 0x92, 0x71, 0x16, 0x85, 0x15, 0x48, 0x9c, 0xd2, 0x9a, 0xb3, 0xa8, 0x73, 0xfe,
	0xec, 0xd3, 0xc9, 0x7f, 0xff, 0x9d, 0xac, 0x83, 0xcb, 0xdb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0xf4, 0x49, 0x22, 0x67, 0x03, 0x00, 0x00,
}
