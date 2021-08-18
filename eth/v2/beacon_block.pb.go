// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: eth/v2/beacon_block.proto

package eth

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	github_com_prysmaticlabs_eth2_types "github.com/prysmaticlabs/eth2-types"
	_ "github.com/prysmaticlabs/ethereumapis/eth/ext"
	v1 "github.com/prysmaticlabs/ethereumapis/eth/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BlockRequestV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
}

func (x *BlockRequestV2) Reset() {
	*x = BlockRequestV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v2_beacon_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequestV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequestV2) ProtoMessage() {}

func (x *BlockRequestV2) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v2_beacon_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequestV2.ProtoReflect.Descriptor instead.
func (*BlockRequestV2) Descriptor() ([]byte, []int) {
	return file_eth_v2_beacon_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockRequestV2) GetBlockId() []byte {
	if x != nil {
		return x.BlockId
	}
	return nil
}

type BlockResponseV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *BeaconBlockContainerV2 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlockResponseV2) Reset() {
	*x = BlockResponseV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v2_beacon_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponseV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponseV2) ProtoMessage() {}

func (x *BlockResponseV2) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v2_beacon_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponseV2.ProtoReflect.Descriptor instead.
func (*BlockResponseV2) Descriptor() ([]byte, []int) {
	return file_eth_v2_beacon_block_proto_rawDescGZIP(), []int{1}
}

func (x *BlockResponseV2) GetData() *BeaconBlockContainerV2 {
	if x != nil {
		return x.Data
	}
	return nil
}

type BlockSSZResponseV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlockSSZResponseV2) Reset() {
	*x = BlockSSZResponseV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v2_beacon_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSSZResponseV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSSZResponseV2) ProtoMessage() {}

func (x *BlockSSZResponseV2) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v2_beacon_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSSZResponseV2.ProtoReflect.Descriptor instead.
func (*BlockSSZResponseV2) Descriptor() ([]byte, []int) {
	return file_eth_v2_beacon_block_proto_rawDescGZIP(), []int{2}
}

func (x *BlockSSZResponseV2) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BeaconBlockContainerV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phase0Block *v1.BeaconBlock    `protobuf:"bytes,1,opt,name=phase0Block,proto3" json:"phase0Block,omitempty"`
	AltairBlock *BeaconBlockAltair `protobuf:"bytes,2,opt,name=altairBlock,proto3" json:"altairBlock,omitempty"`
	Signature   []byte             `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty" ssz-size:"96"`
}

func (x *BeaconBlockContainerV2) Reset() {
	*x = BeaconBlockContainerV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v2_beacon_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconBlockContainerV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconBlockContainerV2) ProtoMessage() {}

func (x *BeaconBlockContainerV2) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v2_beacon_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconBlockContainerV2.ProtoReflect.Descriptor instead.
func (*BeaconBlockContainerV2) Descriptor() ([]byte, []int) {
	return file_eth_v2_beacon_block_proto_rawDescGZIP(), []int{3}
}

func (x *BeaconBlockContainerV2) GetPhase0Block() *v1.BeaconBlock {
	if x != nil {
		return x.Phase0Block
	}
	return nil
}

func (x *BeaconBlockContainerV2) GetAltairBlock() *BeaconBlockAltair {
	if x != nil {
		return x.AltairBlock
	}
	return nil
}

func (x *BeaconBlockContainerV2) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type BeaconBlockAltair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot          github_com_prysmaticlabs_eth2_types.Slot           `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	ProposerIndex github_com_prysmaticlabs_eth2_types.ValidatorIndex `protobuf:"varint,2,opt,name=proposer_index,json=proposerIndex,proto3" json:"proposer_index,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.ValidatorIndex"`
	ParentRoot    []byte                                             `protobuf:"bytes,3,opt,name=parent_root,json=parentRoot,proto3" json:"parent_root,omitempty" ssz-size:"32"`
	StateRoot     []byte                                             `protobuf:"bytes,4,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty" ssz-size:"32"`
	Body          *BeaconBlockBodyAltair                             `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *BeaconBlockAltair) Reset() {
	*x = BeaconBlockAltair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v2_beacon_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconBlockAltair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconBlockAltair) ProtoMessage() {}

func (x *BeaconBlockAltair) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v2_beacon_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconBlockAltair.ProtoReflect.Descriptor instead.
func (*BeaconBlockAltair) Descriptor() ([]byte, []int) {
	return file_eth_v2_beacon_block_proto_rawDescGZIP(), []int{4}
}

func (x *BeaconBlockAltair) GetSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *BeaconBlockAltair) GetProposerIndex() github_com_prysmaticlabs_eth2_types.ValidatorIndex {
	if x != nil {
		return x.ProposerIndex
	}
	return github_com_prysmaticlabs_eth2_types.ValidatorIndex(0)
}

func (x *BeaconBlockAltair) GetParentRoot() []byte {
	if x != nil {
		return x.ParentRoot
	}
	return nil
}

func (x *BeaconBlockAltair) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

func (x *BeaconBlockAltair) GetBody() *BeaconBlockBodyAltair {
	if x != nil {
		return x.Body
	}
	return nil
}

type BeaconBlockBodyAltair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RandaoReveal      []byte                    `protobuf:"bytes,1,opt,name=randao_reveal,json=randaoReveal,proto3" json:"randao_reveal,omitempty" ssz-size:"96"`
	Eth1Data          *v1.Eth1Data              `protobuf:"bytes,2,opt,name=eth1_data,json=eth1Data,proto3" json:"eth1_data,omitempty"`
	Graffiti          []byte                    `protobuf:"bytes,3,opt,name=graffiti,proto3" json:"graffiti,omitempty" ssz-size:"32"`
	ProposerSlashings []*v1.ProposerSlashing    `protobuf:"bytes,4,rep,name=proposer_slashings,json=proposerSlashings,proto3" json:"proposer_slashings,omitempty" ssz-max:"16"`
	AttesterSlashings []*v1.AttesterSlashing    `protobuf:"bytes,5,rep,name=attester_slashings,json=attesterSlashings,proto3" json:"attester_slashings,omitempty" ssz-max:"2"`
	Attestations      []*v1.Attestation         `protobuf:"bytes,6,rep,name=attestations,proto3" json:"attestations,omitempty" ssz-max:"128"`
	Deposits          []*v1.Deposit             `protobuf:"bytes,7,rep,name=deposits,proto3" json:"deposits,omitempty" ssz-max:"16"`
	VoluntaryExits    []*v1.SignedVoluntaryExit `protobuf:"bytes,8,rep,name=voluntary_exits,json=voluntaryExits,proto3" json:"voluntary_exits,omitempty" ssz-max:"16"`
	SyncAggregate     *v1.SyncAggregate         `protobuf:"bytes,9,opt,name=sync_aggregate,json=syncAggregate,proto3" json:"sync_aggregate,omitempty"`
}

func (x *BeaconBlockBodyAltair) Reset() {
	*x = BeaconBlockBodyAltair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v2_beacon_block_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconBlockBodyAltair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconBlockBodyAltair) ProtoMessage() {}

func (x *BeaconBlockBodyAltair) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v2_beacon_block_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconBlockBodyAltair.ProtoReflect.Descriptor instead.
func (*BeaconBlockBodyAltair) Descriptor() ([]byte, []int) {
	return file_eth_v2_beacon_block_proto_rawDescGZIP(), []int{5}
}

func (x *BeaconBlockBodyAltair) GetRandaoReveal() []byte {
	if x != nil {
		return x.RandaoReveal
	}
	return nil
}

func (x *BeaconBlockBodyAltair) GetEth1Data() *v1.Eth1Data {
	if x != nil {
		return x.Eth1Data
	}
	return nil
}

func (x *BeaconBlockBodyAltair) GetGraffiti() []byte {
	if x != nil {
		return x.Graffiti
	}
	return nil
}

func (x *BeaconBlockBodyAltair) GetProposerSlashings() []*v1.ProposerSlashing {
	if x != nil {
		return x.ProposerSlashings
	}
	return nil
}

func (x *BeaconBlockBodyAltair) GetAttesterSlashings() []*v1.AttesterSlashing {
	if x != nil {
		return x.AttesterSlashings
	}
	return nil
}

func (x *BeaconBlockBodyAltair) GetAttestations() []*v1.Attestation {
	if x != nil {
		return x.Attestations
	}
	return nil
}

func (x *BeaconBlockBodyAltair) GetDeposits() []*v1.Deposit {
	if x != nil {
		return x.Deposits
	}
	return nil
}

func (x *BeaconBlockBodyAltair) GetVoluntaryExits() []*v1.SignedVoluntaryExit {
	if x != nil {
		return x.VoluntaryExits
	}
	return nil
}

func (x *BeaconBlockBodyAltair) GetSyncAggregate() *v1.SyncAggregate {
	if x != nil {
		return x.SyncAggregate
	}
	return nil
}

var File_eth_v2_beacon_block_proto protoreflect.FileDescriptor

var file_eth_v2_beacon_block_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x62,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e,
	0x76, 0x32, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78,
	0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65,
	0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2b, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x4e, 0x0a,
	0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x32,
	0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32,
	0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x56, 0x32, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a,
	0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x53, 0x5a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc4, 0x01, 0x0a, 0x16, 0x42, 0x65, 0x61, 0x63,
	0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x56, 0x32, 0x12, 0x3e, 0x0a, 0x0b, 0x70, 0x68, 0x61, 0x73, 0x65, 0x30, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0b, 0x70, 0x68, 0x61, 0x73, 0x65, 0x30, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x44, 0x0a, 0x0b, 0x61, 0x6c, 0x74, 0x61, 0x69, 0x72, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x6c, 0x74, 0x61, 0x69, 0x72, 0x52, 0x0b, 0x61, 0x6c, 0x74,
	0x61, 0x69, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18,
	0x02, 0x39, 0x36, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xc0,
	0x02, 0x0a, 0x11, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x6c,
	0x74, 0x61, 0x69, 0x72, 0x12, 0x40, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74,
	0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x5d, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x36,
	0x82, 0xb5, 0x18, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68,
	0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x27, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02,
	0x33, 0x32, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x25,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x42, 0x6f, 0x64, 0x79, 0x41, 0x6c, 0x74, 0x61, 0x69, 0x72, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0xfa, 0x04, 0x0a, 0x15, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x42, 0x6f, 0x64, 0x79, 0x41, 0x6c, 0x74, 0x61, 0x69, 0x72, 0x12, 0x2b, 0x0a, 0x0d, 0x72,
	0x61, 0x6e, 0x64, 0x61, 0x6f, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x0c, 0x72, 0x61, 0x6e, 0x64,
	0x61, 0x6f, 0x52, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x09, 0x65, 0x74, 0x68, 0x31,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x74,
	0x68, 0x31, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x65, 0x74, 0x68, 0x31, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x22, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x66, 0x66, 0x69, 0x74, 0x69, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x08, 0x67, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x74, 0x69, 0x12, 0x58, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x42, 0x06, 0x92, 0xb5, 0x18, 0x02, 0x31, 0x36, 0x52, 0x11, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x57,
	0x0a, 0x12, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x05, 0x92,
	0xb5, 0x18, 0x01, 0x32, 0x52, 0x11, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x49, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x92, 0xb5, 0x18,
	0x03, 0x31, 0x32, 0x38, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x42, 0x06,
	0x92, 0xb5, 0x18, 0x02, 0x31, 0x36, 0x52, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73,
	0x12, 0x55, 0x0a, 0x0f, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x78,
	0x69, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x61, 0x72, 0x79, 0x45, 0x78, 0x69, 0x74, 0x42,
	0x06, 0x92, 0xb5, 0x18, 0x02, 0x31, 0x36, 0x52, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x61,
	0x72, 0x79, 0x45, 0x78, 0x69, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52,
	0x0d, 0x73, 0x79, 0x6e, 0x63, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x80,
	0x01, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x42, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x3b, 0x65, 0x74, 0x68, 0xaa, 0x02, 0x0f,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x56, 0x32, 0xca,
	0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eth_v2_beacon_block_proto_rawDescOnce sync.Once
	file_eth_v2_beacon_block_proto_rawDescData = file_eth_v2_beacon_block_proto_rawDesc
)

func file_eth_v2_beacon_block_proto_rawDescGZIP() []byte {
	file_eth_v2_beacon_block_proto_rawDescOnce.Do(func() {
		file_eth_v2_beacon_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v2_beacon_block_proto_rawDescData)
	})
	return file_eth_v2_beacon_block_proto_rawDescData
}

var file_eth_v2_beacon_block_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eth_v2_beacon_block_proto_goTypes = []interface{}{
	(*BlockRequestV2)(nil),         // 0: ethereum.eth.v2.BlockRequestV2
	(*BlockResponseV2)(nil),        // 1: ethereum.eth.v2.BlockResponseV2
	(*BlockSSZResponseV2)(nil),     // 2: ethereum.eth.v2.BlockSSZResponseV2
	(*BeaconBlockContainerV2)(nil), // 3: ethereum.eth.v2.BeaconBlockContainerV2
	(*BeaconBlockAltair)(nil),      // 4: ethereum.eth.v2.BeaconBlockAltair
	(*BeaconBlockBodyAltair)(nil),  // 5: ethereum.eth.v2.BeaconBlockBodyAltair
	(*v1.BeaconBlock)(nil),         // 6: ethereum.eth.v1.BeaconBlock
	(*v1.Eth1Data)(nil),            // 7: ethereum.eth.v1.Eth1Data
	(*v1.ProposerSlashing)(nil),    // 8: ethereum.eth.v1.ProposerSlashing
	(*v1.AttesterSlashing)(nil),    // 9: ethereum.eth.v1.AttesterSlashing
	(*v1.Attestation)(nil),         // 10: ethereum.eth.v1.Attestation
	(*v1.Deposit)(nil),             // 11: ethereum.eth.v1.Deposit
	(*v1.SignedVoluntaryExit)(nil), // 12: ethereum.eth.v1.SignedVoluntaryExit
	(*v1.SyncAggregate)(nil),       // 13: ethereum.eth.v1.SyncAggregate
}
var file_eth_v2_beacon_block_proto_depIdxs = []int32{
	3,  // 0: ethereum.eth.v2.BlockResponseV2.data:type_name -> ethereum.eth.v2.BeaconBlockContainerV2
	6,  // 1: ethereum.eth.v2.BeaconBlockContainerV2.phase0Block:type_name -> ethereum.eth.v1.BeaconBlock
	4,  // 2: ethereum.eth.v2.BeaconBlockContainerV2.altairBlock:type_name -> ethereum.eth.v2.BeaconBlockAltair
	5,  // 3: ethereum.eth.v2.BeaconBlockAltair.body:type_name -> ethereum.eth.v2.BeaconBlockBodyAltair
	7,  // 4: ethereum.eth.v2.BeaconBlockBodyAltair.eth1_data:type_name -> ethereum.eth.v1.Eth1Data
	8,  // 5: ethereum.eth.v2.BeaconBlockBodyAltair.proposer_slashings:type_name -> ethereum.eth.v1.ProposerSlashing
	9,  // 6: ethereum.eth.v2.BeaconBlockBodyAltair.attester_slashings:type_name -> ethereum.eth.v1.AttesterSlashing
	10, // 7: ethereum.eth.v2.BeaconBlockBodyAltair.attestations:type_name -> ethereum.eth.v1.Attestation
	11, // 8: ethereum.eth.v2.BeaconBlockBodyAltair.deposits:type_name -> ethereum.eth.v1.Deposit
	12, // 9: ethereum.eth.v2.BeaconBlockBodyAltair.voluntary_exits:type_name -> ethereum.eth.v1.SignedVoluntaryExit
	13, // 10: ethereum.eth.v2.BeaconBlockBodyAltair.sync_aggregate:type_name -> ethereum.eth.v1.SyncAggregate
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eth_v2_beacon_block_proto_init() }
func file_eth_v2_beacon_block_proto_init() {
	if File_eth_v2_beacon_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eth_v2_beacon_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequestV2); i {
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
		file_eth_v2_beacon_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponseV2); i {
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
		file_eth_v2_beacon_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSSZResponseV2); i {
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
		file_eth_v2_beacon_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconBlockContainerV2); i {
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
		file_eth_v2_beacon_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconBlockAltair); i {
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
		file_eth_v2_beacon_block_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconBlockBodyAltair); i {
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
			RawDescriptor: file_eth_v2_beacon_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v2_beacon_block_proto_goTypes,
		DependencyIndexes: file_eth_v2_beacon_block_proto_depIdxs,
		MessageInfos:      file_eth_v2_beacon_block_proto_msgTypes,
	}.Build()
	File_eth_v2_beacon_block_proto = out.File
	file_eth_v2_beacon_block_proto_rawDesc = nil
	file_eth_v2_beacon_block_proto_goTypes = nil
	file_eth_v2_beacon_block_proto_depIdxs = nil
}
