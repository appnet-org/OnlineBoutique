// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.6.1
// source: currency/currency.proto

package currency

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

// Represents an amount of money with its currency type.
type Money struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The whole units of the amount.
	// For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos         int32 `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Money) Reset() {
	*x = Money{}
	mi := &file_currency_currency_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_currency_currency_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_currency_currency_proto_rawDescGZIP(), []int{0}
}

func (x *Money) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Money) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Money) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type GetSupportedCurrenciesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCodes []string `protobuf:"bytes,1,rep,name=currency_codes,json=currencyCodes,proto3" json:"currency_codes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSupportedCurrenciesResponse) Reset() {
	*x = GetSupportedCurrenciesResponse{}
	mi := &file_currency_currency_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSupportedCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedCurrenciesResponse) ProtoMessage() {}

func (x *GetSupportedCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_currency_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*GetSupportedCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_currency_currency_proto_rawDescGZIP(), []int{1}
}

func (x *GetSupportedCurrenciesResponse) GetCurrencyCodes() []string {
	if x != nil {
		return x.CurrencyCodes
	}
	return nil
}

type CurrencyConversionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	From  *Money                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// The 3-letter currency code defined in ISO 4217.
	ToCode        string `protobuf:"bytes,2,opt,name=to_code,json=toCode,proto3" json:"to_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CurrencyConversionRequest) Reset() {
	*x = CurrencyConversionRequest{}
	mi := &file_currency_currency_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrencyConversionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyConversionRequest) ProtoMessage() {}

func (x *CurrencyConversionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currency_currency_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyConversionRequest.ProtoReflect.Descriptor instead.
func (*CurrencyConversionRequest) Descriptor() ([]byte, []int) {
	return file_currency_currency_proto_rawDescGZIP(), []int{2}
}

func (x *CurrencyConversionRequest) GetFrom() *Money {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *CurrencyConversionRequest) GetToCode() string {
	if x != nil {
		return x.ToCode
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_currency_currency_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_currency_currency_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_currency_currency_proto_rawDescGZIP(), []int{3}
}

var File_currency_currency_proto protoreflect.FileDescriptor

var file_currency_currency_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x22, 0x58, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22, 0x47, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x59, 0x0a, 0x19, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xab, 0x01, 0x0a, 0x0f, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x12, 0x23, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_currency_currency_proto_rawDescOnce sync.Once
	file_currency_currency_proto_rawDescData = file_currency_currency_proto_rawDesc
)

func file_currency_currency_proto_rawDescGZIP() []byte {
	file_currency_currency_proto_rawDescOnce.Do(func() {
		file_currency_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_currency_currency_proto_rawDescData)
	})
	return file_currency_currency_proto_rawDescData
}

var file_currency_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_currency_currency_proto_goTypes = []any{
	(*Money)(nil),                          // 0: currency.Money
	(*GetSupportedCurrenciesResponse)(nil), // 1: currency.GetSupportedCurrenciesResponse
	(*CurrencyConversionRequest)(nil),      // 2: currency.CurrencyConversionRequest
	(*Empty)(nil),                          // 3: currency.Empty
}
var file_currency_currency_proto_depIdxs = []int32{
	0, // 0: currency.CurrencyConversionRequest.from:type_name -> currency.Money
	3, // 1: currency.CurrencyService.GetSupportedCurrencies:input_type -> currency.Empty
	2, // 2: currency.CurrencyService.Convert:input_type -> currency.CurrencyConversionRequest
	1, // 3: currency.CurrencyService.GetSupportedCurrencies:output_type -> currency.GetSupportedCurrenciesResponse
	0, // 4: currency.CurrencyService.Convert:output_type -> currency.Money
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_currency_currency_proto_init() }
func file_currency_currency_proto_init() {
	if File_currency_currency_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_currency_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currency_currency_proto_goTypes,
		DependencyIndexes: file_currency_currency_proto_depIdxs,
		MessageInfos:      file_currency_currency_proto_msgTypes,
	}.Build()
	File_currency_currency_proto = out.File
	file_currency_currency_proto_rawDesc = nil
	file_currency_currency_proto_goTypes = nil
	file_currency_currency_proto_depIdxs = nil
}