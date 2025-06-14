// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: optimizeglobalcoin/asset/asset.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

type AssetType int32

const (
	AssetType_REAL_ESTATE     AssetType = 0
	AssetType_FINANCIAL_ASSET AssetType = 1
	AssetType_OTHER_ASSET     AssetType = 2
)

var AssetType_name = map[int32]string{
	0: "REAL_ESTATE",
	1: "FINANCIAL_ASSET",
	2: "OTHER_ASSET",
}

var AssetType_value = map[string]int32{
	"REAL_ESTATE":     0,
	"FINANCIAL_ASSET": 1,
	"OTHER_ASSET":     2,
}

func (x AssetType) String() string {
	return proto.EnumName(AssetType_name, int32(x))
}

func (AssetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_545b8d23845e9518, []int{0}
}

type AssetStatus int32

const (
	AssetStatus_PENDING   AssetStatus = 0
	AssetStatus_VALIDATED AssetStatus = 1
	AssetStatus_INVALID   AssetStatus = 2
)

var AssetStatus_name = map[int32]string{
	0: "PENDING",
	1: "VALIDATED",
	2: "INVALID",
}

var AssetStatus_value = map[string]int32{
	"PENDING":   0,
	"VALIDATED": 1,
	"INVALID":   2,
}

func (x AssetStatus) String() string {
	return proto.EnumName(AssetStatus_name, int32(x))
}

func (AssetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_545b8d23845e9518, []int{1}
}

type ValueValidation struct {
	Validators []string `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	ValueUsd   string   `protobuf:"bytes,2,opt,name=value_usd,json=valueUsd,proto3" json:"value_usd,omitempty"`
	Timestamp  uint64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *ValueValidation) Reset()         { *m = ValueValidation{} }
func (m *ValueValidation) String() string { return proto.CompactTextString(m) }
func (*ValueValidation) ProtoMessage()    {}
func (*ValueValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_545b8d23845e9518, []int{0}
}
func (m *ValueValidation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValueValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValueValidation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValueValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueValidation.Merge(m, src)
}
func (m *ValueValidation) XXX_Size() int {
	return m.Size()
}
func (m *ValueValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueValidation.DiscardUnknown(m)
}

var xxx_messageInfo_ValueValidation proto.InternalMessageInfo

func (m *ValueValidation) GetValidators() []string {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *ValueValidation) GetValueUsd() string {
	if m != nil {
		return m.ValueUsd
	}
	return ""
}

func (m *ValueValidation) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Asset struct {
	Id           uint64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AssetType    AssetType        `protobuf:"varint,3,opt,name=asset_type,json=assetType,proto3,enum=optimizeglobalcoin.asset.AssetType" json:"asset_type,omitempty"`
	Jurisdiction string           `protobuf:"bytes,4,opt,name=jurisdiction,proto3" json:"jurisdiction,omitempty"`
	Owner        string           `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Status       AssetStatus      `protobuf:"varint,6,opt,name=status,proto3,enum=optimizeglobalcoin.asset.AssetStatus" json:"status,omitempty"`
	Value        *ValueValidation `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	Options      string           `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
	Timestamp    uint64           `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_545b8d23845e9518, []int{1}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return m.Size()
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetAssetType() AssetType {
	if m != nil {
		return m.AssetType
	}
	return AssetType_REAL_ESTATE
}

func (m *Asset) GetJurisdiction() string {
	if m != nil {
		return m.Jurisdiction
	}
	return ""
}

func (m *Asset) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Asset) GetStatus() AssetStatus {
	if m != nil {
		return m.Status
	}
	return AssetStatus_PENDING
}

func (m *Asset) GetValue() *ValueValidation {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Asset) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *Asset) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("optimizeglobalcoin.asset.AssetType", AssetType_name, AssetType_value)
	proto.RegisterEnum("optimizeglobalcoin.asset.AssetStatus", AssetStatus_name, AssetStatus_value)
	proto.RegisterType((*ValueValidation)(nil), "optimizeglobalcoin.asset.ValueValidation")
	proto.RegisterType((*Asset)(nil), "optimizeglobalcoin.asset.Asset")
}

func init() {
	proto.RegisterFile("optimizeglobalcoin/asset/asset.proto", fileDescriptor_545b8d23845e9518)
}

var fileDescriptor_545b8d23845e9518 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xf5, 0xba, 0xf9, 0xa8, 0x27, 0xd0, 0x58, 0x03, 0x87, 0x95, 0x40, 0x96, 0x15, 0x40, 0x0a,
	0x3d, 0x04, 0xa9, 0x48, 0x1c, 0x90, 0x10, 0x72, 0x88, 0x01, 0x4b, 0x91, 0x41, 0x1b, 0x93, 0x03,
	0x97, 0x68, 0x5b, 0xaf, 0xd0, 0x22, 0xc7, 0x6b, 0x65, 0x37, 0x85, 0xf2, 0x2b, 0xf8, 0x59, 0x1c,
	0x7b, 0xe4, 0x08, 0xc9, 0x1f, 0x41, 0xde, 0xa4, 0xd0, 0x16, 0x22, 0x2e, 0xd6, 0xce, 0x9b, 0x79,
	0x6f, 0x66, 0x9e, 0x07, 0xee, 0xab, 0xca, 0xc8, 0xb9, 0xfc, 0x22, 0x3e, 0x14, 0xea, 0x98, 0x17,
	0x27, 0x4a, 0x96, 0x8f, 0xb8, 0xd6, 0xc2, 0x6c, 0xbe, 0x83, 0x6a, 0xa1, 0x8c, 0x42, 0xfa, 0x77,
	0xd5, 0xc0, 0xe6, 0x7b, 0x05, 0x74, 0xa7, 0xbc, 0x58, 0x8a, 0x29, 0x2f, 0x64, 0xce, 0x8d, 0x54,
	0x25, 0x06, 0x00, 0xa7, 0x9b, 0x48, 0x2d, 0x34, 0x25, 0xe1, 0x5e, 0xdf, 0x63, 0x97, 0x10, 0xbc,
	0x03, 0xde, 0x69, 0x4d, 0x99, 0x2d, 0x75, 0x4e, 0xdd, 0x90, 0xf4, 0x3d, 0xb6, 0x6f, 0x81, 0x77,
	0x3a, 0xc7, 0xbb, 0xe0, 0x19, 0x39, 0x17, 0xda, 0xf0, 0x79, 0x45, 0xf7, 0x42, 0xd2, 0x6f, 0xb0,
	0x3f, 0x40, 0xef, 0xa7, 0x0b, 0xcd, 0xa8, 0xee, 0x8b, 0x07, 0xe0, 0xca, 0x9c, 0x12, 0x5b, 0xe0,
	0xca, 0x1c, 0x11, 0x1a, 0x25, 0x9f, 0x8b, 0xad, 0x9e, 0x7d, 0xe3, 0x10, 0xc0, 0x0e, 0x39, 0x33,
	0x67, 0x95, 0xb0, 0x62, 0x07, 0x47, 0xf7, 0x06, 0xbb, 0x56, 0x19, 0x58, 0xe1, 0xec, 0xac, 0x12,
	0xcc, 0xe3, 0x17, 0x4f, 0xec, 0xc1, 0x8d, 0x8f, 0xcb, 0x85, 0xd4, 0xb9, 0x3c, 0xa9, 0x97, 0xa3,
	0x0d, 0xab, 0x7f, 0x05, 0xc3, 0xdb, 0xd0, 0x54, 0x9f, 0x4a, 0xb1, 0xa0, 0x4d, 0x9b, 0xdc, 0x04,
	0xf8, 0x0c, 0x5a, 0xda, 0x70, 0xb3, 0xd4, 0xb4, 0x65, 0x3b, 0x3f, 0xf8, 0x4f, 0xe7, 0x89, 0x2d,
	0x66, 0x5b, 0x12, 0x3e, 0x87, 0xa6, 0x35, 0x85, 0xb6, 0x43, 0xd2, 0xef, 0x1c, 0x3d, 0xdc, 0xcd,
	0xbe, 0xe6, 0x3f, 0xdb, 0xf0, 0x90, 0x42, 0xbb, 0xa6, 0xa8, 0x52, 0xd3, 0x7d, 0x3b, 0xd7, 0x45,
	0x78, 0xd5, 0x63, 0xef, 0x9a, 0xc7, 0x87, 0x43, 0xf0, 0x7e, 0x3b, 0x81, 0x5d, 0xe8, 0xb0, 0x38,
	0x1a, 0xcf, 0xe2, 0x49, 0x16, 0x65, 0xb1, 0xef, 0xe0, 0x2d, 0xe8, 0xbe, 0x4c, 0xd2, 0x28, 0x7d,
	0x91, 0x44, 0xe3, 0x59, 0x34, 0x99, 0xc4, 0x99, 0x4f, 0xea, 0xaa, 0x37, 0xd9, 0xeb, 0x98, 0x6d,
	0x01, 0xf7, 0xf0, 0x09, 0x74, 0x2e, 0xed, 0x84, 0x1d, 0x68, 0xbf, 0x8d, 0xd3, 0x51, 0x92, 0xbe,
	0xf2, 0x1d, 0xbc, 0x09, 0xde, 0x34, 0x1a, 0x27, 0xa3, 0x28, 0x8b, 0x47, 0x3e, 0xa9, 0x73, 0x49,
	0x6a, 0x01, 0xdf, 0x1d, 0x3e, 0xfd, 0xb6, 0x0a, 0xc8, 0xf9, 0x2a, 0x20, 0x3f, 0x56, 0x01, 0xf9,
	0xba, 0x0e, 0x9c, 0xf3, 0x75, 0xe0, 0x7c, 0x5f, 0x07, 0xce, 0xfb, 0xf0, 0x1f, 0x77, 0xfa, 0x79,
	0x7b, 0xa9, 0xf5, 0xef, 0xd5, 0xc7, 0x2d, 0x7b, 0xaa, 0x8f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0x26, 0x01, 0x92, 0xd2, 0x02, 0x00, 0x00,
}

func (m *ValueValidation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValueValidation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValueValidation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ValueUsd) > 0 {
		i -= len(m.ValueUsd)
		copy(dAtA[i:], m.ValueUsd)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.ValueUsd)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Validators[iNdEx])
			copy(dAtA[i:], m.Validators[iNdEx])
			i = encodeVarintAsset(dAtA, i, uint64(len(m.Validators[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Asset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Options) > 0 {
		i -= len(m.Options)
		copy(dAtA[i:], m.Options)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Options)))
		i--
		dAtA[i] = 0x42
	}
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAsset(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Status != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Jurisdiction) > 0 {
		i -= len(m.Jurisdiction)
		copy(dAtA[i:], m.Jurisdiction)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Jurisdiction)))
		i--
		dAtA[i] = 0x22
	}
	if m.AssetType != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.AssetType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAsset(dAtA []byte, offset int, v uint64) int {
	offset -= sovAsset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValueValidation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, s := range m.Validators {
			l = len(s)
			n += 1 + l + sovAsset(uint64(l))
		}
	}
	l = len(m.ValueUsd)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovAsset(uint64(m.Timestamp))
	}
	return n
}

func (m *Asset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAsset(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if m.AssetType != 0 {
		n += 1 + sovAsset(uint64(m.AssetType))
	}
	l = len(m.Jurisdiction)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovAsset(uint64(m.Status))
	}
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Options)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovAsset(uint64(m.Timestamp))
	}
	return n
}

func sovAsset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAsset(x uint64) (n int) {
	return sovAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValueValidation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: ValueValidation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValueValidation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueUsd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueUsd = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAsset
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
func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetType", wireType)
			}
			m.AssetType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetType |= AssetType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jurisdiction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jurisdiction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= AssetStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &ValueValidation{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAsset
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
func skipAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
				return 0, ErrInvalidLengthAsset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAsset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAsset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAsset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAsset = fmt.Errorf("proto: unexpected end of group")
)
