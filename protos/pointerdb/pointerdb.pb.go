// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pointerdb.proto

package pointerdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RedundancyScheme_SchemeType int32

const (
	RedundancyScheme_RS RedundancyScheme_SchemeType = 0
)

var RedundancyScheme_SchemeType_name = map[int32]string{
	0: "RS",
}
var RedundancyScheme_SchemeType_value = map[string]int32{
	"RS": 0,
}

func (x RedundancyScheme_SchemeType) String() string {
	return proto.EnumName(RedundancyScheme_SchemeType_name, int32(x))
}
func (RedundancyScheme_SchemeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{0, 0}
}

type EncryptionScheme_EncryptionType int32

const (
	EncryptionScheme_AESGCM    EncryptionScheme_EncryptionType = 0
	EncryptionScheme_SECRETBOX EncryptionScheme_EncryptionType = 1
)

var EncryptionScheme_EncryptionType_name = map[int32]string{
	0: "AESGCM",
	1: "SECRETBOX",
}
var EncryptionScheme_EncryptionType_value = map[string]int32{
	"AESGCM":    0,
	"SECRETBOX": 1,
}

func (x EncryptionScheme_EncryptionType) String() string {
	return proto.EnumName(EncryptionScheme_EncryptionType_name, int32(x))
}
func (EncryptionScheme_EncryptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{1, 0}
}

type Pointer_DataType int32

const (
	Pointer_INLINE Pointer_DataType = 0
	Pointer_REMOTE Pointer_DataType = 1
)

var Pointer_DataType_name = map[int32]string{
	0: "INLINE",
	1: "REMOTE",
}
var Pointer_DataType_value = map[string]int32{
	"INLINE": 0,
	"REMOTE": 1,
}

func (x Pointer_DataType) String() string {
	return proto.EnumName(Pointer_DataType_name, int32(x))
}
func (Pointer_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{4, 0}
}

type RedundancyScheme struct {
	Type RedundancyScheme_SchemeType `protobuf:"varint,1,opt,name=type,proto3,enum=pointerdb.RedundancyScheme_SchemeType" json:"type,omitempty"`
	// these values apply to RS encoding
	MinReq               int64    `protobuf:"varint,2,opt,name=min_req,json=minReq,proto3" json:"min_req,omitempty"`
	Total                int64    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	RepairThreshold      int64    `protobuf:"varint,4,opt,name=repair_threshold,json=repairThreshold,proto3" json:"repair_threshold,omitempty"`
	SuccessThreshold     int64    `protobuf:"varint,5,opt,name=success_threshold,json=successThreshold,proto3" json:"success_threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedundancyScheme) Reset()         { *m = RedundancyScheme{} }
func (m *RedundancyScheme) String() string { return proto.CompactTextString(m) }
func (*RedundancyScheme) ProtoMessage()    {}
func (*RedundancyScheme) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{0}
}
func (m *RedundancyScheme) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedundancyScheme.Unmarshal(m, b)
}
func (m *RedundancyScheme) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedundancyScheme.Marshal(b, m, deterministic)
}
func (dst *RedundancyScheme) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedundancyScheme.Merge(dst, src)
}
func (m *RedundancyScheme) XXX_Size() int {
	return xxx_messageInfo_RedundancyScheme.Size(m)
}
func (m *RedundancyScheme) XXX_DiscardUnknown() {
	xxx_messageInfo_RedundancyScheme.DiscardUnknown(m)
}

var xxx_messageInfo_RedundancyScheme proto.InternalMessageInfo

func (m *RedundancyScheme) GetType() RedundancyScheme_SchemeType {
	if m != nil {
		return m.Type
	}
	return RedundancyScheme_RS
}

func (m *RedundancyScheme) GetMinReq() int64 {
	if m != nil {
		return m.MinReq
	}
	return 0
}

func (m *RedundancyScheme) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RedundancyScheme) GetRepairThreshold() int64 {
	if m != nil {
		return m.RepairThreshold
	}
	return 0
}

func (m *RedundancyScheme) GetSuccessThreshold() int64 {
	if m != nil {
		return m.SuccessThreshold
	}
	return 0
}

type EncryptionScheme struct {
	Type                   EncryptionScheme_EncryptionType `protobuf:"varint,1,opt,name=type,proto3,enum=pointerdb.EncryptionScheme_EncryptionType" json:"type,omitempty"`
	EncryptedEncryptionKey []byte                          `protobuf:"bytes,2,opt,name=encrypted_encryption_key,json=encryptedEncryptionKey,proto3" json:"encrypted_encryption_key,omitempty"`
	EncryptedStartingNonce []byte                          `protobuf:"bytes,3,opt,name=encrypted_starting_nonce,json=encryptedStartingNonce,proto3" json:"encrypted_starting_nonce,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                        `json:"-"`
	XXX_unrecognized       []byte                          `json:"-"`
	XXX_sizecache          int32                           `json:"-"`
}

func (m *EncryptionScheme) Reset()         { *m = EncryptionScheme{} }
func (m *EncryptionScheme) String() string { return proto.CompactTextString(m) }
func (*EncryptionScheme) ProtoMessage()    {}
func (*EncryptionScheme) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{1}
}
func (m *EncryptionScheme) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptionScheme.Unmarshal(m, b)
}
func (m *EncryptionScheme) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptionScheme.Marshal(b, m, deterministic)
}
func (dst *EncryptionScheme) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptionScheme.Merge(dst, src)
}
func (m *EncryptionScheme) XXX_Size() int {
	return xxx_messageInfo_EncryptionScheme.Size(m)
}
func (m *EncryptionScheme) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptionScheme.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptionScheme proto.InternalMessageInfo

func (m *EncryptionScheme) GetType() EncryptionScheme_EncryptionType {
	if m != nil {
		return m.Type
	}
	return EncryptionScheme_AESGCM
}

func (m *EncryptionScheme) GetEncryptedEncryptionKey() []byte {
	if m != nil {
		return m.EncryptedEncryptionKey
	}
	return nil
}

func (m *EncryptionScheme) GetEncryptedStartingNonce() []byte {
	if m != nil {
		return m.EncryptedStartingNonce
	}
	return nil
}

type RemotePiece struct {
	PieceNum             int64    `protobuf:"varint,1,opt,name=piece_num,json=pieceNum,proto3" json:"piece_num,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemotePiece) Reset()         { *m = RemotePiece{} }
func (m *RemotePiece) String() string { return proto.CompactTextString(m) }
func (*RemotePiece) ProtoMessage()    {}
func (*RemotePiece) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{2}
}
func (m *RemotePiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemotePiece.Unmarshal(m, b)
}
func (m *RemotePiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemotePiece.Marshal(b, m, deterministic)
}
func (dst *RemotePiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemotePiece.Merge(dst, src)
}
func (m *RemotePiece) XXX_Size() int {
	return xxx_messageInfo_RemotePiece.Size(m)
}
func (m *RemotePiece) XXX_DiscardUnknown() {
	xxx_messageInfo_RemotePiece.DiscardUnknown(m)
}

var xxx_messageInfo_RemotePiece proto.InternalMessageInfo

func (m *RemotePiece) GetPieceNum() int64 {
	if m != nil {
		return m.PieceNum
	}
	return 0
}

func (m *RemotePiece) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type RemoteSegment struct {
	Redundancy           *RedundancyScheme `protobuf:"bytes,1,opt,name=redundancy,proto3" json:"redundancy,omitempty"`
	PieceId              string            `protobuf:"bytes,2,opt,name=piece_id,json=pieceId,proto3" json:"piece_id,omitempty"`
	RemotePieces         []*RemotePiece    `protobuf:"bytes,3,rep,name=remote_pieces,json=remotePieces,proto3" json:"remote_pieces,omitempty"`
	MerkleRoot           []byte            `protobuf:"bytes,4,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RemoteSegment) Reset()         { *m = RemoteSegment{} }
func (m *RemoteSegment) String() string { return proto.CompactTextString(m) }
func (*RemoteSegment) ProtoMessage()    {}
func (*RemoteSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{3}
}
func (m *RemoteSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteSegment.Unmarshal(m, b)
}
func (m *RemoteSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteSegment.Marshal(b, m, deterministic)
}
func (dst *RemoteSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteSegment.Merge(dst, src)
}
func (m *RemoteSegment) XXX_Size() int {
	return xxx_messageInfo_RemoteSegment.Size(m)
}
func (m *RemoteSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteSegment.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteSegment proto.InternalMessageInfo

func (m *RemoteSegment) GetRedundancy() *RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *RemoteSegment) GetPieceId() string {
	if m != nil {
		return m.PieceId
	}
	return ""
}

func (m *RemoteSegment) GetRemotePieces() []*RemotePiece {
	if m != nil {
		return m.RemotePieces
	}
	return nil
}

func (m *RemoteSegment) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

type Pointer struct {
	Type                 Pointer_DataType     `protobuf:"varint,1,opt,name=type,proto3,enum=pointerdb.Pointer_DataType" json:"type,omitempty"`
	InlineSegment        []byte               `protobuf:"bytes,3,opt,name=inline_segment,json=inlineSegment,proto3" json:"inline_segment,omitempty"`
	Remote               *RemoteSegment       `protobuf:"bytes,4,opt,name=remote,proto3" json:"remote,omitempty"`
	Size                 int64                `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	CreationDate         *timestamp.Timestamp `protobuf:"bytes,6,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	ExpirationDate       *timestamp.Timestamp `protobuf:"bytes,7,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	Metadata             []byte               `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pointer) Reset()         { *m = Pointer{} }
func (m *Pointer) String() string { return proto.CompactTextString(m) }
func (*Pointer) ProtoMessage()    {}
func (*Pointer) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{4}
}
func (m *Pointer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pointer.Unmarshal(m, b)
}
func (m *Pointer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pointer.Marshal(b, m, deterministic)
}
func (dst *Pointer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pointer.Merge(dst, src)
}
func (m *Pointer) XXX_Size() int {
	return xxx_messageInfo_Pointer.Size(m)
}
func (m *Pointer) XXX_DiscardUnknown() {
	xxx_messageInfo_Pointer.DiscardUnknown(m)
}

var xxx_messageInfo_Pointer proto.InternalMessageInfo

func (m *Pointer) GetType() Pointer_DataType {
	if m != nil {
		return m.Type
	}
	return Pointer_INLINE
}

func (m *Pointer) GetInlineSegment() []byte {
	if m != nil {
		return m.InlineSegment
	}
	return nil
}

func (m *Pointer) GetRemote() *RemoteSegment {
	if m != nil {
		return m.Remote
	}
	return nil
}

func (m *Pointer) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Pointer) GetCreationDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

func (m *Pointer) GetExpirationDate() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationDate
	}
	return nil
}

func (m *Pointer) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// PutRequest is a request message for the Put rpc call
type PutRequest struct {
	Path                 []byte   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Pointer              *Pointer `protobuf:"bytes,2,opt,name=pointer,proto3" json:"pointer,omitempty"`
	APIKey               []byte   `protobuf:"bytes,3,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{5}
}
func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (dst *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(dst, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *PutRequest) GetPointer() *Pointer {
	if m != nil {
		return m.Pointer
	}
	return nil
}

func (m *PutRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// GetRequest is a request message for the Get rpc call
type GetRequest struct {
	Path                 []byte   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{6}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *GetRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// ListRequest is a request message for the List rpc call
type ListRequest struct {
	StartingPathKey      []byte   `protobuf:"bytes,1,opt,name=starting_path_key,json=startingPathKey,proto3" json:"starting_path_key,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	APIKey               []byte   `protobuf:"bytes,3,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{7}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetStartingPathKey() []byte {
	if m != nil {
		return m.StartingPathKey
	}
	return nil
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// PutResponse is a response message for the Put rpc call
type PutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{8}
}
func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (dst *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(dst, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

// GetResponse is a response message for the Get rpc call
type GetResponse struct {
	Pointer              []byte   `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{9}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetPointer() []byte {
	if m != nil {
		return m.Pointer
	}
	return nil
}

// ListResponse is a response message for the List rpc call
type ListResponse struct {
	Paths                [][]byte `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	Truncated            bool     `protobuf:"varint,2,opt,name=truncated,proto3" json:"truncated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{10}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *ListResponse) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

type DeleteRequest struct {
	Path                 []byte   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{11}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *DeleteRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// DeleteResponse is a response message for the Delete rpc call
type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_525e73273472a0cd, []int{12}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RedundancyScheme)(nil), "pointerdb.RedundancyScheme")
	proto.RegisterType((*EncryptionScheme)(nil), "pointerdb.EncryptionScheme")
	proto.RegisterType((*RemotePiece)(nil), "pointerdb.RemotePiece")
	proto.RegisterType((*RemoteSegment)(nil), "pointerdb.RemoteSegment")
	proto.RegisterType((*Pointer)(nil), "pointerdb.Pointer")
	proto.RegisterType((*PutRequest)(nil), "pointerdb.PutRequest")
	proto.RegisterType((*GetRequest)(nil), "pointerdb.GetRequest")
	proto.RegisterType((*ListRequest)(nil), "pointerdb.ListRequest")
	proto.RegisterType((*PutResponse)(nil), "pointerdb.PutResponse")
	proto.RegisterType((*GetResponse)(nil), "pointerdb.GetResponse")
	proto.RegisterType((*ListResponse)(nil), "pointerdb.ListResponse")
	proto.RegisterType((*DeleteRequest)(nil), "pointerdb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "pointerdb.DeleteResponse")
	proto.RegisterEnum("pointerdb.RedundancyScheme_SchemeType", RedundancyScheme_SchemeType_name, RedundancyScheme_SchemeType_value)
	proto.RegisterEnum("pointerdb.EncryptionScheme_EncryptionType", EncryptionScheme_EncryptionType_name, EncryptionScheme_EncryptionType_value)
	proto.RegisterEnum("pointerdb.Pointer_DataType", Pointer_DataType_name, Pointer_DataType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PointerDBClient is the client API for PointerDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PointerDBClient interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type pointerDBClient struct {
	cc *grpc.ClientConn
}

func NewPointerDBClient(cc *grpc.ClientConn) PointerDBClient {
	return &pointerDBClient{cc}
}

func (c *pointerDBClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/pointerdb.PointerDB/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointerDBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pointerdb.PointerDB/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointerDBClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/pointerdb.PointerDB/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointerDBClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pointerdb.PointerDB/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointerDBServer is the server API for PointerDB service.
type PointerDBServer interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterPointerDBServer(s *grpc.Server, srv PointerDBServer) {
	s.RegisterService(&_PointerDB_serviceDesc, srv)
}

func _PointerDB_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointerDB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointerDB_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointerDB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PointerDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pointerdb.PointerDB",
	HandlerType: (*PointerDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _PointerDB_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PointerDB_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PointerDB_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PointerDB_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pointerdb.proto",
}

func init() { proto.RegisterFile("pointerdb.proto", fileDescriptor_pointerdb_525e73273472a0cd) }

var fileDescriptor_pointerdb_525e73273472a0cd = []byte{
	// 859 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x6f, 0xdb, 0x46,
	0x0c, 0x8f, 0xa2, 0x44, 0xb6, 0x29, 0xdb, 0x51, 0x0f, 0x59, 0xaa, 0xba, 0x03, 0x1a, 0x08, 0xd8,
	0x96, 0xb5, 0x83, 0x33, 0x78, 0x05, 0x56, 0x2c, 0xd8, 0x86, 0x26, 0x31, 0x02, 0xa3, 0x6d, 0x6a,
	0x9c, 0xfd, 0xb0, 0x37, 0x41, 0x91, 0x58, 0x5b, 0xa8, 0x75, 0x52, 0xa4, 0x13, 0x30, 0xef, 0xfb,
	0xed, 0xd3, 0x0c, 0xc3, 0x1e, 0xf6, 0x05, 0x06, 0xdd, 0x9d, 0xfe, 0x25, 0x5b, 0x06, 0xf4, 0xc9,
	0x22, 0xf9, 0xfb, 0xf1, 0xc8, 0x1f, 0x49, 0xc3, 0x41, 0x12, 0x87, 0x8c, 0x63, 0x1a, 0xdc, 0x8c,
	0x93, 0x34, 0xe6, 0x31, 0xe9, 0x55, 0x8e, 0xd1, 0xb3, 0x55, 0x1c, 0xaf, 0x36, 0x78, 0x2a, 0x02,
	0x37, 0xf9, 0x87, 0x53, 0x1e, 0x46, 0x98, 0x71, 0x2f, 0x4a, 0x24, 0xd6, 0xf9, 0x43, 0x03, 0x8b,
	0x62, 0x90, 0xb3, 0xc0, 0x63, 0xfe, 0x76, 0xe1, 0xaf, 0x31, 0x42, 0xf2, 0x03, 0xec, 0xf1, 0x6d,
	0x82, 0xb6, 0x76, 0xac, 0x9d, 0x0c, 0x27, 0x5f, 0x8e, 0xeb, 0x07, 0xee, 0x42, 0xc7, 0xf2, 0x67,
	0xb9, 0x4d, 0x90, 0x0a, 0x0e, 0x79, 0x0c, 0x9d, 0x28, 0x64, 0x6e, 0x8a, 0xb7, 0xf6, 0xee, 0xb1,
	0x76, 0xa2, 0x53, 0x23, 0x0a, 0x19, 0xc5, 0x5b, 0x72, 0x08, 0xfb, 0x3c, 0xe6, 0xde, 0xc6, 0xd6,
	0x85, 0x5b, 0x1a, 0xe4, 0x6b, 0xb0, 0x52, 0x4c, 0xbc, 0x30, 0x75, 0xf9, 0x3a, 0xc5, 0x6c, 0x1d,
	0x6f, 0x02, 0x7b, 0x4f, 0x00, 0x0e, 0xa4, 0x7f, 0x59, 0xba, 0xc9, 0x0b, 0x78, 0x94, 0xe5, 0xbe,
	0x8f, 0x59, 0xd6, 0xc0, 0xee, 0x0b, 0xac, 0xa5, 0x02, 0x15, 0xd8, 0x39, 0x04, 0xa8, 0x4b, 0x23,
	0x06, 0xec, 0xd2, 0x85, 0xb5, 0xe3, 0xfc, 0xad, 0x81, 0x35, 0x65, 0x7e, 0xba, 0x4d, 0x78, 0x18,
	0x33, 0xd5, 0xed, 0x4f, 0xad, 0x6e, 0x9f, 0x37, 0xba, 0xbd, 0x0b, 0x6d, 0x38, 0x1a, 0x1d, 0xbf,
	0x02, 0x1b, 0xa5, 0x1f, 0x03, 0x17, 0x2b, 0x84, 0xfb, 0x11, 0xb7, 0x42, 0x82, 0x3e, 0x3d, 0xaa,
	0xe2, 0x75, 0x82, 0x37, 0xb8, 0x6d, 0x33, 0x33, 0xee, 0xa5, 0x3c, 0x64, 0x2b, 0x97, 0xc5, 0xcc,
	0x47, 0xa1, 0x52, 0x93, 0xb9, 0x50, 0xe1, 0xeb, 0x22, 0xea, 0xbc, 0x80, 0x61, 0xbb, 0x16, 0x02,
	0x60, 0xbc, 0x9e, 0x2e, 0xae, 0x2e, 0xde, 0x59, 0x3b, 0x64, 0x00, 0xbd, 0xc5, 0xf4, 0x82, 0x4e,
	0x97, 0xe7, 0xef, 0x7f, 0xb1, 0x34, 0xe7, 0x02, 0x4c, 0x8a, 0x51, 0xcc, 0x71, 0x1e, 0xa2, 0x8f,
	0xe4, 0x29, 0xf4, 0x92, 0xe2, 0xc3, 0x65, 0x79, 0x24, 0x9a, 0xd6, 0x69, 0x57, 0x38, 0xae, 0xf3,
	0xa8, 0x18, 0x1f, 0x8b, 0x03, 0x74, 0xc3, 0x40, 0xd4, 0xde, 0xa3, 0x46, 0x61, 0xce, 0x02, 0xe7,
	0x77, 0x0d, 0x06, 0x32, 0xcb, 0x02, 0x57, 0x11, 0x32, 0x4e, 0xce, 0x00, 0xd2, 0x6a, 0x1d, 0x44,
	0x22, 0x73, 0xf2, 0xf4, 0x81, 0x5d, 0xa1, 0x0d, 0x38, 0x79, 0x02, 0xf2, 0xcd, 0xfa, 0xa1, 0x8e,
	0xb0, 0x67, 0x01, 0x39, 0x83, 0x41, 0x2a, 0x1e, 0x72, 0x85, 0x27, 0xb3, 0xf5, 0x63, 0xfd, 0xc4,
	0x9c, 0x1c, 0xb5, 0x52, 0x57, 0xed, 0xd0, 0x7e, 0x5a, 0x1b, 0x19, 0x79, 0x06, 0x66, 0x84, 0xe9,
	0xc7, 0x0d, 0xba, 0x69, 0x1c, 0x73, 0xb1, 0x4a, 0x7d, 0x0a, 0xd2, 0x45, 0xe3, 0x98, 0x3b, 0x7f,
	0xee, 0x42, 0x67, 0x2e, 0x13, 0x91, 0xd3, 0xd6, 0xe4, 0x9b, 0xb5, 0x2b, 0xc4, 0xf8, 0xd2, 0xe3,
	0x5e, 0x63, 0xd4, 0x5f, 0xc0, 0x30, 0x64, 0x9b, 0x90, 0xa1, 0x9b, 0x49, 0x11, 0xd4, 0x98, 0x06,
	0xd2, 0x5b, 0x2a, 0xf3, 0x2d, 0x18, 0xb2, 0x28, 0xf1, 0xbe, 0x39, 0xb1, 0xef, 0x95, 0xae, 0x90,
	0x54, 0xe1, 0x08, 0x81, 0xbd, 0x2c, 0xfc, 0x0d, 0xd5, 0x3a, 0x8b, 0x6f, 0xf2, 0x33, 0x0c, 0xfc,
	0x14, 0x3d, 0xb1, 0x4b, 0x81, 0xc7, 0xd1, 0x36, 0x44, 0xb2, 0xd1, 0x58, 0xde, 0xf4, 0xb8, 0xbc,
	0xe9, 0xf1, 0xb2, 0xbc, 0x69, 0xda, 0x2f, 0x09, 0x97, 0x1e, 0x47, 0x72, 0x01, 0x07, 0xf8, 0x6b,
	0x12, 0xa6, 0x8d, 0x14, 0x9d, 0xff, 0x4d, 0x31, 0xac, 0x29, 0x22, 0xc9, 0x08, 0xba, 0x11, 0x72,
	0x2f, 0xf0, 0xb8, 0x67, 0x77, 0x45, 0xb3, 0x95, 0xed, 0x38, 0xd0, 0x2d, 0x05, 0x2a, 0xf6, 0x6f,
	0x76, 0xfd, 0x76, 0x76, 0x3d, 0xb5, 0x76, 0x8a, 0x6f, 0x3a, 0x7d, 0xf7, 0x7e, 0x39, 0xb5, 0x34,
	0xe7, 0x03, 0xc0, 0x3c, 0xe7, 0x14, 0x6f, 0x73, 0xcc, 0x78, 0xd1, 0x67, 0xe2, 0xf1, 0xb5, 0x50,
	0xbc, 0x4f, 0xc5, 0x37, 0xf9, 0x06, 0x3a, 0x4a, 0x1e, 0xb1, 0x09, 0xe6, 0x84, 0xdc, 0x1f, 0x04,
	0x2d, 0x21, 0xe4, 0x08, 0x8c, 0xd7, 0xf3, 0xd9, 0x1b, 0xdc, 0x2a, 0xe9, 0x95, 0xe5, 0xbc, 0x02,
	0xb8, 0xc2, 0x07, 0xdf, 0xa9, 0x99, 0xbb, 0x2d, 0xe6, 0x0a, 0xcc, 0xb7, 0x61, 0x56, 0x51, 0x9f,
	0xc3, 0xa3, 0xea, 0x14, 0x0b, 0x9e, 0xb8, 0x63, 0x99, 0xe7, 0xa0, 0x0c, 0xcc, 0x3d, 0xbe, 0x2e,
	0x0e, 0xf8, 0x10, 0xf6, 0x37, 0x61, 0x14, 0x72, 0xf5, 0x57, 0x27, 0x8d, 0xff, 0x2c, 0x71, 0x00,
	0xa6, 0x90, 0x22, 0x4b, 0x62, 0x96, 0xa1, 0xf3, 0x15, 0x98, 0xa2, 0x62, 0x69, 0x12, 0xbb, 0x96,
	0x41, 0xbe, 0x56, 0x9a, 0xce, 0x39, 0xf4, 0x65, 0x81, 0x0a, 0x79, 0x08, 0xfb, 0x45, 0x61, 0x99,
	0xad, 0x1d, 0xeb, 0x27, 0x7d, 0x2a, 0x0d, 0xf2, 0x39, 0xf4, 0x78, 0x9a, 0x33, 0xdf, 0xe3, 0x28,
	0x4f, 0xaa, 0x4b, 0x6b, 0x87, 0x73, 0x06, 0x83, 0x4b, 0xdc, 0x20, 0xc7, 0x4f, 0x51, 0xc8, 0x82,
	0x61, 0x49, 0x96, 0x25, 0x4c, 0xfe, 0xd2, 0xa0, 0xa7, 0x46, 0x73, 0x79, 0x4e, 0x5e, 0x82, 0x3e,
	0xcf, 0x39, 0xf9, 0xac, 0x39, 0xb7, 0x6a, 0xe6, 0xa3, 0xa3, 0xbb, 0x6e, 0xd5, 0xc6, 0x4b, 0xd0,
	0xaf, 0xb0, 0xcd, 0xaa, 0x27, 0xd8, 0x62, 0x35, 0x65, 0xfa, 0x1e, 0xf6, 0x0a, 0x31, 0x48, 0x33,
	0xde, 0x18, 0xdf, 0xe8, 0xf1, 0x3d, 0xbf, 0x22, 0xfe, 0x08, 0x86, 0x6c, 0x82, 0x34, 0xcf, 0xb1,
	0x25, 0xca, 0xe8, 0xc9, 0xbf, 0x44, 0x24, 0xfd, 0xc6, 0x10, 0xb7, 0xf2, 0xdd, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x56, 0x72, 0xae, 0x33, 0x6e, 0x07, 0x00, 0x00,
}