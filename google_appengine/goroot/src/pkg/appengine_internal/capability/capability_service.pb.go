// Code generated by protoc-gen-go from "capability_service.proto"
// DO NOT EDIT!

package appengine

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CapabilityConfig_Status int32

const (
	CapabilityConfig_ENABLED   CapabilityConfig_Status = 1
	CapabilityConfig_SCHEDULED CapabilityConfig_Status = 2
	CapabilityConfig_DISABLED  CapabilityConfig_Status = 3
	CapabilityConfig_UNKNOWN   CapabilityConfig_Status = 4
)

var CapabilityConfig_Status_name = map[int32]string{
	1: "ENABLED",
	2: "SCHEDULED",
	3: "DISABLED",
	4: "UNKNOWN",
}
var CapabilityConfig_Status_value = map[string]int32{
	"ENABLED":   1,
	"SCHEDULED": 2,
	"DISABLED":  3,
	"UNKNOWN":   4,
}

func (x CapabilityConfig_Status) Enum() *CapabilityConfig_Status {
	p := new(CapabilityConfig_Status)
	*p = x
	return p
}
func (x CapabilityConfig_Status) String() string {
	return proto.EnumName(CapabilityConfig_Status_name, int32(x))
}
func (x CapabilityConfig_Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CapabilityConfig_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CapabilityConfig_Status_value, data, "CapabilityConfig_Status")
	if err != nil {
		return err
	}
	*x = CapabilityConfig_Status(value)
	return nil
}

type IsEnabledResponse_SummaryStatus int32

const (
	IsEnabledResponse_ENABLED          IsEnabledResponse_SummaryStatus = 1
	IsEnabledResponse_SCHEDULED_FUTURE IsEnabledResponse_SummaryStatus = 2
	IsEnabledResponse_SCHEDULED_NOW    IsEnabledResponse_SummaryStatus = 3
	IsEnabledResponse_DISABLED         IsEnabledResponse_SummaryStatus = 4
	IsEnabledResponse_UNKNOWN          IsEnabledResponse_SummaryStatus = 5
)

var IsEnabledResponse_SummaryStatus_name = map[int32]string{
	1: "ENABLED",
	2: "SCHEDULED_FUTURE",
	3: "SCHEDULED_NOW",
	4: "DISABLED",
	5: "UNKNOWN",
}
var IsEnabledResponse_SummaryStatus_value = map[string]int32{
	"ENABLED":          1,
	"SCHEDULED_FUTURE": 2,
	"SCHEDULED_NOW":    3,
	"DISABLED":         4,
	"UNKNOWN":          5,
}

func (x IsEnabledResponse_SummaryStatus) Enum() *IsEnabledResponse_SummaryStatus {
	p := new(IsEnabledResponse_SummaryStatus)
	*p = x
	return p
}
func (x IsEnabledResponse_SummaryStatus) String() string {
	return proto.EnumName(IsEnabledResponse_SummaryStatus_name, int32(x))
}
func (x IsEnabledResponse_SummaryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *IsEnabledResponse_SummaryStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IsEnabledResponse_SummaryStatus_value, data, "IsEnabledResponse_SummaryStatus")
	if err != nil {
		return err
	}
	*x = IsEnabledResponse_SummaryStatus(value)
	return nil
}

type CapabilityConfigList struct {
	Config           []*CapabilityConfig `protobuf:"bytes,1,rep,name=config" json:"config,omitempty"`
	DefaultConfig    *CapabilityConfig   `protobuf:"bytes,2,opt,name=default_config" json:"default_config,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (this *CapabilityConfigList) Reset()         { *this = CapabilityConfigList{} }
func (this *CapabilityConfigList) String() string { return proto.CompactTextString(this) }
func (*CapabilityConfigList) ProtoMessage()       {}

func (this *CapabilityConfigList) GetDefaultConfig() *CapabilityConfig {
	if this != nil {
		return this.DefaultConfig
	}
	return nil
}

type CapabilityConfig struct {
	Package          *string                  `protobuf:"bytes,1,req,name=package" json:"package,omitempty"`
	Capability       *string                  `protobuf:"bytes,2,req,name=capability" json:"capability,omitempty"`
	Status           *CapabilityConfig_Status `protobuf:"varint,3,opt,name=status,enum=appengine.CapabilityConfig_Status,def=4" json:"status,omitempty"`
	ScheduledTime    *string                  `protobuf:"bytes,7,opt,name=scheduled_time" json:"scheduled_time,omitempty"`
	InternalMessage  *string                  `protobuf:"bytes,4,opt,name=internal_message" json:"internal_message,omitempty"`
	AdminMessage     *string                  `protobuf:"bytes,5,opt,name=admin_message" json:"admin_message,omitempty"`
	ErrorMessage     *string                  `protobuf:"bytes,6,opt,name=error_message" json:"error_message,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (this *CapabilityConfig) Reset()         { *this = CapabilityConfig{} }
func (this *CapabilityConfig) String() string { return proto.CompactTextString(this) }
func (*CapabilityConfig) ProtoMessage()       {}

const Default_CapabilityConfig_Status CapabilityConfig_Status = CapabilityConfig_UNKNOWN

func (this *CapabilityConfig) GetPackage() string {
	if this != nil && this.Package != nil {
		return *this.Package
	}
	return ""
}

func (this *CapabilityConfig) GetCapability() string {
	if this != nil && this.Capability != nil {
		return *this.Capability
	}
	return ""
}

func (this *CapabilityConfig) GetStatus() CapabilityConfig_Status {
	if this != nil && this.Status != nil {
		return *this.Status
	}
	return Default_CapabilityConfig_Status
}

func (this *CapabilityConfig) GetScheduledTime() string {
	if this != nil && this.ScheduledTime != nil {
		return *this.ScheduledTime
	}
	return ""
}

func (this *CapabilityConfig) GetInternalMessage() string {
	if this != nil && this.InternalMessage != nil {
		return *this.InternalMessage
	}
	return ""
}

func (this *CapabilityConfig) GetAdminMessage() string {
	if this != nil && this.AdminMessage != nil {
		return *this.AdminMessage
	}
	return ""
}

func (this *CapabilityConfig) GetErrorMessage() string {
	if this != nil && this.ErrorMessage != nil {
		return *this.ErrorMessage
	}
	return ""
}

type IsEnabledRequest struct {
	Package          *string  `protobuf:"bytes,1,req,name=package" json:"package,omitempty"`
	Capability       []string `protobuf:"bytes,2,rep,name=capability" json:"capability,omitempty"`
	Call             []string `protobuf:"bytes,3,rep,name=call" json:"call,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *IsEnabledRequest) Reset()         { *this = IsEnabledRequest{} }
func (this *IsEnabledRequest) String() string { return proto.CompactTextString(this) }
func (*IsEnabledRequest) ProtoMessage()       {}

func (this *IsEnabledRequest) GetPackage() string {
	if this != nil && this.Package != nil {
		return *this.Package
	}
	return ""
}

type IsEnabledResponse struct {
	SummaryStatus      *IsEnabledResponse_SummaryStatus `protobuf:"varint,1,req,name=summary_status,enum=appengine.IsEnabledResponse_SummaryStatus" json:"summary_status,omitempty"`
	TimeUntilScheduled *int64                           `protobuf:"varint,2,opt,name=time_until_scheduled" json:"time_until_scheduled,omitempty"`
	Config             []*CapabilityConfig              `protobuf:"bytes,3,rep,name=config" json:"config,omitempty"`
	XXX_unrecognized   []byte                           `json:"-"`
}

func (this *IsEnabledResponse) Reset()         { *this = IsEnabledResponse{} }
func (this *IsEnabledResponse) String() string { return proto.CompactTextString(this) }
func (*IsEnabledResponse) ProtoMessage()       {}

func (this *IsEnabledResponse) GetSummaryStatus() IsEnabledResponse_SummaryStatus {
	if this != nil && this.SummaryStatus != nil {
		return *this.SummaryStatus
	}
	return 0
}

func (this *IsEnabledResponse) GetTimeUntilScheduled() int64 {
	if this != nil && this.TimeUntilScheduled != nil {
		return *this.TimeUntilScheduled
	}
	return 0
}

func init() {
	proto.RegisterEnum("appengine.CapabilityConfig_Status", CapabilityConfig_Status_name, CapabilityConfig_Status_value)
	proto.RegisterEnum("appengine.IsEnabledResponse_SummaryStatus", IsEnabledResponse_SummaryStatus_name, IsEnabledResponse_SummaryStatus_value)
}
