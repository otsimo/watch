// Code generated by protoc-gen-gogo.
// source: search.proto
// DO NOT EDIT!

package otsimopb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for SearchService service

type SearchServiceClient interface {
	IndexDatabase(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*Response, error)
	ReindexAll(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*Response, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) IndexDatabase(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.SearchService/IndexDatabase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) ReindexAll(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.SearchService/ReindexAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/apipb.SearchService/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SearchService service

type SearchServiceServer interface {
	IndexDatabase(context.Context, *IndexRequest) (*Response, error)
	ReindexAll(context.Context, *IndexRequest) (*Response, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_IndexDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).IndexDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.SearchService/IndexDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).IndexDatabase(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_ReindexAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).ReindexAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.SearchService/ReindexAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).ReindexAll(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IndexDatabase",
			Handler:    _SearchService_IndexDatabase_Handler,
		},
		{
			MethodName: "ReindexAll",
			Handler:    _SearchService_ReindexAll_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptorSearch = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc8, 0x2c, 0x48, 0x92,
	0xe2, 0xcb, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x86, 0x08, 0x4b, 0xe9, 0xa6, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x85, 0x93,
	0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x28, 0x37, 0xda, 0xce, 0xc8, 0xc5, 0x1b, 0x0c,
	0x36, 0x36, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x9c, 0x8b, 0xd7, 0x33, 0x2f, 0x25,
	0xb5, 0xc2, 0x25, 0xb1, 0x24, 0x31, 0x29, 0xb1, 0x38, 0x55, 0x48, 0x58, 0x0f, 0x6c, 0x93, 0x1e,
	0x58, 0x34, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x8a, 0x1f, 0x2a, 0x18, 0x94, 0x5a, 0x5c,
	0x90, 0x9f, 0x57, 0x9c, 0xaa, 0xc4, 0x20, 0x64, 0xc2, 0xc5, 0x15, 0x94, 0x9a, 0x09, 0x52, 0xe4,
	0x98, 0x93, 0x43, 0xb4, 0x2e, 0x73, 0x2e, 0x36, 0x88, 0xfd, 0x42, 0x22, 0x50, 0x49, 0x08, 0x17,
	0xa6, 0x45, 0x14, 0x4d, 0x14, 0xa6, 0xd1, 0xc9, 0xfa, 0xc4, 0x43, 0x39, 0x86, 0x0b, 0x40, 0x7c,
	0xe2, 0x91, 0x1c, 0xe3, 0x05, 0x20, 0x7e, 0x00, 0xc4, 0x13, 0x1e, 0xcb, 0x31, 0x70, 0xf1, 0x03,
	0xfd, 0xae, 0x97, 0x5f, 0x52, 0x9c, 0x99, 0x9b, 0xaf, 0x97, 0x5e, 0x54, 0x90, 0x1c, 0xc0, 0x18,
	0xc5, 0x01, 0xe1, 0x16, 0x24, 0x2d, 0x62, 0x62, 0xf6, 0x0f, 0x09, 0x4e, 0x62, 0x03, 0xfb, 0xde,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x21, 0xb7, 0xc5, 0x5c, 0x53, 0x01, 0x00, 0x00,
}
