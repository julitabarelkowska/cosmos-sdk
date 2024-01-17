// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/defaults/base/v1/account.proto

package v1

import (
	_ "cosmossdk.io/x/accounts/v1"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	signing "github.com/cosmos/cosmos-sdk/types/tx/signing"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgInit is used to initialize a base account.
type MsgInit struct {
	// pub_key defines the secp256k1 pubkey for the account.
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *MsgInit) Reset()         { *m = MsgInit{} }
func (m *MsgInit) String() string { return proto.CompactTextString(m) }
func (*MsgInit) ProtoMessage()    {}
func (*MsgInit) Descriptor() ([]byte, []int) {
	return fileDescriptor_f112463f97a32614, []int{0}
}
func (m *MsgInit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInit.Merge(m, src)
}
func (m *MsgInit) XXX_Size() int {
	return m.Size()
}
func (m *MsgInit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInit proto.InternalMessageInfo

func (m *MsgInit) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

// MsgInitResponse is the response returned after base account initialization.
// This is empty.
type MsgInitResponse struct {
}

func (m *MsgInitResponse) Reset()         { *m = MsgInitResponse{} }
func (m *MsgInitResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitResponse) ProtoMessage()    {}
func (*MsgInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f112463f97a32614, []int{1}
}
func (m *MsgInitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitResponse.Merge(m, src)
}
func (m *MsgInitResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitResponse proto.InternalMessageInfo

// MsgSwapPubKey is used to change the pubkey for the account.
type MsgSwapPubKey struct {
	// new_pub_key defines the secp256k1 pubkey to swap the account to.
	NewPubKey []byte `protobuf:"bytes,1,opt,name=new_pub_key,json=newPubKey,proto3" json:"new_pub_key,omitempty"`
}

func (m *MsgSwapPubKey) Reset()         { *m = MsgSwapPubKey{} }
func (m *MsgSwapPubKey) String() string { return proto.CompactTextString(m) }
func (*MsgSwapPubKey) ProtoMessage()    {}
func (*MsgSwapPubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_f112463f97a32614, []int{2}
}
func (m *MsgSwapPubKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapPubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapPubKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapPubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapPubKey.Merge(m, src)
}
func (m *MsgSwapPubKey) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapPubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapPubKey.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapPubKey proto.InternalMessageInfo

func (m *MsgSwapPubKey) GetNewPubKey() []byte {
	if m != nil {
		return m.NewPubKey
	}
	return nil
}

// MsgSwapPubKeyResponse is the response for the MsgSwapPubKey message.
// This is empty.
type MsgSwapPubKeyResponse struct {
}

func (m *MsgSwapPubKeyResponse) Reset()         { *m = MsgSwapPubKeyResponse{} }
func (m *MsgSwapPubKeyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapPubKeyResponse) ProtoMessage()    {}
func (*MsgSwapPubKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f112463f97a32614, []int{3}
}
func (m *MsgSwapPubKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapPubKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapPubKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapPubKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapPubKeyResponse.Merge(m, src)
}
func (m *MsgSwapPubKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapPubKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapPubKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapPubKeyResponse proto.InternalMessageInfo

// QuerySequence is the request for the account sequence.
type QuerySequence struct {
}

func (m *QuerySequence) Reset()         { *m = QuerySequence{} }
func (m *QuerySequence) String() string { return proto.CompactTextString(m) }
func (*QuerySequence) ProtoMessage()    {}
func (*QuerySequence) Descriptor() ([]byte, []int) {
	return fileDescriptor_f112463f97a32614, []int{4}
}
func (m *QuerySequence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySequence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySequence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySequence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySequence.Merge(m, src)
}
func (m *QuerySequence) XXX_Size() int {
	return m.Size()
}
func (m *QuerySequence) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySequence.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySequence proto.InternalMessageInfo

// QuerySequenceResponse returns the sequence of the account.
type QuerySequenceResponse struct {
	// sequence is the current sequence of the account.
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *QuerySequenceResponse) Reset()         { *m = QuerySequenceResponse{} }
func (m *QuerySequenceResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySequenceResponse) ProtoMessage()    {}
func (*QuerySequenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f112463f97a32614, []int{5}
}
func (m *QuerySequenceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySequenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySequenceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySequenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySequenceResponse.Merge(m, src)
}
func (m *QuerySequenceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySequenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySequenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySequenceResponse proto.InternalMessageInfo

func (m *QuerySequenceResponse) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// AuthenticationData is what the base account needs to verify signatures
// in an operation.
type AuthenticationData struct {
	// sign_mode defines the sign mode used by the account to authenticate
	// a tx.
	SignMode signing.SignMode `protobuf:"varint,1,opt,name=sign_mode,json=signMode,proto3,enum=cosmos.tx.signing.v1beta1.SignMode" json:"sign_mode,omitempty"`
	// signature defines the signature produced by the provided sign_mode.
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *AuthenticationData) Reset()         { *m = AuthenticationData{} }
func (m *AuthenticationData) String() string { return proto.CompactTextString(m) }
func (*AuthenticationData) ProtoMessage()    {}
func (*AuthenticationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f112463f97a32614, []int{6}
}
func (m *AuthenticationData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticationData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationData.Merge(m, src)
}
func (m *AuthenticationData) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticationData) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationData.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationData proto.InternalMessageInfo

func (m *AuthenticationData) GetSignMode() signing.SignMode {
	if m != nil {
		return m.SignMode
	}
	return signing.SignMode_SIGN_MODE_UNSPECIFIED
}

func (m *AuthenticationData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgInit)(nil), "cosmos.accounts.defaults.base.v1.MsgInit")
	proto.RegisterType((*MsgInitResponse)(nil), "cosmos.accounts.defaults.base.v1.MsgInitResponse")
	proto.RegisterType((*MsgSwapPubKey)(nil), "cosmos.accounts.defaults.base.v1.MsgSwapPubKey")
	proto.RegisterType((*MsgSwapPubKeyResponse)(nil), "cosmos.accounts.defaults.base.v1.MsgSwapPubKeyResponse")
	proto.RegisterType((*QuerySequence)(nil), "cosmos.accounts.defaults.base.v1.QuerySequence")
	proto.RegisterType((*QuerySequenceResponse)(nil), "cosmos.accounts.defaults.base.v1.QuerySequenceResponse")
	proto.RegisterType((*AuthenticationData)(nil), "cosmos.accounts.defaults.base.v1.AuthenticationData")
}

func init() {
	proto.RegisterFile("cosmos/accounts/defaults/base/v1/account.proto", fileDescriptor_f112463f97a32614)
}

var fileDescriptor_f112463f97a32614 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xe9, 0x9b, 0x37, 0xfc, 0xd9, 0xf7, 0x45, 0x62, 0x13, 0x02, 0x12, 0xd3, 0x90, 0x7a,
	0x90, 0x83, 0xd9, 0x4d, 0xe5, 0x0b, 0x28, 0xf1, 0x62, 0x0c, 0x89, 0x96, 0x9b, 0x97, 0x66, 0xdb,
	0x0e, 0xb5, 0x01, 0x76, 0x6b, 0x77, 0x17, 0xe8, 0xb7, 0xf0, 0x63, 0x79, 0xe4, 0xe8, 0xd1, 0xc0,
	0x17, 0x31, 0x74, 0x5b, 0x0c, 0x26, 0xde, 0x76, 0x66, 0x7e, 0xf3, 0x4c, 0xe6, 0x99, 0x45, 0x38,
	0xe0, 0x62, 0xc1, 0x05, 0xa1, 0x41, 0xc0, 0x15, 0x93, 0x82, 0x84, 0x30, 0xa5, 0x6a, 0x2e, 0x05,
	0xf1, 0xa9, 0x00, 0xb2, 0x74, 0xca, 0x0a, 0x4e, 0x52, 0x2e, 0xb9, 0xd9, 0xd7, 0x3c, 0x2e, 0x79,
	0x5c, 0xf2, 0x78, 0xcf, 0xe3, 0xa5, 0xd3, 0x3b, 0x8b, 0x38, 0x8f, 0xe6, 0x40, 0x72, 0xde, 0x57,
	0x53, 0x42, 0x59, 0xa6, 0x9b, 0x7b, 0x57, 0x3f, 0x87, 0x7d, 0xcb, 0x7b, 0xd4, 0x17, 0x32, 0xa5,
	0x81, 0x8c, 0x39, 0x2b, 0xe8, 0xcb, 0x82, 0x96, 0x6b, 0x22, 0xe2, 0x88, 0xc5, 0x2c, 0x22, 0x4b,
	0xc7, 0x07, 0x49, 0x9d, 0x32, 0xd6, 0xa0, 0x6d, 0xa3, 0xda, 0x58, 0x44, 0xf7, 0x2c, 0x96, 0x66,
	0x07, 0xd5, 0x12, 0xe5, 0x7b, 0x33, 0xc8, 0xba, 0x46, 0xdf, 0x18, 0xfc, 0x77, 0xab, 0x89, 0xf2,
	0x1f, 0x20, 0xb3, 0x4f, 0x51, 0xab, 0x60, 0x5c, 0x10, 0x09, 0x67, 0x02, 0x6c, 0x82, 0x9a, 0x63,
	0x11, 0x4d, 0x56, 0x34, 0x79, 0xcc, 0x19, 0xd3, 0x42, 0xff, 0x18, 0xac, 0xbc, 0x63, 0x81, 0x06,
	0x83, 0x95, 0xae, 0xdb, 0x1d, 0xd4, 0x3e, 0x6a, 0x38, 0x28, 0xb5, 0x50, 0xf3, 0x49, 0x41, 0x9a,
	0x4d, 0xe0, 0x55, 0x01, 0x0b, 0xc0, 0x1e, 0xa2, 0xf6, 0x51, 0xa2, 0x24, 0xcd, 0x1e, 0xaa, 0x8b,
	0x22, 0x97, 0xeb, 0xff, 0x75, 0x0f, 0xb1, 0x2d, 0x91, 0x79, 0xab, 0xe4, 0x0b, 0x30, 0x19, 0x07,
	0x74, 0xef, 0xc3, 0x1d, 0x95, 0xd4, 0xbc, 0x41, 0x8d, 0xfd, 0xb6, 0xde, 0x82, 0x87, 0xba, 0xe5,
	0xe4, 0xfa, 0xa2, 0x38, 0x1a, 0x96, 0x6b, 0x5c, 0x3a, 0x51, 0x38, 0x83, 0x27, 0x71, 0xc4, 0xc6,
	0x3c, 0x04, 0xb7, 0x2e, 0x8a, 0x97, 0x79, 0xae, 0x15, 0xa8, 0x54, 0x29, 0x74, 0xff, 0xe8, 0xa5,
	0x0e, 0x89, 0xd1, 0xe8, 0x7d, 0x6b, 0x19, 0x9b, 0xad, 0x65, 0x7c, 0x6e, 0x2d, 0xe3, 0x6d, 0x67,
	0x55, 0x36, 0x3b, 0xab, 0xf2, 0xb1, 0xb3, 0x2a, 0xcf, 0x03, 0x3d, 0x45, 0x84, 0x33, 0x1c, 0x73,
	0xb2, 0xfe, 0xfd, 0x8b, 0xf8, 0xd5, 0xfc, 0x0e, 0xc3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7,
	0xe3, 0x34, 0x53, 0x4d, 0x02, 0x00, 0x00,
}

func (m *MsgInit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSwapPubKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapPubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapPubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewPubKey) > 0 {
		i -= len(m.NewPubKey)
		copy(dAtA[i:], m.NewPubKey)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.NewPubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapPubKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapPubKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapPubKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuerySequence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySequence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySequence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuerySequenceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySequenceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySequenceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AuthenticationData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticationData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthenticationData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if m.SignMode != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.SignMode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func (m *MsgInitResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSwapPubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NewPubKey)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func (m *MsgSwapPubKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuerySequence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuerySequenceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovAccount(uint64(m.Sequence))
	}
	return n
}

func (m *AuthenticationData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignMode != 0 {
		n += 1 + sovAccount(uint64(m.SignMode))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgInitResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSwapPubKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapPubKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapPubKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewPubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewPubKey = append(m.NewPubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.NewPubKey == nil {
				m.NewPubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSwapPubKeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapPubKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapPubKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySequence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySequence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySequence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySequenceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySequenceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySequenceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthenticationData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthenticationData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticationData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignMode", wireType)
			}
			m.SignMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignMode |= signing.SignMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)
