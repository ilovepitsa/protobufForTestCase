// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: Account.proto

package protobufForTestCase

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

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number      string  `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	UsdtBalance float64 `protobuf:"fixed64,2,opt,name=usdtBalance,proto3" json:"usdtBalance,omitempty"`
	UsdBalance  float64 `protobuf:"fixed64,3,opt,name=usdBalance,proto3" json:"usdBalance,omitempty"`
	RubBalance  float64 `protobuf:"fixed64,4,opt,name=rubBalance,proto3" json:"rubBalance,omitempty"`
	EurBalance  float64 `protobuf:"fixed64,5,opt,name=eurBalance,proto3" json:"eurBalance,omitempty"`
	BtcBalance  float64 `protobuf:"fixed64,6,opt,name=btcBalance,proto3" json:"btcBalance,omitempty"`
	Status      bool    `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Invoice) GetUsdtBalance() float64 {
	if x != nil {
		return x.UsdtBalance
	}
	return 0
}

func (x *Invoice) GetUsdBalance() float64 {
	if x != nil {
		return x.UsdBalance
	}
	return 0
}

func (x *Invoice) GetRubBalance() float64 {
	if x != nil {
		return x.RubBalance
	}
	return 0
}

func (x *Invoice) GetEurBalance() float64 {
	if x != nil {
		return x.EurBalance
	}
	return 0
}

func (x *Invoice) GetBtcBalance() float64 {
	if x != nil {
		return x.BtcBalance
	}
	return 0
}

func (x *Invoice) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type ResponceActiveBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int32      `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Invoices  []*Invoice `protobuf:"bytes,2,rep,name=invoices,proto3" json:"invoices,omitempty"`
}

func (x *ResponceActiveBalance) Reset() {
	*x = ResponceActiveBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceActiveBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceActiveBalance) ProtoMessage() {}

func (x *ResponceActiveBalance) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceActiveBalance.ProtoReflect.Descriptor instead.
func (*ResponceActiveBalance) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{1}
}

func (x *ResponceActiveBalance) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *ResponceActiveBalance) GetInvoices() []*Invoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

type ResponceFrozenBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int32      `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Invoices  []*Invoice `protobuf:"bytes,2,rep,name=invoices,proto3" json:"invoices,omitempty"`
}

func (x *ResponceFrozenBalance) Reset() {
	*x = ResponceFrozenBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceFrozenBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceFrozenBalance) ProtoMessage() {}

func (x *ResponceFrozenBalance) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceFrozenBalance.ProtoReflect.Descriptor instead.
func (*ResponceFrozenBalance) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{2}
}

func (x *ResponceFrozenBalance) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *ResponceFrozenBalance) GetInvoices() []*Invoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

type ResponceInvoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransId      int32  `protobuf:"varint,1,opt,name=TransId,proto3" json:"TransId,omitempty"`
	Success      bool   `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
}

func (x *ResponceInvoice) Reset() {
	*x = ResponceInvoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceInvoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceInvoice) ProtoMessage() {}

func (x *ResponceInvoice) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceInvoice.ProtoReflect.Descriptor instead.
func (*ResponceInvoice) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{3}
}

func (x *ResponceInvoice) GetTransId() int32 {
	if x != nil {
		return x.TransId
	}
	return 0
}

func (x *ResponceInvoice) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ResponceInvoice) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type ResponceWithdraw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransId      int32  `protobuf:"varint,1,opt,name=TransId,proto3" json:"TransId,omitempty"`
	Success      bool   `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
}

func (x *ResponceWithdraw) Reset() {
	*x = ResponceWithdraw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceWithdraw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceWithdraw) ProtoMessage() {}

func (x *ResponceWithdraw) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceWithdraw.ProtoReflect.Descriptor instead.
func (*ResponceWithdraw) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{4}
}

func (x *ResponceWithdraw) GetTransId() int32 {
	if x != nil {
		return x.TransId
	}
	return 0
}

func (x *ResponceWithdraw) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ResponceWithdraw) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type RequestTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *RequestTransaction) Reset() {
	*x = RequestTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTransaction) ProtoMessage() {}

func (x *RequestTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTransaction.ProtoReflect.Descriptor instead.
func (*RequestTransaction) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{5}
}

func (x *RequestTransaction) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type RequestActiveBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int32     `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Customer  *Customer `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *RequestActiveBalance) Reset() {
	*x = RequestActiveBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestActiveBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestActiveBalance) ProtoMessage() {}

func (x *RequestActiveBalance) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestActiveBalance.ProtoReflect.Descriptor instead.
func (*RequestActiveBalance) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{6}
}

func (x *RequestActiveBalance) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *RequestActiveBalance) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type RequestFrozenBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int32     `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Customer  *Customer `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *RequestFrozenBalance) Reset() {
	*x = RequestFrozenBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFrozenBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFrozenBalance) ProtoMessage() {}

func (x *RequestFrozenBalance) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFrozenBalance.ProtoReflect.Descriptor instead.
func (*RequestFrozenBalance) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{7}
}

func (x *RequestFrozenBalance) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *RequestFrozenBalance) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type RequestAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Req:
	//	*RequestAccount_ReqTrans
	//	*RequestAccount_ReqAct
	//	*RequestAccount_ReqFroz
	Req isRequestAccount_Req `protobuf_oneof:"req"`
}

func (x *RequestAccount) Reset() {
	*x = RequestAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAccount) ProtoMessage() {}

func (x *RequestAccount) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAccount.ProtoReflect.Descriptor instead.
func (*RequestAccount) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{8}
}

func (m *RequestAccount) GetReq() isRequestAccount_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (x *RequestAccount) GetReqTrans() *RequestTransaction {
	if x, ok := x.GetReq().(*RequestAccount_ReqTrans); ok {
		return x.ReqTrans
	}
	return nil
}

func (x *RequestAccount) GetReqAct() *RequestActiveBalance {
	if x, ok := x.GetReq().(*RequestAccount_ReqAct); ok {
		return x.ReqAct
	}
	return nil
}

func (x *RequestAccount) GetReqFroz() *RequestFrozenBalance {
	if x, ok := x.GetReq().(*RequestAccount_ReqFroz); ok {
		return x.ReqFroz
	}
	return nil
}

type isRequestAccount_Req interface {
	isRequestAccount_Req()
}

type RequestAccount_ReqTrans struct {
	ReqTrans *RequestTransaction `protobuf:"bytes,1,opt,name=reqTrans,proto3,oneof"`
}

type RequestAccount_ReqAct struct {
	ReqAct *RequestActiveBalance `protobuf:"bytes,2,opt,name=reqAct,proto3,oneof"`
}

type RequestAccount_ReqFroz struct {
	ReqFroz *RequestFrozenBalance `protobuf:"bytes,3,opt,name=reqFroz,proto3,oneof"`
}

func (*RequestAccount_ReqTrans) isRequestAccount_Req() {}

func (*RequestAccount_ReqAct) isRequestAccount_Req() {}

func (*RequestAccount_ReqFroz) isRequestAccount_Req() {}

type ResponceAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resp:
	//	*ResponceAccount_RespInv
	//	*ResponceAccount_RespWithdraw
	Resp isResponceAccount_Resp `protobuf_oneof:"resp"`
}

func (x *ResponceAccount) Reset() {
	*x = ResponceAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceAccount) ProtoMessage() {}

func (x *ResponceAccount) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceAccount.ProtoReflect.Descriptor instead.
func (*ResponceAccount) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{9}
}

func (m *ResponceAccount) GetResp() isResponceAccount_Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (x *ResponceAccount) GetRespInv() *ResponceInvoice {
	if x, ok := x.GetResp().(*ResponceAccount_RespInv); ok {
		return x.RespInv
	}
	return nil
}

func (x *ResponceAccount) GetRespWithdraw() *ResponceWithdraw {
	if x, ok := x.GetResp().(*ResponceAccount_RespWithdraw); ok {
		return x.RespWithdraw
	}
	return nil
}

type isResponceAccount_Resp interface {
	isResponceAccount_Resp()
}

type ResponceAccount_RespInv struct {
	RespInv *ResponceInvoice `protobuf:"bytes,1,opt,name=respInv,proto3,oneof"`
}

type ResponceAccount_RespWithdraw struct {
	RespWithdraw *ResponceWithdraw `protobuf:"bytes,2,opt,name=respWithdraw,proto3,oneof"`
}

func (*ResponceAccount_RespInv) isResponceAccount_Resp() {}

func (*ResponceAccount_RespWithdraw) isResponceAccount_Resp() {}

var File_Account_proto protoreflect.FileDescriptor

var file_Account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x64, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x75, 0x73,
	0x64, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x75,
	0x73, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x75, 0x62,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x72,
	0x75, 0x62, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x75, 0x72,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x65,
	0x75, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x74, 0x63,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x62,
	0x74, 0x63, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x6b, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x22, 0x6b,
	0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x22, 0x69, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6a, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x58, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x14,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x6c, 0x0a, 0x14, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x36, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0xde, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x72,
	0x65, 0x71, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x71, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x3f,
	0x0a, 0x06, 0x72, 0x65, 0x71, 0x41, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x71, 0x41, 0x63, 0x74, 0x12,
	0x41, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x46, 0x72, 0x6f, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x46, 0x72,
	0x6f, 0x7a, 0x42, 0x05, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x22, 0xa0, 0x01, 0x0a, 0x0f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x70, 0x49, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x73, 0x70, 0x49, 0x6e, 0x76, 0x12, 0x47, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x70, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6c, 0x6f, 0x76, 0x65,
	0x70, 0x69, 0x74, 0x73, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x46, 0x6f,
	0x72, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Account_proto_rawDescOnce sync.Once
	file_Account_proto_rawDescData = file_Account_proto_rawDesc
)

func file_Account_proto_rawDescGZIP() []byte {
	file_Account_proto_rawDescOnce.Do(func() {
		file_Account_proto_rawDescData = protoimpl.X.CompressGZIP(file_Account_proto_rawDescData)
	})
	return file_Account_proto_rawDescData
}

var file_Account_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_Account_proto_goTypes = []interface{}{
	(*Invoice)(nil),               // 0: account_service.Invoice
	(*ResponceActiveBalance)(nil), // 1: account_service.ResponceActiveBalance
	(*ResponceFrozenBalance)(nil), // 2: account_service.ResponceFrozenBalance
	(*ResponceInvoice)(nil),       // 3: account_service.ResponceInvoice
	(*ResponceWithdraw)(nil),      // 4: account_service.ResponceWithdraw
	(*RequestTransaction)(nil),    // 5: account_service.RequestTransaction
	(*RequestActiveBalance)(nil),  // 6: account_service.RequestActiveBalance
	(*RequestFrozenBalance)(nil),  // 7: account_service.RequestFrozenBalance
	(*RequestAccount)(nil),        // 8: account_service.RequestAccount
	(*ResponceAccount)(nil),       // 9: account_service.ResponceAccount
	(*Transaction)(nil),           // 10: Transaction_service.Transaction
	(*Customer)(nil),              // 11: Customer_service.Customer
}
var file_Account_proto_depIdxs = []int32{
	0,  // 0: account_service.ResponceActiveBalance.invoices:type_name -> account_service.Invoice
	0,  // 1: account_service.ResponceFrozenBalance.invoices:type_name -> account_service.Invoice
	10, // 2: account_service.RequestTransaction.transaction:type_name -> Transaction_service.Transaction
	11, // 3: account_service.RequestActiveBalance.customer:type_name -> Customer_service.Customer
	11, // 4: account_service.RequestFrozenBalance.customer:type_name -> Customer_service.Customer
	5,  // 5: account_service.RequestAccount.reqTrans:type_name -> account_service.RequestTransaction
	6,  // 6: account_service.RequestAccount.reqAct:type_name -> account_service.RequestActiveBalance
	7,  // 7: account_service.RequestAccount.reqFroz:type_name -> account_service.RequestFrozenBalance
	3,  // 8: account_service.ResponceAccount.respInv:type_name -> account_service.ResponceInvoice
	4,  // 9: account_service.ResponceAccount.respWithdraw:type_name -> account_service.ResponceWithdraw
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_Account_proto_init() }
func file_Account_proto_init() {
	if File_Account_proto != nil {
		return
	}
	file_Transaction_proto_init()
	file_Customer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
		file_Account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceActiveBalance); i {
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
		file_Account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceFrozenBalance); i {
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
		file_Account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceInvoice); i {
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
		file_Account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceWithdraw); i {
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
		file_Account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTransaction); i {
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
		file_Account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestActiveBalance); i {
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
		file_Account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFrozenBalance); i {
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
		file_Account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAccount); i {
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
		file_Account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceAccount); i {
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
	file_Account_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*RequestAccount_ReqTrans)(nil),
		(*RequestAccount_ReqAct)(nil),
		(*RequestAccount_ReqFroz)(nil),
	}
	file_Account_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*ResponceAccount_RespInv)(nil),
		(*ResponceAccount_RespWithdraw)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Account_proto_goTypes,
		DependencyIndexes: file_Account_proto_depIdxs,
		MessageInfos:      file_Account_proto_msgTypes,
	}.Build()
	File_Account_proto = out.File
	file_Account_proto_rawDesc = nil
	file_Account_proto_goTypes = nil
	file_Account_proto_depIdxs = nil
}
