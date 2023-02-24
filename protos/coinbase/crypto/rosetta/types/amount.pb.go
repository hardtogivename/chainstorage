// (-- api-linter: core::0140::abbreviations=disabled
//     aip.dev/not-precedent: Matching the naming convention of rosetta --)
// (-- api-linter: core::0143::standardized-codes=disabled
//     aip.dev/not-precedent: Matching the naming convention of rosetta --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.21.12
// source: coinbase/crypto/rosetta/types/amount.proto

// The stable release for rosetta types

package types

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Amount is some Value of a Currency. It is considered invalid to specify a Value without a
// Currency.
type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value of the transaction in atomic units represented as an arbitrary-sized signed integer.
	// For example, 1 BTC would be represented by a value of 100000000.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// Currency is composed of a canonical Symbol and Decimals. This Decimals value is used to
	// convert an Amount.Value from atomic units (Satoshis) to standard units (Bitcoins).
	Currency *Currency `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	// The metadata is protocol specific metadata
	Metadata map[string]*anypb.Any `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_amount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_amount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_amount_proto_rawDescGZIP(), []int{0}
}

func (x *Amount) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Amount) GetCurrency() *Currency {
	if x != nil {
		return x.Currency
	}
	return nil
}

func (x *Amount) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Currency is composed of a canonical Symbol and Decimals. This Decimals value is used to convert
// an Amount.Value from atomic units (Satoshis) to standard units (Bitcoins).
type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Canonical symbol associated with a currency.
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Number of decimal places in the standard unit representation of the amount. For example, BTC
	// has 8 decimals. Note that it is not possible to represent the value of some currency in atomic
	// units that is not base 10.
	Decimals int32 `protobuf:"varint,2,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// Any additional information related to the currency itself. For example, it would be useful to
	// populate this object with the contract address of an ERC-20 token.
	Metadata map[string]*anypb.Any `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_amount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_amount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_amount_proto_rawDescGZIP(), []int{1}
}

func (x *Currency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Currency) GetDecimals() int32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *Currency) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_coinbase_crypto_rosetta_types_amount_proto protoreflect.FileDescriptor

var file_coinbase_crypto_rosetta_types_amount_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f,
	0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65,
	0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x4f, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x51, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xe4, 0x01, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x73, 0x12, 0x51, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x51, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_crypto_rosetta_types_amount_proto_rawDescOnce sync.Once
	file_coinbase_crypto_rosetta_types_amount_proto_rawDescData = file_coinbase_crypto_rosetta_types_amount_proto_rawDesc
)

func file_coinbase_crypto_rosetta_types_amount_proto_rawDescGZIP() []byte {
	file_coinbase_crypto_rosetta_types_amount_proto_rawDescOnce.Do(func() {
		file_coinbase_crypto_rosetta_types_amount_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_crypto_rosetta_types_amount_proto_rawDescData)
	})
	return file_coinbase_crypto_rosetta_types_amount_proto_rawDescData
}

var file_coinbase_crypto_rosetta_types_amount_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_coinbase_crypto_rosetta_types_amount_proto_goTypes = []interface{}{
	(*Amount)(nil),    // 0: coinbase.crypto.rosetta.types.Amount
	(*Currency)(nil),  // 1: coinbase.crypto.rosetta.types.Currency
	nil,               // 2: coinbase.crypto.rosetta.types.Amount.MetadataEntry
	nil,               // 3: coinbase.crypto.rosetta.types.Currency.MetadataEntry
	(*anypb.Any)(nil), // 4: google.protobuf.Any
}
var file_coinbase_crypto_rosetta_types_amount_proto_depIdxs = []int32{
	1, // 0: coinbase.crypto.rosetta.types.Amount.currency:type_name -> coinbase.crypto.rosetta.types.Currency
	2, // 1: coinbase.crypto.rosetta.types.Amount.metadata:type_name -> coinbase.crypto.rosetta.types.Amount.MetadataEntry
	3, // 2: coinbase.crypto.rosetta.types.Currency.metadata:type_name -> coinbase.crypto.rosetta.types.Currency.MetadataEntry
	4, // 3: coinbase.crypto.rosetta.types.Amount.MetadataEntry.value:type_name -> google.protobuf.Any
	4, // 4: coinbase.crypto.rosetta.types.Currency.MetadataEntry.value:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_coinbase_crypto_rosetta_types_amount_proto_init() }
func file_coinbase_crypto_rosetta_types_amount_proto_init() {
	if File_coinbase_crypto_rosetta_types_amount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_crypto_rosetta_types_amount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
		file_coinbase_crypto_rosetta_types_amount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Currency); i {
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
			RawDescriptor: file_coinbase_crypto_rosetta_types_amount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_crypto_rosetta_types_amount_proto_goTypes,
		DependencyIndexes: file_coinbase_crypto_rosetta_types_amount_proto_depIdxs,
		MessageInfos:      file_coinbase_crypto_rosetta_types_amount_proto_msgTypes,
	}.Build()
	File_coinbase_crypto_rosetta_types_amount_proto = out.File
	file_coinbase_crypto_rosetta_types_amount_proto_rawDesc = nil
	file_coinbase_crypto_rosetta_types_amount_proto_goTypes = nil
	file_coinbase_crypto_rosetta_types_amount_proto_depIdxs = nil
}
