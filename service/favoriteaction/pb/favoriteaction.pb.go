// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: favoriteaction.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DouyinFavoriteActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     *int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`             //用户id
	Token      *string `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`                              //用户鉴权token
	VideoId    *int64  `protobuf:"varint,3,req,name=video_id,json=videoId" json:"video_id,omitempty"`          //视频id
	ActionType *int32  `protobuf:"varint,4,req,name=action_type,json=actionType" json:"action_type,omitempty"` //1-点赞，2-取消点赞
}

func (x *DouyinFavoriteActionRequest) Reset() {
	*x = DouyinFavoriteActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favoriteaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteActionRequest) ProtoMessage() {}

func (x *DouyinFavoriteActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favoriteaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteActionRequest) Descriptor() ([]byte, []int) {
	return file_favoriteaction_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinFavoriteActionRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *DouyinFavoriteActionRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *DouyinFavoriteActionRequest) GetVideoId() int64 {
	if x != nil && x.VideoId != nil {
		return *x.VideoId
	}
	return 0
}

func (x *DouyinFavoriteActionRequest) GetActionType() int32 {
	if x != nil && x.ActionType != nil {
		return *x.ActionType
	}
	return 0
}

type DouyinFavoriteActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` //状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     //返回状态描述
}

func (x *DouyinFavoriteActionResponse) Reset() {
	*x = DouyinFavoriteActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favoriteaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteActionResponse) ProtoMessage() {}

func (x *DouyinFavoriteActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favoriteaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteActionResponse) Descriptor() ([]byte, []int) {
	return file_favoriteaction_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinFavoriteActionResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinFavoriteActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

var File_favoriteaction_proto protoreflect.FileDescriptor

var file_favoriteaction_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0x8b, 0x01, 0x0a, 0x1e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x61,
	0x0a, 0x1f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x32, 0x9d, 0x01, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x8a, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62,
}

var (
	file_favoriteaction_proto_rawDescOnce sync.Once
	file_favoriteaction_proto_rawDescData = file_favoriteaction_proto_rawDesc
)

func file_favoriteaction_proto_rawDescGZIP() []byte {
	file_favoriteaction_proto_rawDescOnce.Do(func() {
		file_favoriteaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_favoriteaction_proto_rawDescData)
	})
	return file_favoriteaction_proto_rawDescData
}

var file_favoriteaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_favoriteaction_proto_goTypes = []interface{}{
	(*DouyinFavoriteActionRequest)(nil),  // 0: douyin.core.favoritelist.douyin_favorite_action_request
	(*DouyinFavoriteActionResponse)(nil), // 1: douyin.core.favoritelist.douyin_favorite_action_response
}
var file_favoriteaction_proto_depIdxs = []int32{
	0, // 0: douyin.core.favoritelist.FavoriteAction.GetFavoriteAction:input_type -> douyin.core.favoritelist.douyin_favorite_action_request
	1, // 1: douyin.core.favoritelist.FavoriteAction.GetFavoriteAction:output_type -> douyin.core.favoritelist.douyin_favorite_action_response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_favoriteaction_proto_init() }
func file_favoriteaction_proto_init() {
	if File_favoriteaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favoriteaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteActionRequest); i {
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
		file_favoriteaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteActionResponse); i {
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
			RawDescriptor: file_favoriteaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favoriteaction_proto_goTypes,
		DependencyIndexes: file_favoriteaction_proto_depIdxs,
		MessageInfos:      file_favoriteaction_proto_msgTypes,
	}.Build()
	File_favoriteaction_proto = out.File
	file_favoriteaction_proto_rawDesc = nil
	file_favoriteaction_proto_goTypes = nil
	file_favoriteaction_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FavoriteActionClient is the client API for FavoriteAction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FavoriteActionClient interface {
	GetFavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error)
}

type favoriteActionClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteActionClient(cc grpc.ClientConnInterface) FavoriteActionClient {
	return &favoriteActionClient{cc}
}

func (c *favoriteActionClient) GetFavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error) {
	out := new(DouyinFavoriteActionResponse)
	err := c.cc.Invoke(ctx, "/douyin.core.favoritelist.FavoriteAction/GetFavoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteActionServer is the server API for FavoriteAction service.
type FavoriteActionServer interface {
	GetFavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error)
}

// UnimplementedFavoriteActionServer can be embedded to have forward compatible implementations.
type UnimplementedFavoriteActionServer struct {
}

func (*UnimplementedFavoriteActionServer) GetFavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteAction not implemented")
}

func RegisterFavoriteActionServer(s *grpc.Server, srv FavoriteActionServer) {
	s.RegisterService(&_FavoriteAction_serviceDesc, srv)
}

func _FavoriteAction_GetFavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteActionServer).GetFavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyin.core.favoritelist.FavoriteAction/GetFavoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteActionServer).GetFavoriteAction(ctx, req.(*DouyinFavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FavoriteAction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.core.favoritelist.FavoriteAction",
	HandlerType: (*FavoriteActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFavoriteAction",
			Handler:    _FavoriteAction_GetFavoriteAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favoriteaction.proto",
}
