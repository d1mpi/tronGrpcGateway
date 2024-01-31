// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/contract/asset_issue_contract.proto

package core

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

type AssetIssueContract struct {
	Id                      string                             `protobuf:"bytes,41,opt,name=id,proto3" json:"id,omitempty"`
	OwnerAddress            []byte                             `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Name                    []byte                             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbr                    []byte                             `protobuf:"bytes,3,opt,name=abbr,proto3" json:"abbr,omitempty"`
	TotalSupply             int64                              `protobuf:"varint,4,opt,name=total_supply,json=totalSupply,proto3" json:"total_supply,omitempty"`
	FrozenSupply            []*AssetIssueContract_FrozenSupply `protobuf:"bytes,5,rep,name=frozen_supply,json=frozenSupply,proto3" json:"frozen_supply,omitempty"`
	TrxNum                  int32                              `protobuf:"varint,6,opt,name=trx_num,json=trxNum,proto3" json:"trx_num,omitempty"`
	Precision               int32                              `protobuf:"varint,7,opt,name=precision,proto3" json:"precision,omitempty"`
	Num                     int32                              `protobuf:"varint,8,opt,name=num,proto3" json:"num,omitempty"`
	StartTime               int64                              `protobuf:"varint,9,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime                 int64                              `protobuf:"varint,10,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Order                   int64                              `protobuf:"varint,11,opt,name=order,proto3" json:"order,omitempty"`
	VoteScore               int32                              `protobuf:"varint,16,opt,name=vote_score,json=voteScore,proto3" json:"vote_score,omitempty"`
	Description             []byte                             `protobuf:"bytes,20,opt,name=description,proto3" json:"description,omitempty"`
	Url                     []byte                             `protobuf:"bytes,21,opt,name=url,proto3" json:"url,omitempty"`
	FreeAssetNetLimit       int64                              `protobuf:"varint,22,opt,name=free_asset_net_limit,json=freeAssetNetLimit,proto3" json:"free_asset_net_limit,omitempty"`
	PublicFreeAssetNetLimit int64                              `protobuf:"varint,23,opt,name=public_free_asset_net_limit,json=publicFreeAssetNetLimit,proto3" json:"public_free_asset_net_limit,omitempty"`
	PublicFreeAssetNetUsage int64                              `protobuf:"varint,24,opt,name=public_free_asset_net_usage,json=publicFreeAssetNetUsage,proto3" json:"public_free_asset_net_usage,omitempty"`
	PublicLatestFreeNetTime int64                              `protobuf:"varint,25,opt,name=public_latest_free_net_time,json=publicLatestFreeNetTime,proto3" json:"public_latest_free_net_time,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                           `json:"-"`
	XXX_unrecognized        []byte                             `json:"-"`
	XXX_sizecache           int32                              `json:"-"`
}

func (m *AssetIssueContract) Reset()         { *m = AssetIssueContract{} }
func (m *AssetIssueContract) String() string { return proto.CompactTextString(m) }
func (*AssetIssueContract) ProtoMessage()    {}
func (*AssetIssueContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a53d522aa45c6a0, []int{0}
}

func (m *AssetIssueContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetIssueContract.Unmarshal(m, b)
}
func (m *AssetIssueContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetIssueContract.Marshal(b, m, deterministic)
}
func (m *AssetIssueContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetIssueContract.Merge(m, src)
}
func (m *AssetIssueContract) XXX_Size() int {
	return xxx_messageInfo_AssetIssueContract.Size(m)
}
func (m *AssetIssueContract) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetIssueContract.DiscardUnknown(m)
}

var xxx_messageInfo_AssetIssueContract proto.InternalMessageInfo

func (m *AssetIssueContract) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssetIssueContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *AssetIssueContract) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AssetIssueContract) GetAbbr() []byte {
	if m != nil {
		return m.Abbr
	}
	return nil
}

func (m *AssetIssueContract) GetTotalSupply() int64 {
	if m != nil {
		return m.TotalSupply
	}
	return 0
}

func (m *AssetIssueContract) GetFrozenSupply() []*AssetIssueContract_FrozenSupply {
	if m != nil {
		return m.FrozenSupply
	}
	return nil
}

func (m *AssetIssueContract) GetTrxNum() int32 {
	if m != nil {
		return m.TrxNum
	}
	return 0
}

func (m *AssetIssueContract) GetPrecision() int32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *AssetIssueContract) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *AssetIssueContract) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AssetIssueContract) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *AssetIssueContract) GetOrder() int64 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *AssetIssueContract) GetVoteScore() int32 {
	if m != nil {
		return m.VoteScore
	}
	return 0
}

func (m *AssetIssueContract) GetDescription() []byte {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *AssetIssueContract) GetUrl() []byte {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *AssetIssueContract) GetFreeAssetNetLimit() int64 {
	if m != nil {
		return m.FreeAssetNetLimit
	}
	return 0
}

func (m *AssetIssueContract) GetPublicFreeAssetNetLimit() int64 {
	if m != nil {
		return m.PublicFreeAssetNetLimit
	}
	return 0
}

func (m *AssetIssueContract) GetPublicFreeAssetNetUsage() int64 {
	if m != nil {
		return m.PublicFreeAssetNetUsage
	}
	return 0
}

func (m *AssetIssueContract) GetPublicLatestFreeNetTime() int64 {
	if m != nil {
		return m.PublicLatestFreeNetTime
	}
	return 0
}

type AssetIssueContract_FrozenSupply struct {
	FrozenAmount         int64    `protobuf:"varint,1,opt,name=frozen_amount,json=frozenAmount,proto3" json:"frozen_amount,omitempty"`
	FrozenDays           int64    `protobuf:"varint,2,opt,name=frozen_days,json=frozenDays,proto3" json:"frozen_days,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetIssueContract_FrozenSupply) Reset()         { *m = AssetIssueContract_FrozenSupply{} }
func (m *AssetIssueContract_FrozenSupply) String() string { return proto.CompactTextString(m) }
func (*AssetIssueContract_FrozenSupply) ProtoMessage()    {}
func (*AssetIssueContract_FrozenSupply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a53d522aa45c6a0, []int{0, 0}
}

func (m *AssetIssueContract_FrozenSupply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetIssueContract_FrozenSupply.Unmarshal(m, b)
}
func (m *AssetIssueContract_FrozenSupply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetIssueContract_FrozenSupply.Marshal(b, m, deterministic)
}
func (m *AssetIssueContract_FrozenSupply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetIssueContract_FrozenSupply.Merge(m, src)
}
func (m *AssetIssueContract_FrozenSupply) XXX_Size() int {
	return xxx_messageInfo_AssetIssueContract_FrozenSupply.Size(m)
}
func (m *AssetIssueContract_FrozenSupply) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetIssueContract_FrozenSupply.DiscardUnknown(m)
}

var xxx_messageInfo_AssetIssueContract_FrozenSupply proto.InternalMessageInfo

func (m *AssetIssueContract_FrozenSupply) GetFrozenAmount() int64 {
	if m != nil {
		return m.FrozenAmount
	}
	return 0
}

func (m *AssetIssueContract_FrozenSupply) GetFrozenDays() int64 {
	if m != nil {
		return m.FrozenDays
	}
	return 0
}

type TransferAssetContract struct {
	AssetName            []byte   `protobuf:"bytes,1,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"`
	OwnerAddress         []byte   `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress            []byte   `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount               int64    `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferAssetContract) Reset()         { *m = TransferAssetContract{} }
func (m *TransferAssetContract) String() string { return proto.CompactTextString(m) }
func (*TransferAssetContract) ProtoMessage()    {}
func (*TransferAssetContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a53d522aa45c6a0, []int{1}
}

func (m *TransferAssetContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferAssetContract.Unmarshal(m, b)
}
func (m *TransferAssetContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferAssetContract.Marshal(b, m, deterministic)
}
func (m *TransferAssetContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferAssetContract.Merge(m, src)
}
func (m *TransferAssetContract) XXX_Size() int {
	return xxx_messageInfo_TransferAssetContract.Size(m)
}
func (m *TransferAssetContract) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferAssetContract.DiscardUnknown(m)
}

var xxx_messageInfo_TransferAssetContract proto.InternalMessageInfo

func (m *TransferAssetContract) GetAssetName() []byte {
	if m != nil {
		return m.AssetName
	}
	return nil
}

func (m *TransferAssetContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *TransferAssetContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *TransferAssetContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type UnfreezeAssetContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnfreezeAssetContract) Reset()         { *m = UnfreezeAssetContract{} }
func (m *UnfreezeAssetContract) String() string { return proto.CompactTextString(m) }
func (*UnfreezeAssetContract) ProtoMessage()    {}
func (*UnfreezeAssetContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a53d522aa45c6a0, []int{2}
}

func (m *UnfreezeAssetContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnfreezeAssetContract.Unmarshal(m, b)
}
func (m *UnfreezeAssetContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnfreezeAssetContract.Marshal(b, m, deterministic)
}
func (m *UnfreezeAssetContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnfreezeAssetContract.Merge(m, src)
}
func (m *UnfreezeAssetContract) XXX_Size() int {
	return xxx_messageInfo_UnfreezeAssetContract.Size(m)
}
func (m *UnfreezeAssetContract) XXX_DiscardUnknown() {
	xxx_messageInfo_UnfreezeAssetContract.DiscardUnknown(m)
}

var xxx_messageInfo_UnfreezeAssetContract proto.InternalMessageInfo

func (m *UnfreezeAssetContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

type UpdateAssetContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Description          []byte   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url                  []byte   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	NewLimit             int64    `protobuf:"varint,4,opt,name=new_limit,json=newLimit,proto3" json:"new_limit,omitempty"`
	NewPublicLimit       int64    `protobuf:"varint,5,opt,name=new_public_limit,json=newPublicLimit,proto3" json:"new_public_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAssetContract) Reset()         { *m = UpdateAssetContract{} }
func (m *UpdateAssetContract) String() string { return proto.CompactTextString(m) }
func (*UpdateAssetContract) ProtoMessage()    {}
func (*UpdateAssetContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a53d522aa45c6a0, []int{3}
}

func (m *UpdateAssetContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAssetContract.Unmarshal(m, b)
}
func (m *UpdateAssetContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAssetContract.Marshal(b, m, deterministic)
}
func (m *UpdateAssetContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAssetContract.Merge(m, src)
}
func (m *UpdateAssetContract) XXX_Size() int {
	return xxx_messageInfo_UpdateAssetContract.Size(m)
}
func (m *UpdateAssetContract) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAssetContract.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAssetContract proto.InternalMessageInfo

func (m *UpdateAssetContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *UpdateAssetContract) GetDescription() []byte {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *UpdateAssetContract) GetUrl() []byte {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *UpdateAssetContract) GetNewLimit() int64 {
	if m != nil {
		return m.NewLimit
	}
	return 0
}

func (m *UpdateAssetContract) GetNewPublicLimit() int64 {
	if m != nil {
		return m.NewPublicLimit
	}
	return 0
}

type ParticipateAssetIssueContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress            []byte   `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	AssetName            []byte   `protobuf:"bytes,3,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"`
	Amount               int64    `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParticipateAssetIssueContract) Reset()         { *m = ParticipateAssetIssueContract{} }
func (m *ParticipateAssetIssueContract) String() string { return proto.CompactTextString(m) }
func (*ParticipateAssetIssueContract) ProtoMessage()    {}
func (*ParticipateAssetIssueContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a53d522aa45c6a0, []int{4}
}

func (m *ParticipateAssetIssueContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParticipateAssetIssueContract.Unmarshal(m, b)
}
func (m *ParticipateAssetIssueContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParticipateAssetIssueContract.Marshal(b, m, deterministic)
}
func (m *ParticipateAssetIssueContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParticipateAssetIssueContract.Merge(m, src)
}
func (m *ParticipateAssetIssueContract) XXX_Size() int {
	return xxx_messageInfo_ParticipateAssetIssueContract.Size(m)
}
func (m *ParticipateAssetIssueContract) XXX_DiscardUnknown() {
	xxx_messageInfo_ParticipateAssetIssueContract.DiscardUnknown(m)
}

var xxx_messageInfo_ParticipateAssetIssueContract proto.InternalMessageInfo

func (m *ParticipateAssetIssueContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *ParticipateAssetIssueContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *ParticipateAssetIssueContract) GetAssetName() []byte {
	if m != nil {
		return m.AssetName
	}
	return nil
}

func (m *ParticipateAssetIssueContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*AssetIssueContract)(nil), "protocol.AssetIssueContract")
	proto.RegisterType((*AssetIssueContract_FrozenSupply)(nil), "protocol.AssetIssueContract.FrozenSupply")
	proto.RegisterType((*TransferAssetContract)(nil), "protocol.TransferAssetContract")
	proto.RegisterType((*UnfreezeAssetContract)(nil), "protocol.UnfreezeAssetContract")
	proto.RegisterType((*UpdateAssetContract)(nil), "protocol.UpdateAssetContract")
	proto.RegisterType((*ParticipateAssetIssueContract)(nil), "protocol.ParticipateAssetIssueContract")
}

func init() {
	proto.RegisterFile("core/contract/asset_issue_contract.proto", fileDescriptor_7a53d522aa45c6a0)
}

var fileDescriptor_7a53d522aa45c6a0 = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x4e, 0x1b, 0x3b,
	0x14, 0xd6, 0x64, 0x48, 0x48, 0x4e, 0x02, 0xe2, 0xfa, 0xf2, 0x63, 0x2e, 0x37, 0xba, 0xb9, 0xe9,
	0xa2, 0x61, 0x41, 0x42, 0xdb, 0x2d, 0x1b, 0x4a, 0x85, 0x54, 0x09, 0x45, 0x28, 0xc0, 0xa6, 0x9b,
	0x91, 0x33, 0xe3, 0xa4, 0x96, 0x32, 0xf6, 0xc8, 0xf6, 0x34, 0x84, 0xb7, 0x68, 0x97, 0x7d, 0x8f,
	0xbe, 0x5f, 0xe5, 0xe3, 0x19, 0x48, 0xf8, 0x91, 0x50, 0x57, 0xf1, 0x7c, 0x3f, 0xce, 0x99, 0x73,
	0xbe, 0x33, 0xd0, 0x8b, 0x95, 0xe6, 0x83, 0x58, 0x49, 0xab, 0x59, 0x6c, 0x07, 0xcc, 0x18, 0x6e,
	0x23, 0x61, 0x4c, 0xce, 0xa3, 0x12, 0xec, 0x67, 0x5a, 0x59, 0x45, 0xea, 0xf8, 0x13, 0xab, 0x59,
	0xf7, 0x7b, 0x0d, 0xc8, 0xa9, 0x13, 0x7e, 0x76, 0xba, 0xb3, 0x42, 0x46, 0x36, 0xa1, 0x22, 0x12,
	0x7a, 0xd8, 0x09, 0x7a, 0x8d, 0x51, 0x45, 0x24, 0xe4, 0x0d, 0x6c, 0xa8, 0xb9, 0xe4, 0x3a, 0x62,
	0x49, 0xa2, 0xb9, 0x31, 0x34, 0xe8, 0x04, 0xbd, 0xd6, 0xa8, 0x85, 0xe0, 0xa9, 0xc7, 0x08, 0x81,
	0x35, 0xc9, 0x52, 0x4e, 0x2b, 0xc8, 0xe1, 0xd9, 0x61, 0x6c, 0x3c, 0xd6, 0x34, 0xf4, 0x98, 0x3b,
	0x93, 0xff, 0xa1, 0x65, 0x95, 0x65, 0xb3, 0xc8, 0xe4, 0x59, 0x36, 0x5b, 0xd0, 0xb5, 0x4e, 0xd0,
	0x0b, 0x47, 0x4d, 0xc4, 0xae, 0x10, 0x22, 0x43, 0xd8, 0x98, 0x68, 0x75, 0xc7, 0x65, 0xa9, 0xa9,
	0x76, 0xc2, 0x5e, 0xf3, 0xfd, 0x61, 0xbf, 0x2c, 0xbc, 0xff, 0xb4, 0xe8, 0xfe, 0x39, 0x3a, 0xfc,
	0x0d, 0xa3, 0xd6, 0x64, 0xe9, 0x89, 0xec, 0xc1, 0xba, 0xd5, 0xb7, 0x91, 0xcc, 0x53, 0x5a, 0xeb,
	0x04, 0xbd, 0xea, 0xa8, 0x66, 0xf5, 0xed, 0x30, 0x4f, 0xc9, 0xbf, 0xd0, 0xc8, 0x34, 0x8f, 0x85,
	0x11, 0x4a, 0xd2, 0x75, 0xa4, 0x1e, 0x00, 0xb2, 0x05, 0xa1, 0xb3, 0xd4, 0x11, 0x77, 0x47, 0xd2,
	0x06, 0x30, 0x96, 0x69, 0x1b, 0x59, 0x91, 0x72, 0xda, 0xc0, 0xca, 0x1b, 0x88, 0x5c, 0x8b, 0x94,
	0x93, 0x7d, 0xa8, 0x73, 0x99, 0x78, 0x12, 0x90, 0x5c, 0xe7, 0x32, 0x41, 0x6a, 0x1b, 0xaa, 0x4a,
	0x27, 0x5c, 0xd3, 0x26, 0xe2, 0xfe, 0xc1, 0xdd, 0xf7, 0x4d, 0x59, 0x1e, 0x19, 0x37, 0x3b, 0xba,
	0xe5, 0x0b, 0x70, 0xc8, 0x95, 0x03, 0x48, 0x07, 0x9a, 0x09, 0x37, 0xb1, 0x16, 0x99, 0x75, 0x05,
	0x6e, 0x63, 0x17, 0x97, 0x21, 0x57, 0x62, 0xae, 0x67, 0x74, 0x07, 0x19, 0x77, 0x24, 0x03, 0xd8,
	0x9e, 0x68, 0xce, 0x23, 0x3f, 0x7f, 0xc9, 0x6d, 0x34, 0x13, 0xa9, 0xb0, 0x74, 0x17, 0xff, 0xf7,
	0x2f, 0xc7, 0x61, 0xf3, 0x86, 0xdc, 0x5e, 0x38, 0x82, 0x9c, 0xc0, 0x41, 0x96, 0x8f, 0x67, 0x22,
	0x8e, 0x9e, 0xf5, 0xed, 0xa1, 0x6f, 0xcf, 0x4b, 0xce, 0x5f, 0xef, 0xce, 0x0d, 0x9b, 0x72, 0x4a,
	0x5f, 0x72, 0xdf, 0x38, 0x7a, 0xc9, 0x3d, 0x63, 0x96, 0x1b, 0xeb, 0x2f, 0x71, 0x76, 0xec, 0xe1,
	0xfe, 0xb2, 0xfb, 0x02, 0x15, 0xee, 0x8e, 0x21, 0xc7, 0x76, 0xff, 0x73, 0x0d, 0xad, 0xe5, 0xa1,
	0xbb, 0x98, 0x16, 0xb1, 0x61, 0xa9, 0xca, 0xa5, 0xc5, 0x98, 0x86, 0x65, 0x16, 0x4e, 0x11, 0x23,
	0xff, 0x41, 0xb3, 0x10, 0x25, 0x6c, 0x61, 0x30, 0xad, 0xe1, 0x08, 0x3c, 0xf4, 0x89, 0x2d, 0x4c,
	0xf7, 0x47, 0x00, 0x3b, 0xd7, 0x9a, 0x49, 0x33, 0xe1, 0x1a, 0xab, 0xbd, 0x5f, 0x8b, 0x36, 0x40,
	0xf1, 0x7e, 0x2e, 0xe7, 0x7e, 0x07, 0x1a, 0x88, 0x0c, 0x5d, 0xd8, 0x9f, 0x6c, 0x49, 0xe5, 0x99,
	0x2d, 0x69, 0x03, 0x58, 0x75, 0xaf, 0xf0, 0x7b, 0xd1, 0xb0, 0xaa, 0xa4, 0x77, 0xa1, 0x56, 0xd4,
	0xee, 0xd7, 0xa2, 0x78, 0xea, 0x9e, 0xc0, 0xce, 0x8d, 0x74, 0xcd, 0xb9, 0xe3, 0xab, 0x35, 0xbd,
	0x66, 0x35, 0xbb, 0xbf, 0x02, 0xf8, 0xfb, 0x26, 0x4b, 0x98, 0xfd, 0x03, 0xf3, 0xe3, 0x10, 0x56,
	0x5e, 0x0c, 0x61, 0xf8, 0x10, 0xc2, 0x03, 0x68, 0x48, 0x3e, 0x2f, 0x12, 0xe4, 0xdf, 0xa4, 0x2e,
	0xf9, 0xdc, 0x47, 0xa6, 0x07, 0x5b, 0x8e, 0x2c, 0x07, 0x8f, 0x9a, 0x2a, 0x6a, 0x36, 0x25, 0x9f,
	0x5f, 0xfa, 0x61, 0x3b, 0xb4, 0xfb, 0x33, 0x80, 0xf6, 0x25, 0xd3, 0x56, 0xc4, 0x22, 0x2b, 0x8b,
	0x5f, 0xfd, 0x52, 0xbd, 0xea, 0x0d, 0x56, 0x7b, 0x5e, 0x79, 0xdc, 0xf3, 0xd5, 0xb1, 0x86, 0x8f,
	0xc7, 0xfa, 0xc2, 0x48, 0x3e, 0x9e, 0x01, 0x55, 0x7a, 0xda, 0xb7, 0x5a, 0x49, 0xff, 0x5d, 0x32,
	0xfd, 0xf2, 0x3b, 0xfb, 0xe5, 0xed, 0x54, 0xd8, 0xaf, 0xf9, 0xb8, 0x1f, 0xab, 0x74, 0x70, 0x7c,
	0xfb, 0xee, 0x78, 0x32, 0x70, 0xb2, 0xa3, 0xa9, 0xce, 0xe2, 0xa3, 0x29, 0xb3, 0x7c, 0xce, 0x16,
	0x03, 0xb7, 0xe1, 0xe3, 0x1a, 0x3a, 0x3f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x35, 0x37, 0xe9,
	0xc9, 0xbd, 0x05, 0x00, 0x00,
}
