// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_extension_setting.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An ad group extension setting.
type AdGroupExtensionSetting struct {
	// The resource name of the ad group extension setting.
	// AdGroupExtensionSetting resource names have the form:
	//
	//
	// `customers/{customer_id}/adGroupExtensionSettings/{ad_group_id}~{extension_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The extension type of the ad group extension setting.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v1.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// The resource name of the ad group. The linked extension feed items will
	// serve under this ad group.
	// AdGroup resource names have the form:
	//
	// `customers/{customer_id}/adGroups/{ad_group_id}`
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The resource names of the extension feed items to serve under the ad group.
	// ExtensionFeedItem resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ExtensionFeedItems []*wrappers.StringValue `protobuf:"bytes,4,rep,name=extension_feed_items,json=extensionFeedItems,proto3" json:"extension_feed_items,omitempty"`
	// The device for which the extensions will serve. Optional.
	Device               enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice `protobuf:"varint,5,opt,name=device,proto3,enum=google.ads.googleads.v1.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *AdGroupExtensionSetting) Reset()         { *m = AdGroupExtensionSetting{} }
func (m *AdGroupExtensionSetting) String() string { return proto.CompactTextString(m) }
func (*AdGroupExtensionSetting) ProtoMessage()    {}
func (*AdGroupExtensionSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_extension_setting_c6bc821b4e4f8ffc, []int{0}
}
func (m *AdGroupExtensionSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupExtensionSetting.Unmarshal(m, b)
}
func (m *AdGroupExtensionSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupExtensionSetting.Marshal(b, m, deterministic)
}
func (dst *AdGroupExtensionSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupExtensionSetting.Merge(dst, src)
}
func (m *AdGroupExtensionSetting) XXX_Size() int {
	return xxx_messageInfo_AdGroupExtensionSetting.Size(m)
}
func (m *AdGroupExtensionSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupExtensionSetting.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupExtensionSetting proto.InternalMessageInfo

func (m *AdGroupExtensionSetting) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupExtensionSetting) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *AdGroupExtensionSetting) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupExtensionSetting) GetExtensionFeedItems() []*wrappers.StringValue {
	if m != nil {
		return m.ExtensionFeedItems
	}
	return nil
}

func (m *AdGroupExtensionSetting) GetDevice() enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice {
	if m != nil {
		return m.Device
	}
	return enums.ExtensionSettingDeviceEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupExtensionSetting)(nil), "google.ads.googleads.v1.resources.AdGroupExtensionSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_extension_setting.proto", fileDescriptor_ad_group_extension_setting_c6bc821b4e4f8ffc)
}

var fileDescriptor_ad_group_extension_setting_c6bc821b4e4f8ffc = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0xa5, 0x1d, 0x5d, 0xb5, 0xba, 0xfb, 0x50, 0x04, 0xcb, 0x32, 0xc8, 0xac, 0xb2, 0x30, 0x4f,
	0x29, 0x1d, 0x1f, 0x84, 0x2a, 0x42, 0x07, 0xd7, 0x41, 0x1f, 0x96, 0xa1, 0x2b, 0xf3, 0x20, 0x03,
	0x25, 0x3b, 0xb9, 0x1b, 0x03, 0xd3, 0x24, 0x24, 0xe9, 0xe8, 0x7e, 0x82, 0x7f, 0xe0, 0xb3, 0x8f,
	0x7e, 0x8a, 0x9f, 0xe2, 0x57, 0x48, 0x9b, 0x26, 0x3a, 0xc8, 0xec, 0xce, 0xdb, 0x69, 0xee, 0x39,
	0xe7, 0xde, 0x93, 0x9b, 0x46, 0x53, 0x2a, 0x04, 0x5d, 0x43, 0x8a, 0x89, 0x4e, 0x2d, 0x6c, 0xd1,
	0x26, 0x4b, 0x15, 0x68, 0xd1, 0xa8, 0x15, 0xe8, 0x14, 0x93, 0x8a, 0x2a, 0xd1, 0xc8, 0x0a, 0xbe,
	0x1a, 0xe0, 0x9a, 0x09, 0x5e, 0x69, 0x30, 0x86, 0x71, 0x8a, 0xa4, 0x12, 0x46, 0xc4, 0x27, 0x56,
	0x88, 0x30, 0xd1, 0xc8, 0x7b, 0xa0, 0x4d, 0x86, 0xbc, 0xc7, 0xf1, 0xeb, 0x5d, 0x6d, 0x80, 0x37,
	0xb5, 0x4e, 0xff, 0x73, 0xae, 0x08, 0x6c, 0xd8, 0x0a, 0x6c, 0x83, 0xe3, 0xc9, 0xbe, 0x6a, 0x73,
	0x2d, 0x9d, 0xe6, 0x69, 0xaf, 0xe9, 0xbe, 0x2e, 0x9b, 0xab, 0xf4, 0x8b, 0xc2, 0x52, 0x82, 0xd2,
	0x7d, 0x7d, 0xe8, 0x3c, 0x25, 0x4b, 0x31, 0xe7, 0xc2, 0x60, 0xc3, 0x04, 0xef, 0xab, 0xcf, 0xbe,
	0x0f, 0xa2, 0x27, 0x05, 0x99, 0xb5, 0xb1, 0xcf, 0x9c, 0xfb, 0x85, 0x1d, 0x2d, 0x7e, 0x1e, 0x1d,
	0xba, 0x60, 0x15, 0xc7, 0x35, 0x24, 0xc1, 0x28, 0x18, 0x3f, 0x28, 0x1f, 0xb9, 0xc3, 0x73, 0x5c,
	0x43, 0x0c, 0xd1, 0xd1, 0xf6, 0x58, 0x49, 0x38, 0x0a, 0xc6, 0x47, 0x93, 0x37, 0x68, 0xd7, 0x65,
	0x75, 0x59, 0x90, 0xef, 0xf6, 0xf1, 0x5a, 0xc2, 0x19, 0x6f, 0xea, 0xed, 0x93, 0xf2, 0x10, 0xfe,
	0xfd, 0x8c, 0x5f, 0x46, 0xf7, 0xdd, 0x7a, 0x92, 0xc1, 0x28, 0x18, 0x3f, 0x9c, 0x0c, 0x5d, 0x03,
	0x17, 0x1c, 0x5d, 0x18, 0xc5, 0x38, 0x5d, 0xe0, 0x75, 0x03, 0xe5, 0x3d, 0x6c, 0x43, 0xc5, 0xe7,
	0xd1, 0xe3, 0xbf, 0xf3, 0x5d, 0x01, 0x90, 0x8a, 0x19, 0xa8, 0x75, 0x72, 0x67, 0x34, 0xb8, 0xd5,
	0x24, 0xf6, 0xca, 0x77, 0x00, 0xe4, 0x7d, 0xab, 0x8b, 0x3f, 0x47, 0x07, 0x76, 0x65, 0xc9, 0xdd,
	0x2e, 0xe7, 0x7c, 0xdf, 0x9c, 0xfd, 0xad, 0xbe, 0xed, 0xc4, 0xdb, 0x81, 0xb7, 0x4a, 0x65, 0xef,
	0x3f, 0xfd, 0x16, 0x46, 0xa7, 0x2b, 0x51, 0xa3, 0x5b, 0x1f, 0xdd, 0x74, 0xb8, 0x63, 0x83, 0xf3,
	0x36, 0xd4, 0x3c, 0xf8, 0xf4, 0xa1, 0xb7, 0xa0, 0x62, 0x8d, 0x39, 0x45, 0x42, 0xd1, 0x94, 0x02,
	0xef, 0x22, 0xbb, 0x67, 0x26, 0x99, 0xbe, 0xe1, 0xd7, 0x78, 0xe5, 0xd1, 0x8f, 0x70, 0x30, 0x2b,
	0x8a, 0x9f, 0xe1, 0xc9, 0xcc, 0x5a, 0x16, 0x44, 0x23, 0x0b, 0x5b, 0xb4, 0xc8, 0x50, 0xe9, 0x98,
	0xbf, 0x1c, 0x67, 0x59, 0x10, 0xbd, 0xf4, 0x9c, 0xe5, 0x22, 0x5b, 0x7a, 0xce, 0xef, 0xf0, 0xd4,
	0x16, 0xf2, 0xbc, 0x20, 0x3a, 0xcf, 0x3d, 0x2b, 0xcf, 0x17, 0x59, 0x9e, 0x7b, 0xde, 0xe5, 0x41,
	0x37, 0xec, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xcb, 0x94, 0xc7, 0xc6, 0x03, 0x00,
	0x00,
}
