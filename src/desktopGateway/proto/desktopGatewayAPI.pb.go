// Code generated by protoc-gen-go. DO NOT EDIT.
// source: desktopGateway/proto/desktopGatewayAPI.proto

package desktopGateway

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

// Messages for the estimation service package
type EstimationRequest struct {
	Bla                  string   `protobuf:"bytes,1,opt,name=bla,proto3" json:"bla,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstimationRequest) Reset()         { *m = EstimationRequest{} }
func (m *EstimationRequest) String() string { return proto.CompactTextString(m) }
func (*EstimationRequest) ProtoMessage()    {}
func (*EstimationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4293fa92ac258706, []int{0}
}

func (m *EstimationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstimationRequest.Unmarshal(m, b)
}
func (m *EstimationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstimationRequest.Marshal(b, m, deterministic)
}
func (m *EstimationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimationRequest.Merge(m, src)
}
func (m *EstimationRequest) XXX_Size() int {
	return xxx_messageInfo_EstimationRequest.Size(m)
}
func (m *EstimationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EstimationRequest proto.InternalMessageInfo

func (m *EstimationRequest) GetBla() string {
	if m != nil {
		return m.Bla
	}
	return ""
}

type CostEstimationRespose struct {
	Blabla               string   `protobuf:"bytes,1,opt,name=blabla,proto3" json:"blabla,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CostEstimationRespose) Reset()         { *m = CostEstimationRespose{} }
func (m *CostEstimationRespose) String() string { return proto.CompactTextString(m) }
func (*CostEstimationRespose) ProtoMessage()    {}
func (*CostEstimationRespose) Descriptor() ([]byte, []int) {
	return fileDescriptor_4293fa92ac258706, []int{1}
}

func (m *CostEstimationRespose) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostEstimationRespose.Unmarshal(m, b)
}
func (m *CostEstimationRespose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostEstimationRespose.Marshal(b, m, deterministic)
}
func (m *CostEstimationRespose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostEstimationRespose.Merge(m, src)
}
func (m *CostEstimationRespose) XXX_Size() int {
	return xxx_messageInfo_CostEstimationRespose.Size(m)
}
func (m *CostEstimationRespose) XXX_DiscardUnknown() {
	xxx_messageInfo_CostEstimationRespose.DiscardUnknown(m)
}

var xxx_messageInfo_CostEstimationRespose proto.InternalMessageInfo

func (m *CostEstimationRespose) GetBlabla() string {
	if m != nil {
		return m.Blabla
	}
	return ""
}

type PowerEstimationResponse struct {
	PowerEstimate        []float32 `protobuf:"fixed32,1,rep,packed,name=powerEstimate,proto3" json:"powerEstimate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PowerEstimationResponse) Reset()         { *m = PowerEstimationResponse{} }
func (m *PowerEstimationResponse) String() string { return proto.CompactTextString(m) }
func (*PowerEstimationResponse) ProtoMessage()    {}
func (*PowerEstimationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4293fa92ac258706, []int{2}
}

func (m *PowerEstimationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PowerEstimationResponse.Unmarshal(m, b)
}
func (m *PowerEstimationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PowerEstimationResponse.Marshal(b, m, deterministic)
}
func (m *PowerEstimationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PowerEstimationResponse.Merge(m, src)
}
func (m *PowerEstimationResponse) XXX_Size() int {
	return xxx_messageInfo_PowerEstimationResponse.Size(m)
}
func (m *PowerEstimationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PowerEstimationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PowerEstimationResponse proto.InternalMessageInfo

func (m *PowerEstimationResponse) GetPowerEstimate() []float32 {
	if m != nil {
		return m.PowerEstimate
	}
	return nil
}

// Messages for the frontend login
type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	HashedPassword       string   `protobuf:"bytes,2,opt,name=hashedPassword,proto3" json:"hashedPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4293fa92ac258706, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetHashedPassword() string {
	if m != nil {
		return m.HashedPassword
	}
	return ""
}

type LoginResponse struct {
	Permissions          string   `protobuf:"bytes,1,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4293fa92ac258706, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetPermissions() string {
	if m != nil {
		return m.Permissions
	}
	return ""
}

func init() {
	proto.RegisterType((*EstimationRequest)(nil), "EstimationRequest")
	proto.RegisterType((*CostEstimationRespose)(nil), "CostEstimationRespose")
	proto.RegisterType((*PowerEstimationResponse)(nil), "PowerEstimationResponse")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
}

func init() {
	proto.RegisterFile("desktopGateway/proto/desktopGatewayAPI.proto", fileDescriptor_4293fa92ac258706)
}

var fileDescriptor_4293fa92ac258706 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0x86, 0x43, 0x9f, 0x6e, 0x6c, 0x01, 0x67, 0xd9, 0x69, 0x54, 0x37, 0x76, 0x90,
	0x0c, 0x27, 0x78, 0x11, 0x94, 0x29, 0x22, 0x82, 0x87, 0xd2, 0xdd, 0xbc, 0x65, 0xf6, 0xe1, 0x82,
	0x6b, 0x13, 0xf3, 0x32, 0x8b, 0xdf, 0xc5, 0x0f, 0x2b, 0x96, 0x30, 0xdb, 0xae, 0xde, 0xfa, 0x7e,
	0x7d, 0xf9, 0xe7, 0xbd, 0x1f, 0x81, 0xf3, 0x18, 0xe9, 0xdd, 0x2a, 0xfd, 0x28, 0x2c, 0x66, 0xe2,
	0x6b, 0xaa, 0x8d, 0xb2, 0x6a, 0x5a, 0x86, 0xf3, 0xf0, 0x89, 0xe7, 0x3c, 0x18, 0x41, 0xef, 0x81,
	0xac, 0x4c, 0x84, 0x95, 0x2a, 0x8d, 0xf0, 0x63, 0x83, 0x64, 0x59, 0x17, 0x9a, 0xcb, 0xb5, 0xf0,
	0xbd, 0xa1, 0x37, 0x39, 0x88, 0x7e, 0x3f, 0x83, 0x29, 0x1c, 0xdf, 0x2b, 0xb2, 0xc5, 0x56, 0xd2,
	0x8a, 0x90, 0xf5, 0xa1, 0xb5, 0x5c, 0x8b, 0xbf, 0x6e, 0x57, 0x05, 0xb7, 0x70, 0x12, 0xaa, 0x0c,
	0x4d, 0xe5, 0x44, 0x4a, 0xc8, 0xce, 0xa0, 0xad, 0x0b, 0xbf, 0xd0, 0xf7, 0x86, 0xcd, 0x49, 0x23,
	0x2a, 0xc3, 0x20, 0x82, 0xa3, 0x67, 0xf5, 0x26, 0xb7, 0x33, 0x0d, 0x60, 0x7f, 0x43, 0x68, 0x52,
	0x91, 0xa0, 0xbb, 0x6a, 0x5b, 0xb3, 0x31, 0x74, 0x56, 0x82, 0x56, 0x18, 0x87, 0x82, 0x28, 0x53,
	0x26, 0xf6, 0x1b, 0x79, 0x47, 0x85, 0x06, 0x17, 0xd0, 0x76, 0x99, 0x6e, 0x94, 0x21, 0x1c, 0x6a,
	0x34, 0x89, 0x24, 0x92, 0x2a, 0x25, 0x97, 0x5b, 0x44, 0xb3, 0x6f, 0x6f, 0x67, 0x91, 0x05, 0x9a,
	0x4f, 0xf9, 0x8a, 0xc4, 0x6e, 0xa0, 0x5b, 0x96, 0xb2, 0x08, 0x19, 0xe3, 0x3b, 0x3a, 0x07, 0x7d,
	0x5e, 0xef, 0x6e, 0x0e, 0x3d, 0x5d, 0x89, 0xae, 0x0f, 0xf0, 0xf9, 0x3f, 0x2e, 0x67, 0x57, 0xce,
	0x92, 0x9b, 0x89, 0x8d, 0x61, 0x2f, 0xaf, 0x59, 0x9b, 0x17, 0xed, 0x0d, 0x3a, 0xbc, 0xb4, 0xf8,
	0xdd, 0xe8, 0xe5, 0xb4, 0xee, 0x99, 0x5c, 0x97, 0xe1, 0xb2, 0x95, 0xd3, 0xcb, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x06, 0xa7, 0x9f, 0x5f, 0x54, 0x02, 0x00, 0x00,
}
