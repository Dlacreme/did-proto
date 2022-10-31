// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: demail.proto

package demail_proto

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

type Severity int32

const (
	Severity_DEBUG    Severity = 0   // For debug purpose
	Severity_LOG      Severity = 100 // Normal
	Severity_WARNING  Severity = 200 // Important and impact one or few users
	Severity_CRITICAL Severity = 300 // High impact on users
	Severity_FATAL    Severity = 400 // Application crashed
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0:   "DEBUG",
		100: "LOG",
		200: "WARNING",
		300: "CRITICAL",
		400: "FATAL",
	}
	Severity_value = map[string]int32{
		"DEBUG":    0,
		"LOG":      100,
		"WARNING":  200,
		"CRITICAL": 300,
		"FATAL":    400,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_demail_proto_enumTypes[0].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_demail_proto_enumTypes[0]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_demail_proto_rawDescGZIP(), []int{0}
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Severity Severity `protobuf:"varint,1,opt,name=severity,proto3,enum=Severity" json:"severity,omitempty"`
	Message  string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_demail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_demail_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_DEBUG
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_demail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_demail_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_demail_proto protoreflect.FileDescriptor

var file_demail_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46,
	0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x25, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x47, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x47, 0x10, 0x64, 0x12, 0x0c, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0xc8, 0x01, 0x12, 0x0d, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43,
	0x41, 0x4c, 0x10, 0xac, 0x02, 0x12, 0x0a, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x90,
	0x03, 0x32, 0x24, 0x0a, 0x06, 0x44, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x07, 0x50,
	0x75, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x12, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x07, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6c, 0x61, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x2f, 0x64,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_demail_proto_rawDescOnce sync.Once
	file_demail_proto_rawDescData = file_demail_proto_rawDesc
)

func file_demail_proto_rawDescGZIP() []byte {
	file_demail_proto_rawDescOnce.Do(func() {
		file_demail_proto_rawDescData = protoimpl.X.CompressGZIP(file_demail_proto_rawDescData)
	})
	return file_demail_proto_rawDescData
}

var file_demail_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_demail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_demail_proto_goTypes = []interface{}{
	(Severity)(0),  // 0: Severity
	(*Log)(nil),    // 1: Log
	(*Status)(nil), // 2: Status
}
var file_demail_proto_depIdxs = []int32{
	0, // 0: Log.severity:type_name -> Severity
	1, // 1: Demail.PushLog:input_type -> Log
	2, // 2: Demail.PushLog:output_type -> Status
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_demail_proto_init() }
func file_demail_proto_init() {
	if File_demail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_demail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_demail_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demail_proto_goTypes,
		DependencyIndexes: file_demail_proto_depIdxs,
		EnumInfos:         file_demail_proto_enumTypes,
		MessageInfos:      file_demail_proto_msgTypes,
	}.Build()
	File_demail_proto = out.File
	file_demail_proto_rawDesc = nil
	file_demail_proto_goTypes = nil
	file_demail_proto_depIdxs = nil
}
