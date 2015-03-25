// Code generated by protoc-gen-go.
// source: process.proto
// DO NOT EDIT!

/*
Package process_utility is a generated protocol buffer package.

It is generated from these files:
	process.proto

It has these top-level messages:
	Process
	ProcessSet
*/
package process_utility

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Process struct {
	Pid              *int32  `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	Vsz              *int32  `protobuf:"varint,2,opt,name=vsz" json:"vsz,omitempty"`
	Time             *string `protobuf:"bytes,3,opt,name=time" json:"time,omitempty"`
	Comm             *string `protobuf:"bytes,4,opt,name=comm" json:"comm,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Process) Reset()         { *m = Process{} }
func (m *Process) String() string { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()    {}

func (m *Process) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *Process) GetVsz() int32 {
	if m != nil && m.Vsz != nil {
		return *m.Vsz
	}
	return 0
}

func (m *Process) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

func (m *Process) GetComm() string {
	if m != nil && m.Comm != nil {
		return *m.Comm
	}
	return ""
}

type ProcessSet struct {
	Process          []*Process `protobuf:"bytes,1,rep,name=process" json:"process,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ProcessSet) Reset()         { *m = ProcessSet{} }
func (m *ProcessSet) String() string { return proto.CompactTextString(m) }
func (*ProcessSet) ProtoMessage()    {}

func (m *ProcessSet) GetProcess() []*Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func init() {
}
