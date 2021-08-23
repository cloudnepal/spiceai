// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/aiengine/v1/aiengine.proto

package aiengine_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataConnector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Params map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DataConnector) Reset() {
	*x = DataConnector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataConnector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataConnector) ProtoMessage() {}

func (x *DataConnector) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataConnector.ProtoReflect.Descriptor instead.
func (*DataConnector) Descriptor() ([]byte, []int) {
	return file_proto_aiengine_v1_aiengine_proto_rawDescGZIP(), []int{0}
}

func (x *DataConnector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataConnector) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type DataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connector *DataConnector    `protobuf:"bytes,1,opt,name=connector,proto3" json:"connector,omitempty"`
	Actions   map[string]string `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_proto_aiengine_v1_aiengine_proto_rawDescGZIP(), []int{1}
}

func (x *DataSource) GetConnector() *DataConnector {
	if x != nil {
		return x.Connector
	}
	return nil
}

func (x *DataSource) GetActions() map[string]string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod         string             `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	Period      int64              `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	Interval    int64              `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
	Granularity int64              `protobuf:"varint,4,opt,name=granularity,proto3" json:"granularity,omitempty"`
	EpochTime   int64              `protobuf:"varint,5,opt,name=epoch_time,json=epochTime,proto3" json:"epoch_time,omitempty"`
	Actions     map[string]string  `protobuf:"bytes,6,rep,name=actions,proto3" json:"actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Fields      map[string]float64 `protobuf:"bytes,7,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Laws        []string           `protobuf:"bytes,8,rep,name=laws,proto3" json:"laws,omitempty"`
	Datasources []*DataSource      `protobuf:"bytes,9,rep,name=datasources,proto3" json:"datasources,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_proto_aiengine_v1_aiengine_proto_rawDescGZIP(), []int{2}
}

func (x *InitRequest) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *InitRequest) GetPeriod() int64 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *InitRequest) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *InitRequest) GetGranularity() int64 {
	if x != nil {
		return x.Granularity
	}
	return 0
}

func (x *InitRequest) GetEpochTime() int64 {
	if x != nil {
		return x.EpochTime
	}
	return 0
}

func (x *InitRequest) GetActions() map[string]string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *InitRequest) GetFields() map[string]float64 {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *InitRequest) GetLaws() []string {
	if x != nil {
		return x.Laws
	}
	return nil
}

func (x *InitRequest) GetDatasources() []*DataSource {
	if x != nil {
		return x.Datasources
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error   bool   `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_aiengine_v1_aiengine_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type StartTrainingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod            string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	NumberEpisodes int64  `protobuf:"varint,2,opt,name=number_episodes,json=numberEpisodes,proto3" json:"number_episodes,omitempty"`
	Flight         string `protobuf:"bytes,3,opt,name=flight,proto3" json:"flight,omitempty"`
	TrainingGoal   string `protobuf:"bytes,4,opt,name=training_goal,json=trainingGoal,proto3" json:"training_goal,omitempty"`
	EpochTime      int64  `protobuf:"varint,5,opt,name=epoch_time,json=epochTime,proto3" json:"epoch_time,omitempty"`
}

func (x *StartTrainingRequest) Reset() {
	*x = StartTrainingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTrainingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTrainingRequest) ProtoMessage() {}

func (x *StartTrainingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTrainingRequest.ProtoReflect.Descriptor instead.
func (*StartTrainingRequest) Descriptor() ([]byte, []int) {
	return file_proto_aiengine_v1_aiengine_proto_rawDescGZIP(), []int{4}
}

func (x *StartTrainingRequest) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *StartTrainingRequest) GetNumberEpisodes() int64 {
	if x != nil {
		return x.NumberEpisodes
	}
	return 0
}

func (x *StartTrainingRequest) GetFlight() string {
	if x != nil {
		return x.Flight
	}
	return ""
}

func (x *StartTrainingRequest) GetTrainingGoal() string {
	if x != nil {
		return x.TrainingGoal
	}
	return ""
}

func (x *StartTrainingRequest) GetEpochTime() int64 {
	if x != nil {
		return x.EpochTime
	}
	return 0
}

type InferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *InferenceRequest) Reset() {
	*x = InferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceRequest) ProtoMessage() {}

func (x *InferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceRequest.ProtoReflect.Descriptor instead.
func (*InferenceRequest) Descriptor() ([]byte, []int) {
	return file_proto_aiengine_v1_aiengine_proto_rawDescGZIP(), []int{5}
}

func (x *InferenceRequest) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *InferenceRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type InferenceResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response   *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Start      int64     `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End        int64     `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	Action     string    `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	Confidence float32   `protobuf:"fixed32,5,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Tag        string    `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *InferenceResult) Reset() {
	*x = InferenceResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceResult) ProtoMessage() {}

func (x *InferenceResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceResult.ProtoReflect.Descriptor instead.
func (*InferenceResult) Descriptor() ([]byte, []int) {
	return file_proto_aiengine_v1_aiengine_proto_rawDescGZIP(), []int{6}
}

func (x *InferenceResult) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *InferenceResult) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *InferenceResult) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *InferenceResult) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *InferenceResult) GetConfidence() float32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *InferenceResult) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type AddDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod     string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	CsvData string `protobuf:"bytes,2,opt,name=csv_data,json=csvData,proto3" json:"csv_data,omitempty"`
}

func (x *AddDataRequest) Reset() {
	*x = AddDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDataRequest) ProtoMessage() {}

func (x *AddDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDataRequest.ProtoReflect.Descriptor instead.
func (*AddDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_aiengine_v1_aiengine_proto_rawDescGZIP(), []int{7}
}

func (x *AddDataRequest) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *AddDataRequest) GetCsvData() string {
	if x != nil {
		return x.CsvData
	}
	return ""
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aiengine_v1_aiengine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_proto_aiengine_v1_aiengine_proto_rawDescGZIP(), []int{8}
}

var File_proto_aiengine_v1_aiengine_proto protoreflect.FileDescriptor

var file_proto_aiengine_v1_aiengine_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x22, 0x9b, 0x01, 0x0a,
	0x0d, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbc, 0x01, 0x0a, 0x0a, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61,
	0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x3b, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3a, 0x0a,
	0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd0, 0x03, 0x0a, 0x0b, 0x49, 0x6e,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3c, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x77,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x77, 0x73, 0x12, 0x36, 0x0a,
	0x0b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0xad, 0x01, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x45, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x61,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x36, 0x0a, 0x10, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x3d,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x6f, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x73, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x73, 0x76, 0x44, 0x61, 0x74, 0x61, 0x22, 0x0f, 0x0a,
	0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xbc,
	0x02, 0x0a, 0x08, 0x41, 0x49, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x49,
	0x6e, 0x69, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x69, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x61, 0x69, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x61,
	0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x12, 0x17, 0x2e, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x69, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x63,
	0x65, 0x61, 0x69, 0x2f, 0x73, 0x70, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_aiengine_v1_aiengine_proto_rawDescOnce sync.Once
	file_proto_aiengine_v1_aiengine_proto_rawDescData = file_proto_aiengine_v1_aiengine_proto_rawDesc
)

func file_proto_aiengine_v1_aiengine_proto_rawDescGZIP() []byte {
	file_proto_aiengine_v1_aiengine_proto_rawDescOnce.Do(func() {
		file_proto_aiengine_v1_aiengine_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_aiengine_v1_aiengine_proto_rawDescData)
	})
	return file_proto_aiengine_v1_aiengine_proto_rawDescData
}

var file_proto_aiengine_v1_aiengine_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_aiengine_v1_aiengine_proto_goTypes = []interface{}{
	(*DataConnector)(nil),        // 0: aiengine.DataConnector
	(*DataSource)(nil),           // 1: aiengine.DataSource
	(*InitRequest)(nil),          // 2: aiengine.InitRequest
	(*Response)(nil),             // 3: aiengine.Response
	(*StartTrainingRequest)(nil), // 4: aiengine.StartTrainingRequest
	(*InferenceRequest)(nil),     // 5: aiengine.InferenceRequest
	(*InferenceResult)(nil),      // 6: aiengine.InferenceResult
	(*AddDataRequest)(nil),       // 7: aiengine.AddDataRequest
	(*HealthRequest)(nil),        // 8: aiengine.HealthRequest
	nil,                          // 9: aiengine.DataConnector.ParamsEntry
	nil,                          // 10: aiengine.DataSource.ActionsEntry
	nil,                          // 11: aiengine.InitRequest.ActionsEntry
	nil,                          // 12: aiengine.InitRequest.FieldsEntry
}
var file_proto_aiengine_v1_aiengine_proto_depIdxs = []int32{
	9,  // 0: aiengine.DataConnector.params:type_name -> aiengine.DataConnector.ParamsEntry
	0,  // 1: aiengine.DataSource.connector:type_name -> aiengine.DataConnector
	10, // 2: aiengine.DataSource.actions:type_name -> aiengine.DataSource.ActionsEntry
	11, // 3: aiengine.InitRequest.actions:type_name -> aiengine.InitRequest.ActionsEntry
	12, // 4: aiengine.InitRequest.fields:type_name -> aiengine.InitRequest.FieldsEntry
	1,  // 5: aiengine.InitRequest.datasources:type_name -> aiengine.DataSource
	3,  // 6: aiengine.InferenceResult.response:type_name -> aiengine.Response
	2,  // 7: aiengine.AIEngine.Init:input_type -> aiengine.InitRequest
	7,  // 8: aiengine.AIEngine.AddData:input_type -> aiengine.AddDataRequest
	4,  // 9: aiengine.AIEngine.StartTraining:input_type -> aiengine.StartTrainingRequest
	5,  // 10: aiengine.AIEngine.GetInference:input_type -> aiengine.InferenceRequest
	8,  // 11: aiengine.AIEngine.GetHealth:input_type -> aiengine.HealthRequest
	3,  // 12: aiengine.AIEngine.Init:output_type -> aiengine.Response
	3,  // 13: aiengine.AIEngine.AddData:output_type -> aiengine.Response
	3,  // 14: aiengine.AIEngine.StartTraining:output_type -> aiengine.Response
	6,  // 15: aiengine.AIEngine.GetInference:output_type -> aiengine.InferenceResult
	3,  // 16: aiengine.AIEngine.GetHealth:output_type -> aiengine.Response
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_aiengine_v1_aiengine_proto_init() }
func file_proto_aiengine_v1_aiengine_proto_init() {
	if File_proto_aiengine_v1_aiengine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_aiengine_v1_aiengine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataConnector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_aiengine_v1_aiengine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_aiengine_v1_aiengine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_aiengine_v1_aiengine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_aiengine_v1_aiengine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTrainingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_aiengine_v1_aiengine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_aiengine_v1_aiengine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_aiengine_v1_aiengine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_aiengine_v1_aiengine_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_aiengine_v1_aiengine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_aiengine_v1_aiengine_proto_goTypes,
		DependencyIndexes: file_proto_aiengine_v1_aiengine_proto_depIdxs,
		MessageInfos:      file_proto_aiengine_v1_aiengine_proto_msgTypes,
	}.Build()
	File_proto_aiengine_v1_aiengine_proto = out.File
	file_proto_aiengine_v1_aiengine_proto_rawDesc = nil
	file_proto_aiengine_v1_aiengine_proto_goTypes = nil
	file_proto_aiengine_v1_aiengine_proto_depIdxs = nil
}
