// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: data.proto

package data

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

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataClient interface {
	// 创建临时文件信息文件
	CreateTempInfo(ctx context.Context, in *CreateTempInfoReq, opts ...grpc.CallOption) (*CreateTempInfoResp, error)
	// 上传临时文件
	UploadTempFile(ctx context.Context, opts ...grpc.CallOption) (Data_UploadTempFileClient, error)
	// 转正或删除临时文件
	CommitTempFile(ctx context.Context, in *CommitTempFileReq, opts ...grpc.CallOption) (*CommitTempFileResp, error)
	// 删除临时文件
	DeleteTempFile(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*DeleteTempFileResp, error)
	// 获取临时文件
	GetTempFile(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (Data_GetTempFileClient, error)
	// 获取已上传临时文件的大小
	HeadTempFile(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*HeadTempFileResp, error)
	// 获取文件
	GetMatter(ctx context.Context, in *GetMatterReq, opts ...grpc.CallOption) (Data_GetMatterClient, error)
}

type dataClient struct {
	cc grpc.ClientConnInterface
}

func NewDataClient(cc grpc.ClientConnInterface) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) CreateTempInfo(ctx context.Context, in *CreateTempInfoReq, opts ...grpc.CallOption) (*CreateTempInfoResp, error) {
	out := new(CreateTempInfoResp)
	err := c.cc.Invoke(ctx, "/api.Data/CreateTempInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) UploadTempFile(ctx context.Context, opts ...grpc.CallOption) (Data_UploadTempFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Data_ServiceDesc.Streams[0], "/api.Data/UploadTempFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataUploadTempFileClient{stream}
	return x, nil
}

type Data_UploadTempFileClient interface {
	Send(*UploadTempFileReq) error
	CloseAndRecv() (*UploadTempFileResp, error)
	grpc.ClientStream
}

type dataUploadTempFileClient struct {
	grpc.ClientStream
}

func (x *dataUploadTempFileClient) Send(m *UploadTempFileReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataUploadTempFileClient) CloseAndRecv() (*UploadTempFileResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadTempFileResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataClient) CommitTempFile(ctx context.Context, in *CommitTempFileReq, opts ...grpc.CallOption) (*CommitTempFileResp, error) {
	out := new(CommitTempFileResp)
	err := c.cc.Invoke(ctx, "/api.Data/CommitTempFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) DeleteTempFile(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*DeleteTempFileResp, error) {
	out := new(DeleteTempFileResp)
	err := c.cc.Invoke(ctx, "/api.Data/DeleteTempFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetTempFile(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (Data_GetTempFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Data_ServiceDesc.Streams[1], "/api.Data/GetTempFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataGetTempFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Data_GetTempFileClient interface {
	Recv() (*GetTempFileResp, error)
	grpc.ClientStream
}

type dataGetTempFileClient struct {
	grpc.ClientStream
}

func (x *dataGetTempFileClient) Recv() (*GetTempFileResp, error) {
	m := new(GetTempFileResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataClient) HeadTempFile(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*HeadTempFileResp, error) {
	out := new(HeadTempFileResp)
	err := c.cc.Invoke(ctx, "/api.Data/HeadTempFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetMatter(ctx context.Context, in *GetMatterReq, opts ...grpc.CallOption) (Data_GetMatterClient, error) {
	stream, err := c.cc.NewStream(ctx, &Data_ServiceDesc.Streams[2], "/api.Data/GetMatter", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataGetMatterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Data_GetMatterClient interface {
	Recv() (*GetMatterResp, error)
	grpc.ClientStream
}

type dataGetMatterClient struct {
	grpc.ClientStream
}

func (x *dataGetMatterClient) Recv() (*GetMatterResp, error) {
	m := new(GetMatterResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServer is the server API for Data service.
// All implementations must embed UnimplementedDataServer
// for forward compatibility
type DataServer interface {
	// 创建临时文件信息文件
	CreateTempInfo(context.Context, *CreateTempInfoReq) (*CreateTempInfoResp, error)
	// 上传临时文件
	UploadTempFile(Data_UploadTempFileServer) error
	// 转正或删除临时文件
	CommitTempFile(context.Context, *CommitTempFileReq) (*CommitTempFileResp, error)
	// 删除临时文件
	DeleteTempFile(context.Context, *CommonReq) (*DeleteTempFileResp, error)
	// 获取临时文件
	GetTempFile(*CommonReq, Data_GetTempFileServer) error
	// 获取已上传临时文件的大小
	HeadTempFile(context.Context, *CommonReq) (*HeadTempFileResp, error)
	// 获取文件
	GetMatter(*GetMatterReq, Data_GetMatterServer) error
	mustEmbedUnimplementedDataServer()
}

// UnimplementedDataServer must be embedded to have forward compatible implementations.
type UnimplementedDataServer struct {
}

func (UnimplementedDataServer) CreateTempInfo(context.Context, *CreateTempInfoReq) (*CreateTempInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTempInfo not implemented")
}
func (UnimplementedDataServer) UploadTempFile(Data_UploadTempFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadTempFile not implemented")
}
func (UnimplementedDataServer) CommitTempFile(context.Context, *CommitTempFileReq) (*CommitTempFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitTempFile not implemented")
}
func (UnimplementedDataServer) DeleteTempFile(context.Context, *CommonReq) (*DeleteTempFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTempFile not implemented")
}
func (UnimplementedDataServer) GetTempFile(*CommonReq, Data_GetTempFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTempFile not implemented")
}
func (UnimplementedDataServer) HeadTempFile(context.Context, *CommonReq) (*HeadTempFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadTempFile not implemented")
}
func (UnimplementedDataServer) GetMatter(*GetMatterReq, Data_GetMatterServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMatter not implemented")
}
func (UnimplementedDataServer) mustEmbedUnimplementedDataServer() {}

// UnsafeDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServer will
// result in compilation errors.
type UnsafeDataServer interface {
	mustEmbedUnimplementedDataServer()
}

func RegisterDataServer(s grpc.ServiceRegistrar, srv DataServer) {
	s.RegisterService(&Data_ServiceDesc, srv)
}

func _Data_CreateTempInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTempInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).CreateTempInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Data/CreateTempInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).CreateTempInfo(ctx, req.(*CreateTempInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_UploadTempFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServer).UploadTempFile(&dataUploadTempFileServer{stream})
}

type Data_UploadTempFileServer interface {
	SendAndClose(*UploadTempFileResp) error
	Recv() (*UploadTempFileReq, error)
	grpc.ServerStream
}

type dataUploadTempFileServer struct {
	grpc.ServerStream
}

func (x *dataUploadTempFileServer) SendAndClose(m *UploadTempFileResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataUploadTempFileServer) Recv() (*UploadTempFileReq, error) {
	m := new(UploadTempFileReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Data_CommitTempFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitTempFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).CommitTempFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Data/CommitTempFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).CommitTempFile(ctx, req.(*CommitTempFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_DeleteTempFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).DeleteTempFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Data/DeleteTempFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).DeleteTempFile(ctx, req.(*CommonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_GetTempFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommonReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServer).GetTempFile(m, &dataGetTempFileServer{stream})
}

type Data_GetTempFileServer interface {
	Send(*GetTempFileResp) error
	grpc.ServerStream
}

type dataGetTempFileServer struct {
	grpc.ServerStream
}

func (x *dataGetTempFileServer) Send(m *GetTempFileResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Data_HeadTempFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).HeadTempFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Data/HeadTempFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).HeadTempFile(ctx, req.(*CommonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_GetMatter_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMatterReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServer).GetMatter(m, &dataGetMatterServer{stream})
}

type Data_GetMatterServer interface {
	Send(*GetMatterResp) error
	grpc.ServerStream
}

type dataGetMatterServer struct {
	grpc.ServerStream
}

func (x *dataGetMatterServer) Send(m *GetMatterResp) error {
	return x.ServerStream.SendMsg(m)
}

// Data_ServiceDesc is the grpc.ServiceDesc for Data service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Data_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTempInfo",
			Handler:    _Data_CreateTempInfo_Handler,
		},
		{
			MethodName: "CommitTempFile",
			Handler:    _Data_CommitTempFile_Handler,
		},
		{
			MethodName: "DeleteTempFile",
			Handler:    _Data_DeleteTempFile_Handler,
		},
		{
			MethodName: "HeadTempFile",
			Handler:    _Data_HeadTempFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadTempFile",
			Handler:       _Data_UploadTempFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetTempFile",
			Handler:       _Data_GetTempFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMatter",
			Handler:       _Data_GetMatter_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "data.proto",
}