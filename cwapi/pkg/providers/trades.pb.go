// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/fetcher/v1/trades.proto

package providers

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

type FetchTradesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *FetchTradesRequest) Reset() {
	*x = FetchTradesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_v1_trades_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTradesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTradesRequest) ProtoMessage() {}

func (x *FetchTradesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_v1_trades_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTradesRequest.ProtoReflect.Descriptor instead.
func (*FetchTradesRequest) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_v1_trades_proto_rawDescGZIP(), []int{0}
}

func (x *FetchTradesRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_v1_trades_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_v1_trades_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_v1_trades_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FetchTradesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *FetchTradesResponse) Reset() {
	*x = FetchTradesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_v1_trades_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTradesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTradesResponse) ProtoMessage() {}

func (x *FetchTradesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_v1_trades_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTradesResponse.ProtoReflect.Descriptor instead.
func (*FetchTradesResponse) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_v1_trades_proto_rawDescGZIP(), []int{2}
}

func (x *FetchTradesResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type StopFetchTradesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StopFetchTradesResponse) Reset() {
	*x = StopFetchTradesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_v1_trades_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopFetchTradesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopFetchTradesResponse) ProtoMessage() {}

func (x *StopFetchTradesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_v1_trades_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopFetchTradesResponse.ProtoReflect.Descriptor instead.
func (*StopFetchTradesResponse) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_v1_trades_proto_rawDescGZIP(), []int{3}
}

func (x *StopFetchTradesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_v1_trades_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_v1_trades_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_v1_trades_proto_rawDescGZIP(), []int{4}
}

var File_proto_fetcher_v1_trades_proto protoreflect.FileDescriptor

var file_proto_fetcher_v1_trades_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x2c, 0x0a, 0x12, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x16, 0x0a, 0x04, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3d, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x22, 0x33, 0x0a, 0x17, 0x53, 0x74, 0x6f, 0x70, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xaf,
	0x01, 0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x0b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x12, 0x1e, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x70, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fetcher_v1_trades_proto_rawDescOnce sync.Once
	file_proto_fetcher_v1_trades_proto_rawDescData = file_proto_fetcher_v1_trades_proto_rawDesc
)

func file_proto_fetcher_v1_trades_proto_rawDescGZIP() []byte {
	file_proto_fetcher_v1_trades_proto_rawDescOnce.Do(func() {
		file_proto_fetcher_v1_trades_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fetcher_v1_trades_proto_rawDescData)
	})
	return file_proto_fetcher_v1_trades_proto_rawDescData
}

var file_proto_fetcher_v1_trades_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_fetcher_v1_trades_proto_goTypes = []interface{}{
	(*FetchTradesRequest)(nil),      // 0: fetcher.v1.FetchTradesRequest
	(*Task)(nil),                    // 1: fetcher.v1.Task
	(*FetchTradesResponse)(nil),     // 2: fetcher.v1.FetchTradesResponse
	(*StopFetchTradesResponse)(nil), // 3: fetcher.v1.StopFetchTradesResponse
	(*Empty)(nil),                   // 4: fetcher.v1.Empty
}
var file_proto_fetcher_v1_trades_proto_depIdxs = []int32{
	1, // 0: fetcher.v1.FetchTradesResponse.tasks:type_name -> fetcher.v1.Task
	0, // 1: fetcher.v1.FetcherService.FetchTrades:input_type -> fetcher.v1.FetchTradesRequest
	4, // 2: fetcher.v1.FetcherService.StopFetchTrades:input_type -> fetcher.v1.Empty
	2, // 3: fetcher.v1.FetcherService.FetchTrades:output_type -> fetcher.v1.FetchTradesResponse
	3, // 4: fetcher.v1.FetcherService.StopFetchTrades:output_type -> fetcher.v1.StopFetchTradesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_fetcher_v1_trades_proto_init() }
func file_proto_fetcher_v1_trades_proto_init() {
	if File_proto_fetcher_v1_trades_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_fetcher_v1_trades_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTradesRequest); i {
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
		file_proto_fetcher_v1_trades_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_fetcher_v1_trades_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTradesResponse); i {
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
		file_proto_fetcher_v1_trades_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopFetchTradesResponse); i {
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
		file_proto_fetcher_v1_trades_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_fetcher_v1_trades_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fetcher_v1_trades_proto_goTypes,
		DependencyIndexes: file_proto_fetcher_v1_trades_proto_depIdxs,
		MessageInfos:      file_proto_fetcher_v1_trades_proto_msgTypes,
	}.Build()
	File_proto_fetcher_v1_trades_proto = out.File
	file_proto_fetcher_v1_trades_proto_rawDesc = nil
	file_proto_fetcher_v1_trades_proto_goTypes = nil
	file_proto_fetcher_v1_trades_proto_depIdxs = nil
}
