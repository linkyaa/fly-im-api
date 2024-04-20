// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/im/v1/signin.proto

package v1

import (
	fmt "fmt"
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

// 登录消息
type SignIn struct {
	DeviceType DeviceType `protobuf:"varint,1,opt,name=deviceType,proto3,enum=DeviceType" json:"deviceType,omitempty"`
	UserId     int64      `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	DeviceId   int64      `protobuf:"varint,3,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Signature  []byte     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	SdkVersion int32      `protobuf:"varint,5,opt,name=sdkVersion,proto3" json:"sdkVersion,omitempty"`
}

func (m *SignIn) Reset()         { *m = SignIn{} }
func (m *SignIn) String() string { return proto.CompactTextString(m) }
func (*SignIn) ProtoMessage()    {}
func (*SignIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fbd11d3b75bbc93, []int{0}
}
func (m *SignIn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignIn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignIn.Merge(m, src)
}
func (m *SignIn) XXX_Size() int {
	return m.Size()
}
func (m *SignIn) XXX_DiscardUnknown() {
	xxx_messageInfo_SignIn.DiscardUnknown(m)
}

var xxx_messageInfo_SignIn proto.InternalMessageInfo

func (m *SignIn) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_app
}

func (m *SignIn) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SignIn) GetDeviceId() int64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *SignIn) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignIn) GetSdkVersion() int32 {
	if m != nil {
		return m.SdkVersion
	}
	return 0
}

// 登录响应
type SignInAck struct {
	Code SignInCode `protobuf:"varint,1,opt,name=code,proto3,enum=SignInCode" json:"code,omitempty"`
}

func (m *SignInAck) Reset()         { *m = SignInAck{} }
func (m *SignInAck) String() string { return proto.CompactTextString(m) }
func (*SignInAck) ProtoMessage()    {}
func (*SignInAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fbd11d3b75bbc93, []int{1}
}
func (m *SignInAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignInAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignInAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignInAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInAck.Merge(m, src)
}
func (m *SignInAck) XXX_Size() int {
	return m.Size()
}
func (m *SignInAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInAck.DiscardUnknown(m)
}

var xxx_messageInfo_SignInAck proto.InternalMessageInfo

func (m *SignInAck) GetCode() SignInCode {
	if m != nil {
		return m.Code
	}
	return SignInCode_success
}

func init() {
	proto.RegisterType((*SignIn)(nil), "SignIn")
	proto.RegisterType((*SignInAck)(nil), "SignInAck")
}

func init() { proto.RegisterFile("proto/im/v1/signin.proto", fileDescriptor_8fbd11d3b75bbc93) }

var fileDescriptor_8fbd11d3b75bbc93 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0x47, 0xeb, 0x7f, 0xdb, 0xe8, 0xdf, 0x0b, 0x62, 0xf0, 0x50, 0x59, 0x15, 0x32, 0x51, 0xa7,
	0x48, 0xd0, 0x58, 0xc0, 0xc8, 0xc4, 0xc7, 0x40, 0xd6, 0x80, 0x18, 0xd8, 0x92, 0xd8, 0x84, 0xab,
	0x34, 0x76, 0x94, 0x8f, 0x4a, 0x79, 0x0b, 0xde, 0x82, 0x57, 0x61, 0xec, 0xc8, 0x88, 0x92, 0x17,
	0x41, 0x24, 0x55, 0x9b, 0xf1, 0x77, 0xce, 0x1d, 0x8e, 0x2e, 0xb0, 0x2c, 0x37, 0xa5, 0x11, 0x98,
	0x8a, 0xcd, 0xa5, 0x28, 0x30, 0xd6, 0xa8, 0xdd, 0x0e, 0x2d, 0xe6, 0x43, 0x13, 0x19, 0xbd, 0xe3,
	0xcb, 0x4f, 0x02, 0xd6, 0x13, 0xc6, 0xda, 0xd3, 0xf4, 0x1c, 0x40, 0xaa, 0x0d, 0x46, 0xea, 0xb9,
	0xce, 0x14, 0x23, 0x36, 0x71, 0x4e, 0xae, 0x8e, 0xdc, 0x87, 0x3d, 0xf2, 0x07, 0x9a, 0xce, 0xc1,
	0xaa, 0x0a, 0x95, 0x7b, 0x92, 0xfd, 0xb3, 0x89, 0x33, 0xf6, 0x77, 0x8b, 0x2e, 0xe0, 0x7f, 0x7f,
	0xe5, 0x49, 0x36, 0xee, 0xcc, 0x7e, 0xd3, 0x53, 0x98, 0xfd, 0x35, 0x05, 0x65, 0x95, 0x2b, 0x36,
	0xb1, 0x89, 0x73, 0xec, 0x1f, 0x00, 0xe5, 0x00, 0x85, 0x4c, 0x5e, 0x54, 0x5e, 0xa0, 0xd1, 0x6c,
	0x6a, 0x13, 0x67, 0xea, 0x0f, 0xc8, 0xf2, 0x02, 0x66, 0x7d, 0xe8, 0x6d, 0x94, 0xd0, 0x33, 0x98,
	0x44, 0x46, 0x1e, 0x2a, 0x7b, 0x73, 0x6f, 0xa4, 0xf2, 0x3b, 0x71, 0xf7, 0xf8, 0xd5, 0x70, 0xb2,
	0x6d, 0x38, 0xf9, 0x69, 0x38, 0xf9, 0x68, 0xf9, 0x68, 0xdb, 0xf2, 0xd1, 0x77, 0xcb, 0x47, 0xaf,
	0x6e, 0x8c, 0xe5, 0x7b, 0x15, 0xba, 0x91, 0x49, 0xc5, 0x1a, 0x75, 0x52, 0x07, 0x81, 0x78, 0x5b,
	0xd7, 0x2b, 0x4c, 0x57, 0x41, 0x86, 0x62, 0xf0, 0xa7, 0x1b, 0x4c, 0xb3, 0x30, 0xb4, 0x3a, 0x72,
	0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xa8, 0xc9, 0x05, 0x5c, 0x01, 0x00, 0x00,
}

func (m *SignIn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignIn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignIn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SdkVersion != 0 {
		i = encodeVarintSignin(dAtA, i, uint64(m.SdkVersion))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintSignin(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if m.DeviceId != 0 {
		i = encodeVarintSignin(dAtA, i, uint64(m.DeviceId))
		i--
		dAtA[i] = 0x18
	}
	if m.UserId != 0 {
		i = encodeVarintSignin(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x10
	}
	if m.DeviceType != 0 {
		i = encodeVarintSignin(dAtA, i, uint64(m.DeviceType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SignInAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignInAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignInAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		i = encodeVarintSignin(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSignin(dAtA []byte, offset int, v uint64) int {
	offset -= sovSignin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignIn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeviceType != 0 {
		n += 1 + sovSignin(uint64(m.DeviceType))
	}
	if m.UserId != 0 {
		n += 1 + sovSignin(uint64(m.UserId))
	}
	if m.DeviceId != 0 {
		n += 1 + sovSignin(uint64(m.DeviceId))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovSignin(uint64(l))
	}
	if m.SdkVersion != 0 {
		n += 1 + sovSignin(uint64(m.SdkVersion))
	}
	return n
}

func (m *SignInAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovSignin(uint64(m.Code))
	}
	return n
}

func sovSignin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSignin(x uint64) (n int) {
	return sovSignin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

// Deserialize a specified number of fields, return early as soon as the specified fields are deserialized.
// 反序列化指定数量的字段,达到指定的次数就返回. 参数 targetFileNum <=0 直接返回,不进行反序列化
func (m *SignIn) UnmarshalWithFieldNum(dAtA []byte, targetFileNum int) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignin
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
			return fmt.Errorf("proto: SignIn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignIn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		if targetFileNum <= 0 {
			return nil
		}
		switch fieldNum {
		case 1:
			targetFileNum--
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceType", wireType)
			}
			m.DeviceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeviceType |= DeviceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			targetFileNum--
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			targetFileNum--
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
			}
			m.DeviceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeviceId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			targetFileNum--
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
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
				return ErrInvalidLengthSignin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 5:
			targetFileNum--
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdkVersion", wireType)
			}
			m.SdkVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SdkVersion |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSignin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignin
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

// Deserialize a specified number of fields, return early as soon as the specified fields are deserialized.
// 反序列化指定数量的字段,达到指定的次数就返回. 参数 targetFileNum <=0 直接返回,不进行反序列化
func (m *SignInAck) UnmarshalWithFieldNum(dAtA []byte, targetFileNum int) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignin
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
			return fmt.Errorf("proto: SignInAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignInAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		if targetFileNum <= 0 {
			return nil
		}
		switch fieldNum {
		case 1:
			targetFileNum--
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= SignInCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSignin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignin
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

// Cleanup Clean up the fields of the message.
// 清理消息的字段，该方法在消息被释放时被调用，不对切片中的字段进行处理。
func (m *SignIn) Cleanup() {
	if m == nil {
		return
	}
	m.DeviceType = 0
	m.UserId = 0
	m.DeviceId = 0
	if m.Signature != nil {
		m.Signature = m.Signature[:0]
	}
	m.SdkVersion = 0
}

// Cleanup Clean up the fields of the message.
// 清理消息的字段，该方法在消息被释放时被调用，不对切片中的字段进行处理。
func (m *SignInAck) Cleanup() {
	if m == nil {
		return
	}
	m.Code = 0
}

// DeepCleanup Clean up the fields of the message.
// 清理消息的字段，该方法在消息被释放时被调用，递归处理切片中的字段。
func (m *SignIn) DeepCleanup() {
	if m == nil {
		return
	}
	m.DeviceType = 0
	m.UserId = 0
	m.DeviceId = 0
	if m.Signature != nil {
		m.Signature = m.Signature[:0]
	}
	m.SdkVersion = 0
}

// DeepCleanup Clean up the fields of the message.
// 清理消息的字段，该方法在消息被释放时被调用，递归处理切片中的字段。
func (m *SignInAck) DeepCleanup() {
	if m == nil {
		return
	}
	m.Code = 0
}
func (m *SignIn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignin
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
			return fmt.Errorf("proto: SignIn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignIn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceType", wireType)
			}
			m.DeviceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeviceType |= DeviceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
			}
			m.DeviceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeviceId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
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
				return ErrInvalidLengthSignin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdkVersion", wireType)
			}
			m.SdkVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SdkVersion |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSignin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignin
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
func (m *SignInAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignin
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
			return fmt.Errorf("proto: SignInAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignInAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= SignInCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSignin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignin
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
func skipSignin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignin
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
					return 0, ErrIntOverflowSignin
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
					return 0, ErrIntOverflowSignin
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
				return 0, ErrInvalidLengthSignin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSignin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSignin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSignin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSignin = fmt.Errorf("proto: unexpected end of group")
)
