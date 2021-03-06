// Code generated by protoc-gen-go from "api_base.proto"
// DO NOT EDIT!

package appengine_base

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type StringProto struct {
	Value            *string `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *StringProto) Reset()         { *this = StringProto{} }
func (this *StringProto) String() string { return proto.CompactTextString(this) }
func (*StringProto) ProtoMessage()       {}

func (this *StringProto) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

type Integer32Proto struct {
	Value            *int32 `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *Integer32Proto) Reset()         { *this = Integer32Proto{} }
func (this *Integer32Proto) String() string { return proto.CompactTextString(this) }
func (*Integer32Proto) ProtoMessage()       {}

func (this *Integer32Proto) GetValue() int32 {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return 0
}

type Integer64Proto struct {
	Value            *int64 `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *Integer64Proto) Reset()         { *this = Integer64Proto{} }
func (this *Integer64Proto) String() string { return proto.CompactTextString(this) }
func (*Integer64Proto) ProtoMessage()       {}

func (this *Integer64Proto) GetValue() int64 {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return 0
}

type BoolProto struct {
	Value            *bool  `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *BoolProto) Reset()         { *this = BoolProto{} }
func (this *BoolProto) String() string { return proto.CompactTextString(this) }
func (*BoolProto) ProtoMessage()       {}

func (this *BoolProto) GetValue() bool {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return false
}

type DoubleProto struct {
	Value            *float64 `protobuf:"fixed64,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *DoubleProto) Reset()         { *this = DoubleProto{} }
func (this *DoubleProto) String() string { return proto.CompactTextString(this) }
func (*DoubleProto) ProtoMessage()       {}

func (this *DoubleProto) GetValue() float64 {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return 0
}

type BytesProto struct {
	Value            []byte `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *BytesProto) Reset()         { *this = BytesProto{} }
func (this *BytesProto) String() string { return proto.CompactTextString(this) }
func (*BytesProto) ProtoMessage()       {}

func (this *BytesProto) GetValue() []byte {
	if this != nil {
		return this.Value
	}
	return nil
}

type VoidProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *VoidProto) Reset()         { *this = VoidProto{} }
func (this *VoidProto) String() string { return proto.CompactTextString(this) }
func (*VoidProto) ProtoMessage()       {}

func init() {
}
