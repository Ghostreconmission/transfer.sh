// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/customer_feed_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible customer feed errors.
type CustomerFeedErrorEnum_CustomerFeedError int32

const (
	// Enum unspecified.
	CustomerFeedErrorEnum_UNSPECIFIED CustomerFeedErrorEnum_CustomerFeedError = 0
	// The received error code is not known in this version.
	CustomerFeedErrorEnum_UNKNOWN CustomerFeedErrorEnum_CustomerFeedError = 1
	// An active feed already exists for this customer and place holder type.
	CustomerFeedErrorEnum_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 2
	// The specified feed is removed.
	CustomerFeedErrorEnum_CANNOT_CREATE_FOR_REMOVED_FEED CustomerFeedErrorEnum_CustomerFeedError = 3
	// The CustomerFeed already exists. Update should be used to modify the
	// existing CustomerFeed.
	CustomerFeedErrorEnum_CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 4
	// Cannot update removed customer feed.
	CustomerFeedErrorEnum_CANNOT_MODIFY_REMOVED_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 5
	// Invalid placeholder type.
	CustomerFeedErrorEnum_INVALID_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 6
	// Feed mapping for this placeholder type does not exist.
	CustomerFeedErrorEnum_MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 7
	// Placeholder not allowed at the account level.
	CustomerFeedErrorEnum_PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 8
)

var CustomerFeedErrorEnum_CustomerFeedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	3: "CANNOT_CREATE_FOR_REMOVED_FEED",
	4: "CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED",
	5: "CANNOT_MODIFY_REMOVED_CUSTOMER_FEED",
	6: "INVALID_PLACEHOLDER_TYPE",
	7: "MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE",
	8: "PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED",
}
var CustomerFeedErrorEnum_CustomerFeedError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":      2,
	"CANNOT_CREATE_FOR_REMOVED_FEED":                3,
	"CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED":  4,
	"CANNOT_MODIFY_REMOVED_CUSTOMER_FEED":           5,
	"INVALID_PLACEHOLDER_TYPE":                      6,
	"MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE":      7,
	"PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED": 8,
}

func (x CustomerFeedErrorEnum_CustomerFeedError) String() string {
	return proto.EnumName(CustomerFeedErrorEnum_CustomerFeedError_name, int32(x))
}
func (CustomerFeedErrorEnum_CustomerFeedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_customer_feed_error_4a29e8c44fce075d, []int{0, 0}
}

// Container for enum describing possible customer feed errors.
type CustomerFeedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerFeedErrorEnum) Reset()         { *m = CustomerFeedErrorEnum{} }
func (m *CustomerFeedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerFeedErrorEnum) ProtoMessage()    {}
func (*CustomerFeedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_feed_error_4a29e8c44fce075d, []int{0}
}
func (m *CustomerFeedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerFeedErrorEnum.Unmarshal(m, b)
}
func (m *CustomerFeedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerFeedErrorEnum.Marshal(b, m, deterministic)
}
func (dst *CustomerFeedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerFeedErrorEnum.Merge(dst, src)
}
func (m *CustomerFeedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerFeedErrorEnum.Size(m)
}
func (m *CustomerFeedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerFeedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerFeedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CustomerFeedErrorEnum)(nil), "google.ads.googleads.v1.errors.CustomerFeedErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.CustomerFeedErrorEnum_CustomerFeedError", CustomerFeedErrorEnum_CustomerFeedError_name, CustomerFeedErrorEnum_CustomerFeedError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/customer_feed_error.proto", fileDescriptor_customer_feed_error_4a29e8c44fce075d)
}

var fileDescriptor_customer_feed_error_4a29e8c44fce075d = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x69, 0x06, 0x1b, 0xf2, 0x0e, 0x14, 0x4b, 0x20, 0x84, 0xa6, 0x1e, 0xc2, 0x01, 0x0e,
	0xc3, 0x21, 0xe2, 0x82, 0xc2, 0xc9, 0x8b, 0x9d, 0x12, 0x91, 0xd8, 0x51, 0x92, 0x66, 0x14, 0x55,
	0xb2, 0xc2, 0x12, 0xa2, 0x4a, 0x6b, 0x5c, 0xc5, 0xdd, 0x3e, 0x10, 0x47, 0x3e, 0x0a, 0xdf, 0x81,
	0x2f, 0xc0, 0x8d, 0x13, 0x57, 0xe4, 0xb8, 0xad, 0xd4, 0x16, 0x76, 0xca, 0x5f, 0x2f, 0xbf, 0xff,
	0xff, 0x3d, 0xeb, 0x3d, 0xf0, 0xae, 0x91, 0xb2, 0xb9, 0xae, 0x9d, 0xb2, 0x52, 0x8e, 0x91, 0x5a,
	0xdd, 0xba, 0x4e, 0xdd, 0x75, 0xb2, 0x53, 0xce, 0xd5, 0x8d, 0x5a, 0xc9, 0x45, 0xdd, 0x89, 0xaf,
	0x75, 0x5d, 0x89, 0xbe, 0x88, 0x96, 0x9d, 0x5c, 0x49, 0x38, 0x32, 0x38, 0x2a, 0x2b, 0x85, 0xb6,
	0x4e, 0x74, 0xeb, 0x22, 0xe3, 0x7c, 0x7e, 0xb6, 0x49, 0x5e, 0xce, 0x9d, 0xb2, 0x6d, 0xe5, 0xaa,
	0x5c, 0xcd, 0x65, 0xab, 0x8c, 0xdb, 0xfe, 0x63, 0x81, 0x27, 0xfe, 0x3a, 0x3b, 0xa8, 0xeb, 0x8a,
	0x6a, 0x13, 0x6d, 0x6f, 0x16, 0xf6, 0x4f, 0x0b, 0x3c, 0x3e, 0xf8, 0x03, 0x1f, 0x81, 0xd3, 0x09,
	0xcb, 0x12, 0xea, 0x87, 0x41, 0x48, 0xc9, 0xf0, 0x1e, 0x3c, 0x05, 0x27, 0x13, 0xf6, 0x91, 0xf1,
	0x4b, 0x36, 0x1c, 0xc0, 0x73, 0xf0, 0x2a, 0xa0, 0x94, 0x08, 0x1c, 0xa5, 0x14, 0x93, 0xa9, 0xa0,
	0x9f, 0xc2, 0x2c, 0xcf, 0x44, 0xc0, 0x53, 0x91, 0x44, 0xd8, 0xa7, 0x1f, 0x78, 0x44, 0x68, 0x2a,
	0xf2, 0x69, 0x42, 0x87, 0x16, 0xb4, 0xc1, 0xc8, 0xc7, 0x8c, 0xf1, 0x5c, 0xf8, 0x29, 0xc5, 0x39,
	0xed, 0xb9, 0x94, 0xc6, 0xbc, 0xa0, 0x44, 0xe8, 0x9c, 0xe1, 0x11, 0x7c, 0x03, 0xce, 0x77, 0x99,
	0x9d, 0xe8, 0x90, 0x8d, 0x85, 0x3f, 0xc9, 0x72, 0x1e, 0xd3, 0xd4, 0x38, 0xee, 0xc3, 0x97, 0xe0,
	0xc5, 0xda, 0x11, 0x73, 0x12, 0x06, 0xd3, 0x6d, 0xe2, 0x2e, 0xf8, 0x00, 0x9e, 0x81, 0x67, 0x21,
	0x2b, 0x70, 0x14, 0x92, 0xc3, 0xe1, 0x8e, 0xf5, 0x53, 0xe2, 0x30, 0xcb, 0x74, 0x07, 0xcd, 0xc7,
	0x38, 0x49, 0x7a, 0xfd, 0xaf, 0xa7, 0x9c, 0x40, 0x17, 0xbc, 0xde, 0xaf, 0x0a, 0x3d, 0x02, 0x8e,
	0x22, 0x7e, 0x49, 0x89, 0xe0, 0x6c, 0xaf, 0xfd, 0xc3, 0x8b, 0xdf, 0x03, 0x60, 0x5f, 0xc9, 0x05,
	0xba, 0x7b, 0x7d, 0x17, 0x4f, 0x0f, 0x76, 0x90, 0xe8, 0xc5, 0x25, 0x83, 0xcf, 0x64, 0xed, 0x6c,
	0xe4, 0x75, 0xd9, 0x36, 0x48, 0x76, 0x8d, 0xd3, 0xd4, 0x6d, 0xbf, 0xd6, 0xcd, 0x09, 0x2d, 0xe7,
	0xea, 0x7f, 0x17, 0xf5, 0xde, 0x7c, 0xbe, 0x59, 0x47, 0x63, 0x8c, 0xbf, 0x5b, 0xa3, 0xb1, 0x09,
	0xc3, 0x95, 0x42, 0x46, 0x6a, 0x55, 0xb8, 0xa8, 0x6f, 0xa9, 0x7e, 0x6c, 0x80, 0x19, 0xae, 0xd4,
	0x6c, 0x0b, 0xcc, 0x0a, 0x77, 0x66, 0x80, 0x5f, 0x96, 0x6d, 0xaa, 0x9e, 0x87, 0x2b, 0xe5, 0x79,
	0x5b, 0xc4, 0xf3, 0x0a, 0xd7, 0xf3, 0x0c, 0xf4, 0xe5, 0xb8, 0x9f, 0xee, 0xed, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x64, 0xe1, 0x58, 0x11, 0xee, 0x02, 0x00, 0x00,
}
