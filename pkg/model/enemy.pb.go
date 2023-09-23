// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pb/model/enemy.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnemyRank int32

const (
	EnemyRank_RANK_INVALID EnemyRank = 0
	EnemyRank_MINION       EnemyRank = 1
	EnemyRank_ELITE        EnemyRank = 2
	EnemyRank_LITTLE_BOSS  EnemyRank = 3
	EnemyRank_BIG_BOSS     EnemyRank = 4
)

// Enum value maps for EnemyRank.
var (
	EnemyRank_name = map[int32]string{
		0: "RANK_INVALID",
		1: "MINION",
		2: "ELITE",
		3: "LITTLE_BOSS",
		4: "BIG_BOSS",
	}
	EnemyRank_value = map[string]int32{
		"RANK_INVALID": 0,
		"MINION":       1,
		"ELITE":        2,
		"LITTLE_BOSS":  3,
		"BIG_BOSS":     4,
	}
)

func (x EnemyRank) Enum() *EnemyRank {
	p := new(EnemyRank)
	*p = x
	return p
}

func (x EnemyRank) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnemyRank) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_model_enemy_proto_enumTypes[0].Descriptor()
}

func (EnemyRank) Type() protoreflect.EnumType {
	return &file_pb_model_enemy_proto_enumTypes[0]
}

func (x EnemyRank) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnemyRank.Descriptor instead.
func (EnemyRank) EnumDescriptor() ([]byte, []int) {
	return file_pb_model_enemy_proto_rawDescGZIP(), []int{0}
}

type Enemy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Level      uint32           `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Weaknesses []DamageType     `protobuf:"varint,3,rep,packed,name=weaknesses,proto3,enum=model.DamageType" json:"weaknesses,omitempty"`
	DebuffRes  []*DebuffRES     `protobuf:"bytes,4,rep,name=debuff_res,json=debuffRes,proto3" json:"debuff_res,omitempty"`
	DamageRes  []*DamageRES     `protobuf:"bytes,5,rep,name=damage_res,json=damageRes,proto3" json:"damage_res,omitempty"`
	Rank       EnemyRank        `protobuf:"varint,6,opt,name=rank,proto3,enum=model.EnemyRank" json:"rank,omitempty"`
	BaseStats  *BaseStats       `protobuf:"bytes,7,opt,name=base_stats,json=baseStats,proto3" json:"base_stats,omitempty"`
	Parameters *structpb.Struct `protobuf:"bytes,10,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *Enemy) Reset() {
	*x = Enemy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_enemy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enemy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enemy) ProtoMessage() {}

func (x *Enemy) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_enemy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enemy.ProtoReflect.Descriptor instead.
func (*Enemy) Descriptor() ([]byte, []int) {
	return file_pb_model_enemy_proto_rawDescGZIP(), []int{0}
}

func (x *Enemy) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Enemy) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Enemy) GetWeaknesses() []DamageType {
	if x != nil {
		return x.Weaknesses
	}
	return nil
}

func (x *Enemy) GetDebuffRes() []*DebuffRES {
	if x != nil {
		return x.DebuffRes
	}
	return nil
}

func (x *Enemy) GetDamageRes() []*DamageRES {
	if x != nil {
		return x.DamageRes
	}
	return nil
}

func (x *Enemy) GetRank() EnemyRank {
	if x != nil {
		return x.Rank
	}
	return EnemyRank_RANK_INVALID
}

func (x *Enemy) GetBaseStats() *BaseStats {
	if x != nil {
		return x.BaseStats
	}
	return nil
}

func (x *Enemy) GetParameters() *structpb.Struct {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type DebuffRES struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag   BehaviorFlag `protobuf:"varint,1,opt,name=flag,proto3,enum=model.BehaviorFlag" json:"flag,omitempty"`
	Amount float64      `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DebuffRES) Reset() {
	*x = DebuffRES{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_enemy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebuffRES) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebuffRES) ProtoMessage() {}

func (x *DebuffRES) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_enemy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebuffRES.ProtoReflect.Descriptor instead.
func (*DebuffRES) Descriptor() ([]byte, []int) {
	return file_pb_model_enemy_proto_rawDescGZIP(), []int{1}
}

func (x *DebuffRES) GetFlag() BehaviorFlag {
	if x != nil {
		return x.Flag
	}
	return BehaviorFlag_INVALID_FLAG
}

func (x *DebuffRES) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type DamageRES struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   DamageType `protobuf:"varint,1,opt,name=type,proto3,enum=model.DamageType" json:"type,omitempty"`
	Amount float64    `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DamageRES) Reset() {
	*x = DamageRES{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_enemy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DamageRES) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DamageRES) ProtoMessage() {}

func (x *DamageRES) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_enemy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DamageRES.ProtoReflect.Descriptor instead.
func (*DamageRES) Descriptor() ([]byte, []int) {
	return file_pb_model_enemy_proto_rawDescGZIP(), []int{2}
}

func (x *DamageRES) GetType() DamageType {
	if x != nil {
		return x.Type
	}
	return DamageType_INVALID_DAMAGE_TYPE
}

func (x *DamageRES) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BaseStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Atk        float64 `protobuf:"fixed64,1,opt,name=atk,proto3" json:"atk,omitempty"`
	Def        float64 `protobuf:"fixed64,2,opt,name=def,proto3" json:"def,omitempty"`
	Hp         float64 `protobuf:"fixed64,3,opt,name=hp,proto3" json:"hp,omitempty"`
	Spd        float64 `protobuf:"fixed64,4,opt,name=spd,proto3" json:"spd,omitempty"`
	Stance     float64 `protobuf:"fixed64,5,opt,name=stance,proto3" json:"stance,omitempty"`
	CritChance float64 `protobuf:"fixed64,6,opt,name=crit_chance,json=critChance,proto3" json:"crit_chance,omitempty"`
	CritDmg    float64 `protobuf:"fixed64,7,opt,name=crit_dmg,json=critDmg,proto3" json:"crit_dmg,omitempty"`
	EffectRes  float64 `protobuf:"fixed64,8,opt,name=effect_res,json=effectRes,proto3" json:"effect_res,omitempty"`
	MinFatigue float64 `protobuf:"fixed64,9,opt,name=min_fatigue,json=minFatigue,proto3" json:"min_fatigue,omitempty"`
}

func (x *BaseStats) Reset() {
	*x = BaseStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_enemy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseStats) ProtoMessage() {}

func (x *BaseStats) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_enemy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseStats.ProtoReflect.Descriptor instead.
func (*BaseStats) Descriptor() ([]byte, []int) {
	return file_pb_model_enemy_proto_rawDescGZIP(), []int{3}
}

func (x *BaseStats) GetAtk() float64 {
	if x != nil {
		return x.Atk
	}
	return 0
}

func (x *BaseStats) GetDef() float64 {
	if x != nil {
		return x.Def
	}
	return 0
}

func (x *BaseStats) GetHp() float64 {
	if x != nil {
		return x.Hp
	}
	return 0
}

func (x *BaseStats) GetSpd() float64 {
	if x != nil {
		return x.Spd
	}
	return 0
}

func (x *BaseStats) GetStance() float64 {
	if x != nil {
		return x.Stance
	}
	return 0
}

func (x *BaseStats) GetCritChance() float64 {
	if x != nil {
		return x.CritChance
	}
	return 0
}

func (x *BaseStats) GetCritDmg() float64 {
	if x != nil {
		return x.CritDmg
	}
	return 0
}

func (x *BaseStats) GetEffectRes() float64 {
	if x != nil {
		return x.EffectRes
	}
	return 0
}

func (x *BaseStats) GetMinFatigue() float64 {
	if x != nil {
		return x.MinFatigue
	}
	return 0
}

var File_pb_model_enemy_proto protoreflect.FileDescriptor

var file_pb_model_enemy_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x65, 0x6d, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x62, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd4, 0x02, 0x0a, 0x05, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x31, 0x0a, 0x0a, 0x77, 0x65, 0x61, 0x6b, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x77, 0x65, 0x61, 0x6b, 0x6e,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x66, 0x66, 0x5f,
	0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x66, 0x66, 0x52, 0x45, 0x53, 0x52, 0x09, 0x64, 0x65, 0x62,
	0x75, 0x66, 0x66, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x0a, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x45, 0x53, 0x52, 0x09, 0x64, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e,
	0x65, 0x6d, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x2f, 0x0a,
	0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x37,
	0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x4c, 0x0a, 0x09, 0x44, 0x65, 0x62, 0x75, 0x66,
	0x66, 0x52, 0x45, 0x53, 0x12, 0x27, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x09, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x45, 0x53, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xe5, 0x01, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x74, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x74,
	0x6b, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x64, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x02, 0x68, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x70, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x73, 0x70, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x72, 0x69, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x5f, 0x64, 0x6d, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x63, 0x72, 0x69, 0x74, 0x44, 0x6d, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f,
	0x66, 0x61, 0x74, 0x69, 0x67, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6d,
	0x69, 0x6e, 0x46, 0x61, 0x74, 0x69, 0x67, 0x75, 0x65, 0x2a, 0x53, 0x0a, 0x09, 0x45, 0x6e, 0x65,
	0x6d, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x49,
	0x4f, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4c, 0x49, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x4c, 0x49, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x42, 0x4f, 0x53, 0x53, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x42, 0x49, 0x47, 0x5f, 0x42, 0x4f, 0x53, 0x53, 0x10, 0x04, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d,
	0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x2f, 0x73, 0x72, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_model_enemy_proto_rawDescOnce sync.Once
	file_pb_model_enemy_proto_rawDescData = file_pb_model_enemy_proto_rawDesc
)

func file_pb_model_enemy_proto_rawDescGZIP() []byte {
	file_pb_model_enemy_proto_rawDescOnce.Do(func() {
		file_pb_model_enemy_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_model_enemy_proto_rawDescData)
	})
	return file_pb_model_enemy_proto_rawDescData
}

var file_pb_model_enemy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_model_enemy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_model_enemy_proto_goTypes = []interface{}{
	(EnemyRank)(0),          // 0: model.EnemyRank
	(*Enemy)(nil),           // 1: model.Enemy
	(*DebuffRES)(nil),       // 2: model.DebuffRES
	(*DamageRES)(nil),       // 3: model.DamageRES
	(*BaseStats)(nil),       // 4: model.BaseStats
	(DamageType)(0),         // 5: model.DamageType
	(*structpb.Struct)(nil), // 6: google.protobuf.Struct
	(BehaviorFlag)(0),       // 7: model.BehaviorFlag
}
var file_pb_model_enemy_proto_depIdxs = []int32{
	5, // 0: model.Enemy.weaknesses:type_name -> model.DamageType
	2, // 1: model.Enemy.debuff_res:type_name -> model.DebuffRES
	3, // 2: model.Enemy.damage_res:type_name -> model.DamageRES
	0, // 3: model.Enemy.rank:type_name -> model.EnemyRank
	4, // 4: model.Enemy.base_stats:type_name -> model.BaseStats
	6, // 5: model.Enemy.parameters:type_name -> google.protobuf.Struct
	7, // 6: model.DebuffRES.flag:type_name -> model.BehaviorFlag
	5, // 7: model.DamageRES.type:type_name -> model.DamageType
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pb_model_enemy_proto_init() }
func file_pb_model_enemy_proto_init() {
	if File_pb_model_enemy_proto != nil {
		return
	}
	file_pb_model_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_model_enemy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enemy); i {
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
		file_pb_model_enemy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebuffRES); i {
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
		file_pb_model_enemy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DamageRES); i {
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
		file_pb_model_enemy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseStats); i {
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
			RawDescriptor: file_pb_model_enemy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_model_enemy_proto_goTypes,
		DependencyIndexes: file_pb_model_enemy_proto_depIdxs,
		EnumInfos:         file_pb_model_enemy_proto_enumTypes,
		MessageInfos:      file_pb_model_enemy_proto_msgTypes,
	}.Build()
	File_pb_model_enemy_proto = out.File
	file_pb_model_enemy_proto_rawDesc = nil
	file_pb_model_enemy_proto_goTypes = nil
	file_pb_model_enemy_proto_depIdxs = nil
}
