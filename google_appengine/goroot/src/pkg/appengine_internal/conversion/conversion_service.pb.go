// Code generated by protoc-gen-go from "conversion_service.proto"
// DO NOT EDIT!

package appengine

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ConversionServiceError_ErrorCode int32

const (
	ConversionServiceError_OK                     ConversionServiceError_ErrorCode = 0
	ConversionServiceError_TIMEOUT                ConversionServiceError_ErrorCode = 1
	ConversionServiceError_TRANSIENT_ERROR        ConversionServiceError_ErrorCode = 2
	ConversionServiceError_INTERNAL_ERROR         ConversionServiceError_ErrorCode = 3
	ConversionServiceError_UNSUPPORTED_CONVERSION ConversionServiceError_ErrorCode = 4
	ConversionServiceError_CONVERSION_TOO_LARGE   ConversionServiceError_ErrorCode = 5
	ConversionServiceError_TOO_MANY_CONVERSIONS   ConversionServiceError_ErrorCode = 6
	ConversionServiceError_INVALID_REQUEST        ConversionServiceError_ErrorCode = 7
)

var ConversionServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "TIMEOUT",
	2: "TRANSIENT_ERROR",
	3: "INTERNAL_ERROR",
	4: "UNSUPPORTED_CONVERSION",
	5: "CONVERSION_TOO_LARGE",
	6: "TOO_MANY_CONVERSIONS",
	7: "INVALID_REQUEST",
}
var ConversionServiceError_ErrorCode_value = map[string]int32{
	"OK":                     0,
	"TIMEOUT":                1,
	"TRANSIENT_ERROR":        2,
	"INTERNAL_ERROR":         3,
	"UNSUPPORTED_CONVERSION": 4,
	"CONVERSION_TOO_LARGE":   5,
	"TOO_MANY_CONVERSIONS":   6,
	"INVALID_REQUEST":        7,
}

func (x ConversionServiceError_ErrorCode) Enum() *ConversionServiceError_ErrorCode {
	p := new(ConversionServiceError_ErrorCode)
	*p = x
	return p
}
func (x ConversionServiceError_ErrorCode) String() string {
	return proto.EnumName(ConversionServiceError_ErrorCode_name, int32(x))
}
func (x ConversionServiceError_ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *ConversionServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ConversionServiceError_ErrorCode_value, data, "ConversionServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = ConversionServiceError_ErrorCode(value)
	return nil
}

type ConversionServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *ConversionServiceError) Reset()         { *this = ConversionServiceError{} }
func (this *ConversionServiceError) String() string { return proto.CompactTextString(this) }
func (*ConversionServiceError) ProtoMessage()       {}

type AssetInfo struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	MimeType         *string `protobuf:"bytes,3,opt,name=mime_type" json:"mime_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *AssetInfo) Reset()         { *this = AssetInfo{} }
func (this *AssetInfo) String() string { return proto.CompactTextString(this) }
func (*AssetInfo) ProtoMessage()       {}

func (this *AssetInfo) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *AssetInfo) GetData() []byte {
	if this != nil {
		return this.Data
	}
	return nil
}

func (this *AssetInfo) GetMimeType() string {
	if this != nil && this.MimeType != nil {
		return *this.MimeType
	}
	return ""
}

type DocumentInfo struct {
	Asset            []*AssetInfo `protobuf:"bytes,1,rep,name=asset" json:"asset,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (this *DocumentInfo) Reset()         { *this = DocumentInfo{} }
func (this *DocumentInfo) String() string { return proto.CompactTextString(this) }
func (*DocumentInfo) ProtoMessage()       {}

type ConversionInput struct {
	Input            *DocumentInfo              `protobuf:"bytes,1,req,name=input" json:"input,omitempty"`
	OutputMimeType   *string                    `protobuf:"bytes,2,req,name=output_mime_type" json:"output_mime_type,omitempty"`
	Flag             []*ConversionInput_AuxData `protobuf:"bytes,3,rep,name=flag" json:"flag,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (this *ConversionInput) Reset()         { *this = ConversionInput{} }
func (this *ConversionInput) String() string { return proto.CompactTextString(this) }
func (*ConversionInput) ProtoMessage()       {}

func (this *ConversionInput) GetInput() *DocumentInfo {
	if this != nil {
		return this.Input
	}
	return nil
}

func (this *ConversionInput) GetOutputMimeType() string {
	if this != nil && this.OutputMimeType != nil {
		return *this.OutputMimeType
	}
	return ""
}

type ConversionInput_AuxData struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ConversionInput_AuxData) Reset()         { *this = ConversionInput_AuxData{} }
func (this *ConversionInput_AuxData) String() string { return proto.CompactTextString(this) }
func (*ConversionInput_AuxData) ProtoMessage()       {}

func (this *ConversionInput_AuxData) GetKey() string {
	if this != nil && this.Key != nil {
		return *this.Key
	}
	return ""
}

func (this *ConversionInput_AuxData) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

type ConversionOutput struct {
	ErrorCode        *ConversionServiceError_ErrorCode `protobuf:"varint,1,req,name=error_code,enum=appengine.ConversionServiceError_ErrorCode" json:"error_code,omitempty"`
	Output           *DocumentInfo                     `protobuf:"bytes,2,opt,name=output" json:"output,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (this *ConversionOutput) Reset()         { *this = ConversionOutput{} }
func (this *ConversionOutput) String() string { return proto.CompactTextString(this) }
func (*ConversionOutput) ProtoMessage()       {}

func (this *ConversionOutput) GetErrorCode() ConversionServiceError_ErrorCode {
	if this != nil && this.ErrorCode != nil {
		return *this.ErrorCode
	}
	return 0
}

func (this *ConversionOutput) GetOutput() *DocumentInfo {
	if this != nil {
		return this.Output
	}
	return nil
}

type ConversionRequest struct {
	Conversion       []*ConversionInput `protobuf:"bytes,1,rep,name=conversion" json:"conversion,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (this *ConversionRequest) Reset()         { *this = ConversionRequest{} }
func (this *ConversionRequest) String() string { return proto.CompactTextString(this) }
func (*ConversionRequest) ProtoMessage()       {}

type ConversionResponse struct {
	Result           []*ConversionOutput `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (this *ConversionResponse) Reset()         { *this = ConversionResponse{} }
func (this *ConversionResponse) String() string { return proto.CompactTextString(this) }
func (*ConversionResponse) ProtoMessage()       {}

func init() {
	proto.RegisterEnum("appengine.ConversionServiceError_ErrorCode", ConversionServiceError_ErrorCode_name, ConversionServiceError_ErrorCode_value)
}