// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: dcs/controller/v0/controller.proto

package controller

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

type SetAlarmStateRequest_AlarmState int32

const (
	SetAlarmStateRequest_ALARM_STATE_UNSPECIFIED SetAlarmStateRequest_AlarmState = 0
	SetAlarmStateRequest_ALARM_STATE_AUTO        SetAlarmStateRequest_AlarmState = 1
	SetAlarmStateRequest_ALARM_STATE_GREEN       SetAlarmStateRequest_AlarmState = 2
	SetAlarmStateRequest_ALARM_STATE_RED         SetAlarmStateRequest_AlarmState = 3
)

// Enum value maps for SetAlarmStateRequest_AlarmState.
var (
	SetAlarmStateRequest_AlarmState_name = map[int32]string{
		0: "ALARM_STATE_UNSPECIFIED",
		1: "ALARM_STATE_AUTO",
		2: "ALARM_STATE_GREEN",
		3: "ALARM_STATE_RED",
	}
	SetAlarmStateRequest_AlarmState_value = map[string]int32{
		"ALARM_STATE_UNSPECIFIED": 0,
		"ALARM_STATE_AUTO":        1,
		"ALARM_STATE_GREEN":       2,
		"ALARM_STATE_RED":         3,
	}
)

func (x SetAlarmStateRequest_AlarmState) Enum() *SetAlarmStateRequest_AlarmState {
	p := new(SetAlarmStateRequest_AlarmState)
	*p = x
	return p
}

func (x SetAlarmStateRequest_AlarmState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetAlarmStateRequest_AlarmState) Descriptor() protoreflect.EnumDescriptor {
	return file_dcs_controller_v0_controller_proto_enumTypes[0].Descriptor()
}

func (SetAlarmStateRequest_AlarmState) Type() protoreflect.EnumType {
	return &file_dcs_controller_v0_controller_proto_enumTypes[0]
}

func (x SetAlarmStateRequest_AlarmState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetAlarmStateRequest_AlarmState.Descriptor instead.
func (SetAlarmStateRequest_AlarmState) EnumDescriptor() ([]byte, []int) {
	return file_dcs_controller_v0_controller_proto_rawDescGZIP(), []int{0, 0}
}

type SetAlarmStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Name:
	//	*SetAlarmStateRequest_GroupName
	//	*SetAlarmStateRequest_UnitName
	Name       isSetAlarmStateRequest_Name     `protobuf_oneof:"name"`
	AlarmState SetAlarmStateRequest_AlarmState `protobuf:"varint,3,opt,name=alarm_state,json=alarmState,proto3,enum=dcs.controller.v0.SetAlarmStateRequest_AlarmState" json:"alarm_state,omitempty"`
}

func (x *SetAlarmStateRequest) Reset() {
	*x = SetAlarmStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_controller_v0_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAlarmStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAlarmStateRequest) ProtoMessage() {}

func (x *SetAlarmStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_controller_v0_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAlarmStateRequest.ProtoReflect.Descriptor instead.
func (*SetAlarmStateRequest) Descriptor() ([]byte, []int) {
	return file_dcs_controller_v0_controller_proto_rawDescGZIP(), []int{0}
}

func (m *SetAlarmStateRequest) GetName() isSetAlarmStateRequest_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (x *SetAlarmStateRequest) GetGroupName() string {
	if x, ok := x.GetName().(*SetAlarmStateRequest_GroupName); ok {
		return x.GroupName
	}
	return ""
}

func (x *SetAlarmStateRequest) GetUnitName() string {
	if x, ok := x.GetName().(*SetAlarmStateRequest_UnitName); ok {
		return x.UnitName
	}
	return ""
}

func (x *SetAlarmStateRequest) GetAlarmState() SetAlarmStateRequest_AlarmState {
	if x != nil {
		return x.AlarmState
	}
	return SetAlarmStateRequest_ALARM_STATE_UNSPECIFIED
}

type isSetAlarmStateRequest_Name interface {
	isSetAlarmStateRequest_Name()
}

type SetAlarmStateRequest_GroupName struct {
	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3,oneof"`
}

type SetAlarmStateRequest_UnitName struct {
	UnitName string `protobuf:"bytes,2,opt,name=unit_name,json=unitName,proto3,oneof"`
}

func (*SetAlarmStateRequest_GroupName) isSetAlarmStateRequest_Name() {}

func (*SetAlarmStateRequest_UnitName) isSetAlarmStateRequest_Name() {}

type SetAlarmStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetAlarmStateResponse) Reset() {
	*x = SetAlarmStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_controller_v0_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAlarmStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAlarmStateResponse) ProtoMessage() {}

func (x *SetAlarmStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_controller_v0_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAlarmStateResponse.ProtoReflect.Descriptor instead.
func (*SetAlarmStateResponse) Descriptor() ([]byte, []int) {
	return file_dcs_controller_v0_controller_proto_rawDescGZIP(), []int{1}
}

var File_dcs_controller_v0_controller_proto protoreflect.FileDescriptor

var file_dcs_controller_v0_controller_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x76, 0x30, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x30, 0x22, 0xa0, 0x02, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x41,
	0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x53, 0x0a, 0x0b, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x6c, 0x61,
	0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41,
	0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x6b, 0x0a, 0x0a, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x4c, 0x41, 0x52, 0x4d, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x41, 0x4c, 0x41, 0x52, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x41, 0x55, 0x54, 0x4f, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x4c, 0x41, 0x52, 0x4d, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x13, 0x0a,
	0x0f, 0x41, 0x4c, 0x41, 0x52, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x44,
	0x10, 0x03, 0x42, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65,
	0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x79, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x41,
	0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x64, 0x63, 0x73, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65,
	0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x59,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x43, 0x53,
	0x2d, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x64, 0x63, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0xaa, 0x02, 0x23, 0x52, 0x75, 0x72, 0x6f, 0x75, 0x6e, 0x69, 0x4a, 0x6f, 0x6e,
	0x65, 0x73, 0x2e, 0x44, 0x63, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x30, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dcs_controller_v0_controller_proto_rawDescOnce sync.Once
	file_dcs_controller_v0_controller_proto_rawDescData = file_dcs_controller_v0_controller_proto_rawDesc
)

func file_dcs_controller_v0_controller_proto_rawDescGZIP() []byte {
	file_dcs_controller_v0_controller_proto_rawDescOnce.Do(func() {
		file_dcs_controller_v0_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_dcs_controller_v0_controller_proto_rawDescData)
	})
	return file_dcs_controller_v0_controller_proto_rawDescData
}

var file_dcs_controller_v0_controller_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dcs_controller_v0_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dcs_controller_v0_controller_proto_goTypes = []interface{}{
	(SetAlarmStateRequest_AlarmState)(0), // 0: dcs.controller.v0.SetAlarmStateRequest.AlarmState
	(*SetAlarmStateRequest)(nil),         // 1: dcs.controller.v0.SetAlarmStateRequest
	(*SetAlarmStateResponse)(nil),        // 2: dcs.controller.v0.SetAlarmStateResponse
}
var file_dcs_controller_v0_controller_proto_depIdxs = []int32{
	0, // 0: dcs.controller.v0.SetAlarmStateRequest.alarm_state:type_name -> dcs.controller.v0.SetAlarmStateRequest.AlarmState
	1, // 1: dcs.controller.v0.ControllerService.SetAlarmState:input_type -> dcs.controller.v0.SetAlarmStateRequest
	2, // 2: dcs.controller.v0.ControllerService.SetAlarmState:output_type -> dcs.controller.v0.SetAlarmStateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dcs_controller_v0_controller_proto_init() }
func file_dcs_controller_v0_controller_proto_init() {
	if File_dcs_controller_v0_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dcs_controller_v0_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAlarmStateRequest); i {
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
		file_dcs_controller_v0_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAlarmStateResponse); i {
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
	file_dcs_controller_v0_controller_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SetAlarmStateRequest_GroupName)(nil),
		(*SetAlarmStateRequest_UnitName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dcs_controller_v0_controller_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dcs_controller_v0_controller_proto_goTypes,
		DependencyIndexes: file_dcs_controller_v0_controller_proto_depIdxs,
		EnumInfos:         file_dcs_controller_v0_controller_proto_enumTypes,
		MessageInfos:      file_dcs_controller_v0_controller_proto_msgTypes,
	}.Build()
	File_dcs_controller_v0_controller_proto = out.File
	file_dcs_controller_v0_controller_proto_rawDesc = nil
	file_dcs_controller_v0_controller_proto_goTypes = nil
	file_dcs_controller_v0_controller_proto_depIdxs = nil
}
