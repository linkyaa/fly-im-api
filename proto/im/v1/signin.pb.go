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

type SignInCode int32

const (
	SignInCode_success       SignInCode = 0
	SignInCode_tokenExpired  SignInCode = 1
	SignInCode_tokenInvalid  SignInCode = 2
	SignInCode_updateVersion SignInCode = 3
)

var SignInCode_name = map[int32]string{
	0: "success",
	1: "tokenExpired",
	2: "tokenInvalid",
	3: "updateVersion",
}

var SignInCode_value = map[string]int32{
	"success":       0,
	"tokenExpired":  1,
	"tokenInvalid":  2,
	"updateVersion": 3,
}

func (x SignInCode) String() string {
	return proto.EnumName(SignInCode_name, int32(x))
}

func (SignInCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8fbd11d3b75bbc93, []int{0}
}

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
	proto.RegisterEnum("SignInCode", SignInCode_name, SignInCode_value)
	proto.RegisterType((*SignIn)(nil), "SignIn")
	proto.RegisterType((*SignInAck)(nil), "SignInAck")
}

func init() { proto.RegisterFile("proto/im/v1/signin.proto", fileDescriptor_8fbd11d3b75bbc93) }

var fileDescriptor_8fbd11d3b75bbc93 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xbb, 0x4e, 0xe3, 0x40,
	0x14, 0xc6, 0x71, 0x4f, 0x6e, 0xbb, 0x39, 0xc9, 0xae, 0xbc, 0x53, 0x44, 0x56, 0xb4, 0x1a, 0xac,
	0x54, 0x16, 0x10, 0x5b, 0x40, 0x49, 0xc5, 0x4d, 0xc2, 0x1d, 0x32, 0x88, 0x82, 0xce, 0x99, 0x19,
	0xc2, 0x91, 0xe3, 0x19, 0xcb, 0x63, 0x47, 0xe4, 0x2d, 0x78, 0x0b, 0x5e, 0x85, 0x32, 0x25, 0x25,
	0x4a, 0x5e, 0x04, 0xe1, 0x44, 0x89, 0xcb, 0xf3, 0xfb, 0x5b, 0xf2, 0xa7, 0x01, 0x27, 0xcb, 0x75,
	0xa1, 0x03, 0x4c, 0x83, 0xf9, 0x49, 0x60, 0x70, 0xaa, 0x50, 0xf9, 0x15, 0x0d, 0x07, 0xf5, 0xc2,
	0xb5, 0xda, 0xfa, 0xe8, 0x9d, 0x40, 0xe7, 0x1e, 0xa7, 0x2a, 0x54, 0xf4, 0x08, 0x40, 0xc8, 0x39,
	0x72, 0xf9, 0xb0, 0xc8, 0xa4, 0x43, 0x5c, 0xe2, 0xfd, 0x3d, 0xed, 0xf9, 0xd7, 0x3b, 0x8a, 0x6a,
	0x99, 0x0e, 0xa0, 0x53, 0x1a, 0x99, 0x87, 0xc2, 0x69, 0xb8, 0xc4, 0x6b, 0x46, 0xdb, 0x8b, 0x0e,
	0xe1, 0xf7, 0xe6, 0xab, 0x50, 0x38, 0xcd, 0xaa, 0xec, 0x6e, 0xfa, 0x1f, 0xba, 0x3f, 0x9b, 0xe2,
	0xa2, 0xcc, 0xa5, 0xd3, 0x72, 0x89, 0xd7, 0x8f, 0xf6, 0x40, 0x19, 0x80, 0x11, 0xc9, 0xa3, 0xcc,
	0x0d, 0x6a, 0xe5, 0xb4, 0x5d, 0xe2, 0xb5, 0xa3, 0x9a, 0x8c, 0x8e, 0xa1, 0xbb, 0x19, 0x7a, 0xc1,
	0x13, 0x7a, 0x00, 0x2d, 0xae, 0xc5, 0x7e, 0xe5, 0xa6, 0x5c, 0x69, 0x21, 0xa3, 0x2a, 0x1c, 0xde,
	0x01, 0xec, 0x8d, 0xf6, 0xe0, 0x97, 0x29, 0x39, 0x97, 0xc6, 0xd8, 0x16, 0xb5, 0xa1, 0x5f, 0xe8,
	0x44, 0xaa, 0x9b, 0xd7, 0x0c, 0x73, 0x29, 0x6c, 0xb2, 0x93, 0x50, 0xcd, 0xe3, 0x19, 0x0a, 0xbb,
	0x41, 0xff, 0xc1, 0x9f, 0x32, 0x13, 0x71, 0x21, 0xb7, 0x7f, 0xb7, 0x9b, 0x97, 0xb7, 0x1f, 0x2b,
	0x46, 0x96, 0x2b, 0x46, 0xbe, 0x56, 0x8c, 0xbc, 0xad, 0x99, 0xb5, 0x5c, 0x33, 0xeb, 0x73, 0xcd,
	0xac, 0x27, 0x7f, 0x8a, 0xc5, 0x4b, 0x39, 0xf1, 0xb9, 0x4e, 0x83, 0x19, 0xaa, 0x64, 0x11, 0xc7,
	0xc1, 0xf3, 0x6c, 0x31, 0xc6, 0x74, 0x1c, 0x67, 0x18, 0xd4, 0x5e, 0xfe, 0x1c, 0xd3, 0x6c, 0x32,
	0xe9, 0x54, 0x72, 0xf6, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x29, 0x08, 0xf2, 0xae, 0x01, 0x00,
	0x00,
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
