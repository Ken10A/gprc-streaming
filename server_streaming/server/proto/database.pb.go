// Code generated by protoc-gen-go. DO NOT EDIT.
// source: database.proto

/*
Package practical_grpc_v1 is a generated protocol buffer package.

It is generated from these files:
	database.proto

It has these top-level messages:
	SearchRequest
	SearchResponse
*/
package practical_grpc_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SearchRequest struct {
	Term       string `protobuf:"bytes,1,opt,name=term" json:"term,omitempty"`
	MaxResults int64  `protobuf:"varint,2,opt,name=max_results,json=maxResults" json:"max_results,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SearchRequest) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

func (m *SearchRequest) GetMaxResults() int64 {
	if m != nil {
		return m.MaxResults
	}
	return 0
}

type SearchResponse struct {
	MatchedTerm string `protobuf:"bytes,1,opt,name=matched_term,json=matchedTerm" json:"matched_term,omitempty"`
	Content     string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	Rank        int32  `protobuf:"varint,3,opt,name=rank" json:"rank,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchResponse) GetMatchedTerm() string {
	if m != nil {
		return m.MatchedTerm
	}
	return ""
}

func (m *SearchResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SearchResponse) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "practical.grpc.v1.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "practical.grpc.v1.SearchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Database service

type DatabaseClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type databaseClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseClient(cc *grpc.ClientConn) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/practical.grpc.v1.Database/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Database service

type DatabaseServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

func _Database_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/practical.grpc.v1.Database/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "practical.grpc.v1.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Database_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "database.proto",
}

func init() { proto.RegisterFile("database.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x89, 0xd5, 0x6a, 0xa7, 0x5a, 0x30, 0xa7, 0xe0, 0xc5, 0xb4, 0xa7, 0x3d, 0x05, 0xd4,
	0xaf, 0xd0, 0xab, 0x97, 0xe8, 0xc5, 0x53, 0x99, 0xa6, 0x83, 0x15, 0x9b, 0x3f, 0x26, 0x53, 0xd9,
	0x8f, 0x2f, 0x9b, 0xdd, 0x05, 0x45, 0xe8, 0xed, 0xcd, 0x63, 0xf8, 0xcd, 0x7b, 0x03, 0x8b, 0x1d,
	0x32, 0x6e, 0xb1, 0x90, 0x49, 0x39, 0x72, 0x94, 0xb7, 0x29, 0xa3, 0xe3, 0x0f, 0x87, 0x07, 0xf3,
	0x9e, 0x93, 0x33, 0xdf, 0x0f, 0xab, 0x35, 0xdc, 0xbc, 0x10, 0x66, 0xb7, 0xb7, 0xf4, 0x75, 0xa4,
	0xc2, 0x52, 0xc2, 0x39, 0x53, 0xf6, 0x4a, 0x68, 0xd1, 0xcc, 0x6c, 0xd5, 0xf2, 0x1e, 0xe6, 0x1e,
	0xdb, 0x4d, 0xa6, 0x72, 0x3c, 0x70, 0x51, 0x67, 0x5a, 0x34, 0x13, 0x0b, 0x1e, 0x5b, 0xdb, 0x3b,
	0x2b, 0x84, 0xc5, 0x48, 0x29, 0x29, 0x86, 0x42, 0x72, 0x09, 0xd7, 0x1e, 0xd9, 0xed, 0x69, 0xb7,
	0xf9, 0x85, 0x9b, 0x0f, 0xde, 0x6b, 0x47, 0x55, 0x70, 0xe9, 0x62, 0x60, 0x0a, 0x5c, 0x89, 0x33,
	0x3b, 0x8e, 0x5d, 0x86, 0x8c, 0xe1, 0x53, 0x4d, 0xb4, 0x68, 0x2e, 0x6c, 0xd5, 0x8f, 0x6f, 0x70,
	0xb5, 0x1e, 0xda, 0xc8, 0x67, 0x98, 0xf6, 0xe7, 0xa4, 0x36, 0xff, 0x2a, 0x99, 0x3f, 0x7d, 0xee,
	0x96, 0x27, 0x36, 0xfa, 0xac, 0xdb, 0x69, 0xfd, 0xce, 0xd3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0x57, 0x22, 0x2e, 0x2f, 0x01, 0x00, 0x00,
}