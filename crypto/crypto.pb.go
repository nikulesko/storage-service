// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: crypto.proto

package crypto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// currency code for observing for example BTC
	BaseCurrencyCode string `protobuf:"bytes,1,opt,name=base_currency_code,json=baseCurrencyCode,proto3" json:"base_currency_code,omitempty"`
	// exchange currency code for example USD
	ExchangeCurrencyCode string `protobuf:"bytes,2,opt,name=exchange_currency_code,json=exchangeCurrencyCode,proto3" json:"exchange_currency_code,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_crypto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetBaseCurrencyCode() string {
	if x != nil {
		return x.BaseCurrencyCode
	}
	return ""
}

func (x *Request) GetExchangeCurrencyCode() string {
	if x != nil {
		return x.ExchangeCurrencyCode
	}
	return ""
}

type Response struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// unix time in ms
	Datetime int64 `protobuf:"varint,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	// exchange rate, we use the float type, but it must be something else to represent money
	Rate          float32 `protobuf:"fixed32,2,opt,name=rate,proto3" json:"rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_crypto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetDatetime() int64 {
	if x != nil {
		return x.Datetime
	}
	return 0
}

func (x *Response) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_crypto_proto protoreflect.FileDescriptor

var file_crypto_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x32, 0x39, 0x0a, 0x0a, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x08, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6b, 0x75, 0x6c, 0x65, 0x73, 0x6b, 0x6f, 0x2f, 0x67, 0x6f, 0x6c,
	0x65, 0x61, 0x72, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_crypto_proto_rawDescOnce sync.Once
	file_crypto_proto_rawDescData []byte
)

func file_crypto_proto_rawDescGZIP() []byte {
	file_crypto_proto_rawDescOnce.Do(func() {
		file_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_crypto_proto_rawDesc), len(file_crypto_proto_rawDesc)))
	})
	return file_crypto_proto_rawDescData
}

var file_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crypto_proto_goTypes = []any{
	(*Request)(nil),  // 0: Request
	(*Response)(nil), // 1: Response
}
var file_crypto_proto_depIdxs = []int32{
	0, // 0: CryptoRate.GetDataStreaming:input_type -> Request
	1, // 1: CryptoRate.GetDataStreaming:output_type -> Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crypto_proto_init() }
func file_crypto_proto_init() {
	if File_crypto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_crypto_proto_rawDesc), len(file_crypto_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crypto_proto_goTypes,
		DependencyIndexes: file_crypto_proto_depIdxs,
		MessageInfos:      file_crypto_proto_msgTypes,
	}.Build()
	File_crypto_proto = out.File
	file_crypto_proto_goTypes = nil
	file_crypto_proto_depIdxs = nil
}
