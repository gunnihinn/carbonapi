// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: carbonapi.proto

/*
Package carbonapipb is a generated protocol buffer package.

It is generated from these files:
	carbonapi.proto

It has these top-level messages:
	AccessLogDetails
*/
package carbonapipb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AccessLogDetails struct {
	Handler                       string   `protobuf:"bytes,1,opt,name=handler,proto3" json:"handler,omitempty"`
	CarbonapiUuid                 string   `protobuf:"bytes,2,opt,name=carbonapi_uuid,json=carbonapiUuid,proto3" json:"carbonapi_uuid,omitempty"`
	Username                      string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Url                           string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	PeerIp                        string   `protobuf:"bytes,5,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	PeerPort                      string   `protobuf:"bytes,6,opt,name=peer_port,json=peerPort,proto3" json:"peer_port,omitempty"`
	Host                          string   `protobuf:"bytes,7,opt,name=host,proto3" json:"host,omitempty"`
	Referer                       string   `protobuf:"bytes,8,opt,name=referer,proto3" json:"referer,omitempty"`
	Format                        string   `protobuf:"bytes,9,opt,name=format,proto3" json:"format,omitempty"`
	UseCache                      bool     `protobuf:"varint,10,opt,name=use_cache,json=useCache,proto3" json:"use_cache,omitempty"`
	Targets                       []string `protobuf:"bytes,11,rep,name=targets" json:"targets,omitempty"`
	CacheTimeout                  int32    `protobuf:"varint,12,opt,name=cache_timeout,json=cacheTimeout,proto3" json:"cache_timeout,omitempty"`
	Metrics                       []string `protobuf:"bytes,13,rep,name=metrics" json:"metrics,omitempty"`
	HaveNonFatalErrors            bool     `protobuf:"varint,14,opt,name=have_non_fatal_errors,json=haveNonFatalErrors,proto3" json:"have_non_fatal_errors,omitempty"`
	Runtime                       float64  `protobuf:"fixed64,15,opt,name=runtime,proto3" json:"runtime,omitempty"`
	HttpCode                      int32    `protobuf:"varint,16,opt,name=http_code,json=httpCode,proto3" json:"http_code,omitempty"`
	CarbonzipperResponseSizeBytes int64    `protobuf:"varint,17,opt,name=carbonzipper_response_size_bytes,json=carbonzipperResponseSizeBytes,proto3" json:"carbonzipper_response_size_bytes,omitempty"`
	CarbonapiResponseSizeBytes    int64    `protobuf:"varint,18,opt,name=carbonapi_response_size_bytes,json=carbonapiResponseSizeBytes,proto3" json:"carbonapi_response_size_bytes,omitempty"`
	Reason                        string   `protobuf:"bytes,19,opt,name=reason,proto3" json:"reason,omitempty"`
	SendGlobs                     bool     `protobuf:"varint,20,opt,name=send_globs,json=sendGlobs,proto3" json:"send_globs,omitempty"`
	From                          int32    `protobuf:"varint,21,opt,name=from,proto3" json:"from,omitempty"`
	Until                         int32    `protobuf:"varint,22,opt,name=until,proto3" json:"until,omitempty"`
	Tz                            string   `protobuf:"bytes,23,opt,name=tz,proto3" json:"tz,omitempty"`
	FromRaw                       string   `protobuf:"bytes,24,opt,name=from_raw,json=fromRaw,proto3" json:"from_raw,omitempty"`
	UntilRaw                      string   `protobuf:"bytes,25,opt,name=until_raw,json=untilRaw,proto3" json:"until_raw,omitempty"`
	Uri                           string   `protobuf:"bytes,26,opt,name=uri,proto3" json:"uri,omitempty"`
	FromCache                     bool     `protobuf:"varint,27,opt,name=from_cache,json=fromCache,proto3" json:"from_cache,omitempty"`
}

func (m *AccessLogDetails) Reset()                    { *m = AccessLogDetails{} }
func (m *AccessLogDetails) String() string            { return proto.CompactTextString(m) }
func (*AccessLogDetails) ProtoMessage()               {}
func (*AccessLogDetails) Descriptor() ([]byte, []int) { return fileDescriptorCarbonapi, []int{0} }

func (m *AccessLogDetails) GetHandler() string {
	if m != nil {
		return m.Handler
	}
	return ""
}

func (m *AccessLogDetails) GetCarbonapiUuid() string {
	if m != nil {
		return m.CarbonapiUuid
	}
	return ""
}

func (m *AccessLogDetails) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccessLogDetails) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AccessLogDetails) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

func (m *AccessLogDetails) GetPeerPort() string {
	if m != nil {
		return m.PeerPort
	}
	return ""
}

func (m *AccessLogDetails) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AccessLogDetails) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

func (m *AccessLogDetails) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *AccessLogDetails) GetUseCache() bool {
	if m != nil {
		return m.UseCache
	}
	return false
}

func (m *AccessLogDetails) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *AccessLogDetails) GetCacheTimeout() int32 {
	if m != nil {
		return m.CacheTimeout
	}
	return 0
}

func (m *AccessLogDetails) GetMetrics() []string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *AccessLogDetails) GetHaveNonFatalErrors() bool {
	if m != nil {
		return m.HaveNonFatalErrors
	}
	return false
}

func (m *AccessLogDetails) GetRuntime() float64 {
	if m != nil {
		return m.Runtime
	}
	return 0
}

func (m *AccessLogDetails) GetHttpCode() int32 {
	if m != nil {
		return m.HttpCode
	}
	return 0
}

func (m *AccessLogDetails) GetCarbonzipperResponseSizeBytes() int64 {
	if m != nil {
		return m.CarbonzipperResponseSizeBytes
	}
	return 0
}

func (m *AccessLogDetails) GetCarbonapiResponseSizeBytes() int64 {
	if m != nil {
		return m.CarbonapiResponseSizeBytes
	}
	return 0
}

func (m *AccessLogDetails) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *AccessLogDetails) GetSendGlobs() bool {
	if m != nil {
		return m.SendGlobs
	}
	return false
}

func (m *AccessLogDetails) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *AccessLogDetails) GetUntil() int32 {
	if m != nil {
		return m.Until
	}
	return 0
}

func (m *AccessLogDetails) GetTz() string {
	if m != nil {
		return m.Tz
	}
	return ""
}

func (m *AccessLogDetails) GetFromRaw() string {
	if m != nil {
		return m.FromRaw
	}
	return ""
}

func (m *AccessLogDetails) GetUntilRaw() string {
	if m != nil {
		return m.UntilRaw
	}
	return ""
}

func (m *AccessLogDetails) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *AccessLogDetails) GetFromCache() bool {
	if m != nil {
		return m.FromCache
	}
	return false
}

func init() {
	proto.RegisterType((*AccessLogDetails)(nil), "carbonapipb.AccessLogDetails")
}

func init() { proto.RegisterFile("carbonapi.proto", fileDescriptorCarbonapi) }

var fileDescriptorCarbonapi = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x61, 0x6b, 0x13, 0x4d,
	0x10, 0xc7, 0xb9, 0xa6, 0xbd, 0x24, 0xdb, 0x26, 0xcd, 0xb3, 0x4f, 0x9b, 0x4e, 0x53, 0x02, 0x87,
	0x22, 0xe4, 0x95, 0x20, 0x7e, 0x82, 0x5a, 0xb5, 0x08, 0x22, 0x72, 0xea, 0xeb, 0x65, 0x73, 0x37,
	0x49, 0x16, 0x2e, 0xb7, 0xcb, 0xee, 0x9e, 0xc5, 0x7c, 0x3d, 0xbf, 0x98, 0xcc, 0x6c, 0x12, 0x05,
	0x7d, 0xb7, 0xff, 0xdf, 0xcc, 0xfc, 0x77, 0xf6, 0x66, 0x4e, 0x5c, 0x56, 0xda, 0x2f, 0x6d, 0xab,
	0x9d, 0x79, 0xe9, 0xbc, 0x8d, 0x56, 0x9e, 0x1f, 0x81, 0x5b, 0x3e, 0xfb, 0x99, 0x8b, 0xc9, 0x7d,
	0x55, 0x61, 0x08, 0x1f, 0xed, 0xfa, 0x2d, 0x46, 0x6d, 0x9a, 0x20, 0x41, 0xf4, 0x37, 0xba, 0xad,
	0x1b, 0xf4, 0x90, 0x15, 0xd9, 0x62, 0x58, 0x1e, 0xa4, 0x7c, 0x21, 0xc6, 0xc7, 0x6a, 0xd5, 0x75,
	0xa6, 0x86, 0x13, 0x4e, 0x18, 0x1d, 0xe9, 0xb7, 0xce, 0xd4, 0x72, 0x26, 0x06, 0x5d, 0x40, 0xdf,
	0xea, 0x2d, 0x42, 0x8f, 0x13, 0x8e, 0x5a, 0x4e, 0x44, 0xaf, 0xf3, 0x0d, 0x9c, 0x32, 0xa6, 0xa3,
	0xbc, 0x11, 0x7d, 0x87, 0xe8, 0x95, 0x71, 0x70, 0xc6, 0x34, 0x27, 0xf9, 0xc1, 0xc9, 0x3b, 0x31,
	0xe4, 0x80, 0xb3, 0x3e, 0x42, 0x9e, 0x7c, 0x08, 0x7c, 0xb6, 0x3e, 0x4a, 0x29, 0x4e, 0x37, 0x36,
	0x44, 0xe8, 0x33, 0xe7, 0x33, 0x35, 0xee, 0x71, 0x85, 0x1e, 0x3d, 0x0c, 0x52, 0xe3, 0x7b, 0x29,
	0xa7, 0x22, 0x5f, 0x59, 0xbf, 0xd5, 0x11, 0x86, 0xe9, 0x8a, 0xa4, 0xe8, 0x8a, 0x2e, 0xa0, 0xaa,
	0x74, 0xb5, 0x41, 0x10, 0x45, 0xb6, 0x18, 0x70, 0xab, 0x0f, 0xa4, 0xc9, 0x2e, 0x6a, 0xbf, 0xc6,
	0x18, 0xe0, 0xbc, 0xe8, 0x91, 0xdd, 0x5e, 0xca, 0xe7, 0x62, 0xc4, 0x25, 0x2a, 0x9a, 0x2d, 0xda,
	0x2e, 0xc2, 0x45, 0x91, 0x2d, 0xce, 0xca, 0x0b, 0x86, 0x5f, 0x13, 0xa3, 0xf2, 0x2d, 0x46, 0x6f,
	0xaa, 0x00, 0xa3, 0x54, 0xbe, 0x97, 0xf2, 0x95, 0xb8, 0xde, 0xe8, 0xef, 0xa8, 0x5a, 0xdb, 0xaa,
	0x95, 0x8e, 0xba, 0x51, 0xe8, 0xbd, 0xf5, 0x01, 0xc6, 0xdc, 0x81, 0xa4, 0xe0, 0x27, 0xdb, 0xbe,
	0xa7, 0xd0, 0x3b, 0x8e, 0xf0, 0xd3, 0xba, 0x96, 0xae, 0x83, 0xcb, 0x22, 0x5b, 0x64, 0xe5, 0x41,
	0xd2, 0x13, 0x36, 0x31, 0x3a, 0x55, 0xd9, 0x1a, 0x61, 0xc2, 0x7d, 0x0c, 0x08, 0x3c, 0xd8, 0x1a,
	0xe5, 0xa3, 0x28, 0xd2, 0x68, 0x76, 0xc6, 0x39, 0xf4, 0xca, 0x63, 0x70, 0xb6, 0x0d, 0xa8, 0x82,
	0xd9, 0xa1, 0x5a, 0xfe, 0x88, 0x18, 0xe0, 0xbf, 0x22, 0x5b, 0xf4, 0xca, 0xf9, 0x9f, 0x79, 0xe5,
	0x3e, 0xed, 0x8b, 0xd9, 0xe1, 0x1b, 0x4a, 0x92, 0xf7, 0x62, 0xfe, 0x7b, 0xf2, 0xff, 0x72, 0x91,
	0xec, 0x32, 0x3b, 0x26, 0xfd, 0x6d, 0x31, 0x15, 0xb9, 0x47, 0x1d, 0x6c, 0x0b, 0xff, 0xa7, 0x19,
	0x24, 0x25, 0xe7, 0x42, 0x04, 0x6c, 0x6b, 0xb5, 0x6e, 0xec, 0x32, 0xc0, 0x15, 0x7f, 0x82, 0x21,
	0x91, 0x47, 0x02, 0x34, 0xe8, 0x95, 0xb7, 0x5b, 0xb8, 0xe6, 0xa7, 0xf1, 0x59, 0x5e, 0x89, 0x33,
	0x7a, 0x7d, 0x03, 0x53, 0x86, 0x49, 0xc8, 0xb1, 0x38, 0x89, 0x3b, 0xb8, 0x61, 0xf3, 0x93, 0xb8,
	0x93, 0xb7, 0x62, 0x40, 0xd9, 0xca, 0xeb, 0x27, 0x80, 0xb4, 0x0f, 0xa4, 0x4b, 0xfd, 0xc4, 0x73,
	0xa7, 0x1a, 0x8e, 0xdd, 0xee, 0x57, 0x94, 0x00, 0x05, 0x79, 0x45, 0x0d, 0xcc, 0x0e, 0x2b, 0x6a,
	0xa8, 0x45, 0x76, 0x4a, 0x7b, 0x72, 0x97, 0x5a, 0x24, 0xc2, 0x8b, 0xb2, 0xcc, 0xf9, 0xcf, 0x7a,
	0xfd, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x41, 0x99, 0x55, 0x6c, 0x03, 0x00, 0x00,
}