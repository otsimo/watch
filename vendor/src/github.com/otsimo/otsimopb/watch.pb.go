// Code generated by protoc-gen-gogo.
// source: watch.proto
// DO NOT EDIT!

package otsimopb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WatchEvent_EventType int32

const (
	PROFILE_UPDATED     WatchEvent_EventType = 0
	CHILD_UPDATED       WatchEvent_EventType = 1
	CHILD_GAMES_UPDATED WatchEvent_EventType = 2
	CHILD_SOUND_UPDATED WatchEvent_EventType = 3
)

var WatchEvent_EventType_name = map[int32]string{
	0: "PROFILE_UPDATED",
	1: "CHILD_UPDATED",
	2: "CHILD_GAMES_UPDATED",
	3: "CHILD_SOUND_UPDATED",
}
var WatchEvent_EventType_value = map[string]int32{
	"PROFILE_UPDATED":     0,
	"CHILD_UPDATED":       1,
	"CHILD_GAMES_UPDATED": 2,
	"CHILD_SOUND_UPDATED": 3,
}

func (x WatchEvent_EventType) String() string {
	return proto.EnumName(WatchEvent_EventType_name, int32(x))
}
func (WatchEvent_EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorWatch, []int{3, 0} }

type EmitRequest struct {
	ProfileId string      `protobuf:"bytes,1,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	Event     *WatchEvent `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
}

func (m *EmitRequest) Reset()                    { *m = EmitRequest{} }
func (m *EmitRequest) String() string            { return proto.CompactTextString(m) }
func (*EmitRequest) ProtoMessage()               {}
func (*EmitRequest) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{0} }

type EmitResponse struct {
}

func (m *EmitResponse) Reset()                    { *m = EmitResponse{} }
func (m *EmitResponse) String() string            { return proto.CompactTextString(m) }
func (*EmitResponse) ProtoMessage()               {}
func (*EmitResponse) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{1} }

type WatchRequest struct {
	// profile id is for Create request
	ProfileId string `protobuf:"bytes,2,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
}

func (m *WatchRequest) Reset()                    { *m = WatchRequest{} }
func (m *WatchRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()               {}
func (*WatchRequest) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{2} }

type WatchEvent struct {
	Type      WatchEvent_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=apipb.WatchEvent_EventType" json:"type,omitempty"`
	ProfileId string               `protobuf:"bytes,2,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	ChildId   string               `protobuf:"bytes,3,opt,name=child_id,json=childId,proto3" json:"child_id,omitempty"`
	GameId    string               `protobuf:"bytes,4,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (m *WatchEvent) Reset()                    { *m = WatchEvent{} }
func (m *WatchEvent) String() string            { return proto.CompactTextString(m) }
func (*WatchEvent) ProtoMessage()               {}
func (*WatchEvent) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{3} }

type WatchResponse struct {
	Created  bool        `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Canceled bool        `protobuf:"varint,2,opt,name=canceled,proto3" json:"canceled,omitempty"`
	Event    *WatchEvent `protobuf:"bytes,3,opt,name=event" json:"event,omitempty"`
}

func (m *WatchResponse) Reset()                    { *m = WatchResponse{} }
func (m *WatchResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()               {}
func (*WatchResponse) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{4} }

func init() {
	proto.RegisterType((*EmitRequest)(nil), "apipb.EmitRequest")
	proto.RegisterType((*EmitResponse)(nil), "apipb.EmitResponse")
	proto.RegisterType((*WatchRequest)(nil), "apipb.WatchRequest")
	proto.RegisterType((*WatchEvent)(nil), "apipb.WatchEvent")
	proto.RegisterType((*WatchResponse)(nil), "apipb.WatchResponse")
	proto.RegisterEnum("apipb.WatchEvent_EventType", WatchEvent_EventType_name, WatchEvent_EventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for WatchService service

type WatchServiceClient interface {
	Emit(ctx context.Context, in *EmitRequest, opts ...grpc.CallOption) (*EmitResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (WatchService_WatchClient, error)
}

type watchServiceClient struct {
	cc *grpc.ClientConn
}

func NewWatchServiceClient(cc *grpc.ClientConn) WatchServiceClient {
	return &watchServiceClient{cc}
}

func (c *watchServiceClient) Emit(ctx context.Context, in *EmitRequest, opts ...grpc.CallOption) (*EmitResponse, error) {
	out := new(EmitResponse)
	err := grpc.Invoke(ctx, "/apipb.WatchService/Emit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchServiceClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (WatchService_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_WatchService_serviceDesc.Streams[0], c.cc, "/apipb.WatchService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchService_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type watchServiceWatchClient struct {
	grpc.ClientStream
}

func (x *watchServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for WatchService service

type WatchServiceServer interface {
	Emit(context.Context, *EmitRequest) (*EmitResponse, error)
	Watch(*WatchRequest, WatchService_WatchServer) error
}

func RegisterWatchServiceServer(s *grpc.Server, srv WatchServiceServer) {
	s.RegisterService(&_WatchService_serviceDesc, srv)
}

func _WatchService_Emit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchServiceServer).Emit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.WatchService/Emit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchServiceServer).Emit(ctx, req.(*EmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WatchService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatchServiceServer).Watch(m, &watchServiceWatchServer{stream})
}

type WatchService_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type watchServiceWatchServer struct {
	grpc.ServerStream
}

func (x *watchServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _WatchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.WatchService",
	HandlerType: (*WatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Emit",
			Handler:    _WatchService_Emit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _WatchService_Watch_Handler,
			ServerStreams: true,
		},
	},
}

func (m *EmitRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EmitRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProfileId) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.ProfileId)))
		i += copy(data[i:], m.ProfileId)
	}
	if m.Event != nil {
		data[i] = 0x12
		i++
		i = encodeVarintWatch(data, i, uint64(m.Event.Size()))
		n1, err := m.Event.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *EmitResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EmitResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *WatchRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WatchRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProfileId) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.ProfileId)))
		i += copy(data[i:], m.ProfileId)
	}
	return i, nil
}

func (m *WatchEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WatchEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintWatch(data, i, uint64(m.Type))
	}
	if len(m.ProfileId) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.ProfileId)))
		i += copy(data[i:], m.ProfileId)
	}
	if len(m.ChildId) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.ChildId)))
		i += copy(data[i:], m.ChildId)
	}
	if len(m.GameId) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.GameId)))
		i += copy(data[i:], m.GameId)
	}
	return i, nil
}

func (m *WatchResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WatchResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Created {
		data[i] = 0x8
		i++
		if m.Created {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Canceled {
		data[i] = 0x10
		i++
		if m.Canceled {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Event != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintWatch(data, i, uint64(m.Event.Size()))
		n2, err := m.Event.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Watch(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Watch(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintWatch(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *EmitRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProfileId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovWatch(uint64(l))
	}
	return n
}

func (m *EmitResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *WatchRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProfileId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	return n
}

func (m *WatchEvent) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovWatch(uint64(m.Type))
	}
	l = len(m.ProfileId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	l = len(m.ChildId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	l = len(m.GameId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	return n
}

func (m *WatchResponse) Size() (n int) {
	var l int
	_ = l
	if m.Created {
		n += 2
	}
	if m.Canceled {
		n += 2
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovWatch(uint64(l))
	}
	return n
}

func sovWatch(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWatch(x uint64) (n int) {
	return sovWatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EmitRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &WatchEvent{}
			}
			if err := m.Event.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmitResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WatchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WatchEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (WatchEvent_EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChildId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Created = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Canceled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Canceled = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &WatchEvent{}
			}
			if err := m.Event.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWatch(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWatch
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWatch
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWatch(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWatch = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWatch   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorWatch = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x6d, 0x92, 0xe9, 0x34, 0xbd, 0x9d, 0x07, 0xe3, 0x20, 0x51, 0x8a, 0xa8, 0x50, 0x36, 0xb0,
	0x99, 0x14, 0x15, 0x76, 0xac, 0x06, 0x1a, 0xa0, 0xd2, 0x40, 0xab, 0xa4, 0x15, 0x12, 0x9b, 0x51,
	0x92, 0x7a, 0x52, 0x4b, 0x4d, 0x1c, 0x12, 0x67, 0x10, 0x7f, 0xc1, 0x37, 0xf0, 0x35, 0x5d, 0xf6,
	0x13, 0x78, 0x7c, 0x04, 0x5b, 0x1c, 0x3b, 0x4d, 0x5a, 0x40, 0xb0, 0x70, 0x94, 0x73, 0x8e, 0xcf,
	0x3d, 0xd7, 0xbe, 0x86, 0xce, 0x47, 0x8f, 0x05, 0x4b, 0x2b, 0x49, 0x29, 0xa3, 0xa8, 0xe9, 0x25,
	0x24, 0xf1, 0x7b, 0xe7, 0x21, 0x61, 0xcb, 0xdc, 0xb7, 0x02, 0x1a, 0x0d, 0x42, 0x1a, 0xd2, 0x81,
	0x50, 0xfd, 0xfc, 0x5a, 0x20, 0x01, 0xc4, 0x9f, 0x74, 0x99, 0x73, 0xe8, 0xd8, 0x11, 0x61, 0x0e,
	0xfe, 0x90, 0xe3, 0x8c, 0xa1, 0xfb, 0x00, 0x9c, 0xbf, 0x26, 0x2b, 0x7c, 0x45, 0x16, 0x5d, 0xe5,
	0x81, 0xf2, 0xa8, 0xed, 0xb4, 0x4b, 0x66, 0xbc, 0x40, 0x0f, 0xa1, 0x89, 0x6f, 0x70, 0xcc, 0xba,
	0x2a, 0x57, 0x3a, 0xc3, 0x33, 0x4b, 0x64, 0x5a, 0xef, 0x8a, 0x36, 0xec, 0x42, 0x70, 0xa4, 0x6e,
	0x9e, 0xc0, 0x91, 0x2c, 0x9b, 0x25, 0x34, 0xce, 0xb0, 0x79, 0x0e, 0x47, 0x62, 0xd3, 0xdf, 0x73,
	0xd4, 0xdf, 0x72, 0xcc, 0x9f, 0x0a, 0x40, 0x5d, 0x14, 0x0d, 0xe0, 0x80, 0x7d, 0x4a, 0xb0, 0xe8,
	0xe7, 0x64, 0x78, 0xef, 0x8f, 0x54, 0x4b, 0x7c, 0x67, 0x7c, 0x8b, 0x23, 0x36, 0xfe, 0xa7, 0x3c,
	0xba, 0x0b, 0x7a, 0xb0, 0x24, 0xab, 0x45, 0x21, 0x6a, 0x42, 0x6c, 0x09, 0xcc, 0xa5, 0x3b, 0xd0,
	0x0a, 0xbd, 0x48, 0xd8, 0x0e, 0x84, 0x72, 0x58, 0x40, 0xde, 0x12, 0x86, 0x76, 0x95, 0x82, 0x0c,
	0x38, 0x9d, 0x3a, 0x93, 0x97, 0xe3, 0x4b, 0xfb, 0x6a, 0x3e, 0x1d, 0x5d, 0xcc, 0xec, 0xd1, 0xad,
	0x06, 0x3a, 0x83, 0xe3, 0x17, 0xaf, 0xc7, 0x97, 0xa3, 0x8a, 0x52, 0x78, 0x35, 0x43, 0x52, 0xaf,
	0x2e, 0xde, 0xd8, 0x6e, 0x25, 0xa8, 0xb5, 0xe0, 0x4e, 0xe6, 0x6f, 0x6b, 0x87, 0x66, 0xc6, 0x70,
	0x5c, 0x5e, 0x94, 0xbc, 0x39, 0xd4, 0x85, 0x56, 0x90, 0x62, 0x8f, 0x61, 0x39, 0x0e, 0xdd, 0xd9,
	0x42, 0xd4, 0xe3, 0xa7, 0xf0, 0xe2, 0x00, 0xaf, 0xb0, 0x3c, 0xa2, 0xee, 0x54, 0xb8, 0x1e, 0x94,
	0xf6, 0xef, 0x41, 0x0d, 0xf3, 0x72, 0x30, 0x2e, 0x4e, 0x6f, 0x48, 0x80, 0x8b, 0xab, 0x2e, 0x06,
	0x87, 0x50, 0xe9, 0xd8, 0x79, 0x1c, 0x3d, 0x63, 0x8f, 0x2b, 0xfb, 0x7b, 0x0a, 0x4d, 0x51, 0x00,
	0x19, 0xbb, 0x19, 0x5b, 0xcb, 0xed, 0x7d, 0x52, 0x7a, 0x1e, 0x2b, 0xcf, 0x9f, 0xad, 0xbf, 0xf5,
	0x1b, 0x1b, 0xbe, 0xd6, 0xdf, 0xfb, 0xca, 0x86, 0xaf, 0xaf, 0x7c, 0x7d, 0xfe, 0xd1, 0x6f, 0xc0,
	0x29, 0x7f, 0xb8, 0x16, 0x65, 0x19, 0x89, 0xa8, 0x15, 0xa6, 0x49, 0x30, 0x55, 0xde, 0xeb, 0x12,
	0x26, 0xfe, 0x17, 0x55, 0x9b, 0xcc, 0x5c, 0xff, 0x50, 0x3c, 0xdd, 0x27, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0xcc, 0x93, 0x13, 0xff, 0x02, 0x00, 0x00,
}