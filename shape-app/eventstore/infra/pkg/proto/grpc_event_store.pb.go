// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: grpc_event_store.proto

package evenstore

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregateId *UUID `protobuf:"bytes,1,opt,name=aggregateId,proto3" json:"aggregateId,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_grpc_event_store_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetAggregateId() *UUID {
	if x != nil {
		return x.AggregateId
	}
	return nil
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_grpc_event_store_proto_rawDescGZIP(), []int{1}
}

func (x *UUID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Events  *Events `protobuf:"bytes,3,opt,name=events,proto3" json:"events,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_event_store_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetEvents() *Events {
	if x != nil {
		return x.Events
	}
	return nil
}

type Events struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event []*Event `protobuf:"bytes,1,rep,name=event,proto3" json:"event,omitempty"`
}

func (x *Events) Reset() {
	*x = Events{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Events) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events) ProtoMessage() {}

func (x *Events) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events.ProtoReflect.Descriptor instead.
func (*Events) Descriptor() ([]byte, []int) {
	return file_grpc_event_store_proto_rawDescGZIP(), []int{3}
}

func (x *Events) GetEvent() []*Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_grpc_event_store_proto protoreflect.FileDescriptor

var file_grpc_event_store_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x22, 0x3a, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x0b,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22,
	0x16, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x30, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x32, 0x6a, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x13, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x0f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b,
	0x5a, 0x19, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x32, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpc_event_store_proto_rawDescOnce sync.Once
	file_grpc_event_store_proto_rawDescData = file_grpc_event_store_proto_rawDesc
)

func file_grpc_event_store_proto_rawDescGZIP() []byte {
	file_grpc_event_store_proto_rawDescOnce.Do(func() {
		file_grpc_event_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_event_store_proto_rawDescData)
	})
	return file_grpc_event_store_proto_rawDescData
}

var file_grpc_event_store_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_event_store_proto_goTypes = []interface{}{
	(*Event)(nil),    // 0: evenstore.Event
	(*UUID)(nil),     // 1: evenstore.UUID
	(*Response)(nil), // 2: evenstore.Response
	(*Events)(nil),   // 3: evenstore.Events
}
var file_grpc_event_store_proto_depIdxs = []int32{
	1, // 0: evenstore.Event.aggregateId:type_name -> evenstore.UUID
	3, // 1: evenstore.Response.events:type_name -> evenstore.Events
	0, // 2: evenstore.Events.event:type_name -> evenstore.Event
	3, // 3: evenstore.EventStore.Save:input_type -> evenstore.Events
	1, // 4: evenstore.EventStore.Load:input_type -> evenstore.UUID
	2, // 5: evenstore.EventStore.Save:output_type -> evenstore.Response
	2, // 6: evenstore.EventStore.Load:output_type -> evenstore.Response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_event_store_proto_init() }
func file_grpc_event_store_proto_init() {
	if File_grpc_event_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_event_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_grpc_event_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_grpc_event_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_grpc_event_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Events); i {
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
			RawDescriptor: file_grpc_event_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_event_store_proto_goTypes,
		DependencyIndexes: file_grpc_event_store_proto_depIdxs,
		MessageInfos:      file_grpc_event_store_proto_msgTypes,
	}.Build()
	File_grpc_event_store_proto = out.File
	file_grpc_event_store_proto_rawDesc = nil
	file_grpc_event_store_proto_goTypes = nil
	file_grpc_event_store_proto_depIdxs = nil
}