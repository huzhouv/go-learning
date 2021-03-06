// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: arith.proto

package mygrpc

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

// 客户端发送给服务端
type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"` // 每一个message的字段都有一个唯一编号，用以区分消息二进制格式，而且不能更改
	B int32 `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arith_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_arith_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_arith_proto_rawDescGZIP(), []int{0}
}

func (x *Param) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Param) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

// 服务端返回给客户端
type Int struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret int32 `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
}

func (x *Int) Reset() {
	*x = Int{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arith_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int) ProtoMessage() {}

func (x *Int) ProtoReflect() protoreflect.Message {
	mi := &file_arith_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int.ProtoReflect.Descriptor instead.
func (*Int) Descriptor() ([]byte, []int) {
	return file_arith_proto_rawDescGZIP(), []int{1}
}

func (x *Int) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

// 服务端返回给客户端
type Quotient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quo int32 `protobuf:"varint,1,opt,name=Quo,proto3" json:"Quo,omitempty"`
	Rem int32 `protobuf:"varint,2,opt,name=Rem,proto3" json:"Rem,omitempty"`
}

func (x *Quotient) Reset() {
	*x = Quotient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arith_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quotient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quotient) ProtoMessage() {}

func (x *Quotient) ProtoReflect() protoreflect.Message {
	mi := &file_arith_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quotient.ProtoReflect.Descriptor instead.
func (*Quotient) Descriptor() ([]byte, []int) {
	return file_arith_proto_rawDescGZIP(), []int{2}
}

func (x *Quotient) GetQuo() int32 {
	if x != nil {
		return x.Quo
	}
	return 0
}

func (x *Quotient) GetRem() int32 {
	if x != nil {
		return x.Rem
	}
	return 0
}

var File_arith_proto protoreflect.FileDescriptor

var file_arith_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x72, 0x69, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a,
	0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x41, 0x12, 0x0c, 0x0a, 0x01, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x42, 0x22, 0x17, 0x0a, 0x03, 0x49, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x22, 0x2e, 0x0a, 0x08, 0x51,
	0x75, 0x6f, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x51, 0x75, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x51, 0x75, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x65, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x52, 0x65, 0x6d, 0x32, 0x42, 0x0a, 0x05, 0x41,
	0x72, 0x69, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79,
	0x12, 0x06, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x04, 0x2e, 0x49, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x1d, 0x0a, 0x06, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x12, 0x06, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x1a, 0x09, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arith_proto_rawDescOnce sync.Once
	file_arith_proto_rawDescData = file_arith_proto_rawDesc
)

func file_arith_proto_rawDescGZIP() []byte {
	file_arith_proto_rawDescOnce.Do(func() {
		file_arith_proto_rawDescData = protoimpl.X.CompressGZIP(file_arith_proto_rawDescData)
	})
	return file_arith_proto_rawDescData
}

var file_arith_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_arith_proto_goTypes = []interface{}{
	(*Param)(nil),    // 0: Param
	(*Int)(nil),      // 1: Int
	(*Quotient)(nil), // 2: Quotient
}
var file_arith_proto_depIdxs = []int32{
	0, // 0: Arith.Multiply:input_type -> Param
	0, // 1: Arith.Divide:input_type -> Param
	1, // 2: Arith.Multiply:output_type -> Int
	2, // 3: Arith.Divide:output_type -> Quotient
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_arith_proto_init() }
func file_arith_proto_init() {
	if File_arith_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arith_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
		file_arith_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int); i {
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
		file_arith_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quotient); i {
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
			RawDescriptor: file_arith_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arith_proto_goTypes,
		DependencyIndexes: file_arith_proto_depIdxs,
		MessageInfos:      file_arith_proto_msgTypes,
	}.Build()
	File_arith_proto = out.File
	file_arith_proto_rawDesc = nil
	file_arith_proto_goTypes = nil
	file_arith_proto_depIdxs = nil
}
