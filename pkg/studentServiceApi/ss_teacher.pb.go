// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: ss_teacher.proto

package student

import (
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

type PositionType int32

const (
	PositionType_POSTGRADUATE PositionType = 0
	PositionType_ASSISTANT    PositionType = 1
	PositionType_DEAN         PositionType = 2
)

// Enum value maps for PositionType.
var (
	PositionType_name = map[int32]string{
		0: "POSTGRADUATE",
		1: "ASSISTANT",
		2: "DEAN",
	}
	PositionType_value = map[string]int32{
		"POSTGRADUATE": 0,
		"ASSISTANT":    1,
		"DEAN":         2,
	}
)

func (x PositionType) Enum() *PositionType {
	p := new(PositionType)
	*p = x
	return p
}

func (x PositionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PositionType) Descriptor() protoreflect.EnumDescriptor {
	return file_ss_teacher_proto_enumTypes[0].Descriptor()
}

func (PositionType) Type() protoreflect.EnumType {
	return &file_ss_teacher_proto_enumTypes[0]
}

func (x PositionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PositionType.Descriptor instead.
func (PositionType) EnumDescriptor() ([]byte, []int) {
	return file_ss_teacher_proto_rawDescGZIP(), []int{0}
}

type CreateTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionType PositionType `protobuf:"varint,1,opt,name=position_type,json=positionType,proto3,enum=student.PositionType" json:"position_type,omitempty"`
	FullName     string       `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	StudentId    string       `protobuf:"bytes,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
}

func (x *CreateTeacherRequest) Reset() {
	*x = CreateTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_teacher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeacherRequest) ProtoMessage() {}

func (x *CreateTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ss_teacher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeacherRequest.ProtoReflect.Descriptor instead.
func (*CreateTeacherRequest) Descriptor() ([]byte, []int) {
	return file_ss_teacher_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTeacherRequest) GetPositionType() PositionType {
	if x != nil {
		return x.PositionType
	}
	return PositionType_POSTGRADUATE
}

func (x *CreateTeacherRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateTeacherRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PositionType PositionType `protobuf:"varint,2,opt,name=position_type,json=positionType,proto3,enum=student.PositionType" json:"position_type,omitempty"`
	FullName     string       `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	StudentId    string       `protobuf:"bytes,4,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_teacher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_ss_teacher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_ss_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *Teacher) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Teacher) GetPositionType() PositionType {
	if x != nil {
		return x.PositionType
	}
	return PositionType_POSTGRADUATE
}

func (x *Teacher) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Teacher) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type UpdateTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PositionType PositionType `protobuf:"varint,2,opt,name=position_type,json=positionType,proto3,enum=student.PositionType" json:"position_type,omitempty"`
	FullName     string       `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
}

func (x *UpdateTeacherRequest) Reset() {
	*x = UpdateTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_teacher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeacherRequest) ProtoMessage() {}

func (x *UpdateTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ss_teacher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeacherRequest.ProtoReflect.Descriptor instead.
func (*UpdateTeacherRequest) Descriptor() ([]byte, []int) {
	return file_ss_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTeacherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTeacherRequest) GetPositionType() PositionType {
	if x != nil {
		return x.PositionType
	}
	return PositionType_POSTGRADUATE
}

func (x *UpdateTeacherRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

type ListTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherIds []string `protobuf:"bytes,1,rep,name=teacher_ids,json=teacherIds,proto3" json:"teacher_ids,omitempty"`
}

func (x *ListTeacherRequest) Reset() {
	*x = ListTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_teacher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeacherRequest) ProtoMessage() {}

func (x *ListTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ss_teacher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeacherRequest.ProtoReflect.Descriptor instead.
func (*ListTeacherRequest) Descriptor() ([]byte, []int) {
	return file_ss_teacher_proto_rawDescGZIP(), []int{3}
}

func (x *ListTeacherRequest) GetTeacherIds() []string {
	if x != nil {
		return x.TeacherIds
	}
	return nil
}

type ListTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teachers []*Teacher `protobuf:"bytes,1,rep,name=teachers,proto3" json:"teachers,omitempty"`
}

func (x *ListTeacherResponse) Reset() {
	*x = ListTeacherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_teacher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeacherResponse) ProtoMessage() {}

func (x *ListTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ss_teacher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeacherResponse.ProtoReflect.Descriptor instead.
func (*ListTeacherResponse) Descriptor() ([]byte, []int) {
	return file_ss_teacher_proto_rawDescGZIP(), []int{4}
}

func (x *ListTeacherResponse) GetTeachers() []*Teacher {
	if x != nil {
		return x.Teachers
	}
	return nil
}

var File_ss_teacher_proto protoreflect.FileDescriptor

var file_ss_teacher_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x73, 0x5f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a,
	0x07, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x7f, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x35, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x43, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2a, 0x39, 0x0a,
	0x0c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x0c, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x41, 0x44, 0x55, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x41, 0x53, 0x53, 0x49, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x45, 0x41, 0x4e, 0x10, 0x02, 0x32, 0xe4, 0x01, 0x0a, 0x0e, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x0c, 0x50, 0x61, 0x74, 0x63, 0x68, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61,
	0x76, 0x65, 0x6c, 0x7a, 0x68, 0x69, 0x72, 0x6e, 0x6f, 0x76, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69,
	0x3b, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ss_teacher_proto_rawDescOnce sync.Once
	file_ss_teacher_proto_rawDescData = file_ss_teacher_proto_rawDesc
)

func file_ss_teacher_proto_rawDescGZIP() []byte {
	file_ss_teacher_proto_rawDescOnce.Do(func() {
		file_ss_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_ss_teacher_proto_rawDescData)
	})
	return file_ss_teacher_proto_rawDescData
}

var file_ss_teacher_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ss_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ss_teacher_proto_goTypes = []interface{}{
	(PositionType)(0),            // 0: student.PositionType
	(*CreateTeacherRequest)(nil), // 1: student.CreateTeacherRequest
	(*Teacher)(nil),              // 2: student.Teacher
	(*UpdateTeacherRequest)(nil), // 3: student.UpdateTeacherRequest
	(*ListTeacherRequest)(nil),   // 4: student.ListTeacherRequest
	(*ListTeacherResponse)(nil),  // 5: student.ListTeacherResponse
}
var file_ss_teacher_proto_depIdxs = []int32{
	0, // 0: student.CreateTeacherRequest.position_type:type_name -> student.PositionType
	0, // 1: student.Teacher.position_type:type_name -> student.PositionType
	0, // 2: student.UpdateTeacherRequest.position_type:type_name -> student.PositionType
	2, // 3: student.ListTeacherResponse.teachers:type_name -> student.Teacher
	1, // 4: student.TeacherService.CreateTeacher:input_type -> student.CreateTeacherRequest
	3, // 5: student.TeacherService.PatchTeacher:input_type -> student.UpdateTeacherRequest
	4, // 6: student.TeacherService.ListTeachers:input_type -> student.ListTeacherRequest
	2, // 7: student.TeacherService.CreateTeacher:output_type -> student.Teacher
	2, // 8: student.TeacherService.PatchTeacher:output_type -> student.Teacher
	5, // 9: student.TeacherService.ListTeachers:output_type -> student.ListTeacherResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ss_teacher_proto_init() }
func file_ss_teacher_proto_init() {
	if File_ss_teacher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ss_teacher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeacherRequest); i {
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
		file_ss_teacher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teacher); i {
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
		file_ss_teacher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTeacherRequest); i {
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
		file_ss_teacher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeacherRequest); i {
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
		file_ss_teacher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeacherResponse); i {
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
			RawDescriptor: file_ss_teacher_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ss_teacher_proto_goTypes,
		DependencyIndexes: file_ss_teacher_proto_depIdxs,
		EnumInfos:         file_ss_teacher_proto_enumTypes,
		MessageInfos:      file_ss_teacher_proto_msgTypes,
	}.Build()
	File_ss_teacher_proto = out.File
	file_ss_teacher_proto_rawDesc = nil
	file_ss_teacher_proto_goTypes = nil
	file_ss_teacher_proto_depIdxs = nil
}
