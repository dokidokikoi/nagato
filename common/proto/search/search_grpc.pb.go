// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: search.proto

package search

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

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchClient interface {
	GetDoc(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*DocResponse, error)
	CreateDocByID(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*Response, error)
	UpdateDoc(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*Response, error)
	DelDoc(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*Response, error)
	BulkDoc(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*Response, error)
	CreateIndex(ctx context.Context, in *IndexReqest, opts ...grpc.CallOption) (*Response, error)
	DelIndices(ctx context.Context, in *DelIndexReqest, opts ...grpc.CallOption) (*Response, error)
	SearchDoc(ctx context.Context, in *SearchReqest, opts ...grpc.CallOption) (*DocResponse, error)
}

type searchClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchClient(cc grpc.ClientConnInterface) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) GetDoc(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*DocResponse, error) {
	out := new(DocResponse)
	err := c.cc.Invoke(ctx, "/api.Search/GetDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) CreateDocByID(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Search/CreateDocByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) UpdateDoc(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Search/UpdateDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) DelDoc(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Search/DelDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) BulkDoc(ctx context.Context, in *DocReqest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Search/BulkDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) CreateIndex(ctx context.Context, in *IndexReqest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Search/CreateIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) DelIndices(ctx context.Context, in *DelIndexReqest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Search/DelIndices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) SearchDoc(ctx context.Context, in *SearchReqest, opts ...grpc.CallOption) (*DocResponse, error) {
	out := new(DocResponse)
	err := c.cc.Invoke(ctx, "/api.Search/SearchDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServer is the server API for Search service.
// All implementations must embed UnimplementedSearchServer
// for forward compatibility
type SearchServer interface {
	GetDoc(context.Context, *DocReqest) (*DocResponse, error)
	CreateDocByID(context.Context, *DocReqest) (*Response, error)
	UpdateDoc(context.Context, *DocReqest) (*Response, error)
	DelDoc(context.Context, *DocReqest) (*Response, error)
	BulkDoc(context.Context, *DocReqest) (*Response, error)
	CreateIndex(context.Context, *IndexReqest) (*Response, error)
	DelIndices(context.Context, *DelIndexReqest) (*Response, error)
	SearchDoc(context.Context, *SearchReqest) (*DocResponse, error)
	mustEmbedUnimplementedSearchServer()
}

// UnimplementedSearchServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServer struct {
}

func (UnimplementedSearchServer) GetDoc(context.Context, *DocReqest) (*DocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoc not implemented")
}
func (UnimplementedSearchServer) CreateDocByID(context.Context, *DocReqest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocByID not implemented")
}
func (UnimplementedSearchServer) UpdateDoc(context.Context, *DocReqest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoc not implemented")
}
func (UnimplementedSearchServer) DelDoc(context.Context, *DocReqest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelDoc not implemented")
}
func (UnimplementedSearchServer) BulkDoc(context.Context, *DocReqest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkDoc not implemented")
}
func (UnimplementedSearchServer) CreateIndex(context.Context, *IndexReqest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIndex not implemented")
}
func (UnimplementedSearchServer) DelIndices(context.Context, *DelIndexReqest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelIndices not implemented")
}
func (UnimplementedSearchServer) SearchDoc(context.Context, *SearchReqest) (*DocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDoc not implemented")
}
func (UnimplementedSearchServer) mustEmbedUnimplementedSearchServer() {}

// UnsafeSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServer will
// result in compilation errors.
type UnsafeSearchServer interface {
	mustEmbedUnimplementedSearchServer()
}

func RegisterSearchServer(s grpc.ServiceRegistrar, srv SearchServer) {
	s.RegisterService(&Search_ServiceDesc, srv)
}

func _Search_GetDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).GetDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Search/GetDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).GetDoc(ctx, req.(*DocReqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_CreateDocByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).CreateDocByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Search/CreateDocByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).CreateDocByID(ctx, req.(*DocReqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_UpdateDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).UpdateDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Search/UpdateDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).UpdateDoc(ctx, req.(*DocReqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_DelDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).DelDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Search/DelDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).DelDoc(ctx, req.(*DocReqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_BulkDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).BulkDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Search/BulkDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).BulkDoc(ctx, req.(*DocReqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_CreateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).CreateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Search/CreateIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).CreateIndex(ctx, req.(*IndexReqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_DelIndices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelIndexReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).DelIndices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Search/DelIndices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).DelIndices(ctx, req.(*DelIndexReqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_SearchDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).SearchDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Search/SearchDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).SearchDoc(ctx, req.(*SearchReqest))
	}
	return interceptor(ctx, in, info, handler)
}

// Search_ServiceDesc is the grpc.ServiceDesc for Search service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Search_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDoc",
			Handler:    _Search_GetDoc_Handler,
		},
		{
			MethodName: "CreateDocByID",
			Handler:    _Search_CreateDocByID_Handler,
		},
		{
			MethodName: "UpdateDoc",
			Handler:    _Search_UpdateDoc_Handler,
		},
		{
			MethodName: "DelDoc",
			Handler:    _Search_DelDoc_Handler,
		},
		{
			MethodName: "BulkDoc",
			Handler:    _Search_BulkDoc_Handler,
		},
		{
			MethodName: "CreateIndex",
			Handler:    _Search_CreateIndex_Handler,
		},
		{
			MethodName: "DelIndices",
			Handler:    _Search_DelIndices_Handler,
		},
		{
			MethodName: "SearchDoc",
			Handler:    _Search_SearchDoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}