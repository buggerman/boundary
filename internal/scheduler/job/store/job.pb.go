// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: controller/storage/job/store/v1/job.proto

// Package store provides protobufs for storing types in the job package.

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
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

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// plugin_id is the primary key of the plugin that registered and owns the job
	// @inject_tag: `gorm:"primary_key"`
	PluginId string `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty" gorm:"primary_key"`
	// The name of the job.
	// @inject_tag: `gorm:"primary_key"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"primary_key"`
	// The human friendly description of the job.
	// @inject_tag: `gorm:"not_null"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" gorm:"not_null"`
	// next_scheduled_run is the time that the next run should be created.
	// @inject_tag: `gorm:"default:current_timestamp"`
	NextScheduledRun *timestamp.Timestamp `protobuf:"bytes,4,opt,name=next_scheduled_run,json=nextScheduledRun,proto3" json:"next_scheduled_run,omitempty" gorm:"default:current_timestamp"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_job_store_v1_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_job_store_v1_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_controller_storage_job_store_v1_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Job) GetNextScheduledRun() *timestamp.Timestamp {
	if x != nil {
		return x.NextScheduledRun
	}
	return nil
}

type JobRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// private_id is used to access the job run but not intended to be available via the API
	// @inject_tag: `gorm:"primary_key;default:null"`
	PrivateId string `protobuf:"bytes,1,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty" gorm:"primary_key;default:null"`
	// The job_plugin_id is the id of the plugin that owns the job.
	// @inject_tag: `gorm:"not_null"`
	JobPluginId string `protobuf:"bytes,2,opt,name=job_plugin_id,json=jobPluginId,proto3" json:"job_plugin_id,omitempty" gorm:"not_null"`
	// The job_name is the name of the job and must be set.
	// @inject_tag: `gorm:"not_null"`
	JobName string `protobuf:"bytes,3,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty" gorm:"not_null"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// The end_time is set when the job run completes.
	// @inject_tag: `gorm:"default:null"`
	EndTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" gorm:"default:null"`
	// completed_count is set during an update to indicate the current progress.
	// @inject_tag: `gorm:"default:0"`
	CompletedCount uint32 `protobuf:"varint,8,opt,name=completed_count,json=completedCount,proto3" json:"completed_count,omitempty" gorm:"default:0"`
	// total_count is set during an update to indicate the progress goal.
	// @inject_tag: `gorm:"default:0"`
	TotalCount uint32 `protobuf:"varint,9,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty" gorm:"default:0"`
	// status of the job run (running, completed, failed or interrupted).
	// @inject_tag: `gorm:"not_null"`
	Status string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty" gorm:"not_null"`
	// The controller_id of the controller running the job and must be set.
	// @inject_tag: `gorm:"not_null"`
	ControllerId string `protobuf:"bytes,11,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty" gorm:"not_null"`
}

func (x *JobRun) Reset() {
	*x = JobRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_job_store_v1_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRun) ProtoMessage() {}

func (x *JobRun) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_job_store_v1_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRun.ProtoReflect.Descriptor instead.
func (*JobRun) Descriptor() ([]byte, []int) {
	return file_controller_storage_job_store_v1_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobRun) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *JobRun) GetJobPluginId() string {
	if x != nil {
		return x.JobPluginId
	}
	return ""
}

func (x *JobRun) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *JobRun) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *JobRun) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *JobRun) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *JobRun) GetCompletedCount() uint32 {
	if x != nil {
		return x.CompletedCount
	}
	return 0
}

func (x *JobRun) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *JobRun) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *JobRun) GetControllerId() string {
	if x != nil {
		return x.ControllerId
	}
	return ""
}

var File_controller_storage_job_store_v1_job_proto protoreflect.FileDescriptor

var file_controller_storage_job_store_v1_job_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01,
	0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x12, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x52,
	0x75, 0x6e, 0x22, 0xd9, 0x03, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d,
	0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x42,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_job_store_v1_job_proto_rawDescOnce sync.Once
	file_controller_storage_job_store_v1_job_proto_rawDescData = file_controller_storage_job_store_v1_job_proto_rawDesc
)

func file_controller_storage_job_store_v1_job_proto_rawDescGZIP() []byte {
	file_controller_storage_job_store_v1_job_proto_rawDescOnce.Do(func() {
		file_controller_storage_job_store_v1_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_job_store_v1_job_proto_rawDescData)
	})
	return file_controller_storage_job_store_v1_job_proto_rawDescData
}

var file_controller_storage_job_store_v1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_storage_job_store_v1_job_proto_goTypes = []interface{}{
	(*Job)(nil),                 // 0: controller.storage.job.store.v1.Job
	(*JobRun)(nil),              // 1: controller.storage.job.store.v1.JobRun
	(*timestamp.Timestamp)(nil), // 2: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_job_store_v1_job_proto_depIdxs = []int32{
	2, // 0: controller.storage.job.store.v1.Job.next_scheduled_run:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 1: controller.storage.job.store.v1.JobRun.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 2: controller.storage.job.store.v1.JobRun.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 3: controller.storage.job.store.v1.JobRun.end_time:type_name -> controller.storage.timestamp.v1.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_controller_storage_job_store_v1_job_proto_init() }
func file_controller_storage_job_store_v1_job_proto_init() {
	if File_controller_storage_job_store_v1_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_job_store_v1_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_controller_storage_job_store_v1_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobRun); i {
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
			RawDescriptor: file_controller_storage_job_store_v1_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_job_store_v1_job_proto_goTypes,
		DependencyIndexes: file_controller_storage_job_store_v1_job_proto_depIdxs,
		MessageInfos:      file_controller_storage_job_store_v1_job_proto_msgTypes,
	}.Build()
	File_controller_storage_job_store_v1_job_proto = out.File
	file_controller_storage_job_store_v1_job_proto_rawDesc = nil
	file_controller_storage_job_store_v1_job_proto_goTypes = nil
	file_controller_storage_job_store_v1_job_proto_depIdxs = nil
}