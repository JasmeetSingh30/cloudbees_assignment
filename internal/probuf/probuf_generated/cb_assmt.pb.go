// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: cb_assmt.proto

package probuf_generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SeatModificationRequest_SectionEnum int32

const (
	SeatModificationRequest_SECTION_A SeatModificationRequest_SectionEnum = 0
	SeatModificationRequest_SECTION_B SeatModificationRequest_SectionEnum = 1
)

// Enum value maps for SeatModificationRequest_SectionEnum.
var (
	SeatModificationRequest_SectionEnum_name = map[int32]string{
		0: "SECTION_A",
		1: "SECTION_B",
	}
	SeatModificationRequest_SectionEnum_value = map[string]int32{
		"SECTION_A": 0,
		"SECTION_B": 1,
	}
)

func (x SeatModificationRequest_SectionEnum) Enum() *SeatModificationRequest_SectionEnum {
	p := new(SeatModificationRequest_SectionEnum)
	*p = x
	return p
}

func (x SeatModificationRequest_SectionEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeatModificationRequest_SectionEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cb_assmt_proto_enumTypes[0].Descriptor()
}

func (SeatModificationRequest_SectionEnum) Type() protoreflect.EnumType {
	return &file_cb_assmt_proto_enumTypes[0]
}

func (x SeatModificationRequest_SectionEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeatModificationRequest_SectionEnum.Descriptor instead.
func (SeatModificationRequest_SectionEnum) EnumDescriptor() ([]byte, []int) {
	return file_cb_assmt_proto_rawDescGZIP(), []int{5, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_assmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cb_assmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_cb_assmt_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TicketCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *TicketCreateRequest) Reset() {
	*x = TicketCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_assmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketCreateRequest) ProtoMessage() {}

func (x *TicketCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cb_assmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketCreateRequest.ProtoReflect.Descriptor instead.
func (*TicketCreateRequest) Descriptor() ([]byte, []int) {
	return file_cb_assmt_proto_rawDescGZIP(), []int{1}
}

func (x *TicketCreateRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TicketCreateRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TicketCreateRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *User   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	From     string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To       string  `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	TicketId int32   `protobuf:"varint,4,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
	Section  string  `protobuf:"bytes,5,opt,name=section,proto3" json:"section,omitempty"`
	SeatNo   int32   `protobuf:"varint,6,opt,name=seatNo,proto3" json:"seatNo,omitempty"`
	Price    float32 `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_assmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_cb_assmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_cb_assmt_proto_rawDescGZIP(), []int{2}
}

func (x *Receipt) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Receipt) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Receipt) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Receipt) GetTicketId() int32 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *Receipt) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Receipt) GetSeatNo() int32 {
	if x != nil {
		return x.SeatNo
	}
	return 0
}

func (x *Receipt) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Seat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	SeatNo int32 `protobuf:"varint,2,opt,name=seatNo,proto3" json:"seatNo,omitempty"`
}

func (x *Seat) Reset() {
	*x = Seat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_assmt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seat) ProtoMessage() {}

func (x *Seat) ProtoReflect() protoreflect.Message {
	mi := &file_cb_assmt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seat.ProtoReflect.Descriptor instead.
func (*Seat) Descriptor() ([]byte, []int) {
	return file_cb_assmt_proto_rawDescGZIP(), []int{3}
}

func (x *Seat) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Seat) GetSeatNo() int32 {
	if x != nil {
		return x.SeatNo
	}
	return 0
}

type Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionName string  `protobuf:"bytes,1,opt,name=section_name,json=sectionName,proto3" json:"section_name,omitempty"`
	Seats       []*Seat `protobuf:"bytes,2,rep,name=seats,proto3" json:"seats,omitempty"`
}

func (x *Section) Reset() {
	*x = Section{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_assmt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_cb_assmt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Section.ProtoReflect.Descriptor instead.
func (*Section) Descriptor() ([]byte, []int) {
	return file_cb_assmt_proto_rawDescGZIP(), []int{4}
}

func (x *Section) GetSectionName() string {
	if x != nil {
		return x.SectionName
	}
	return ""
}

func (x *Section) GetSeats() []*Seat {
	if x != nil {
		return x.Seats
	}
	return nil
}

type SeatModificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *User                               `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	TicketId int32                               `protobuf:"varint,2,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
	Section  SeatModificationRequest_SectionEnum `protobuf:"varint,3,opt,name=section,proto3,enum=railwayTicketing.SeatModificationRequest_SectionEnum" json:"section,omitempty"`
	Seat     int32                               `protobuf:"varint,4,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *SeatModificationRequest) Reset() {
	*x = SeatModificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_assmt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatModificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatModificationRequest) ProtoMessage() {}

func (x *SeatModificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cb_assmt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatModificationRequest.ProtoReflect.Descriptor instead.
func (*SeatModificationRequest) Descriptor() ([]byte, []int) {
	return file_cb_assmt_proto_rawDescGZIP(), []int{5}
}

func (x *SeatModificationRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SeatModificationRequest) GetTicketId() int32 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *SeatModificationRequest) GetSection() SeatModificationRequest_SectionEnum {
	if x != nil {
		return x.Section
	}
	return SeatModificationRequest_SECTION_A
}

func (x *SeatModificationRequest) GetSeat() int32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

type TicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	TicketId int32 `protobuf:"varint,2,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
}

func (x *TicketRequest) Reset() {
	*x = TicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_assmt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketRequest) ProtoMessage() {}

func (x *TicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cb_assmt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketRequest.ProtoReflect.Descriptor instead.
func (*TicketRequest) Descriptor() ([]byte, []int) {
	return file_cb_assmt_proto_rawDescGZIP(), []int{6}
}

func (x *TicketRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TicketRequest) GetTicketId() int32 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type TrainSectionAllocated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string                              `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To      string                              `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Section SeatModificationRequest_SectionEnum `protobuf:"varint,3,opt,name=Section,proto3,enum=railwayTicketing.SeatModificationRequest_SectionEnum" json:"Section,omitempty"`
}

func (x *TrainSectionAllocated) Reset() {
	*x = TrainSectionAllocated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_assmt_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainSectionAllocated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainSectionAllocated) ProtoMessage() {}

func (x *TrainSectionAllocated) ProtoReflect() protoreflect.Message {
	mi := &file_cb_assmt_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainSectionAllocated.ProtoReflect.Descriptor instead.
func (*TrainSectionAllocated) Descriptor() ([]byte, []int) {
	return file_cb_assmt_proto_rawDescGZIP(), []int{7}
}

func (x *TrainSectionAllocated) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TrainSectionAllocated) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TrainSectionAllocated) GetSection() SeatModificationRequest_SectionEnum {
	if x != nil {
		return x.Section
	}
	return SeatModificationRequest_SECTION_A
}

var File_cb_assmt_proto protoreflect.FileDescriptor

var file_cb_assmt_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x62, 0x5f, 0x61, 0x73, 0x73, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x58, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x65, 0x0a, 0x13, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x22, 0xbd, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x2a, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x61, 0x69,
	0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x4a, 0x0a, 0x04, 0x53, 0x65, 0x61, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x22, 0x5a, 0x0a, 0x07,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x65,
	0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x61, 0x69, 0x6c,
	0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61,
	0x74, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x22, 0xf3, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x61,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e,
	0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x65, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x65, 0x61,
	0x74, 0x22, 0x2b, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x10, 0x01, 0x22, 0x57,
	0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x4f, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x07, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xa1, 0x03, 0x0a, 0x0e, 0x52, 0x61, 0x69, 0x6c, 0x77,
	0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x09, 0x42, 0x75, 0x79,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x49, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61,
	0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77,
	0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x12, 0x58, 0x0a, 0x12, 0x56, 0x69, 0x65, 0x77, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x72, 0x61, 0x69, 0x6c,
	0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x1a, 0x19, 0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a,
	0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1f, 0x2e,
	0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x61, 0x74, 0x12, 0x29, 0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x72, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cb_assmt_proto_rawDescOnce sync.Once
	file_cb_assmt_proto_rawDescData = file_cb_assmt_proto_rawDesc
)

func file_cb_assmt_proto_rawDescGZIP() []byte {
	file_cb_assmt_proto_rawDescOnce.Do(func() {
		file_cb_assmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_cb_assmt_proto_rawDescData)
	})
	return file_cb_assmt_proto_rawDescData
}

var file_cb_assmt_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cb_assmt_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cb_assmt_proto_goTypes = []any{
	(SeatModificationRequest_SectionEnum)(0), // 0: railwayTicketing.SeatModificationRequest.SectionEnum
	(*User)(nil),                             // 1: railwayTicketing.User
	(*TicketCreateRequest)(nil),              // 2: railwayTicketing.TicketCreateRequest
	(*Receipt)(nil),                          // 3: railwayTicketing.Receipt
	(*Seat)(nil),                             // 4: railwayTicketing.Seat
	(*Section)(nil),                          // 5: railwayTicketing.Section
	(*SeatModificationRequest)(nil),          // 6: railwayTicketing.SeatModificationRequest
	(*TicketRequest)(nil),                    // 7: railwayTicketing.TicketRequest
	(*TrainSectionAllocated)(nil),            // 8: railwayTicketing.TrainSectionAllocated
	(*emptypb.Empty)(nil),                    // 9: google.protobuf.Empty
}
var file_cb_assmt_proto_depIdxs = []int32{
	1,  // 0: railwayTicketing.TicketCreateRequest.user:type_name -> railwayTicketing.User
	1,  // 1: railwayTicketing.Receipt.user:type_name -> railwayTicketing.User
	1,  // 2: railwayTicketing.Seat.user:type_name -> railwayTicketing.User
	4,  // 3: railwayTicketing.Section.seats:type_name -> railwayTicketing.Seat
	1,  // 4: railwayTicketing.SeatModificationRequest.user:type_name -> railwayTicketing.User
	0,  // 5: railwayTicketing.SeatModificationRequest.section:type_name -> railwayTicketing.SeatModificationRequest.SectionEnum
	1,  // 6: railwayTicketing.TicketRequest.user:type_name -> railwayTicketing.User
	0,  // 7: railwayTicketing.TrainSectionAllocated.Section:type_name -> railwayTicketing.SeatModificationRequest.SectionEnum
	2,  // 8: railwayTicketing.RailwayService.BuyTicket:input_type -> railwayTicketing.TicketCreateRequest
	7,  // 9: railwayTicketing.RailwayService.ViewReceipt:input_type -> railwayTicketing.TicketRequest
	8,  // 10: railwayTicketing.RailwayService.ViewAllocatedSeats:input_type -> railwayTicketing.TrainSectionAllocated
	7,  // 11: railwayTicketing.RailwayService.CancelTicket:input_type -> railwayTicketing.TicketRequest
	6,  // 12: railwayTicketing.RailwayService.ModifySeat:input_type -> railwayTicketing.SeatModificationRequest
	3,  // 13: railwayTicketing.RailwayService.BuyTicket:output_type -> railwayTicketing.Receipt
	3,  // 14: railwayTicketing.RailwayService.ViewReceipt:output_type -> railwayTicketing.Receipt
	5,  // 15: railwayTicketing.RailwayService.ViewAllocatedSeats:output_type -> railwayTicketing.Section
	9,  // 16: railwayTicketing.RailwayService.CancelTicket:output_type -> google.protobuf.Empty
	3,  // 17: railwayTicketing.RailwayService.ModifySeat:output_type -> railwayTicketing.Receipt
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_cb_assmt_proto_init() }
func file_cb_assmt_proto_init() {
	if File_cb_assmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cb_assmt_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_cb_assmt_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TicketCreateRequest); i {
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
		file_cb_assmt_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Receipt); i {
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
		file_cb_assmt_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Seat); i {
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
		file_cb_assmt_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Section); i {
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
		file_cb_assmt_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SeatModificationRequest); i {
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
		file_cb_assmt_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*TicketRequest); i {
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
		file_cb_assmt_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*TrainSectionAllocated); i {
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
			RawDescriptor: file_cb_assmt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cb_assmt_proto_goTypes,
		DependencyIndexes: file_cb_assmt_proto_depIdxs,
		EnumInfos:         file_cb_assmt_proto_enumTypes,
		MessageInfos:      file_cb_assmt_proto_msgTypes,
	}.Build()
	File_cb_assmt_proto = out.File
	file_cb_assmt_proto_rawDesc = nil
	file_cb_assmt_proto_goTypes = nil
	file_cb_assmt_proto_depIdxs = nil
}
