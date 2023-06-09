// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: derivatives/perpetual_options.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type OptionType int32

const (
	OptionType_OPTION_UNKNOWN OptionType = 0
	OptionType_CALL           OptionType = 1
	OptionType_PUT            OptionType = 2
)

var OptionType_name = map[int32]string{
	0: "OPTION_UNKNOWN",
	1: "CALL",
	2: "PUT",
}

var OptionType_value = map[string]int32{
	"OPTION_UNKNOWN": 0,
	"CALL":           1,
	"PUT":            2,
}

func (x OptionType) String() string {
	return proto.EnumName(OptionType_name, int32(x))
}

func (OptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8ecd39830a7920f, []int{0}
}

type PerpetualOptionsParams struct {
	PremiumCommissionRate                       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=premium_commission_rate,json=premiumCommissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"premium_commission_rate" yaml:"premium_commission_rate"`
	StrikeCommissionRate                        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=strike_commission_rate,json=strikeCommissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"strike_commission_rate" yaml:"premium_commission_rate"`
	MarginMaintenanceRate                       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=margin_maintenance_rate,json=marginMaintenanceRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"margin_maintenance_rate" yaml:"margin_maintenance_rate"`
	ImaginaryFundingRateProportionalCoefficient github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=imaginary_funding_rate_proportional_coefficient,json=imaginaryFundingRateProportionalCoefficient,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"imaginary_funding_rate_proportional_coefficient" yaml:"imaginary_funding_rate_proportonal_coefficient"`
	Markets                                     []*Market                              `protobuf:"bytes,5,rep,name=markets,proto3" json:"markets,omitempty" yaml:"markets"`
}

func (m *PerpetualOptionsParams) Reset()         { *m = PerpetualOptionsParams{} }
func (m *PerpetualOptionsParams) String() string { return proto.CompactTextString(m) }
func (*PerpetualOptionsParams) ProtoMessage()    {}
func (*PerpetualOptionsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ecd39830a7920f, []int{0}
}
func (m *PerpetualOptionsParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerpetualOptionsParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerpetualOptionsParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerpetualOptionsParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerpetualOptionsParams.Merge(m, src)
}
func (m *PerpetualOptionsParams) XXX_Size() int {
	return m.Size()
}
func (m *PerpetualOptionsParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PerpetualOptionsParams.DiscardUnknown(m)
}

var xxx_messageInfo_PerpetualOptionsParams proto.InternalMessageInfo

func (m *PerpetualOptionsParams) GetMarkets() []*Market {
	if m != nil {
		return m.Markets
	}
	return nil
}

type PerpetualOptionsPositionInstance struct {
	OptionType   OptionType                             `protobuf:"varint,1,opt,name=option_type,json=optionType,proto3,enum=ununifi.derivatives.OptionType" json:"option_type,omitempty" yaml:"option_type"`
	PositionType PositionType                           `protobuf:"varint,2,opt,name=position_type,json=positionType,proto3,enum=ununifi.derivatives.PositionType" json:"position_type,omitempty" yaml:"position_type"`
	StrikePrice  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=strike_price,json=strikePrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"strike_price" yaml:"strike_price"`
	Premium      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=premium,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"premium" yaml:"premium"`
}

func (m *PerpetualOptionsPositionInstance) Reset()         { *m = PerpetualOptionsPositionInstance{} }
func (m *PerpetualOptionsPositionInstance) String() string { return proto.CompactTextString(m) }
func (*PerpetualOptionsPositionInstance) ProtoMessage()    {}
func (*PerpetualOptionsPositionInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ecd39830a7920f, []int{1}
}
func (m *PerpetualOptionsPositionInstance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerpetualOptionsPositionInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerpetualOptionsPositionInstance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerpetualOptionsPositionInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerpetualOptionsPositionInstance.Merge(m, src)
}
func (m *PerpetualOptionsPositionInstance) XXX_Size() int {
	return m.Size()
}
func (m *PerpetualOptionsPositionInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_PerpetualOptionsPositionInstance.DiscardUnknown(m)
}

var xxx_messageInfo_PerpetualOptionsPositionInstance proto.InternalMessageInfo

func (m *PerpetualOptionsPositionInstance) GetOptionType() OptionType {
	if m != nil {
		return m.OptionType
	}
	return OptionType_OPTION_UNKNOWN
}

func (m *PerpetualOptionsPositionInstance) GetPositionType() PositionType {
	if m != nil {
		return m.PositionType
	}
	return PositionType_POSITION_UNKNOWN
}

func init() {
	proto.RegisterEnum("ununifi.derivatives.OptionType", OptionType_name, OptionType_value)
	proto.RegisterType((*PerpetualOptionsParams)(nil), "ununifi.derivatives.PerpetualOptionsParams")
	proto.RegisterType((*PerpetualOptionsPositionInstance)(nil), "ununifi.derivatives.PerpetualOptionsPositionInstance")
}

func init() {
	proto.RegisterFile("derivatives/perpetual_options.proto", fileDescriptor_a8ecd39830a7920f)
}

var fileDescriptor_a8ecd39830a7920f = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xe3, 0xa6, 0xa5, 0x30, 0x2d, 0x51, 0x35, 0xbd, 0x10, 0x8a, 0xb0, 0x8b, 0x91, 0x50,
	0xc5, 0x25, 0x56, 0x5b, 0xb1, 0x61, 0x05, 0x29, 0x54, 0xaa, 0x68, 0x13, 0xcb, 0x6a, 0x04, 0xea,
	0xc6, 0x4c, 0xdc, 0x89, 0x3b, 0x6a, 0xe6, 0xa2, 0x99, 0x71, 0x45, 0x5e, 0x00, 0xb1, 0xe4, 0x21,
	0x78, 0x98, 0x6e, 0x90, 0x2a, 0x56, 0x88, 0x45, 0x84, 0xda, 0x37, 0xc8, 0x13, 0x20, 0x7b, 0x9c,
	0xd4, 0xf4, 0x82, 0x54, 0x21, 0x56, 0x19, 0xe7, 0xfc, 0xe7, 0x3f, 0xdf, 0x1c, 0x9f, 0x63, 0xf0,
	0x70, 0x0f, 0x4b, 0x72, 0x88, 0x34, 0x39, 0xc4, 0xca, 0x13, 0x58, 0x0a, 0xac, 0x13, 0xd4, 0x0d,
	0xb9, 0xd0, 0x84, 0x33, 0x55, 0x13, 0x92, 0x6b, 0x0e, 0x67, 0x13, 0x96, 0x30, 0xd2, 0x21, 0xb5,
	0x82, 0x78, 0xf1, 0x6e, 0xcc, 0x79, 0xdc, 0xc5, 0x5e, 0x26, 0x69, 0x27, 0x1d, 0x0f, 0xb1, 0x9e,
	0xd1, 0x2f, 0xce, 0xc5, 0x3c, 0xe6, 0xd9, 0xd1, 0x4b, 0x4f, 0xf9, 0xbf, 0xce, 0xf9, 0x04, 0x4d,
	0x28, 0x56, 0x1a, 0x51, 0x91, 0x0b, 0xec, 0x88, 0x2b, 0xca, 0x95, 0xd7, 0x46, 0x0a, 0x7b, 0x87,
	0x2b, 0x6d, 0xac, 0xd1, 0x8a, 0x17, 0x71, 0xc2, 0xf2, 0xf8, 0xfd, 0x22, 0x6b, 0xe1, 0x6c, 0xc2,
	0xee, 0xf7, 0x09, 0xb0, 0xe0, 0x0f, 0x6f, 0xd0, 0x34, 0x17, 0xf0, 0x91, 0x44, 0x54, 0xc1, 0xcf,
	0x16, 0xb8, 0x23, 0x24, 0xa6, 0x24, 0xa1, 0x61, 0xc4, 0x29, 0x25, 0x4a, 0x11, 0xce, 0x42, 0x89,
	0x34, 0xae, 0x5a, 0x4b, 0xd6, 0xf2, 0xad, 0xba, 0x7f, 0xd4, 0x77, 0x4a, 0x3f, 0xfb, 0xce, 0xa3,
	0x98, 0xe8, 0xfd, 0xa4, 0x5d, 0x8b, 0x38, 0xf5, 0x72, 0x1c, 0xf3, 0xf3, 0x4c, 0xed, 0x1d, 0x78,
	0xba, 0x27, 0xb0, 0xaa, 0xbd, 0xc6, 0xd1, 0xa0, 0xef, 0xd8, 0x3d, 0x44, 0xbb, 0x2f, 0xdc, 0x2b,
	0x6c, 0xdd, 0x60, 0x3e, 0x8f, 0xac, 0x8f, 0x02, 0x01, 0xd2, 0x18, 0x7e, 0xb2, 0xc0, 0x82, 0xd2,
	0x92, 0x1c, 0xe0, 0x0b, 0x24, 0x63, 0xff, 0x89, 0x64, 0xce, 0xd4, 0x3b, 0x07, 0x92, 0xf6, 0x84,
	0x22, 0x19, 0x13, 0x16, 0x52, 0x44, 0x98, 0xc6, 0x0c, 0xb1, 0x08, 0x1b, 0x92, 0xf2, 0xbf, 0x91,
	0x5c, 0x61, 0xeb, 0x06, 0xf3, 0x26, 0xb2, 0x7d, 0x16, 0xc8, 0x50, 0xbe, 0x59, 0xc0, 0x23, 0x14,
	0xc5, 0x84, 0x21, 0xd9, 0x0b, 0x3b, 0x09, 0xdb, 0x23, 0x2c, 0xce, 0x52, 0x42, 0x21, 0xb9, 0xe0,
	0x32, 0x7d, 0x97, 0xa8, 0x1b, 0x46, 0x1c, 0x77, 0x3a, 0x24, 0x22, 0x98, 0xe9, 0xea, 0x78, 0x86,
	0x18, 0x5f, 0x1b, 0xf1, 0xb9, 0x41, 0xfc, 0x7b, 0xb9, 0xf3, 0xd5, 0xdc, 0xe0, 0xc9, 0x28, 0x61,
	0xc3, 0xe8, 0x53, 0x70, 0xbf, 0x00, 0xb7, 0x7e, 0xa6, 0x86, 0x9b, 0x60, 0x92, 0x22, 0x79, 0x80,
	0xb5, 0xaa, 0x4e, 0x2c, 0x95, 0x97, 0xa7, 0x56, 0xef, 0xd5, 0x2e, 0xd9, 0xa0, 0xda, 0x76, 0xa6,
	0xa9, 0xc3, 0x41, 0xdf, 0xa9, 0x8c, 0x1a, 0x97, 0x66, 0xb9, 0xc1, 0x30, 0xdf, 0xfd, 0x5a, 0x06,
	0x4b, 0x17, 0x86, 0x9a, 0x2b, 0x92, 0x1e, 0x36, 0x99, 0xd2, 0x69, 0x0f, 0xe1, 0x7b, 0x30, 0x65,
	0x16, 0x36, 0x4c, 0x6f, 0x9a, 0x4d, 0x74, 0x65, 0xd5, 0xb9, 0xb4, 0xa6, 0xb1, 0xd8, 0xe9, 0x09,
	0x5c, 0x5f, 0x18, 0xf4, 0x1d, 0x68, 0xea, 0x16, 0xb2, 0xdd, 0x00, 0xf0, 0x91, 0x06, 0x7e, 0x00,
	0xb7, 0x45, 0x5e, 0xcd, 0x78, 0x8f, 0x65, 0xde, 0x0f, 0x2e, 0xf5, 0x1e, 0x72, 0x65, 0xee, 0xd5,
	0x41, 0xdf, 0x99, 0xcb, 0x07, 0xb3, 0xe8, 0xe0, 0x06, 0xd3, 0xa2, 0xa0, 0x83, 0xfb, 0x60, 0x3a,
	0x5f, 0x07, 0x21, 0x49, 0x34, 0x1c, 0xbd, 0x37, 0xd7, 0x7e, 0xaf, 0xb3, 0xa6, 0x56, 0xd1, 0xcb,
	0x0d, 0xa6, 0xcc, 0xa3, 0x9f, 0x3e, 0xc1, 0x5d, 0x30, 0x99, 0xaf, 0x48, 0x3e, 0x3c, 0x2f, 0xaf,
	0x5d, 0xa4, 0xf2, 0xc7, 0xa6, 0xb9, 0xc1, 0xd0, 0xf0, 0xf1, 0x1a, 0x00, 0x67, 0x9d, 0x85, 0x10,
	0x54, 0x9a, 0xfe, 0xce, 0x66, 0xb3, 0x11, 0xb6, 0x1a, 0x6f, 0x1b, 0xcd, 0x77, 0x8d, 0x99, 0x12,
	0xbc, 0x09, 0xc6, 0xd7, 0x5f, 0x6d, 0x6d, 0xcd, 0x58, 0x70, 0x12, 0x94, 0xfd, 0xd6, 0xce, 0xcc,
	0x58, 0x7d, 0xe3, 0xe8, 0xc4, 0xb6, 0x8e, 0x4f, 0x6c, 0xeb, 0xd7, 0x89, 0x6d, 0x7d, 0x39, 0xb5,
	0x4b, 0xc7, 0xa7, 0x76, 0xe9, 0xc7, 0xa9, 0x5d, 0xda, 0x7d, 0x5a, 0x20, 0x6a, 0xb1, 0x16, 0x23,
	0x1b, 0xc4, 0x8b, 0xf6, 0x11, 0x61, 0xde, 0xc7, 0xe2, 0x87, 0xcf, 0xb0, 0xb5, 0x6f, 0x64, 0xdf,
	0xbf, 0xb5, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x0a, 0xab, 0x92, 0xcc, 0x05, 0x00, 0x00,
}

func (m *PerpetualOptionsParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerpetualOptionsParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerpetualOptionsParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Markets) > 0 {
		for iNdEx := len(m.Markets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Markets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPerpetualOptions(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size := m.ImaginaryFundingRateProportionalCoefficient.Size()
		i -= size
		if _, err := m.ImaginaryFundingRateProportionalCoefficient.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualOptions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.MarginMaintenanceRate.Size()
		i -= size
		if _, err := m.MarginMaintenanceRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualOptions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.StrikeCommissionRate.Size()
		i -= size
		if _, err := m.StrikeCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualOptions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.PremiumCommissionRate.Size()
		i -= size
		if _, err := m.PremiumCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualOptions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PerpetualOptionsPositionInstance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerpetualOptionsPositionInstance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerpetualOptionsPositionInstance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Premium.Size()
		i -= size
		if _, err := m.Premium.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualOptions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.StrikePrice.Size()
		i -= size
		if _, err := m.StrikePrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualOptions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.PositionType != 0 {
		i = encodeVarintPerpetualOptions(dAtA, i, uint64(m.PositionType))
		i--
		dAtA[i] = 0x10
	}
	if m.OptionType != 0 {
		i = encodeVarintPerpetualOptions(dAtA, i, uint64(m.OptionType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPerpetualOptions(dAtA []byte, offset int, v uint64) int {
	offset -= sovPerpetualOptions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PerpetualOptionsParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PremiumCommissionRate.Size()
	n += 1 + l + sovPerpetualOptions(uint64(l))
	l = m.StrikeCommissionRate.Size()
	n += 1 + l + sovPerpetualOptions(uint64(l))
	l = m.MarginMaintenanceRate.Size()
	n += 1 + l + sovPerpetualOptions(uint64(l))
	l = m.ImaginaryFundingRateProportionalCoefficient.Size()
	n += 1 + l + sovPerpetualOptions(uint64(l))
	if len(m.Markets) > 0 {
		for _, e := range m.Markets {
			l = e.Size()
			n += 1 + l + sovPerpetualOptions(uint64(l))
		}
	}
	return n
}

func (m *PerpetualOptionsPositionInstance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OptionType != 0 {
		n += 1 + sovPerpetualOptions(uint64(m.OptionType))
	}
	if m.PositionType != 0 {
		n += 1 + sovPerpetualOptions(uint64(m.PositionType))
	}
	l = m.StrikePrice.Size()
	n += 1 + l + sovPerpetualOptions(uint64(l))
	l = m.Premium.Size()
	n += 1 + l + sovPerpetualOptions(uint64(l))
	return n
}

func sovPerpetualOptions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPerpetualOptions(x uint64) (n int) {
	return sovPerpetualOptions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PerpetualOptionsParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerpetualOptions
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
			return fmt.Errorf("proto: PerpetualOptionsParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerpetualOptionsParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PremiumCommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualOptions
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
				return ErrInvalidLengthPerpetualOptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PremiumCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrikeCommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualOptions
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
				return ErrInvalidLengthPerpetualOptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StrikeCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarginMaintenanceRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualOptions
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
				return ErrInvalidLengthPerpetualOptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MarginMaintenanceRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImaginaryFundingRateProportionalCoefficient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualOptions
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
				return ErrInvalidLengthPerpetualOptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ImaginaryFundingRateProportionalCoefficient.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Markets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualOptions
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
				return ErrInvalidLengthPerpetualOptions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Markets = append(m.Markets, &Market{})
			if err := m.Markets[len(m.Markets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerpetualOptions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPerpetualOptions
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
func (m *PerpetualOptionsPositionInstance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerpetualOptions
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
			return fmt.Errorf("proto: PerpetualOptionsPositionInstance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerpetualOptionsPositionInstance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionType", wireType)
			}
			m.OptionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualOptions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OptionType |= OptionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionType", wireType)
			}
			m.PositionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualOptions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PositionType |= PositionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrikePrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualOptions
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
				return ErrInvalidLengthPerpetualOptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StrikePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Premium", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualOptions
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
				return ErrInvalidLengthPerpetualOptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Premium.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerpetualOptions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPerpetualOptions
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
func skipPerpetualOptions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPerpetualOptions
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
					return 0, ErrIntOverflowPerpetualOptions
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
					return 0, ErrIntOverflowPerpetualOptions
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
				return 0, ErrInvalidLengthPerpetualOptions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPerpetualOptions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPerpetualOptions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPerpetualOptions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPerpetualOptions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPerpetualOptions = fmt.Errorf("proto: unexpected end of group")
)
