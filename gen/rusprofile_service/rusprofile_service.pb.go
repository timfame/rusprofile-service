// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.3
// source: rusprofile_service/rusprofile_service.proto

package rusprofile_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCompanyByINNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inn string `protobuf:"bytes,1,opt,name=inn,proto3" json:"inn,omitempty"`
}

func (x *GetCompanyByINNRequest) Reset() {
	*x = GetCompanyByINNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rusprofile_service_rusprofile_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyByINNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyByINNRequest) ProtoMessage() {}

func (x *GetCompanyByINNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rusprofile_service_rusprofile_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyByINNRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyByINNRequest) Descriptor() ([]byte, []int) {
	return file_rusprofile_service_rusprofile_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCompanyByINNRequest) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

type GetCompanyByINNResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Inn      string `protobuf:"bytes,2,opt,name=inn,proto3" json:"inn,omitempty"`
	Kpp      string `protobuf:"bytes,3,opt,name=kpp,proto3" json:"kpp,omitempty"`
	Director string `protobuf:"bytes,4,opt,name=director,proto3" json:"director,omitempty"`
}

func (x *GetCompanyByINNResponse) Reset() {
	*x = GetCompanyByINNResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rusprofile_service_rusprofile_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyByINNResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyByINNResponse) ProtoMessage() {}

func (x *GetCompanyByINNResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rusprofile_service_rusprofile_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyByINNResponse.ProtoReflect.Descriptor instead.
func (*GetCompanyByINNResponse) Descriptor() ([]byte, []int) {
	return file_rusprofile_service_rusprofile_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCompanyByINNResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCompanyByINNResponse) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *GetCompanyByINNResponse) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *GetCompanyByINNResponse) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

var File_rusprofile_service_rusprofile_service_proto protoreflect.FileDescriptor

var file_rusprofile_service_rusprofile_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72,
	0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49,
	0x4e, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x22, 0x6d, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49, 0x4e, 0x4e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x70, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x70, 0x70, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x97, 0x01, 0x0a, 0x11, 0x52,
	0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x81, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42,
	0x79, 0x49, 0x4e, 0x4e, 0x12, 0x2a, 0x2e, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49, 0x4e, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x42, 0x79, 0x49, 0x4e, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x69, 0x65, 0x73, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rusprofile_service_rusprofile_service_proto_rawDescOnce sync.Once
	file_rusprofile_service_rusprofile_service_proto_rawDescData = file_rusprofile_service_rusprofile_service_proto_rawDesc
)

func file_rusprofile_service_rusprofile_service_proto_rawDescGZIP() []byte {
	file_rusprofile_service_rusprofile_service_proto_rawDescOnce.Do(func() {
		file_rusprofile_service_rusprofile_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rusprofile_service_rusprofile_service_proto_rawDescData)
	})
	return file_rusprofile_service_rusprofile_service_proto_rawDescData
}

var file_rusprofile_service_rusprofile_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rusprofile_service_rusprofile_service_proto_goTypes = []interface{}{
	(*GetCompanyByINNRequest)(nil),  // 0: rusprofile_service.GetCompanyByINNRequest
	(*GetCompanyByINNResponse)(nil), // 1: rusprofile_service.GetCompanyByINNResponse
}
var file_rusprofile_service_rusprofile_service_proto_depIdxs = []int32{
	0, // 0: rusprofile_service.RusprofileService.GetCompanyByINN:input_type -> rusprofile_service.GetCompanyByINNRequest
	1, // 1: rusprofile_service.RusprofileService.GetCompanyByINN:output_type -> rusprofile_service.GetCompanyByINNResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rusprofile_service_rusprofile_service_proto_init() }
func file_rusprofile_service_rusprofile_service_proto_init() {
	if File_rusprofile_service_rusprofile_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rusprofile_service_rusprofile_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyByINNRequest); i {
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
		file_rusprofile_service_rusprofile_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyByINNResponse); i {
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
			RawDescriptor: file_rusprofile_service_rusprofile_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rusprofile_service_rusprofile_service_proto_goTypes,
		DependencyIndexes: file_rusprofile_service_rusprofile_service_proto_depIdxs,
		MessageInfos:      file_rusprofile_service_rusprofile_service_proto_msgTypes,
	}.Build()
	File_rusprofile_service_rusprofile_service_proto = out.File
	file_rusprofile_service_rusprofile_service_proto_rawDesc = nil
	file_rusprofile_service_rusprofile_service_proto_goTypes = nil
	file_rusprofile_service_rusprofile_service_proto_depIdxs = nil
}
