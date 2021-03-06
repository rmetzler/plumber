// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/http.proto

package records

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HTTPRecord struct {
	Method               string            `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Headers              map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Path                 string            `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Source               string            `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Body                 []byte            `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Timestamp            int64             `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HTTPRecord) Reset()         { *m = HTTPRecord{} }
func (m *HTTPRecord) String() string { return proto.CompactTextString(m) }
func (*HTTPRecord) ProtoMessage()    {}
func (*HTTPRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_655362cba3b74fb9, []int{0}
}

func (m *HTTPRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPRecord.Unmarshal(m, b)
}
func (m *HTTPRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPRecord.Marshal(b, m, deterministic)
}
func (m *HTTPRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPRecord.Merge(m, src)
}
func (m *HTTPRecord) XXX_Size() int {
	return xxx_messageInfo_HTTPRecord.Size(m)
}
func (m *HTTPRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPRecord.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPRecord proto.InternalMessageInfo

func (m *HTTPRecord) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HTTPRecord) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPRecord) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HTTPRecord) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *HTTPRecord) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *HTTPRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*HTTPRecord)(nil), "records.HTTPRecord")
	proto.RegisterMapType((map[string]string)(nil), "records.HTTPRecord.HeadersEntry")
}

func init() { proto.RegisterFile("records/http.proto", fileDescriptor_655362cba3b74fb9) }

var fileDescriptor_655362cba3b74fb9 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4b, 0xf4, 0x30,
	0x10, 0xc5, 0x49, 0xbb, 0xdb, 0x65, 0xe7, 0xdb, 0xc3, 0x47, 0x10, 0x09, 0xe2, 0xa1, 0x78, 0xea,
	0xc5, 0x06, 0x54, 0x44, 0x7a, 0x14, 0x84, 0x3d, 0x4a, 0xd9, 0x93, 0xb7, 0xa4, 0x1d, 0xcc, 0xe2,
	0xd6, 0x84, 0x64, 0x2a, 0xf4, 0x5f, 0xf7, 0x24, 0x4d, 0x23, 0xeb, 0xed, 0xfd, 0x1e, 0x33, 0xef,
	0x25, 0x03, 0xdc, 0x63, 0x67, 0x7d, 0x1f, 0xa4, 0x21, 0x72, 0xb5, 0xf3, 0x96, 0x2c, 0xdf, 0x24,
	0xef, 0xe6, 0x9b, 0x01, 0xec, 0x0f, 0x87, 0xd7, 0x36, 0x32, 0xbf, 0x84, 0x62, 0x40, 0x32, 0xb6,
	0x17, 0xac, 0x64, 0xd5, 0xb6, 0x4d, 0xc4, 0x1b, 0xd8, 0x18, 0x54, 0x3d, 0xfa, 0x20, 0xb2, 0x32,
	0xaf, 0xfe, 0xdd, 0x95, 0x75, 0x4a, 0xa8, 0xcf, 0xdb, 0xf5, 0x7e, 0x19, 0x79, 0xf9, 0x24, 0x3f,
	0xb5, 0xbf, 0x0b, 0x9c, 0xc3, 0xca, 0x29, 0x32, 0x22, 0x8f, 0x89, 0x51, 0xcf, 0x3d, 0xc1, 0x8e,
	0xbe, 0x43, 0xb1, 0x5a, 0x7a, 0x16, 0x9a, 0x67, 0xb5, 0xed, 0x27, 0xb1, 0x2e, 0x59, 0xb5, 0x6b,
	0xa3, 0xe6, 0xd7, 0xb0, 0xa5, 0xe3, 0x80, 0x81, 0xd4, 0xe0, 0x44, 0x51, 0xb2, 0x2a, 0x6f, 0xcf,
	0xc6, 0x55, 0x03, 0xbb, 0xbf, 0xb5, 0xfc, 0x3f, 0xe4, 0x1f, 0x38, 0xa5, 0xe7, 0xcf, 0x92, 0x5f,
	0xc0, 0xfa, 0x4b, 0x9d, 0x46, 0x14, 0x59, 0xf4, 0x16, 0x68, 0xb2, 0x27, 0xf6, 0xfc, 0xf8, 0xf6,
	0xf0, 0x7e, 0x24, 0x33, 0xea, 0xba, 0xb3, 0x83, 0xd4, 0x8a, 0x3a, 0xd3, 0x59, 0xef, 0xa4, 0x3b,
	0x8d, 0x83, 0x46, 0x2f, 0x09, 0x03, 0xdd, 0xaa, 0x10, 0x90, 0x82, 0x74, 0x3a, 0xc8, 0xf4, 0x65,
	0x5d, 0xc4, 0x23, 0xde, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x3d, 0x50, 0xbb, 0x5a, 0x01,
	0x00, 0x00,
}
