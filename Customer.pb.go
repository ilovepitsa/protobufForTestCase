// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: Customer.proto

package protobuf

import (
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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_Customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_Customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RequestAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int32  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RequestAdd) Reset() {
	*x = RequestAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAdd) ProtoMessage() {}

func (x *RequestAdd) ProtoReflect() protoreflect.Message {
	mi := &file_Customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAdd.ProtoReflect.Descriptor instead.
func (*RequestAdd) Descriptor() ([]byte, []int) {
	return file_Customer_proto_rawDescGZIP(), []int{1}
}

func (x *RequestAdd) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *RequestAdd) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResponceAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int32 `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	IsSuccess bool  `protobuf:"varint,2,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
}

func (x *ResponceAdd) Reset() {
	*x = ResponceAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceAdd) ProtoMessage() {}

func (x *ResponceAdd) ProtoReflect() protoreflect.Message {
	mi := &file_Customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceAdd.ProtoReflect.Descriptor instead.
func (*ResponceAdd) Descriptor() ([]byte, []int) {
	return file_Customer_proto_rawDescGZIP(), []int{2}
}

func (x *ResponceAdd) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *ResponceAdd) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type RequestGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId  int32 `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	CustomerId int32 `protobuf:"varint,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
}

func (x *RequestGet) Reset() {
	*x = RequestGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGet) ProtoMessage() {}

func (x *RequestGet) ProtoReflect() protoreflect.Message {
	mi := &file_Customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGet.ProtoReflect.Descriptor instead.
func (*RequestGet) Descriptor() ([]byte, []int) {
	return file_Customer_proto_rawDescGZIP(), []int{3}
}

func (x *RequestGet) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *RequestGet) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type ResponceGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int32  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ResponceGet) Reset() {
	*x = ResponceGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceGet) ProtoMessage() {}

func (x *ResponceGet) ProtoReflect() protoreflect.Message {
	mi := &file_Customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceGet.ProtoReflect.Descriptor instead.
func (*ResponceGet) Descriptor() ([]byte, []int) {
	return file_Customer_proto_rawDescGZIP(), []int{4}
}

func (x *ResponceGet) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *ResponceGet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RequestGetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int32 `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *RequestGetAll) Reset() {
	*x = RequestGetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGetAll) ProtoMessage() {}

func (x *RequestGetAll) ProtoReflect() protoreflect.Message {
	mi := &file_Customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGetAll.ProtoReflect.Descriptor instead.
func (*RequestGetAll) Descriptor() ([]byte, []int) {
	return file_Customer_proto_rawDescGZIP(), []int{5}
}

func (x *RequestGetAll) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

type ResponceGetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customers []*Customer `protobuf:"bytes,2,rep,name=customers,proto3" json:"customers,omitempty"`
}

func (x *ResponceGetAll) Reset() {
	*x = ResponceGetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceGetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceGetAll) ProtoMessage() {}

func (x *ResponceGetAll) ProtoReflect() protoreflect.Message {
	mi := &file_Customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceGetAll.ProtoReflect.Descriptor instead.
func (*ResponceGetAll) Descriptor() ([]byte, []int) {
	return file_Customer_proto_rawDescGZIP(), []int{6}
}

func (x *ResponceGetAll) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Req:
	//	*Request_ReqAdd
	//	*Request_ReqGet
	//	*Request_ReqGetAll
	Req isRequest_Req `protobuf_oneof:"req"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Customer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Customer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Customer_proto_rawDescGZIP(), []int{7}
}

func (m *Request) GetReq() isRequest_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (x *Request) GetReqAdd() *RequestAdd {
	if x, ok := x.GetReq().(*Request_ReqAdd); ok {
		return x.ReqAdd
	}
	return nil
}

func (x *Request) GetReqGet() *RequestGet {
	if x, ok := x.GetReq().(*Request_ReqGet); ok {
		return x.ReqGet
	}
	return nil
}

func (x *Request) GetReqGetAll() *RequestGetAll {
	if x, ok := x.GetReq().(*Request_ReqGetAll); ok {
		return x.ReqGetAll
	}
	return nil
}

type isRequest_Req interface {
	isRequest_Req()
}

type Request_ReqAdd struct {
	ReqAdd *RequestAdd `protobuf:"bytes,1,opt,name=reqAdd,proto3,oneof"`
}

type Request_ReqGet struct {
	ReqGet *RequestGet `protobuf:"bytes,2,opt,name=reqGet,proto3,oneof"`
}

type Request_ReqGetAll struct {
	ReqGetAll *RequestGetAll `protobuf:"bytes,3,opt,name=reqGetAll,proto3,oneof"`
}

func (*Request_ReqAdd) isRequest_Req() {}

func (*Request_ReqGet) isRequest_Req() {}

func (*Request_ReqGetAll) isRequest_Req() {}

type Responce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resp:
	//	*Responce_RespAdd
	//	*Responce_RespGet
	//	*Responce_RespGetAll
	Resp isResponce_Resp `protobuf_oneof:"resp"`
}

func (x *Responce) Reset() {
	*x = Responce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Customer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Responce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Responce) ProtoMessage() {}

func (x *Responce) ProtoReflect() protoreflect.Message {
	mi := &file_Customer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Responce.ProtoReflect.Descriptor instead.
func (*Responce) Descriptor() ([]byte, []int) {
	return file_Customer_proto_rawDescGZIP(), []int{8}
}

func (m *Responce) GetResp() isResponce_Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (x *Responce) GetRespAdd() *ResponceAdd {
	if x, ok := x.GetResp().(*Responce_RespAdd); ok {
		return x.RespAdd
	}
	return nil
}

func (x *Responce) GetRespGet() *ResponceGet {
	if x, ok := x.GetResp().(*Responce_RespGet); ok {
		return x.RespGet
	}
	return nil
}

func (x *Responce) GetRespGetAll() *ResponceGetAll {
	if x, ok := x.GetResp().(*Responce_RespGetAll); ok {
		return x.RespGetAll
	}
	return nil
}

type isResponce_Resp interface {
	isResponce_Resp()
}

type Responce_RespAdd struct {
	RespAdd *ResponceAdd `protobuf:"bytes,1,opt,name=respAdd,proto3,oneof"`
}

type Responce_RespGet struct {
	RespGet *ResponceGet `protobuf:"bytes,2,opt,name=respGet,proto3,oneof"`
}

type Responce_RespGetAll struct {
	RespGetAll *ResponceGetAll `protobuf:"bytes,3,opt,name=respGetAll,proto3,oneof"`
}

func (*Responce_RespAdd) isResponce_Resp() {}

func (*Responce_RespGet) isResponce_Resp() {}

func (*Responce_RespGetAll) isResponce_Resp() {}

var File_Customer_proto protoreflect.FileDescriptor

var file_Customer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x2e, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x49, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x41, 0x64,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4a, 0x0a,
	0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x0d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x41, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x48,
	0x00, 0x52, 0x06, 0x72, 0x65, 0x71, 0x41, 0x64, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x72, 0x65, 0x71,
	0x47, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x71, 0x47, 0x65,
	0x74, 0x12, 0x3f, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x71, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x42, 0x05, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x22, 0xcc, 0x01, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x70, 0x41, 0x64,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x41, 0x64, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x73, 0x70, 0x41, 0x64,
	0x64, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x70, 0x47, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x47, 0x65,
	0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x73, 0x70, 0x47, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x0a,
	0x72, 0x65, 0x73, 0x70, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x06, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x70, 0x69, 0x74, 0x73,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Customer_proto_rawDescOnce sync.Once
	file_Customer_proto_rawDescData = file_Customer_proto_rawDesc
)

func file_Customer_proto_rawDescGZIP() []byte {
	file_Customer_proto_rawDescOnce.Do(func() {
		file_Customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_Customer_proto_rawDescData)
	})
	return file_Customer_proto_rawDescData
}

var file_Customer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_Customer_proto_goTypes = []interface{}{
	(*Customer)(nil),       // 0: Customer_service.Customer
	(*RequestAdd)(nil),     // 1: Customer_service.RequestAdd
	(*ResponceAdd)(nil),    // 2: Customer_service.ResponceAdd
	(*RequestGet)(nil),     // 3: Customer_service.RequestGet
	(*ResponceGet)(nil),    // 4: Customer_service.ResponceGet
	(*RequestGetAll)(nil),  // 5: Customer_service.RequestGetAll
	(*ResponceGetAll)(nil), // 6: Customer_service.ResponceGetAll
	(*Request)(nil),        // 7: Customer_service.Request
	(*Responce)(nil),       // 8: Customer_service.Responce
}
var file_Customer_proto_depIdxs = []int32{
	0, // 0: Customer_service.ResponceGetAll.customers:type_name -> Customer_service.Customer
	1, // 1: Customer_service.Request.reqAdd:type_name -> Customer_service.RequestAdd
	3, // 2: Customer_service.Request.reqGet:type_name -> Customer_service.RequestGet
	5, // 3: Customer_service.Request.reqGetAll:type_name -> Customer_service.RequestGetAll
	2, // 4: Customer_service.Responce.respAdd:type_name -> Customer_service.ResponceAdd
	4, // 5: Customer_service.Responce.respGet:type_name -> Customer_service.ResponceGet
	6, // 6: Customer_service.Responce.respGetAll:type_name -> Customer_service.ResponceGetAll
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_Customer_proto_init() }
func file_Customer_proto_init() {
	if File_Customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_Customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAdd); i {
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
		file_Customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceAdd); i {
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
		file_Customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGet); i {
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
		file_Customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceGet); i {
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
		file_Customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGetAll); i {
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
		file_Customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceGetAll); i {
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
		file_Customer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Customer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Responce); i {
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
	file_Customer_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*Request_ReqAdd)(nil),
		(*Request_ReqGet)(nil),
		(*Request_ReqGetAll)(nil),
	}
	file_Customer_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Responce_RespAdd)(nil),
		(*Responce_RespGet)(nil),
		(*Responce_RespGetAll)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Customer_proto_goTypes,
		DependencyIndexes: file_Customer_proto_depIdxs,
		MessageInfos:      file_Customer_proto_msgTypes,
	}.Build()
	File_Customer_proto = out.File
	file_Customer_proto_rawDesc = nil
	file_Customer_proto_goTypes = nil
	file_Customer_proto_depIdxs = nil
}
