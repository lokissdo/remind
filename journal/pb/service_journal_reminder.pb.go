// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: service_journal_reminder.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_journal_reminder_proto protoreflect.FileDescriptor

var file_service_journal_reminder_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x6a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x70,
	0x63, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xb4, 0x0d, 0x0a, 0x0d, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x12, 0xa6, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x92, 0x41,
	0x38, 0x12, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x1a, 0x24, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65,
	0x77, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22,
	0x1a, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0xa2,
	0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x92, 0x41, 0x34, 0x12, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x20, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x32, 0x1a, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x3a, 0x01, 0x2a, 0x12, 0x9f, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x92, 0x41, 0x34, 0x12,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x1a, 0x20, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a, 0x1a, 0x2f, 0x6a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0xa0, 0x01, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x92,
	0x41, 0x46, 0x12, 0x19, 0x41, 0x64, 0x64, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x29, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x61,
	0x64, 0x64, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x61,
	0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15,
	0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xb3, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x73, 0x92, 0x41, 0x50, 0x12, 0x1e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20,
	0x66, 0x72, 0x6f, 0x6d, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x2e,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20,
	0x66, 0x72, 0x6f, 0x6d, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0xa0,
	0x01, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x64, 0x64, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x92, 0x41, 0x46, 0x12, 0x19, 0x41, 0x64, 0x64,
	0x20, 0x61, 0x6e, 0x20, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x6a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x29, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x64, 0x64, 0x20, 0x61, 0x6e, 0x20, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x3a, 0x01,
	0x2a, 0x12, 0xb3, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x73, 0x92, 0x41, 0x50, 0x12, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x61, 0x6e, 0x20, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x61, 0x20,
	0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x2e, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x61, 0x6e, 0x20, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x61, 0x20,
	0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f,
	0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x99, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x32,
	0x12, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x1a, 0x1f, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x61, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x12, 0x9b, 0x01, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x92, 0x41, 0x30, 0x12,
	0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x1a,
	0x1e, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x73, 0x12, 0xc6, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x70, 0x92, 0x41, 0x42, 0x12, 0x17, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x1a, 0x27, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x42, 0x97, 0x01, 0x5a, 0x11, 0x72,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62,
	0x92, 0x41, 0x80, 0x01, 0x12, 0x7e, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x20, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x41,
	0x50, 0x49, 0x22, 0x5b, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x20, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x12, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x6b, 0x69, 0x73, 0x73, 0x64,
	0x6f, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x1a, 0x1b, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x68,
	0x61, 0x6e, 0x68, 0x67, 0x71, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_journal_reminder_proto_goTypes = []interface{}{
	(*CreateJournalRequest)(nil),          // 0: pb.CreateJournalRequest
	(*UpdateJournalRequest)(nil),          // 1: pb.UpdateJournalRequest
	(*DeleteJournalRequest)(nil),          // 2: pb.DeleteJournalRequest
	(*AddImageRequest)(nil),               // 3: pb.AddImageRequest
	(*DeleteImageRequest)(nil),            // 4: pb.DeleteImageRequest
	(*AddAudioRequest)(nil),               // 5: pb.AddAudioRequest
	(*DeleteAudioRequest)(nil),            // 6: pb.DeleteAudioRequest
	(*QueryJournalRequest)(nil),           // 7: pb.QueryJournalRequest
	(*QueryJournalsRequest)(nil),          // 8: pb.QueryJournalsRequest
	(*CreateJournalResponse)(nil),         // 9: pb.CreateJournalResponse
	(*UpdateJournalResponse)(nil),         // 10: pb.UpdateJournalResponse
	(*DeleteJournalResponse)(nil),         // 11: pb.DeleteJournalResponse
	(*AddImageResponse)(nil),              // 12: pb.AddImageResponse
	(*DeleteImageResponse)(nil),           // 13: pb.DeleteImageResponse
	(*AddAudioResponse)(nil),              // 14: pb.AddAudioResponse
	(*DeleteAudioResponse)(nil),           // 15: pb.DeleteAudioResponse
	(*QueryJournalResponse)(nil),          // 16: pb.QueryJournalResponse
	(*QueryJournalsResponse)(nil),         // 17: pb.QueryJournalsResponse
	(*AdvancedQueryJournalsResponse)(nil), // 18: pb.AdvancedQueryJournalsResponse
}
var file_service_journal_reminder_proto_depIdxs = []int32{
	0,  // 0: pb.JournalRemind.CreateJournal:input_type -> pb.CreateJournalRequest
	1,  // 1: pb.JournalRemind.UpdateJournal:input_type -> pb.UpdateJournalRequest
	2,  // 2: pb.JournalRemind.DeleteJournal:input_type -> pb.DeleteJournalRequest
	3,  // 3: pb.JournalRemind.AddImage:input_type -> pb.AddImageRequest
	4,  // 4: pb.JournalRemind.DeleteImage:input_type -> pb.DeleteImageRequest
	5,  // 5: pb.JournalRemind.AddAudio:input_type -> pb.AddAudioRequest
	6,  // 6: pb.JournalRemind.DeleteAudio:input_type -> pb.DeleteAudioRequest
	7,  // 7: pb.JournalRemind.QueryJournal:input_type -> pb.QueryJournalRequest
	8,  // 8: pb.JournalRemind.QueryJournals:input_type -> pb.QueryJournalsRequest
	8,  // 9: pb.JournalRemind.AdvancedQueryJournals:input_type -> pb.QueryJournalsRequest
	9,  // 10: pb.JournalRemind.CreateJournal:output_type -> pb.CreateJournalResponse
	10, // 11: pb.JournalRemind.UpdateJournal:output_type -> pb.UpdateJournalResponse
	11, // 12: pb.JournalRemind.DeleteJournal:output_type -> pb.DeleteJournalResponse
	12, // 13: pb.JournalRemind.AddImage:output_type -> pb.AddImageResponse
	13, // 14: pb.JournalRemind.DeleteImage:output_type -> pb.DeleteImageResponse
	14, // 15: pb.JournalRemind.AddAudio:output_type -> pb.AddAudioResponse
	15, // 16: pb.JournalRemind.DeleteAudio:output_type -> pb.DeleteAudioResponse
	16, // 17: pb.JournalRemind.QueryJournal:output_type -> pb.QueryJournalResponse
	17, // 18: pb.JournalRemind.QueryJournals:output_type -> pb.QueryJournalsResponse
	18, // 19: pb.JournalRemind.AdvancedQueryJournals:output_type -> pb.AdvancedQueryJournalsResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_journal_reminder_proto_init() }
func file_service_journal_reminder_proto_init() {
	if File_service_journal_reminder_proto != nil {
		return
	}
	file_rpc_crud_journal_proto_init()
	file_rpc_query_journal_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_journal_reminder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_journal_reminder_proto_goTypes,
		DependencyIndexes: file_service_journal_reminder_proto_depIdxs,
	}.Build()
	File_service_journal_reminder_proto = out.File
	file_service_journal_reminder_proto_rawDesc = nil
	file_service_journal_reminder_proto_goTypes = nil
	file_service_journal_reminder_proto_depIdxs = nil
}
