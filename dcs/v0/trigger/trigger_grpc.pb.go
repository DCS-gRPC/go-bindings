// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trigger

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

// TriggerServiceClient is the client API for TriggerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TriggerServiceClient interface {
	// https://wiki.hoggitworld.com/view/DCS_func_outText
	OutText(ctx context.Context, in *OutTextRequest, opts ...grpc.CallOption) (*OutTextResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_outTextForCoalition
	OutTextForCoalition(ctx context.Context, in *OutTextForCoalitionRequest, opts ...grpc.CallOption) (*OutTextForCoalitionResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_outTextForGroup
	OutTextForGroup(ctx context.Context, in *OutTextForGroupRequest, opts ...grpc.CallOption) (*OutTextForGroupResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getUserFlag
	GetUserFlag(ctx context.Context, in *GetUserFlagRequest, opts ...grpc.CallOption) (*GetUserFlagResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_setUserFlag
	SetUserFlag(ctx context.Context, in *SetUserFlagRequest, opts ...grpc.CallOption) (*SetUserFlagResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_markToAll
	MarkToAll(ctx context.Context, in *MarkToAllRequest, opts ...grpc.CallOption) (*MarkToAllResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_markToCoalition
	MarkToCoalition(ctx context.Context, in *MarkToCoalitionRequest, opts ...grpc.CallOption) (*MarkToCoalitionResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_markToGroup
	MarkToGroup(ctx context.Context, in *MarkToGroupRequest, opts ...grpc.CallOption) (*MarkToGroupResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_removeMark
	RemoveMark(ctx context.Context, in *RemoveMarkRequest, opts ...grpc.CallOption) (*RemoveMarkResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_explosion
	Explosion(ctx context.Context, in *ExplosionRequest, opts ...grpc.CallOption) (*ExplosionResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_smoke
	Smoke(ctx context.Context, in *SmokeRequest, opts ...grpc.CallOption) (*SmokeResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_illuminationBomb
	IlluminationBomb(ctx context.Context, in *IlluminationBombRequest, opts ...grpc.CallOption) (*IlluminationBombResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_signalFlare
	SignalFlare(ctx context.Context, in *SignalFlareRequest, opts ...grpc.CallOption) (*SignalFlareResponse, error)
}

type triggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTriggerServiceClient(cc grpc.ClientConnInterface) TriggerServiceClient {
	return &triggerServiceClient{cc}
}

func (c *triggerServiceClient) OutText(ctx context.Context, in *OutTextRequest, opts ...grpc.CallOption) (*OutTextResponse, error) {
	out := new(OutTextResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/OutText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) OutTextForCoalition(ctx context.Context, in *OutTextForCoalitionRequest, opts ...grpc.CallOption) (*OutTextForCoalitionResponse, error) {
	out := new(OutTextForCoalitionResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/OutTextForCoalition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) OutTextForGroup(ctx context.Context, in *OutTextForGroupRequest, opts ...grpc.CallOption) (*OutTextForGroupResponse, error) {
	out := new(OutTextForGroupResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/OutTextForGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) GetUserFlag(ctx context.Context, in *GetUserFlagRequest, opts ...grpc.CallOption) (*GetUserFlagResponse, error) {
	out := new(GetUserFlagResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/GetUserFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) SetUserFlag(ctx context.Context, in *SetUserFlagRequest, opts ...grpc.CallOption) (*SetUserFlagResponse, error) {
	out := new(SetUserFlagResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/SetUserFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) MarkToAll(ctx context.Context, in *MarkToAllRequest, opts ...grpc.CallOption) (*MarkToAllResponse, error) {
	out := new(MarkToAllResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/MarkToAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) MarkToCoalition(ctx context.Context, in *MarkToCoalitionRequest, opts ...grpc.CallOption) (*MarkToCoalitionResponse, error) {
	out := new(MarkToCoalitionResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/MarkToCoalition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) MarkToGroup(ctx context.Context, in *MarkToGroupRequest, opts ...grpc.CallOption) (*MarkToGroupResponse, error) {
	out := new(MarkToGroupResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/MarkToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) RemoveMark(ctx context.Context, in *RemoveMarkRequest, opts ...grpc.CallOption) (*RemoveMarkResponse, error) {
	out := new(RemoveMarkResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/RemoveMark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) Explosion(ctx context.Context, in *ExplosionRequest, opts ...grpc.CallOption) (*ExplosionResponse, error) {
	out := new(ExplosionResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/Explosion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) Smoke(ctx context.Context, in *SmokeRequest, opts ...grpc.CallOption) (*SmokeResponse, error) {
	out := new(SmokeResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/Smoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) IlluminationBomb(ctx context.Context, in *IlluminationBombRequest, opts ...grpc.CallOption) (*IlluminationBombResponse, error) {
	out := new(IlluminationBombResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/IlluminationBomb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) SignalFlare(ctx context.Context, in *SignalFlareRequest, opts ...grpc.CallOption) (*SignalFlareResponse, error) {
	out := new(SignalFlareResponse)
	err := c.cc.Invoke(ctx, "/dcs.trigger.v0.TriggerService/SignalFlare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TriggerServiceServer is the server API for TriggerService service.
// All implementations must embed UnimplementedTriggerServiceServer
// for forward compatibility
type TriggerServiceServer interface {
	// https://wiki.hoggitworld.com/view/DCS_func_outText
	OutText(context.Context, *OutTextRequest) (*OutTextResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_outTextForCoalition
	OutTextForCoalition(context.Context, *OutTextForCoalitionRequest) (*OutTextForCoalitionResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_outTextForGroup
	OutTextForGroup(context.Context, *OutTextForGroupRequest) (*OutTextForGroupResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getUserFlag
	GetUserFlag(context.Context, *GetUserFlagRequest) (*GetUserFlagResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_setUserFlag
	SetUserFlag(context.Context, *SetUserFlagRequest) (*SetUserFlagResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_markToAll
	MarkToAll(context.Context, *MarkToAllRequest) (*MarkToAllResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_markToCoalition
	MarkToCoalition(context.Context, *MarkToCoalitionRequest) (*MarkToCoalitionResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_markToGroup
	MarkToGroup(context.Context, *MarkToGroupRequest) (*MarkToGroupResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_removeMark
	RemoveMark(context.Context, *RemoveMarkRequest) (*RemoveMarkResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_explosion
	Explosion(context.Context, *ExplosionRequest) (*ExplosionResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_smoke
	Smoke(context.Context, *SmokeRequest) (*SmokeResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_illuminationBomb
	IlluminationBomb(context.Context, *IlluminationBombRequest) (*IlluminationBombResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_signalFlare
	SignalFlare(context.Context, *SignalFlareRequest) (*SignalFlareResponse, error)
	mustEmbedUnimplementedTriggerServiceServer()
}

// UnimplementedTriggerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTriggerServiceServer struct {
}

func (UnimplementedTriggerServiceServer) OutText(context.Context, *OutTextRequest) (*OutTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OutText not implemented")
}
func (UnimplementedTriggerServiceServer) OutTextForCoalition(context.Context, *OutTextForCoalitionRequest) (*OutTextForCoalitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OutTextForCoalition not implemented")
}
func (UnimplementedTriggerServiceServer) OutTextForGroup(context.Context, *OutTextForGroupRequest) (*OutTextForGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OutTextForGroup not implemented")
}
func (UnimplementedTriggerServiceServer) GetUserFlag(context.Context, *GetUserFlagRequest) (*GetUserFlagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFlag not implemented")
}
func (UnimplementedTriggerServiceServer) SetUserFlag(context.Context, *SetUserFlagRequest) (*SetUserFlagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserFlag not implemented")
}
func (UnimplementedTriggerServiceServer) MarkToAll(context.Context, *MarkToAllRequest) (*MarkToAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkToAll not implemented")
}
func (UnimplementedTriggerServiceServer) MarkToCoalition(context.Context, *MarkToCoalitionRequest) (*MarkToCoalitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkToCoalition not implemented")
}
func (UnimplementedTriggerServiceServer) MarkToGroup(context.Context, *MarkToGroupRequest) (*MarkToGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkToGroup not implemented")
}
func (UnimplementedTriggerServiceServer) RemoveMark(context.Context, *RemoveMarkRequest) (*RemoveMarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMark not implemented")
}
func (UnimplementedTriggerServiceServer) Explosion(context.Context, *ExplosionRequest) (*ExplosionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Explosion not implemented")
}
func (UnimplementedTriggerServiceServer) Smoke(context.Context, *SmokeRequest) (*SmokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Smoke not implemented")
}
func (UnimplementedTriggerServiceServer) IlluminationBomb(context.Context, *IlluminationBombRequest) (*IlluminationBombResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IlluminationBomb not implemented")
}
func (UnimplementedTriggerServiceServer) SignalFlare(context.Context, *SignalFlareRequest) (*SignalFlareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalFlare not implemented")
}
func (UnimplementedTriggerServiceServer) mustEmbedUnimplementedTriggerServiceServer() {}

// UnsafeTriggerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TriggerServiceServer will
// result in compilation errors.
type UnsafeTriggerServiceServer interface {
	mustEmbedUnimplementedTriggerServiceServer()
}

func RegisterTriggerServiceServer(s grpc.ServiceRegistrar, srv TriggerServiceServer) {
	s.RegisterService(&TriggerService_ServiceDesc, srv)
}

func _TriggerService_OutText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).OutText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/OutText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).OutText(ctx, req.(*OutTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_OutTextForCoalition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutTextForCoalitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).OutTextForCoalition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/OutTextForCoalition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).OutTextForCoalition(ctx, req.(*OutTextForCoalitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_OutTextForGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutTextForGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).OutTextForGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/OutTextForGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).OutTextForGroup(ctx, req.(*OutTextForGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_GetUserFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).GetUserFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/GetUserFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).GetUserFlag(ctx, req.(*GetUserFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_SetUserFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).SetUserFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/SetUserFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).SetUserFlag(ctx, req.(*SetUserFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_MarkToAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkToAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).MarkToAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/MarkToAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).MarkToAll(ctx, req.(*MarkToAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_MarkToCoalition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkToCoalitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).MarkToCoalition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/MarkToCoalition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).MarkToCoalition(ctx, req.(*MarkToCoalitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_MarkToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).MarkToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/MarkToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).MarkToGroup(ctx, req.(*MarkToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_RemoveMark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).RemoveMark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/RemoveMark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).RemoveMark(ctx, req.(*RemoveMarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_Explosion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExplosionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).Explosion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/Explosion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).Explosion(ctx, req.(*ExplosionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_Smoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).Smoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/Smoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).Smoke(ctx, req.(*SmokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_IlluminationBomb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IlluminationBombRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).IlluminationBomb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/IlluminationBomb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).IlluminationBomb(ctx, req.(*IlluminationBombRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_SignalFlare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalFlareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).SignalFlare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.trigger.v0.TriggerService/SignalFlare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).SignalFlare(ctx, req.(*SignalFlareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TriggerService_ServiceDesc is the grpc.ServiceDesc for TriggerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TriggerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dcs.trigger.v0.TriggerService",
	HandlerType: (*TriggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OutText",
			Handler:    _TriggerService_OutText_Handler,
		},
		{
			MethodName: "OutTextForCoalition",
			Handler:    _TriggerService_OutTextForCoalition_Handler,
		},
		{
			MethodName: "OutTextForGroup",
			Handler:    _TriggerService_OutTextForGroup_Handler,
		},
		{
			MethodName: "GetUserFlag",
			Handler:    _TriggerService_GetUserFlag_Handler,
		},
		{
			MethodName: "SetUserFlag",
			Handler:    _TriggerService_SetUserFlag_Handler,
		},
		{
			MethodName: "MarkToAll",
			Handler:    _TriggerService_MarkToAll_Handler,
		},
		{
			MethodName: "MarkToCoalition",
			Handler:    _TriggerService_MarkToCoalition_Handler,
		},
		{
			MethodName: "MarkToGroup",
			Handler:    _TriggerService_MarkToGroup_Handler,
		},
		{
			MethodName: "RemoveMark",
			Handler:    _TriggerService_RemoveMark_Handler,
		},
		{
			MethodName: "Explosion",
			Handler:    _TriggerService_Explosion_Handler,
		},
		{
			MethodName: "Smoke",
			Handler:    _TriggerService_Smoke_Handler,
		},
		{
			MethodName: "IlluminationBomb",
			Handler:    _TriggerService_IlluminationBomb_Handler,
		},
		{
			MethodName: "SignalFlare",
			Handler:    _TriggerService_SignalFlare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcs/trigger/v0/trigger.proto",
}
