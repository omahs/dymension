// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/delayedack/rollapp_packet.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RollappPacket_Status int32

const (
	RollappPacket_PENDING  RollappPacket_Status = 0
	RollappPacket_ACCEPTED RollappPacket_Status = 1
	RollappPacket_REJECTED RollappPacket_Status = 2
)

var RollappPacket_Status_name = map[int32]string{
	0: "PENDING",
	1: "ACCEPTED",
	2: "REJECTED",
}

var RollappPacket_Status_value = map[string]int32{
	"PENDING":  0,
	"ACCEPTED": 1,
	"REJECTED": 2,
}

func (x RollappPacket_Status) String() string {
	return proto.EnumName(RollappPacket_Status_name, int32(x))
}

func (RollappPacket_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b801a01bc7222719, []int{0, 0}
}

type RollappPacket struct {
	Packet      *types.Packet        `protobuf:"bytes,1,opt,name=packet,proto3" json:"packet,omitempty"`
	Status      RollappPacket_Status `protobuf:"varint,2,opt,name=status,proto3,enum=dymensionxyz.dymension.delayedack.RollappPacket_Status" json:"status,omitempty"`
	ProofHeight uint64               `protobuf:"varint,3,opt,name=ProofHeight,proto3" json:"ProofHeight,omitempty"`
	Error       string               `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Relayer     []byte               `protobuf:"bytes,5,opt,name=relayer,proto3" json:"relayer,omitempty"`
}

func (m *RollappPacket) Reset()         { *m = RollappPacket{} }
func (m *RollappPacket) String() string { return proto.CompactTextString(m) }
func (*RollappPacket) ProtoMessage()    {}
func (*RollappPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_b801a01bc7222719, []int{0}
}
func (m *RollappPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RollappPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RollappPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RollappPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollappPacket.Merge(m, src)
}
func (m *RollappPacket) XXX_Size() int {
	return m.Size()
}
func (m *RollappPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_RollappPacket.DiscardUnknown(m)
}

var xxx_messageInfo_RollappPacket proto.InternalMessageInfo

func (m *RollappPacket) GetPacket() *types.Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *RollappPacket) GetStatus() RollappPacket_Status {
	if m != nil {
		return m.Status
	}
	return RollappPacket_PENDING
}

func (m *RollappPacket) GetProofHeight() uint64 {
	if m != nil {
		return m.ProofHeight
	}
	return 0
}

func (m *RollappPacket) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RollappPacket) GetRelayer() []byte {
	if m != nil {
		return m.Relayer
	}
	return nil
}

func init() {
	proto.RegisterEnum("dymensionxyz.dymension.delayedack.RollappPacket_Status", RollappPacket_Status_name, RollappPacket_Status_value)
	proto.RegisterType((*RollappPacket)(nil), "dymensionxyz.dymension.delayedack.RollappPacket")
}

func init() {
	proto.RegisterFile("dymension/delayedack/rollapp_packet.proto", fileDescriptor_b801a01bc7222719)
}

var fileDescriptor_b801a01bc7222719 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbd, 0x6e, 0xea, 0x30,
	0x18, 0x86, 0x63, 0x0e, 0x84, 0x73, 0x0c, 0xe7, 0x08, 0x59, 0x0c, 0x11, 0x47, 0x8a, 0x02, 0x53,
	0xba, 0xd8, 0x02, 0x06, 0xe6, 0x16, 0xa2, 0xfe, 0x0c, 0x34, 0x4d, 0x3b, 0x75, 0xa9, 0x92, 0xe0,
	0x86, 0x88, 0x10, 0x47, 0x8e, 0x41, 0xa4, 0x57, 0xd1, 0xb9, 0x57, 0xd4, 0x91, 0xb1, 0x63, 0x05,
	0x37, 0x52, 0x25, 0xe6, 0xaf, 0x53, 0xb7, 0xef, 0x89, 0xde, 0xf7, 0x7b, 0x62, 0x1b, 0x9e, 0x4d,
	0xb2, 0x39, 0x8d, 0xd3, 0x90, 0xc5, 0x64, 0x42, 0x23, 0x37, 0xa3, 0x13, 0xd7, 0x9f, 0x11, 0xce,
	0xa2, 0xc8, 0x4d, 0x92, 0xa7, 0xc4, 0xf5, 0x67, 0x54, 0xe0, 0x84, 0x33, 0xc1, 0x50, 0xfb, 0x10,
	0x5d, 0x65, 0x2f, 0xf8, 0x00, 0xf8, 0xd8, 0x6b, 0x35, 0x03, 0x16, 0xb0, 0x22, 0x4d, 0xf2, 0x49,
	0x16, 0x5b, 0xed, 0xd0, 0xf3, 0x89, 0xcf, 0x38, 0x25, 0xfe, 0xd4, 0x8d, 0x63, 0x1a, 0x91, 0x65,
	0x77, 0x3f, 0xca, 0x48, 0xe7, 0xad, 0x04, 0xff, 0x3a, 0x52, 0x6a, 0x17, 0x4e, 0xd4, 0x87, 0xaa,
	0xb4, 0x6b, 0xc0, 0x00, 0x66, 0xad, 0xf7, 0x1f, 0x87, 0x9e, 0x8f, 0xf3, 0x2d, 0x78, 0x5f, 0x5d,
	0x76, 0xb1, 0x0c, 0x3b, 0xbb, 0x28, 0xba, 0x85, 0x6a, 0x2a, 0x5c, 0xb1, 0x48, 0xb5, 0x92, 0x01,
	0xcc, 0x7f, 0xbd, 0x01, 0xfe, 0xf1, 0x9f, 0xf1, 0x37, 0x2d, 0xbe, 0x2f, 0xea, 0xce, 0x6e, 0x0d,
	0x32, 0x60, 0xcd, 0xe6, 0x8c, 0x3d, 0x5f, 0xd1, 0x30, 0x98, 0x0a, 0xed, 0x97, 0x01, 0xcc, 0xb2,
	0x73, 0xfa, 0x09, 0x35, 0x61, 0x85, 0x72, 0xce, 0xb8, 0x56, 0x36, 0x80, 0xf9, 0xc7, 0x91, 0x80,
	0x34, 0x58, 0xe5, 0x85, 0x82, 0x6b, 0x15, 0x03, 0x98, 0x75, 0x67, 0x8f, 0x9d, 0x2e, 0x54, 0xa5,
	0x03, 0xd5, 0x60, 0xd5, 0xb6, 0xc6, 0xa3, 0xeb, 0xf1, 0x65, 0x43, 0x41, 0x75, 0xf8, 0xfb, 0x7c,
	0x38, 0xb4, 0xec, 0x07, 0x6b, 0xd4, 0x00, 0x39, 0x39, 0xd6, 0x8d, 0x35, 0xcc, 0xa9, 0x74, 0x71,
	0xf7, 0xbe, 0xd1, 0xc1, 0x7a, 0xa3, 0x83, 0xcf, 0x8d, 0x0e, 0x5e, 0xb7, 0xba, 0xb2, 0xde, 0xea,
	0xca, 0xc7, 0x56, 0x57, 0x1e, 0x07, 0x41, 0x28, 0xa6, 0x0b, 0x0f, 0xfb, 0x6c, 0x4e, 0x4e, 0x4f,
	0x7a, 0x04, 0xb2, 0xec, 0x93, 0xd5, 0xe9, 0xd3, 0x8a, 0x2c, 0xa1, 0xa9, 0xa7, 0x16, 0xd7, 0xde,
	0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x5e, 0x9c, 0xfb, 0xff, 0x01, 0x00, 0x00,
}

func (m *RollappPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RollappPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RollappPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayer) > 0 {
		i -= len(m.Relayer)
		copy(dAtA[i:], m.Relayer)
		i = encodeVarintRollappPacket(dAtA, i, uint64(len(m.Relayer)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintRollappPacket(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x22
	}
	if m.ProofHeight != 0 {
		i = encodeVarintRollappPacket(dAtA, i, uint64(m.ProofHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.Status != 0 {
		i = encodeVarintRollappPacket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.Packet != nil {
		{
			size, err := m.Packet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRollappPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRollappPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovRollappPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RollappPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		l = m.Packet.Size()
		n += 1 + l + sovRollappPacket(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovRollappPacket(uint64(m.Status))
	}
	if m.ProofHeight != 0 {
		n += 1 + sovRollappPacket(uint64(m.ProofHeight))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovRollappPacket(uint64(l))
	}
	l = len(m.Relayer)
	if l > 0 {
		n += 1 + l + sovRollappPacket(uint64(l))
	}
	return n
}

func sovRollappPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRollappPacket(x uint64) (n int) {
	return sovRollappPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RollappPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRollappPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RollappPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RollappPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Packet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollappPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRollappPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollappPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Packet == nil {
				m.Packet = &types.Packet{}
			}
			if err := m.Packet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollappPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= RollappPacket_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofHeight", wireType)
			}
			m.ProofHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollappPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProofHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollappPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRollappPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollappPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollappPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRollappPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRollappPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayer = append(m.Relayer[:0], dAtA[iNdEx:postIndex]...)
			if m.Relayer == nil {
				m.Relayer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRollappPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRollappPacket
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
func skipRollappPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRollappPacket
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
					return 0, ErrIntOverflowRollappPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRollappPacket
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
			if length < 0 {
				return 0, ErrInvalidLengthRollappPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRollappPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRollappPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRollappPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRollappPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRollappPacket = fmt.Errorf("proto: unexpected end of group")
)
