// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sigchain.proto

package por

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type SigAlgo int32

const (
	SigAlgo_ECDSA SigAlgo = 0
)

var SigAlgo_name = map[int32]string{
	0: "ECDSA",
}
var SigAlgo_value = map[string]int32{
	"ECDSA": 0,
}

func (x SigAlgo) String() string {
	return proto.EnumName(SigAlgo_name, int32(x))
}
func (SigAlgo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sigchain_199507578bf67c64, []int{0}
}

type SigChainElem struct {
	Addr                 []byte   `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	NextPubkey           []byte   `protobuf:"bytes,2,opt,name=NextPubkey,proto3" json:"NextPubkey,omitempty"`
	Mining               bool     `protobuf:"varint,3,opt,name=Mining,proto3" json:"Mining,omitempty"`
	SigAlgo              SigAlgo  `protobuf:"varint,4,opt,name=SigAlgo,proto3,enum=por.SigAlgo" json:"SigAlgo,omitempty"`
	Signature            []byte   `protobuf:"bytes,5,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigChainElem) Reset()         { *m = SigChainElem{} }
func (m *SigChainElem) String() string { return proto.CompactTextString(m) }
func (*SigChainElem) ProtoMessage()    {}
func (*SigChainElem) Descriptor() ([]byte, []int) {
	return fileDescriptor_sigchain_199507578bf67c64, []int{0}
}
func (m *SigChainElem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigChainElem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigChainElem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SigChainElem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigChainElem.Merge(dst, src)
}
func (m *SigChainElem) XXX_Size() int {
	return m.Size()
}
func (m *SigChainElem) XXX_DiscardUnknown() {
	xxx_messageInfo_SigChainElem.DiscardUnknown(m)
}

var xxx_messageInfo_SigChainElem proto.InternalMessageInfo

func (m *SigChainElem) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *SigChainElem) GetNextPubkey() []byte {
	if m != nil {
		return m.NextPubkey
	}
	return nil
}

func (m *SigChainElem) GetMining() bool {
	if m != nil {
		return m.Mining
	}
	return false
}

func (m *SigChainElem) GetSigAlgo() SigAlgo {
	if m != nil {
		return m.SigAlgo
	}
	return SigAlgo_ECDSA
}

func (m *SigChainElem) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SigChain struct {
	Nonce                uint32          `protobuf:"varint,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	DataSize             uint32          `protobuf:"varint,2,opt,name=DataSize,proto3" json:"DataSize,omitempty"`
	DataHash             []byte          `protobuf:"bytes,3,opt,name=DataHash,proto3" json:"DataHash,omitempty"`
	BlockHash            []byte          `protobuf:"bytes,4,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	SrcPubkey            []byte          `protobuf:"bytes,5,opt,name=SrcPubkey,proto3" json:"SrcPubkey,omitempty"`
	DestPubkey           []byte          `protobuf:"bytes,6,opt,name=DestPubkey,proto3" json:"DestPubkey,omitempty"`
	Elems                []*SigChainElem `protobuf:"bytes,7,rep,name=Elems" json:"Elems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SigChain) Reset()         { *m = SigChain{} }
func (m *SigChain) String() string { return proto.CompactTextString(m) }
func (*SigChain) ProtoMessage()    {}
func (*SigChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_sigchain_199507578bf67c64, []int{1}
}
func (m *SigChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SigChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigChain.Merge(dst, src)
}
func (m *SigChain) XXX_Size() int {
	return m.Size()
}
func (m *SigChain) XXX_DiscardUnknown() {
	xxx_messageInfo_SigChain.DiscardUnknown(m)
}

var xxx_messageInfo_SigChain proto.InternalMessageInfo

func (m *SigChain) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *SigChain) GetDataSize() uint32 {
	if m != nil {
		return m.DataSize
	}
	return 0
}

func (m *SigChain) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

func (m *SigChain) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *SigChain) GetSrcPubkey() []byte {
	if m != nil {
		return m.SrcPubkey
	}
	return nil
}

func (m *SigChain) GetDestPubkey() []byte {
	if m != nil {
		return m.DestPubkey
	}
	return nil
}

func (m *SigChain) GetElems() []*SigChainElem {
	if m != nil {
		return m.Elems
	}
	return nil
}

func init() {
	proto.RegisterType((*SigChainElem)(nil), "por.SigChainElem")
	proto.RegisterType((*SigChain)(nil), "por.SigChain")
	proto.RegisterEnum("por.SigAlgo", SigAlgo_name, SigAlgo_value)
}
func (m *SigChainElem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigChainElem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.Addr)))
		i += copy(dAtA[i:], m.Addr)
	}
	if len(m.NextPubkey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.NextPubkey)))
		i += copy(dAtA[i:], m.NextPubkey)
	}
	if m.Mining {
		dAtA[i] = 0x18
		i++
		if m.Mining {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.SigAlgo != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(m.SigAlgo))
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SigChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigChain) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(m.Nonce))
	}
	if m.DataSize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(m.DataSize))
	}
	if len(m.DataHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.DataHash)))
		i += copy(dAtA[i:], m.DataHash)
	}
	if len(m.BlockHash) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.BlockHash)))
		i += copy(dAtA[i:], m.BlockHash)
	}
	if len(m.SrcPubkey) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.SrcPubkey)))
		i += copy(dAtA[i:], m.SrcPubkey)
	}
	if len(m.DestPubkey) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSigchain(dAtA, i, uint64(len(m.DestPubkey)))
		i += copy(dAtA[i:], m.DestPubkey)
	}
	if len(m.Elems) > 0 {
		for _, msg := range m.Elems {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintSigchain(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSigchain(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SigChainElem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.NextPubkey)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	if m.Mining {
		n += 2
	}
	if m.SigAlgo != 0 {
		n += 1 + sovSigchain(uint64(m.SigAlgo))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SigChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovSigchain(uint64(m.Nonce))
	}
	if m.DataSize != 0 {
		n += 1 + sovSigchain(uint64(m.DataSize))
	}
	l = len(m.DataHash)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.SrcPubkey)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	l = len(m.DestPubkey)
	if l > 0 {
		n += 1 + l + sovSigchain(uint64(l))
	}
	if len(m.Elems) > 0 {
		for _, e := range m.Elems {
			l = e.Size()
			n += 1 + l + sovSigchain(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSigchain(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSigchain(x uint64) (n int) {
	return sovSigchain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SigChainElem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSigchain
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
			return fmt.Errorf("proto: SigChainElem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigChainElem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
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
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = append(m.Addr[:0], dAtA[iNdEx:postIndex]...)
			if m.Addr == nil {
				m.Addr = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
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
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextPubkey = append(m.NextPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.NextPubkey == nil {
				m.NextPubkey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mining", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Mining = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigAlgo", wireType)
			}
			m.SigAlgo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigAlgo |= (SigAlgo(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
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
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSigchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSigchain
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
func (m *SigChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSigchain
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
			return fmt.Errorf("proto: SigChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataSize", wireType)
			}
			m.DataSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
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
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataHash = append(m.DataHash[:0], dAtA[iNdEx:postIndex]...)
			if m.DataHash == nil {
				m.DataHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
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
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
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
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcPubkey = append(m.SrcPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.SrcPubkey == nil {
				m.SrcPubkey = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
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
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestPubkey = append(m.DestPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.DestPubkey == nil {
				m.DestPubkey = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigchain
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
				return ErrInvalidLengthSigchain
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elems = append(m.Elems, &SigChainElem{})
			if err := m.Elems[len(m.Elems)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSigchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSigchain
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
func skipSigchain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSigchain
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
					return 0, ErrIntOverflowSigchain
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
					return 0, ErrIntOverflowSigchain
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
				return 0, ErrInvalidLengthSigchain
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSigchain
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
				next, err := skipSigchain(dAtA[start:])
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
	ErrInvalidLengthSigchain = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSigchain   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sigchain.proto", fileDescriptor_sigchain_199507578bf67c64) }

var fileDescriptor_sigchain_199507578bf67c64 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0x17, 0xb7, 0xee, 0xcf, 0xb1, 0x1b, 0x33, 0x0c, 0x09, 0x22, 0xa5, 0xec, 0x42, 0x8b,
	0x17, 0xbb, 0x98, 0x4f, 0xb0, 0xad, 0x03, 0x6f, 0x1c, 0x92, 0x3e, 0x41, 0xd7, 0x85, 0x2c, 0xac,
	0x36, 0x25, 0xed, 0x40, 0x7d, 0x12, 0x2f, 0x7d, 0x1c, 0x2f, 0x7d, 0x01, 0x41, 0xea, 0x8b, 0x48,
	0xd3, 0xa6, 0xf3, 0x2e, 0xdf, 0xef, 0x83, 0x9c, 0xef, 0x7c, 0x07, 0x46, 0x99, 0xe0, 0xd1, 0x3e,
	0x14, 0xc9, 0x2c, 0x55, 0x32, 0x97, 0xb8, 0x9d, 0x4a, 0x35, 0xfd, 0x40, 0x60, 0x07, 0x82, 0xaf,
	0x4a, 0xbe, 0x8e, 0xd9, 0x33, 0xc6, 0xd0, 0x59, 0xec, 0x76, 0x8a, 0x20, 0x17, 0x79, 0x36, 0xd5,
	0x6f, 0xec, 0x00, 0x6c, 0xd8, 0x4b, 0xfe, 0x74, 0xdc, 0x1e, 0xd8, 0x2b, 0x39, 0xd3, 0xce, 0x3f,
	0x82, 0x2f, 0xa1, 0xfb, 0x28, 0x12, 0x91, 0x70, 0xd2, 0x76, 0x91, 0xd7, 0xa7, 0xb5, 0xc2, 0x37,
	0xd0, 0x0b, 0x04, 0x5f, 0xc4, 0x5c, 0x92, 0x8e, 0x8b, 0xbc, 0xd1, 0xdc, 0x9e, 0xa5, 0x52, 0xcd,
	0x6a, 0x46, 0x8d, 0x89, 0xaf, 0x61, 0x10, 0x08, 0x9e, 0x84, 0xf9, 0x51, 0x31, 0x62, 0xe9, 0xef,
	0x4f, 0x60, 0xfa, 0x8d, 0xa0, 0x6f, 0x22, 0xe2, 0x09, 0x58, 0x1b, 0x99, 0x44, 0x4c, 0xe7, 0x1b,
	0xd2, 0x4a, 0xe0, 0x2b, 0xe8, 0xfb, 0x61, 0x1e, 0x06, 0xe2, 0x8d, 0xe9, 0x78, 0x43, 0xda, 0x68,
	0xe3, 0x3d, 0x84, 0xd9, 0x5e, 0xc7, 0xb3, 0x69, 0xa3, 0xcb, 0xc1, 0xcb, 0x58, 0x46, 0x07, 0x6d,
	0x76, 0xaa, 0xc1, 0x0d, 0xd0, 0xb1, 0x54, 0x54, 0x6f, 0x6d, 0x62, 0x19, 0x50, 0x96, 0xe2, 0xb3,
	0xcc, 0x94, 0xd2, 0xad, 0x4a, 0x39, 0x11, 0x7c, 0x0b, 0x56, 0x59, 0x68, 0x46, 0x7a, 0x6e, 0xdb,
	0x3b, 0x9f, 0x5f, 0x98, 0xd5, 0x9b, 0xaa, 0x69, 0xe5, 0xdf, 0x4d, 0x9a, 0x96, 0xf0, 0x00, 0xac,
	0xf5, 0xca, 0x0f, 0x16, 0xe3, 0xd6, 0x72, 0xfc, 0x59, 0x38, 0xe8, 0xab, 0x70, 0xd0, 0x4f, 0xe1,
	0xa0, 0xf7, 0x5f, 0xa7, 0xb5, 0xed, 0xea, 0xb3, 0xdd, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x93,
	0xa3, 0x43, 0x64, 0xc8, 0x01, 0x00, 0x00,
}