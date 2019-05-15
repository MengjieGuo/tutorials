// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/orders/orders.proto

package mu_micro_book_srv_orders

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

type Request struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	BookId               int64    `protobuf:"varint,2,opt,name=bookId,proto3" json:"bookId,omitempty"`
	OrderId              int64    `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e484a55ca159a606, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *Request) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type Order struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	BookId               int64    `protobuf:"varint,3,opt,name=bookId,proto3" json:"bookId,omitempty"`
	InvHistoryId         int64    `protobuf:"varint,4,opt,name=invHistoryId,proto3" json:"invHistoryId,omitempty"`
	State                int64    `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
	CreatedTime          int64    `protobuf:"varint,6,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          int64    `protobuf:"varint,7,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_e484a55ca159a606, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *Order) GetInvHistoryId() int64 {
	if m != nil {
		return m.InvHistoryId
	}
	return 0
}

func (m *Order) GetState() int64 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *Order) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Order) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Order                *Order   `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e484a55ca159a606, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_e484a55ca159a606, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "mu.micro.book.srv.orders.Request")
	proto.RegisterType((*Order)(nil), "mu.micro.book.srv.orders.Order")
	proto.RegisterType((*Response)(nil), "mu.micro.book.srv.orders.Response")
	proto.RegisterType((*Error)(nil), "mu.micro.book.srv.orders.Error")
}

func init() { proto.RegisterFile("proto/orders/orders.proto", fileDescriptor_e484a55ca159a606) }

var fileDescriptor_e484a55ca159a606 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x4e, 0xeb, 0x30,
	0x10, 0x7e, 0x49, 0x9a, 0xb4, 0x6f, 0x8a, 0x58, 0x58, 0x08, 0x19, 0x36, 0x14, 0xaf, 0x58, 0x19,
	0xa9, 0x15, 0x47, 0x40, 0x90, 0x4d, 0x91, 0x5c, 0x2e, 0xd0, 0xc6, 0xb3, 0xb0, 0xa0, 0x75, 0xb1,
	0x9d, 0x22, 0xae, 0xc1, 0x3d, 0x38, 0x06, 0xf7, 0x42, 0x1e, 0xa7, 0xa8, 0x5d, 0xb4, 0x2b, 0x56,
	0xf1, 0xf7, 0x33, 0xdf, 0x78, 0x26, 0x86, 0x8b, 0xb5, 0xb3, 0xc1, 0xde, 0x5a, 0xa7, 0xd1, 0xf9,
	0xee, 0x23, 0x89, 0x63, 0x7c, 0xd9, 0xca, 0xa5, 0x69, 0x9c, 0x95, 0x0b, 0x6b, 0x5f, 0xa4, 0x77,
	0x1b, 0x99, 0x74, 0x31, 0x83, 0xbe, 0xc2, 0xb7, 0x16, 0x7d, 0x60, 0xe7, 0x50, 0xb5, 0x1e, 0x5d,
	0xad, 0x79, 0x36, 0xca, 0x6e, 0x0a, 0xd5, 0xa1, 0xc8, 0xc7, 0xaa, 0x5a, 0xf3, 0x3c, 0xf1, 0x09,
	0x31, 0x0e, 0x7d, 0x0a, 0xa9, 0x35, 0x2f, 0x48, 0xd8, 0x42, 0xf1, 0x9d, 0x41, 0xf9, 0x14, 0xcf,
	0xec, 0x14, 0x72, 0xb3, 0xcd, 0xcb, 0x8d, 0xde, 0xe9, 0x91, 0x1f, 0xe8, 0x51, 0xec, 0xf5, 0x10,
	0x70, 0x62, 0x56, 0x9b, 0x47, 0xe3, 0x83, 0x75, 0x1f, 0xb5, 0xe6, 0x3d, 0x52, 0xf7, 0x38, 0x76,
	0x06, 0xa5, 0x0f, 0xf3, 0x80, 0xbc, 0x24, 0x31, 0x01, 0x36, 0x82, 0x61, 0xe3, 0x70, 0x1e, 0x50,
	0x3f, 0x9b, 0x25, 0xf2, 0x8a, 0xb4, 0x5d, 0x2a, 0x3a, 0xda, 0xb5, 0xfe, 0x75, 0xf4, 0x93, 0x63,
	0x87, 0x12, 0x9f, 0x19, 0x0c, 0x14, 0xfa, 0xb5, 0x5d, 0x79, 0x8c, 0xe3, 0xfa, 0xb6, 0x69, 0xd0,
	0x7b, 0x9a, 0x67, 0xa0, 0xb6, 0x90, 0xdd, 0x41, 0x89, 0xce, 0x59, 0x47, 0x33, 0x0d, 0xc7, 0x57,
	0xf2, 0xd0, 0xb6, 0xe5, 0x7d, 0xb4, 0xa9, 0xe4, 0x8e, 0x65, 0x44, 0xd3, 0xc8, 0x47, 0xcb, 0x68,
	0x97, 0x2a, 0xb9, 0xc5, 0x04, 0x4a, 0x8a, 0x61, 0x0c, 0x7a, 0x8d, 0xd5, 0x48, 0xb7, 0x29, 0x15,
	0x9d, 0xe3, 0x1e, 0x35, 0x86, 0xb9, 0x79, 0xa5, 0xbb, 0xfc, 0x57, 0x1d, 0x1a, 0x7f, 0x65, 0x50,
	0x51, 0x8a, 0x67, 0x53, 0x28, 0xa6, 0xf8, 0xce, 0xae, 0x0f, 0xb7, 0xeb, 0x1e, 0xc4, 0xa5, 0x38,
	0x66, 0x49, 0x5b, 0x11, 0xff, 0xd8, 0x0c, 0x06, 0x0f, 0x18, 0xd2, 0xef, 0xfe, 0xab, 0xd0, 0x45,
	0x45, 0xef, 0x76, 0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0x55, 0x9e, 0x67, 0x76, 0xd4, 0x02, 0x00,
	0x00,
}
