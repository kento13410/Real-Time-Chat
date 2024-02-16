// Code generated by entproto. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: entpb/entpb.proto

package entpb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetChatRequest_View int32

const (
	GetChatRequest_VIEW_UNSPECIFIED GetChatRequest_View = 0
	GetChatRequest_BASIC            GetChatRequest_View = 1
	GetChatRequest_WITH_EDGE_IDS    GetChatRequest_View = 2
)

// Enum value maps for GetChatRequest_View.
var (
	GetChatRequest_View_name = map[int32]string{
		0: "VIEW_UNSPECIFIED",
		1: "BASIC",
		2: "WITH_EDGE_IDS",
	}
	GetChatRequest_View_value = map[string]int32{
		"VIEW_UNSPECIFIED": 0,
		"BASIC":            1,
		"WITH_EDGE_IDS":    2,
	}
)

func (x GetChatRequest_View) Enum() *GetChatRequest_View {
	p := new(GetChatRequest_View)
	*p = x
	return p
}

func (x GetChatRequest_View) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetChatRequest_View) Descriptor() protoreflect.EnumDescriptor {
	return file_entpb_entpb_proto_enumTypes[0].Descriptor()
}

func (GetChatRequest_View) Type() protoreflect.EnumType {
	return &file_entpb_entpb_proto_enumTypes[0]
}

func (x GetChatRequest_View) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetChatRequest_View.Descriptor instead.
func (GetChatRequest_View) EnumDescriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{2, 0}
}

type ListChatRequest_View int32

const (
	ListChatRequest_VIEW_UNSPECIFIED ListChatRequest_View = 0
	ListChatRequest_BASIC            ListChatRequest_View = 1
	ListChatRequest_WITH_EDGE_IDS    ListChatRequest_View = 2
)

// Enum value maps for ListChatRequest_View.
var (
	ListChatRequest_View_name = map[int32]string{
		0: "VIEW_UNSPECIFIED",
		1: "BASIC",
		2: "WITH_EDGE_IDS",
	}
	ListChatRequest_View_value = map[string]int32{
		"VIEW_UNSPECIFIED": 0,
		"BASIC":            1,
		"WITH_EDGE_IDS":    2,
	}
)

func (x ListChatRequest_View) Enum() *ListChatRequest_View {
	p := new(ListChatRequest_View)
	*p = x
	return p
}

func (x ListChatRequest_View) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListChatRequest_View) Descriptor() protoreflect.EnumDescriptor {
	return file_entpb_entpb_proto_enumTypes[1].Descriptor()
}

func (ListChatRequest_View) Type() protoreflect.EnumType {
	return &file_entpb_entpb_proto_enumTypes[1]
}

func (x ListChatRequest_View) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListChatRequest_View.Descriptor instead.
func (ListChatRequest_View) EnumDescriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{5, 0}
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderId   int64                `protobuf:"varint,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	ReceiverId int64                `protobuf:"varint,3,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	Message    string               `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	SentAt     *timestamp.Timestamp `protobuf:"bytes,5,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Chat) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *Chat) GetReceiverId() int64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *Chat) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Chat) GetSentAt() *timestamp.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

type CreateChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *CreateChatRequest) Reset() {
	*x = CreateChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest) ProtoMessage() {}

func (x *CreateChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRequest.ProtoReflect.Descriptor instead.
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChatRequest) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type GetChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	View GetChatRequest_View `protobuf:"varint,2,opt,name=view,proto3,enum=entpb.GetChatRequest_View" json:"view,omitempty"`
}

func (x *GetChatRequest) Reset() {
	*x = GetChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatRequest) ProtoMessage() {}

func (x *GetChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatRequest.ProtoReflect.Descriptor instead.
func (*GetChatRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{2}
}

func (x *GetChatRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetChatRequest) GetView() GetChatRequest_View {
	if x != nil {
		return x.View
	}
	return GetChatRequest_VIEW_UNSPECIFIED
}

type UpdateChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *UpdateChatRequest) Reset() {
	*x = UpdateChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChatRequest) ProtoMessage() {}

func (x *UpdateChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChatRequest.ProtoReflect.Descriptor instead.
func (*UpdateChatRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateChatRequest) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type DeleteChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteChatRequest) Reset() {
	*x = DeleteChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatRequest) ProtoMessage() {}

func (x *DeleteChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatRequest.ProtoReflect.Descriptor instead.
func (*DeleteChatRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteChatRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32                `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string               `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	View      ListChatRequest_View `protobuf:"varint,3,opt,name=view,proto3,enum=entpb.ListChatRequest_View" json:"view,omitempty"`
}

func (x *ListChatRequest) Reset() {
	*x = ListChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatRequest) ProtoMessage() {}

func (x *ListChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatRequest.ProtoReflect.Descriptor instead.
func (*ListChatRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{5}
}

func (x *ListChatRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListChatRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListChatRequest) GetView() ListChatRequest_View {
	if x != nil {
		return x.View
	}
	return ListChatRequest_VIEW_UNSPECIFIED
}

type ListChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatList      []*Chat `protobuf:"bytes,1,rep,name=chat_list,json=chatList,proto3" json:"chat_list,omitempty"`
	NextPageToken string  `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListChatResponse) Reset() {
	*x = ListChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatResponse) ProtoMessage() {}

func (x *ListChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatResponse.ProtoReflect.Descriptor instead.
func (*ListChatResponse) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{6}
}

func (x *ListChatResponse) GetChatList() []*Chat {
	if x != nil {
		return x.ChatList
	}
	return nil
}

func (x *ListChatResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type BatchCreateChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*CreateChatRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *BatchCreateChatsRequest) Reset() {
	*x = BatchCreateChatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateChatsRequest) ProtoMessage() {}

func (x *BatchCreateChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateChatsRequest.ProtoReflect.Descriptor instead.
func (*BatchCreateChatsRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{7}
}

func (x *BatchCreateChatsRequest) GetRequests() []*CreateChatRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

type BatchCreateChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*Chat `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *BatchCreateChatsResponse) Reset() {
	*x = BatchCreateChatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateChatsResponse) ProtoMessage() {}

func (x *BatchCreateChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateChatsResponse.ProtoReflect.Descriptor instead.
func (*BatchCreateChatsResponse) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{8}
}

func (x *BatchCreateChatsResponse) GetChats() []*Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

var File_entpb_entpb_proto protoreflect.FileDescriptor

var file_entpb_entpb_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0x34,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04,
	0x63, 0x68, 0x61, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x69, 0x65,
	0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x3a, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12,
	0x14, 0x0a, 0x10, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x45, 0x44, 0x47, 0x45, 0x5f, 0x49, 0x44,
	0x53, 0x10, 0x02, 0x22, 0x34, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xba,
	0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2f,
	0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22,
	0x3a, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x45, 0x57, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x49, 0x54, 0x48,
	0x5f, 0x45, 0x44, 0x47, 0x45, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x02, 0x22, 0x64, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x08, 0x63, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x4f, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x22, 0x3d, 0x0a, 0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x32, 0xdf, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x29, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x2f, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x3a,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x16, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x52, 0x65, 0x61, 0x6c, 0x2d, 0x54, 0x69, 0x6d, 0x65,
	0x2d, 0x43, 0x68, 0x61, 0x74, 0x2f, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entpb_entpb_proto_rawDescOnce sync.Once
	file_entpb_entpb_proto_rawDescData = file_entpb_entpb_proto_rawDesc
)

func file_entpb_entpb_proto_rawDescGZIP() []byte {
	file_entpb_entpb_proto_rawDescOnce.Do(func() {
		file_entpb_entpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_entpb_entpb_proto_rawDescData)
	})
	return file_entpb_entpb_proto_rawDescData
}

var file_entpb_entpb_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_entpb_entpb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_entpb_entpb_proto_goTypes = []interface{}{
	(GetChatRequest_View)(0),         // 0: entpb.GetChatRequest.View
	(ListChatRequest_View)(0),        // 1: entpb.ListChatRequest.View
	(*Chat)(nil),                     // 2: entpb.Chat
	(*CreateChatRequest)(nil),        // 3: entpb.CreateChatRequest
	(*GetChatRequest)(nil),           // 4: entpb.GetChatRequest
	(*UpdateChatRequest)(nil),        // 5: entpb.UpdateChatRequest
	(*DeleteChatRequest)(nil),        // 6: entpb.DeleteChatRequest
	(*ListChatRequest)(nil),          // 7: entpb.ListChatRequest
	(*ListChatResponse)(nil),         // 8: entpb.ListChatResponse
	(*BatchCreateChatsRequest)(nil),  // 9: entpb.BatchCreateChatsRequest
	(*BatchCreateChatsResponse)(nil), // 10: entpb.BatchCreateChatsResponse
	(*timestamp.Timestamp)(nil),      // 11: google.protobuf.Timestamp
	(*empty.Empty)(nil),              // 12: google.protobuf.Empty
}
var file_entpb_entpb_proto_depIdxs = []int32{
	11, // 0: entpb.Chat.sent_at:type_name -> google.protobuf.Timestamp
	2,  // 1: entpb.CreateChatRequest.chat:type_name -> entpb.Chat
	0,  // 2: entpb.GetChatRequest.view:type_name -> entpb.GetChatRequest.View
	2,  // 3: entpb.UpdateChatRequest.chat:type_name -> entpb.Chat
	1,  // 4: entpb.ListChatRequest.view:type_name -> entpb.ListChatRequest.View
	2,  // 5: entpb.ListChatResponse.chat_list:type_name -> entpb.Chat
	3,  // 6: entpb.BatchCreateChatsRequest.requests:type_name -> entpb.CreateChatRequest
	2,  // 7: entpb.BatchCreateChatsResponse.chats:type_name -> entpb.Chat
	3,  // 8: entpb.ChatService.Create:input_type -> entpb.CreateChatRequest
	4,  // 9: entpb.ChatService.Get:input_type -> entpb.GetChatRequest
	5,  // 10: entpb.ChatService.Update:input_type -> entpb.UpdateChatRequest
	6,  // 11: entpb.ChatService.Delete:input_type -> entpb.DeleteChatRequest
	7,  // 12: entpb.ChatService.List:input_type -> entpb.ListChatRequest
	9,  // 13: entpb.ChatService.BatchCreate:input_type -> entpb.BatchCreateChatsRequest
	2,  // 14: entpb.ChatService.Create:output_type -> entpb.Chat
	2,  // 15: entpb.ChatService.Get:output_type -> entpb.Chat
	2,  // 16: entpb.ChatService.Update:output_type -> entpb.Chat
	12, // 17: entpb.ChatService.Delete:output_type -> google.protobuf.Empty
	8,  // 18: entpb.ChatService.List:output_type -> entpb.ListChatResponse
	10, // 19: entpb.ChatService.BatchCreate:output_type -> entpb.BatchCreateChatsResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_entpb_entpb_proto_init() }
func file_entpb_entpb_proto_init() {
	if File_entpb_entpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entpb_entpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_entpb_entpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatRequest); i {
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
		file_entpb_entpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatRequest); i {
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
		file_entpb_entpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChatRequest); i {
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
		file_entpb_entpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatRequest); i {
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
		file_entpb_entpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatRequest); i {
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
		file_entpb_entpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatResponse); i {
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
		file_entpb_entpb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchCreateChatsRequest); i {
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
		file_entpb_entpb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchCreateChatsResponse); i {
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
			RawDescriptor: file_entpb_entpb_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entpb_entpb_proto_goTypes,
		DependencyIndexes: file_entpb_entpb_proto_depIdxs,
		EnumInfos:         file_entpb_entpb_proto_enumTypes,
		MessageInfos:      file_entpb_entpb_proto_msgTypes,
	}.Build()
	File_entpb_entpb_proto = out.File
	file_entpb_entpb_proto_rawDesc = nil
	file_entpb_entpb_proto_goTypes = nil
	file_entpb_entpb_proto_depIdxs = nil
}