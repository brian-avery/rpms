// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/test/spyAdapter/template/apa/tmpl_handler_service.proto

/*
	Package sampleapa is a generated protocol buffer package.

	It is generated from these files:
		mixer/test/spyAdapter/template/apa/tmpl_handler_service.proto

	It has these top-level messages:
		InstanceParam
*/
package sampleapa

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "istio.io/api/mixer/adapter/model/v1beta1"

import strings "strings"
import reflect "reflect"
import sortkeys "github.com/gogo/protobuf/sortkeys"

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

// Represents instance configuration schema for 'sampleapa' template.
type InstanceParam struct {
	Int64Primitive  string `protobuf:"bytes,1,opt,name=int64Primitive,proto3" json:"int64Primitive,omitempty"`
	BoolPrimitive   string `protobuf:"bytes,2,opt,name=boolPrimitive,proto3" json:"boolPrimitive,omitempty"`
	DoublePrimitive string `protobuf:"bytes,3,opt,name=doublePrimitive,proto3" json:"doublePrimitive,omitempty"`
	StringPrimitive string `protobuf:"bytes,4,opt,name=stringPrimitive,proto3" json:"stringPrimitive,omitempty"`
	// Attribute names to expression mapping. These expressions can use the fields from the output object
	// returned by the attribute producing adapters using $out.<fieldName> notation. For example:
	// source.ip : $out.source_pod_ip
	// In the above example, source.ip attribute will be added to the existing attribute list and its value will be set to
	// the value of source_pod_ip field of the output returned by the adapter.
	AttributeBindings map[string]string `protobuf:"bytes,72295728,rep,name=attribute_bindings,json=attributeBindings" json:"attribute_bindings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *InstanceParam) Reset()                    { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage()               {}
func (*InstanceParam) Descriptor() ([]byte, []int) { return fileDescriptorTmplHandlerService, []int{0} }

func init() {
	proto.RegisterType((*InstanceParam)(nil), "sampleapa.InstanceParam")
}
func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Int64Primitive) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(m.Int64Primitive)))
		i += copy(dAtA[i:], m.Int64Primitive)
	}
	if len(m.BoolPrimitive) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(m.BoolPrimitive)))
		i += copy(dAtA[i:], m.BoolPrimitive)
	}
	if len(m.DoublePrimitive) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(m.DoublePrimitive)))
		i += copy(dAtA[i:], m.DoublePrimitive)
	}
	if len(m.StringPrimitive) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(m.StringPrimitive)))
		i += copy(dAtA[i:], m.StringPrimitive)
	}
	if len(m.AttributeBindings) > 0 {
		for k, _ := range m.AttributeBindings {
			dAtA[i] = 0x82
			i++
			dAtA[i] = 0xd3
			i++
			dAtA[i] = 0xe4
			i++
			dAtA[i] = 0x93
			i++
			dAtA[i] = 0x2
			i++
			v := m.AttributeBindings[k]
			mapSize := 1 + len(k) + sovTmplHandlerService(uint64(len(k))) + 1 + len(v) + sovTmplHandlerService(uint64(len(v)))
			i = encodeVarintTmplHandlerService(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintTmplHandlerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	l = len(m.Int64Primitive)
	if l > 0 {
		n += 1 + l + sovTmplHandlerService(uint64(l))
	}
	l = len(m.BoolPrimitive)
	if l > 0 {
		n += 1 + l + sovTmplHandlerService(uint64(l))
	}
	l = len(m.DoublePrimitive)
	if l > 0 {
		n += 1 + l + sovTmplHandlerService(uint64(l))
	}
	l = len(m.StringPrimitive)
	if l > 0 {
		n += 1 + l + sovTmplHandlerService(uint64(l))
	}
	if len(m.AttributeBindings) > 0 {
		for k, v := range m.AttributeBindings {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTmplHandlerService(uint64(len(k))) + 1 + len(v) + sovTmplHandlerService(uint64(len(v)))
			n += mapEntrySize + 5 + sovTmplHandlerService(uint64(mapEntrySize))
		}
	}
	return n
}

func sovTmplHandlerService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTmplHandlerService(x uint64) (n int) {
	return sovTmplHandlerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	keysForAttributeBindings := make([]string, 0, len(this.AttributeBindings))
	for k, _ := range this.AttributeBindings {
		keysForAttributeBindings = append(keysForAttributeBindings, k)
	}
	sortkeys.Strings(keysForAttributeBindings)
	mapStringForAttributeBindings := "map[string]string{"
	for _, k := range keysForAttributeBindings {
		mapStringForAttributeBindings += fmt.Sprintf("%v: %v,", k, this.AttributeBindings[k])
	}
	mapStringForAttributeBindings += "}"
	s := strings.Join([]string{`&InstanceParam{`,
		`Int64Primitive:` + fmt.Sprintf("%v", this.Int64Primitive) + `,`,
		`BoolPrimitive:` + fmt.Sprintf("%v", this.BoolPrimitive) + `,`,
		`DoublePrimitive:` + fmt.Sprintf("%v", this.DoublePrimitive) + `,`,
		`StringPrimitive:` + fmt.Sprintf("%v", this.StringPrimitive) + `,`,
		`AttributeBindings:` + mapStringForAttributeBindings + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTmplHandlerService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTmplHandlerService
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
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64Primitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
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
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Int64Primitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolPrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
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
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BoolPrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DoublePrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
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
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DoublePrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringPrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
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
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringPrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 72295728:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttributeBindings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
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
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttributeBindings == nil {
				m.AttributeBindings = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTmplHandlerService
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTmplHandlerService
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTmplHandlerService
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTmplHandlerService
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTmplHandlerService
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTmplHandlerService(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTmplHandlerService
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.AttributeBindings[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTmplHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTmplHandlerService
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
func skipTmplHandlerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTmplHandlerService
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
					return 0, ErrIntOverflowTmplHandlerService
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
					return 0, ErrIntOverflowTmplHandlerService
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
				return 0, ErrInvalidLengthTmplHandlerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTmplHandlerService
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
				next, err := skipTmplHandlerService(dAtA[start:])
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
	ErrInvalidLengthTmplHandlerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTmplHandlerService   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/test/spyAdapter/template/apa/tmpl_handler_service.proto", fileDescriptorTmplHandlerService)
}

var fileDescriptorTmplHandlerService = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6a, 0x14, 0x41,
	0x10, 0x87, 0xa7, 0x67, 0x55, 0x48, 0x4b, 0xfc, 0xd3, 0x04, 0x09, 0x7b, 0x68, 0x42, 0x10, 0xd9,
	0x83, 0x6c, 0x13, 0x15, 0x11, 0xc1, 0x43, 0x82, 0x39, 0x78, 0x0b, 0x79, 0x81, 0xa5, 0x26, 0x53,
	0x8c, 0x8d, 0x3d, 0xdd, 0x4d, 0x77, 0xed, 0x90, 0xbd, 0x89, 0x4f, 0x20, 0xe4, 0x25, 0x3c, 0xfa,
	0x00, 0x3e, 0x40, 0xf0, 0x14, 0x3c, 0x89, 0x20, 0x38, 0x63, 0x0e, 0x1e, 0x73, 0xf4, 0x28, 0x3b,
	0xb3, 0x71, 0xc8, 0x90, 0x5b, 0xd5, 0xc7, 0x57, 0x05, 0xf5, 0x2b, 0xfe, 0xaa, 0xd4, 0xc7, 0x18,
	0x14, 0x61, 0x24, 0x15, 0xfd, 0x62, 0x37, 0x07, 0x4f, 0x6d, 0x5f, 0x7a, 0x03, 0x84, 0x0a, 0x3c,
	0x28, 0x2a, 0xbd, 0x99, 0xbd, 0x05, 0x9b, 0x1b, 0x0c, 0xb3, 0x88, 0xa1, 0xd2, 0x47, 0x38, 0xf5,
	0xc1, 0x91, 0x13, 0x6b, 0x11, 0x4a, 0x6f, 0x10, 0x3c, 0x8c, 0x37, 0x0a, 0x57, 0xb8, 0x96, 0xaa,
	0x65, 0xd5, 0x09, 0xe3, 0xc7, 0xdd, 0x7e, 0x58, 0xed, 0x2d, 0x5d, 0x8e, 0x46, 0x55, 0x3b, 0x19,
	0x12, 0xec, 0x28, 0x3c, 0x26, 0xb4, 0x51, 0x3b, 0x1b, 0x3b, 0x7b, 0xfb, 0x47, 0xca, 0xd7, 0xdf,
	0xd8, 0x48, 0x60, 0x8f, 0xf0, 0x00, 0x02, 0x94, 0xe2, 0x11, 0xbf, 0xa3, 0x2d, 0x3d, 0x7f, 0x76,
	0x10, 0x74, 0xa9, 0x49, 0x57, 0xb8, 0xc9, 0xb6, 0xd8, 0x64, 0xed, 0x70, 0x40, 0xc5, 0x43, 0xbe,
	0x9e, 0x39, 0x67, 0x7a, 0x2d, 0x6d, 0xb5, 0xab, 0x50, 0x4c, 0xf8, 0xdd, 0xdc, 0xcd, 0x33, 0x83,
	0xbd, 0x37, 0x6a, 0xbd, 0x21, 0x5e, 0x9a, 0x91, 0x82, 0xb6, 0x45, 0x6f, 0xde, 0xe8, 0xcc, 0x01,
	0x16, 0xc0, 0x05, 0x10, 0x05, 0x9d, 0xcd, 0x09, 0x67, 0x99, 0xb6, 0xb9, 0xb6, 0x45, 0xdc, 0xfc,
	0xfc, 0xf5, 0xcb, 0xf6, 0xd6, 0x68, 0x72, 0xfb, 0x89, 0x9a, 0xfe, 0x8f, 0x68, 0x7a, 0xe5, 0xb4,
	0xe9, 0xee, 0xe5, 0xd4, 0xde, 0x6a, 0x68, 0xdf, 0x52, 0x58, 0x1c, 0xde, 0x87, 0x21, 0x1f, 0xbf,
	0xe6, 0x0f, 0xae, 0x97, 0xc5, 0x3d, 0x3e, 0x7a, 0x87, 0x8b, 0x55, 0x26, 0xcb, 0x52, 0x6c, 0xf0,
	0x9b, 0x15, 0x98, 0xf9, 0x65, 0x00, 0x5d, 0xf3, 0x32, 0x7d, 0xc1, 0xf6, 0xf6, 0x4f, 0x6b, 0x99,
	0x9c, 0xd5, 0x32, 0xf9, 0x5e, 0xcb, 0xe4, 0xa2, 0x96, 0xc9, 0xfb, 0x46, 0xb2, 0x4f, 0x8d, 0x4c,
	0x4e, 0x1b, 0xc9, 0xce, 0x1a, 0xc9, 0x7e, 0x35, 0x92, 0xfd, 0x69, 0x64, 0x72, 0xd1, 0x48, 0xf6,
	0xf1, 0xb7, 0x4c, 0xfe, 0x7e, 0x3b, 0x3f, 0x49, 0x47, 0x1f, 0x7e, 0x9e, 0x9f, 0xa4, 0xfd, 0x9f,
	0xb3, 0x5b, 0xed, 0xab, 0x9e, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x18, 0x20, 0xa8, 0xa0, 0x3a,
	0x02, 0x00, 0x00,
}