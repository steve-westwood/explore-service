// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: explore-service.proto

package explore_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListLikedYouRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RecipientUserId string                 `protobuf:"bytes,1,opt,name=recipient_user_id,json=recipientUserId,proto3" json:"recipient_user_id,omitempty"`
	PaginationToken *string                `protobuf:"bytes,2,opt,name=pagination_token,json=paginationToken,proto3,oneof" json:"pagination_token,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListLikedYouRequest) Reset() {
	*x = ListLikedYouRequest{}
	mi := &file_explore_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLikedYouRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLikedYouRequest) ProtoMessage() {}

func (x *ListLikedYouRequest) ProtoReflect() protoreflect.Message {
	mi := &file_explore_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLikedYouRequest.ProtoReflect.Descriptor instead.
func (*ListLikedYouRequest) Descriptor() ([]byte, []int) {
	return file_explore_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListLikedYouRequest) GetRecipientUserId() string {
	if x != nil {
		return x.RecipientUserId
	}
	return ""
}

func (x *ListLikedYouRequest) GetPaginationToken() string {
	if x != nil && x.PaginationToken != nil {
		return *x.PaginationToken
	}
	return ""
}

type ListLikedYouResponse struct {
	state               protoimpl.MessageState        `protogen:"open.v1"`
	Likers              []*ListLikedYouResponse_Liker `protobuf:"bytes,1,rep,name=likers,proto3" json:"likers,omitempty"`
	NextPaginationToken *string                       `protobuf:"bytes,2,opt,name=next_pagination_token,json=nextPaginationToken,proto3,oneof" json:"next_pagination_token,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ListLikedYouResponse) Reset() {
	*x = ListLikedYouResponse{}
	mi := &file_explore_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLikedYouResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLikedYouResponse) ProtoMessage() {}

func (x *ListLikedYouResponse) ProtoReflect() protoreflect.Message {
	mi := &file_explore_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLikedYouResponse.ProtoReflect.Descriptor instead.
func (*ListLikedYouResponse) Descriptor() ([]byte, []int) {
	return file_explore_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListLikedYouResponse) GetLikers() []*ListLikedYouResponse_Liker {
	if x != nil {
		return x.Likers
	}
	return nil
}

func (x *ListLikedYouResponse) GetNextPaginationToken() string {
	if x != nil && x.NextPaginationToken != nil {
		return *x.NextPaginationToken
	}
	return ""
}

type CountLikedYouRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RecipientUserId string                 `protobuf:"bytes,1,opt,name=recipient_user_id,json=recipientUserId,proto3" json:"recipient_user_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CountLikedYouRequest) Reset() {
	*x = CountLikedYouRequest{}
	mi := &file_explore_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountLikedYouRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountLikedYouRequest) ProtoMessage() {}

func (x *CountLikedYouRequest) ProtoReflect() protoreflect.Message {
	mi := &file_explore_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountLikedYouRequest.ProtoReflect.Descriptor instead.
func (*CountLikedYouRequest) Descriptor() ([]byte, []int) {
	return file_explore_service_proto_rawDescGZIP(), []int{2}
}

func (x *CountLikedYouRequest) GetRecipientUserId() string {
	if x != nil {
		return x.RecipientUserId
	}
	return ""
}

type CountLikedYouResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         uint64                 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CountLikedYouResponse) Reset() {
	*x = CountLikedYouResponse{}
	mi := &file_explore_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountLikedYouResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountLikedYouResponse) ProtoMessage() {}

func (x *CountLikedYouResponse) ProtoReflect() protoreflect.Message {
	mi := &file_explore_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountLikedYouResponse.ProtoReflect.Descriptor instead.
func (*CountLikedYouResponse) Descriptor() ([]byte, []int) {
	return file_explore_service_proto_rawDescGZIP(), []int{3}
}

func (x *CountLikedYouResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type PutDecisionRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ActorUserId     string                 `protobuf:"bytes,1,opt,name=actor_user_id,json=actorUserId,proto3" json:"actor_user_id,omitempty"`
	RecipientUserId string                 `protobuf:"bytes,2,opt,name=recipient_user_id,json=recipientUserId,proto3" json:"recipient_user_id,omitempty"`
	LikedRecipient  bool                   `protobuf:"varint,3,opt,name=liked_recipient,json=likedRecipient,proto3" json:"liked_recipient,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *PutDecisionRequest) Reset() {
	*x = PutDecisionRequest{}
	mi := &file_explore_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutDecisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutDecisionRequest) ProtoMessage() {}

func (x *PutDecisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_explore_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutDecisionRequest.ProtoReflect.Descriptor instead.
func (*PutDecisionRequest) Descriptor() ([]byte, []int) {
	return file_explore_service_proto_rawDescGZIP(), []int{4}
}

func (x *PutDecisionRequest) GetActorUserId() string {
	if x != nil {
		return x.ActorUserId
	}
	return ""
}

func (x *PutDecisionRequest) GetRecipientUserId() string {
	if x != nil {
		return x.RecipientUserId
	}
	return ""
}

func (x *PutDecisionRequest) GetLikedRecipient() bool {
	if x != nil {
		return x.LikedRecipient
	}
	return false
}

type PutDecisionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MutualLikes   bool                   `protobuf:"varint,1,opt,name=mutual_likes,json=mutualLikes,proto3" json:"mutual_likes,omitempty"` // True if both users like each other
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutDecisionResponse) Reset() {
	*x = PutDecisionResponse{}
	mi := &file_explore_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutDecisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutDecisionResponse) ProtoMessage() {}

func (x *PutDecisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_explore_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutDecisionResponse.ProtoReflect.Descriptor instead.
func (*PutDecisionResponse) Descriptor() ([]byte, []int) {
	return file_explore_service_proto_rawDescGZIP(), []int{5}
}

func (x *PutDecisionResponse) GetMutualLikes() bool {
	if x != nil {
		return x.MutualLikes
	}
	return false
}

type ListLikedYouResponse_Liker struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ActorId       string                 `protobuf:"bytes,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	UnixTimestamp uint64                 `protobuf:"varint,2,opt,name=unix_timestamp,json=unixTimestamp,proto3" json:"unix_timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLikedYouResponse_Liker) Reset() {
	*x = ListLikedYouResponse_Liker{}
	mi := &file_explore_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLikedYouResponse_Liker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLikedYouResponse_Liker) ProtoMessage() {}

func (x *ListLikedYouResponse_Liker) ProtoReflect() protoreflect.Message {
	mi := &file_explore_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLikedYouResponse_Liker.ProtoReflect.Descriptor instead.
func (*ListLikedYouResponse_Liker) Descriptor() ([]byte, []int) {
	return file_explore_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListLikedYouResponse_Liker) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

func (x *ListLikedYouResponse_Liker) GetUnixTimestamp() uint64 {
	if x != nil {
		return x.UnixTimestamp
	}
	return 0
}

var File_explore_service_proto protoreflect.FileDescriptor

var file_explore_service_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65,
	0x22, 0x86, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f,
	0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf1, 0x01, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x6c, 0x69, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x69, 0x6b, 0x65, 0x72, 0x73, 0x12,
	0x37, 0x0a, 0x15, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x13, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x1a, 0x49, 0x0a, 0x05, 0x4c, 0x69, 0x6b, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x75, 0x6e, 0x69, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a,
	0x14, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x2d, 0x0a, 0x15, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59,
	0x6f, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x8d, 0x01, 0x0a, 0x12, 0x50, 0x75, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x69, 0x6b, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0x38, 0x0a, 0x13, 0x50, 0x75, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x75, 0x74, 0x75, 0x61,
	0x6c, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6d,
	0x75, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x32, 0xc7, 0x02, 0x0a, 0x0e, 0x45,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x12, 0x1c, 0x2e,
	0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x59, 0x6f, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59,
	0x6f, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x65, 0x77, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x12, 0x1c, 0x2e,
	0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x59, 0x6f, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59,
	0x6f, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x12, 0x1d, 0x2e, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64,
	0x59, 0x6f, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59,
	0x6f, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x50, 0x75,
	0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x70, 0x6c,
	0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65,
	0x2e, 0x50, 0x75, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x2d, 0x77, 0x65, 0x73, 0x74, 0x77, 0x6f, 0x6f,
	0x64, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_explore_service_proto_rawDescOnce sync.Once
	file_explore_service_proto_rawDescData []byte
)

func file_explore_service_proto_rawDescGZIP() []byte {
	file_explore_service_proto_rawDescOnce.Do(func() {
		file_explore_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_explore_service_proto_rawDesc), len(file_explore_service_proto_rawDesc)))
	})
	return file_explore_service_proto_rawDescData
}

var file_explore_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_explore_service_proto_goTypes = []any{
	(*ListLikedYouRequest)(nil),        // 0: explore.ListLikedYouRequest
	(*ListLikedYouResponse)(nil),       // 1: explore.ListLikedYouResponse
	(*CountLikedYouRequest)(nil),       // 2: explore.CountLikedYouRequest
	(*CountLikedYouResponse)(nil),      // 3: explore.CountLikedYouResponse
	(*PutDecisionRequest)(nil),         // 4: explore.PutDecisionRequest
	(*PutDecisionResponse)(nil),        // 5: explore.PutDecisionResponse
	(*ListLikedYouResponse_Liker)(nil), // 6: explore.ListLikedYouResponse.Liker
}
var file_explore_service_proto_depIdxs = []int32{
	6, // 0: explore.ListLikedYouResponse.likers:type_name -> explore.ListLikedYouResponse.Liker
	0, // 1: explore.ExploreService.ListLikedYou:input_type -> explore.ListLikedYouRequest
	0, // 2: explore.ExploreService.ListNewLikedYou:input_type -> explore.ListLikedYouRequest
	2, // 3: explore.ExploreService.CountLikedYou:input_type -> explore.CountLikedYouRequest
	4, // 4: explore.ExploreService.PutDecision:input_type -> explore.PutDecisionRequest
	1, // 5: explore.ExploreService.ListLikedYou:output_type -> explore.ListLikedYouResponse
	1, // 6: explore.ExploreService.ListNewLikedYou:output_type -> explore.ListLikedYouResponse
	3, // 7: explore.ExploreService.CountLikedYou:output_type -> explore.CountLikedYouResponse
	5, // 8: explore.ExploreService.PutDecision:output_type -> explore.PutDecisionResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_explore_service_proto_init() }
func file_explore_service_proto_init() {
	if File_explore_service_proto != nil {
		return
	}
	file_explore_service_proto_msgTypes[0].OneofWrappers = []any{}
	file_explore_service_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_explore_service_proto_rawDesc), len(file_explore_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_explore_service_proto_goTypes,
		DependencyIndexes: file_explore_service_proto_depIdxs,
		MessageInfos:      file_explore_service_proto_msgTypes,
	}.Build()
	File_explore_service_proto = out.File
	file_explore_service_proto_goTypes = nil
	file_explore_service_proto_depIdxs = nil
}
