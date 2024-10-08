// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: relay/relay.proto

package relay

import (
	common "github.com/Layr-Labs/eigenda/api/grpc/common"
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

// Request a specific chunk
type GetChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash of the blob's header.
	HeaderHash []byte `protobuf:"bytes,1,opt,name=header_hash,json=headerHash,proto3" json:"header_hash,omitempty"`
	// The index of the chunk within the blob.
	ChunkIndex uint32 `protobuf:"varint,2,opt,name=chunk_index,json=chunkIndex,proto3" json:"chunk_index,omitempty"`
}

func (x *GetChunkRequest) Reset() {
	*x = GetChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkRequest) ProtoMessage() {}

func (x *GetChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkRequest.ProtoReflect.Descriptor instead.
func (*GetChunkRequest) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{0}
}

func (x *GetChunkRequest) GetHeaderHash() []byte {
	if x != nil {
		return x.HeaderHash
	}
	return nil
}

func (x *GetChunkRequest) GetChunkIndex() uint32 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

// Reply to GetChunkRequest
type GetChunkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chunk requested.
	Chunk *common.ChunkData `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *GetChunkReply) Reset() {
	*x = GetChunkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkReply) ProtoMessage() {}

func (x *GetChunkReply) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkReply.ProtoReflect.Descriptor instead.
func (*GetChunkReply) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{1}
}

func (x *GetChunkReply) GetChunk() *common.ChunkData {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_relay_relay_proto protoreflect.FileDescriptor

var file_relay_relay_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x38, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0x41, 0x0a,
	0x05, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x15, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c,
	0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x64, 0x61,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_relay_proto_rawDescOnce sync.Once
	file_relay_relay_proto_rawDescData = file_relay_relay_proto_rawDesc
)

func file_relay_relay_proto_rawDescGZIP() []byte {
	file_relay_relay_proto_rawDescOnce.Do(func() {
		file_relay_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_relay_proto_rawDescData)
	})
	return file_relay_relay_proto_rawDescData
}

var file_relay_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relay_relay_proto_goTypes = []interface{}{
	(*GetChunkRequest)(nil),  // 0: node.GetChunkRequest
	(*GetChunkReply)(nil),    // 1: node.GetChunkReply
	(*common.ChunkData)(nil), // 2: common.ChunkData
}
var file_relay_relay_proto_depIdxs = []int32{
	2, // 0: node.GetChunkReply.chunk:type_name -> common.ChunkData
	0, // 1: node.Relay.GetChunk:input_type -> node.GetChunkRequest
	1, // 2: node.Relay.GetChunk:output_type -> node.GetChunkReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relay_relay_proto_init() }
func file_relay_relay_proto_init() {
	if File_relay_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChunkRequest); i {
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
		file_relay_relay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChunkReply); i {
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
			RawDescriptor: file_relay_relay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relay_relay_proto_goTypes,
		DependencyIndexes: file_relay_relay_proto_depIdxs,
		MessageInfos:      file_relay_relay_proto_msgTypes,
	}.Build()
	File_relay_relay_proto = out.File
	file_relay_relay_proto_rawDesc = nil
	file_relay_relay_proto_goTypes = nil
	file_relay_relay_proto_depIdxs = nil
}
