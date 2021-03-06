// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bq_field.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Message containing options related to BigQuery schema generation
// and management via Protobuf.
type BigQueryFieldOptions struct {
	// Flag to specify that a field should be marked as 'REQUIRED' when
	// used to generate schema for BigQuery.
	Require bool `protobuf:"varint,1,opt,name=require" json:"require,omitempty"`
	// Optionally override whatever type is resolved by the schema
	// generator. This is useful, for example, to store epoch timestamps
	// with the underlying 'TIMESTAMP' type, when normally, they would
	// be structured as 'INTEGER' fields.
	TypeOverride string `protobuf:"bytes,2,opt,name=type_override,json=typeOverride" json:"type_override,omitempty"`
	// Optionally omit a field from BigQuery schema.
	Ignore bool `protobuf:"varint,3,opt,name=ignore" json:"ignore,omitempty"`
	// Set the description for a field in BigQuery schema.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
}

func (m *BigQueryFieldOptions) Reset()                    { *m = BigQueryFieldOptions{} }
func (m *BigQueryFieldOptions) String() string            { return proto.CompactTextString(m) }
func (*BigQueryFieldOptions) ProtoMessage()               {}
func (*BigQueryFieldOptions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *BigQueryFieldOptions) GetRequire() bool {
	if m != nil {
		return m.Require
	}
	return false
}

func (m *BigQueryFieldOptions) GetTypeOverride() string {
	if m != nil {
		return m.TypeOverride
	}
	return ""
}

func (m *BigQueryFieldOptions) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

func (m *BigQueryFieldOptions) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

var E_Bigquery = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*BigQueryFieldOptions)(nil),
	Field:         1021,
	Name:          "gen_bq_schema.bigquery",
	Tag:           "bytes,1021,opt,name=bigquery",
	Filename:      "bq_field.proto",
}

func init() {
	proto.RegisterType((*BigQueryFieldOptions)(nil), "gen_bq_schema.BigQueryFieldOptions")
	proto.RegisterExtension(E_Bigquery)
}

func init() { proto.RegisterFile("bq_field.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x59, 0x95, 0x36, 0x6e, 0xad, 0x87, 0x45, 0x64, 0x11, 0x84, 0xc5, 0x5e, 0x72, 0xda,
	0x82, 0xde, 0x3c, 0xf6, 0xe0, 0xb5, 0x98, 0xa3, 0x97, 0xc5, 0x6d, 0xa6, 0xeb, 0x40, 0xcd, 0x24,
	0x93, 0x8d, 0xd0, 0x5f, 0xe1, 0x2f, 0x16, 0x24, 0x9b, 0x46, 0x2a, 0xf4, 0x34, 0xbc, 0x37, 0xbc,
	0xc7, 0xfb, 0xe4, 0xb5, 0x6f, 0xdc, 0x16, 0x61, 0x57, 0xda, 0x9a, 0x29, 0x92, 0x9a, 0x07, 0xa8,
	0x9c, 0x6f, 0x5c, 0xbb, 0xf9, 0x80, 0xcf, 0xf7, 0x3b, 0x13, 0x88, 0xc2, 0x0e, 0x96, 0xe9, 0xe9,
	0xbb, 0xed, 0xb2, 0x84, 0x76, 0xc3, 0x58, 0x47, 0xe2, 0x21, 0xf0, 0xf0, 0x2d, 0xe4, 0xcd, 0x0a,
	0xc3, 0x6b, 0x07, 0xbc, 0x7f, 0xe9, 0x8b, 0xd6, 0x75, 0x44, 0xaa, 0x5a, 0xa5, 0xe5, 0x94, 0xa1,
	0xe9, 0x90, 0x41, 0x0b, 0x23, 0xf2, 0xac, 0x18, 0xa5, 0x5a, 0xc8, 0x79, 0xdc, 0xd7, 0xe0, 0xe8,
	0x0b, 0x98, 0xb1, 0x04, 0x7d, 0x66, 0x44, 0x7e, 0x59, 0x5c, 0xf5, 0xe6, 0xfa, 0xe0, 0xa9, 0x5b,
	0x39, 0xc1, 0x50, 0x11, 0x83, 0x3e, 0x4f, 0xe9, 0x83, 0x52, 0x46, 0xce, 0xc6, 0x0d, 0x48, 0x95,
	0xbe, 0x48, 0xd1, 0x63, 0xeb, 0xd9, 0xc9, 0xcc, 0x63, 0x68, 0xfa, 0x41, 0xea, 0xde, 0x0e, 0x00,
	0x76, 0x04, 0xb0, 0xc7, 0x1b, 0xf5, 0xcf, 0xd4, 0x88, 0x7c, 0xf6, 0xb8, 0xb0, 0xff, 0xa8, 0xed,
	0x29, 0x9e, 0xe2, 0xaf, 0x74, 0x95, 0xbd, 0x4d, 0x52, 0x5d, 0xeb, 0x87, 0xfb, 0xf4, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0x9a, 0xe0, 0x3e, 0x46, 0x01, 0x00, 0x00,
}
