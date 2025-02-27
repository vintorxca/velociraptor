// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: acl.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApiClientACL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllQuery             bool     `protobuf:"varint,1,opt,name=all_query,json=allQuery,proto3" json:"all_query,omitempty"`
	AnyQuery             bool     `protobuf:"varint,10,opt,name=any_query,json=anyQuery,proto3" json:"any_query,omitempty"`
	PublishQueues        []string `protobuf:"bytes,2,rep,name=publish_queues,json=publishQueues,proto3" json:"publish_queues,omitempty"`
	ReadResults          bool     `protobuf:"varint,3,opt,name=read_results,json=readResults,proto3" json:"read_results,omitempty"`
	LabelClients         bool     `protobuf:"varint,11,opt,name=label_clients,json=labelClients,proto3" json:"label_clients,omitempty"`
	CollectClient        bool     `protobuf:"varint,4,opt,name=collect_client,json=collectClient,proto3" json:"collect_client,omitempty"`
	CollectServer        bool     `protobuf:"varint,5,opt,name=collect_server,json=collectServer,proto3" json:"collect_server,omitempty"`
	ArtifactWriter       bool     `protobuf:"varint,6,opt,name=artifact_writer,json=artifactWriter,proto3" json:"artifact_writer,omitempty"`
	ServerArtifactWriter bool     `protobuf:"varint,15,opt,name=server_artifact_writer,json=serverArtifactWriter,proto3" json:"server_artifact_writer,omitempty"`
	Execve               bool     `protobuf:"varint,7,opt,name=execve,proto3" json:"execve,omitempty"`
	NotebookEditor       bool     `protobuf:"varint,8,opt,name=notebook_editor,json=notebookEditor,proto3" json:"notebook_editor,omitempty"`
	ServerAdmin          bool     `protobuf:"varint,12,opt,name=server_admin,json=serverAdmin,proto3" json:"server_admin,omitempty"`
	FilesystemRead       bool     `protobuf:"varint,13,opt,name=filesystem_read,json=filesystemRead,proto3" json:"filesystem_read,omitempty"`
	FilesystemWrite      bool     `protobuf:"varint,14,opt,name=filesystem_write,json=filesystemWrite,proto3" json:"filesystem_write,omitempty"`
	MachineState         bool     `protobuf:"varint,16,opt,name=machine_state,json=machineState,proto3" json:"machine_state,omitempty"`
	PrepareResults       bool     `protobuf:"varint,17,opt,name=prepare_results,json=prepareResults,proto3" json:"prepare_results,omitempty"`
	DatastoreAccess      bool     `protobuf:"varint,18,opt,name=datastore_access,json=datastoreAccess,proto3" json:"datastore_access,omitempty"`
	// A list of roles in lieu of the permissions above. These will be
	// interpolated into this ACL object.
	Roles []string `protobuf:"bytes,9,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *ApiClientACL) Reset() {
	*x = ApiClientACL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiClientACL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiClientACL) ProtoMessage() {}

func (x *ApiClientACL) ProtoReflect() protoreflect.Message {
	mi := &file_acl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiClientACL.ProtoReflect.Descriptor instead.
func (*ApiClientACL) Descriptor() ([]byte, []int) {
	return file_acl_proto_rawDescGZIP(), []int{0}
}

func (x *ApiClientACL) GetAllQuery() bool {
	if x != nil {
		return x.AllQuery
	}
	return false
}

func (x *ApiClientACL) GetAnyQuery() bool {
	if x != nil {
		return x.AnyQuery
	}
	return false
}

func (x *ApiClientACL) GetPublishQueues() []string {
	if x != nil {
		return x.PublishQueues
	}
	return nil
}

func (x *ApiClientACL) GetReadResults() bool {
	if x != nil {
		return x.ReadResults
	}
	return false
}

func (x *ApiClientACL) GetLabelClients() bool {
	if x != nil {
		return x.LabelClients
	}
	return false
}

func (x *ApiClientACL) GetCollectClient() bool {
	if x != nil {
		return x.CollectClient
	}
	return false
}

func (x *ApiClientACL) GetCollectServer() bool {
	if x != nil {
		return x.CollectServer
	}
	return false
}

func (x *ApiClientACL) GetArtifactWriter() bool {
	if x != nil {
		return x.ArtifactWriter
	}
	return false
}

func (x *ApiClientACL) GetServerArtifactWriter() bool {
	if x != nil {
		return x.ServerArtifactWriter
	}
	return false
}

func (x *ApiClientACL) GetExecve() bool {
	if x != nil {
		return x.Execve
	}
	return false
}

func (x *ApiClientACL) GetNotebookEditor() bool {
	if x != nil {
		return x.NotebookEditor
	}
	return false
}

func (x *ApiClientACL) GetServerAdmin() bool {
	if x != nil {
		return x.ServerAdmin
	}
	return false
}

func (x *ApiClientACL) GetFilesystemRead() bool {
	if x != nil {
		return x.FilesystemRead
	}
	return false
}

func (x *ApiClientACL) GetFilesystemWrite() bool {
	if x != nil {
		return x.FilesystemWrite
	}
	return false
}

func (x *ApiClientACL) GetMachineState() bool {
	if x != nil {
		return x.MachineState
	}
	return false
}

func (x *ApiClientACL) GetPrepareResults() bool {
	if x != nil {
		return x.PrepareResults
	}
	return false
}

func (x *ApiClientACL) GetDatastoreAccess() bool {
	if x != nil {
		return x.DatastoreAccess
	}
	return false
}

func (x *ApiClientACL) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

// A role is a named sets of ACL permissions. A user may possess
// multiple roles.
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permissions *ApiClientACL `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_acl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_acl_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetPermissions() *ApiClientACL {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_acl_proto protoreflect.FileDescriptor

var file_acl_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x06, 0x0a, 0x0c, 0x41, 0x70, 0x69,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x43, 0x4c, 0x12, 0x4b, 0x0a, 0x09, 0x61, 0x6c, 0x6c,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x2e, 0xe2, 0xfc,
	0xe3, 0xc4, 0x01, 0x28, 0x12, 0x26, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x61,
	0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x72, 0x79, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x08, 0x61, 0x6c,
	0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e, 0x79, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x6e, 0x79, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x58, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x31, 0xe2, 0xfc, 0xe3,
	0xc4, 0x01, 0x2b, 0x12, 0x29, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x63,
	0x61, 0x6e, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x2e, 0x52, 0x0d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x65, 0x63, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x65, 0x63, 0x76, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f,
	0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x45, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x51, 0x0a, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x43, 0x4c, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x32, 0x5a, 0x30,
	0x77, 0x77, 0x77, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x72,
	0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x63, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acl_proto_rawDescOnce sync.Once
	file_acl_proto_rawDescData = file_acl_proto_rawDesc
)

func file_acl_proto_rawDescGZIP() []byte {
	file_acl_proto_rawDescOnce.Do(func() {
		file_acl_proto_rawDescData = protoimpl.X.CompressGZIP(file_acl_proto_rawDescData)
	})
	return file_acl_proto_rawDescData
}

var file_acl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_acl_proto_goTypes = []interface{}{
	(*ApiClientACL)(nil), // 0: proto.ApiClientACL
	(*Role)(nil),         // 1: proto.Role
}
var file_acl_proto_depIdxs = []int32{
	0, // 0: proto.Role.permissions:type_name -> proto.ApiClientACL
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_acl_proto_init() }
func file_acl_proto_init() {
	if File_acl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiClientACL); i {
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
		file_acl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
			RawDescriptor: file_acl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acl_proto_goTypes,
		DependencyIndexes: file_acl_proto_depIdxs,
		MessageInfos:      file_acl_proto_msgTypes,
	}.Build()
	File_acl_proto = out.File
	file_acl_proto_rawDesc = nil
	file_acl_proto_goTypes = nil
	file_acl_proto_depIdxs = nil
}
