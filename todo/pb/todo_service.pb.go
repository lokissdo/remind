// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: todo_service.proto

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

var File_todo_service_proto protoreflect.FileDescriptor

var file_todo_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72,
	0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xef, 0x05, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x34, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x20, 0x6e, 0x65, 0x77, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x95, 0x01, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x92, 0x41, 0x36, 0x12, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x27, 0x55, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20,
	0x74, 0x6f, 0x64, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x32, 0x14, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x6f, 0x64, 0x6f, 0x12, 0x8e, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x92, 0x41, 0x32, 0x12, 0x0e, 0x47, 0x65, 0x74, 0x20, 0x74,
	0x6f, 0x64, 0x6f, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x1a, 0x20, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61,
	0x20, 0x74, 0x6f, 0x64, 0x6f, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x12, 0x11, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f,
	0x74, 0x6f, 0x64, 0x6f, 0x12, 0x8f, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64,
	0x6f, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x55, 0x92, 0x41, 0x37, 0x12, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x1a, 0x29, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74,
	0x6f, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x12, 0x13, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x92, 0x41, 0x32, 0x12, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x23, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x74, 0x6f, 0x64, 0x6f, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x2a, 0x14, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x42, 0x87, 0x01, 0x92, 0x41, 0x74, 0x12, 0x72, 0x0a,
	0x17, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x20, 0x54, 0x6f, 0x64, 0x6f, 0x20, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x20, 0x41, 0x50, 0x49, 0x22, 0x52, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x20, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x29, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x6b, 0x69,
	0x73, 0x73, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x1a, 0x18, 0x64, 0x6f, 0x6b, 0x68, 0x61, 0x69, 0x68, 0x75, 0x6e, 0x67, 0x32, 0x30,
	0x30, 0x33, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x5a, 0x0e, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_todo_service_proto_goTypes = []any{
	(*CreateTodoRequest)(nil),   // 0: pb.CreateTodoRequest
	(*UpdateTodoRequest)(nil),   // 1: pb.UpdateTodoRequest
	(*GetTodoByIDRequest)(nil),  // 2: pb.GetTodoByIDRequest
	(*ListTodosRequest)(nil),    // 3: pb.ListTodosRequest
	(*DeleteTodoRequest)(nil),   // 4: pb.DeleteTodoRequest
	(*CreateTodoResponse)(nil),  // 5: pb.CreateTodoResponse
	(*UpdateTodoResponse)(nil),  // 6: pb.UpdateTodoResponse
	(*GetTodoByIDResponse)(nil), // 7: pb.GetTodoByIDResponse
	(*ListTodosResponse)(nil),   // 8: pb.ListTodosResponse
	(*DeleteTodoResponse)(nil),  // 9: pb.DeleteTodoResponse
}
var file_todo_service_proto_depIdxs = []int32{
	0, // 0: pb.TodoService.CreateTodo:input_type -> pb.CreateTodoRequest
	1, // 1: pb.TodoService.UpdateTodo:input_type -> pb.UpdateTodoRequest
	2, // 2: pb.TodoService.GetTodoByID:input_type -> pb.GetTodoByIDRequest
	3, // 3: pb.TodoService.ListTodos:input_type -> pb.ListTodosRequest
	4, // 4: pb.TodoService.DeleteTodo:input_type -> pb.DeleteTodoRequest
	5, // 5: pb.TodoService.CreateTodo:output_type -> pb.CreateTodoResponse
	6, // 6: pb.TodoService.UpdateTodo:output_type -> pb.UpdateTodoResponse
	7, // 7: pb.TodoService.GetTodoByID:output_type -> pb.GetTodoByIDResponse
	8, // 8: pb.TodoService.ListTodos:output_type -> pb.ListTodosResponse
	9, // 9: pb.TodoService.DeleteTodo:output_type -> pb.DeleteTodoResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_todo_service_proto_init() }
func file_todo_service_proto_init() {
	if File_todo_service_proto != nil {
		return
	}
	file_rpc_create_todo_proto_init()
	file_rpc_update_todo_proto_init()
	file_rpc_get_todo_proto_init()
	file_rpc_list_todos_proto_init()
	file_rpc_delete_todo_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_todo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_service_proto_goTypes,
		DependencyIndexes: file_todo_service_proto_depIdxs,
	}.Build()
	File_todo_service_proto = out.File
	file_todo_service_proto_rawDesc = nil
	file_todo_service_proto_goTypes = nil
	file_todo_service_proto_depIdxs = nil
}
