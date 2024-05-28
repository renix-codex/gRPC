// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: org.proto

package protobufs

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

type OrgIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrgIDRequest) Reset() {
	*x = OrgIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgIDRequest) ProtoMessage() {}

func (x *OrgIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgIDRequest.ProtoReflect.Descriptor instead.
func (*OrgIDRequest) Descriptor() ([]byte, []int) {
	return file_org_proto_rawDescGZIP(), []int{0}
}

func (x *OrgIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OrgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Meta *Meta  `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *OrgResponse) Reset() {
	*x = OrgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgResponse) ProtoMessage() {}

func (x *OrgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_org_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgResponse.ProtoReflect.Descriptor instead.
func (*OrgResponse) Descriptor() ([]byte, []int) {
	return file_org_proto_rawDescGZIP(), []int{1}
}

func (x *OrgResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrgResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrgResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Pin     string `protobuf:"bytes,2,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_org_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_org_proto_rawDescGZIP(), []int{2}
}

func (x *Meta) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Meta) GetPin() string {
	if x != nil {
		return x.Pin
	}
	return ""
}

var File_org_proto protoreflect.FileDescriptor

var file_org_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x0b, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x32, 0x0a, 0x04, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x32, 0x49,
	0x0a, 0x0a, 0x4f, 0x72, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x42, 0x79, 0x49, 0x44, 0x12, 0x15, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x67, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_org_proto_rawDescOnce sync.Once
	file_org_proto_rawDescData = file_org_proto_rawDesc
)

func file_org_proto_rawDescGZIP() []byte {
	file_org_proto_rawDescOnce.Do(func() {
		file_org_proto_rawDescData = protoimpl.X.CompressGZIP(file_org_proto_rawDescData)
	})
	return file_org_proto_rawDescData
}

var file_org_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_org_proto_goTypes = []interface{}{
	(*OrgIDRequest)(nil), // 0: example.OrgIDRequest
	(*OrgResponse)(nil),  // 1: example.OrgResponse
	(*Meta)(nil),         // 2: example.Meta
}
var file_org_proto_depIdxs = []int32{
	2, // 0: example.OrgResponse.meta:type_name -> example.Meta
	0, // 1: example.OrgService.GetOrgByID:input_type -> example.OrgIDRequest
	1, // 2: example.OrgService.GetOrgByID:output_type -> example.OrgResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_org_proto_init() }
func file_org_proto_init() {
	if File_org_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_org_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgIDRequest); i {
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
		file_org_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgResponse); i {
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
		file_org_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
			RawDescriptor: file_org_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_org_proto_goTypes,
		DependencyIndexes: file_org_proto_depIdxs,
		MessageInfos:      file_org_proto_msgTypes,
	}.Build()
	File_org_proto = out.File
	file_org_proto_rawDesc = nil
	file_org_proto_goTypes = nil
	file_org_proto_depIdxs = nil
}
