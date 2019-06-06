// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

package rpc

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	validator "github.com/hyperledger/burrow/acm/validator"
	bcm "github.com/hyperledger/burrow/bcm"
	github_com_hyperledger_burrow_binary "github.com/hyperledger/burrow/binary"
	tendermint "github.com/hyperledger/burrow/consensus/tendermint"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ResultStatus struct {
	ChainID       string                                        `protobuf:"bytes,1,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	RunID         string                                        `protobuf:"bytes,2,opt,name=RunID,proto3" json:"RunID,omitempty"`
	BurrowVersion string                                        `protobuf:"bytes,3,opt,name=BurrowVersion,proto3" json:"BurrowVersion,omitempty"`
	GenesisHash   github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,4,opt,name=GenesisHash,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"GenesisHash"`
	NodeInfo      *tendermint.NodeInfo                          `protobuf:"bytes,5,opt,name=NodeInfo,proto3" json:"NodeInfo,omitempty"`
	SyncInfo      *bcm.SyncInfo                                 `protobuf:"bytes,6,opt,name=SyncInfo,proto3" json:"SyncInfo,omitempty"`
	// When catching up in fast sync
	CatchingUp           bool                 `protobuf:"varint,8,opt,name=CatchingUp,proto3" json:""`
	ValidatorInfo        *validator.Validator `protobuf:"bytes,7,opt,name=ValidatorInfo,proto3" json:"ValidatorInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ResultStatus) Reset()         { *m = ResultStatus{} }
func (m *ResultStatus) String() string { return proto.CompactTextString(m) }
func (*ResultStatus) ProtoMessage()    {}
func (*ResultStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}
func (m *ResultStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResultStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResultStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResultStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultStatus.Merge(m, src)
}
func (m *ResultStatus) XXX_Size() int {
	return m.Size()
}
func (m *ResultStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ResultStatus proto.InternalMessageInfo

func (m *ResultStatus) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *ResultStatus) GetRunID() string {
	if m != nil {
		return m.RunID
	}
	return ""
}

func (m *ResultStatus) GetBurrowVersion() string {
	if m != nil {
		return m.BurrowVersion
	}
	return ""
}

func (m *ResultStatus) GetNodeInfo() *tendermint.NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

func (m *ResultStatus) GetSyncInfo() *bcm.SyncInfo {
	if m != nil {
		return m.SyncInfo
	}
	return nil
}

func (m *ResultStatus) GetCatchingUp() bool {
	if m != nil {
		return m.CatchingUp
	}
	return false
}

func (m *ResultStatus) GetValidatorInfo() *validator.Validator {
	if m != nil {
		return m.ValidatorInfo
	}
	return nil
}

func (*ResultStatus) XXX_MessageName() string {
	return "rpc.ResultStatus"
}
func init() {
	proto.RegisterType((*ResultStatus)(nil), "rpc.ResultStatus")
	golang_proto.RegisterType((*ResultStatus)(nil), "rpc.ResultStatus")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }
func init() { golang_proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3f, 0x6f, 0xda, 0x40,
	0x18, 0xc6, 0x39, 0xfe, 0x9a, 0x03, 0xd4, 0xea, 0xc4, 0x60, 0x31, 0x18, 0xb7, 0x62, 0x70, 0x87,
	0xda, 0x55, 0xab, 0xaa, 0x52, 0x47, 0x13, 0x29, 0xb0, 0x64, 0x38, 0x14, 0x22, 0x65, 0xf3, 0x9f,
	0xc3, 0x3e, 0x09, 0xee, 0xac, 0xf3, 0x39, 0x89, 0xbf, 0x5d, 0x46, 0xc6, 0xcc, 0x19, 0x50, 0x04,
	0x5b, 0x3e, 0x43, 0x86, 0x88, 0x03, 0x13, 0xb3, 0x64, 0x7b, 0x9f, 0xdf, 0xf3, 0xbe, 0x8f, 0xec,
	0xe7, 0x60, 0x5b, 0x24, 0x81, 0x9d, 0x08, 0x2e, 0x39, 0xaa, 0x89, 0x24, 0x18, 0xfc, 0x8c, 0xa8,
	0x8c, 0x33, 0xdf, 0x0e, 0xf8, 0xca, 0x89, 0x78, 0xc4, 0x1d, 0xe5, 0xf9, 0xd9, 0x42, 0x29, 0x25,
	0xd4, 0x74, 0xb8, 0x19, 0x7c, 0x95, 0x84, 0x85, 0x44, 0xac, 0x28, 0x93, 0x47, 0xf2, 0xe5, 0xce,
	0x5b, 0xd2, 0xd0, 0x93, 0x5c, 0x1c, 0x41, 0xdb, 0x0f, 0x56, 0x87, 0xf1, 0xfb, 0x5b, 0x15, 0x76,
	0x31, 0x49, 0xb3, 0xa5, 0x9c, 0x49, 0x4f, 0x66, 0x29, 0xd2, 0x61, 0x6b, 0x1c, 0x7b, 0x94, 0x4d,
	0x2f, 0x74, 0x60, 0x02, 0xab, 0x8d, 0x0b, 0x89, 0xfa, 0xb0, 0x81, 0xb3, 0x3d, 0xaf, 0x2a, 0x7e,
	0x10, 0x68, 0x04, 0x7b, 0x6e, 0x26, 0x04, 0xbf, 0x9f, 0x13, 0x91, 0x52, 0xce, 0xf4, 0x9a, 0x72,
	0xcf, 0x21, 0xba, 0x81, 0x9d, 0x4b, 0xc2, 0x48, 0x4a, 0xd3, 0x89, 0x97, 0xc6, 0x7a, 0xdd, 0x04,
	0x56, 0xd7, 0xfd, 0xbb, 0xde, 0x0c, 0x2b, 0xcf, 0x9b, 0x61, 0xf9, 0x07, 0xe3, 0x3c, 0x21, 0x62,
	0x49, 0xc2, 0x88, 0x08, 0xc7, 0x57, 0x11, 0x8e, 0x4f, 0x99, 0x27, 0x72, 0x7b, 0x42, 0x1e, 0xdc,
	0x5c, 0x92, 0x14, 0x97, 0x93, 0xd0, 0x2f, 0xa8, 0x5d, 0xf1, 0x90, 0x4c, 0xd9, 0x82, 0xeb, 0x0d,
	0x13, 0x58, 0x9d, 0xdf, 0x7d, 0xbb, 0x54, 0x40, 0xe1, 0xe1, 0xd3, 0x16, 0xfa, 0x01, 0xb5, 0x59,
	0xce, 0x02, 0x75, 0xd1, 0x54, 0x17, 0x3d, 0x7b, 0xdf, 0x47, 0x01, 0xf1, 0xc9, 0x46, 0x23, 0x08,
	0xc7, 0x9e, 0x0c, 0x62, 0xca, 0xa2, 0xeb, 0x44, 0xd7, 0x4c, 0x60, 0x69, 0x6e, 0xfd, 0x75, 0x33,
	0xac, 0xe0, 0x12, 0x47, 0xff, 0x61, 0x6f, 0x5e, 0x14, 0xac, 0x52, 0x5b, 0xc7, 0xef, 0xf8, 0xa8,
	0xfd, 0xe4, 0xe3, 0xf3, 0x55, 0xf7, 0xdf, 0x7a, 0x6b, 0x80, 0xa7, 0xad, 0x01, 0x5e, 0xb6, 0x06,
	0x78, 0xdc, 0x19, 0x60, 0xbd, 0x33, 0xc0, 0xed, 0xb7, 0xcf, 0x0b, 0x11, 0x49, 0xe0, 0x37, 0xd5,
	0xf3, 0xfd, 0x79, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x85, 0xc2, 0xfd, 0xda, 0x2d, 0x02, 0x00, 0x00,
}

func (m *ResultStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResultStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ChainID)))
		i += copy(dAtA[i:], m.ChainID)
	}
	if len(m.RunID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RunID)))
		i += copy(dAtA[i:], m.RunID)
	}
	if len(m.BurrowVersion) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.BurrowVersion)))
		i += copy(dAtA[i:], m.BurrowVersion)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.GenesisHash.Size()))
	n1, err := m.GenesisHash.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.NodeInfo != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.NodeInfo.Size()))
		n2, err := m.NodeInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.SyncInfo != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SyncInfo.Size()))
		n3, err := m.SyncInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.ValidatorInfo != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ValidatorInfo.Size()))
		n4, err := m.ValidatorInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.CatchingUp {
		dAtA[i] = 0x40
		i++
		if m.CatchingUp {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResultStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RunID)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.BurrowVersion)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = m.GenesisHash.Size()
	n += 1 + l + sovRpc(uint64(l))
	if m.NodeInfo != nil {
		l = m.NodeInfo.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.SyncInfo != nil {
		l = m.SyncInfo.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.ValidatorInfo != nil {
		l = m.ValidatorInfo.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.CatchingUp {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRpc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResultStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: ResultStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResultStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurrowVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurrowVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NodeInfo == nil {
				m.NodeInfo = &tendermint.NodeInfo{}
			}
			if err := m.NodeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SyncInfo == nil {
				m.SyncInfo = &bcm.SyncInfo{}
			}
			if err := m.SyncInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidatorInfo == nil {
				m.ValidatorInfo = &validator.Validator{}
			}
			if err := m.ValidatorInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CatchingUp", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CatchingUp = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRpc
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
func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
				return 0, ErrInvalidLengthRpc
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRpc
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
				next, err := skipRpc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRpc
				}
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
	ErrInvalidLengthRpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc   = fmt.Errorf("proto: integer overflow")
)
