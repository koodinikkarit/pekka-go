// Code generated by protoc-gen-go.
// source: event_timer.proto
// DO NOT EDIT!

package PekkaService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EventTimer struct {
}

func (m *EventTimer) Reset()                    { *m = EventTimer{} }
func (m *EventTimer) String() string            { return proto.CompactTextString(m) }
func (*EventTimer) ProtoMessage()               {}
func (*EventTimer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterType((*EventTimer)(nil), "PekkaService.EventTimer")
}

func init() { proto.RegisterFile("event_timer.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 72 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2d, 0x4b, 0xcd,
	0x2b, 0x89, 0x2f, 0xc9, 0xcc, 0x4d, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x09,
	0x48, 0xcd, 0xce, 0x4e, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0xe2, 0xe1, 0xe2, 0x72,
	0x05, 0x29, 0x09, 0x01, 0xa9, 0x48, 0x62, 0x03, 0x2b, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x71, 0x74, 0xa6, 0x02, 0x37, 0x00, 0x00, 0x00,
}
