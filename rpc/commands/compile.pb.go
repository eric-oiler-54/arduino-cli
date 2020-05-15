// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands/compile.proto

package commands

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CompileReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Fqbn                 string    `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	SketchPath           string    `protobuf:"bytes,3,opt,name=sketchPath,proto3" json:"sketchPath,omitempty"`
	ShowProperties       bool      `protobuf:"varint,4,opt,name=showProperties,proto3" json:"showProperties,omitempty"`
	Preprocess           bool      `protobuf:"varint,5,opt,name=preprocess,proto3" json:"preprocess,omitempty"`
	BuildCachePath       string    `protobuf:"bytes,6,opt,name=buildCachePath,proto3" json:"buildCachePath,omitempty"`
	BuildPath            string    `protobuf:"bytes,7,opt,name=buildPath,proto3" json:"buildPath,omitempty"`
	BuildProperties      []string  `protobuf:"bytes,8,rep,name=buildProperties,proto3" json:"buildProperties,omitempty"`
	Warnings             string    `protobuf:"bytes,9,opt,name=warnings,proto3" json:"warnings,omitempty"`
	Verbose              bool      `protobuf:"varint,10,opt,name=verbose,proto3" json:"verbose,omitempty"`
	Quiet                bool      `protobuf:"varint,11,opt,name=quiet,proto3" json:"quiet,omitempty"`
	VidPid               string    `protobuf:"bytes,12,opt,name=vidPid,proto3" json:"vidPid,omitempty"`
	ExportFile           string    `protobuf:"bytes,13,opt,name=exportFile,proto3" json:"exportFile,omitempty"` // Deprecated: Do not use.
	Jobs                 int32     `protobuf:"varint,14,opt,name=jobs,proto3" json:"jobs,omitempty"`
	Libraries            []string  `protobuf:"bytes,15,rep,name=libraries,proto3" json:"libraries,omitempty"`
	OptimizeForDebug     bool      `protobuf:"varint,16,opt,name=optimizeForDebug,proto3" json:"optimizeForDebug,omitempty"`
	DryRun               bool      `protobuf:"varint,17,opt,name=dryRun,proto3" json:"dryRun,omitempty"`
	ExportDir            string    `protobuf:"bytes,18,opt,name=export_dir,json=exportDir,proto3" json:"export_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CompileReq) Reset()         { *m = CompileReq{} }
func (m *CompileReq) String() string { return proto.CompactTextString(m) }
func (*CompileReq) ProtoMessage()    {}
func (*CompileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bc582849c76c3d, []int{0}
}

func (m *CompileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompileReq.Unmarshal(m, b)
}
func (m *CompileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompileReq.Marshal(b, m, deterministic)
}
func (m *CompileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompileReq.Merge(m, src)
}
func (m *CompileReq) XXX_Size() int {
	return xxx_messageInfo_CompileReq.Size(m)
}
func (m *CompileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CompileReq.DiscardUnknown(m)
}

var xxx_messageInfo_CompileReq proto.InternalMessageInfo

func (m *CompileReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *CompileReq) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

func (m *CompileReq) GetSketchPath() string {
	if m != nil {
		return m.SketchPath
	}
	return ""
}

func (m *CompileReq) GetShowProperties() bool {
	if m != nil {
		return m.ShowProperties
	}
	return false
}

func (m *CompileReq) GetPreprocess() bool {
	if m != nil {
		return m.Preprocess
	}
	return false
}

func (m *CompileReq) GetBuildCachePath() string {
	if m != nil {
		return m.BuildCachePath
	}
	return ""
}

func (m *CompileReq) GetBuildPath() string {
	if m != nil {
		return m.BuildPath
	}
	return ""
}

func (m *CompileReq) GetBuildProperties() []string {
	if m != nil {
		return m.BuildProperties
	}
	return nil
}

func (m *CompileReq) GetWarnings() string {
	if m != nil {
		return m.Warnings
	}
	return ""
}

func (m *CompileReq) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *CompileReq) GetQuiet() bool {
	if m != nil {
		return m.Quiet
	}
	return false
}

func (m *CompileReq) GetVidPid() string {
	if m != nil {
		return m.VidPid
	}
	return ""
}

// Deprecated: Do not use.
func (m *CompileReq) GetExportFile() string {
	if m != nil {
		return m.ExportFile
	}
	return ""
}

func (m *CompileReq) GetJobs() int32 {
	if m != nil {
		return m.Jobs
	}
	return 0
}

func (m *CompileReq) GetLibraries() []string {
	if m != nil {
		return m.Libraries
	}
	return nil
}

func (m *CompileReq) GetOptimizeForDebug() bool {
	if m != nil {
		return m.OptimizeForDebug
	}
	return false
}

func (m *CompileReq) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *CompileReq) GetExportDir() string {
	if m != nil {
		return m.ExportDir
	}
	return ""
}

type CompileResp struct {
	OutStream            []byte   `protobuf:"bytes,1,opt,name=out_stream,json=outStream,proto3" json:"out_stream,omitempty"`
	ErrStream            []byte   `protobuf:"bytes,2,opt,name=err_stream,json=errStream,proto3" json:"err_stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompileResp) Reset()         { *m = CompileResp{} }
func (m *CompileResp) String() string { return proto.CompactTextString(m) }
func (*CompileResp) ProtoMessage()    {}
func (*CompileResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_86bc582849c76c3d, []int{1}
}

func (m *CompileResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompileResp.Unmarshal(m, b)
}
func (m *CompileResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompileResp.Marshal(b, m, deterministic)
}
func (m *CompileResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompileResp.Merge(m, src)
}
func (m *CompileResp) XXX_Size() int {
	return xxx_messageInfo_CompileResp.Size(m)
}
func (m *CompileResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CompileResp.DiscardUnknown(m)
}

var xxx_messageInfo_CompileResp proto.InternalMessageInfo

func (m *CompileResp) GetOutStream() []byte {
	if m != nil {
		return m.OutStream
	}
	return nil
}

func (m *CompileResp) GetErrStream() []byte {
	if m != nil {
		return m.ErrStream
	}
	return nil
}

func init() {
	proto.RegisterType((*CompileReq)(nil), "cc.arduino.cli.commands.CompileReq")
	proto.RegisterType((*CompileResp)(nil), "cc.arduino.cli.commands.CompileResp")
}

func init() { proto.RegisterFile("commands/compile.proto", fileDescriptor_86bc582849c76c3d) }

var fileDescriptor_86bc582849c76c3d = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0xe5, 0x34, 0x49, 0xe3, 0x49, 0x69, 0xcb, 0x0a, 0xca, 0xaa, 0x02, 0x64, 0x72, 0x40,
	0x16, 0xa8, 0x8e, 0x04, 0x67, 0x2e, 0x6d, 0x55, 0x09, 0x71, 0x89, 0xcc, 0x8d, 0x4b, 0x65, 0xaf,
	0x97, 0x78, 0xc0, 0xf6, 0x3a, 0xb3, 0xeb, 0x06, 0xf8, 0x9d, 0xfc, 0x20, 0xe4, 0x71, 0x9c, 0x44,
	0x41, 0x3d, 0xc5, 0xf3, 0xcd, 0xdb, 0xb7, 0x2f, 0xab, 0x07, 0x17, 0xca, 0x94, 0x65, 0x52, 0x65,
	0x76, 0xae, 0x4c, 0x59, 0x63, 0xa1, 0xa3, 0x9a, 0x8c, 0x33, 0xe2, 0x85, 0x52, 0x51, 0x42, 0x59,
	0x83, 0x95, 0x89, 0x54, 0x81, 0x51, 0x2f, 0xbb, 0x7c, 0xbe, 0x7f, 0xa0, 0x34, 0x55, 0xa7, 0x9f,
	0xfd, 0x1d, 0x02, 0xdc, 0x74, 0x0e, 0xb1, 0x5e, 0x89, 0x4f, 0x30, 0xc1, 0xca, 0xba, 0xa4, 0x52,
	0x5a, 0x7a, 0x81, 0x17, 0x4e, 0x3f, 0xbc, 0x89, 0x1e, 0x71, 0x8c, 0x3e, 0x6f, 0x84, 0xf1, 0xf6,
	0x88, 0x10, 0x30, 0xfc, 0xbe, 0x4a, 0x2b, 0x39, 0x08, 0xbc, 0xd0, 0x8f, 0xf9, 0x5b, 0xbc, 0x06,
	0xb0, 0x3f, 0xb5, 0x53, 0xf9, 0x22, 0x71, 0xb9, 0x3c, 0xe2, 0xcd, 0x1e, 0x11, 0x6f, 0xe1, 0xd4,
	0xe6, 0x66, 0xbd, 0x20, 0x53, 0x6b, 0x72, 0xa8, 0xad, 0x1c, 0x06, 0x5e, 0x38, 0x89, 0x0f, 0x68,
	0xeb, 0x53, 0x93, 0xae, 0xc9, 0x28, 0x6d, 0xad, 0x1c, 0xb1, 0x66, 0x8f, 0xb4, 0x3e, 0x69, 0x83,
	0x45, 0x76, 0x93, 0xa8, 0x5c, 0xf3, 0x5d, 0x63, 0xbe, 0xeb, 0x80, 0x8a, 0x97, 0xe0, 0x33, 0x61,
	0xc9, 0x31, 0x4b, 0x76, 0x40, 0x84, 0x70, 0xd6, 0x0d, 0xbb, 0x38, 0x93, 0xe0, 0x28, 0xf4, 0xe3,
	0x43, 0x2c, 0x2e, 0x61, 0xb2, 0x4e, 0xa8, 0xc2, 0x6a, 0x69, 0xa5, 0xcf, 0x36, 0xdb, 0x59, 0x48,
	0x38, 0x7e, 0xd0, 0x94, 0x1a, 0xab, 0x25, 0x70, 0xd0, 0x7e, 0x14, 0xcf, 0x60, 0xb4, 0x6a, 0x50,
	0x3b, 0x39, 0x65, 0xde, 0x0d, 0xe2, 0x02, 0xc6, 0x0f, 0x98, 0x2d, 0x30, 0x93, 0x27, 0xec, 0xb4,
	0x99, 0xc4, 0x0c, 0x40, 0xff, 0xaa, 0x0d, 0xb9, 0x3b, 0x2c, 0xb4, 0x7c, 0xd2, 0xee, 0xae, 0x07,
	0xd2, 0x8b, 0xf7, 0x68, 0xfb, 0xe6, 0x3f, 0x4c, 0x6a, 0xe5, 0x69, 0xe0, 0x85, 0xa3, 0x98, 0xbf,
	0xdb, 0xff, 0x58, 0x60, 0x4a, 0x09, 0xb5, 0xf9, 0xcf, 0x38, 0xff, 0x0e, 0x88, 0x77, 0x70, 0x6e,
	0x6a, 0x87, 0x25, 0xfe, 0xd1, 0x77, 0x86, 0x6e, 0x75, 0xda, 0x2c, 0xe5, 0x39, 0xc7, 0xf9, 0x8f,
	0xb7, 0xc9, 0x32, 0xfa, 0x1d, 0x37, 0x95, 0x7c, 0xca, 0x8a, 0xcd, 0x24, 0x5e, 0xf5, 0xc9, 0xee,
	0x33, 0x24, 0x29, 0xba, 0x67, 0xec, 0xc8, 0x2d, 0xd2, 0xec, 0x0b, 0x4c, 0xb7, 0xad, 0xb2, 0x75,
	0xab, 0x36, 0x8d, 0xbb, 0xb7, 0x8e, 0x74, 0x52, 0x72, 0xb1, 0x4e, 0x62, 0xdf, 0x34, 0xee, 0x2b,
	0x03, 0x36, 0x23, 0xea, 0xd7, 0x83, 0x6e, 0xad, 0x89, 0xba, 0xf5, 0xf5, 0xd5, 0xb7, 0xf7, 0x4b,
	0x74, 0x79, 0x93, 0xb6, 0xdd, 0x9b, 0x6f, 0xba, 0xd8, 0xff, 0x5e, 0xa9, 0x02, 0xe7, 0x54, 0xab,
	0x79, 0xdf, 0xcb, 0x74, 0xcc, 0xcd, 0xfe, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x40, 0x39, 0x20,
	0x0a, 0x23, 0x03, 0x00, 0x00,
}
