// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.10.1
// source: Course.proto

package Course

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CourseModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId       int32   `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	CourseName     string  `protobuf:"bytes,2,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
	CourseDispName string  `protobuf:"bytes,3,opt,name=course_disp_name,json=courseDispName,proto3" json:"course_disp_name,omitempty"`
	CourseIntr     string  `protobuf:"bytes,4,opt,name=course_intr,json=courseIntr,proto3" json:"course_intr,omitempty"`
	CoursePrice    float32 `protobuf:"fixed32,5,opt,name=course_price,json=coursePrice,proto3" json:"course_price,omitempty"`
	CoursePrice2   float32 `protobuf:"fixed32,6,opt,name=course_price2,json=coursePrice2,proto3" json:"course_price2,omitempty"`
	// @inject_tag: gorm:"type:timestamp"
	Addtime *Timestamp `protobuf:"bytes,7,opt,name=addtime,proto3" json:"addtime,omitempty" gorm:"type:timestamp"`
}

func (x *CourseModel) Reset() {
	*x = CourseModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseModel) ProtoMessage() {}

func (x *CourseModel) ProtoReflect() protoreflect.Message {
	mi := &file_Course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseModel.ProtoReflect.Descriptor instead.
func (*CourseModel) Descriptor() ([]byte, []int) {
	return file_Course_proto_rawDescGZIP(), []int{0}
}

func (x *CourseModel) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *CourseModel) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *CourseModel) GetCourseDispName() string {
	if x != nil {
		return x.CourseDispName
	}
	return ""
}

func (x *CourseModel) GetCourseIntr() string {
	if x != nil {
		return x.CourseIntr
	}
	return ""
}

func (x *CourseModel) GetCoursePrice() float32 {
	if x != nil {
		return x.CoursePrice
	}
	return 0
}

func (x *CourseModel) GetCoursePrice2() float32 {
	if x != nil {
		return x.CoursePrice2
	}
	return 0
}

func (x *CourseModel) GetAddtime() *Timestamp {
	if x != nil {
		return x.Addtime
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: form:"size"
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty" form:"size"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_Course_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"course_id"
	CourseId int32 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty" uri:"course_id"`
	// @inject_tag: header:"fetch_type"
	FetchType int32 `protobuf:"varint,2,opt,name=fetch_type,json=fetchType,proto3" json:"fetch_type,omitempty" header:"fetch_type"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_Course_proto_rawDescGZIP(), []int{2}
}

func (x *DetailRequest) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *DetailRequest) GetFetchType() int32 {
	if x != nil {
		return x.FetchType
	}
	return 0
}

//计数
type CourseCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"-"
	CountId int32 `protobuf:"varint,1,opt,name=count_id,json=countId,proto3" json:"-"`
	// @inject_tag: json:"-"
	CourseId   int32  `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"-"`
	CountKey   string `protobuf:"bytes,3,opt,name=count_key,json=countKey,proto3" json:"count_key,omitempty"`
	CountValue int32  `protobuf:"varint,4,opt,name=count_value,json=countValue,proto3" json:"count_value,omitempty"`
}

func (x *CourseCounts) Reset() {
	*x = CourseCounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Course_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseCounts) ProtoMessage() {}

func (x *CourseCounts) ProtoReflect() protoreflect.Message {
	mi := &file_Course_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseCounts.ProtoReflect.Descriptor instead.
func (*CourseCounts) Descriptor() ([]byte, []int) {
	return file_Course_proto_rawDescGZIP(), []int{3}
}

func (x *CourseCounts) GetCountId() int32 {
	if x != nil {
		return x.CountId
	}
	return 0
}

func (x *CourseCounts) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *CourseCounts) GetCountKey() string {
	if x != nil {
		return x.CountKey
	}
	return ""
}

func (x *CourseCounts) GetCountValue() int32 {
	if x != nil {
		return x.CountValue
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CourseModel `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Course_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Course_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_Course_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetResult() []*CourseModel {
	if x != nil {
		return x.Result
	}
	return nil
}

type DetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *CourseModel    `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
	Counts []*CourseCounts `protobuf:"bytes,2,rep,name=counts,proto3" json:"counts,omitempty"`
}

func (x *DetailResponse) Reset() {
	*x = DetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Course_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse) ProtoMessage() {}

func (x *DetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Course_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse.ProtoReflect.Descriptor instead.
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return file_Course_proto_rawDescGZIP(), []int{5}
}

func (x *DetailResponse) GetCourse() *CourseModel {
	if x != nil {
		return x.Course
	}
	return nil
}

func (x *DetailResponse) GetCounts() []*CourseCounts {
	if x != nil {
		return x.Counts
	}
	return nil
}

var File_Course_proto protoreflect.FileDescriptor

var file_Course_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x64, 0x69, 0x73,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x44, 0x69, 0x73, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x32, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x64, 0x64, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x4b, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6b, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x32, 0x84, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72,
	0x54, 0x6f, 0x70, 0x12, 0x13, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x15, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Course_proto_rawDescOnce sync.Once
	file_Course_proto_rawDescData = file_Course_proto_rawDesc
)

func file_Course_proto_rawDescGZIP() []byte {
	file_Course_proto_rawDescOnce.Do(func() {
		file_Course_proto_rawDescData = protoimpl.X.CompressGZIP(file_Course_proto_rawDescData)
	})
	return file_Course_proto_rawDescData
}

var file_Course_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Course_proto_goTypes = []interface{}{
	(*CourseModel)(nil),    // 0: Course.CourseModel
	(*ListRequest)(nil),    // 1: Course.ListRequest
	(*DetailRequest)(nil),  // 2: Course.DetailRequest
	(*CourseCounts)(nil),   // 3: Course.CourseCounts
	(*ListResponse)(nil),   // 4: Course.ListResponse
	(*DetailResponse)(nil), // 5: Course.DetailResponse
	(*Timestamp)(nil),      // 6: Course.Timestamp
}
var file_Course_proto_depIdxs = []int32{
	6, // 0: Course.CourseModel.addtime:type_name -> Course.Timestamp
	0, // 1: Course.ListResponse.result:type_name -> Course.CourseModel
	0, // 2: Course.DetailResponse.course:type_name -> Course.CourseModel
	3, // 3: Course.DetailResponse.counts:type_name -> Course.CourseCounts
	1, // 4: Course.CourseService.ListForTop:input_type -> Course.ListRequest
	2, // 5: Course.CourseService.GetDetail:input_type -> Course.DetailRequest
	4, // 6: Course.CourseService.ListForTop:output_type -> Course.ListResponse
	5, // 7: Course.CourseService.GetDetail:output_type -> Course.DetailResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_Course_proto_init() }
func file_Course_proto_init() {
	if File_Course_proto != nil {
		return
	}
	file_Common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseModel); i {
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
		file_Course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_Course_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRequest); i {
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
		file_Course_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseCounts); i {
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
		file_Course_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_Course_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailResponse); i {
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
			RawDescriptor: file_Course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Course_proto_goTypes,
		DependencyIndexes: file_Course_proto_depIdxs,
		MessageInfos:      file_Course_proto_msgTypes,
	}.Build()
	File_Course_proto = out.File
	file_Course_proto_rawDesc = nil
	file_Course_proto_goTypes = nil
	file_Course_proto_depIdxs = nil
}