// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/postgresql/v1/database.proto

package postgresql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A PostgreSQL Database resource. For more information, see
// the [Developer's Guide](/docs/managed-postgresql/concepts).
type Database struct {
	// Name of the database.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the PostgreSQL cluster that the database belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the user assigned as the owner of the database.
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// POSIX locale for string sorting order.
	// Can only be set at creation time.
	LcCollate string `protobuf:"bytes,4,opt,name=lc_collate,json=lcCollate,proto3" json:"lc_collate,omitempty"`
	// POSIX locale for character classification.
	// Can only be set at creation time.
	LcCtype string `protobuf:"bytes,5,opt,name=lc_ctype,json=lcCtype,proto3" json:"lc_ctype,omitempty"`
	// PostgreSQL extensions enabled for the database.
	Extensions           []*Extension `protobuf:"bytes,6,rep,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Database) Reset()         { *m = Database{} }
func (m *Database) String() string { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()    {}
func (*Database) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f036bfe45d2a329, []int{0}
}

func (m *Database) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Database.Unmarshal(m, b)
}
func (m *Database) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Database.Marshal(b, m, deterministic)
}
func (m *Database) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Database.Merge(m, src)
}
func (m *Database) XXX_Size() int {
	return xxx_messageInfo_Database.Size(m)
}
func (m *Database) XXX_DiscardUnknown() {
	xxx_messageInfo_Database.DiscardUnknown(m)
}

var xxx_messageInfo_Database proto.InternalMessageInfo

func (m *Database) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Database) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *Database) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Database) GetLcCollate() string {
	if m != nil {
		return m.LcCollate
	}
	return ""
}

func (m *Database) GetLcCtype() string {
	if m != nil {
		return m.LcCtype
	}
	return ""
}

func (m *Database) GetExtensions() []*Extension {
	if m != nil {
		return m.Extensions
	}
	return nil
}

type Extension struct {
	// Name of the extension, e.g. `pg_trgm` or `pg_btree`.
	// Extensions supported by Managed Service for PostgreSQL are [listed in the Developer's Guide](/docs/managed-postgresql/operations/cluster-extensions).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the extension.
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Extension) Reset()         { *m = Extension{} }
func (m *Extension) String() string { return proto.CompactTextString(m) }
func (*Extension) ProtoMessage()    {}
func (*Extension) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f036bfe45d2a329, []int{1}
}

func (m *Extension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extension.Unmarshal(m, b)
}
func (m *Extension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extension.Marshal(b, m, deterministic)
}
func (m *Extension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extension.Merge(m, src)
}
func (m *Extension) XXX_Size() int {
	return xxx_messageInfo_Extension.Size(m)
}
func (m *Extension) XXX_DiscardUnknown() {
	xxx_messageInfo_Extension.DiscardUnknown(m)
}

var xxx_messageInfo_Extension proto.InternalMessageInfo

func (m *Extension) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Extension) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type DatabaseSpec struct {
	// Name of the PostgreSQL database. 1-63 characters long.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of the user to be assigned as the owner of the database.
	// To get the list of available PostgreSQL users, make a [UserService.List] request.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// POSIX locale for string sorting order.
	// Can only be set at creation time.
	LcCollate string `protobuf:"bytes,3,opt,name=lc_collate,json=lcCollate,proto3" json:"lc_collate,omitempty"`
	// POSIX locale for character classification.
	// Can only be set at creation time.
	LcCtype string `protobuf:"bytes,4,opt,name=lc_ctype,json=lcCtype,proto3" json:"lc_ctype,omitempty"`
	// PostgreSQL extensions to be enabled for the database.
	Extensions           []*Extension `protobuf:"bytes,5,rep,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DatabaseSpec) Reset()         { *m = DatabaseSpec{} }
func (m *DatabaseSpec) String() string { return proto.CompactTextString(m) }
func (*DatabaseSpec) ProtoMessage()    {}
func (*DatabaseSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f036bfe45d2a329, []int{2}
}

func (m *DatabaseSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseSpec.Unmarshal(m, b)
}
func (m *DatabaseSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseSpec.Marshal(b, m, deterministic)
}
func (m *DatabaseSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseSpec.Merge(m, src)
}
func (m *DatabaseSpec) XXX_Size() int {
	return xxx_messageInfo_DatabaseSpec.Size(m)
}
func (m *DatabaseSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseSpec proto.InternalMessageInfo

func (m *DatabaseSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatabaseSpec) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DatabaseSpec) GetLcCollate() string {
	if m != nil {
		return m.LcCollate
	}
	return ""
}

func (m *DatabaseSpec) GetLcCtype() string {
	if m != nil {
		return m.LcCtype
	}
	return ""
}

func (m *DatabaseSpec) GetExtensions() []*Extension {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func init() {
	proto.RegisterType((*Database)(nil), "yandex.cloud.mdb.postgresql.v1.Database")
	proto.RegisterType((*Extension)(nil), "yandex.cloud.mdb.postgresql.v1.Extension")
	proto.RegisterType((*DatabaseSpec)(nil), "yandex.cloud.mdb.postgresql.v1.DatabaseSpec")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/postgresql/v1/database.proto", fileDescriptor_5f036bfe45d2a329)
}

var fileDescriptor_5f036bfe45d2a329 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0x87, 0xc9, 0xbd, 0xe9, 0xbd, 0xb7, 0xe3, 0x9f, 0xc5, 0x28, 0x18, 0x85, 0x96, 0xd2, 0x55,
	0x55, 0x66, 0x62, 0x5a, 0x28, 0x16, 0x75, 0x61, 0xaa, 0x42, 0x17, 0x22, 0x54, 0xdd, 0x54, 0x4a,
	0x98, 0xcc, 0x0c, 0x31, 0x30, 0xc9, 0xc4, 0x64, 0x1a, 0x5b, 0xe9, 0x13, 0xf8, 0x30, 0xbe, 0x46,
	0x7d, 0x04, 0x1f, 0xc1, 0x75, 0x9f, 0x40, 0x32, 0x49, 0x6c, 0x0a, 0xda, 0xcd, 0xdd, 0xcd, 0x99,
	0xdf, 0xf9, 0x0e, 0x9c, 0x8f, 0x03, 0xd0, 0x86, 0xc4, 0x8c, 0xaf, 0x6d, 0x2a, 0xe4, 0x8a, 0xd9,
	0x11, 0xf3, 0xed, 0x44, 0x66, 0x2a, 0x48, 0x79, 0xf6, 0x45, 0xd8, 0xb9, 0x63, 0x33, 0xa2, 0x88,
	0x4f, 0x32, 0x8e, 0x93, 0x54, 0x2a, 0x09, 0xbb, 0x65, 0x3b, 0xd6, 0xed, 0x38, 0x62, 0x3e, 0x3e,
	0xb4, 0xe3, 0xdc, 0x79, 0xd0, 0x39, 0x1a, 0x97, 0x13, 0x11, 0x32, 0xa2, 0x42, 0x19, 0x97, 0x78,
	0xff, 0x97, 0x01, 0xae, 0x5e, 0x55, 0x13, 0x21, 0x04, 0x66, 0x4c, 0x22, 0x6e, 0x19, 0x3d, 0x63,
	0xd0, 0x9e, 0xeb, 0x37, 0xec, 0x00, 0x40, 0xc5, 0x2a, 0x53, 0x3c, 0xf5, 0x42, 0x66, 0x9d, 0xe9,
	0xa4, 0x5d, 0xfd, 0xcc, 0x18, 0xbc, 0x0b, 0x5a, 0xf2, 0x6b, 0xcc, 0x53, 0xeb, 0x5c, 0x27, 0x65,
	0x51, 0x40, 0x82, 0x7a, 0x54, 0x0a, 0x41, 0x14, 0xb7, 0xcc, 0x12, 0x12, 0x74, 0x5a, 0x7e, 0xc0,
	0xfb, 0xe0, 0xaa, 0x88, 0xd5, 0x26, 0xe1, 0x56, 0x4b, 0x87, 0x97, 0x82, 0x4e, 0x8b, 0x12, 0xce,
	0x00, 0xe0, 0x6b, 0xc5, 0xe3, 0x2c, 0x94, 0x71, 0x66, 0x5d, 0xf4, 0xce, 0x07, 0x37, 0x86, 0x0f,
	0xf1, 0xe9, 0x1d, 0xf1, 0xeb, 0x9a, 0x98, 0x37, 0xe0, 0xfe, 0x04, 0xb4, 0xff, 0x06, 0xff, 0x5c,
	0xcd, 0x02, 0x97, 0x39, 0x4f, 0x8b, 0xb8, 0xda, 0xab, 0x2e, 0xfb, 0x3f, 0xce, 0xc0, 0xcd, 0xda,
	0xca, 0xfb, 0x84, 0x53, 0x38, 0x6c, 0xe2, 0x6e, 0xf7, 0xf7, 0xce, 0x31, 0xf6, 0x3b, 0xe7, 0xf6,
	0x27, 0x82, 0xbe, 0xbd, 0x44, 0x8b, 0x27, 0x68, 0xe2, 0xa1, 0xe5, 0xa3, 0xef, 0x3f, 0x1d, 0xf3,
	0xf9, 0x8b, 0xf1, 0xa8, 0x1a, 0x3f, 0xaa, 0xd5, 0xe8, 0xe1, 0x6e, 0xa7, 0x82, 0x6e, 0x35, 0xa0,
	0x06, 0x53, 0x99, 0x1b, 0x1f, 0x99, 0xd3, 0x52, 0xdd, 0x7b, 0xfb, 0x9d, 0x73, 0x67, 0x5b, 0x61,
	0xde, 0xf2, 0x31, 0xfe, 0xf8, 0xe1, 0x0d, 0x7a, 0xba, 0x9d, 0x36, 0x95, 0x0e, 0x1b, 0x4a, 0xcd,
	0xd3, 0xd4, 0x7f, 0x5c, 0xb7, 0xae, 0xe1, 0xda, 0x7d, 0xb7, 0x78, 0x1b, 0x84, 0xea, 0xf3, 0xca,
	0xc7, 0x54, 0x46, 0x76, 0x39, 0x02, 0x95, 0x27, 0x17, 0x48, 0x14, 0xf0, 0x58, 0x5f, 0x9b, 0x7d,
	0xfa, 0xb4, 0x9f, 0x1d, 0x2a, 0xff, 0x42, 0x03, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01,
	0x69, 0x06, 0xe9, 0x0e, 0x03, 0x00, 0x00,
}
