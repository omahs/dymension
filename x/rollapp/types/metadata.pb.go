// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/rollapp/metadata.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type RollappMetadata struct {
	// website is the rollapp website
	Website string `protobuf:"bytes,1,opt,name=website,proto3" json:"website,omitempty"`
	// description is the rollapp description. should be limited to 512 chars
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// logo_data_uri is a base64 rep with a URI prefix to the rollapp logo. Size limited
	LogoDataUri string `protobuf:"bytes,3,opt,name=logo_data_uri,json=logoDataUri,proto3" json:"logo_data_uri,omitempty"`
	// token_logo_data_uri is a URI to the native token logo. Size limited
	TokenLogoDataUri string `protobuf:"bytes,4,opt,name=token_logo_data_uri,json=tokenLogoDataUri,proto3" json:"token_logo_data_uri,omitempty"`
	// telegram is the rollapp telegram link
	Telegram string `protobuf:"bytes,5,opt,name=telegram,proto3" json:"telegram,omitempty"`
	// x is the rollapp twitter link
	X string `protobuf:"bytes,6,opt,name=x,proto3" json:"x,omitempty"`
	// genesis_url has the genesis file
	GenesisUrl string `protobuf:"bytes,7,opt,name=genesis_url,json=genesisUrl,proto3" json:"genesis_url,omitempty"`
	// display_name is a non semantic name for displaying on gui etc. Size limited.
	DisplayName string `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// tagline is a non semantic tagline/catch-phrase. Size limited.
	Tagline string `protobuf:"bytes,9,opt,name=tagline,proto3" json:"tagline,omitempty"`
	// token_symbol is the native token symbol
	TokenSymbol string `protobuf:"bytes,10,opt,name=token_symbol,json=tokenSymbol,proto3" json:"token_symbol,omitempty"`
}

func (m *RollappMetadata) Reset()         { *m = RollappMetadata{} }
func (m *RollappMetadata) String() string { return proto.CompactTextString(m) }
func (*RollappMetadata) ProtoMessage()    {}
func (*RollappMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_481e8ecd49e74695, []int{0}
}
func (m *RollappMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RollappMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RollappMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RollappMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollappMetadata.Merge(m, src)
}
func (m *RollappMetadata) XXX_Size() int {
	return m.Size()
}
func (m *RollappMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RollappMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RollappMetadata proto.InternalMessageInfo

func (m *RollappMetadata) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *RollappMetadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RollappMetadata) GetLogoDataUri() string {
	if m != nil {
		return m.LogoDataUri
	}
	return ""
}

func (m *RollappMetadata) GetTokenLogoDataUri() string {
	if m != nil {
		return m.TokenLogoDataUri
	}
	return ""
}

func (m *RollappMetadata) GetTelegram() string {
	if m != nil {
		return m.Telegram
	}
	return ""
}

func (m *RollappMetadata) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *RollappMetadata) GetGenesisUrl() string {
	if m != nil {
		return m.GenesisUrl
	}
	return ""
}

func (m *RollappMetadata) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *RollappMetadata) GetTagline() string {
	if m != nil {
		return m.Tagline
	}
	return ""
}

func (m *RollappMetadata) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func init() {
	proto.RegisterType((*RollappMetadata)(nil), "dymensionxyz.dymension.rollapp.RollappMetadata")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/rollapp/metadata.proto", fileDescriptor_481e8ecd49e74695)
}

var fileDescriptor_481e8ecd49e74695 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbd, 0x4e, 0xeb, 0x30,
	0x14, 0xc7, 0x9b, 0xde, 0x7b, 0xfb, 0xe1, 0xf6, 0x0a, 0x64, 0x16, 0x8b, 0xc1, 0x94, 0x4e, 0x2c,
	0x4d, 0x86, 0xf2, 0x04, 0x88, 0x11, 0x3a, 0x14, 0x75, 0x61, 0x89, 0x9c, 0xe6, 0x28, 0x58, 0xf8,
	0x23, 0xb2, 0x5d, 0x48, 0x78, 0x0a, 0x1e, 0x0b, 0x89, 0xa5, 0x23, 0x23, 0x6a, 0x5f, 0x04, 0xc5,
	0x09, 0x51, 0x19, 0xd8, 0xf2, 0xff, 0x9d, 0xdf, 0x89, 0x7c, 0xf4, 0x47, 0xb3, 0xb4, 0x94, 0xa0,
	0x2c, 0xd7, 0xaa, 0x28, 0x5f, 0xa2, 0x36, 0x44, 0x46, 0x0b, 0xc1, 0xf2, 0x3c, 0x92, 0xe0, 0x58,
	0xca, 0x1c, 0x0b, 0x73, 0xa3, 0x9d, 0xc6, 0xf4, 0x50, 0x0f, 0xdb, 0x10, 0x36, 0xfa, 0xf4, 0xbd,
	0x8b, 0x8e, 0x96, 0xf5, 0xf7, 0x6d, 0xb3, 0x89, 0x09, 0xea, 0x3f, 0x43, 0x62, 0xb9, 0x03, 0x12,
	0x4c, 0x82, 0x8b, 0xe1, 0xf2, 0x3b, 0xe2, 0x09, 0x1a, 0xa5, 0x60, 0xd7, 0x86, 0xe7, 0x8e, 0x6b,
	0x45, 0xba, 0x7e, 0x7a, 0x88, 0xf0, 0x14, 0xfd, 0x17, 0x3a, 0xd3, 0x71, 0xf5, 0xa3, 0x78, 0x63,
	0x38, 0xf9, 0x53, 0x3b, 0x15, 0xbc, 0x66, 0x8e, 0xad, 0x0c, 0xc7, 0x33, 0x74, 0xe2, 0xf4, 0x23,
	0xa8, 0xf8, 0xa7, 0xf9, 0xd7, 0x9b, 0xc7, 0x7e, 0x74, 0x73, 0xa0, 0x9f, 0xa2, 0x81, 0x03, 0x01,
	0x99, 0x61, 0x92, 0xfc, 0xf3, 0x4e, 0x9b, 0xf1, 0x18, 0x05, 0x05, 0xe9, 0x79, 0x18, 0x14, 0xf8,
	0x0c, 0x8d, 0x32, 0x50, 0x60, 0xb9, 0x8d, 0x37, 0x46, 0x90, 0xbe, 0xe7, 0xa8, 0x41, 0x2b, 0x23,
	0xf0, 0x39, 0x1a, 0xa7, 0xdc, 0xe6, 0x82, 0x95, 0xb1, 0x62, 0x12, 0xc8, 0xa0, 0x39, 0xa0, 0x66,
	0x0b, 0x26, 0xa1, 0x3a, 0xde, 0xb1, 0x4c, 0x70, 0x05, 0x64, 0x58, 0x1f, 0xdf, 0xc4, 0x6a, 0xb9,
	0x7e, 0xb6, 0x2d, 0x65, 0xa2, 0x05, 0x41, 0xf5, 0xb2, 0x67, 0x77, 0x1e, 0x5d, 0x2d, 0xde, 0x76,
	0x34, 0xd8, 0xee, 0x68, 0xf0, 0xb9, 0xa3, 0xc1, 0xeb, 0x9e, 0x76, 0xb6, 0x7b, 0xda, 0xf9, 0xd8,
	0xd3, 0xce, 0xfd, 0x65, 0xc6, 0xdd, 0xc3, 0x26, 0x09, 0xd7, 0x5a, 0x46, 0xbf, 0x34, 0xf8, 0x34,
	0x8f, 0x8a, 0xb6, 0x46, 0x57, 0xe6, 0x60, 0x93, 0x9e, 0x2f, 0x71, 0xfe, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0x38, 0x45, 0x1a, 0xf5, 0x01, 0x00, 0x00,
}

func (m *RollappMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RollappMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RollappMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenSymbol) > 0 {
		i -= len(m.TokenSymbol)
		copy(dAtA[i:], m.TokenSymbol)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.TokenSymbol)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Tagline) > 0 {
		i -= len(m.Tagline)
		copy(dAtA[i:], m.Tagline)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Tagline)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.DisplayName) > 0 {
		i -= len(m.DisplayName)
		copy(dAtA[i:], m.DisplayName)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.DisplayName)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.GenesisUrl) > 0 {
		i -= len(m.GenesisUrl)
		copy(dAtA[i:], m.GenesisUrl)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.GenesisUrl)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.X) > 0 {
		i -= len(m.X)
		copy(dAtA[i:], m.X)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.X)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Telegram) > 0 {
		i -= len(m.Telegram)
		copy(dAtA[i:], m.Telegram)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Telegram)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TokenLogoDataUri) > 0 {
		i -= len(m.TokenLogoDataUri)
		copy(dAtA[i:], m.TokenLogoDataUri)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.TokenLogoDataUri)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LogoDataUri) > 0 {
		i -= len(m.LogoDataUri)
		copy(dAtA[i:], m.LogoDataUri)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.LogoDataUri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RollappMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.LogoDataUri)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.TokenLogoDataUri)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Telegram)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.X)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.GenesisUrl)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Tagline)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.TokenSymbol)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RollappMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: RollappMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RollappMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogoDataUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogoDataUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenLogoDataUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenLogoDataUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Telegram", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Telegram = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.X = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tagline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tagline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenSymbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenSymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadata = fmt.Errorf("proto: unexpected end of group")
)
