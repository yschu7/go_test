// Code generated by protoc-gen-go from "urlfetch_service.proto"
// DO NOT EDIT!

package appengine

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type URLFetchServiceError_ErrorCode int32

const (
	URLFetchServiceError_OK                       URLFetchServiceError_ErrorCode = 0
	URLFetchServiceError_INVALID_URL              URLFetchServiceError_ErrorCode = 1
	URLFetchServiceError_FETCH_ERROR              URLFetchServiceError_ErrorCode = 2
	URLFetchServiceError_UNSPECIFIED_ERROR        URLFetchServiceError_ErrorCode = 3
	URLFetchServiceError_RESPONSE_TOO_LARGE       URLFetchServiceError_ErrorCode = 4
	URLFetchServiceError_DEADLINE_EXCEEDED        URLFetchServiceError_ErrorCode = 5
	URLFetchServiceError_SSL_CERTIFICATE_ERROR    URLFetchServiceError_ErrorCode = 6
	URLFetchServiceError_DNS_ERROR                URLFetchServiceError_ErrorCode = 7
	URLFetchServiceError_CLOSED                   URLFetchServiceError_ErrorCode = 8
	URLFetchServiceError_INTERNAL_TRANSIENT_ERROR URLFetchServiceError_ErrorCode = 9
	URLFetchServiceError_TOO_MANY_REDIRECTS       URLFetchServiceError_ErrorCode = 10
	URLFetchServiceError_MALFORMED_REPLY          URLFetchServiceError_ErrorCode = 11
	URLFetchServiceError_CONNECTION_ERROR         URLFetchServiceError_ErrorCode = 12
)

var URLFetchServiceError_ErrorCode_name = map[int32]string{
	0:  "OK",
	1:  "INVALID_URL",
	2:  "FETCH_ERROR",
	3:  "UNSPECIFIED_ERROR",
	4:  "RESPONSE_TOO_LARGE",
	5:  "DEADLINE_EXCEEDED",
	6:  "SSL_CERTIFICATE_ERROR",
	7:  "DNS_ERROR",
	8:  "CLOSED",
	9:  "INTERNAL_TRANSIENT_ERROR",
	10: "TOO_MANY_REDIRECTS",
	11: "MALFORMED_REPLY",
	12: "CONNECTION_ERROR",
}
var URLFetchServiceError_ErrorCode_value = map[string]int32{
	"OK":                       0,
	"INVALID_URL":              1,
	"FETCH_ERROR":              2,
	"UNSPECIFIED_ERROR":        3,
	"RESPONSE_TOO_LARGE":       4,
	"DEADLINE_EXCEEDED":        5,
	"SSL_CERTIFICATE_ERROR":    6,
	"DNS_ERROR":                7,
	"CLOSED":                   8,
	"INTERNAL_TRANSIENT_ERROR": 9,
	"TOO_MANY_REDIRECTS":       10,
	"MALFORMED_REPLY":          11,
	"CONNECTION_ERROR":         12,
}

func (x URLFetchServiceError_ErrorCode) Enum() *URLFetchServiceError_ErrorCode {
	p := new(URLFetchServiceError_ErrorCode)
	*p = x
	return p
}
func (x URLFetchServiceError_ErrorCode) String() string {
	return proto.EnumName(URLFetchServiceError_ErrorCode_name, int32(x))
}
func (x URLFetchServiceError_ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *URLFetchServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(URLFetchServiceError_ErrorCode_value, data, "URLFetchServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = URLFetchServiceError_ErrorCode(value)
	return nil
}

type URLFetchRequest_RequestMethod int32

const (
	URLFetchRequest_GET    URLFetchRequest_RequestMethod = 1
	URLFetchRequest_POST   URLFetchRequest_RequestMethod = 2
	URLFetchRequest_HEAD   URLFetchRequest_RequestMethod = 3
	URLFetchRequest_PUT    URLFetchRequest_RequestMethod = 4
	URLFetchRequest_DELETE URLFetchRequest_RequestMethod = 5
)

var URLFetchRequest_RequestMethod_name = map[int32]string{
	1: "GET",
	2: "POST",
	3: "HEAD",
	4: "PUT",
	5: "DELETE",
}
var URLFetchRequest_RequestMethod_value = map[string]int32{
	"GET":    1,
	"POST":   2,
	"HEAD":   3,
	"PUT":    4,
	"DELETE": 5,
}

func (x URLFetchRequest_RequestMethod) Enum() *URLFetchRequest_RequestMethod {
	p := new(URLFetchRequest_RequestMethod)
	*p = x
	return p
}
func (x URLFetchRequest_RequestMethod) String() string {
	return proto.EnumName(URLFetchRequest_RequestMethod_name, int32(x))
}
func (x URLFetchRequest_RequestMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *URLFetchRequest_RequestMethod) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(URLFetchRequest_RequestMethod_value, data, "URLFetchRequest_RequestMethod")
	if err != nil {
		return err
	}
	*x = URLFetchRequest_RequestMethod(value)
	return nil
}

type URLFetchServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *URLFetchServiceError) Reset()         { *this = URLFetchServiceError{} }
func (this *URLFetchServiceError) String() string { return proto.CompactTextString(this) }
func (*URLFetchServiceError) ProtoMessage()       {}

type URLFetchRequest struct {
	Method                        *URLFetchRequest_RequestMethod `protobuf:"varint,1,req,enum=appengine.URLFetchRequest_RequestMethod" json:"Method,omitempty"`
	Url                           *string                        `protobuf:"bytes,2,req" json:"Url,omitempty"`
	Header                        []*URLFetchRequest_Header      `protobuf:"group,3,rep" json:"header,omitempty"`
	Payload                       []byte                         `protobuf:"bytes,6,opt" json:"Payload,omitempty"`
	FollowRedirects               *bool                          `protobuf:"varint,7,opt,def=1" json:"FollowRedirects,omitempty"`
	Deadline                      *float64                       `protobuf:"fixed64,8,opt" json:"Deadline,omitempty"`
	MustValidateServerCertificate *bool                          `protobuf:"varint,9,opt,def=1" json:"MustValidateServerCertificate,omitempty"`
	XXX_unrecognized              []byte                         `json:"-"`
}

func (this *URLFetchRequest) Reset()         { *this = URLFetchRequest{} }
func (this *URLFetchRequest) String() string { return proto.CompactTextString(this) }
func (*URLFetchRequest) ProtoMessage()       {}

const Default_URLFetchRequest_FollowRedirects bool = true
const Default_URLFetchRequest_MustValidateServerCertificate bool = true

func (this *URLFetchRequest) GetMethod() URLFetchRequest_RequestMethod {
	if this != nil && this.Method != nil {
		return *this.Method
	}
	return 0
}

func (this *URLFetchRequest) GetUrl() string {
	if this != nil && this.Url != nil {
		return *this.Url
	}
	return ""
}

func (this *URLFetchRequest) GetPayload() []byte {
	if this != nil {
		return this.Payload
	}
	return nil
}

func (this *URLFetchRequest) GetFollowRedirects() bool {
	if this != nil && this.FollowRedirects != nil {
		return *this.FollowRedirects
	}
	return Default_URLFetchRequest_FollowRedirects
}

func (this *URLFetchRequest) GetDeadline() float64 {
	if this != nil && this.Deadline != nil {
		return *this.Deadline
	}
	return 0
}

func (this *URLFetchRequest) GetMustValidateServerCertificate() bool {
	if this != nil && this.MustValidateServerCertificate != nil {
		return *this.MustValidateServerCertificate
	}
	return Default_URLFetchRequest_MustValidateServerCertificate
}

type URLFetchRequest_Header struct {
	Key              *string `protobuf:"bytes,4,req" json:"Key,omitempty"`
	Value            *string `protobuf:"bytes,5,req" json:"Value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *URLFetchRequest_Header) Reset()         { *this = URLFetchRequest_Header{} }
func (this *URLFetchRequest_Header) String() string { return proto.CompactTextString(this) }
func (*URLFetchRequest_Header) ProtoMessage()       {}

func (this *URLFetchRequest_Header) GetKey() string {
	if this != nil && this.Key != nil {
		return *this.Key
	}
	return ""
}

func (this *URLFetchRequest_Header) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

type URLFetchResponse struct {
	Content               []byte                     `protobuf:"bytes,1,opt" json:"Content,omitempty"`
	StatusCode            *int32                     `protobuf:"varint,2,req" json:"StatusCode,omitempty"`
	Header                []*URLFetchResponse_Header `protobuf:"group,3,rep" json:"header,omitempty"`
	ContentWasTruncated   *bool                      `protobuf:"varint,6,opt,def=0" json:"ContentWasTruncated,omitempty"`
	ExternalBytesSent     *int64                     `protobuf:"varint,7,opt" json:"ExternalBytesSent,omitempty"`
	ExternalBytesReceived *int64                     `protobuf:"varint,8,opt" json:"ExternalBytesReceived,omitempty"`
	FinalUrl              *string                    `protobuf:"bytes,9,opt" json:"FinalUrl,omitempty"`
	ApiCpuMilliseconds    *int64                     `protobuf:"varint,10,opt,def=0" json:"ApiCpuMilliseconds,omitempty"`
	ApiBytesSent          *int64                     `protobuf:"varint,11,opt,def=0" json:"ApiBytesSent,omitempty"`
	ApiBytesReceived      *int64                     `protobuf:"varint,12,opt,def=0" json:"ApiBytesReceived,omitempty"`
	XXX_unrecognized      []byte                     `json:"-"`
}

func (this *URLFetchResponse) Reset()         { *this = URLFetchResponse{} }
func (this *URLFetchResponse) String() string { return proto.CompactTextString(this) }
func (*URLFetchResponse) ProtoMessage()       {}

const Default_URLFetchResponse_ContentWasTruncated bool = false
const Default_URLFetchResponse_ApiCpuMilliseconds int64 = 0
const Default_URLFetchResponse_ApiBytesSent int64 = 0
const Default_URLFetchResponse_ApiBytesReceived int64 = 0

func (this *URLFetchResponse) GetContent() []byte {
	if this != nil {
		return this.Content
	}
	return nil
}

func (this *URLFetchResponse) GetStatusCode() int32 {
	if this != nil && this.StatusCode != nil {
		return *this.StatusCode
	}
	return 0
}

func (this *URLFetchResponse) GetContentWasTruncated() bool {
	if this != nil && this.ContentWasTruncated != nil {
		return *this.ContentWasTruncated
	}
	return Default_URLFetchResponse_ContentWasTruncated
}

func (this *URLFetchResponse) GetExternalBytesSent() int64 {
	if this != nil && this.ExternalBytesSent != nil {
		return *this.ExternalBytesSent
	}
	return 0
}

func (this *URLFetchResponse) GetExternalBytesReceived() int64 {
	if this != nil && this.ExternalBytesReceived != nil {
		return *this.ExternalBytesReceived
	}
	return 0
}

func (this *URLFetchResponse) GetFinalUrl() string {
	if this != nil && this.FinalUrl != nil {
		return *this.FinalUrl
	}
	return ""
}

func (this *URLFetchResponse) GetApiCpuMilliseconds() int64 {
	if this != nil && this.ApiCpuMilliseconds != nil {
		return *this.ApiCpuMilliseconds
	}
	return Default_URLFetchResponse_ApiCpuMilliseconds
}

func (this *URLFetchResponse) GetApiBytesSent() int64 {
	if this != nil && this.ApiBytesSent != nil {
		return *this.ApiBytesSent
	}
	return Default_URLFetchResponse_ApiBytesSent
}

func (this *URLFetchResponse) GetApiBytesReceived() int64 {
	if this != nil && this.ApiBytesReceived != nil {
		return *this.ApiBytesReceived
	}
	return Default_URLFetchResponse_ApiBytesReceived
}

type URLFetchResponse_Header struct {
	Key              *string `protobuf:"bytes,4,req" json:"Key,omitempty"`
	Value            *string `protobuf:"bytes,5,req" json:"Value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *URLFetchResponse_Header) Reset()         { *this = URLFetchResponse_Header{} }
func (this *URLFetchResponse_Header) String() string { return proto.CompactTextString(this) }
func (*URLFetchResponse_Header) ProtoMessage()       {}

func (this *URLFetchResponse_Header) GetKey() string {
	if this != nil && this.Key != nil {
		return *this.Key
	}
	return ""
}

func (this *URLFetchResponse_Header) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("appengine.URLFetchServiceError_ErrorCode", URLFetchServiceError_ErrorCode_name, URLFetchServiceError_ErrorCode_value)
	proto.RegisterEnum("appengine.URLFetchRequest_RequestMethod", URLFetchRequest_RequestMethod_name, URLFetchRequest_RequestMethod_value)
}
