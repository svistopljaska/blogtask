// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: Blog.proto

package blogrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Posts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *Posts) Reset() {
	*x = Posts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Posts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Posts) ProtoMessage() {}

func (x *Posts) ProtoReflect() protoreflect.Message {
	mi := &file_Blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Posts.ProtoReflect.Descriptor instead.
func (*Posts) Descriptor() ([]byte, []int) {
	return file_Blog_proto_rawDescGZIP(), []int{0}
}

func (x *Posts) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *int64                 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Tags      []*Tag                 `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Createdon *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdon,proto3,oneof" json:"createdon,omitempty"`
	Author    string                 `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_Blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_Blog_proto_rawDescGZIP(), []int{1}
}

func (x *Post) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Post) GetCreatedon() *timestamppb.Timestamp {
	if x != nil {
		return x.Createdon
	}
	return nil
}

func (x *Post) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type PostId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PostId) Reset() {
	*x = PostId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostId) ProtoMessage() {}

func (x *PostId) ProtoReflect() protoreflect.Message {
	mi := &file_Blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostId.ProtoReflect.Descriptor instead.
func (*PostId) Descriptor() ([]byte, []int) {
	return file_Blog_proto_rawDescGZIP(), []int{2}
}

func (x *PostId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *int64 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_Blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_Blog_proto_rawDescGZIP(), []int{3}
}

func (x *Tag) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_Blog_proto protoreflect.FileDescriptor

var file_Blog_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6c,
	0x6f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x22, 0xd9, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x6f, 0x6e, 0x22, 0x18, 0x0a,
	0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x13,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x32, 0x68,
	0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x32, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x42, 0x12, 0x5a, 0x10, 0x62, 0x6c, 0x6f, 0x67,
	0x74, 0x61, 0x73, 0x6b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Blog_proto_rawDescOnce sync.Once
	file_Blog_proto_rawDescData = file_Blog_proto_rawDesc
)

func file_Blog_proto_rawDescGZIP() []byte {
	file_Blog_proto_rawDescOnce.Do(func() {
		file_Blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_Blog_proto_rawDescData)
	})
	return file_Blog_proto_rawDescData
}

var file_Blog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Blog_proto_goTypes = []interface{}{
	(*Posts)(nil),                 // 0: blogrpc.Posts
	(*Post)(nil),                  // 1: blogrpc.Post
	(*PostId)(nil),                // 2: blogrpc.PostId
	(*Tag)(nil),                   // 3: blogrpc.Tag
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_Blog_proto_depIdxs = []int32{
	1, // 0: blogrpc.Posts.posts:type_name -> blogrpc.Post
	3, // 1: blogrpc.Post.tags:type_name -> blogrpc.Tag
	4, // 2: blogrpc.Post.createdon:type_name -> google.protobuf.Timestamp
	5, // 3: blogrpc.Blog.getPosts:input_type -> google.protobuf.Empty
	1, // 4: blogrpc.Blog.createPost:input_type -> blogrpc.Post
	0, // 5: blogrpc.Blog.getPosts:output_type -> blogrpc.Posts
	2, // 6: blogrpc.Blog.createPost:output_type -> blogrpc.PostId
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Blog_proto_init() }
func file_Blog_proto_init() {
	if File_Blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Posts); i {
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
		file_Blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_Blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostId); i {
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
		file_Blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
	file_Blog_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_Blog_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Blog_proto_goTypes,
		DependencyIndexes: file_Blog_proto_depIdxs,
		MessageInfos:      file_Blog_proto_msgTypes,
	}.Build()
	File_Blog_proto = out.File
	file_Blog_proto_rawDesc = nil
	file_Blog_proto_goTypes = nil
	file_Blog_proto_depIdxs = nil
}