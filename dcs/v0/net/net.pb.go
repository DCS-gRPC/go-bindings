// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: dcs/net/v0/net.proto

package net

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

type SendChatToRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the message to send in the chat
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// the target player of the direct message
	TargetPlayerId uint32 `protobuf:"varint,2,opt,name=target_player_id,json=targetPlayerId,proto3" json:"target_player_id,omitempty"`
}

func (x *SendChatToRequest) Reset() {
	*x = SendChatToRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChatToRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChatToRequest) ProtoMessage() {}

func (x *SendChatToRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChatToRequest.ProtoReflect.Descriptor instead.
func (*SendChatToRequest) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{0}
}

func (x *SendChatToRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendChatToRequest) GetTargetPlayerId() uint32 {
	if x != nil {
		return x.TargetPlayerId
	}
	return 0
}

type SendChatToResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendChatToResponse) Reset() {
	*x = SendChatToResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChatToResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChatToResponse) ProtoMessage() {}

func (x *SendChatToResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChatToResponse.ProtoReflect.Descriptor instead.
func (*SendChatToResponse) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{1}
}

type SendChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the message to send in the chat
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// which coalition? DCS only supports ALL or NEUTRAL
	// (only applicable to send_chat)
	Coalition common.Coalition `protobuf:"varint,2,opt,name=coalition,proto3,enum=dcs.common.v0.Coalition" json:"coalition,omitempty"`
}

func (x *SendChatRequest) Reset() {
	*x = SendChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChatRequest) ProtoMessage() {}

func (x *SendChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChatRequest.ProtoReflect.Descriptor instead.
func (*SendChatRequest) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{2}
}

func (x *SendChatRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendChatRequest) GetCoalition() common.Coalition {
	if x != nil {
		return x.Coalition
	}
	return common.Coalition(0)
}

type SendChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendChatResponse) Reset() {
	*x = SendChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChatResponse) ProtoMessage() {}

func (x *SendChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChatResponse.ProtoReflect.Descriptor instead.
func (*SendChatResponse) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{3}
}

type GetPlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlayersRequest) Reset() {
	*x = GetPlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayersRequest) ProtoMessage() {}

func (x *GetPlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayersRequest.ProtoReflect.Descriptor instead.
func (*GetPlayersRequest) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{4}
}

type GetPlayersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of all the players connected to the server
	Players []*GetPlayersResponse_GetPlayerInfo `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *GetPlayersResponse) Reset() {
	*x = GetPlayersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayersResponse) ProtoMessage() {}

func (x *GetPlayersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayersResponse.ProtoReflect.Descriptor instead.
func (*GetPlayersResponse) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlayersResponse) GetPlayers() []*GetPlayersResponse_GetPlayerInfo {
	if x != nil {
		return x.Players
	}
	return nil
}

type ForcePlayerSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId  uint32           `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Coalition common.Coalition `protobuf:"varint,2,opt,name=coalition,proto3,enum=dcs.common.v0.Coalition" json:"coalition,omitempty"`
	SlotId    string           `protobuf:"bytes,3,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
}

func (x *ForcePlayerSlotRequest) Reset() {
	*x = ForcePlayerSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForcePlayerSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForcePlayerSlotRequest) ProtoMessage() {}

func (x *ForcePlayerSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForcePlayerSlotRequest.ProtoReflect.Descriptor instead.
func (*ForcePlayerSlotRequest) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{6}
}

func (x *ForcePlayerSlotRequest) GetPlayerId() uint32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *ForcePlayerSlotRequest) GetCoalition() common.Coalition {
	if x != nil {
		return x.Coalition
	}
	return common.Coalition(0)
}

func (x *ForcePlayerSlotRequest) GetSlotId() string {
	if x != nil {
		return x.SlotId
	}
	return ""
}

type ForcePlayerSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForcePlayerSlotResponse) Reset() {
	*x = ForcePlayerSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForcePlayerSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForcePlayerSlotResponse) ProtoMessage() {}

func (x *ForcePlayerSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForcePlayerSlotResponse.ProtoReflect.Descriptor instead.
func (*ForcePlayerSlotResponse) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{7}
}

type KickPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *KickPlayerRequest) Reset() {
	*x = KickPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickPlayerRequest) ProtoMessage() {}

func (x *KickPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickPlayerRequest.ProtoReflect.Descriptor instead.
func (*KickPlayerRequest) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{8}
}

func (x *KickPlayerRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KickPlayerRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type KickPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KickPlayerResponse) Reset() {
	*x = KickPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickPlayerResponse) ProtoMessage() {}

func (x *KickPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickPlayerResponse.ProtoReflect.Descriptor instead.
func (*KickPlayerResponse) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{9}
}

type GetPlayersResponse_GetPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the player id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// player's online name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// coalition which player is slotted in
	Coalition common.Coalition `protobuf:"varint,3,opt,name=coalition,proto3,enum=dcs.common.v0.Coalition" json:"coalition,omitempty"`
	// the slot identifier
	Slot string `protobuf:"bytes,4,opt,name=slot,proto3" json:"slot,omitempty"`
	// the ping of the player
	Ping uint32 `protobuf:"varint,5,opt,name=ping,proto3" json:"ping,omitempty"`
	// the connection ip address and port the client
	// has established with the server
	RemoteAddress string `protobuf:"bytes,6,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	// the unique identifier for the player
	Ucid string `protobuf:"bytes,7,opt,name=ucid,proto3" json:"ucid,omitempty"`
	// abbreviated language (locale) e.g. "en"
	Locale string `protobuf:"bytes,8,opt,name=locale,proto3" json:"locale,omitempty"`
}

func (x *GetPlayersResponse_GetPlayerInfo) Reset() {
	*x = GetPlayersResponse_GetPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dcs_net_v0_net_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayersResponse_GetPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayersResponse_GetPlayerInfo) ProtoMessage() {}

func (x *GetPlayersResponse_GetPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dcs_net_v0_net_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayersResponse_GetPlayerInfo.ProtoReflect.Descriptor instead.
func (*GetPlayersResponse_GetPlayerInfo) Descriptor() ([]byte, []int) {
	return file_dcs_net_v0_net_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetPlayersResponse_GetPlayerInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetPlayersResponse_GetPlayerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPlayersResponse_GetPlayerInfo) GetCoalition() common.Coalition {
	if x != nil {
		return x.Coalition
	}
	return common.Coalition(0)
}

func (x *GetPlayersResponse_GetPlayerInfo) GetSlot() string {
	if x != nil {
		return x.Slot
	}
	return ""
}

func (x *GetPlayersResponse_GetPlayerInfo) GetPing() uint32 {
	if x != nil {
		return x.Ping
	}
	return 0
}

func (x *GetPlayersResponse_GetPlayerInfo) GetRemoteAddress() string {
	if x != nil {
		return x.RemoteAddress
	}
	return ""
}

func (x *GetPlayersResponse_GetPlayerInfo) GetUcid() string {
	if x != nil {
		return x.Ucid
	}
	return ""
}

func (x *GetPlayersResponse_GetPlayerInfo) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

var File_dcs_net_v0_net_proto protoreflect.FileDescriptor

var file_dcs_net_v0_net_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x63, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x76, 0x30, 0x2f, 0x6e, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x63, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x2e,
	0x76, 0x30, 0x1a, 0x1a, 0x64, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x30, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57,
	0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x43,
	0x68, 0x61, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x0a,
	0x0f, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x63, 0x6f,
	0x61, 0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x6f,
	0x61, 0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x61, 0x6c, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc5, 0x02, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x76, 0x30,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x1a, 0xe6, 0x01, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x36, 0x0a, 0x09, 0x63, 0x6f, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x30, 0x2e, 0x43, 0x6f, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63,
	0x6f, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x63, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x63, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x63,
	0x6f, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x64, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x43,
	0x6f, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x61, 0x6c, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17,
	0x46, 0x6f, 0x72, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x11, 0x4b, 0x69, 0x63, 0x6b, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x4b, 0x69, 0x63, 0x6b, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa0, 0x03, 0x0a,
	0x0a, 0x4e, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x54, 0x6f, 0x12, 0x1d, 0x2e, 0x64, 0x63, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x54,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x54, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x08, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1b, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x76, 0x30,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x1d, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x76, 0x30, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x4b, 0x69, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x76, 0x30, 0x2e, 0x4b, 0x69,
	0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x76, 0x30, 0x2e, 0x4b, 0x69, 0x63,
	0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5c, 0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x53, 0x6c, 0x6f, 0x74, 0x12, 0x22, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x76,
	0x30, 0x2e, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x63, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x2e, 0x76, 0x30, 0x2e, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x4b, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x43,
	0x53, 0x2d, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x64, 0x63, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x6e, 0x65, 0x74, 0xaa, 0x02, 0x1c,
	0x52, 0x75, 0x72, 0x6f, 0x75, 0x6e, 0x69, 0x4a, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x44, 0x63, 0x73,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x30, 0x2e, 0x4e, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dcs_net_v0_net_proto_rawDescOnce sync.Once
	file_dcs_net_v0_net_proto_rawDescData = file_dcs_net_v0_net_proto_rawDesc
)

func file_dcs_net_v0_net_proto_rawDescGZIP() []byte {
	file_dcs_net_v0_net_proto_rawDescOnce.Do(func() {
		file_dcs_net_v0_net_proto_rawDescData = protoimpl.X.CompressGZIP(file_dcs_net_v0_net_proto_rawDescData)
	})
	return file_dcs_net_v0_net_proto_rawDescData
}

var file_dcs_net_v0_net_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_dcs_net_v0_net_proto_goTypes = []interface{}{
	(*SendChatToRequest)(nil),                // 0: dcs.net.v0.SendChatToRequest
	(*SendChatToResponse)(nil),               // 1: dcs.net.v0.SendChatToResponse
	(*SendChatRequest)(nil),                  // 2: dcs.net.v0.SendChatRequest
	(*SendChatResponse)(nil),                 // 3: dcs.net.v0.SendChatResponse
	(*GetPlayersRequest)(nil),                // 4: dcs.net.v0.GetPlayersRequest
	(*GetPlayersResponse)(nil),               // 5: dcs.net.v0.GetPlayersResponse
	(*ForcePlayerSlotRequest)(nil),           // 6: dcs.net.v0.ForcePlayerSlotRequest
	(*ForcePlayerSlotResponse)(nil),          // 7: dcs.net.v0.ForcePlayerSlotResponse
	(*KickPlayerRequest)(nil),                // 8: dcs.net.v0.KickPlayerRequest
	(*KickPlayerResponse)(nil),               // 9: dcs.net.v0.KickPlayerResponse
	(*GetPlayersResponse_GetPlayerInfo)(nil), // 10: dcs.net.v0.GetPlayersResponse.GetPlayerInfo
	(common.Coalition)(0),                    // 11: dcs.common.v0.Coalition
}
var file_dcs_net_v0_net_proto_depIdxs = []int32{
	11, // 0: dcs.net.v0.SendChatRequest.coalition:type_name -> dcs.common.v0.Coalition
	10, // 1: dcs.net.v0.GetPlayersResponse.players:type_name -> dcs.net.v0.GetPlayersResponse.GetPlayerInfo
	11, // 2: dcs.net.v0.ForcePlayerSlotRequest.coalition:type_name -> dcs.common.v0.Coalition
	11, // 3: dcs.net.v0.GetPlayersResponse.GetPlayerInfo.coalition:type_name -> dcs.common.v0.Coalition
	0,  // 4: dcs.net.v0.NetService.SendChatTo:input_type -> dcs.net.v0.SendChatToRequest
	2,  // 5: dcs.net.v0.NetService.SendChat:input_type -> dcs.net.v0.SendChatRequest
	4,  // 6: dcs.net.v0.NetService.GetPlayers:input_type -> dcs.net.v0.GetPlayersRequest
	8,  // 7: dcs.net.v0.NetService.KickPlayer:input_type -> dcs.net.v0.KickPlayerRequest
	6,  // 8: dcs.net.v0.NetService.ForcePlayerSlot:input_type -> dcs.net.v0.ForcePlayerSlotRequest
	1,  // 9: dcs.net.v0.NetService.SendChatTo:output_type -> dcs.net.v0.SendChatToResponse
	3,  // 10: dcs.net.v0.NetService.SendChat:output_type -> dcs.net.v0.SendChatResponse
	5,  // 11: dcs.net.v0.NetService.GetPlayers:output_type -> dcs.net.v0.GetPlayersResponse
	9,  // 12: dcs.net.v0.NetService.KickPlayer:output_type -> dcs.net.v0.KickPlayerResponse
	7,  // 13: dcs.net.v0.NetService.ForcePlayerSlot:output_type -> dcs.net.v0.ForcePlayerSlotResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_dcs_net_v0_net_proto_init() }
func file_dcs_net_v0_net_proto_init() {
	if File_dcs_net_v0_net_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dcs_net_v0_net_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendChatToRequest); i {
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
		file_dcs_net_v0_net_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendChatToResponse); i {
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
		file_dcs_net_v0_net_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendChatRequest); i {
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
		file_dcs_net_v0_net_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendChatResponse); i {
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
		file_dcs_net_v0_net_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayersRequest); i {
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
		file_dcs_net_v0_net_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayersResponse); i {
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
		file_dcs_net_v0_net_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForcePlayerSlotRequest); i {
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
		file_dcs_net_v0_net_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForcePlayerSlotResponse); i {
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
		file_dcs_net_v0_net_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickPlayerRequest); i {
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
		file_dcs_net_v0_net_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickPlayerResponse); i {
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
		file_dcs_net_v0_net_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayersResponse_GetPlayerInfo); i {
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
			RawDescriptor: file_dcs_net_v0_net_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dcs_net_v0_net_proto_goTypes,
		DependencyIndexes: file_dcs_net_v0_net_proto_depIdxs,
		MessageInfos:      file_dcs_net_v0_net_proto_msgTypes,
	}.Build()
	File_dcs_net_v0_net_proto = out.File
	file_dcs_net_v0_net_proto_rawDesc = nil
	file_dcs_net_v0_net_proto_goTypes = nil
	file_dcs_net_v0_net_proto_depIdxs = nil
}
