// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slice.proto

package events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SessionSlice is a slice of submitted chunks
type SessionSlice struct {
	// Namespace is a session namespace
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,json=namespace,proto3" json:"Namespace,omitempty"`
	// SessionID is a session ID associated with this chunk
	SessionID string `protobuf:"bytes,2,opt,name=SessionID,json=sessionID,proto3" json:"SessionID,omitempty"`
	// Chunks is a list of submitted session chunks
	Chunks []*SessionChunk `protobuf:"bytes,3,rep,name=Chunks,json=chunks" json:"Chunks,omitempty"`
	// Version specifies session slice version
	Version              int64    `protobuf:"varint,4,opt,name=Version,json=version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionSlice) Reset()         { *m = SessionSlice{} }
func (m *SessionSlice) String() string { return proto.CompactTextString(m) }
func (*SessionSlice) ProtoMessage()    {}
func (*SessionSlice) Descriptor() ([]byte, []int) {
	return fileDescriptor_slice_5be80a50d4502c85, []int{0}
}
func (m *SessionSlice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionSlice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionSlice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SessionSlice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionSlice.Merge(dst, src)
}
func (m *SessionSlice) XXX_Size() int {
	return m.Size()
}
func (m *SessionSlice) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionSlice.DiscardUnknown(m)
}

var xxx_messageInfo_SessionSlice proto.InternalMessageInfo

func (m *SessionSlice) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SessionSlice) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *SessionSlice) GetChunks() []*SessionChunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *SessionSlice) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

// SessionChunk is a chunk to be posted in the context of the session
type SessionChunk struct {
	// Time is the occurence of this event
	Time int64 `protobuf:"varint,2,opt,name=Time,json=time,proto3" json:"Time,omitempty"`
	// Data is captured data, contains event fields in case of event, session data otherwise
	Data []byte `protobuf:"bytes,3,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
	// EventType is event type
	EventType string `protobuf:"bytes,4,opt,name=EventType,json=eventType,proto3" json:"EventType,omitempty"`
	// EventIndex is the event global index
	EventIndex int64 `protobuf:"varint,5,opt,name=EventIndex,json=eventIndex,proto3" json:"EventIndex,omitempty"`
	// Index is the autoincremented chunk index
	ChunkIndex int64 `protobuf:"varint,6,opt,name=ChunkIndex,json=chunkIndex,proto3" json:"ChunkIndex,omitempty"`
	// Offset is an offset from the previous chunk in bytes
	Offset int64 `protobuf:"varint,7,opt,name=Offset,json=offset,proto3" json:"Offset,omitempty"`
	// Delay is a delay from the previous event in milliseconds
	Delay                int64    `protobuf:"varint,8,opt,name=Delay,json=delay,proto3" json:"Delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionChunk) Reset()         { *m = SessionChunk{} }
func (m *SessionChunk) String() string { return proto.CompactTextString(m) }
func (*SessionChunk) ProtoMessage()    {}
func (*SessionChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_slice_5be80a50d4502c85, []int{1}
}
func (m *SessionChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SessionChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionChunk.Merge(dst, src)
}
func (m *SessionChunk) XXX_Size() int {
	return m.Size()
}
func (m *SessionChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionChunk.DiscardUnknown(m)
}

var xxx_messageInfo_SessionChunk proto.InternalMessageInfo

func (m *SessionChunk) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *SessionChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SessionChunk) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *SessionChunk) GetEventIndex() int64 {
	if m != nil {
		return m.EventIndex
	}
	return 0
}

func (m *SessionChunk) GetChunkIndex() int64 {
	if m != nil {
		return m.ChunkIndex
	}
	return 0
}

func (m *SessionChunk) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SessionChunk) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func init() {
	proto.RegisterType((*SessionSlice)(nil), "events.SessionSlice")
	proto.RegisterType((*SessionChunk)(nil), "events.SessionChunk")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuditLog service

type AuditLogClient interface {
	SubmitSessionSlice(ctx context.Context, opts ...grpc.CallOption) (AuditLog_SubmitSessionSliceClient, error)
}

type auditLogClient struct {
	cc *grpc.ClientConn
}

func NewAuditLogClient(cc *grpc.ClientConn) AuditLogClient {
	return &auditLogClient{cc}
}

func (c *auditLogClient) SubmitSessionSlice(ctx context.Context, opts ...grpc.CallOption) (AuditLog_SubmitSessionSliceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AuditLog_serviceDesc.Streams[0], "/events.AuditLog/SubmitSessionSlice", opts...)
	if err != nil {
		return nil, err
	}
	x := &auditLogSubmitSessionSliceClient{stream}
	return x, nil
}

type AuditLog_SubmitSessionSliceClient interface {
	Send(*SessionSlice) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type auditLogSubmitSessionSliceClient struct {
	grpc.ClientStream
}

func (x *auditLogSubmitSessionSliceClient) Send(m *SessionSlice) error {
	return x.ClientStream.SendMsg(m)
}

func (x *auditLogSubmitSessionSliceClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AuditLog service

type AuditLogServer interface {
	SubmitSessionSlice(AuditLog_SubmitSessionSliceServer) error
}

func RegisterAuditLogServer(s *grpc.Server, srv AuditLogServer) {
	s.RegisterService(&_AuditLog_serviceDesc, srv)
}

func _AuditLog_SubmitSessionSlice_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuditLogServer).SubmitSessionSlice(&auditLogSubmitSessionSliceServer{stream})
}

type AuditLog_SubmitSessionSliceServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*SessionSlice, error)
	grpc.ServerStream
}

type auditLogSubmitSessionSliceServer struct {
	grpc.ServerStream
}

func (x *auditLogSubmitSessionSliceServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *auditLogSubmitSessionSliceServer) Recv() (*SessionSlice, error) {
	m := new(SessionSlice)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AuditLog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.AuditLog",
	HandlerType: (*AuditLogServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubmitSessionSlice",
			Handler:       _AuditLog_SubmitSessionSlice_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "slice.proto",
}

func (m *SessionSlice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionSlice) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSlice(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.SessionID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSlice(dAtA, i, uint64(len(m.SessionID)))
		i += copy(dAtA[i:], m.SessionID)
	}
	if len(m.Chunks) > 0 {
		for _, msg := range m.Chunks {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintSlice(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Version != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SessionChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionChunk) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.Time))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSlice(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.EventType) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSlice(dAtA, i, uint64(len(m.EventType)))
		i += copy(dAtA[i:], m.EventType)
	}
	if m.EventIndex != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.EventIndex))
	}
	if m.ChunkIndex != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.ChunkIndex))
	}
	if m.Offset != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.Offset))
	}
	if m.Delay != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.Delay))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSlice(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SessionSlice) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovSlice(uint64(l))
		}
	}
	if m.Version != 0 {
		n += 1 + sovSlice(uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SessionChunk) Size() (n int) {
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovSlice(uint64(m.Time))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	l = len(m.EventType)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	if m.EventIndex != 0 {
		n += 1 + sovSlice(uint64(m.EventIndex))
	}
	if m.ChunkIndex != 0 {
		n += 1 + sovSlice(uint64(m.ChunkIndex))
	}
	if m.Offset != 0 {
		n += 1 + sovSlice(uint64(m.Offset))
	}
	if m.Delay != 0 {
		n += 1 + sovSlice(uint64(m.Delay))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSlice(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSlice(x uint64) (n int) {
	return sovSlice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionSlice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionSlice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionSlice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, &SessionChunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SessionChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventIndex", wireType)
			}
			m.EventIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventIndex |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkIndex", wireType)
			}
			m.ChunkIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChunkIndex |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delay", wireType)
			}
			m.Delay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Delay |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSlice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSlice
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSlice
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipSlice(dAtA[start:])
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
	ErrInvalidLengthSlice = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlice   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("slice.proto", fileDescriptor_slice_5be80a50d4502c85) }

var fileDescriptor_slice_5be80a50d4502c85 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x6a, 0xf2, 0x40,
	0x10, 0xc7, 0xdd, 0x2f, 0x31, 0x9a, 0xd5, 0xc3, 0xc7, 0x22, 0xb2, 0xf8, 0x7d, 0x84, 0xe0, 0x29,
	0x87, 0x12, 0xc1, 0x3e, 0x41, 0x5b, 0x2d, 0x08, 0xa5, 0x85, 0x28, 0xbd, 0xc7, 0x64, 0xb4, 0xa1,
	0x26, 0x1b, 0xdc, 0x8d, 0xd4, 0xd7, 0xe8, 0xa9, 0xcf, 0xd3, 0x53, 0x8f, 0x7d, 0x84, 0x62, 0x5f,
	0xa4, 0xec, 0xac, 0xb1, 0xb4, 0xc7, 0xf9, 0xfd, 0xfe, 0xb3, 0x33, 0xec, 0xd0, 0x8e, 0xdc, 0x64,
	0x09, 0x84, 0xe5, 0x56, 0x28, 0xc1, 0x1c, 0xd8, 0x41, 0xa1, 0xe4, 0xe0, 0xdf, 0x5a, 0x88, 0xf5,
	0x06, 0x46, 0x48, 0x97, 0xd5, 0x6a, 0x04, 0x79, 0xa9, 0xf6, 0x26, 0x34, 0x7c, 0x26, 0xb4, 0x3b,
	0x07, 0x29, 0x33, 0x51, 0xcc, 0x75, 0x2f, 0xfb, 0x4f, 0xdd, 0xdb, 0x38, 0x07, 0x59, 0xc6, 0x09,
	0x70, 0xe2, 0x93, 0xc0, 0x8d, 0xdc, 0xa2, 0x06, 0xda, 0x1e, 0xd3, 0xb3, 0x09, 0xff, 0x63, 0xac,
	0xac, 0x01, 0x3b, 0xa3, 0xce, 0xd5, 0x43, 0x55, 0x3c, 0x4a, 0x6e, 0xf9, 0x56, 0xd0, 0x19, 0xf7,
	0x42, 0xb3, 0x42, 0x78, 0xec, 0x41, 0x19, 0x39, 0x09, 0x66, 0x18, 0xa7, 0xad, 0x7b, 0xd8, 0x6a,
	0xce, 0x6d, 0x9f, 0x04, 0x56, 0xd4, 0xda, 0x99, 0x72, 0xf8, 0xfa, 0xbd, 0x14, 0xb6, 0x30, 0x46,
	0xed, 0x45, 0x96, 0x03, 0x4e, 0xb4, 0x22, 0x5b, 0x65, 0x39, 0x68, 0x36, 0x89, 0x55, 0xcc, 0x2d,
	0x9f, 0x04, 0xdd, 0xc8, 0x4e, 0x63, 0x15, 0xeb, 0xf5, 0xa6, 0x7a, 0xe2, 0x62, 0x5f, 0x02, 0x3e,
	0xea, 0x46, 0x2e, 0xd4, 0x80, 0x79, 0x94, 0xa2, 0x9d, 0x15, 0x29, 0x3c, 0xf1, 0x26, 0xbe, 0x45,
	0xe1, 0x44, 0xb4, 0xc7, 0x71, 0xc6, 0x3b, 0xc6, 0x27, 0x27, 0xc2, 0xfa, 0xd4, 0xb9, 0x5b, 0xad,
	0x24, 0x28, 0xde, 0x42, 0xe7, 0x08, 0xac, 0x58, 0x8f, 0x36, 0x27, 0xb0, 0x89, 0xf7, 0xbc, 0x8d,
	0xb8, 0x99, 0xea, 0x62, 0x1c, 0xd1, 0xf6, 0x45, 0x95, 0x66, 0xea, 0x46, 0xac, 0xd9, 0x35, 0x65,
	0xf3, 0x6a, 0x99, 0x67, 0xea, 0xc7, 0x57, 0xff, 0xfe, 0x1e, 0xa4, 0x83, 0x7e, 0x68, 0xee, 0x15,
	0xd6, 0xf7, 0x0a, 0xa7, 0xfa, 0x5e, 0xc3, 0x46, 0x40, 0x2e, 0xff, 0xbe, 0x1d, 0x3c, 0xf2, 0x7e,
	0xf0, 0xc8, 0xc7, 0xc1, 0x23, 0x2f, 0x9f, 0x5e, 0x63, 0xe9, 0x60, 0xea, 0xfc, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x7f, 0x3c, 0xc0, 0x23, 0xfa, 0x01, 0x00, 0x00,
}
