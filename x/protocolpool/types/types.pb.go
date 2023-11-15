// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/protocolpool/v1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Budget defines the fields of a budget proposal.
type Budget struct {
	// recipient_address is the address of the recipient who can claim the budget.
	RecipientAddress string `protobuf:"bytes,1,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	// total_budget is the total amount allocated for the budget.
	TotalBudget *types.Coin `protobuf:"bytes,2,opt,name=total_budget,json=totalBudget,proto3" json:"total_budget,omitempty"`
	// claimed_amount is the total amount claimed from the total budget amount requested.
	ClaimedAmount *types.Coin `protobuf:"bytes,3,opt,name=claimed_amount,json=claimedAmount,proto3" json:"claimed_amount,omitempty"`
	// start_time is the time when the budget becomes claimable.
	StartTime *time.Time `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time,omitempty"`
	// next_claim_from is the time when the budget was last successfully claimed or distributed.
	// It is used to track the next starting claim time for fund distribution. If set, it cannot be less than start_time.
	NextClaimFrom *time.Time `protobuf:"bytes,5,opt,name=next_claim_from,json=nextClaimFrom,proto3,stdtime" json:"next_claim_from,omitempty"`
	// tranches is the number of times the total budget amount is to be distributed.
	Tranches uint64 `protobuf:"varint,6,opt,name=tranches,proto3" json:"tranches,omitempty"`
	// tranches_left is the number of tranches left for the amount to be distributed.
	TranchesLeft uint64 `protobuf:"varint,7,opt,name=tranches_left,json=tranchesLeft,proto3" json:"tranches_left,omitempty"`
	// Period is the time interval(number of seconds) at which funds distribution should be performed.
	// For example, if a period is set to 3600, it represents an action that
	// should occur every hour (3600 seconds).
	Period *time.Duration `protobuf:"bytes,8,opt,name=period,proto3,stdduration" json:"period,omitempty"`
}

func (m *Budget) Reset()         { *m = Budget{} }
func (m *Budget) String() string { return proto.CompactTextString(m) }
func (*Budget) ProtoMessage()    {}
func (*Budget) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b7d0ea246d7f44, []int{0}
}
func (m *Budget) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Budget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Budget.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Budget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Budget.Merge(m, src)
}
func (m *Budget) XXX_Size() int {
	return m.Size()
}
func (m *Budget) XXX_DiscardUnknown() {
	xxx_messageInfo_Budget.DiscardUnknown(m)
}

var xxx_messageInfo_Budget proto.InternalMessageInfo

func (m *Budget) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *Budget) GetTotalBudget() *types.Coin {
	if m != nil {
		return m.TotalBudget
	}
	return nil
}

func (m *Budget) GetClaimedAmount() *types.Coin {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

func (m *Budget) GetStartTime() *time.Time {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Budget) GetNextClaimFrom() *time.Time {
	if m != nil {
		return m.NextClaimFrom
	}
	return nil
}

func (m *Budget) GetTranches() uint64 {
	if m != nil {
		return m.Tranches
	}
	return 0
}

func (m *Budget) GetTranchesLeft() uint64 {
	if m != nil {
		return m.TranchesLeft
	}
	return 0
}

func (m *Budget) GetPeriod() *time.Duration {
	if m != nil {
		return m.Period
	}
	return nil
}

func init() {
	proto.RegisterType((*Budget)(nil), "cosmos.protocolpool.v1.Budget")
}

func init() {
	proto.RegisterFile("cosmos/protocolpool/v1/types.proto", fileDescriptor_c1b7d0ea246d7f44)
}

var fileDescriptor_c1b7d0ea246d7f44 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3f, 0x6f, 0xd4, 0x30,
	0x1c, 0xbd, 0xd0, 0xe3, 0x68, 0xdd, 0x1e, 0x7f, 0xa2, 0x0a, 0xb9, 0x19, 0xd2, 0xa3, 0x2c, 0xb7,
	0xe0, 0xe8, 0x60, 0x60, 0x00, 0x09, 0x9a, 0x02, 0x62, 0x60, 0x3a, 0x98, 0x58, 0x2c, 0x27, 0xf1,
	0x05, 0x8b, 0x24, 0xbf, 0xc8, 0xfe, 0xdd, 0xa9, 0x7c, 0x8b, 0x8e, 0x7c, 0x10, 0x3e, 0x04, 0x63,
	0x85, 0x18, 0xd8, 0x40, 0x77, 0x5f, 0x04, 0xc5, 0x76, 0xaa, 0xd2, 0x85, 0x6e, 0xce, 0xfb, 0xbd,
	0xf7, 0xf2, 0xfc, 0xfc, 0x23, 0x47, 0x39, 0x98, 0x1a, 0x4c, 0xd2, 0x6a, 0x40, 0xc8, 0xa1, 0x6a,
	0x01, 0xaa, 0x64, 0x35, 0x4b, 0xf0, 0x4b, 0x2b, 0x0d, 0xb3, 0x68, 0x78, 0xdf, 0x71, 0xd8, 0x65,
	0x0e, 0x5b, 0xcd, 0xa2, 0xfd, 0x12, 0x4a, 0xb0, 0x60, 0xd2, 0x9d, 0xdc, 0x3c, 0x3a, 0x70, 0x6c,
	0xee, 0x06, 0x97, 0xa5, 0x51, 0xec, 0x7f, 0x96, 0x09, 0x23, 0x93, 0xd5, 0x2c, 0x93, 0x28, 0x66,
	0x49, 0x0e, 0xaa, 0xf1, 0xf3, 0xc3, 0x12, 0xa0, 0xac, 0xa4, 0x0b, 0x93, 0x2d, 0x17, 0x09, 0xaa,
	0x5a, 0x1a, 0x14, 0x75, 0xdb, 0x1b, 0x5c, 0x25, 0x14, 0x4b, 0x2d, 0x50, 0x81, 0x37, 0x38, 0xfa,
	0xb9, 0x45, 0x46, 0xe9, 0xb2, 0x28, 0x25, 0x86, 0xaf, 0xc9, 0x3d, 0x2d, 0x73, 0xd5, 0x2a, 0xd9,
	0x20, 0x17, 0x45, 0xa1, 0xa5, 0x31, 0x34, 0x98, 0x04, 0xd3, 0x9d, 0x94, 0xfe, 0xf8, 0xf6, 0x68,
	0xdf, 0x07, 0x3b, 0x76, 0x93, 0xf7, 0xa8, 0x55, 0x53, 0xce, 0xef, 0x5e, 0x48, 0x3c, 0x1e, 0x3e,
	0x27, 0x7b, 0x08, 0x28, 0x2a, 0x9e, 0x59, 0x5b, 0x7a, 0x63, 0x12, 0x4c, 0x77, 0x1f, 0x1f, 0x30,
	0x2f, 0xef, 0x6e, 0xc2, 0xfc, 0x4d, 0xd8, 0x09, 0xa8, 0x66, 0xbe, 0x6b, 0xe9, 0x3e, 0xc4, 0x4b,
	0x72, 0x3b, 0xaf, 0x84, 0xaa, 0x65, 0xc1, 0x45, 0x0d, 0xcb, 0x06, 0xe9, 0xd6, 0xff, 0xf4, 0x63,
	0x2f, 0x38, 0xb6, 0xfc, 0xf0, 0x05, 0x21, 0x06, 0x85, 0x46, 0xde, 0x55, 0x41, 0x87, 0x56, 0x1d,
	0x31, 0x57, 0x03, 0xeb, 0x6b, 0x60, 0x1f, 0xfa, 0x9e, 0xd2, 0xe1, 0xd9, 0xef, 0xc3, 0x60, 0xbe,
	0x63, 0x35, 0x1d, 0x1a, 0xbe, 0x25, 0x77, 0x1a, 0x79, 0x8a, 0xdc, 0xda, 0xf2, 0x85, 0x86, 0x9a,
	0xde, 0xbc, 0xa6, 0xcb, 0xb8, 0x13, 0x9e, 0x74, 0xba, 0x37, 0x1a, 0xea, 0x30, 0x22, 0xdb, 0xa8,
	0x45, 0x93, 0x7f, 0x92, 0x86, 0x8e, 0x26, 0xc1, 0x74, 0x38, 0xbf, 0xf8, 0x0e, 0x1f, 0x92, 0x71,
	0x7f, 0xe6, 0x95, 0x5c, 0x20, 0xbd, 0x65, 0x09, 0x7b, 0x3d, 0xf8, 0x4e, 0x2e, 0x30, 0x7c, 0x4a,
	0x46, 0xad, 0xd4, 0x0a, 0x0a, 0xba, 0xed, 0x5b, 0xb8, 0x9a, 0xe0, 0x95, 0x7f, 0xce, 0x74, 0xf8,
	0xb5, 0x0b, 0xe0, 0xe9, 0xe9, 0xb3, 0xef, 0xeb, 0x38, 0x38, 0x5f, 0xc7, 0xc1, 0x9f, 0x75, 0x1c,
	0x9c, 0x6d, 0xe2, 0xc1, 0xf9, 0x26, 0x1e, 0xfc, 0xda, 0xc4, 0x83, 0x8f, 0x0f, 0x5c, 0x8f, 0xa6,
	0xf8, 0xcc, 0x14, 0x24, 0xa7, 0xff, 0xae, 0xb1, 0xdd, 0xe1, 0x6c, 0x64, 0xb1, 0x27, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x07, 0x16, 0x5b, 0x67, 0xea, 0x02, 0x00, 0x00,
}

func (m *Budget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Budget) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Budget) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Period != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(*m.Period, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(*m.Period):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintTypes(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x42
	}
	if m.TranchesLeft != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.TranchesLeft))
		i--
		dAtA[i] = 0x38
	}
	if m.Tranches != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Tranches))
		i--
		dAtA[i] = 0x30
	}
	if m.NextClaimFrom != nil {
		n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.NextClaimFrom, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.NextClaimFrom):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintTypes(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x2a
	}
	if m.StartTime != nil {
		n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.StartTime):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintTypes(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x22
	}
	if m.ClaimedAmount != nil {
		{
			size, err := m.ClaimedAmount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.TotalBudget != nil {
		{
			size, err := m.TotalBudget.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Budget) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.TotalBudget != nil {
		l = m.TotalBudget.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ClaimedAmount != nil {
		l = m.ClaimedAmount.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.StartTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.StartTime)
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.NextClaimFrom != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.NextClaimFrom)
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Tranches != 0 {
		n += 1 + sovTypes(uint64(m.Tranches))
	}
	if m.TranchesLeft != 0 {
		n += 1 + sovTypes(uint64(m.TranchesLeft))
	}
	if m.Period != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(*m.Period)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Budget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Budget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Budget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBudget", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TotalBudget == nil {
				m.TotalBudget = &types.Coin{}
			}
			if err := m.TotalBudget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClaimedAmount == nil {
				m.ClaimedAmount = &types.Coin{}
			}
			if err := m.ClaimedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartTime == nil {
				m.StartTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextClaimFrom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NextClaimFrom == nil {
				m.NextClaimFrom = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.NextClaimFrom, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tranches", wireType)
			}
			m.Tranches = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tranches |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TranchesLeft", wireType)
			}
			m.TranchesLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TranchesLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Period == nil {
				m.Period = new(time.Duration)
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(m.Period, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
