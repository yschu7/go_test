// Code generated by protoc-gen-go from "log_service.proto"
// DO NOT EDIT!

package appengine

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type LogServiceError_ErrorCode int32

const (
	LogServiceError_OK              LogServiceError_ErrorCode = 0
	LogServiceError_INVALID_REQUEST LogServiceError_ErrorCode = 1
	LogServiceError_STORAGE_ERROR   LogServiceError_ErrorCode = 2
)

var LogServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INVALID_REQUEST",
	2: "STORAGE_ERROR",
}
var LogServiceError_ErrorCode_value = map[string]int32{
	"OK":              0,
	"INVALID_REQUEST": 1,
	"STORAGE_ERROR":   2,
}

func (x LogServiceError_ErrorCode) Enum() *LogServiceError_ErrorCode {
	p := new(LogServiceError_ErrorCode)
	*p = x
	return p
}
func (x LogServiceError_ErrorCode) String() string {
	return proto.EnumName(LogServiceError_ErrorCode_name, int32(x))
}
func (x LogServiceError_ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LogServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LogServiceError_ErrorCode_value, data, "LogServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = LogServiceError_ErrorCode(value)
	return nil
}

type LogServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *LogServiceError) Reset()         { *this = LogServiceError{} }
func (this *LogServiceError) String() string { return proto.CompactTextString(this) }
func (*LogServiceError) ProtoMessage()       {}

type UserAppLogLine struct {
	TimestampUsec    *int64  `protobuf:"varint,1,req,name=timestamp_usec" json:"timestamp_usec,omitempty"`
	Level            *int64  `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Message          *string `protobuf:"bytes,3,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *UserAppLogLine) Reset()         { *this = UserAppLogLine{} }
func (this *UserAppLogLine) String() string { return proto.CompactTextString(this) }
func (*UserAppLogLine) ProtoMessage()       {}

func (this *UserAppLogLine) GetTimestampUsec() int64 {
	if this != nil && this.TimestampUsec != nil {
		return *this.TimestampUsec
	}
	return 0
}

func (this *UserAppLogLine) GetLevel() int64 {
	if this != nil && this.Level != nil {
		return *this.Level
	}
	return 0
}

func (this *UserAppLogLine) GetMessage() string {
	if this != nil && this.Message != nil {
		return *this.Message
	}
	return ""
}

type UserAppLogGroup struct {
	LogLine          []*UserAppLogLine `protobuf:"bytes,2,rep,name=log_line" json:"log_line,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (this *UserAppLogGroup) Reset()         { *this = UserAppLogGroup{} }
func (this *UserAppLogGroup) String() string { return proto.CompactTextString(this) }
func (*UserAppLogGroup) ProtoMessage()       {}

type FlushRequest struct {
	Logs             []byte `protobuf:"bytes,1,opt,name=logs" json:"logs,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *FlushRequest) Reset()         { *this = FlushRequest{} }
func (this *FlushRequest) String() string { return proto.CompactTextString(this) }
func (*FlushRequest) ProtoMessage()       {}

func (this *FlushRequest) GetLogs() []byte {
	if this != nil {
		return this.Logs
	}
	return nil
}

type SetStatusRequest struct {
	Status           *string `protobuf:"bytes,1,req,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *SetStatusRequest) Reset()         { *this = SetStatusRequest{} }
func (this *SetStatusRequest) String() string { return proto.CompactTextString(this) }
func (*SetStatusRequest) ProtoMessage()       {}

func (this *SetStatusRequest) GetStatus() string {
	if this != nil && this.Status != nil {
		return *this.Status
	}
	return ""
}

type LogOffset struct {
	RequestId        []byte `protobuf:"bytes,1,opt,name=request_id" json:"request_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *LogOffset) Reset()         { *this = LogOffset{} }
func (this *LogOffset) String() string { return proto.CompactTextString(this) }
func (*LogOffset) ProtoMessage()       {}

func (this *LogOffset) GetRequestId() []byte {
	if this != nil {
		return this.RequestId
	}
	return nil
}

type LogLine struct {
	Time             *int64  `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	Level            *int32  `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	LogMessage       *string `protobuf:"bytes,3,req,name=log_message" json:"log_message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *LogLine) Reset()         { *this = LogLine{} }
func (this *LogLine) String() string { return proto.CompactTextString(this) }
func (*LogLine) ProtoMessage()       {}

func (this *LogLine) GetTime() int64 {
	if this != nil && this.Time != nil {
		return *this.Time
	}
	return 0
}

func (this *LogLine) GetLevel() int32 {
	if this != nil && this.Level != nil {
		return *this.Level
	}
	return 0
}

func (this *LogLine) GetLogMessage() string {
	if this != nil && this.LogMessage != nil {
		return *this.LogMessage
	}
	return ""
}

type RequestLog struct {
	AppId                   *string    `protobuf:"bytes,1,req,name=app_id" json:"app_id,omitempty"`
	VersionId               *string    `protobuf:"bytes,2,req,name=version_id" json:"version_id,omitempty"`
	RequestId               []byte     `protobuf:"bytes,3,req,name=request_id" json:"request_id,omitempty"`
	Offset                  *LogOffset `protobuf:"bytes,35,opt,name=offset" json:"offset,omitempty"`
	Ip                      *string    `protobuf:"bytes,4,req,name=ip" json:"ip,omitempty"`
	Nickname                *string    `protobuf:"bytes,5,opt,name=nickname" json:"nickname,omitempty"`
	StartTime               *int64     `protobuf:"varint,6,req,name=start_time" json:"start_time,omitempty"`
	EndTime                 *int64     `protobuf:"varint,7,req,name=end_time" json:"end_time,omitempty"`
	Latency                 *int64     `protobuf:"varint,8,req,name=latency" json:"latency,omitempty"`
	Mcycles                 *int64     `protobuf:"varint,9,req,name=mcycles" json:"mcycles,omitempty"`
	Method                  *string    `protobuf:"bytes,10,req,name=method" json:"method,omitempty"`
	Resource                *string    `protobuf:"bytes,11,req,name=resource" json:"resource,omitempty"`
	HttpVersion             *string    `protobuf:"bytes,12,req,name=http_version" json:"http_version,omitempty"`
	Status                  *int32     `protobuf:"varint,13,req,name=status" json:"status,omitempty"`
	ResponseSize            *int64     `protobuf:"varint,14,req,name=response_size" json:"response_size,omitempty"`
	Referrer                *string    `protobuf:"bytes,15,opt,name=referrer" json:"referrer,omitempty"`
	UserAgent               *string    `protobuf:"bytes,16,opt,name=user_agent" json:"user_agent,omitempty"`
	UrlMapEntry             *string    `protobuf:"bytes,17,req,name=url_map_entry" json:"url_map_entry,omitempty"`
	Combined                *string    `protobuf:"bytes,18,req,name=combined" json:"combined,omitempty"`
	ApiMcycles              *int64     `protobuf:"varint,19,opt,name=api_mcycles" json:"api_mcycles,omitempty"`
	Host                    *string    `protobuf:"bytes,20,opt,name=host" json:"host,omitempty"`
	Cost                    *float64   `protobuf:"fixed64,21,opt,name=cost" json:"cost,omitempty"`
	TaskQueueName           *string    `protobuf:"bytes,22,opt,name=task_queue_name" json:"task_queue_name,omitempty"`
	TaskName                *string    `protobuf:"bytes,23,opt,name=task_name" json:"task_name,omitempty"`
	WasLoadingRequest       *bool      `protobuf:"varint,24,opt,name=was_loading_request" json:"was_loading_request,omitempty"`
	PendingTime             *int64     `protobuf:"varint,25,opt,name=pending_time" json:"pending_time,omitempty"`
	ReplicaIndex            *int32     `protobuf:"varint,26,opt,name=replica_index,def=-1" json:"replica_index,omitempty"`
	Finished                *bool      `protobuf:"varint,27,opt,name=finished,def=1" json:"finished,omitempty"`
	CloneKey                []byte     `protobuf:"bytes,28,opt,name=clone_key" json:"clone_key,omitempty"`
	Line                    []*LogLine `protobuf:"bytes,29,rep,name=line" json:"line,omitempty"`
	LinesIncomplete         *bool      `protobuf:"varint,36,opt,name=lines_incomplete" json:"lines_incomplete,omitempty"`
	ExitReason              *int32     `protobuf:"varint,30,opt,name=exit_reason" json:"exit_reason,omitempty"`
	WasThrottledForTime     *bool      `protobuf:"varint,31,opt,name=was_throttled_for_time" json:"was_throttled_for_time,omitempty"`
	WasThrottledForRequests *bool      `protobuf:"varint,32,opt,name=was_throttled_for_requests" json:"was_throttled_for_requests,omitempty"`
	ThrottledTime           *int64     `protobuf:"varint,33,opt,name=throttled_time" json:"throttled_time,omitempty"`
	ServerName              []byte     `protobuf:"bytes,34,opt,name=server_name" json:"server_name,omitempty"`
	XXX_unrecognized        []byte     `json:"-"`
}

func (this *RequestLog) Reset()         { *this = RequestLog{} }
func (this *RequestLog) String() string { return proto.CompactTextString(this) }
func (*RequestLog) ProtoMessage()       {}

const Default_RequestLog_ReplicaIndex int32 = -1
const Default_RequestLog_Finished bool = true

func (this *RequestLog) GetAppId() string {
	if this != nil && this.AppId != nil {
		return *this.AppId
	}
	return ""
}

func (this *RequestLog) GetVersionId() string {
	if this != nil && this.VersionId != nil {
		return *this.VersionId
	}
	return ""
}

func (this *RequestLog) GetRequestId() []byte {
	if this != nil {
		return this.RequestId
	}
	return nil
}

func (this *RequestLog) GetOffset() *LogOffset {
	if this != nil {
		return this.Offset
	}
	return nil
}

func (this *RequestLog) GetIp() string {
	if this != nil && this.Ip != nil {
		return *this.Ip
	}
	return ""
}

func (this *RequestLog) GetNickname() string {
	if this != nil && this.Nickname != nil {
		return *this.Nickname
	}
	return ""
}

func (this *RequestLog) GetStartTime() int64 {
	if this != nil && this.StartTime != nil {
		return *this.StartTime
	}
	return 0
}

func (this *RequestLog) GetEndTime() int64 {
	if this != nil && this.EndTime != nil {
		return *this.EndTime
	}
	return 0
}

func (this *RequestLog) GetLatency() int64 {
	if this != nil && this.Latency != nil {
		return *this.Latency
	}
	return 0
}

func (this *RequestLog) GetMcycles() int64 {
	if this != nil && this.Mcycles != nil {
		return *this.Mcycles
	}
	return 0
}

func (this *RequestLog) GetMethod() string {
	if this != nil && this.Method != nil {
		return *this.Method
	}
	return ""
}

func (this *RequestLog) GetResource() string {
	if this != nil && this.Resource != nil {
		return *this.Resource
	}
	return ""
}

func (this *RequestLog) GetHttpVersion() string {
	if this != nil && this.HttpVersion != nil {
		return *this.HttpVersion
	}
	return ""
}

func (this *RequestLog) GetStatus() int32 {
	if this != nil && this.Status != nil {
		return *this.Status
	}
	return 0
}

func (this *RequestLog) GetResponseSize() int64 {
	if this != nil && this.ResponseSize != nil {
		return *this.ResponseSize
	}
	return 0
}

func (this *RequestLog) GetReferrer() string {
	if this != nil && this.Referrer != nil {
		return *this.Referrer
	}
	return ""
}

func (this *RequestLog) GetUserAgent() string {
	if this != nil && this.UserAgent != nil {
		return *this.UserAgent
	}
	return ""
}

func (this *RequestLog) GetUrlMapEntry() string {
	if this != nil && this.UrlMapEntry != nil {
		return *this.UrlMapEntry
	}
	return ""
}

func (this *RequestLog) GetCombined() string {
	if this != nil && this.Combined != nil {
		return *this.Combined
	}
	return ""
}

func (this *RequestLog) GetApiMcycles() int64 {
	if this != nil && this.ApiMcycles != nil {
		return *this.ApiMcycles
	}
	return 0
}

func (this *RequestLog) GetHost() string {
	if this != nil && this.Host != nil {
		return *this.Host
	}
	return ""
}

func (this *RequestLog) GetCost() float64 {
	if this != nil && this.Cost != nil {
		return *this.Cost
	}
	return 0
}

func (this *RequestLog) GetTaskQueueName() string {
	if this != nil && this.TaskQueueName != nil {
		return *this.TaskQueueName
	}
	return ""
}

func (this *RequestLog) GetTaskName() string {
	if this != nil && this.TaskName != nil {
		return *this.TaskName
	}
	return ""
}

func (this *RequestLog) GetWasLoadingRequest() bool {
	if this != nil && this.WasLoadingRequest != nil {
		return *this.WasLoadingRequest
	}
	return false
}

func (this *RequestLog) GetPendingTime() int64 {
	if this != nil && this.PendingTime != nil {
		return *this.PendingTime
	}
	return 0
}

func (this *RequestLog) GetReplicaIndex() int32 {
	if this != nil && this.ReplicaIndex != nil {
		return *this.ReplicaIndex
	}
	return Default_RequestLog_ReplicaIndex
}

func (this *RequestLog) GetFinished() bool {
	if this != nil && this.Finished != nil {
		return *this.Finished
	}
	return Default_RequestLog_Finished
}

func (this *RequestLog) GetCloneKey() []byte {
	if this != nil {
		return this.CloneKey
	}
	return nil
}

func (this *RequestLog) GetLinesIncomplete() bool {
	if this != nil && this.LinesIncomplete != nil {
		return *this.LinesIncomplete
	}
	return false
}

func (this *RequestLog) GetExitReason() int32 {
	if this != nil && this.ExitReason != nil {
		return *this.ExitReason
	}
	return 0
}

func (this *RequestLog) GetWasThrottledForTime() bool {
	if this != nil && this.WasThrottledForTime != nil {
		return *this.WasThrottledForTime
	}
	return false
}

func (this *RequestLog) GetWasThrottledForRequests() bool {
	if this != nil && this.WasThrottledForRequests != nil {
		return *this.WasThrottledForRequests
	}
	return false
}

func (this *RequestLog) GetThrottledTime() int64 {
	if this != nil && this.ThrottledTime != nil {
		return *this.ThrottledTime
	}
	return 0
}

func (this *RequestLog) GetServerName() []byte {
	if this != nil {
		return this.ServerName
	}
	return nil
}

type LogReadRequest struct {
	AppId             *string    `protobuf:"bytes,1,req,name=app_id" json:"app_id,omitempty"`
	VersionId         []string   `protobuf:"bytes,2,rep,name=version_id" json:"version_id,omitempty"`
	StartTime         *int64     `protobuf:"varint,3,opt,name=start_time" json:"start_time,omitempty"`
	EndTime           *int64     `protobuf:"varint,4,opt,name=end_time" json:"end_time,omitempty"`
	Offset            *LogOffset `protobuf:"bytes,5,opt,name=offset" json:"offset,omitempty"`
	RequestId         [][]byte   `protobuf:"bytes,6,rep,name=request_id" json:"request_id,omitempty"`
	MinimumLogLevel   *int32     `protobuf:"varint,7,opt,name=minimum_log_level" json:"minimum_log_level,omitempty"`
	IncludeIncomplete *bool      `protobuf:"varint,8,opt,name=include_incomplete" json:"include_incomplete,omitempty"`
	Count             *int64     `protobuf:"varint,9,opt,name=count" json:"count,omitempty"`
	CombinedLogRegex  *string    `protobuf:"bytes,14,opt,name=combined_log_regex" json:"combined_log_regex,omitempty"`
	HostRegex         *string    `protobuf:"bytes,15,opt,name=host_regex" json:"host_regex,omitempty"`
	ReplicaIndex      *int32     `protobuf:"varint,16,opt,name=replica_index" json:"replica_index,omitempty"`
	IncludeAppLogs    *bool      `protobuf:"varint,10,opt,name=include_app_logs" json:"include_app_logs,omitempty"`
	AppLogsPerRequest *int32     `protobuf:"varint,17,opt,name=app_logs_per_request" json:"app_logs_per_request,omitempty"`
	IncludeHost       *bool      `protobuf:"varint,11,opt,name=include_host" json:"include_host,omitempty"`
	IncludeAll        *bool      `protobuf:"varint,12,opt,name=include_all" json:"include_all,omitempty"`
	CacheIterator     *bool      `protobuf:"varint,13,opt,name=cache_iterator" json:"cache_iterator,omitempty"`
	NumShards         *int32     `protobuf:"varint,18,opt,name=num_shards" json:"num_shards,omitempty"`
	XXX_unrecognized  []byte     `json:"-"`
}

func (this *LogReadRequest) Reset()         { *this = LogReadRequest{} }
func (this *LogReadRequest) String() string { return proto.CompactTextString(this) }
func (*LogReadRequest) ProtoMessage()       {}

func (this *LogReadRequest) GetAppId() string {
	if this != nil && this.AppId != nil {
		return *this.AppId
	}
	return ""
}

func (this *LogReadRequest) GetStartTime() int64 {
	if this != nil && this.StartTime != nil {
		return *this.StartTime
	}
	return 0
}

func (this *LogReadRequest) GetEndTime() int64 {
	if this != nil && this.EndTime != nil {
		return *this.EndTime
	}
	return 0
}

func (this *LogReadRequest) GetOffset() *LogOffset {
	if this != nil {
		return this.Offset
	}
	return nil
}

func (this *LogReadRequest) GetMinimumLogLevel() int32 {
	if this != nil && this.MinimumLogLevel != nil {
		return *this.MinimumLogLevel
	}
	return 0
}

func (this *LogReadRequest) GetIncludeIncomplete() bool {
	if this != nil && this.IncludeIncomplete != nil {
		return *this.IncludeIncomplete
	}
	return false
}

func (this *LogReadRequest) GetCount() int64 {
	if this != nil && this.Count != nil {
		return *this.Count
	}
	return 0
}

func (this *LogReadRequest) GetCombinedLogRegex() string {
	if this != nil && this.CombinedLogRegex != nil {
		return *this.CombinedLogRegex
	}
	return ""
}

func (this *LogReadRequest) GetHostRegex() string {
	if this != nil && this.HostRegex != nil {
		return *this.HostRegex
	}
	return ""
}

func (this *LogReadRequest) GetReplicaIndex() int32 {
	if this != nil && this.ReplicaIndex != nil {
		return *this.ReplicaIndex
	}
	return 0
}

func (this *LogReadRequest) GetIncludeAppLogs() bool {
	if this != nil && this.IncludeAppLogs != nil {
		return *this.IncludeAppLogs
	}
	return false
}

func (this *LogReadRequest) GetAppLogsPerRequest() int32 {
	if this != nil && this.AppLogsPerRequest != nil {
		return *this.AppLogsPerRequest
	}
	return 0
}

func (this *LogReadRequest) GetIncludeHost() bool {
	if this != nil && this.IncludeHost != nil {
		return *this.IncludeHost
	}
	return false
}

func (this *LogReadRequest) GetIncludeAll() bool {
	if this != nil && this.IncludeAll != nil {
		return *this.IncludeAll
	}
	return false
}

func (this *LogReadRequest) GetCacheIterator() bool {
	if this != nil && this.CacheIterator != nil {
		return *this.CacheIterator
	}
	return false
}

func (this *LogReadRequest) GetNumShards() int32 {
	if this != nil && this.NumShards != nil {
		return *this.NumShards
	}
	return 0
}

type LogReadResponse struct {
	Log              []*RequestLog `protobuf:"bytes,1,rep,name=log" json:"log,omitempty"`
	Offset           *LogOffset    `protobuf:"bytes,2,opt,name=offset" json:"offset,omitempty"`
	LastEndTime      *int64        `protobuf:"varint,3,opt,name=last_end_time" json:"last_end_time,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (this *LogReadResponse) Reset()         { *this = LogReadResponse{} }
func (this *LogReadResponse) String() string { return proto.CompactTextString(this) }
func (*LogReadResponse) ProtoMessage()       {}

func (this *LogReadResponse) GetOffset() *LogOffset {
	if this != nil {
		return this.Offset
	}
	return nil
}

func (this *LogReadResponse) GetLastEndTime() int64 {
	if this != nil && this.LastEndTime != nil {
		return *this.LastEndTime
	}
	return 0
}

type LogUsageRecord struct {
	VersionId        *string `protobuf:"bytes,1,opt,name=version_id" json:"version_id,omitempty"`
	StartTime        *int32  `protobuf:"varint,2,opt,name=start_time" json:"start_time,omitempty"`
	EndTime          *int32  `protobuf:"varint,3,opt,name=end_time" json:"end_time,omitempty"`
	Count            *int64  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	TotalSize        *int64  `protobuf:"varint,5,opt,name=total_size" json:"total_size,omitempty"`
	Records          *int32  `protobuf:"varint,6,opt,name=records" json:"records,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *LogUsageRecord) Reset()         { *this = LogUsageRecord{} }
func (this *LogUsageRecord) String() string { return proto.CompactTextString(this) }
func (*LogUsageRecord) ProtoMessage()       {}

func (this *LogUsageRecord) GetVersionId() string {
	if this != nil && this.VersionId != nil {
		return *this.VersionId
	}
	return ""
}

func (this *LogUsageRecord) GetStartTime() int32 {
	if this != nil && this.StartTime != nil {
		return *this.StartTime
	}
	return 0
}

func (this *LogUsageRecord) GetEndTime() int32 {
	if this != nil && this.EndTime != nil {
		return *this.EndTime
	}
	return 0
}

func (this *LogUsageRecord) GetCount() int64 {
	if this != nil && this.Count != nil {
		return *this.Count
	}
	return 0
}

func (this *LogUsageRecord) GetTotalSize() int64 {
	if this != nil && this.TotalSize != nil {
		return *this.TotalSize
	}
	return 0
}

func (this *LogUsageRecord) GetRecords() int32 {
	if this != nil && this.Records != nil {
		return *this.Records
	}
	return 0
}

type LogUsageRequest struct {
	AppId            *string  `protobuf:"bytes,1,req,name=app_id" json:"app_id,omitempty"`
	VersionId        []string `protobuf:"bytes,2,rep,name=version_id" json:"version_id,omitempty"`
	StartTime        *int32   `protobuf:"varint,3,opt,name=start_time" json:"start_time,omitempty"`
	EndTime          *int32   `protobuf:"varint,4,opt,name=end_time" json:"end_time,omitempty"`
	ResolutionHours  *uint32  `protobuf:"varint,5,opt,name=resolution_hours,def=1" json:"resolution_hours,omitempty"`
	CombineVersions  *bool    `protobuf:"varint,6,opt,name=combine_versions" json:"combine_versions,omitempty"`
	UsageVersion     *int32   `protobuf:"varint,7,opt,name=usage_version" json:"usage_version,omitempty"`
	VersionsOnly     *bool    `protobuf:"varint,8,opt,name=versions_only" json:"versions_only,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *LogUsageRequest) Reset()         { *this = LogUsageRequest{} }
func (this *LogUsageRequest) String() string { return proto.CompactTextString(this) }
func (*LogUsageRequest) ProtoMessage()       {}

const Default_LogUsageRequest_ResolutionHours uint32 = 1

func (this *LogUsageRequest) GetAppId() string {
	if this != nil && this.AppId != nil {
		return *this.AppId
	}
	return ""
}

func (this *LogUsageRequest) GetStartTime() int32 {
	if this != nil && this.StartTime != nil {
		return *this.StartTime
	}
	return 0
}

func (this *LogUsageRequest) GetEndTime() int32 {
	if this != nil && this.EndTime != nil {
		return *this.EndTime
	}
	return 0
}

func (this *LogUsageRequest) GetResolutionHours() uint32 {
	if this != nil && this.ResolutionHours != nil {
		return *this.ResolutionHours
	}
	return Default_LogUsageRequest_ResolutionHours
}

func (this *LogUsageRequest) GetCombineVersions() bool {
	if this != nil && this.CombineVersions != nil {
		return *this.CombineVersions
	}
	return false
}

func (this *LogUsageRequest) GetUsageVersion() int32 {
	if this != nil && this.UsageVersion != nil {
		return *this.UsageVersion
	}
	return 0
}

func (this *LogUsageRequest) GetVersionsOnly() bool {
	if this != nil && this.VersionsOnly != nil {
		return *this.VersionsOnly
	}
	return false
}

type LogUsageResponse struct {
	Usage            []*LogUsageRecord `protobuf:"bytes,1,rep,name=usage" json:"usage,omitempty"`
	Summary          *LogUsageRecord   `protobuf:"bytes,2,opt,name=summary" json:"summary,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (this *LogUsageResponse) Reset()         { *this = LogUsageResponse{} }
func (this *LogUsageResponse) String() string { return proto.CompactTextString(this) }
func (*LogUsageResponse) ProtoMessage()       {}

func (this *LogUsageResponse) GetSummary() *LogUsageRecord {
	if this != nil {
		return this.Summary
	}
	return nil
}

func init() {
	proto.RegisterEnum("appengine.LogServiceError_ErrorCode", LogServiceError_ErrorCode_name, LogServiceError_ErrorCode_value)
}
