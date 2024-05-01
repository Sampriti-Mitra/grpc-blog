// ./proto/blogs/blog.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v5.26.1
// source: blogs/blog.proto

package blogs

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	date "google.golang.org/genproto/googleapis/type/date"
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

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId          string     `protobuf:"bytes,1,opt,name=post_id,proto3" json:"post_id,omitempty"`
	Title           string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content         string     `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Author          string     `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	PublicationDate *date.Date `protobuf:"bytes,5,opt,name=publication_date,proto3" json:"publication_date,omitempty"`
	Tags            []string   `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogs_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_blogs_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Blog) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Blog) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Blog) GetPublicationDate() *date.Date {
	if x != nil {
		return x.PublicationDate
	}
	return nil
}

func (x *Blog) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogs_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_blogs_blog_proto_rawDescGZIP(), []int{1}
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogs_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_blogs_blog_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Result) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BlogUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId  string   `protobuf:"bytes,1,opt,name=post_id,proto3" json:"post_id,omitempty"`
	Title   string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Author  string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Tags    []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *BlogUpdateRequest) Reset() {
	*x = BlogUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogs_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogUpdateRequest) ProtoMessage() {}

func (x *BlogUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogUpdateRequest.ProtoReflect.Descriptor instead.
func (*BlogUpdateRequest) Descriptor() ([]byte, []int) {
	return file_blogs_blog_proto_rawDescGZIP(), []int{3}
}

func (x *BlogUpdateRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *BlogUpdateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogUpdateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *BlogUpdateRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *BlogUpdateRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type BlogFetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,proto3" json:"post_id,omitempty"`
}

func (x *BlogFetchRequest) Reset() {
	*x = BlogFetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogs_blog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogFetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogFetchRequest) ProtoMessage() {}

func (x *BlogFetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogFetchRequest.ProtoReflect.Descriptor instead.
func (*BlogFetchRequest) Descriptor() ([]byte, []int) {
	return file_blogs_blog_proto_rawDescGZIP(), []int{4}
}

func (x *BlogFetchRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

var File_blogs_blog_proto protoreflect.FileDescriptor

var file_blogs_blog_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3c,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x89, 0x01, 0x0a,
	0x11, 0x42, 0x6c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x67,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x32, 0x89, 0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x12, 0x2d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x05, 0x2e, 0x42, 0x6c,
	0x6f, 0x67, 0x1a, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x12,
	0x41, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x42, 0x6c,
	0x6f, 0x67, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05,
	0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x32, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x45, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x42, 0x6c, 0x6f, 0x67,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x61, 0x6d, 0x70, 0x72, 0x69, 0x74, 0x69, 0x2d, 0x4d, 0x69, 0x74, 0x72, 0x61, 0x2f,
	0x62, 0x6c, 0x6f, 0x67, 0x2d, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blogs_blog_proto_rawDescOnce sync.Once
	file_blogs_blog_proto_rawDescData = file_blogs_blog_proto_rawDesc
)

func file_blogs_blog_proto_rawDescGZIP() []byte {
	file_blogs_blog_proto_rawDescOnce.Do(func() {
		file_blogs_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_blogs_blog_proto_rawDescData)
	})
	return file_blogs_blog_proto_rawDescData
}

var file_blogs_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_blogs_blog_proto_goTypes = []interface{}{
	(*Blog)(nil),              // 0: Blog
	(*Empty)(nil),             // 1: Empty
	(*Result)(nil),            // 2: Result
	(*BlogUpdateRequest)(nil), // 3: BlogUpdateRequest
	(*BlogFetchRequest)(nil),  // 4: BlogFetchRequest
	(*date.Date)(nil),         // 5: google.type.Date
}
var file_blogs_blog_proto_depIdxs = []int32{
	5, // 0: Blog.publication_date:type_name -> google.type.Date
	0, // 1: Blogs.AddPost:input_type -> Blog
	4, // 2: Blogs.ReadPost:input_type -> BlogFetchRequest
	3, // 3: Blogs.UpdatePost:input_type -> BlogUpdateRequest
	4, // 4: Blogs.DeletePost:input_type -> BlogFetchRequest
	0, // 5: Blogs.AddPost:output_type -> Blog
	0, // 6: Blogs.ReadPost:output_type -> Blog
	0, // 7: Blogs.UpdatePost:output_type -> Blog
	2, // 8: Blogs.DeletePost:output_type -> Result
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blogs_blog_proto_init() }
func file_blogs_blog_proto_init() {
	if File_blogs_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blogs_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
		file_blogs_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_blogs_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_blogs_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogUpdateRequest); i {
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
		file_blogs_blog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogFetchRequest); i {
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
			RawDescriptor: file_blogs_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blogs_blog_proto_goTypes,
		DependencyIndexes: file_blogs_blog_proto_depIdxs,
		MessageInfos:      file_blogs_blog_proto_msgTypes,
	}.Build()
	File_blogs_blog_proto = out.File
	file_blogs_blog_proto_rawDesc = nil
	file_blogs_blog_proto_goTypes = nil
	file_blogs_blog_proto_depIdxs = nil
}