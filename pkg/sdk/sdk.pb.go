// Copyright 2019 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sdk.proto

package sdk

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// I am Empty
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_d5ced4af94219733, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// Key, Value entry
type KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_d5ced4af94219733, []int{1}
}
func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (dst *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(dst, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// A GameServer Custom Resource Definition object
// We will only export those resources that make the most
// sense. Can always expand to more as needed.
type GameServer struct {
	ObjectMeta           *GameServer_ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Spec                 *GameServer_Spec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status               *GameServer_Status     `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GameServer) Reset()         { *m = GameServer{} }
func (m *GameServer) String() string { return proto.CompactTextString(m) }
func (*GameServer) ProtoMessage()    {}
func (*GameServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_d5ced4af94219733, []int{2}
}
func (m *GameServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer.Unmarshal(m, b)
}
func (m *GameServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer.Marshal(b, m, deterministic)
}
func (dst *GameServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer.Merge(dst, src)
}
func (m *GameServer) XXX_Size() int {
	return xxx_messageInfo_GameServer.Size(m)
}
func (m *GameServer) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer proto.InternalMessageInfo

func (m *GameServer) GetObjectMeta() *GameServer_ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *GameServer) GetSpec() *GameServer_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *GameServer) GetStatus() *GameServer_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// representation of the K8s ObjectMeta resource
type GameServer_ObjectMeta struct {
	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace       string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid             string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	ResourceVersion string `protobuf:"bytes,4,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	Generation      int64  `protobuf:"varint,5,opt,name=generation,proto3" json:"generation,omitempty"`
	// timestamp is in Epoch format, unit: seconds
	CreationTimestamp int64 `protobuf:"varint,6,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`
	// optional deletion timestamp in Epoch format, unit: seconds
	DeletionTimestamp    int64             `protobuf:"varint,7,opt,name=deletion_timestamp,json=deletionTimestamp,proto3" json:"deletion_timestamp,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,8,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GameServer_ObjectMeta) Reset()         { *m = GameServer_ObjectMeta{} }
func (m *GameServer_ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*GameServer_ObjectMeta) ProtoMessage()    {}
func (*GameServer_ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_d5ced4af94219733, []int{2, 0}
}
func (m *GameServer_ObjectMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_ObjectMeta.Unmarshal(m, b)
}
func (m *GameServer_ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_ObjectMeta.Marshal(b, m, deterministic)
}
func (dst *GameServer_ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_ObjectMeta.Merge(dst, src)
}
func (m *GameServer_ObjectMeta) XXX_Size() int {
	return xxx_messageInfo_GameServer_ObjectMeta.Size(m)
}
func (m *GameServer_ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_ObjectMeta proto.InternalMessageInfo

func (m *GameServer_ObjectMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameServer_ObjectMeta) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GameServer_ObjectMeta) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GameServer_ObjectMeta) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *GameServer_ObjectMeta) GetGeneration() int64 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func (m *GameServer_ObjectMeta) GetCreationTimestamp() int64 {
	if m != nil {
		return m.CreationTimestamp
	}
	return 0
}

func (m *GameServer_ObjectMeta) GetDeletionTimestamp() int64 {
	if m != nil {
		return m.DeletionTimestamp
	}
	return 0
}

func (m *GameServer_ObjectMeta) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *GameServer_ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type GameServer_Spec struct {
	Health               *GameServer_Spec_Health `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GameServer_Spec) Reset()         { *m = GameServer_Spec{} }
func (m *GameServer_Spec) String() string { return proto.CompactTextString(m) }
func (*GameServer_Spec) ProtoMessage()    {}
func (*GameServer_Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_d5ced4af94219733, []int{2, 1}
}
func (m *GameServer_Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_Spec.Unmarshal(m, b)
}
func (m *GameServer_Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_Spec.Marshal(b, m, deterministic)
}
func (dst *GameServer_Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_Spec.Merge(dst, src)
}
func (m *GameServer_Spec) XXX_Size() int {
	return xxx_messageInfo_GameServer_Spec.Size(m)
}
func (m *GameServer_Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_Spec proto.InternalMessageInfo

func (m *GameServer_Spec) GetHealth() *GameServer_Spec_Health {
	if m != nil {
		return m.Health
	}
	return nil
}

type GameServer_Spec_Health struct {
	Disabled             bool     `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	PeriodSeconds        int32    `protobuf:"varint,2,opt,name=period_seconds,json=periodSeconds,proto3" json:"period_seconds,omitempty"`
	FailureThreshold     int32    `protobuf:"varint,3,opt,name=failure_threshold,json=failureThreshold,proto3" json:"failure_threshold,omitempty"`
	InitialDelaySeconds  int32    `protobuf:"varint,4,opt,name=initial_delay_seconds,json=initialDelaySeconds,proto3" json:"initial_delay_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameServer_Spec_Health) Reset()         { *m = GameServer_Spec_Health{} }
func (m *GameServer_Spec_Health) String() string { return proto.CompactTextString(m) }
func (*GameServer_Spec_Health) ProtoMessage()    {}
func (*GameServer_Spec_Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_d5ced4af94219733, []int{2, 1, 0}
}
func (m *GameServer_Spec_Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_Spec_Health.Unmarshal(m, b)
}
func (m *GameServer_Spec_Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_Spec_Health.Marshal(b, m, deterministic)
}
func (dst *GameServer_Spec_Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_Spec_Health.Merge(dst, src)
}
func (m *GameServer_Spec_Health) XXX_Size() int {
	return xxx_messageInfo_GameServer_Spec_Health.Size(m)
}
func (m *GameServer_Spec_Health) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_Spec_Health.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_Spec_Health proto.InternalMessageInfo

func (m *GameServer_Spec_Health) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *GameServer_Spec_Health) GetPeriodSeconds() int32 {
	if m != nil {
		return m.PeriodSeconds
	}
	return 0
}

func (m *GameServer_Spec_Health) GetFailureThreshold() int32 {
	if m != nil {
		return m.FailureThreshold
	}
	return 0
}

func (m *GameServer_Spec_Health) GetInitialDelaySeconds() int32 {
	if m != nil {
		return m.InitialDelaySeconds
	}
	return 0
}

type GameServer_Status struct {
	State                string                    `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Address              string                    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Ports                []*GameServer_Status_Port `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GameServer_Status) Reset()         { *m = GameServer_Status{} }
func (m *GameServer_Status) String() string { return proto.CompactTextString(m) }
func (*GameServer_Status) ProtoMessage()    {}
func (*GameServer_Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_d5ced4af94219733, []int{2, 2}
}
func (m *GameServer_Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_Status.Unmarshal(m, b)
}
func (m *GameServer_Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_Status.Marshal(b, m, deterministic)
}
func (dst *GameServer_Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_Status.Merge(dst, src)
}
func (m *GameServer_Status) XXX_Size() int {
	return xxx_messageInfo_GameServer_Status.Size(m)
}
func (m *GameServer_Status) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_Status.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_Status proto.InternalMessageInfo

func (m *GameServer_Status) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *GameServer_Status) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GameServer_Status) GetPorts() []*GameServer_Status_Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type GameServer_Status_Port struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameServer_Status_Port) Reset()         { *m = GameServer_Status_Port{} }
func (m *GameServer_Status_Port) String() string { return proto.CompactTextString(m) }
func (*GameServer_Status_Port) ProtoMessage()    {}
func (*GameServer_Status_Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_d5ced4af94219733, []int{2, 2, 0}
}
func (m *GameServer_Status_Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_Status_Port.Unmarshal(m, b)
}
func (m *GameServer_Status_Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_Status_Port.Marshal(b, m, deterministic)
}
func (dst *GameServer_Status_Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_Status_Port.Merge(dst, src)
}
func (m *GameServer_Status_Port) XXX_Size() int {
	return xxx_messageInfo_GameServer_Status_Port.Size(m)
}
func (m *GameServer_Status_Port) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_Status_Port.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_Status_Port proto.InternalMessageInfo

func (m *GameServer_Status_Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameServer_Status_Port) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "stable.agones.dev.sdk.Empty")
	proto.RegisterType((*KeyValue)(nil), "stable.agones.dev.sdk.KeyValue")
	proto.RegisterType((*GameServer)(nil), "stable.agones.dev.sdk.GameServer")
	proto.RegisterType((*GameServer_ObjectMeta)(nil), "stable.agones.dev.sdk.GameServer.ObjectMeta")
	proto.RegisterMapType((map[string]string)(nil), "stable.agones.dev.sdk.GameServer.ObjectMeta.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "stable.agones.dev.sdk.GameServer.ObjectMeta.LabelsEntry")
	proto.RegisterType((*GameServer_Spec)(nil), "stable.agones.dev.sdk.GameServer.Spec")
	proto.RegisterType((*GameServer_Spec_Health)(nil), "stable.agones.dev.sdk.GameServer.Spec.Health")
	proto.RegisterType((*GameServer_Status)(nil), "stable.agones.dev.sdk.GameServer.Status")
	proto.RegisterType((*GameServer_Status_Port)(nil), "stable.agones.dev.sdk.GameServer.Status.Port")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SDKClient is the client API for SDK service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SDKClient interface {
	// Call when the GameServer is ready
	Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Call to self Allocation the GameServer
	Allocate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Call when the GameServer is shutting down
	Shutdown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Send a Empty every d Duration to declare that this GameSever is healthy
	Health(ctx context.Context, opts ...grpc.CallOption) (SDK_HealthClient, error)
	// Retrieve the current GameServer data
	GetGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameServer, error)
	// Send GameServer details whenever the GameServer is updated
	WatchGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SDK_WatchGameServerClient, error)
	// Apply a Label to the backing GameServer metadata
	SetLabel(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error)
	// Apply a Annotation to the backing GameServer metadata
	SetAnnotation(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error)
}

type sDKClient struct {
	cc *grpc.ClientConn
}

func NewSDKClient(cc *grpc.ClientConn) SDKClient {
	return &sDKClient{cc}
}

func (c *sDKClient) Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/stable.agones.dev.sdk.SDK/Ready", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Allocate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/stable.agones.dev.sdk.SDK/Allocate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Shutdown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/stable.agones.dev.sdk.SDK/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Health(ctx context.Context, opts ...grpc.CallOption) (SDK_HealthClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SDK_serviceDesc.Streams[0], "/stable.agones.dev.sdk.SDK/Health", opts...)
	if err != nil {
		return nil, err
	}
	x := &sDKHealthClient{stream}
	return x, nil
}

type SDK_HealthClient interface {
	Send(*Empty) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type sDKHealthClient struct {
	grpc.ClientStream
}

func (x *sDKHealthClient) Send(m *Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sDKHealthClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sDKClient) GetGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameServer, error) {
	out := new(GameServer)
	err := c.cc.Invoke(ctx, "/stable.agones.dev.sdk.SDK/GetGameServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) WatchGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SDK_WatchGameServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SDK_serviceDesc.Streams[1], "/stable.agones.dev.sdk.SDK/WatchGameServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &sDKWatchGameServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SDK_WatchGameServerClient interface {
	Recv() (*GameServer, error)
	grpc.ClientStream
}

type sDKWatchGameServerClient struct {
	grpc.ClientStream
}

func (x *sDKWatchGameServerClient) Recv() (*GameServer, error) {
	m := new(GameServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sDKClient) SetLabel(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/stable.agones.dev.sdk.SDK/SetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) SetAnnotation(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/stable.agones.dev.sdk.SDK/SetAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SDKServer is the server API for SDK service.
type SDKServer interface {
	// Call when the GameServer is ready
	Ready(context.Context, *Empty) (*Empty, error)
	// Call to self Allocation the GameServer
	Allocate(context.Context, *Empty) (*Empty, error)
	// Call when the GameServer is shutting down
	Shutdown(context.Context, *Empty) (*Empty, error)
	// Send a Empty every d Duration to declare that this GameSever is healthy
	Health(SDK_HealthServer) error
	// Retrieve the current GameServer data
	GetGameServer(context.Context, *Empty) (*GameServer, error)
	// Send GameServer details whenever the GameServer is updated
	WatchGameServer(*Empty, SDK_WatchGameServerServer) error
	// Apply a Label to the backing GameServer metadata
	SetLabel(context.Context, *KeyValue) (*Empty, error)
	// Apply a Annotation to the backing GameServer metadata
	SetAnnotation(context.Context, *KeyValue) (*Empty, error)
}

func RegisterSDKServer(s *grpc.Server, srv SDKServer) {
	s.RegisterService(&_SDK_serviceDesc, srv)
}

func _SDK_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stable.agones.dev.sdk.SDK/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Ready(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stable.agones.dev.sdk.SDK/Allocate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Allocate(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stable.agones.dev.sdk.SDK/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Shutdown(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Health_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SDKServer).Health(&sDKHealthServer{stream})
}

type SDK_HealthServer interface {
	SendAndClose(*Empty) error
	Recv() (*Empty, error)
	grpc.ServerStream
}

type sDKHealthServer struct {
	grpc.ServerStream
}

func (x *sDKHealthServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sDKHealthServer) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SDK_GetGameServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).GetGameServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stable.agones.dev.sdk.SDK/GetGameServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).GetGameServer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_WatchGameServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SDKServer).WatchGameServer(m, &sDKWatchGameServerServer{stream})
}

type SDK_WatchGameServerServer interface {
	Send(*GameServer) error
	grpc.ServerStream
}

type sDKWatchGameServerServer struct {
	grpc.ServerStream
}

func (x *sDKWatchGameServerServer) Send(m *GameServer) error {
	return x.ServerStream.SendMsg(m)
}

func _SDK_SetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).SetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stable.agones.dev.sdk.SDK/SetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).SetLabel(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_SetAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).SetAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stable.agones.dev.sdk.SDK/SetAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).SetAnnotation(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _SDK_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stable.agones.dev.sdk.SDK",
	HandlerType: (*SDKServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ready",
			Handler:    _SDK_Ready_Handler,
		},
		{
			MethodName: "Allocate",
			Handler:    _SDK_Allocate_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _SDK_Shutdown_Handler,
		},
		{
			MethodName: "GetGameServer",
			Handler:    _SDK_GetGameServer_Handler,
		},
		{
			MethodName: "SetLabel",
			Handler:    _SDK_SetLabel_Handler,
		},
		{
			MethodName: "SetAnnotation",
			Handler:    _SDK_SetAnnotation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Health",
			Handler:       _SDK_Health_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "WatchGameServer",
			Handler:       _SDK_WatchGameServer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sdk.proto",
}

func init() { proto.RegisterFile("sdk.proto", fileDescriptor_sdk_d5ced4af94219733) }

var fileDescriptor_sdk_d5ced4af94219733 = []byte{
	// 815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x51, 0x8f, 0xdb, 0x44,
	0x10, 0x96, 0x2f, 0xb6, 0x2f, 0x19, 0x73, 0xbd, 0xbb, 0xbd, 0xab, 0xe4, 0x5a, 0x27, 0x7a, 0x58,
	0x02, 0x85, 0x83, 0xda, 0x28, 0xbc, 0x94, 0x48, 0x20, 0x0a, 0x3d, 0x15, 0xa9, 0x54, 0x54, 0x4e,
	0x75, 0x48, 0x15, 0x52, 0xb4, 0xe7, 0x1d, 0x12, 0x13, 0xc7, 0x6b, 0xed, 0x6e, 0x52, 0xe5, 0x95,
	0xbf, 0xc0, 0x6f, 0x40, 0xbc, 0xc0, 0x0b, 0x7f, 0x85, 0xbf, 0xc0, 0xcf, 0xe0, 0x01, 0xed, 0xda,
	0x8e, 0xa3, 0xaa, 0xed, 0xe5, 0x74, 0xf7, 0x94, 0xd9, 0x99, 0xf9, 0xbe, 0x6f, 0x3d, 0x3b, 0x3b,
	0x1b, 0xe8, 0x49, 0x36, 0x8b, 0x4a, 0xc1, 0x15, 0x27, 0x77, 0xa5, 0xa2, 0x97, 0x39, 0x46, 0x74,
	0xc2, 0x0b, 0x94, 0x11, 0xc3, 0x65, 0x24, 0xd9, 0x2c, 0x38, 0x99, 0x70, 0x3e, 0xc9, 0x31, 0xa6,
	0x65, 0x16, 0xd3, 0xa2, 0xe0, 0x8a, 0xaa, 0x8c, 0x17, 0xb2, 0x02, 0x85, 0xbb, 0xe0, 0x9c, 0xcf,
	0x4b, 0xb5, 0x0a, 0x07, 0xd0, 0x7d, 0x8a, 0xab, 0x0b, 0x9a, 0x2f, 0x90, 0x1c, 0x40, 0x67, 0x86,
	0x2b, 0xdf, 0x3a, 0xb5, 0xfa, 0xbd, 0x44, 0x9b, 0xe4, 0x18, 0x9c, 0xa5, 0x0e, 0xf9, 0x3b, 0xc6,
	0x57, 0x2d, 0xc2, 0xbf, 0x7b, 0x00, 0x4f, 0xe8, 0x1c, 0x47, 0x28, 0x96, 0x28, 0xc8, 0x33, 0xf0,
	0xf8, 0xe5, 0x2f, 0x98, 0xaa, 0xf1, 0x1c, 0x15, 0x35, 0x70, 0x6f, 0xf0, 0x69, 0xf4, 0xc6, 0x6d,
	0x45, 0x2d, 0x2e, 0xfa, 0xc1, 0x80, 0x9e, 0xa1, 0xa2, 0x09, 0xf0, 0xb5, 0x4d, 0x86, 0x60, 0xcb,
	0x12, 0x53, 0x23, 0xe9, 0x0d, 0x3e, 0xba, 0x9a, 0x67, 0x54, 0x62, 0x9a, 0x18, 0x0c, 0xf9, 0x1a,
	0x5c, 0xa9, 0xa8, 0x5a, 0x48, 0xbf, 0x63, 0xd0, 0xfd, 0x2d, 0xd0, 0x26, 0x3f, 0xa9, 0x71, 0xc1,
	0x1f, 0x36, 0x40, 0xbb, 0x31, 0x42, 0xc0, 0x2e, 0xe8, 0x1c, 0xeb, 0x9a, 0x18, 0x9b, 0x9c, 0x40,
	0x4f, 0xff, 0xca, 0x92, 0xa6, 0x4d, 0x61, 0x5a, 0x87, 0x2e, 0xe2, 0x22, 0x63, 0x46, 0xbf, 0x97,
	0x68, 0x93, 0x7c, 0x0c, 0x07, 0x02, 0x25, 0x5f, 0x88, 0x14, 0xc7, 0x4b, 0x14, 0x32, 0xe3, 0x85,
	0x6f, 0x9b, 0xf0, 0x7e, 0xe3, 0xbf, 0xa8, 0xdc, 0xe4, 0x7d, 0x80, 0x09, 0x16, 0x28, 0xcc, 0x59,
	0xf9, 0xce, 0xa9, 0xd5, 0xef, 0x24, 0x1b, 0x1e, 0xf2, 0x00, 0x48, 0x2a, 0xd0, 0xd8, 0x63, 0x95,
	0xcd, 0x51, 0x2a, 0x3a, 0x2f, 0x7d, 0xd7, 0xe4, 0x1d, 0x36, 0x91, 0x17, 0x4d, 0x40, 0xa7, 0x33,
	0xcc, 0xf1, 0xb5, 0xf4, 0xdd, 0x2a, 0xbd, 0x89, 0xb4, 0xe9, 0x63, 0xf0, 0x36, 0x3a, 0xc5, 0xef,
	0x9e, 0x76, 0xfa, 0xde, 0xe0, 0xcb, 0xeb, 0x1c, 0x64, 0xf4, 0xa8, 0xc5, 0x9f, 0x17, 0x4a, 0xac,
	0x92, 0x4d, 0x46, 0xf2, 0x1c, 0xdc, 0x9c, 0x5e, 0x62, 0x2e, 0xfd, 0x9e, 0xe1, 0x7e, 0x78, 0x2d,
	0xee, 0xef, 0x0d, 0xb4, 0xa2, 0xad, 0x79, 0x82, 0xaf, 0xe0, 0xe0, 0x75, 0xc9, 0x6d, 0xdb, 0x78,
	0xb8, 0xf3, 0xd0, 0x0a, 0xbe, 0x00, 0x6f, 0x83, 0xf6, 0x5a, 0xd0, 0xff, 0x2c, 0xb0, 0x75, 0xeb,
	0x91, 0x73, 0x70, 0xa7, 0x48, 0x73, 0x35, 0xad, 0x5b, 0xff, 0xc1, 0x76, 0x2d, 0x1b, 0x7d, 0x67,
	0x40, 0x49, 0x0d, 0x0e, 0xfe, 0xb4, 0xc0, 0xad, 0x5c, 0x24, 0x80, 0x2e, 0xcb, 0xa4, 0xe6, 0x60,
	0x86, 0xb3, 0x9b, 0xac, 0xd7, 0xe4, 0x43, 0xb8, 0x53, 0xa2, 0xc8, 0x38, 0x1b, 0x4b, 0x4c, 0x79,
	0xc1, 0xa4, 0xd9, 0x99, 0x93, 0xec, 0x55, 0xde, 0x51, 0xe5, 0x24, 0x9f, 0xc0, 0xe1, 0xcf, 0x34,
	0xcb, 0x17, 0x02, 0xc7, 0x6a, 0x2a, 0x50, 0x4e, 0x79, 0x5e, 0x35, 0xa5, 0x93, 0x1c, 0xd4, 0x81,
	0x17, 0x8d, 0x9f, 0x0c, 0xe0, 0x6e, 0x56, 0x64, 0x2a, 0xa3, 0xf9, 0x98, 0x61, 0x4e, 0x57, 0x6b,
	0x6a, 0xdb, 0x00, 0x8e, 0xea, 0xe0, 0x63, 0x1d, 0xab, 0x05, 0x82, 0xbf, 0x2c, 0x70, 0xab, 0xbb,
	0xa3, 0x6b, 0xa4, 0x6f, 0x4f, 0x73, 0x4b, 0xaa, 0x05, 0xf1, 0x61, 0x97, 0x32, 0x26, 0x50, 0xca,
	0xba, 0x76, 0xcd, 0x92, 0x7c, 0x0b, 0x4e, 0xc9, 0x85, 0xd2, 0x97, 0xb4, 0xb3, 0x65, 0xbd, 0x8c,
	0x50, 0xf4, 0x9c, 0x0b, 0x95, 0x54, 0xd8, 0x20, 0x02, 0x5b, 0x2f, 0xdf, 0x78, 0x43, 0x09, 0xd8,
	0x3a, 0xa9, 0xae, 0x8c, 0xb1, 0x07, 0xbf, 0xbb, 0xd0, 0x19, 0x3d, 0x7e, 0x4a, 0x2e, 0xc0, 0x49,
	0x90, 0xb2, 0x15, 0x39, 0x79, 0x8b, 0xac, 0x99, 0x8b, 0xc1, 0x3b, 0xa3, 0xe1, 0xe1, 0xaf, 0xff,
	0xfc, 0xfb, 0xdb, 0x8e, 0x17, 0xba, 0xb1, 0xd0, 0x5c, 0x43, 0xeb, 0x8c, 0xfc, 0x04, 0xdd, 0x47,
	0x79, 0xce, 0x53, 0xfd, 0xe9, 0x37, 0xa1, 0x3e, 0x36, 0xd4, 0x77, 0xc2, 0x5e, 0x4c, 0x6b, 0xba,
	0x9a, 0x7d, 0x34, 0x5d, 0x28, 0xc6, 0x5f, 0x15, 0xb7, 0xc4, 0x2e, 0x6b, 0x3a, 0xcd, 0xfe, 0x72,
	0xdd, 0x79, 0x37, 0xe1, 0x26, 0x86, 0xfb, 0xbd, 0x70, 0x37, 0xae, 0x3a, 0x7a, 0x68, 0x9d, 0xf5,
	0x2d, 0x82, 0xb0, 0xf7, 0x04, 0xd5, 0xc6, 0x73, 0xf1, 0x6e, 0x89, 0x0f, 0xae, 0x6c, 0x86, 0xf0,
	0xc8, 0xe8, 0xec, 0x11, 0x2f, 0x9e, 0xe8, 0xa9, 0x5b, 0xb1, 0x72, 0xd8, 0xff, 0x91, 0xaa, 0x74,
	0x7a, 0x9b, 0x42, 0xf7, 0x8c, 0xd0, 0x11, 0x39, 0x8c, 0x5f, 0x69, 0xea, 0x0d, 0xb9, 0xcf, 0xf4,
	0x77, 0x75, 0x47, 0xa8, 0xcc, 0xf0, 0x20, 0xf7, 0xdf, 0xc2, 0xd5, 0xbc, 0xac, 0x57, 0x14, 0x2e,
	0x30, 0x3a, 0xc7, 0xc1, 0x7e, 0xac, 0x1f, 0x50, 0x46, 0x15, 0x8d, 0xcd, 0x78, 0xd3, 0x47, 0xc3,
	0x61, 0x6f, 0x84, 0xaa, 0x9d, 0x71, 0x37, 0xd5, 0xba, 0x6f, 0xb4, 0xee, 0x05, 0xc7, 0xad, 0x56,
	0x3b, 0xa1, 0x87, 0xd6, 0xd9, 0x37, 0xce, 0xcb, 0x8e, 0x64, 0xb3, 0x4b, 0xd7, 0xfc, 0x4f, 0xf8,
	0xfc, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x06, 0x65, 0x87, 0x69, 0x08, 0x00, 0x00,
}
