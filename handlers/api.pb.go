// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package handlers is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	ApiPic
	ApiPicTag
	LookupPicDetailsResponse
	IndexResponse
	AddPicTagsResponse
	CreatePicResponse
	FindSimilarPicsResponse
*/
package handlers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ApiPic struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Width                int32                     `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height               int32                     `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Version              int64                     `protobuf:"fixed64,4,opt,name=version" json:"version,omitempty"`
	Type                 string                    `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	RelativeUrl          string                    `protobuf:"bytes,6,opt,name=relative_url,json=relativeUrl" json:"relative_url,omitempty"`
	ThumbnailRelativeUrl string                    `protobuf:"bytes,7,opt,name=thumbnail_relative_url,json=thumbnailRelativeUrl" json:"thumbnail_relative_url,omitempty"`
	PendingDeletion      bool                      `protobuf:"varint,9,opt,name=pending_deletion,json=pendingDeletion" json:"pending_deletion,omitempty"`
	ViewCount            int64                     `protobuf:"varint,10,opt,name=view_count,json=viewCount" json:"view_count,omitempty"`
	Duration             *google_protobuf.Duration `protobuf:"bytes,11,opt,name=duration" json:"duration,omitempty"`
}

func (m *ApiPic) Reset()                    { *m = ApiPic{} }
func (m *ApiPic) String() string            { return proto.CompactTextString(m) }
func (*ApiPic) ProtoMessage()               {}
func (*ApiPic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApiPic) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type ApiPicTag struct {
	PicId        string                      `protobuf:"bytes,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
	TagId        string                      `protobuf:"bytes,2,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
	Name         string                      `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	CreatedTime  *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=created_time,json=createdTime" json:"created_time,omitempty"`
	ModifiedTime *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=modified_time,json=modifiedTime" json:"modified_time,omitempty"`
	Version      int64                       `protobuf:"fixed64,6,opt,name=version" json:"version,omitempty"`
}

func (m *ApiPicTag) Reset()                    { *m = ApiPicTag{} }
func (m *ApiPicTag) String() string            { return proto.CompactTextString(m) }
func (*ApiPicTag) ProtoMessage()               {}
func (*ApiPicTag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ApiPicTag) GetCreatedTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *ApiPicTag) GetModifiedTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ModifiedTime
	}
	return nil
}

type LookupPicDetailsResponse struct {
	Pic    *ApiPic      `protobuf:"bytes,1,opt,name=pic" json:"pic,omitempty"`
	PicTag []*ApiPicTag `protobuf:"bytes,2,rep,name=pic_tag,json=picTag" json:"pic_tag,omitempty"`
}

func (m *LookupPicDetailsResponse) Reset()                    { *m = LookupPicDetailsResponse{} }
func (m *LookupPicDetailsResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupPicDetailsResponse) ProtoMessage()               {}
func (*LookupPicDetailsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LookupPicDetailsResponse) GetPic() *ApiPic {
	if m != nil {
		return m.Pic
	}
	return nil
}

func (m *LookupPicDetailsResponse) GetPicTag() []*ApiPicTag {
	if m != nil {
		return m.PicTag
	}
	return nil
}

type IndexResponse struct {
	Pic []*ApiPic `protobuf:"bytes,1,rep,name=pic" json:"pic,omitempty"`
}

func (m *IndexResponse) Reset()                    { *m = IndexResponse{} }
func (m *IndexResponse) String() string            { return proto.CompactTextString(m) }
func (*IndexResponse) ProtoMessage()               {}
func (*IndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IndexResponse) GetPic() []*ApiPic {
	if m != nil {
		return m.Pic
	}
	return nil
}

type AddPicTagsResponse struct {
}

func (m *AddPicTagsResponse) Reset()                    { *m = AddPicTagsResponse{} }
func (m *AddPicTagsResponse) String() string            { return proto.CompactTextString(m) }
func (*AddPicTagsResponse) ProtoMessage()               {}
func (*AddPicTagsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type CreatePicResponse struct {
	Pic *ApiPic `protobuf:"bytes,1,opt,name=pic" json:"pic,omitempty"`
}

func (m *CreatePicResponse) Reset()                    { *m = CreatePicResponse{} }
func (m *CreatePicResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePicResponse) ProtoMessage()               {}
func (*CreatePicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreatePicResponse) GetPic() *ApiPic {
	if m != nil {
		return m.Pic
	}
	return nil
}

type FindSimilarPicsResponse struct {
	Id []string `protobuf:"bytes,1,rep,name=id" json:"id,omitempty"`
}

func (m *FindSimilarPicsResponse) Reset()                    { *m = FindSimilarPicsResponse{} }
func (m *FindSimilarPicsResponse) String() string            { return proto.CompactTextString(m) }
func (*FindSimilarPicsResponse) ProtoMessage()               {}
func (*FindSimilarPicsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*ApiPic)(nil), "pixur.api.ApiPic")
	proto.RegisterType((*ApiPicTag)(nil), "pixur.api.ApiPicTag")
	proto.RegisterType((*LookupPicDetailsResponse)(nil), "pixur.api.LookupPicDetailsResponse")
	proto.RegisterType((*IndexResponse)(nil), "pixur.api.IndexResponse")
	proto.RegisterType((*AddPicTagsResponse)(nil), "pixur.api.AddPicTagsResponse")
	proto.RegisterType((*CreatePicResponse)(nil), "pixur.api.CreatePicResponse")
	proto.RegisterType((*FindSimilarPicsResponse)(nil), "pixur.api.FindSimilarPicsResponse")
}

var fileDescriptor0 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd4, 0x3c,
	0x10, 0xc7, 0x95, 0xdd, 0x6e, 0xba, 0x99, 0xb4, 0xdf, 0xd7, 0x5a, 0x4b, 0x31, 0x95, 0x80, 0x10,
	0x2e, 0xe9, 0x81, 0x54, 0x5a, 0x8a, 0xc4, 0x05, 0xa1, 0xd2, 0x0a, 0xa9, 0x88, 0x43, 0x65, 0x96,
	0x0b, 0x97, 0xc8, 0x8d, 0xdd, 0xec, 0x88, 0xc4, 0xb1, 0x12, 0xa7, 0x2d, 0xef, 0xc9, 0x3b, 0xf0,
	0x1a, 0xc8, 0xce, 0x66, 0x69, 0xa9, 0x10, 0xe2, 0x96, 0xf9, 0xcf, 0x7f, 0xc6, 0xe3, 0xdf, 0xc4,
	0x10, 0x70, 0x8d, 0xa9, 0x6e, 0x6a, 0x53, 0x93, 0x40, 0xe3, 0x4d, 0xd7, 0xa4, 0x5c, 0xe3, 0xfe,
	0x93, 0xa2, 0xae, 0x8b, 0x52, 0x1e, 0xba, 0xc4, 0x45, 0x77, 0x79, 0x28, 0xba, 0x86, 0x1b, 0xac,
	0x55, 0x6f, 0xdd, 0x7f, 0xfa, 0x7b, 0xde, 0x60, 0x25, 0x5b, 0xc3, 0x2b, 0xdd, 0x1b, 0xe2, 0xef,
	0x23, 0xf0, 0x8f, 0x35, 0x9e, 0x63, 0x4e, 0xfe, 0x83, 0x11, 0x0a, 0xea, 0x45, 0x5e, 0x12, 0xb0,
	0x11, 0x0a, 0x32, 0x83, 0xc9, 0x35, 0x0a, 0xb3, 0xa4, 0xa3, 0xc8, 0x4b, 0x26, 0xac, 0x0f, 0xc8,
	0x1e, 0xf8, 0x4b, 0x89, 0xc5, 0xd2, 0xd0, 0xb1, 0x93, 0x57, 0x11, 0xa1, 0xb0, 0x79, 0x25, 0x9b,
	0x16, 0x6b, 0x45, 0x37, 0x22, 0x2f, 0xd9, 0x61, 0x43, 0x48, 0x08, 0x6c, 0x98, 0x6f, 0x5a, 0xd2,
	0x89, 0xeb, 0xec, 0xbe, 0xc9, 0x33, 0xd8, 0x6a, 0x64, 0xc9, 0x0d, 0x5e, 0xc9, 0xac, 0x6b, 0x4a,
	0xea, 0xbb, 0x5c, 0x38, 0x68, 0x9f, 0x9b, 0x92, 0x1c, 0xc1, 0x9e, 0x59, 0x76, 0xd5, 0x85, 0xe2,
	0x58, 0x66, 0x77, 0xcc, 0x9b, 0xce, 0x3c, 0x5b, 0x67, 0xd9, 0xad, 0xaa, 0x03, 0xd8, 0xd1, 0x52,
	0x09, 0x54, 0x45, 0x26, 0x64, 0x29, 0x2d, 0x0a, 0x1a, 0x44, 0x5e, 0x32, 0x65, 0xff, 0xaf, 0xf4,
	0xd3, 0x95, 0x4c, 0x1e, 0x03, 0x5c, 0xa1, 0xbc, 0xce, 0xf2, 0xba, 0x53, 0x86, 0x42, 0xe4, 0x25,
	0x63, 0x16, 0x58, 0xe5, 0xc4, 0x0a, 0xe4, 0x15, 0x4c, 0x07, 0x98, 0x34, 0x8c, 0xbc, 0x24, 0x9c,
	0x3f, 0x4a, 0x7b, 0x9a, 0xe9, 0x40, 0x33, 0x3d, 0x5d, 0x19, 0xd8, 0xda, 0xfa, 0x61, 0x63, 0x3a,
	0xdd, 0x09, 0xe2, 0x1f, 0x1e, 0x04, 0x3d, 0xd6, 0x05, 0x2f, 0xc8, 0x03, 0xf0, 0x35, 0xe6, 0xd9,
	0x9a, 0xee, 0x44, 0x63, 0x7e, 0x26, 0xac, 0x6c, 0x78, 0x61, 0xe5, 0x51, 0x2f, 0x1b, 0x5e, 0x9c,
	0x09, 0xcb, 0x4b, 0xf1, 0x4a, 0x3a, 0xbe, 0x01, 0x73, 0xdf, 0xe4, 0x0d, 0x6c, 0xe5, 0x8d, 0xe4,
	0x46, 0x8a, 0xcc, 0x6e, 0xd0, 0x21, 0x0e, 0xe7, 0xfb, 0xf7, 0x06, 0x5a, 0x0c, 0xeb, 0x65, 0xe1,
	0xca, 0x6f, 0x15, 0xf2, 0x16, 0xb6, 0xab, 0x5a, 0xe0, 0x25, 0x0e, 0xf5, 0x93, 0xbf, 0xd6, 0x6f,
	0x0d, 0x05, 0xae, 0xc1, 0xad, 0xed, 0xfa, 0x77, 0xb6, 0x1b, 0x2b, 0xa0, 0x1f, 0xeb, 0xfa, 0x6b,
	0xa7, 0xcf, 0x31, 0x3f, 0x95, 0x86, 0x63, 0xd9, 0x32, 0xd9, 0xea, 0x5a, 0xb5, 0x92, 0x3c, 0x87,
	0xb1, 0xc6, 0xdc, 0x5d, 0x3a, 0x9c, 0xef, 0xa6, 0xeb, 0xdf, 0x36, 0xed, 0xd1, 0x30, 0x9b, 0x25,
	0x2f, 0x60, 0xd3, 0xc2, 0x31, 0xbc, 0xa0, 0xa3, 0x68, 0x9c, 0x84, 0xf3, 0xd9, 0x3d, 0xe3, 0x82,
	0x17, 0xcc, 0x12, 0x5c, 0xf0, 0x22, 0x3e, 0x82, 0xed, 0x33, 0x25, 0xe4, 0xcd, 0xfd, 0x43, 0xc6,
	0x7f, 0x3e, 0x24, 0x9e, 0x01, 0x39, 0x16, 0xa2, 0x6f, 0xb5, 0x9e, 0x2f, 0x7e, 0x0d, 0xbb, 0x27,
	0x8e, 0x92, 0xf5, 0xfd, 0xcb, 0xd0, 0xf1, 0x01, 0x3c, 0x7c, 0x8f, 0x4a, 0x7c, 0xc2, 0x0a, 0x4b,
	0xde, 0x9c, 0x63, 0xfe, 0xeb, 0xd2, 0xc3, 0x33, 0x1a, 0xf7, 0xcf, 0xe8, 0x1d, 0x7c, 0x99, 0x2e,
	0xb9, 0x12, 0xa5, 0x6c, 0xda, 0x0b, 0xdf, 0x81, 0x7e, 0xf9, 0x33, 0x00, 0x00, 0xff, 0xff, 0x41,
	0x06, 0x74, 0x3c, 0xcd, 0x03, 0x00, 0x00,
}
