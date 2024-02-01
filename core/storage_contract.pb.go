// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: protocol/core/contract/storage_contract.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BuyStorageBytesContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Bytes        int64  `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"` // storage bytes for buy
}

func (x *BuyStorageBytesContract) Reset() {
	*x = BuyStorageBytesContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_storage_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyStorageBytesContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyStorageBytesContract) ProtoMessage() {}

func (x *BuyStorageBytesContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_storage_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyStorageBytesContract.ProtoReflect.Descriptor instead.
func (*BuyStorageBytesContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_storage_contract_proto_rawDescGZIP(), []int{0}
}

func (x *BuyStorageBytesContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *BuyStorageBytesContract) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

type BuyStorageContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Quant        int64  `protobuf:"varint,2,opt,name=quant,proto3" json:"quant,omitempty"` // trx quantity for buy storage (sun)
}

func (x *BuyStorageContract) Reset() {
	*x = BuyStorageContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_storage_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyStorageContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyStorageContract) ProtoMessage() {}

func (x *BuyStorageContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_storage_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyStorageContract.ProtoReflect.Descriptor instead.
func (*BuyStorageContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_storage_contract_proto_rawDescGZIP(), []int{1}
}

func (x *BuyStorageContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *BuyStorageContract) GetQuant() int64 {
	if x != nil {
		return x.Quant
	}
	return 0
}

type SellStorageContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	StorageBytes int64  `protobuf:"varint,2,opt,name=storage_bytes,json=storageBytes,proto3" json:"storage_bytes,omitempty"`
}

func (x *SellStorageContract) Reset() {
	*x = SellStorageContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_storage_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellStorageContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellStorageContract) ProtoMessage() {}

func (x *SellStorageContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_storage_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellStorageContract.ProtoReflect.Descriptor instead.
func (*SellStorageContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_storage_contract_proto_rawDescGZIP(), []int{2}
}

func (x *SellStorageContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *SellStorageContract) GetStorageBytes() int64 {
	if x != nil {
		return x.StorageBytes
	}
	return 0
}

type UpdateBrokerageContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Brokerage    int32  `protobuf:"varint,2,opt,name=brokerage,proto3" json:"brokerage,omitempty"` // 1 mean 1%
}

func (x *UpdateBrokerageContract) Reset() {
	*x = UpdateBrokerageContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_storage_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBrokerageContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBrokerageContract) ProtoMessage() {}

func (x *UpdateBrokerageContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_storage_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBrokerageContract.ProtoReflect.Descriptor instead.
func (*UpdateBrokerageContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_storage_contract_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBrokerageContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *UpdateBrokerageContract) GetBrokerage() int32 {
	if x != nil {
		return x.Brokerage
	}
	return 0
}

var File_protocol_core_contract_storage_contract_proto protoreflect.FileDescriptor

var file_protocol_core_contract_storage_contract_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x54, 0x0a, 0x17, 0x42, 0x75, 0x79,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22,
	0x4f, 0x0a, 0x12, 0x42, 0x75, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x22, 0x5f, 0x0a, 0x13, 0x53, 0x65, 0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x22, 0x5c, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x42,
	0x45, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_core_contract_storage_contract_proto_rawDescOnce sync.Once
	file_protocol_core_contract_storage_contract_proto_rawDescData = file_protocol_core_contract_storage_contract_proto_rawDesc
)

func file_protocol_core_contract_storage_contract_proto_rawDescGZIP() []byte {
	file_protocol_core_contract_storage_contract_proto_rawDescOnce.Do(func() {
		file_protocol_core_contract_storage_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_core_contract_storage_contract_proto_rawDescData)
	})
	return file_protocol_core_contract_storage_contract_proto_rawDescData
}

var file_protocol_core_contract_storage_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_core_contract_storage_contract_proto_goTypes = []interface{}{
	(*BuyStorageBytesContract)(nil), // 0: protocol.BuyStorageBytesContract
	(*BuyStorageContract)(nil),      // 1: protocol.BuyStorageContract
	(*SellStorageContract)(nil),     // 2: protocol.SellStorageContract
	(*UpdateBrokerageContract)(nil), // 3: protocol.UpdateBrokerageContract
}
var file_protocol_core_contract_storage_contract_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_core_contract_storage_contract_proto_init() }
func file_protocol_core_contract_storage_contract_proto_init() {
	if File_protocol_core_contract_storage_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_core_contract_storage_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyStorageBytesContract); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protocol_core_contract_storage_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyStorageContract); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protocol_core_contract_storage_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellStorageContract); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protocol_core_contract_storage_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBrokerageContract); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protocol_core_contract_storage_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_core_contract_storage_contract_proto_goTypes,
		DependencyIndexes: file_protocol_core_contract_storage_contract_proto_depIdxs,
		MessageInfos:      file_protocol_core_contract_storage_contract_proto_msgTypes,
	}.Build()
	File_protocol_core_contract_storage_contract_proto = out.File
	file_protocol_core_contract_storage_contract_proto_rawDesc = nil
	file_protocol_core_contract_storage_contract_proto_goTypes = nil
	file_protocol_core_contract_storage_contract_proto_depIdxs = nil
}
