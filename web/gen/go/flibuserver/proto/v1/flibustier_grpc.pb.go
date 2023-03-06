// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: flibuserver/proto/v1/flibustier.proto

package proto

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

// FlibustierServiceClient is the client API for FlibustierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlibustierServiceClient interface {
	GlobalSearch(ctx context.Context, in *GlobalSearchRequest, opts ...grpc.CallOption) (*GlobalSearchResponse, error)
	CheckUpdates(ctx context.Context, in *CheckUpdatesRequest, opts ...grpc.CallOption) (*CheckUpdatesResponse, error)
	GetSeriesBooks(ctx context.Context, in *GetSeriesBooksRequest, opts ...grpc.CallOption) (*GetSeriesBooksResponse, error)
	GetAuthorBooks(ctx context.Context, in *GetAuthorBooksRequest, opts ...grpc.CallOption) (*GetAuthorBooksResponse, error)
	TrackEntry(ctx context.Context, in *TrackEntryRequest, opts ...grpc.CallOption) (*TrackEntryResponse, error)
	ListTrackedEntries(ctx context.Context, in *ListTrackedEntriesRequest, opts ...grpc.CallOption) (*ListTrackedEntriesResponse, error)
	UntrackEntry(ctx context.Context, in *UntrackEntryRequest, opts ...grpc.CallOption) (*UntrackEntryResponse, error)
	UpdateEntry(ctx context.Context, in *UpdateTrackedEntryRequest, opts ...grpc.CallOption) (*UpdateTrackedEntryResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	// Added for testing only. Do not use in production.
	DeleteAllUsers(ctx context.Context, in *DeleteAllUsersRequest, opts ...grpc.CallOption) (*DeleteAllUsersResponse, error)
	DeleteAllTracked(ctx context.Context, in *DeleteAllTrackedRequest, opts ...grpc.CallOption) (*DeleteAllTrackedResponse, error)
}

type flibustierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlibustierServiceClient(cc grpc.ClientConnInterface) FlibustierServiceClient {
	return &flibustierServiceClient{cc}
}

func (c *flibustierServiceClient) GlobalSearch(ctx context.Context, in *GlobalSearchRequest, opts ...grpc.CallOption) (*GlobalSearchResponse, error) {
	out := new(GlobalSearchResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/GlobalSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) CheckUpdates(ctx context.Context, in *CheckUpdatesRequest, opts ...grpc.CallOption) (*CheckUpdatesResponse, error) {
	out := new(CheckUpdatesResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/CheckUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) GetSeriesBooks(ctx context.Context, in *GetSeriesBooksRequest, opts ...grpc.CallOption) (*GetSeriesBooksResponse, error) {
	out := new(GetSeriesBooksResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/GetSeriesBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) GetAuthorBooks(ctx context.Context, in *GetAuthorBooksRequest, opts ...grpc.CallOption) (*GetAuthorBooksResponse, error) {
	out := new(GetAuthorBooksResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/GetAuthorBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) TrackEntry(ctx context.Context, in *TrackEntryRequest, opts ...grpc.CallOption) (*TrackEntryResponse, error) {
	out := new(TrackEntryResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/TrackEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) ListTrackedEntries(ctx context.Context, in *ListTrackedEntriesRequest, opts ...grpc.CallOption) (*ListTrackedEntriesResponse, error) {
	out := new(ListTrackedEntriesResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/ListTrackedEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) UntrackEntry(ctx context.Context, in *UntrackEntryRequest, opts ...grpc.CallOption) (*UntrackEntryResponse, error) {
	out := new(UntrackEntryResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/UntrackEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) UpdateEntry(ctx context.Context, in *UpdateTrackedEntryRequest, opts ...grpc.CallOption) (*UpdateTrackedEntryResponse, error) {
	out := new(UpdateTrackedEntryResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/UpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) DeleteAllUsers(ctx context.Context, in *DeleteAllUsersRequest, opts ...grpc.CallOption) (*DeleteAllUsersResponse, error) {
	out := new(DeleteAllUsersResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/DeleteAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flibustierServiceClient) DeleteAllTracked(ctx context.Context, in *DeleteAllTrackedRequest, opts ...grpc.CallOption) (*DeleteAllTrackedResponse, error) {
	out := new(DeleteAllTrackedResponse)
	err := c.cc.Invoke(ctx, "/flibuserver.proto.v1.FlibustierService/DeleteAllTracked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlibustierServiceServer is the server API for FlibustierService service.
// All implementations must embed UnimplementedFlibustierServiceServer
// for forward compatibility
type FlibustierServiceServer interface {
	GlobalSearch(context.Context, *GlobalSearchRequest) (*GlobalSearchResponse, error)
	CheckUpdates(context.Context, *CheckUpdatesRequest) (*CheckUpdatesResponse, error)
	GetSeriesBooks(context.Context, *GetSeriesBooksRequest) (*GetSeriesBooksResponse, error)
	GetAuthorBooks(context.Context, *GetAuthorBooksRequest) (*GetAuthorBooksResponse, error)
	TrackEntry(context.Context, *TrackEntryRequest) (*TrackEntryResponse, error)
	ListTrackedEntries(context.Context, *ListTrackedEntriesRequest) (*ListTrackedEntriesResponse, error)
	UntrackEntry(context.Context, *UntrackEntryRequest) (*UntrackEntryResponse, error)
	UpdateEntry(context.Context, *UpdateTrackedEntryRequest) (*UpdateTrackedEntryResponse, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// Added for testing only. Do not use in production.
	DeleteAllUsers(context.Context, *DeleteAllUsersRequest) (*DeleteAllUsersResponse, error)
	DeleteAllTracked(context.Context, *DeleteAllTrackedRequest) (*DeleteAllTrackedResponse, error)
	mustEmbedUnimplementedFlibustierServiceServer()
}

// UnimplementedFlibustierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlibustierServiceServer struct {
}

func (UnimplementedFlibustierServiceServer) GlobalSearch(context.Context, *GlobalSearchRequest) (*GlobalSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalSearch not implemented")
}
func (UnimplementedFlibustierServiceServer) CheckUpdates(context.Context, *CheckUpdatesRequest) (*CheckUpdatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUpdates not implemented")
}
func (UnimplementedFlibustierServiceServer) GetSeriesBooks(context.Context, *GetSeriesBooksRequest) (*GetSeriesBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeriesBooks not implemented")
}
func (UnimplementedFlibustierServiceServer) GetAuthorBooks(context.Context, *GetAuthorBooksRequest) (*GetAuthorBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorBooks not implemented")
}
func (UnimplementedFlibustierServiceServer) TrackEntry(context.Context, *TrackEntryRequest) (*TrackEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackEntry not implemented")
}
func (UnimplementedFlibustierServiceServer) ListTrackedEntries(context.Context, *ListTrackedEntriesRequest) (*ListTrackedEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrackedEntries not implemented")
}
func (UnimplementedFlibustierServiceServer) UntrackEntry(context.Context, *UntrackEntryRequest) (*UntrackEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UntrackEntry not implemented")
}
func (UnimplementedFlibustierServiceServer) UpdateEntry(context.Context, *UpdateTrackedEntryRequest) (*UpdateTrackedEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntry not implemented")
}
func (UnimplementedFlibustierServiceServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedFlibustierServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedFlibustierServiceServer) DeleteAllUsers(context.Context, *DeleteAllUsersRequest) (*DeleteAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllUsers not implemented")
}
func (UnimplementedFlibustierServiceServer) DeleteAllTracked(context.Context, *DeleteAllTrackedRequest) (*DeleteAllTrackedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllTracked not implemented")
}
func (UnimplementedFlibustierServiceServer) mustEmbedUnimplementedFlibustierServiceServer() {}

// UnsafeFlibustierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlibustierServiceServer will
// result in compilation errors.
type UnsafeFlibustierServiceServer interface {
	mustEmbedUnimplementedFlibustierServiceServer()
}

func RegisterFlibustierServiceServer(s grpc.ServiceRegistrar, srv FlibustierServiceServer) {
	s.RegisterService(&FlibustierService_ServiceDesc, srv)
}

func _FlibustierService_GlobalSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).GlobalSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/GlobalSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).GlobalSearch(ctx, req.(*GlobalSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_CheckUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).CheckUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/CheckUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).CheckUpdates(ctx, req.(*CheckUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_GetSeriesBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSeriesBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).GetSeriesBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/GetSeriesBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).GetSeriesBooks(ctx, req.(*GetSeriesBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_GetAuthorBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).GetAuthorBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/GetAuthorBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).GetAuthorBooks(ctx, req.(*GetAuthorBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_TrackEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).TrackEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/TrackEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).TrackEntry(ctx, req.(*TrackEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_ListTrackedEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTrackedEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).ListTrackedEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/ListTrackedEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).ListTrackedEntries(ctx, req.(*ListTrackedEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_UntrackEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UntrackEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).UntrackEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/UntrackEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).UntrackEntry(ctx, req.(*UntrackEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrackedEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).UpdateEntry(ctx, req.(*UpdateTrackedEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_DeleteAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).DeleteAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/DeleteAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).DeleteAllUsers(ctx, req.(*DeleteAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlibustierService_DeleteAllTracked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllTrackedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlibustierServiceServer).DeleteAllTracked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flibuserver.proto.v1.FlibustierService/DeleteAllTracked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlibustierServiceServer).DeleteAllTracked(ctx, req.(*DeleteAllTrackedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlibustierService_ServiceDesc is the grpc.ServiceDesc for FlibustierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlibustierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flibuserver.proto.v1.FlibustierService",
	HandlerType: (*FlibustierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GlobalSearch",
			Handler:    _FlibustierService_GlobalSearch_Handler,
		},
		{
			MethodName: "CheckUpdates",
			Handler:    _FlibustierService_CheckUpdates_Handler,
		},
		{
			MethodName: "GetSeriesBooks",
			Handler:    _FlibustierService_GetSeriesBooks_Handler,
		},
		{
			MethodName: "GetAuthorBooks",
			Handler:    _FlibustierService_GetAuthorBooks_Handler,
		},
		{
			MethodName: "TrackEntry",
			Handler:    _FlibustierService_TrackEntry_Handler,
		},
		{
			MethodName: "ListTrackedEntries",
			Handler:    _FlibustierService_ListTrackedEntries_Handler,
		},
		{
			MethodName: "UntrackEntry",
			Handler:    _FlibustierService_UntrackEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _FlibustierService_UpdateEntry_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _FlibustierService_GetUserInfo_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _FlibustierService_ListUsers_Handler,
		},
		{
			MethodName: "DeleteAllUsers",
			Handler:    _FlibustierService_DeleteAllUsers_Handler,
		},
		{
			MethodName: "DeleteAllTracked",
			Handler:    _FlibustierService_DeleteAllTracked_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flibuserver/proto/v1/flibustier.proto",
}
