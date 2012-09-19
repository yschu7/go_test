// Code generated by protoc-gen-go from "mail_service.proto"
// DO NOT EDIT!

package appengine

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type MailServiceError_ErrorCode int32

const (
	MailServiceError_OK                      MailServiceError_ErrorCode = 0
	MailServiceError_INTERNAL_ERROR          MailServiceError_ErrorCode = 1
	MailServiceError_BAD_REQUEST             MailServiceError_ErrorCode = 2
	MailServiceError_UNAUTHORIZED_SENDER     MailServiceError_ErrorCode = 3
	MailServiceError_INVALID_ATTACHMENT_TYPE MailServiceError_ErrorCode = 4
	MailServiceError_INVALID_HEADER_NAME     MailServiceError_ErrorCode = 5
)

var MailServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "BAD_REQUEST",
	3: "UNAUTHORIZED_SENDER",
	4: "INVALID_ATTACHMENT_TYPE",
	5: "INVALID_HEADER_NAME",
}
var MailServiceError_ErrorCode_value = map[string]int32{
	"OK":                      0,
	"INTERNAL_ERROR":          1,
	"BAD_REQUEST":             2,
	"UNAUTHORIZED_SENDER":     3,
	"INVALID_ATTACHMENT_TYPE": 4,
	"INVALID_HEADER_NAME":     5,
}

func (x MailServiceError_ErrorCode) Enum() *MailServiceError_ErrorCode {
	p := new(MailServiceError_ErrorCode)
	*p = x
	return p
}
func (x MailServiceError_ErrorCode) String() string {
	return proto.EnumName(MailServiceError_ErrorCode_name, int32(x))
}
func (x MailServiceError_ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *MailServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MailServiceError_ErrorCode_value, data, "MailServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = MailServiceError_ErrorCode(value)
	return nil
}

type MailServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *MailServiceError) Reset()         { *this = MailServiceError{} }
func (this *MailServiceError) String() string { return proto.CompactTextString(this) }
func (*MailServiceError) ProtoMessage()       {}

type MailAttachment struct {
	FileName         *string `protobuf:"bytes,1,req" json:"FileName,omitempty"`
	Data             []byte  `protobuf:"bytes,2,req" json:"Data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *MailAttachment) Reset()         { *this = MailAttachment{} }
func (this *MailAttachment) String() string { return proto.CompactTextString(this) }
func (*MailAttachment) ProtoMessage()       {}

func (this *MailAttachment) GetFileName() string {
	if this != nil && this.FileName != nil {
		return *this.FileName
	}
	return ""
}

func (this *MailAttachment) GetData() []byte {
	if this != nil {
		return this.Data
	}
	return nil
}

type MailHeader struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *MailHeader) Reset()         { *this = MailHeader{} }
func (this *MailHeader) String() string { return proto.CompactTextString(this) }
func (*MailHeader) ProtoMessage()       {}

func (this *MailHeader) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *MailHeader) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

type MailMessage struct {
	Sender           *string           `protobuf:"bytes,1,req" json:"Sender,omitempty"`
	ReplyTo          *string           `protobuf:"bytes,2,opt" json:"ReplyTo,omitempty"`
	To               []string          `protobuf:"bytes,3,rep" json:"To,omitempty"`
	Cc               []string          `protobuf:"bytes,4,rep" json:"Cc,omitempty"`
	Bcc              []string          `protobuf:"bytes,5,rep" json:"Bcc,omitempty"`
	Subject          *string           `protobuf:"bytes,6,req" json:"Subject,omitempty"`
	TextBody         *string           `protobuf:"bytes,7,opt" json:"TextBody,omitempty"`
	HtmlBody         *string           `protobuf:"bytes,8,opt" json:"HtmlBody,omitempty"`
	Attachment       []*MailAttachment `protobuf:"bytes,9,rep" json:"Attachment,omitempty"`
	Header           []*MailHeader     `protobuf:"bytes,10,rep" json:"Header,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (this *MailMessage) Reset()         { *this = MailMessage{} }
func (this *MailMessage) String() string { return proto.CompactTextString(this) }
func (*MailMessage) ProtoMessage()       {}

func (this *MailMessage) GetSender() string {
	if this != nil && this.Sender != nil {
		return *this.Sender
	}
	return ""
}

func (this *MailMessage) GetReplyTo() string {
	if this != nil && this.ReplyTo != nil {
		return *this.ReplyTo
	}
	return ""
}

func (this *MailMessage) GetSubject() string {
	if this != nil && this.Subject != nil {
		return *this.Subject
	}
	return ""
}

func (this *MailMessage) GetTextBody() string {
	if this != nil && this.TextBody != nil {
		return *this.TextBody
	}
	return ""
}

func (this *MailMessage) GetHtmlBody() string {
	if this != nil && this.HtmlBody != nil {
		return *this.HtmlBody
	}
	return ""
}

func init() {
	proto.RegisterEnum("appengine.MailServiceError_ErrorCode", MailServiceError_ErrorCode_name, MailServiceError_ErrorCode_value)
}