// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/nft/blockchain/codec.proto

/*
	Package blockchain is a generated protocol buffer package.

	It is generated from these files:
		x/nft/blockchain/codec.proto

	It has these top-level messages:
		BlockchainToken
		IssueTokenMsg
		UpdateTokenMsg
		TokenDetails
		Node
*/
package blockchain

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import nft "github.com/iov-one/weave/x/nft"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BlockchainToken struct {
	Base    *nft.NonFungibleToken `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	Details *TokenDetails         `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
}

func (m *BlockchainToken) Reset()                    { *m = BlockchainToken{} }
func (m *BlockchainToken) String() string            { return proto.CompactTextString(m) }
func (*BlockchainToken) ProtoMessage()               {}
func (*BlockchainToken) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{0} }

func (m *BlockchainToken) GetBase() *nft.NonFungibleToken {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *BlockchainToken) GetDetails() *TokenDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

type IssueTokenMsg struct {
	Owner     []byte          `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Id        []byte          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Details   TokenDetails    `protobuf:"bytes,3,opt,name=details" json:"details"`
	Approvals []*nft.Approval `protobuf:"bytes,4,rep,name=approvals" json:"approvals,omitempty"`
}

func (m *IssueTokenMsg) Reset()                    { *m = IssueTokenMsg{} }
func (m *IssueTokenMsg) String() string            { return proto.CompactTextString(m) }
func (*IssueTokenMsg) ProtoMessage()               {}
func (*IssueTokenMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{1} }

func (m *IssueTokenMsg) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *IssueTokenMsg) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *IssueTokenMsg) GetDetails() TokenDetails {
	if m != nil {
		return m.Details
	}
	return TokenDetails{}
}

func (m *IssueTokenMsg) GetApprovals() []*nft.Approval {
	if m != nil {
		return m.Approvals
	}
	return nil
}

type UpdateTokenMsg struct {
	Actor      []byte       `protobuf:"bytes,1,opt,name=actor,proto3" json:"actor,omitempty"`
	Id         []byte       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NewDetails TokenDetails `protobuf:"bytes,3,opt,name=newDetails" json:"newDetails"`
}

func (m *UpdateTokenMsg) Reset()                    { *m = UpdateTokenMsg{} }
func (m *UpdateTokenMsg) String() string            { return proto.CompactTextString(m) }
func (*UpdateTokenMsg) ProtoMessage()               {}
func (*UpdateTokenMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{2} }

func (m *UpdateTokenMsg) GetActor() []byte {
	if m != nil {
		return m.Actor
	}
	return nil
}

func (m *UpdateTokenMsg) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UpdateTokenMsg) GetNewDetails() TokenDetails {
	if m != nil {
		return m.NewDetails
	}
	return TokenDetails{}
}

type TokenDetails struct {
	Nodes   []Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes"`
	ChainID []byte `protobuf:"bytes,2,opt,name=chainID,proto3" json:"chainID,omitempty"`
}

func (m *TokenDetails) Reset()                    { *m = TokenDetails{} }
func (m *TokenDetails) String() string            { return proto.CompactTextString(m) }
func (*TokenDetails) ProtoMessage()               {}
func (*TokenDetails) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{3} }

func (m *TokenDetails) GetNodes() []Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *TokenDetails) GetChainID() []byte {
	if m != nil {
		return m.ChainID
	}
	return nil
}

type Node struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{4} }

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*BlockchainToken)(nil), "blockchain.BlockchainToken")
	proto.RegisterType((*IssueTokenMsg)(nil), "blockchain.IssueTokenMsg")
	proto.RegisterType((*UpdateTokenMsg)(nil), "blockchain.UpdateTokenMsg")
	proto.RegisterType((*TokenDetails)(nil), "blockchain.TokenDetails")
	proto.RegisterType((*Node)(nil), "blockchain.Node")
}
func (m *BlockchainToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockchainToken) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Base != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Base.Size()))
		n1, err := m.Base.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Details != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Details.Size()))
		n2, err := m.Details.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *IssueTokenMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IssueTokenMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintCodec(dAtA, i, uint64(m.Details.Size()))
	n3, err := m.Details.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Approvals) > 0 {
		for _, msg := range m.Approvals {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *UpdateTokenMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateTokenMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Actor) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Actor)))
		i += copy(dAtA[i:], m.Actor)
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintCodec(dAtA, i, uint64(m.NewDetails.Size()))
	n4, err := m.NewDetails.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *TokenDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, msg := range m.Nodes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ChainID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.ChainID)))
		i += copy(dAtA[i:], m.ChainID)
	}
	return i, nil
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BlockchainToken) Size() (n int) {
	var l int
	_ = l
	if m.Base != nil {
		l = m.Base.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Details != nil {
		l = m.Details.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *IssueTokenMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = m.Details.Size()
	n += 1 + l + sovCodec(uint64(l))
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *UpdateTokenMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Actor)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = m.NewDetails.Size()
	n += 1 + l + sovCodec(uint64(l))
	return n
}

func (m *TokenDetails) Size() (n int) {
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *Node) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockchainToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: BlockchainToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockchainToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Base == nil {
				m.Base = &nft.NonFungibleToken{}
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Details == nil {
				m.Details = &TokenDetails{}
			}
			if err := m.Details.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *IssueTokenMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: IssueTokenMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueTokenMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Details.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, &nft.Approval{})
			if err := m.Approvals[len(m.Approvals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *UpdateTokenMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: UpdateTokenMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateTokenMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actor", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actor = append(m.Actor[:0], dAtA[iNdEx:postIndex]...)
			if m.Actor == nil {
				m.Actor = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NewDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *TokenDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: TokenDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = append(m.ChainID[:0], dAtA[iNdEx:postIndex]...)
			if m.ChainID == nil {
				m.ChainID = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
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
				next, err := skipCodec(dAtA[start:])
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
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("x/nft/blockchain/codec.proto", fileDescriptorCodec) }

var fileDescriptorCodec = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6a, 0xe2, 0x50,
	0x14, 0xc6, 0x8d, 0xc6, 0x11, 0x8f, 0x7f, 0x46, 0x2e, 0x33, 0x10, 0x64, 0xc8, 0x48, 0x56, 0xce,
	0x1f, 0x13, 0x70, 0x36, 0xb3, 0x2a, 0x54, 0xa4, 0xe0, 0xa2, 0x2e, 0x42, 0xdb, 0xfd, 0x4d, 0x72,
	0x8d, 0xc1, 0x78, 0x4f, 0x30, 0x7f, 0xec, 0x63, 0xf4, 0x01, 0xfa, 0x40, 0x2e, 0xfb, 0x04, 0xa5,
	0xd8, 0x17, 0x29, 0x3d, 0x51, 0x93, 0x52, 0x28, 0x74, 0x77, 0xbf, 0x73, 0x7e, 0xf7, 0xfb, 0xbe,
	0x90, 0x0b, 0x3f, 0x6e, 0x2d, 0xb9, 0x48, 0x2c, 0x27, 0x44, 0x77, 0xe5, 0x2e, 0x79, 0x20, 0x2d,
	0x17, 0x3d, 0xe1, 0x9a, 0xd1, 0x06, 0x13, 0x64, 0x50, 0xcc, 0xfb, 0xbf, 0xfd, 0x20, 0x59, 0xa6,
	0x8e, 0xe9, 0xe2, 0xda, 0x0a, 0x30, 0x1b, 0xa1, 0x14, 0xd6, 0x56, 0xf0, 0x4c, 0x58, 0xb9, 0x45,
	0xe9, 0x5e, 0x7f, 0x54, 0x62, 0x7d, 0xf4, 0xd1, 0xa2, 0xb1, 0x93, 0x2e, 0x48, 0x91, 0xa0, 0x53,
	0x8e, 0x1b, 0x11, 0x7c, 0x9d, 0x9c, 0x82, 0xae, 0x70, 0x25, 0x24, 0xfb, 0x05, 0xaa, 0xc3, 0x63,
	0xa1, 0x29, 0x03, 0x65, 0xd8, 0x1a, 0x7f, 0x37, 0xe5, 0x22, 0x31, 0xe7, 0x28, 0x2f, 0x52, 0xe9,
	0x07, 0x4e, 0x28, 0x08, 0xb2, 0x09, 0x61, 0x63, 0x68, 0x78, 0x22, 0xe1, 0x41, 0x18, 0x6b, 0x55,
	0xa2, 0x35, 0xb3, 0xa8, 0x6d, 0x12, 0x39, 0xcd, 0xf7, 0xf6, 0x11, 0x34, 0xee, 0x15, 0xe8, 0xcc,
	0xe2, 0x38, 0xcd, 0x8d, 0x2e, 0x63, 0x9f, 0x7d, 0x83, 0x3a, 0x6e, 0xa5, 0xd8, 0x50, 0x62, 0xdb,
	0xce, 0x05, 0xeb, 0x42, 0x35, 0xf0, 0xc8, 0xb6, 0x6d, 0x57, 0x03, 0x8f, 0xfd, 0x2f, 0xb2, 0x6a,
	0x1f, 0x67, 0x4d, 0xd4, 0xdd, 0xe3, 0xcf, 0xca, 0x29, 0x91, 0xfd, 0x81, 0x26, 0x8f, 0xa2, 0x0d,
	0x66, 0x3c, 0x8c, 0x35, 0x75, 0x50, 0x1b, 0xb6, 0xc6, 0x1d, 0xfa, 0xaa, 0xf3, 0xc3, 0xd4, 0x2e,
	0xf6, 0x46, 0x06, 0xdd, 0xeb, 0xc8, 0xe3, 0xc9, 0x9b, 0x7a, 0xdc, 0x4d, 0xf0, 0x54, 0x8f, 0xc4,
	0xbb, 0x7a, 0x67, 0x00, 0x52, 0x6c, 0xa7, 0x9f, 0x6a, 0x58, 0xba, 0x61, 0xdc, 0x40, 0xbb, 0x4c,
	0xb0, 0xbf, 0x50, 0x97, 0xe8, 0x89, 0x58, 0x53, 0xa8, 0x70, 0xaf, 0x6c, 0x35, 0x47, 0x4f, 0x1c,
	0x2c, 0x72, 0x88, 0x69, 0xd0, 0xa0, 0xd5, 0x6c, 0x7a, 0xa8, 0x74, 0x94, 0x46, 0x1f, 0xd4, 0x57,
	0x9c, 0x31, 0x50, 0x25, 0x5f, 0xe7, 0x7f, 0xb5, 0x69, 0xd3, 0x79, 0xd2, 0xdb, 0xed, 0x75, 0xe5,
	0x61, 0xaf, 0x2b, 0x4f, 0x7b, 0x5d, 0xb9, 0x7b, 0xd6, 0x2b, 0xce, 0x17, 0x7a, 0x15, 0xff, 0x5e,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xef, 0xae, 0x90, 0xc2, 0x9c, 0x02, 0x00, 0x00,
}
