// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: ProtoLogistic/ProtoLogistic.proto

package ProtoLogistic

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TruckType int32

const (
	TruckType_UNDEFINED_TRUCK TruckType = 0
	TruckType_RETAIL_TRUCK    TruckType = 1
	TruckType_NORMAL_TRUCK    TruckType = 2
)

// Enum value maps for TruckType.
var (
	TruckType_name = map[int32]string{
		0: "UNDEFINED_TRUCK",
		1: "RETAIL_TRUCK",
		2: "NORMAL_TRUCK",
	}
	TruckType_value = map[string]int32{
		"UNDEFINED_TRUCK": 0,
		"RETAIL_TRUCK":    1,
		"NORMAL_TRUCK":    2,
	}
)

func (x TruckType) Enum() *TruckType {
	p := new(TruckType)
	*p = x
	return p
}

func (x TruckType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TruckType) Descriptor() protoreflect.EnumDescriptor {
	return file_ProtoLogistic_ProtoLogistic_proto_enumTypes[0].Descriptor()
}

func (TruckType) Type() protoreflect.EnumType {
	return &file_ProtoLogistic_ProtoLogistic_proto_enumTypes[0]
}

func (x TruckType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TruckType.Descriptor instead.
func (TruckType) EnumDescriptor() ([]byte, []int) {
	return file_ProtoLogistic_ProtoLogistic_proto_rawDescGZIP(), []int{0}
}

type PackageType int32

const (
	PackageType_UNDEFINED_PACKAGE PackageType = 0
	PackageType_RETAIL_PACKAGE    PackageType = 1
	PackageType_PRIORITY_PACKAGE  PackageType = 2
	PackageType_NORMAL_PACKAGE    PackageType = 3
)

// Enum value maps for PackageType.
var (
	PackageType_name = map[int32]string{
		0: "UNDEFINED_PACKAGE",
		1: "RETAIL_PACKAGE",
		2: "PRIORITY_PACKAGE",
		3: "NORMAL_PACKAGE",
	}
	PackageType_value = map[string]int32{
		"UNDEFINED_PACKAGE": 0,
		"RETAIL_PACKAGE":    1,
		"PRIORITY_PACKAGE":  2,
		"NORMAL_PACKAGE":    3,
	}
)

func (x PackageType) Enum() *PackageType {
	p := new(PackageType)
	*p = x
	return p
}

func (x PackageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackageType) Descriptor() protoreflect.EnumDescriptor {
	return file_ProtoLogistic_ProtoLogistic_proto_enumTypes[1].Descriptor()
}

func (PackageType) Type() protoreflect.EnumType {
	return &file_ProtoLogistic_ProtoLogistic_proto_enumTypes[1]
}

func (x PackageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackageType.Descriptor instead.
func (PackageType) EnumDescriptor() ([]byte, []int) {
	return file_ProtoLogistic_ProtoLogistic_proto_rawDescGZIP(), []int{1}
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDPaquete   string      `protobuf:"bytes,1,opt,name=IDPaquete,proto3" json:"IDPaquete,omitempty"`
	Tipo        PackageType `protobuf:"varint,2,opt,name=Tipo,proto3,enum=ProtoLogistic.PackageType" json:"Tipo,omitempty"`
	Valor       int32       `protobuf:"varint,3,opt,name=Valor,proto3" json:"Valor,omitempty"`
	Origen      string      `protobuf:"bytes,4,opt,name=Origen,proto3" json:"Origen,omitempty"`
	Destino     string      `protobuf:"bytes,5,opt,name=Destino,proto3" json:"Destino,omitempty"`
	Intentos    int32       `protobuf:"varint,6,opt,name=Intentos,proto3" json:"Intentos,omitempty"`
	Estado      string      `protobuf:"bytes,7,opt,name=Estado,proto3" json:"Estado,omitempty"`
	Seguimiento string      `protobuf:"bytes,8,opt,name=Seguimiento,proto3" json:"Seguimiento,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoLogistic_ProtoLogistic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoLogistic_ProtoLogistic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_ProtoLogistic_ProtoLogistic_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetIDPaquete() string {
	if x != nil {
		return x.IDPaquete
	}
	return ""
}

func (x *Package) GetTipo() PackageType {
	if x != nil {
		return x.Tipo
	}
	return PackageType_UNDEFINED_PACKAGE
}

func (x *Package) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Package) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *Package) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *Package) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *Package) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *Package) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

type Truck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type TruckType `protobuf:"varint,1,opt,name=Type,proto3,enum=ProtoLogistic.TruckType" json:"Type,omitempty"`
}

func (x *Truck) Reset() {
	*x = Truck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoLogistic_ProtoLogistic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Truck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Truck) ProtoMessage() {}

func (x *Truck) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoLogistic_ProtoLogistic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Truck.ProtoReflect.Descriptor instead.
func (*Truck) Descriptor() ([]byte, []int) {
	return file_ProtoLogistic_ProtoLogistic_proto_rawDescGZIP(), []int{1}
}

func (x *Truck) GetType() TruckType {
	if x != nil {
		return x.Type
	}
	return TruckType_UNDEFINED_TRUCK
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoLogistic_ProtoLogistic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoLogistic_ProtoLogistic_proto_msgTypes[2]
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
	return file_ProtoLogistic_ProtoLogistic_proto_rawDescGZIP(), []int{2}
}

var File_ProtoLogistic_ProtoLogistic_proto protoreflect.FileDescriptor

var file_ProtoLogistic_ProtoLogistic_proto_rawDesc = []byte{
	0x0a, 0x21, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x22, 0xf5, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x49, 0x44, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x49, 0x44, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04,
	0x54, 0x69, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53,
	0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x05, 0x54, 0x72,
	0x75, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x54, 0x72, 0x75, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2a, 0x44, 0x0a, 0x09, 0x54, 0x72,
	0x75, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x5f, 0x54, 0x52, 0x55, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x52, 0x55, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x54, 0x52, 0x55, 0x43, 0x4b, 0x10, 0x02,
	0x2a, 0x62, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x43,
	0x4b, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c,
	0x5f, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52,
	0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x50, 0x41, 0x43, 0x4b, 0x41,
	0x47, 0x45, 0x10, 0x03, 0x32, 0x98, 0x02, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x54,
	0x72, 0x75, 0x63, 0x6b, 0x1a, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x53,
	0x6f, 0x6c, 0x61, 0x72, 0x56, 0x2f, 0x54, 0x61, 0x72, 0x65, 0x61, 0x31, 0x2f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ProtoLogistic_ProtoLogistic_proto_rawDescOnce sync.Once
	file_ProtoLogistic_ProtoLogistic_proto_rawDescData = file_ProtoLogistic_ProtoLogistic_proto_rawDesc
)

func file_ProtoLogistic_ProtoLogistic_proto_rawDescGZIP() []byte {
	file_ProtoLogistic_ProtoLogistic_proto_rawDescOnce.Do(func() {
		file_ProtoLogistic_ProtoLogistic_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProtoLogistic_ProtoLogistic_proto_rawDescData)
	})
	return file_ProtoLogistic_ProtoLogistic_proto_rawDescData
}

var file_ProtoLogistic_ProtoLogistic_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ProtoLogistic_ProtoLogistic_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ProtoLogistic_ProtoLogistic_proto_goTypes = []interface{}{
	(TruckType)(0),   // 0: ProtoLogistic.TruckType
	(PackageType)(0), // 1: ProtoLogistic.PackageType
	(*Package)(nil),  // 2: ProtoLogistic.Package
	(*Truck)(nil),    // 3: ProtoLogistic.Truck
	(*Empty)(nil),    // 4: ProtoLogistic.Empty
}
var file_ProtoLogistic_ProtoLogistic_proto_depIdxs = []int32{
	1, // 0: ProtoLogistic.Package.Tipo:type_name -> ProtoLogistic.PackageType
	0, // 1: ProtoLogistic.Truck.Type:type_name -> ProtoLogistic.TruckType
	2, // 2: ProtoLogistic.ProtoLogisticService.DeliverPackage:input_type -> ProtoLogistic.Package
	2, // 3: ProtoLogistic.ProtoLogisticService.CheckStatus:input_type -> ProtoLogistic.Package
	3, // 4: ProtoLogistic.ProtoLogisticService.AskPackage:input_type -> ProtoLogistic.Truck
	2, // 5: ProtoLogistic.ProtoLogisticService.FinishPackage:input_type -> ProtoLogistic.Package
	4, // 6: ProtoLogistic.ProtoLogisticService.DeliverPackage:output_type -> ProtoLogistic.Empty
	2, // 7: ProtoLogistic.ProtoLogisticService.CheckStatus:output_type -> ProtoLogistic.Package
	2, // 8: ProtoLogistic.ProtoLogisticService.AskPackage:output_type -> ProtoLogistic.Package
	4, // 9: ProtoLogistic.ProtoLogisticService.FinishPackage:output_type -> ProtoLogistic.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ProtoLogistic_ProtoLogistic_proto_init() }
func file_ProtoLogistic_ProtoLogistic_proto_init() {
	if File_ProtoLogistic_ProtoLogistic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProtoLogistic_ProtoLogistic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_ProtoLogistic_ProtoLogistic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Truck); i {
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
		file_ProtoLogistic_ProtoLogistic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_ProtoLogistic_ProtoLogistic_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ProtoLogistic_ProtoLogistic_proto_goTypes,
		DependencyIndexes: file_ProtoLogistic_ProtoLogistic_proto_depIdxs,
		EnumInfos:         file_ProtoLogistic_ProtoLogistic_proto_enumTypes,
		MessageInfos:      file_ProtoLogistic_ProtoLogistic_proto_msgTypes,
	}.Build()
	File_ProtoLogistic_ProtoLogistic_proto = out.File
	file_ProtoLogistic_ProtoLogistic_proto_rawDesc = nil
	file_ProtoLogistic_ProtoLogistic_proto_goTypes = nil
	file_ProtoLogistic_ProtoLogistic_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProtoLogisticServiceClient is the client API for ProtoLogisticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProtoLogisticServiceClient interface {
	DeliverPackage(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Empty, error)
	CheckStatus(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error)
	AskPackage(ctx context.Context, in *Truck, opts ...grpc.CallOption) (*Package, error)
	FinishPackage(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Empty, error)
}

type protoLogisticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProtoLogisticServiceClient(cc grpc.ClientConnInterface) ProtoLogisticServiceClient {
	return &protoLogisticServiceClient{cc}
}

func (c *protoLogisticServiceClient) DeliverPackage(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ProtoLogistic.ProtoLogisticService/DeliverPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoLogisticServiceClient) CheckStatus(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, "/ProtoLogistic.ProtoLogisticService/CheckStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoLogisticServiceClient) AskPackage(ctx context.Context, in *Truck, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, "/ProtoLogistic.ProtoLogisticService/AskPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoLogisticServiceClient) FinishPackage(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ProtoLogistic.ProtoLogisticService/FinishPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtoLogisticServiceServer is the server API for ProtoLogisticService service.
type ProtoLogisticServiceServer interface {
	DeliverPackage(context.Context, *Package) (*Empty, error)
	CheckStatus(context.Context, *Package) (*Package, error)
	AskPackage(context.Context, *Truck) (*Package, error)
	FinishPackage(context.Context, *Package) (*Empty, error)
}

// UnimplementedProtoLogisticServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProtoLogisticServiceServer struct {
}

func (*UnimplementedProtoLogisticServiceServer) DeliverPackage(context.Context, *Package) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverPackage not implemented")
}
func (*UnimplementedProtoLogisticServiceServer) CheckStatus(context.Context, *Package) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}
func (*UnimplementedProtoLogisticServiceServer) AskPackage(context.Context, *Truck) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskPackage not implemented")
}
func (*UnimplementedProtoLogisticServiceServer) FinishPackage(context.Context, *Package) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishPackage not implemented")
}

func RegisterProtoLogisticServiceServer(s *grpc.Server, srv ProtoLogisticServiceServer) {
	s.RegisterService(&_ProtoLogisticService_serviceDesc, srv)
}

func _ProtoLogisticService_DeliverPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoLogisticServiceServer).DeliverPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProtoLogistic.ProtoLogisticService/DeliverPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoLogisticServiceServer).DeliverPackage(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoLogisticService_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoLogisticServiceServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProtoLogistic.ProtoLogisticService/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoLogisticServiceServer).CheckStatus(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoLogisticService_AskPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Truck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoLogisticServiceServer).AskPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProtoLogistic.ProtoLogisticService/AskPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoLogisticServiceServer).AskPackage(ctx, req.(*Truck))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoLogisticService_FinishPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoLogisticServiceServer).FinishPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProtoLogistic.ProtoLogisticService/FinishPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoLogisticServiceServer).FinishPackage(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProtoLogisticService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProtoLogistic.ProtoLogisticService",
	HandlerType: (*ProtoLogisticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeliverPackage",
			Handler:    _ProtoLogisticService_DeliverPackage_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _ProtoLogisticService_CheckStatus_Handler,
		},
		{
			MethodName: "AskPackage",
			Handler:    _ProtoLogisticService_AskPackage_Handler,
		},
		{
			MethodName: "FinishPackage",
			Handler:    _ProtoLogisticService_FinishPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ProtoLogistic/ProtoLogistic.proto",
}