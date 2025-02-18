// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: dcs/hook/v0/hook.proto

package hook

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HookServiceClient is the client API for HookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HookServiceClient interface {
	// https://wiki.hoggitworld.com/view/DCS_func_getMissionName
	GetMissionName(ctx context.Context, in *GetMissionNameRequest, opts ...grpc.CallOption) (*GetMissionNameResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getMissionFilename
	GetMissionFilename(ctx context.Context, in *GetMissionFilenameRequest, opts ...grpc.CallOption) (*GetMissionFilenameResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getMissionDescription
	GetMissionDescription(ctx context.Context, in *GetMissionDescriptionRequest, opts ...grpc.CallOption) (*GetMissionDescriptionResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getPause
	GetPaused(ctx context.Context, in *GetPausedRequest, opts ...grpc.CallOption) (*GetPausedResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_setPause
	SetPaused(ctx context.Context, in *SetPausedRequest, opts ...grpc.CallOption) (*SetPausedResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_stopMission
	StopMission(ctx context.Context, in *StopMissionRequest, opts ...grpc.CallOption) (*StopMissionResponse, error)
	// Reload the currently running mission
	ReloadCurrentMission(ctx context.Context, in *ReloadCurrentMissionRequest, opts ...grpc.CallOption) (*ReloadCurrentMissionResponse, error)
	// Load the next mission in the server mission list. Note that it does
	// not loop back to the first mission once the end of the mission list
	// has been reached
	LoadNextMission(ctx context.Context, in *LoadNextMissionRequest, opts ...grpc.CallOption) (*LoadNextMissionResponse, error)
	// Load a specific mission file. This does not need to be in the mission
	// list.
	LoadMission(ctx context.Context, in *LoadMissionRequest, opts ...grpc.CallOption) (*LoadMissionResponse, error)
	// Evaluate some Lua inside of the hook environment and return the result as a
	// JSON string. Disabled by default.
	Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_exitProcess
	ExitProcess(ctx context.Context, in *ExitProcessRequest, opts ...grpc.CallOption) (*ExitProcessResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_isMultiplayer
	IsMultiplayer(ctx context.Context, in *IsMultiplayerRequest, opts ...grpc.CallOption) (*IsMultiplayerResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_isServer
	IsServer(ctx context.Context, in *IsServerRequest, opts ...grpc.CallOption) (*IsServerResponse, error)
	// Bans a player that is currently connected to the server
	BanPlayer(ctx context.Context, in *BanPlayerRequest, opts ...grpc.CallOption) (*BanPlayerResponse, error)
	// Unbans a player via their globally unique ID
	UnbanPlayer(ctx context.Context, in *UnbanPlayerRequest, opts ...grpc.CallOption) (*UnbanPlayerResponse, error)
	// Get a list of all the banned players
	GetBannedPlayers(ctx context.Context, in *GetBannedPlayersRequest, opts ...grpc.CallOption) (*GetBannedPlayersResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getUnitType
	GetUnitType(ctx context.Context, in *GetUnitTypeRequest, opts ...grpc.CallOption) (*GetUnitTypeResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getRealTime
	GetRealTime(ctx context.Context, in *GetRealTimeRequest, opts ...grpc.CallOption) (*GetRealTimeResponse, error)
	// Get a count of ballistics objects
	GetBallisticsCount(ctx context.Context, in *GetBallisticsCountRequest, opts ...grpc.CallOption) (*GetBallisticsCountResponse, error)
}

type hookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHookServiceClient(cc grpc.ClientConnInterface) HookServiceClient {
	return &hookServiceClient{cc}
}

func (c *hookServiceClient) GetMissionName(ctx context.Context, in *GetMissionNameRequest, opts ...grpc.CallOption) (*GetMissionNameResponse, error) {
	out := new(GetMissionNameResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/GetMissionName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) GetMissionFilename(ctx context.Context, in *GetMissionFilenameRequest, opts ...grpc.CallOption) (*GetMissionFilenameResponse, error) {
	out := new(GetMissionFilenameResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/GetMissionFilename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) GetMissionDescription(ctx context.Context, in *GetMissionDescriptionRequest, opts ...grpc.CallOption) (*GetMissionDescriptionResponse, error) {
	out := new(GetMissionDescriptionResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/GetMissionDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) GetPaused(ctx context.Context, in *GetPausedRequest, opts ...grpc.CallOption) (*GetPausedResponse, error) {
	out := new(GetPausedResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/GetPaused", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) SetPaused(ctx context.Context, in *SetPausedRequest, opts ...grpc.CallOption) (*SetPausedResponse, error) {
	out := new(SetPausedResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/SetPaused", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) StopMission(ctx context.Context, in *StopMissionRequest, opts ...grpc.CallOption) (*StopMissionResponse, error) {
	out := new(StopMissionResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/StopMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) ReloadCurrentMission(ctx context.Context, in *ReloadCurrentMissionRequest, opts ...grpc.CallOption) (*ReloadCurrentMissionResponse, error) {
	out := new(ReloadCurrentMissionResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/ReloadCurrentMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) LoadNextMission(ctx context.Context, in *LoadNextMissionRequest, opts ...grpc.CallOption) (*LoadNextMissionResponse, error) {
	out := new(LoadNextMissionResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/LoadNextMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) LoadMission(ctx context.Context, in *LoadMissionRequest, opts ...grpc.CallOption) (*LoadMissionResponse, error) {
	out := new(LoadMissionResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/LoadMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error) {
	out := new(EvalResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/Eval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) ExitProcess(ctx context.Context, in *ExitProcessRequest, opts ...grpc.CallOption) (*ExitProcessResponse, error) {
	out := new(ExitProcessResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/ExitProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) IsMultiplayer(ctx context.Context, in *IsMultiplayerRequest, opts ...grpc.CallOption) (*IsMultiplayerResponse, error) {
	out := new(IsMultiplayerResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/IsMultiplayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) IsServer(ctx context.Context, in *IsServerRequest, opts ...grpc.CallOption) (*IsServerResponse, error) {
	out := new(IsServerResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/IsServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) BanPlayer(ctx context.Context, in *BanPlayerRequest, opts ...grpc.CallOption) (*BanPlayerResponse, error) {
	out := new(BanPlayerResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/BanPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) UnbanPlayer(ctx context.Context, in *UnbanPlayerRequest, opts ...grpc.CallOption) (*UnbanPlayerResponse, error) {
	out := new(UnbanPlayerResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/UnbanPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) GetBannedPlayers(ctx context.Context, in *GetBannedPlayersRequest, opts ...grpc.CallOption) (*GetBannedPlayersResponse, error) {
	out := new(GetBannedPlayersResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/GetBannedPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) GetUnitType(ctx context.Context, in *GetUnitTypeRequest, opts ...grpc.CallOption) (*GetUnitTypeResponse, error) {
	out := new(GetUnitTypeResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/GetUnitType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) GetRealTime(ctx context.Context, in *GetRealTimeRequest, opts ...grpc.CallOption) (*GetRealTimeResponse, error) {
	out := new(GetRealTimeResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/GetRealTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) GetBallisticsCount(ctx context.Context, in *GetBallisticsCountRequest, opts ...grpc.CallOption) (*GetBallisticsCountResponse, error) {
	out := new(GetBallisticsCountResponse)
	err := c.cc.Invoke(ctx, "/dcs.hook.v0.HookService/GetBallisticsCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HookServiceServer is the server API for HookService service.
// All implementations must embed UnimplementedHookServiceServer
// for forward compatibility
type HookServiceServer interface {
	// https://wiki.hoggitworld.com/view/DCS_func_getMissionName
	GetMissionName(context.Context, *GetMissionNameRequest) (*GetMissionNameResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getMissionFilename
	GetMissionFilename(context.Context, *GetMissionFilenameRequest) (*GetMissionFilenameResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getMissionDescription
	GetMissionDescription(context.Context, *GetMissionDescriptionRequest) (*GetMissionDescriptionResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getPause
	GetPaused(context.Context, *GetPausedRequest) (*GetPausedResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_setPause
	SetPaused(context.Context, *SetPausedRequest) (*SetPausedResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_stopMission
	StopMission(context.Context, *StopMissionRequest) (*StopMissionResponse, error)
	// Reload the currently running mission
	ReloadCurrentMission(context.Context, *ReloadCurrentMissionRequest) (*ReloadCurrentMissionResponse, error)
	// Load the next mission in the server mission list. Note that it does
	// not loop back to the first mission once the end of the mission list
	// has been reached
	LoadNextMission(context.Context, *LoadNextMissionRequest) (*LoadNextMissionResponse, error)
	// Load a specific mission file. This does not need to be in the mission
	// list.
	LoadMission(context.Context, *LoadMissionRequest) (*LoadMissionResponse, error)
	// Evaluate some Lua inside of the hook environment and return the result as a
	// JSON string. Disabled by default.
	Eval(context.Context, *EvalRequest) (*EvalResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_exitProcess
	ExitProcess(context.Context, *ExitProcessRequest) (*ExitProcessResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_isMultiplayer
	IsMultiplayer(context.Context, *IsMultiplayerRequest) (*IsMultiplayerResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_isServer
	IsServer(context.Context, *IsServerRequest) (*IsServerResponse, error)
	// Bans a player that is currently connected to the server
	BanPlayer(context.Context, *BanPlayerRequest) (*BanPlayerResponse, error)
	// Unbans a player via their globally unique ID
	UnbanPlayer(context.Context, *UnbanPlayerRequest) (*UnbanPlayerResponse, error)
	// Get a list of all the banned players
	GetBannedPlayers(context.Context, *GetBannedPlayersRequest) (*GetBannedPlayersResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getUnitType
	GetUnitType(context.Context, *GetUnitTypeRequest) (*GetUnitTypeResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getRealTime
	GetRealTime(context.Context, *GetRealTimeRequest) (*GetRealTimeResponse, error)
	// Get a count of ballistics objects
	GetBallisticsCount(context.Context, *GetBallisticsCountRequest) (*GetBallisticsCountResponse, error)
	mustEmbedUnimplementedHookServiceServer()
}

// UnimplementedHookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHookServiceServer struct {
}

func (UnimplementedHookServiceServer) GetMissionName(context.Context, *GetMissionNameRequest) (*GetMissionNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMissionName not implemented")
}
func (UnimplementedHookServiceServer) GetMissionFilename(context.Context, *GetMissionFilenameRequest) (*GetMissionFilenameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMissionFilename not implemented")
}
func (UnimplementedHookServiceServer) GetMissionDescription(context.Context, *GetMissionDescriptionRequest) (*GetMissionDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMissionDescription not implemented")
}
func (UnimplementedHookServiceServer) GetPaused(context.Context, *GetPausedRequest) (*GetPausedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaused not implemented")
}
func (UnimplementedHookServiceServer) SetPaused(context.Context, *SetPausedRequest) (*SetPausedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPaused not implemented")
}
func (UnimplementedHookServiceServer) StopMission(context.Context, *StopMissionRequest) (*StopMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopMission not implemented")
}
func (UnimplementedHookServiceServer) ReloadCurrentMission(context.Context, *ReloadCurrentMissionRequest) (*ReloadCurrentMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadCurrentMission not implemented")
}
func (UnimplementedHookServiceServer) LoadNextMission(context.Context, *LoadNextMissionRequest) (*LoadNextMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadNextMission not implemented")
}
func (UnimplementedHookServiceServer) LoadMission(context.Context, *LoadMissionRequest) (*LoadMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadMission not implemented")
}
func (UnimplementedHookServiceServer) Eval(context.Context, *EvalRequest) (*EvalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Eval not implemented")
}
func (UnimplementedHookServiceServer) ExitProcess(context.Context, *ExitProcessRequest) (*ExitProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitProcess not implemented")
}
func (UnimplementedHookServiceServer) IsMultiplayer(context.Context, *IsMultiplayerRequest) (*IsMultiplayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsMultiplayer not implemented")
}
func (UnimplementedHookServiceServer) IsServer(context.Context, *IsServerRequest) (*IsServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsServer not implemented")
}
func (UnimplementedHookServiceServer) BanPlayer(context.Context, *BanPlayerRequest) (*BanPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanPlayer not implemented")
}
func (UnimplementedHookServiceServer) UnbanPlayer(context.Context, *UnbanPlayerRequest) (*UnbanPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbanPlayer not implemented")
}
func (UnimplementedHookServiceServer) GetBannedPlayers(context.Context, *GetBannedPlayersRequest) (*GetBannedPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBannedPlayers not implemented")
}
func (UnimplementedHookServiceServer) GetUnitType(context.Context, *GetUnitTypeRequest) (*GetUnitTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnitType not implemented")
}
func (UnimplementedHookServiceServer) GetRealTime(context.Context, *GetRealTimeRequest) (*GetRealTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealTime not implemented")
}
func (UnimplementedHookServiceServer) GetBallisticsCount(context.Context, *GetBallisticsCountRequest) (*GetBallisticsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBallisticsCount not implemented")
}
func (UnimplementedHookServiceServer) mustEmbedUnimplementedHookServiceServer() {}

// UnsafeHookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HookServiceServer will
// result in compilation errors.
type UnsafeHookServiceServer interface {
	mustEmbedUnimplementedHookServiceServer()
}

func RegisterHookServiceServer(s grpc.ServiceRegistrar, srv HookServiceServer) {
	s.RegisterService(&HookService_ServiceDesc, srv)
}

func _HookService_GetMissionName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMissionNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).GetMissionName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/GetMissionName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).GetMissionName(ctx, req.(*GetMissionNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_GetMissionFilename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMissionFilenameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).GetMissionFilename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/GetMissionFilename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).GetMissionFilename(ctx, req.(*GetMissionFilenameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_GetMissionDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMissionDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).GetMissionDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/GetMissionDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).GetMissionDescription(ctx, req.(*GetMissionDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_GetPaused_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPausedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).GetPaused(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/GetPaused",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).GetPaused(ctx, req.(*GetPausedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_SetPaused_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPausedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).SetPaused(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/SetPaused",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).SetPaused(ctx, req.(*SetPausedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_StopMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).StopMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/StopMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).StopMission(ctx, req.(*StopMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_ReloadCurrentMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadCurrentMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).ReloadCurrentMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/ReloadCurrentMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).ReloadCurrentMission(ctx, req.(*ReloadCurrentMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_LoadNextMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadNextMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).LoadNextMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/LoadNextMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).LoadNextMission(ctx, req.(*LoadNextMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_LoadMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).LoadMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/LoadMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).LoadMission(ctx, req.(*LoadMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_Eval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).Eval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/Eval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).Eval(ctx, req.(*EvalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_ExitProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExitProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).ExitProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/ExitProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).ExitProcess(ctx, req.(*ExitProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_IsMultiplayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsMultiplayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).IsMultiplayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/IsMultiplayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).IsMultiplayer(ctx, req.(*IsMultiplayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_IsServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).IsServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/IsServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).IsServer(ctx, req.(*IsServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_BanPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).BanPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/BanPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).BanPlayer(ctx, req.(*BanPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_UnbanPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbanPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).UnbanPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/UnbanPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).UnbanPlayer(ctx, req.(*UnbanPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_GetBannedPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBannedPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).GetBannedPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/GetBannedPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).GetBannedPlayers(ctx, req.(*GetBannedPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_GetUnitType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnitTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).GetUnitType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/GetUnitType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).GetUnitType(ctx, req.(*GetUnitTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_GetRealTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRealTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).GetRealTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/GetRealTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).GetRealTime(ctx, req.(*GetRealTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_GetBallisticsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBallisticsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).GetBallisticsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.hook.v0.HookService/GetBallisticsCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).GetBallisticsCount(ctx, req.(*GetBallisticsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HookService_ServiceDesc is the grpc.ServiceDesc for HookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dcs.hook.v0.HookService",
	HandlerType: (*HookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMissionName",
			Handler:    _HookService_GetMissionName_Handler,
		},
		{
			MethodName: "GetMissionFilename",
			Handler:    _HookService_GetMissionFilename_Handler,
		},
		{
			MethodName: "GetMissionDescription",
			Handler:    _HookService_GetMissionDescription_Handler,
		},
		{
			MethodName: "GetPaused",
			Handler:    _HookService_GetPaused_Handler,
		},
		{
			MethodName: "SetPaused",
			Handler:    _HookService_SetPaused_Handler,
		},
		{
			MethodName: "StopMission",
			Handler:    _HookService_StopMission_Handler,
		},
		{
			MethodName: "ReloadCurrentMission",
			Handler:    _HookService_ReloadCurrentMission_Handler,
		},
		{
			MethodName: "LoadNextMission",
			Handler:    _HookService_LoadNextMission_Handler,
		},
		{
			MethodName: "LoadMission",
			Handler:    _HookService_LoadMission_Handler,
		},
		{
			MethodName: "Eval",
			Handler:    _HookService_Eval_Handler,
		},
		{
			MethodName: "ExitProcess",
			Handler:    _HookService_ExitProcess_Handler,
		},
		{
			MethodName: "IsMultiplayer",
			Handler:    _HookService_IsMultiplayer_Handler,
		},
		{
			MethodName: "IsServer",
			Handler:    _HookService_IsServer_Handler,
		},
		{
			MethodName: "BanPlayer",
			Handler:    _HookService_BanPlayer_Handler,
		},
		{
			MethodName: "UnbanPlayer",
			Handler:    _HookService_UnbanPlayer_Handler,
		},
		{
			MethodName: "GetBannedPlayers",
			Handler:    _HookService_GetBannedPlayers_Handler,
		},
		{
			MethodName: "GetUnitType",
			Handler:    _HookService_GetUnitType_Handler,
		},
		{
			MethodName: "GetRealTime",
			Handler:    _HookService_GetRealTime_Handler,
		},
		{
			MethodName: "GetBallisticsCount",
			Handler:    _HookService_GetBallisticsCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcs/hook/v0/hook.proto",
}
