// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/api/storage/v1/generated.proto

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/api/storage/v1/generated.proto

	It has these top-level messages:
		StorageClass
		StorageClassList
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gopro/ext-k8s/apis/storage/v1beta1"
import k8s_io_apimachinery_pkg_apis_meta_v1 "github.com/gopro/ext-k8s/apis/meta/v1"
import _ "github.com/gopro/ext-k8s/runtime"
import _ "github.com/gopro/ext-k8s/runtime/schema"
import _ "github.com/gopro/ext-k8s/util/intstr"

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

// StorageClass describes the parameters for a class of storage for
// which PersistentVolumes can be dynamically provisioned.
//
// StorageClasses are non-namespaced; the name of the storage class
// according to etcd is in ObjectMeta.Name.
type StorageClass struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Provisioner indicates the type of the provisioner.
	Provisioner *string `protobuf:"bytes,2,opt,name=provisioner" json:"provisioner,omitempty"`
	// Parameters holds the parameters for the provisioner that should
	// create volumes of this storage class.
	// +optional
	Parameters map[string]string `protobuf:"bytes,3,rep,name=parameters" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Dynamically provisioned PersistentVolumes of this storage class are
	// created with this reclaimPolicy. Defaults to Delete.
	// +optional
	ReclaimPolicy *string `protobuf:"bytes,4,opt,name=reclaimPolicy" json:"reclaimPolicy,omitempty"`
	// Dynamically provisioned PersistentVolumes of this storage class are
	// created with these mountOptions, e.g. ["ro", "soft"]. Not validated -
	// mount of the PVs will simply fail if one is invalid.
	// +optional
	MountOptions []string `protobuf:"bytes,5,rep,name=mountOptions" json:"mountOptions,omitempty"`
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	// +optional
	AllowVolumeExpansion *bool `protobuf:"varint,6,opt,name=allowVolumeExpansion" json:"allowVolumeExpansion,omitempty"`
	// VolumeBindingMode indicates how PersistentVolumeClaims should be
	// provisioned and bound.  When unset, VolumeBindingImmediate is used.
	// This field is alpha-level and is only honored by servers that enable
	// the VolumeScheduling feature.
	// +optional
	VolumeBindingMode *string `protobuf:"bytes,7,opt,name=volumeBindingMode" json:"volumeBindingMode,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *StorageClass) Reset()                    { *m = StorageClass{} }
func (m *StorageClass) String() string            { return proto.CompactTextString(m) }
func (*StorageClass) ProtoMessage()               {}
func (*StorageClass) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *StorageClass) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StorageClass) GetProvisioner() string {
	if m != nil && m.Provisioner != nil {
		return *m.Provisioner
	}
	return ""
}

func (m *StorageClass) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *StorageClass) GetReclaimPolicy() string {
	if m != nil && m.ReclaimPolicy != nil {
		return *m.ReclaimPolicy
	}
	return ""
}

func (m *StorageClass) GetMountOptions() []string {
	if m != nil {
		return m.MountOptions
	}
	return nil
}

func (m *StorageClass) GetAllowVolumeExpansion() bool {
	if m != nil && m.AllowVolumeExpansion != nil {
		return *m.AllowVolumeExpansion
	}
	return false
}

func (m *StorageClass) GetVolumeBindingMode() string {
	if m != nil && m.VolumeBindingMode != nil {
		return *m.VolumeBindingMode
	}
	return ""
}

// StorageClassList is a collection of storage classes.
type StorageClassList struct {
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Items is the list of StorageClasses
	Items            []*StorageClass `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *StorageClassList) Reset()                    { *m = StorageClassList{} }
func (m *StorageClassList) String() string            { return proto.CompactTextString(m) }
func (*StorageClassList) ProtoMessage()               {}
func (*StorageClassList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *StorageClassList) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StorageClassList) GetItems() []*StorageClass {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*StorageClass)(nil), "k8s.io.api.storage.v1.StorageClass")
	proto.RegisterType((*StorageClassList)(nil), "k8s.io.api.storage.v1.StorageClassList")
}
func (m *StorageClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageClass) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Provisioner != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Provisioner)))
		i += copy(dAtA[i:], *m.Provisioner)
	}
	if len(m.Parameters) > 0 {
		for k, _ := range m.Parameters {
			dAtA[i] = 0x1a
			i++
			v := m.Parameters[k]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.ReclaimPolicy != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.ReclaimPolicy)))
		i += copy(dAtA[i:], *m.ReclaimPolicy)
	}
	if len(m.MountOptions) > 0 {
		for _, s := range m.MountOptions {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.AllowVolumeExpansion != nil {
		dAtA[i] = 0x30
		i++
		if *m.AllowVolumeExpansion {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.VolumeBindingMode != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.VolumeBindingMode)))
		i += copy(dAtA[i:], *m.VolumeBindingMode)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StorageClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageClassList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n2, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
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

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StorageClass) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Provisioner != nil {
		l = len(*m.Provisioner)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Parameters) > 0 {
		for k, v := range m.Parameters {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if m.ReclaimPolicy != nil {
		l = len(*m.ReclaimPolicy)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.MountOptions) > 0 {
		for _, s := range m.MountOptions {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.AllowVolumeExpansion != nil {
		n += 2
	}
	if m.VolumeBindingMode != nil {
		l = len(*m.VolumeBindingMode)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StorageClassList) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StorageClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: StorageClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provisioner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Provisioner = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Parameters == nil {
				m.Parameters = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
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
							return ErrIntOverflowGenerated
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
						return ErrInvalidLengthGenerated
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
							return ErrIntOverflowGenerated
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
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Parameters[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReclaimPolicy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ReclaimPolicy = &s
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MountOptions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MountOptions = append(m.MountOptions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowVolumeExpansion", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			b := bool(v != 0)
			m.AllowVolumeExpansion = &b
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeBindingMode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.VolumeBindingMode = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *StorageClassList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: StorageClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &StorageClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("k8s.io/api/storage/v1/generated.proto", fileDescriptorGenerated) }

var fileDescriptorGenerated = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x4c, 0xa0, 0xd9, 0x14, 0x51, 0x56, 0x41, 0xb2, 0x72, 0x88, 0xac, 0x00, 0x92,
	0x85, 0xd0, 0x9a, 0xa4, 0x08, 0x15, 0x24, 0x2e, 0x45, 0xbd, 0xa0, 0x56, 0xad, 0x5c, 0x89, 0x03,
	0xb7, 0xa9, 0x33, 0x72, 0x97, 0xd8, 0xbb, 0xd6, 0xee, 0xd8, 0x90, 0x37, 0x81, 0x97, 0xe1, 0xcc,
	0x91, 0x47, 0x40, 0xe1, 0x45, 0x90, 0x9d, 0x2a, 0x4d, 0xe3, 0x54, 0xcd, 0xcd, 0xfe, 0x76, 0x7e,
	0xf3, 0xe7, 0xfb, 0xd8, 0x8b, 0xe9, 0x81, 0x15, 0x52, 0x87, 0x90, 0xcb, 0xd0, 0x92, 0x36, 0x90,
	0x60, 0x58, 0x8e, 0xc2, 0x04, 0x15, 0x1a, 0x20, 0x9c, 0x88, 0xdc, 0x68, 0xd2, 0xfc, 0xe9, 0xa2,
	0x4c, 0x40, 0x2e, 0xc5, 0x55, 0x99, 0x28, 0x47, 0xfd, 0x97, 0x1b, 0xe9, 0x0b, 0x24, 0x68, 0xb4,
	0xe8, 0xbf, 0xb9, 0xae, 0xcd, 0x20, 0xbe, 0x94, 0x0a, 0xcd, 0x2c, 0xcc, 0xa7, 0x49, 0x25, 0xd8,
	0x30, 0x43, 0x82, 0x0d, 0x83, 0xfb, 0xe1, 0x6d, 0x94, 0x29, 0x14, 0xc9, 0x0c, 0x1b, 0xc0, 0xdb,
	0xbb, 0x00, 0x1b, 0x5f, 0x62, 0x06, 0x0d, 0x6e, 0xff, 0x36, 0xae, 0x20, 0x99, 0x86, 0x52, 0x91,
	0x25, 0xb3, 0x0e, 0x0d, 0x7f, 0xb9, 0x6c, 0xf7, 0x7c, 0x71, 0xf7, 0xc7, 0x14, 0xac, 0xe5, 0xc7,
	0x6c, 0xa7, 0xba, 0x64, 0x02, 0x04, 0x9e, 0xe3, 0x3b, 0x41, 0x77, 0xfc, 0x5a, 0x5c, 0x5b, 0xb7,
	0x6c, 0x2c, 0xf2, 0x69, 0x52, 0x09, 0x56, 0x54, 0xd5, 0xa2, 0x1c, 0x89, 0xd3, 0x8b, 0xaf, 0x18,
	0xd3, 0x09, 0x12, 0x44, 0xcb, 0x0e, 0xdc, 0x67, 0xdd, 0xdc, 0xe8, 0x52, 0x5a, 0xa9, 0x15, 0x1a,
	0xaf, 0xe5, 0x3b, 0x41, 0x27, 0x5a, 0x95, 0xf8, 0x39, 0x63, 0x39, 0x18, 0xc8, 0x90, 0xd0, 0x58,
	0xcf, 0xf5, 0xdd, 0xa0, 0x3b, 0xde, 0x17, 0x1b, 0xc3, 0x12, 0xab, 0x8b, 0x8a, 0xb3, 0x25, 0x75,
	0xa4, 0xc8, 0xcc, 0xa2, 0x95, 0x36, 0xfc, 0x39, 0x7b, 0x64, 0x30, 0x4e, 0x41, 0x66, 0x67, 0x3a,
	0x95, 0xf1, 0xcc, 0xbb, 0x5f, 0x0f, 0xbe, 0x29, 0xf2, 0x21, 0xdb, 0xcd, 0x74, 0xa1, 0xe8, 0x34,
	0x27, 0xa9, 0x95, 0xf5, 0xda, 0xbe, 0x1b, 0x74, 0xa2, 0x1b, 0x1a, 0x1f, 0xb3, 0x1e, 0xa4, 0xa9,
	0xfe, 0xf6, 0x59, 0xa7, 0x45, 0x86, 0x47, 0xdf, 0x73, 0x50, 0xd5, 0xe2, 0xde, 0x03, 0xdf, 0x09,
	0x76, 0xa2, 0x8d, 0x6f, 0xfc, 0x15, 0x7b, 0x52, 0xd6, 0xd2, 0xa1, 0x54, 0x13, 0xa9, 0x92, 0x13,
	0x3d, 0x41, 0xef, 0x61, 0xbd, 0x41, 0xf3, 0xa1, 0xff, 0x81, 0x3d, 0x5e, 0x3b, 0x85, 0xef, 0x31,
	0x77, 0x8a, 0xb3, 0xda, 0xfe, 0x4e, 0x54, 0x7d, 0xf2, 0x1e, 0x6b, 0x97, 0x90, 0x16, 0x78, 0xe5,
	0xe0, 0xe2, 0xe7, 0x7d, 0xeb, 0xc0, 0x19, 0xfe, 0x74, 0xd8, 0xde, 0xaa, 0x2f, 0xc7, 0xd2, 0x12,
	0xff, 0xd4, 0x08, 0x51, 0x6c, 0x17, 0x62, 0x45, 0xaf, 0x45, 0xf8, 0x8e, 0xb5, 0x25, 0x61, 0x66,
	0xbd, 0x56, 0x9d, 0xcd, 0xb3, 0x2d, 0xb2, 0x89, 0x16, 0xc4, 0x61, 0xef, 0xf7, 0x7c, 0xe0, 0xfc,
	0x99, 0x0f, 0x9c, 0xbf, 0xf3, 0x81, 0xf3, 0xe3, 0xdf, 0xe0, 0xde, 0x97, 0x56, 0x39, 0xfa, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0xa2, 0xef, 0xf0, 0x1a, 0xb1, 0x03, 0x00, 0x00,
}
