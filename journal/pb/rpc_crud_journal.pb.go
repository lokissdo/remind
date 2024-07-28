// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: rpc_crud_journal.proto

package pb

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

type CreateJournalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content  string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status   bool     `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Images   []string `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`
	Audios   []string `protobuf:"bytes,7,rep,name=audios,proto3" json:"audios,omitempty"`
}

func (x *CreateJournalRequest) Reset() {
	*x = CreateJournalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJournalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJournalRequest) ProtoMessage() {}

func (x *CreateJournalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJournalRequest.ProtoReflect.Descriptor instead.
func (*CreateJournalRequest) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJournalRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateJournalRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateJournalRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateJournalRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *CreateJournalRequest) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *CreateJournalRequest) GetAudios() []string {
	if x != nil {
		return x.Audios
	}
	return nil
}

type CreateJournalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Journal *Journal `protobuf:"bytes,1,opt,name=journal,proto3" json:"journal,omitempty"`
	Images  []*Image `protobuf:"bytes,2,rep,name=images,proto3" json:"images,omitempty"`
	Audios  []*Audio `protobuf:"bytes,3,rep,name=audios,proto3" json:"audios,omitempty"`
}

func (x *CreateJournalResponse) Reset() {
	*x = CreateJournalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJournalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJournalResponse) ProtoMessage() {}

func (x *CreateJournalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJournalResponse.ProtoReflect.Descriptor instead.
func (*CreateJournalResponse) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJournalResponse) GetJournal() *Journal {
	if x != nil {
		return x.Journal
	}
	return nil
}

func (x *CreateJournalResponse) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *CreateJournalResponse) GetAudios() []*Audio {
	if x != nil {
		return x.Audios
	}
	return nil
}

type DeleteJournalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *DeleteJournalRequest) Reset() {
	*x = DeleteJournalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJournalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJournalRequest) ProtoMessage() {}

func (x *DeleteJournalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJournalRequest.ProtoReflect.Descriptor instead.
func (*DeleteJournalRequest) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteJournalRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteJournalRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type DeleteJournalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteJournalResponse) Reset() {
	*x = DeleteJournalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJournalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJournalResponse) ProtoMessage() {}

func (x *DeleteJournalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJournalResponse.ProtoReflect.Descriptor instead.
func (*DeleteJournalResponse) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteJournalResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UpdateJournalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string  `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Title    *string `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Content  *string `protobuf:"bytes,4,opt,name=content,proto3,oneof" json:"content,omitempty"`
	Status   *bool   `protobuf:"varint,5,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *UpdateJournalRequest) Reset() {
	*x = UpdateJournalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJournalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJournalRequest) ProtoMessage() {}

func (x *UpdateJournalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJournalRequest.ProtoReflect.Descriptor instead.
func (*UpdateJournalRequest) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateJournalRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateJournalRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateJournalRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdateJournalRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *UpdateJournalRequest) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

type UpdateJournalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Journal *Journal `protobuf:"bytes,1,opt,name=journal,proto3" json:"journal,omitempty"`
}

func (x *UpdateJournalResponse) Reset() {
	*x = UpdateJournalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJournalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJournalResponse) ProtoMessage() {}

func (x *UpdateJournalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJournalResponse.ProtoReflect.Descriptor instead.
func (*UpdateJournalResponse) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateJournalResponse) GetJournal() *Journal {
	if x != nil {
		return x.Journal
	}
	return nil
}

type AddImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JournalId int64  `protobuf:"varint,1,opt,name=journal_id,json=journalId,proto3" json:"journal_id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AddImageRequest) Reset() {
	*x = AddImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddImageRequest) ProtoMessage() {}

func (x *AddImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddImageRequest.ProtoReflect.Descriptor instead.
func (*AddImageRequest) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{6}
}

func (x *AddImageRequest) GetJournalId() int64 {
	if x != nil {
		return x.JournalId
	}
	return 0
}

func (x *AddImageRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddImageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AddImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *AddImageResponse) Reset() {
	*x = AddImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddImageResponse) ProtoMessage() {}

func (x *AddImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddImageResponse.ProtoReflect.Descriptor instead.
func (*AddImageResponse) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{7}
}

func (x *AddImageResponse) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type DeleteImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JournalId int64  `protobuf:"varint,2,opt,name=journal_id,json=journalId,proto3" json:"journal_id,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *DeleteImageRequest) Reset() {
	*x = DeleteImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImageRequest) ProtoMessage() {}

func (x *DeleteImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImageRequest.ProtoReflect.Descriptor instead.
func (*DeleteImageRequest) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteImageRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteImageRequest) GetJournalId() int64 {
	if x != nil {
		return x.JournalId
	}
	return 0
}

func (x *DeleteImageRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type DeleteImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteImageResponse) Reset() {
	*x = DeleteImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImageResponse) ProtoMessage() {}

func (x *DeleteImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImageResponse.ProtoReflect.Descriptor instead.
func (*DeleteImageResponse) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteImageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type AddAudioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JournalId int64  `protobuf:"varint,1,opt,name=journal_id,json=journalId,proto3" json:"journal_id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AddAudioRequest) Reset() {
	*x = AddAudioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAudioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAudioRequest) ProtoMessage() {}

func (x *AddAudioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAudioRequest.ProtoReflect.Descriptor instead.
func (*AddAudioRequest) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{10}
}

func (x *AddAudioRequest) GetJournalId() int64 {
	if x != nil {
		return x.JournalId
	}
	return 0
}

func (x *AddAudioRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddAudioRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AddAudioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audio *Audio `protobuf:"bytes,1,opt,name=audio,proto3" json:"audio,omitempty"`
}

func (x *AddAudioResponse) Reset() {
	*x = AddAudioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAudioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAudioResponse) ProtoMessage() {}

func (x *AddAudioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAudioResponse.ProtoReflect.Descriptor instead.
func (*AddAudioResponse) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{11}
}

func (x *AddAudioResponse) GetAudio() *Audio {
	if x != nil {
		return x.Audio
	}
	return nil
}

type DeleteAudioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JournalId int64  `protobuf:"varint,2,opt,name=journal_id,json=journalId,proto3" json:"journal_id,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *DeleteAudioRequest) Reset() {
	*x = DeleteAudioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAudioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAudioRequest) ProtoMessage() {}

func (x *DeleteAudioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAudioRequest.ProtoReflect.Descriptor instead.
func (*DeleteAudioRequest) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteAudioRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteAudioRequest) GetJournalId() int64 {
	if x != nil {
		return x.JournalId
	}
	return 0
}

func (x *DeleteAudioRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type DeleteAudioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteAudioResponse) Reset() {
	*x = DeleteAudioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_crud_journal_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAudioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAudioResponse) ProtoMessage() {}

func (x *DeleteAudioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_crud_journal_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAudioResponse.ProtoReflect.Descriptor instead.
func (*DeleteAudioResponse) Descriptor() ([]byte, []int) {
	return file_rpc_crud_journal_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteAudioResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_rpc_crud_journal_proto protoreflect.FileDescriptor

var file_rpc_crud_journal_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0d, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x07, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x07,
	0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x06, 0x61, 0x75,
	0x64, 0x69, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x06, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x22, 0x42, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x3e, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x07, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x22, 0x66, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x5f,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x66, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x22, 0x5f, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x13, 0x5a, 0x11, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_crud_journal_proto_rawDescOnce sync.Once
	file_rpc_crud_journal_proto_rawDescData = file_rpc_crud_journal_proto_rawDesc
)

func file_rpc_crud_journal_proto_rawDescGZIP() []byte {
	file_rpc_crud_journal_proto_rawDescOnce.Do(func() {
		file_rpc_crud_journal_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_crud_journal_proto_rawDescData)
	})
	return file_rpc_crud_journal_proto_rawDescData
}

var file_rpc_crud_journal_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_rpc_crud_journal_proto_goTypes = []interface{}{
	(*CreateJournalRequest)(nil),  // 0: pb.CreateJournalRequest
	(*CreateJournalResponse)(nil), // 1: pb.CreateJournalResponse
	(*DeleteJournalRequest)(nil),  // 2: pb.DeleteJournalRequest
	(*DeleteJournalResponse)(nil), // 3: pb.DeleteJournalResponse
	(*UpdateJournalRequest)(nil),  // 4: pb.UpdateJournalRequest
	(*UpdateJournalResponse)(nil), // 5: pb.UpdateJournalResponse
	(*AddImageRequest)(nil),       // 6: pb.AddImageRequest
	(*AddImageResponse)(nil),      // 7: pb.AddImageResponse
	(*DeleteImageRequest)(nil),    // 8: pb.DeleteImageRequest
	(*DeleteImageResponse)(nil),   // 9: pb.DeleteImageResponse
	(*AddAudioRequest)(nil),       // 10: pb.AddAudioRequest
	(*AddAudioResponse)(nil),      // 11: pb.AddAudioResponse
	(*DeleteAudioRequest)(nil),    // 12: pb.DeleteAudioRequest
	(*DeleteAudioResponse)(nil),   // 13: pb.DeleteAudioResponse
	(*Journal)(nil),               // 14: pb.Journal
	(*Image)(nil),                 // 15: pb.Image
	(*Audio)(nil),                 // 16: pb.Audio
}
var file_rpc_crud_journal_proto_depIdxs = []int32{
	14, // 0: pb.CreateJournalResponse.journal:type_name -> pb.Journal
	15, // 1: pb.CreateJournalResponse.images:type_name -> pb.Image
	16, // 2: pb.CreateJournalResponse.audios:type_name -> pb.Audio
	14, // 3: pb.UpdateJournalResponse.journal:type_name -> pb.Journal
	15, // 4: pb.AddImageResponse.image:type_name -> pb.Image
	16, // 5: pb.AddAudioResponse.audio:type_name -> pb.Audio
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_rpc_crud_journal_proto_init() }
func file_rpc_crud_journal_proto_init() {
	if File_rpc_crud_journal_proto != nil {
		return
	}
	file_journal_proto_init()
	file_dtypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_crud_journal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJournalRequest); i {
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
		file_rpc_crud_journal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJournalResponse); i {
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
		file_rpc_crud_journal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJournalRequest); i {
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
		file_rpc_crud_journal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJournalResponse); i {
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
		file_rpc_crud_journal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJournalRequest); i {
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
		file_rpc_crud_journal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJournalResponse); i {
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
		file_rpc_crud_journal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddImageRequest); i {
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
		file_rpc_crud_journal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddImageResponse); i {
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
		file_rpc_crud_journal_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImageRequest); i {
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
		file_rpc_crud_journal_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImageResponse); i {
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
		file_rpc_crud_journal_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAudioRequest); i {
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
		file_rpc_crud_journal_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAudioResponse); i {
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
		file_rpc_crud_journal_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAudioRequest); i {
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
		file_rpc_crud_journal_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAudioResponse); i {
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
	file_rpc_crud_journal_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_crud_journal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_crud_journal_proto_goTypes,
		DependencyIndexes: file_rpc_crud_journal_proto_depIdxs,
		MessageInfos:      file_rpc_crud_journal_proto_msgTypes,
	}.Build()
	File_rpc_crud_journal_proto = out.File
	file_rpc_crud_journal_proto_rawDesc = nil
	file_rpc_crud_journal_proto_goTypes = nil
	file_rpc_crud_journal_proto_depIdxs = nil
}
