// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.3
// source: dcs/group/v0/group.proto

package group

import (
	common "github.com/DCS-gRPC/go-bindings/dcs/v0/common"
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

type GetUnitsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	// Whether the response should include only active units (`true`), only
	// inactive units (`false`), or all units (`nil`).
	Active *bool `protobuf:"varint,2,opt,name=active,proto3,oneof" json:"active,omitempty"`
}

func (x *GetUnitsRequest) Reset() {
	*x = GetUnitsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_group_v0_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnitsRequest) ProtoMessage() {}

func (x *GetUnitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_group_v0_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnitsRequest.ProtoReflect.Descriptor instead.
func (*GetUnitsRequest) Descriptor() ([]byte, []int) {
	return file_dcs_group_v0_group_proto_rawDescGZIP(), []int{0}
}

func (x *GetUnitsRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GetUnitsRequest) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

type GetUnitsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units []*common.Unit `protobuf:"bytes,1,rep,name=units,proto3" json:"units,omitempty"`
}

func (x *GetUnitsResponse) Reset() {
	*x = GetUnitsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_group_v0_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnitsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnitsResponse) ProtoMessage() {}

func (x *GetUnitsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_group_v0_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnitsResponse.ProtoReflect.Descriptor instead.
func (*GetUnitsResponse) Descriptor() ([]byte, []int) {
	return file_dcs_group_v0_group_proto_rawDescGZIP(), []int{1}
}

func (x *GetUnitsResponse) GetUnits() []*common.Unit {
	if x != nil {
		return x.Units
	}
	return nil
}

var File_dcs_group_v0_group_proto protoreflect.FileDescriptor

var file_dcs_group_v0_group_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x63, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x30, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x63, 0x73, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x30, 0x1a, 0x1a, 0x64, 0x63, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x3d,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x30, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x32, 0x5b, 0x0a,
	0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x64, 0x63, 0x73, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4f, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x43, 0x53, 0x2d, 0x67, 0x52, 0x50,
	0x43, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x64, 0x63,
	0x73, 0x2f, 0x76, 0x30, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0xaa, 0x02, 0x1e, 0x52, 0x75, 0x72,
	0x6f, 0x75, 0x6e, 0x69, 0x4a, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x44, 0x63, 0x73, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x2e, 0x56, 0x30, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_dcs_group_v0_group_proto_rawDescOnce sync.Once
	file_dcs_group_v0_group_proto_rawDescData = file_dcs_group_v0_group_proto_rawDesc
)

func file_dcs_group_v0_group_proto_rawDescGZIP() []byte {
	file_dcs_group_v0_group_proto_rawDescOnce.Do(func() {
		file_dcs_group_v0_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_dcs_group_v0_group_proto_rawDescData)
	})
	return file_dcs_group_v0_group_proto_rawDescData
}

var file_dcs_group_v0_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dcs_group_v0_group_proto_goTypes = []interface{}{
	(*GetUnitsRequest)(nil),  // 0: dcs.group.v0.GetUnitsRequest
	(*GetUnitsResponse)(nil), // 1: dcs.group.v0.GetUnitsResponse
	(*common.Unit)(nil),      // 2: dcs.common.v0.Unit
}
var file_dcs_group_v0_group_proto_depIdxs = []int32{
	2, // 0: dcs.group.v0.GetUnitsResponse.units:type_name -> dcs.common.v0.Unit
	0, // 1: dcs.group.v0.GroupService.GetUnits:input_type -> dcs.group.v0.GetUnitsRequest
	1, // 2: dcs.group.v0.GroupService.GetUnits:output_type -> dcs.group.v0.GetUnitsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dcs_group_v0_group_proto_init() }
func file_dcs_group_v0_group_proto_init() {
	if File_dcs_group_v0_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dcs_group_v0_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnitsRequest); i {
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
		file_dcs_group_v0_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnitsResponse); i {
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
	file_dcs_group_v0_group_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dcs_group_v0_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dcs_group_v0_group_proto_goTypes,
		DependencyIndexes: file_dcs_group_v0_group_proto_depIdxs,
		MessageInfos:      file_dcs_group_v0_group_proto_msgTypes,
	}.Build()
	File_dcs_group_v0_group_proto = out.File
	file_dcs_group_v0_group_proto_rawDesc = nil
	file_dcs_group_v0_group_proto_goTypes = nil
	file_dcs_group_v0_group_proto_depIdxs = nil
}
