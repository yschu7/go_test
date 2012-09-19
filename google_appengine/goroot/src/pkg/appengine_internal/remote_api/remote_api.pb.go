// Code generated by protoc-gen-go from "remote_api.proto"
// DO NOT EDIT!

package remote_api

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	ServiceName      *string `protobuf:"bytes,2,req,name=service_name" json:"service_name,omitempty"`
	Method           *string `protobuf:"bytes,3,req,name=method" json:"method,omitempty"`
	Request          []byte  `protobuf:"bytes,4,req,name=request" json:"request,omitempty"`
	RequestId        *string `protobuf:"bytes,5,opt,name=request_id" json:"request_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Request) Reset()         { *this = Request{} }
func (this *Request) String() string { return proto.CompactTextString(this) }
func (*Request) ProtoMessage()       {}

func (this *Request) GetServiceName() string {
	if this != nil && this.ServiceName != nil {
		return *this.ServiceName
	}
	return ""
}

func (this *Request) GetMethod() string {
	if this != nil && this.Method != nil {
		return *this.Method
	}
	return ""
}

func (this *Request) GetRequest() []byte {
	if this != nil {
		return this.Request
	}
	return nil
}

func (this *Request) GetRequestId() string {
	if this != nil && this.RequestId != nil {
		return *this.RequestId
	}
	return ""
}

type ApplicationError struct {
	Code             *int32  `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Detail           *string `protobuf:"bytes,2,req,name=detail" json:"detail,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ApplicationError) Reset()         { *this = ApplicationError{} }
func (this *ApplicationError) String() string { return proto.CompactTextString(this) }
func (*ApplicationError) ProtoMessage()       {}

func (this *ApplicationError) GetCode() int32 {
	if this != nil && this.Code != nil {
		return *this.Code
	}
	return 0
}

func (this *ApplicationError) GetDetail() string {
	if this != nil && this.Detail != nil {
		return *this.Detail
	}
	return ""
}

type Response struct {
	Response         []byte            `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Exception        []byte            `protobuf:"bytes,2,opt,name=exception" json:"exception,omitempty"`
	ApplicationError *ApplicationError `protobuf:"bytes,3,opt,name=application_error" json:"application_error,omitempty"`
	JavaException    []byte            `protobuf:"bytes,4,opt,name=java_exception" json:"java_exception,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (this *Response) Reset()         { *this = Response{} }
func (this *Response) String() string { return proto.CompactTextString(this) }
func (*Response) ProtoMessage()       {}

func (this *Response) GetResponse() []byte {
	if this != nil {
		return this.Response
	}
	return nil
}

func (this *Response) GetException() []byte {
	if this != nil {
		return this.Exception
	}
	return nil
}

func (this *Response) GetApplicationError() *ApplicationError {
	if this != nil {
		return this.ApplicationError
	}
	return nil
}

func (this *Response) GetJavaException() []byte {
	if this != nil {
		return this.JavaException
	}
	return nil
}

func init() {
}
