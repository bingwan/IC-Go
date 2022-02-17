// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.2
// source: types.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A non-interactive distributed key generation (NI-DKG) tag.
type NiDkgTag int32

const (
	NiDkgTag_NI_DKG_TAG_UNSPECIFIED    NiDkgTag = 0
	NiDkgTag_NI_DKG_TAG_LOW_THRESHOLD  NiDkgTag = 1
	NiDkgTag_NI_DKG_TAG_HIGH_THRESHOLD NiDkgTag = 2
)

// Enum value maps for NiDkgTag.
var (
	NiDkgTag_name = map[int32]string{
		0: "NI_DKG_TAG_UNSPECIFIED",
		1: "NI_DKG_TAG_LOW_THRESHOLD",
		2: "NI_DKG_TAG_HIGH_THRESHOLD",
	}
	NiDkgTag_value = map[string]int32{
		"NI_DKG_TAG_UNSPECIFIED":    0,
		"NI_DKG_TAG_LOW_THRESHOLD":  1,
		"NI_DKG_TAG_HIGH_THRESHOLD": 2,
	}
)

func (x NiDkgTag) Enum() *NiDkgTag {
	p := new(NiDkgTag)
	*p = x
	return p
}

func (x NiDkgTag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NiDkgTag) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[0].Descriptor()
}

func (NiDkgTag) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[0]
}

func (x NiDkgTag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NiDkgTag.Descriptor instead.
func (NiDkgTag) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

type PrincipalId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw []byte `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (x *PrincipalId) Reset() {
	*x = PrincipalId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrincipalId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipalId) ProtoMessage() {}

func (x *PrincipalId) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipalId.ProtoReflect.Descriptor instead.
func (*PrincipalId) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *PrincipalId) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

type CanisterId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrincipalId *PrincipalId `protobuf:"bytes,1,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
}

func (x *CanisterId) Reset() {
	*x = CanisterId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanisterId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanisterId) ProtoMessage() {}

func (x *CanisterId) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanisterId.ProtoReflect.Descriptor instead.
func (*CanisterId) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *CanisterId) GetPrincipalId() *PrincipalId {
	if x != nil {
		return x.PrincipalId
	}
	return nil
}

type SubnetId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrincipalId *PrincipalId `protobuf:"bytes,1,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
}

func (x *SubnetId) Reset() {
	*x = SubnetId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubnetId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetId) ProtoMessage() {}

func (x *SubnetId) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetId.ProtoReflect.Descriptor instead.
func (*SubnetId) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *SubnetId) GetPrincipalId() *PrincipalId {
	if x != nil {
		return x.PrincipalId
	}
	return nil
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrincipalId *PrincipalId `protobuf:"bytes,1,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *UserId) GetPrincipalId() *PrincipalId {
	if x != nil {
		return x.PrincipalId
	}
	return nil
}

type NodeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrincipalId *PrincipalId `protobuf:"bytes,1,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
}

func (x *NodeId) Reset() {
	*x = NodeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeId) ProtoMessage() {}

func (x *NodeId) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeId.ProtoReflect.Descriptor instead.
func (*NodeId) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{4}
}

func (x *NodeId) GetPrincipalId() *PrincipalId {
	if x != nil {
		return x.PrincipalId
	}
	return nil
}

// A non-interactive distributed key generation (NI-DKG) ID.
type NiDkgId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBlockHeight uint64                 `protobuf:"varint,1,opt,name=start_block_height,json=startBlockHeight,proto3" json:"start_block_height,omitempty"`
	DealerSubnet     []byte                 `protobuf:"bytes,2,opt,name=dealer_subnet,json=dealerSubnet,proto3" json:"dealer_subnet,omitempty"`
	DkgTag           NiDkgTag               `protobuf:"varint,4,opt,name=dkg_tag,json=dkgTag,proto3,enum=ic_registry.NiDkgTag" json:"dkg_tag,omitempty"`
	RemoteTargetId   *wrapperspb.BytesValue `protobuf:"bytes,5,opt,name=remote_target_id,json=remoteTargetId,proto3" json:"remote_target_id,omitempty"`
}

func (x *NiDkgId) Reset() {
	*x = NiDkgId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NiDkgId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NiDkgId) ProtoMessage() {}

func (x *NiDkgId) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NiDkgId.ProtoReflect.Descriptor instead.
func (*NiDkgId) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{5}
}

func (x *NiDkgId) GetStartBlockHeight() uint64 {
	if x != nil {
		return x.StartBlockHeight
	}
	return 0
}

func (x *NiDkgId) GetDealerSubnet() []byte {
	if x != nil {
		return x.DealerSubnet
	}
	return nil
}

func (x *NiDkgId) GetDkgTag() NiDkgTag {
	if x != nil {
		return x.DkgTag
	}
	return NiDkgTag_NI_DKG_TAG_UNSPECIFIED
}

func (x *NiDkgId) GetRemoteTargetId() *wrapperspb.BytesValue {
	if x != nil {
		return x.RemoteTargetId
	}
	return nil
}

type NominalCycles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	High uint64 `protobuf:"varint,1,opt,name=high,proto3" json:"high,omitempty"`
	Low  uint64 `protobuf:"varint,2,opt,name=low,proto3" json:"low,omitempty"`
}

func (x *NominalCycles) Reset() {
	*x = NominalCycles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NominalCycles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NominalCycles) ProtoMessage() {}

func (x *NominalCycles) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NominalCycles.ProtoReflect.Descriptor instead.
func (*NominalCycles) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{6}
}

func (x *NominalCycles) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *NominalCycles) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69,
	0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0b, 0x50, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x61, 0x77, 0x22, 0x49, 0x0a, 0x0a, 0x43,
	0x61, 0x6e, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0c, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x63, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x49, 0x64, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x22,
	0x45, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0c, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x3b, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64,
	0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xea, 0x01,
	0x0a, 0x07, 0x4e, 0x69, 0x44, 0x6b, 0x67, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x07,
	0x64, 0x6b, 0x67, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x69, 0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4e, 0x69, 0x44, 0x6b,
	0x67, 0x54, 0x61, 0x67, 0x52, 0x06, 0x64, 0x6b, 0x67, 0x54, 0x61, 0x67, 0x12, 0x45, 0x0a, 0x10,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x22, 0x35, 0x0a, 0x0d, 0x4e, 0x6f,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x69, 0x67, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c, 0x6f,
	0x77, 0x2a, 0x63, 0x0a, 0x08, 0x4e, 0x69, 0x44, 0x6b, 0x67, 0x54, 0x61, 0x67, 0x12, 0x1a, 0x0a,
	0x16, 0x4e, 0x49, 0x5f, 0x44, 0x4b, 0x47, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x4e, 0x49, 0x5f,
	0x44, 0x4b, 0x47, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x48, 0x52, 0x45,
	0x53, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4e, 0x49, 0x5f, 0x44, 0x4b,
	0x47, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53,
	0x48, 0x4f, 0x4c, 0x44, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_types_proto_goTypes = []interface{}{
	(NiDkgTag)(0),                 // 0: ic_registry.NiDkgTag
	(*PrincipalId)(nil),           // 1: ic_registry.PrincipalId
	(*CanisterId)(nil),            // 2: ic_registry.CanisterId
	(*SubnetId)(nil),              // 3: ic_registry.SubnetId
	(*UserId)(nil),                // 4: ic_registry.UserId
	(*NodeId)(nil),                // 5: ic_registry.NodeId
	(*NiDkgId)(nil),               // 6: ic_registry.NiDkgId
	(*NominalCycles)(nil),         // 7: ic_registry.NominalCycles
	(*wrapperspb.BytesValue)(nil), // 8: google.protobuf.BytesValue
}
var file_types_proto_depIdxs = []int32{
	1, // 0: ic_registry.CanisterId.principal_id:type_name -> ic_registry.PrincipalId
	1, // 1: ic_registry.SubnetId.principal_id:type_name -> ic_registry.PrincipalId
	1, // 2: ic_registry.UserId.principal_id:type_name -> ic_registry.PrincipalId
	1, // 3: ic_registry.NodeId.principal_id:type_name -> ic_registry.PrincipalId
	0, // 4: ic_registry.NiDkgId.dkg_tag:type_name -> ic_registry.NiDkgTag
	8, // 5: ic_registry.NiDkgId.remote_target_id:type_name -> google.protobuf.BytesValue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrincipalId); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanisterId); i {
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
		file_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubnetId); i {
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
		file_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeId); i {
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
		file_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NiDkgId); i {
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
		file_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NominalCycles); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		EnumInfos:         file_types_proto_enumTypes,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}