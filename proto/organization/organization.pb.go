// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: proto/organization/organization.proto

package organization_lms

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

type NewOrganizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Contact   string `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	SourceTag string `protobuf:"bytes,5,opt,name=sourceTag,proto3" json:"sourceTag,omitempty"`
	LogoLink  string `protobuf:"bytes,6,opt,name=logoLink,proto3" json:"logoLink,omitempty"`
	Password  string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *NewOrganizationRequest) Reset() {
	*x = NewOrganizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_organization_organization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOrganizationRequest) ProtoMessage() {}

func (x *NewOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_organization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOrganizationRequest.ProtoReflect.Descriptor instead.
func (*NewOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_proto_organization_organization_proto_rawDescGZIP(), []int{0}
}

func (x *NewOrganizationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewOrganizationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewOrganizationRequest) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *NewOrganizationRequest) GetSourceTag() string {
	if x != nil {
		return x.SourceTag
	}
	return ""
}

func (x *NewOrganizationRequest) GetLogoLink() string {
	if x != nil {
		return x.LogoLink
	}
	return ""
}

func (x *NewOrganizationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type OrganizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status  bool   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrganizationResponse) Reset() {
	*x = OrganizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_organization_organization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationResponse) ProtoMessage() {}

func (x *OrganizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_organization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationResponse.ProtoReflect.Descriptor instead.
func (*OrganizationResponse) Descriptor() ([]byte, []int) {
	return file_proto_organization_organization_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrganizationResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type OrganizationByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId string `protobuf:"bytes,1,opt,name=organizationId,proto3" json:"organizationId,omitempty"`
}

func (x *OrganizationByIdRequest) Reset() {
	*x = OrganizationByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_organization_organization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationByIdRequest) ProtoMessage() {}

func (x *OrganizationByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_organization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationByIdRequest.ProtoReflect.Descriptor instead.
func (*OrganizationByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_organization_organization_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationByIdRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

type OrganizationDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email          string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Contact        string   `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	LogoLink       string   `protobuf:"bytes,5,opt,name=logoLink,proto3" json:"logoLink,omitempty"`
	SourceTag      string   `protobuf:"bytes,6,opt,name=sourceTag,proto3" json:"sourceTag,omitempty"`
	AllowedSources []string `protobuf:"bytes,7,rep,name=allowedSources,proto3" json:"allowedSources,omitempty"`
}

func (x *OrganizationDetails) Reset() {
	*x = OrganizationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_organization_organization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationDetails) ProtoMessage() {}

func (x *OrganizationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_organization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationDetails.ProtoReflect.Descriptor instead.
func (*OrganizationDetails) Descriptor() ([]byte, []int) {
	return file_proto_organization_organization_proto_rawDescGZIP(), []int{3}
}

func (x *OrganizationDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrganizationDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OrganizationDetails) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *OrganizationDetails) GetLogoLink() string {
	if x != nil {
		return x.LogoLink
	}
	return ""
}

func (x *OrganizationDetails) GetSourceTag() string {
	if x != nil {
		return x.SourceTag
	}
	return ""
}

func (x *OrganizationDetails) GetAllowedSources() []string {
	if x != nil {
		return x.AllowedSources
	}
	return nil
}

type UpdateOrganizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Contact        string   `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
	LogoLink       string   `protobuf:"bytes,4,opt,name=logoLink,proto3" json:"logoLink,omitempty"`
	AllowedSources []string `protobuf:"bytes,5,rep,name=allowedSources,proto3" json:"allowedSources,omitempty"`
}

func (x *UpdateOrganizationRequest) Reset() {
	*x = UpdateOrganizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_organization_organization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrganizationRequest) ProtoMessage() {}

func (x *UpdateOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_organization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrganizationRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_proto_organization_organization_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOrganizationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateOrganizationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateOrganizationRequest) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *UpdateOrganizationRequest) GetLogoLink() string {
	if x != nil {
		return x.LogoLink
	}
	return ""
}

func (x *UpdateOrganizationRequest) GetAllowedSources() []string {
	if x != nil {
		return x.AllowedSources
	}
	return nil
}

var File_proto_organization_organization_proto protoreflect.FileDescriptor

var file_proto_organization_organization_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0xb2,
	0x01, 0x0a, 0x16, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x48, 0x0a, 0x14, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x41, 0x0a,
	0x17, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0xcb, 0x01, 0x0a, 0x13, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x61, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x9d,
	0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x32, 0xca,
	0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x66, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00,
	0x12, 0x69, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x3b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_organization_organization_proto_rawDescOnce sync.Once
	file_proto_organization_organization_proto_rawDescData = file_proto_organization_organization_proto_rawDesc
)

func file_proto_organization_organization_proto_rawDescGZIP() []byte {
	file_proto_organization_organization_proto_rawDescOnce.Do(func() {
		file_proto_organization_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_organization_organization_proto_rawDescData)
	})
	return file_proto_organization_organization_proto_rawDescData
}

var file_proto_organization_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_organization_organization_proto_goTypes = []interface{}{
	(*NewOrganizationRequest)(nil),    // 0: go.micro.service.user.NewOrganizationRequest
	(*OrganizationResponse)(nil),      // 1: go.micro.service.user.OrganizationResponse
	(*OrganizationByIdRequest)(nil),   // 2: go.micro.service.user.OrganizationByIdRequest
	(*OrganizationDetails)(nil),       // 3: go.micro.service.user.OrganizationDetails
	(*UpdateOrganizationRequest)(nil), // 4: go.micro.service.user.UpdateOrganizationRequest
}
var file_proto_organization_organization_proto_depIdxs = []int32{
	0, // 0: go.micro.service.user.Organization.Create:input_type -> go.micro.service.user.NewOrganizationRequest
	2, // 1: go.micro.service.user.Organization.GetByID:input_type -> go.micro.service.user.OrganizationByIdRequest
	4, // 2: go.micro.service.user.Organization.Update:input_type -> go.micro.service.user.UpdateOrganizationRequest
	1, // 3: go.micro.service.user.Organization.Create:output_type -> go.micro.service.user.OrganizationResponse
	3, // 4: go.micro.service.user.Organization.GetByID:output_type -> go.micro.service.user.OrganizationDetails
	1, // 5: go.micro.service.user.Organization.Update:output_type -> go.micro.service.user.OrganizationResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_organization_organization_proto_init() }
func file_proto_organization_organization_proto_init() {
	if File_proto_organization_organization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_organization_organization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOrganizationRequest); i {
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
		file_proto_organization_organization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationResponse); i {
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
		file_proto_organization_organization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationByIdRequest); i {
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
		file_proto_organization_organization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationDetails); i {
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
		file_proto_organization_organization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrganizationRequest); i {
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
			RawDescriptor: file_proto_organization_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_organization_organization_proto_goTypes,
		DependencyIndexes: file_proto_organization_organization_proto_depIdxs,
		MessageInfos:      file_proto_organization_organization_proto_msgTypes,
	}.Build()
	File_proto_organization_organization_proto = out.File
	file_proto_organization_organization_proto_rawDesc = nil
	file_proto_organization_organization_proto_goTypes = nil
	file_proto_organization_organization_proto_depIdxs = nil
}
