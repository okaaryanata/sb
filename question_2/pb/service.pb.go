// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Pagination           string   `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	SearchWord           string   `protobuf:"bytes,2,opt,name=search_word,proto3" json:"search_word,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPagination() string {
	if m != nil {
		return m.Pagination
	}
	return ""
}

func (m *Request) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

type Response struct {
	IsSuccess            bool              `protobuf:"varint,1,opt,name=is_success,proto3" json:"is_success,omitempty"`
	ErrorMessage         string            `protobuf:"bytes,2,opt,name=error_message,json=message,proto3" json:"error_message,omitempty"`
	TotalResults         string            `protobuf:"bytes,3,opt,name=totalResults,proto3" json:"totalResults,omitempty"`
	Response             string            `protobuf:"bytes,4,opt,name=Response,proto3" json:"Response,omitempty"`
	Search               []*Response_MOVIE `protobuf:"bytes,5,rep,name=Search,proto3" json:"Search,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *Response) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *Response) GetTotalResults() string {
	if m != nil {
		return m.TotalResults
	}
	return ""
}

func (m *Response) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *Response) GetSearch() []*Response_MOVIE {
	if m != nil {
		return m.Search
	}
	return nil
}

type Response_MOVIE struct {
	Title                string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year                 string   `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID               string   `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster               string   `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_MOVIE) Reset()         { *m = Response_MOVIE{} }
func (m *Response_MOVIE) String() string { return proto.CompactTextString(m) }
func (*Response_MOVIE) ProtoMessage()    {}
func (*Response_MOVIE) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{1, 0}
}

func (m *Response_MOVIE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_MOVIE.Unmarshal(m, b)
}
func (m *Response_MOVIE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_MOVIE.Marshal(b, m, deterministic)
}
func (m *Response_MOVIE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_MOVIE.Merge(m, src)
}
func (m *Response_MOVIE) XXX_Size() int {
	return xxx_messageInfo_Response_MOVIE.Size(m)
}
func (m *Response_MOVIE) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_MOVIE.DiscardUnknown(m)
}

var xxx_messageInfo_Response_MOVIE proto.InternalMessageInfo

func (m *Response_MOVIE) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Response_MOVIE) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *Response_MOVIE) GetImdbID() string {
	if m != nil {
		return m.ImdbID
	}
	return ""
}

func (m *Response_MOVIE) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Response_MOVIE) GetPoster() string {
	if m != nil {
		return m.Poster
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "pb.request")
	proto.RegisterType((*Response)(nil), "pb.response")
	proto.RegisterType((*Response_MOVIE)(nil), "pb.response.MOVIE")
}

func init() { proto.RegisterFile("pb/service.proto", fileDescriptor_6ff5ab49d8a5fcc4) }

var fileDescriptor_6ff5ab49d8a5fcc4 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xdf, 0x4a, 0xf3, 0x40,
	0x10, 0xc5, 0xbf, 0xfe, 0x49, 0x9b, 0x6f, 0x5a, 0x41, 0x06, 0x91, 0xa5, 0x17, 0xa5, 0xe4, 0xc6,
	0xe2, 0x45, 0x84, 0x7a, 0xe1, 0x0b, 0x28, 0x52, 0xa4, 0x28, 0x69, 0x11, 0xbc, 0x2a, 0x9b, 0x74,
	0xa8, 0x0b, 0x6d, 0x77, 0xdd, 0xd9, 0x54, 0x7c, 0x28, 0xdf, 0x51, 0xb2, 0xd9, 0x4a, 0xbc, 0x9b,
	0x73, 0xe6, 0x30, 0xfb, 0x5b, 0x0e, 0x9c, 0x9b, 0xfc, 0x86, 0xc9, 0x1e, 0x55, 0x41, 0xa9, 0xb1,
	0xda, 0x69, 0x6c, 0x9b, 0x3c, 0x79, 0x82, 0xbe, 0xa5, 0x8f, 0x92, 0xd8, 0xe1, 0x18, 0xc0, 0xc8,
	0xad, 0x3a, 0x48, 0xa7, 0xf4, 0x41, 0xb4, 0x26, 0xad, 0xe9, 0xff, 0xac, 0xe1, 0xe0, 0x04, 0x06,
	0x4c, 0xd2, 0x16, 0xef, 0xeb, 0x4f, 0x6d, 0x37, 0xa2, 0xed, 0x03, 0x4d, 0x2b, 0xf9, 0x6e, 0x43,
	0x6c, 0x89, 0x8d, 0x3e, 0x30, 0x55, 0xe7, 0x14, 0xaf, 0xb9, 0x2c, 0x0a, 0x62, 0xf6, 0xe7, 0xe2,
	0xac, 0xe1, 0xe0, 0x18, 0xce, 0xc8, 0x5a, 0x6d, 0xd7, 0x7b, 0x62, 0x96, 0x5b, 0x0a, 0x07, 0xfb,
	0x41, 0x62, 0x02, 0x43, 0xa7, 0x9d, 0xdc, 0x65, 0xc4, 0xe5, 0xce, 0xb1, 0xe8, 0xf8, 0xf5, 0x1f,
	0x0f, 0x47, 0x10, 0x67, 0xe1, 0x3d, 0xd1, 0xf5, 0xfb, 0x5f, 0x8d, 0xd7, 0xd0, 0x5b, 0x7a, 0x36,
	0x11, 0x4d, 0x3a, 0xd3, 0xc1, 0x0c, 0x53, 0x93, 0xa7, 0x27, 0xba, 0x74, 0xf1, 0xfc, 0x3a, 0x7f,
	0xc8, 0x42, 0x62, 0x54, 0x42, 0xe4, 0x0d, 0xbc, 0x80, 0x68, 0xa5, 0xdc, 0x8e, 0xc2, 0xf7, 0x6b,
	0x81, 0x08, 0xdd, 0x37, 0x92, 0x36, 0x10, 0xfa, 0x19, 0x2f, 0xa1, 0xa7, 0xf6, 0x9b, 0x7c, 0x7e,
	0x1f, 0xc0, 0x82, 0xaa, 0xb2, 0xab, 0x2f, 0x73, 0xc2, 0xf1, 0x73, 0x95, 0x7d, 0xd1, 0xec, 0xc8,
	0x8a, 0xa8, 0xce, 0xd6, 0x6a, 0x76, 0x07, 0xc3, 0x85, 0x3e, 0x2a, 0x5a, 0xd6, 0xb5, 0xe0, 0x15,
	0xc4, 0x8f, 0xe4, 0xbc, 0x85, 0x83, 0x1a, 0xd7, 0x57, 0x33, 0x1a, 0x36, 0xd9, 0x93, 0x7f, 0x79,
	0xcf, 0x17, 0x78, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xa6, 0x0f, 0xd6, 0xd4, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieServiceClient interface {
	GetMovie(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) GetMovie(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.MovieService/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
type MovieServiceServer interface {
	GetMovie(context.Context, *Request) (*Response, error)
}

// UnimplementedMovieServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (*UnimplementedMovieServiceServer) GetMovie(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}

func RegisterMovieServiceServer(s *grpc.Server, srv MovieServiceServer) {
	s.RegisterService(&_MovieService_serviceDesc, srv)
}

func _MovieService_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MovieService/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovie(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _MovieService_GetMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/service.proto",
}
