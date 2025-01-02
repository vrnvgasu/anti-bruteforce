// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: api/AntiBruteforceService.proto

package pb

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

type ListType int32

const (
	ListType_BLACK ListType = 0
	ListType_WHITE ListType = 1
)

// Enum value maps for ListType.
var (
	ListType_name = map[int32]string{
		0: "BLACK",
		1: "WHITE",
	}
	ListType_value = map[string]int32{
		"BLACK": 0,
		"WHITE": 1,
	}
)

func (x ListType) Enum() *ListType {
	p := new(ListType)
	*p = x
	return p
}

func (x ListType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_AntiBruteforceService_proto_enumTypes[0].Descriptor()
}

func (ListType) Type() protoreflect.EnumType {
	return &file_api_AntiBruteforceService_proto_enumTypes[0]
}

func (x ListType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListType.Descriptor instead.
func (ListType) EnumDescriptor() ([]byte, []int) {
	return file_api_AntiBruteforceService_proto_rawDescGZIP(), []int{0}
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	mi := &file_api_AntiBruteforceService_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_AntiBruteforceService_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_api_AntiBruteforceService_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CheckRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CheckRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	mi := &file_api_AntiBruteforceService_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_AntiBruteforceService_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_api_AntiBruteforceService_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type ClearRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Ip    string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *ClearRequest) Reset() {
	*x = ClearRequest{}
	mi := &file_api_AntiBruteforceService_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearRequest) ProtoMessage() {}

func (x *ClearRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_AntiBruteforceService_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearRequest.ProtoReflect.Descriptor instead.
func (*ClearRequest) Descriptor() ([]byte, []int) {
	return file_api_AntiBruteforceService_proto_rawDescGZIP(), []int{2}
}

func (x *ClearRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *ClearRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   *IP      `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Type ListType `protobuf:"varint,2,opt,name=type,proto3,enum=event.ListType" json:"type,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_api_AntiBruteforceService_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_AntiBruteforceService_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_api_AntiBruteforceService_proto_rawDescGZIP(), []int{3}
}

func (x *ListRequest) GetIp() *IP {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *ListRequest) GetType() ListType {
	if x != nil {
		return x.Type
	}
	return ListType_BLACK
}

type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Mask string `protobuf:"bytes,2,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *IP) Reset() {
	*x = IP{}
	mi := &file_api_AntiBruteforceService_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_api_AntiBruteforceService_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_api_AntiBruteforceService_proto_rawDescGZIP(), []int{4}
}

func (x *IP) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *IP) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

type OK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OK) Reset() {
	*x = OK{}
	mi := &file_api_AntiBruteforceService_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OK) ProtoMessage() {}

func (x *OK) ProtoReflect() protoreflect.Message {
	mi := &file_api_AntiBruteforceService_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OK.ProtoReflect.Descriptor instead.
func (*OK) Descriptor() ([]byte, []int) {
	return file_api_AntiBruteforceService_proto_rawDescGZIP(), []int{5}
}

var File_api_AntiBruteforceService_proto protoreflect.FileDescriptor

var file_api_AntiBruteforceService_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x41, 0x6e, 0x74, 0x69, 0x42, 0x72, 0x75, 0x74, 0x65, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x34, 0x0a, 0x0c, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x4d, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x50, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x28, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61,
	0x73, 0x6b, 0x22, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x2a, 0x20, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x57, 0x48, 0x49, 0x54, 0x45, 0x10, 0x01, 0x32, 0xc9, 0x01, 0x0a, 0x0e, 0x41,
	0x6e, 0x74, 0x69, 0x42, 0x72, 0x75, 0x74, 0x65, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a,
	0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x05, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x13, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x09, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x4b, 0x22, 0x00, 0x12, 0x2c,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x09, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x4b, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x09,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x50, 0x1a, 0x09, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x4f, 0x4b, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_AntiBruteforceService_proto_rawDescOnce sync.Once
	file_api_AntiBruteforceService_proto_rawDescData = file_api_AntiBruteforceService_proto_rawDesc
)

func file_api_AntiBruteforceService_proto_rawDescGZIP() []byte {
	file_api_AntiBruteforceService_proto_rawDescOnce.Do(func() {
		file_api_AntiBruteforceService_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_AntiBruteforceService_proto_rawDescData)
	})
	return file_api_AntiBruteforceService_proto_rawDescData
}

var file_api_AntiBruteforceService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_AntiBruteforceService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_AntiBruteforceService_proto_goTypes = []any{
	(ListType)(0),         // 0: event.ListType
	(*CheckRequest)(nil),  // 1: event.CheckRequest
	(*CheckResponse)(nil), // 2: event.CheckResponse
	(*ClearRequest)(nil),  // 3: event.ClearRequest
	(*ListRequest)(nil),   // 4: event.ListRequest
	(*IP)(nil),            // 5: event.IP
	(*OK)(nil),            // 6: event.OK
}
var file_api_AntiBruteforceService_proto_depIdxs = []int32{
	5, // 0: event.ListRequest.ip:type_name -> event.IP
	0, // 1: event.ListRequest.type:type_name -> event.ListType
	1, // 2: event.AntiBruteforce.Check:input_type -> event.CheckRequest
	3, // 3: event.AntiBruteforce.Clear:input_type -> event.ClearRequest
	4, // 4: event.AntiBruteforce.AddToList:input_type -> event.ListRequest
	5, // 5: event.AntiBruteforce.RemoveFromList:input_type -> event.IP
	2, // 6: event.AntiBruteforce.Check:output_type -> event.CheckResponse
	6, // 7: event.AntiBruteforce.Clear:output_type -> event.OK
	6, // 8: event.AntiBruteforce.AddToList:output_type -> event.OK
	6, // 9: event.AntiBruteforce.RemoveFromList:output_type -> event.OK
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_AntiBruteforceService_proto_init() }
func file_api_AntiBruteforceService_proto_init() {
	if File_api_AntiBruteforceService_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_AntiBruteforceService_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_AntiBruteforceService_proto_goTypes,
		DependencyIndexes: file_api_AntiBruteforceService_proto_depIdxs,
		EnumInfos:         file_api_AntiBruteforceService_proto_enumTypes,
		MessageInfos:      file_api_AntiBruteforceService_proto_msgTypes,
	}.Build()
	File_api_AntiBruteforceService_proto = out.File
	file_api_AntiBruteforceService_proto_rawDesc = nil
	file_api_AntiBruteforceService_proto_goTypes = nil
	file_api_AntiBruteforceService_proto_depIdxs = nil
}