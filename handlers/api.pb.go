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
	LookupPicDetailsRequest
	LookupPicDetailsResponse
	IndexRequest
	IndexResponse
	AddPicTagsRequest
	AddPicTagsResponse
	CreatePicRequest
	CreatePicResponse
	FindSimilarPicsRequest
	FindSimilarPicsResponse
	IncrementViewCountRequest
	IncrementViewCountResponse
	PurgePicRequest
	PurgePicResponse
	SoftDeletePicRequest
	SoftDeletePicResponse
	UpsertPicRequest
	UpsertPicResponse
	GetXsrfTokenRequest
	GetXsrfTokenResponse
	CreateUserRequest
	CreateUserResponse
	GetRefreshTokenRequest
	GetRefreshTokenResponse
	PwtHeader
	PwtPayload
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

// Copy of schema.proto
type DeletionReason int32

const (
	// The reason is not know, due to limitations of proto
	DeletionReason_UNKNOWN DeletionReason = 0
	// No specific reason.  This is a catch-all reason.
	DeletionReason_NONE DeletionReason = 1
	// The pic is in violation of the rules.
	DeletionReason_RULE_VIOLATION DeletionReason = 2
)

var DeletionReason_name = map[int32]string{
	0: "UNKNOWN",
	1: "NONE",
	2: "RULE_VIOLATION",
}
var DeletionReason_value = map[string]int32{
	"UNKNOWN":        0,
	"NONE":           1,
	"RULE_VIOLATION": 2,
}

func (x DeletionReason) String() string {
	return proto.EnumName(DeletionReason_name, int32(x))
}
func (DeletionReason) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PwtHeader_Algorithm int32

const (
	PwtHeader_UNKNOWN PwtHeader_Algorithm = 0
	PwtHeader_HS256   PwtHeader_Algorithm = 1
	PwtHeader_RS256   PwtHeader_Algorithm = 2
)

var PwtHeader_Algorithm_name = map[int32]string{
	0: "UNKNOWN",
	1: "HS256",
	2: "RS256",
}
var PwtHeader_Algorithm_value = map[string]int32{
	"UNKNOWN": 0,
	"HS256":   1,
	"RS256":   2,
}

func (x PwtHeader_Algorithm) String() string {
	return proto.EnumName(PwtHeader_Algorithm_name, int32(x))
}
func (PwtHeader_Algorithm) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{26, 0} }

type PwtPayload_Type int32

const (
	PwtPayload_UNKNOWN PwtPayload_Type = 0
	PwtPayload_REFRESH PwtPayload_Type = 1
	PwtPayload_AUTH    PwtPayload_Type = 2
	PwtPayload_PIX     PwtPayload_Type = 3
)

var PwtPayload_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "REFRESH",
	2: "AUTH",
	3: "PIX",
}
var PwtPayload_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"REFRESH": 1,
	"AUTH":    2,
	"PIX":     3,
}

func (x PwtPayload_Type) String() string {
	return proto.EnumName(PwtPayload_Type_name, int32(x))
}
func (PwtPayload_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{27, 0} }

type ApiPic struct {
	// id is the unique identifier for the pic, in varint form
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// width is the width of pic
	Width int32 `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	// height is the height of the pic
	Height int32 `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	// version is the version of the pic.  It is used when updating the pic.
	Version int64 `protobuf:"fixed64,4,opt,name=version" json:"version,omitempty"`
	// type is the mime type of the pic. (JPEG, GIF, PNG, or WEBM)
	Type string `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	// relative_url is the location of the pic file relative to the root dir.
	RelativeUrl string `protobuf:"bytes,6,opt,name=relative_url,json=relativeUrl" json:"relative_url,omitempty"`
	// thumbnail_relative_url is the location of the the pic thumbnail
	// relative to the root dir.
	ThumbnailRelativeUrl string `protobuf:"bytes,7,opt,name=thumbnail_relative_url,json=thumbnailRelativeUrl" json:"thumbnail_relative_url,omitempty"`
	// pending_deletion indicates if the pic may be deleted soon.
	PendingDeletion bool `protobuf:"varint,9,opt,name=pending_deletion,json=pendingDeletion" json:"pending_deletion,omitempty"`
	// view_count is the number of views this picture has received.
	ViewCount int64 `protobuf:"varint,10,opt,name=view_count,json=viewCount" json:"view_count,omitempty"`
	// duration is present if the image is animated (GIF or WEBM).  Note that
	// GIFs duration is not well defined and is subject to reinterpretation.
	Duration *google_protobuf.Duration `protobuf:"bytes,11,opt,name=duration" json:"duration,omitempty"`
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
	// pic_id is the unique identifier for the pic, in varint form
	PicId string `protobuf:"bytes,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
	// tag_id is the unique identifier for the tag, in varint form
	TagId string `protobuf:"bytes,2,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
	// name is the tag name in utf8 form
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// created_time is when the tag was created.
	CreatedTime *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=created_time,json=createdTime" json:"created_time,omitempty"`
	// modified_time is when the tag was last modified.
	ModifiedTime *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=modified_time,json=modifiedTime" json:"modified_time,omitempty"`
	// version is the version of the tag.  It is used when updating the tag.
	Version int64 `protobuf:"fixed64,6,opt,name=version" json:"version,omitempty"`
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

type LookupPicDetailsRequest struct {
	PicId string `protobuf:"bytes,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
}

func (m *LookupPicDetailsRequest) Reset()                    { *m = LookupPicDetailsRequest{} }
func (m *LookupPicDetailsRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupPicDetailsRequest) ProtoMessage()               {}
func (*LookupPicDetailsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type LookupPicDetailsResponse struct {
	Pic    *ApiPic      `protobuf:"bytes,1,opt,name=pic" json:"pic,omitempty"`
	PicTag []*ApiPicTag `protobuf:"bytes,2,rep,name=pic_tag,json=picTag" json:"pic_tag,omitempty"`
}

func (m *LookupPicDetailsResponse) Reset()                    { *m = LookupPicDetailsResponse{} }
func (m *LookupPicDetailsResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupPicDetailsResponse) ProtoMessage()               {}
func (*LookupPicDetailsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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

type IndexRequest struct {
	StartPicId string `protobuf:"bytes,1,opt,name=start_pic_id,json=startPicId" json:"start_pic_id,omitempty"`
}

func (m *IndexRequest) Reset()                    { *m = IndexRequest{} }
func (m *IndexRequest) String() string            { return proto.CompactTextString(m) }
func (*IndexRequest) ProtoMessage()               {}
func (*IndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type IndexResponse struct {
	Pic []*ApiPic `protobuf:"bytes,1,rep,name=pic" json:"pic,omitempty"`
}

func (m *IndexResponse) Reset()                    { *m = IndexResponse{} }
func (m *IndexResponse) String() string            { return proto.CompactTextString(m) }
func (*IndexResponse) ProtoMessage()               {}
func (*IndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IndexResponse) GetPic() []*ApiPic {
	if m != nil {
		return m.Pic
	}
	return nil
}

type AddPicTagsRequest struct {
	PicId string   `protobuf:"bytes,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
	Tag   []string `protobuf:"bytes,2,rep,name=tag" json:"tag,omitempty"`
}

func (m *AddPicTagsRequest) Reset()                    { *m = AddPicTagsRequest{} }
func (m *AddPicTagsRequest) String() string            { return proto.CompactTextString(m) }
func (*AddPicTagsRequest) ProtoMessage()               {}
func (*AddPicTagsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type AddPicTagsResponse struct {
}

func (m *AddPicTagsResponse) Reset()                    { *m = AddPicTagsResponse{} }
func (m *AddPicTagsResponse) String() string            { return proto.CompactTextString(m) }
func (*AddPicTagsResponse) ProtoMessage()               {}
func (*AddPicTagsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type CreatePicRequest struct {
	// file_data is the contents of the pic.  Either file_data or file_url must be present.
	// If sent as part of a stream, this field will be the only one read after the first message.
	FileData []byte `protobuf:"bytes,1,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
	// optional, only used if file_data is present.
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	// URL to get the pic from.  Either file_data or file_url must be present.
	FileUrl string   `protobuf:"bytes,3,opt,name=file_url,json=fileUrl" json:"file_url,omitempty"`
	Tag     []string `protobuf:"bytes,4,rep,name=tag" json:"tag,omitempty"`
}

func (m *CreatePicRequest) Reset()                    { *m = CreatePicRequest{} }
func (m *CreatePicRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePicRequest) ProtoMessage()               {}
func (*CreatePicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type CreatePicResponse struct {
	Pic *ApiPic `protobuf:"bytes,1,opt,name=pic" json:"pic,omitempty"`
}

func (m *CreatePicResponse) Reset()                    { *m = CreatePicResponse{} }
func (m *CreatePicResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePicResponse) ProtoMessage()               {}
func (*CreatePicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CreatePicResponse) GetPic() *ApiPic {
	if m != nil {
		return m.Pic
	}
	return nil
}

type FindSimilarPicsRequest struct {
	PicId string `protobuf:"bytes,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
}

func (m *FindSimilarPicsRequest) Reset()                    { *m = FindSimilarPicsRequest{} }
func (m *FindSimilarPicsRequest) String() string            { return proto.CompactTextString(m) }
func (*FindSimilarPicsRequest) ProtoMessage()               {}
func (*FindSimilarPicsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type FindSimilarPicsResponse struct {
	PicId []string `protobuf:"bytes,1,rep,name=pic_id,json=picId" json:"pic_id,omitempty"`
}

func (m *FindSimilarPicsResponse) Reset()                    { *m = FindSimilarPicsResponse{} }
func (m *FindSimilarPicsResponse) String() string            { return proto.CompactTextString(m) }
func (*FindSimilarPicsResponse) ProtoMessage()               {}
func (*FindSimilarPicsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type IncrementViewCountRequest struct {
	PicId string `protobuf:"bytes,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
}

func (m *IncrementViewCountRequest) Reset()                    { *m = IncrementViewCountRequest{} }
func (m *IncrementViewCountRequest) String() string            { return proto.CompactTextString(m) }
func (*IncrementViewCountRequest) ProtoMessage()               {}
func (*IncrementViewCountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type IncrementViewCountResponse struct {
}

func (m *IncrementViewCountResponse) Reset()                    { *m = IncrementViewCountResponse{} }
func (m *IncrementViewCountResponse) String() string            { return proto.CompactTextString(m) }
func (*IncrementViewCountResponse) ProtoMessage()               {}
func (*IncrementViewCountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type PurgePicRequest struct {
	PicId string `protobuf:"bytes,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
}

func (m *PurgePicRequest) Reset()                    { *m = PurgePicRequest{} }
func (m *PurgePicRequest) String() string            { return proto.CompactTextString(m) }
func (*PurgePicRequest) ProtoMessage()               {}
func (*PurgePicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type PurgePicResponse struct {
}

func (m *PurgePicResponse) Reset()                    { *m = PurgePicResponse{} }
func (m *PurgePicResponse) String() string            { return proto.CompactTextString(m) }
func (*PurgePicResponse) ProtoMessage()               {}
func (*PurgePicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type SoftDeletePicRequest struct {
	PicId        string                      `protobuf:"bytes,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
	Details      string                      `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
	Reason       DeletionReason              `protobuf:"varint,3,opt,name=reason,enum=pixur.api.DeletionReason" json:"reason,omitempty"`
	DeletionTime *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=deletion_time,json=deletionTime" json:"deletion_time,omitempty"`
}

func (m *SoftDeletePicRequest) Reset()                    { *m = SoftDeletePicRequest{} }
func (m *SoftDeletePicRequest) String() string            { return proto.CompactTextString(m) }
func (*SoftDeletePicRequest) ProtoMessage()               {}
func (*SoftDeletePicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *SoftDeletePicRequest) GetDeletionTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.DeletionTime
	}
	return nil
}

type SoftDeletePicResponse struct {
}

func (m *SoftDeletePicResponse) Reset()                    { *m = SoftDeletePicResponse{} }
func (m *SoftDeletePicResponse) String() string            { return proto.CompactTextString(m) }
func (*SoftDeletePicResponse) ProtoMessage()               {}
func (*SoftDeletePicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type UpsertPicRequest struct {
	Url     string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Data    []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Md5Hash []byte   `protobuf:"bytes,4,opt,name=md5_hash,json=md5Hash,proto3" json:"md5_hash,omitempty"`
	Tag     []string `protobuf:"bytes,5,rep,name=tag" json:"tag,omitempty"`
}

func (m *UpsertPicRequest) Reset()                    { *m = UpsertPicRequest{} }
func (m *UpsertPicRequest) String() string            { return proto.CompactTextString(m) }
func (*UpsertPicRequest) ProtoMessage()               {}
func (*UpsertPicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type UpsertPicResponse struct {
	Pic *ApiPic `protobuf:"bytes,1,opt,name=pic" json:"pic,omitempty"`
}

func (m *UpsertPicResponse) Reset()                    { *m = UpsertPicResponse{} }
func (m *UpsertPicResponse) String() string            { return proto.CompactTextString(m) }
func (*UpsertPicResponse) ProtoMessage()               {}
func (*UpsertPicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *UpsertPicResponse) GetPic() *ApiPic {
	if m != nil {
		return m.Pic
	}
	return nil
}

type GetXsrfTokenRequest struct {
}

func (m *GetXsrfTokenRequest) Reset()                    { *m = GetXsrfTokenRequest{} }
func (m *GetXsrfTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*GetXsrfTokenRequest) ProtoMessage()               {}
func (*GetXsrfTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

type GetXsrfTokenResponse struct {
	XsrfToken string `protobuf:"bytes,1,opt,name=xsrf_token,json=xsrfToken" json:"xsrf_token,omitempty"`
}

func (m *GetXsrfTokenResponse) Reset()                    { *m = GetXsrfTokenResponse{} }
func (m *GetXsrfTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*GetXsrfTokenResponse) ProtoMessage()               {}
func (*GetXsrfTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

type CreateUserRequest struct {
	// ident is the unique identity of the user being created, usually an email address
	Ident string `protobuf:"bytes,1,opt,name=ident" json:"ident,omitempty"`
	// secret is the secret string used to authenticate the user, usually a password
	Secret string `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

type CreateUserResponse struct {
}

func (m *CreateUserResponse) Reset()                    { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()               {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

type GetRefreshTokenRequest struct {
	// ident is the unique identity of the user being created, usually an email address
	Ident string `protobuf:"bytes,1,opt,name=ident" json:"ident,omitempty"`
	// secret is the secret string used to authenticate the user, usually a password
	Secret       string `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *GetRefreshTokenRequest) Reset()                    { *m = GetRefreshTokenRequest{} }
func (m *GetRefreshTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRefreshTokenRequest) ProtoMessage()               {}
func (*GetRefreshTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

type GetRefreshTokenResponse struct {
	RefreshToken   string      `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	AuthToken      string      `protobuf:"bytes,2,opt,name=auth_token,json=authToken" json:"auth_token,omitempty"`
	PixToken       string      `protobuf:"bytes,5,opt,name=pix_token,json=pixToken" json:"pix_token,omitempty"`
	RefreshPayload *PwtPayload `protobuf:"bytes,3,opt,name=refresh_payload,json=refreshPayload" json:"refresh_payload,omitempty"`
	AuthPayload    *PwtPayload `protobuf:"bytes,4,opt,name=auth_payload,json=authPayload" json:"auth_payload,omitempty"`
	PixPayload     *PwtPayload `protobuf:"bytes,6,opt,name=pix_payload,json=pixPayload" json:"pix_payload,omitempty"`
}

func (m *GetRefreshTokenResponse) Reset()                    { *m = GetRefreshTokenResponse{} }
func (m *GetRefreshTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRefreshTokenResponse) ProtoMessage()               {}
func (*GetRefreshTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

func (m *GetRefreshTokenResponse) GetRefreshPayload() *PwtPayload {
	if m != nil {
		return m.RefreshPayload
	}
	return nil
}

func (m *GetRefreshTokenResponse) GetAuthPayload() *PwtPayload {
	if m != nil {
		return m.AuthPayload
	}
	return nil
}

func (m *GetRefreshTokenResponse) GetPixPayload() *PwtPayload {
	if m != nil {
		return m.PixPayload
	}
	return nil
}

type PwtHeader struct {
	Algorithm PwtHeader_Algorithm `protobuf:"varint,1,opt,name=algorithm,enum=pixur.api.PwtHeader_Algorithm" json:"algorithm,omitempty"`
	Version   int64               `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *PwtHeader) Reset()                    { *m = PwtHeader{} }
func (m *PwtHeader) String() string            { return proto.CompactTextString(m) }
func (*PwtHeader) ProtoMessage()               {}
func (*PwtHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{26} }

type PwtPayload struct {
	Subject   string                      `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	NotBefore *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=not_before,json=notBefore" json:"not_before,omitempty"`
	NotAfter  *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=not_after,json=notAfter" json:"not_after,omitempty"`
	// represents when this token should be reverified.  Optional.
	SoftNotAfter  *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=soft_not_after,json=softNotAfter" json:"soft_not_after,omitempty"`
	Issuer        string                      `protobuf:"bytes,4,opt,name=issuer" json:"issuer,omitempty"`
	TokenId       int64                       `protobuf:"varint,5,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	TokenParentId int64                       `protobuf:"varint,8,opt,name=token_parent_id,json=tokenParentId" json:"token_parent_id,omitempty"`
	Type          PwtPayload_Type             `protobuf:"varint,7,opt,name=type,enum=pixur.api.PwtPayload_Type" json:"type,omitempty"`
}

func (m *PwtPayload) Reset()                    { *m = PwtPayload{} }
func (m *PwtPayload) String() string            { return proto.CompactTextString(m) }
func (*PwtPayload) ProtoMessage()               {}
func (*PwtPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{27} }

func (m *PwtPayload) GetNotBefore() *google_protobuf1.Timestamp {
	if m != nil {
		return m.NotBefore
	}
	return nil
}

func (m *PwtPayload) GetNotAfter() *google_protobuf1.Timestamp {
	if m != nil {
		return m.NotAfter
	}
	return nil
}

func (m *PwtPayload) GetSoftNotAfter() *google_protobuf1.Timestamp {
	if m != nil {
		return m.SoftNotAfter
	}
	return nil
}

func init() {
	proto.RegisterType((*ApiPic)(nil), "pixur.api.ApiPic")
	proto.RegisterType((*ApiPicTag)(nil), "pixur.api.ApiPicTag")
	proto.RegisterType((*LookupPicDetailsRequest)(nil), "pixur.api.LookupPicDetailsRequest")
	proto.RegisterType((*LookupPicDetailsResponse)(nil), "pixur.api.LookupPicDetailsResponse")
	proto.RegisterType((*IndexRequest)(nil), "pixur.api.IndexRequest")
	proto.RegisterType((*IndexResponse)(nil), "pixur.api.IndexResponse")
	proto.RegisterType((*AddPicTagsRequest)(nil), "pixur.api.AddPicTagsRequest")
	proto.RegisterType((*AddPicTagsResponse)(nil), "pixur.api.AddPicTagsResponse")
	proto.RegisterType((*CreatePicRequest)(nil), "pixur.api.CreatePicRequest")
	proto.RegisterType((*CreatePicResponse)(nil), "pixur.api.CreatePicResponse")
	proto.RegisterType((*FindSimilarPicsRequest)(nil), "pixur.api.FindSimilarPicsRequest")
	proto.RegisterType((*FindSimilarPicsResponse)(nil), "pixur.api.FindSimilarPicsResponse")
	proto.RegisterType((*IncrementViewCountRequest)(nil), "pixur.api.IncrementViewCountRequest")
	proto.RegisterType((*IncrementViewCountResponse)(nil), "pixur.api.IncrementViewCountResponse")
	proto.RegisterType((*PurgePicRequest)(nil), "pixur.api.PurgePicRequest")
	proto.RegisterType((*PurgePicResponse)(nil), "pixur.api.PurgePicResponse")
	proto.RegisterType((*SoftDeletePicRequest)(nil), "pixur.api.SoftDeletePicRequest")
	proto.RegisterType((*SoftDeletePicResponse)(nil), "pixur.api.SoftDeletePicResponse")
	proto.RegisterType((*UpsertPicRequest)(nil), "pixur.api.UpsertPicRequest")
	proto.RegisterType((*UpsertPicResponse)(nil), "pixur.api.UpsertPicResponse")
	proto.RegisterType((*GetXsrfTokenRequest)(nil), "pixur.api.GetXsrfTokenRequest")
	proto.RegisterType((*GetXsrfTokenResponse)(nil), "pixur.api.GetXsrfTokenResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "pixur.api.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "pixur.api.CreateUserResponse")
	proto.RegisterType((*GetRefreshTokenRequest)(nil), "pixur.api.GetRefreshTokenRequest")
	proto.RegisterType((*GetRefreshTokenResponse)(nil), "pixur.api.GetRefreshTokenResponse")
	proto.RegisterType((*PwtHeader)(nil), "pixur.api.PwtHeader")
	proto.RegisterType((*PwtPayload)(nil), "pixur.api.PwtPayload")
	proto.RegisterEnum("pixur.api.DeletionReason", DeletionReason_name, DeletionReason_value)
	proto.RegisterEnum("pixur.api.PwtHeader_Algorithm", PwtHeader_Algorithm_name, PwtHeader_Algorithm_value)
	proto.RegisterEnum("pixur.api.PwtPayload_Type", PwtPayload_Type_name, PwtPayload_Type_value)
}

var fileDescriptor0 = []byte{
	// 1530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x57, 0xdb, 0x72, 0x1b, 0x37,
	0x12, 0xdd, 0xe1, 0x7d, 0x9a, 0x94, 0x44, 0x61, 0x75, 0xa1, 0x28, 0xd9, 0xa6, 0xc7, 0xbb, 0x5b,
	0xdc, 0xad, 0x5a, 0xda, 0xa1, 0x2d, 0xc7, 0xa9, 0x38, 0x17, 0xd9, 0x92, 0x2d, 0x26, 0x2e, 0x8a,
	0x35, 0xa2, 0x1c, 0x57, 0xf2, 0xc0, 0x82, 0x38, 0x20, 0x89, 0x98, 0x9c, 0x99, 0x60, 0x40, 0x5d,
	0xbe, 0x23, 0xaf, 0x79, 0xce, 0x53, 0xfe, 0x20, 0xbf, 0x91, 0x87, 0xfc, 0x41, 0x7e, 0x23, 0x05,
	0x0c, 0x40, 0xce, 0xf0, 0x22, 0xc5, 0x95, 0x37, 0xa0, 0xfb, 0x9c, 0x46, 0x4f, 0x77, 0xa3, 0xd1,
	0x03, 0x26, 0xf6, 0x69, 0xcd, 0x67, 0x1e, 0xf7, 0x90, 0xe9, 0xd3, 0xab, 0x31, 0xab, 0x61, 0x9f,
	0x96, 0xef, 0xf6, 0x3d, 0xaf, 0x3f, 0x24, 0x0f, 0xa5, 0xe2, 0x7c, 0xdc, 0x7b, 0xe8, 0x8c, 0x19,
	0xe6, 0xd4, 0x73, 0x43, 0x68, 0xf9, 0xde, 0xac, 0x9e, 0xd3, 0x11, 0x09, 0x38, 0x1e, 0xf9, 0x21,
	0xc0, 0xfa, 0x2d, 0x01, 0x99, 0x03, 0x9f, 0xb6, 0x68, 0x17, 0xad, 0x42, 0x82, 0x3a, 0x25, 0xa3,
	0x62, 0x54, 0x4d, 0x3b, 0x41, 0x1d, 0xb4, 0x01, 0xe9, 0x4b, 0xea, 0xf0, 0x41, 0x29, 0x51, 0x31,
	0xaa, 0x69, 0x3b, 0xdc, 0xa0, 0x2d, 0xc8, 0x0c, 0x08, 0xed, 0x0f, 0x78, 0x29, 0x29, 0xc5, 0x6a,
	0x87, 0x4a, 0x90, 0xbd, 0x20, 0x2c, 0xa0, 0x9e, 0x5b, 0x4a, 0x55, 0x8c, 0x6a, 0xd1, 0xd6, 0x5b,
	0x84, 0x20, 0xc5, 0xaf, 0x7d, 0x52, 0x4a, 0x4b, 0xcb, 0x72, 0x8d, 0xee, 0x43, 0x81, 0x91, 0x21,
	0xe6, 0xf4, 0x82, 0x74, 0xc6, 0x6c, 0x58, 0xca, 0x48, 0x5d, 0x5e, 0xcb, 0xce, 0xd8, 0x10, 0x3d,
	0x81, 0x2d, 0x3e, 0x18, 0x8f, 0xce, 0x5d, 0x4c, 0x87, 0x9d, 0x18, 0x38, 0x2b, 0xc1, 0x1b, 0x13,
	0xad, 0x1d, 0x61, 0xfd, 0x17, 0x8a, 0x3e, 0x71, 0x1d, 0xea, 0xf6, 0x3b, 0x0e, 0x19, 0x12, 0x11,
	0x8a, 0x92, 0x59, 0x31, 0xaa, 0x39, 0x7b, 0x4d, 0xc9, 0x0f, 0x95, 0x18, 0xdd, 0x01, 0xb8, 0xa0,
	0xe4, 0xb2, 0xd3, 0xf5, 0xc6, 0x2e, 0x2f, 0x41, 0xc5, 0xa8, 0x26, 0x6d, 0x53, 0x48, 0x5e, 0x0a,
	0x01, 0xda, 0x87, 0x9c, 0x0e, 0x66, 0x29, 0x5f, 0x31, 0xaa, 0xf9, 0xfa, 0x4e, 0x2d, 0x8c, 0x66,
	0x4d, 0x47, 0xb3, 0x76, 0xa8, 0x00, 0xf6, 0x04, 0xfa, 0x55, 0x2a, 0x97, 0x2b, 0x9a, 0xd6, 0x1f,
	0x06, 0x98, 0x61, 0x58, 0xdb, 0xb8, 0x8f, 0x36, 0x21, 0xe3, 0xd3, 0x6e, 0x67, 0x12, 0xdd, 0xb4,
	0x4f, 0xbb, 0x0d, 0x47, 0x88, 0x39, 0xee, 0x0b, 0x71, 0x22, 0x14, 0x73, 0xdc, 0x6f, 0x38, 0x22,
	0x5e, 0x2e, 0x1e, 0x11, 0x19, 0x5f, 0xd3, 0x96, 0x6b, 0xf4, 0x19, 0x14, 0xba, 0x8c, 0x60, 0x4e,
	0x9c, 0x8e, 0xc8, 0xa0, 0x0c, 0x71, 0xbe, 0x5e, 0x9e, 0x73, 0xa8, 0xad, 0xd3, 0x6b, 0xe7, 0x15,
	0x5e, 0x48, 0xd0, 0x17, 0xb0, 0x32, 0xf2, 0x1c, 0xda, 0xa3, 0x9a, 0x9f, 0xbe, 0x95, 0x5f, 0xd0,
	0x04, 0x69, 0x20, 0x92, 0xdd, 0x4c, 0x2c, 0xbb, 0xd6, 0x23, 0xd8, 0x7e, 0xe3, 0x79, 0xef, 0xc7,
	0x7e, 0x8b, 0x76, 0x0f, 0x09, 0xc7, 0x74, 0x18, 0xd8, 0xe4, 0x87, 0x31, 0x09, 0xf8, 0x92, 0xcf,
	0xb6, 0x5c, 0x28, 0xcd, 0x33, 0x02, 0xdf, 0x73, 0x03, 0x82, 0x1e, 0x40, 0xd2, 0xa7, 0x5d, 0x89,
	0xcf, 0xd7, 0xd7, 0x6b, 0x93, 0x42, 0xaf, 0x85, 0xc1, 0xb4, 0x85, 0x16, 0xfd, 0x1f, 0xb2, 0xc2,
	0x2e, 0xc7, 0xfd, 0x52, 0xa2, 0x92, 0xac, 0xe6, 0xeb, 0x1b, 0x73, 0xc0, 0x36, 0xee, 0xdb, 0xe2,
	0xf0, 0x36, 0xee, 0x5b, 0x8f, 0xa0, 0xd0, 0x70, 0x1d, 0x72, 0xa5, 0xdd, 0xaa, 0x40, 0x21, 0xe0,
	0x98, 0xf1, 0x4e, 0xcc, 0x39, 0x90, 0xb2, 0x96, 0xf4, 0xf0, 0x09, 0xac, 0x28, 0xc6, 0xac, 0x5b,
	0xc9, 0xe5, 0x6e, 0x59, 0xcf, 0x61, 0xfd, 0xc0, 0x71, 0xc2, 0xc3, 0x6f, 0x89, 0x01, 0x2a, 0x42,
	0x52, 0xbb, 0x6f, 0xda, 0x62, 0x69, 0x6d, 0x00, 0x8a, 0xb2, 0xc3, 0x83, 0xad, 0x6b, 0x28, 0xbe,
	0x94, 0x79, 0x14, 0xa7, 0x28, 0x93, 0xbb, 0x60, 0xf6, 0xe8, 0x90, 0x74, 0x1c, 0xcc, 0xb1, 0xb4,
	0x5a, 0xb0, 0x73, 0x42, 0x70, 0x88, 0x39, 0x9e, 0x28, 0x65, 0x05, 0x85, 0x65, 0x25, 0x95, 0x4d,
	0x51, 0x45, 0x3b, 0x20, 0xd7, 0xf2, 0x12, 0x85, 0xd5, 0x95, 0x15, 0x7b, 0x71, 0x6f, 0x94, 0x43,
	0xa9, 0xa9, 0x43, 0xcf, 0x60, 0x3d, 0x72, 0xf4, 0x07, 0xe4, 0xc7, 0x7a, 0x08, 0x5b, 0xaf, 0xa8,
	0xeb, 0x9c, 0xd2, 0x11, 0x1d, 0x62, 0xd6, 0xa2, 0xdd, 0xdb, 0x2a, 0xe2, 0x11, 0x6c, 0xcf, 0x11,
	0xd4, 0x81, 0x51, 0x46, 0x72, 0xca, 0xa8, 0xc3, 0x4e, 0xc3, 0xed, 0x32, 0x32, 0x22, 0x2e, 0x7f,
	0xab, 0xaf, 0xec, 0x2d, 0xa7, 0xec, 0x41, 0x79, 0x11, 0x47, 0x45, 0xba, 0x0a, 0x6b, 0xad, 0x31,
	0xeb, 0x47, 0x03, 0xbd, 0xc4, 0x0e, 0x82, 0xe2, 0x14, 0xa9, 0xd8, 0xbf, 0x1a, 0xb0, 0x71, 0xea,
	0xf5, 0xb8, 0x6c, 0x2e, 0xb7, 0xdb, 0x10, 0xf7, 0xc9, 0x09, 0x4b, 0x5f, 0x25, 0x49, 0x6f, 0xd1,
	0x47, 0x90, 0x61, 0x04, 0x07, 0x9e, 0x2b, 0x33, 0xb4, 0x5a, 0xdf, 0x89, 0x04, 0x59, 0xb7, 0x2e,
	0x5b, 0x02, 0x6c, 0x05, 0x14, 0xb7, 0x5b, 0xf7, 0xba, 0xbf, 0xda, 0x1d, 0x0a, 0x9a, 0x20, 0x44,
	0xd6, 0x36, 0x6c, 0xce, 0x38, 0x3f, 0x2d, 0xbf, 0x33, 0x3f, 0x20, 0xf2, 0x5e, 0xe8, 0x2f, 0x2a,
	0x42, 0x52, 0xd4, 0x4f, 0xf8, 0x39, 0x62, 0x39, 0x69, 0x58, 0x89, 0x48, 0xc3, 0x42, 0x90, 0x92,
	0xf5, 0x99, 0x94, 0xf5, 0x29, 0xd7, 0xa2, 0xfc, 0x46, 0xce, 0x7e, 0x67, 0x80, 0x83, 0x81, 0x74,
	0xb1, 0x60, 0x67, 0x47, 0xce, 0xfe, 0x31, 0x0e, 0x06, 0xba, 0xfc, 0xd2, 0xb1, 0xf2, 0x8b, 0x1c,
	0xfd, 0x21, 0xe5, 0xb7, 0x09, 0xff, 0x7c, 0x4d, 0xf8, 0xbb, 0x80, 0xf5, 0xda, 0xde, 0x7b, 0xe2,
	0x2a, 0xbf, 0xad, 0x7d, 0xd8, 0x88, 0x8b, 0x95, 0xcd, 0x3b, 0x00, 0x57, 0x01, 0xeb, 0x75, 0xb8,
	0x90, 0xaa, 0xcf, 0x32, 0xaf, 0x34, 0xcc, 0x3a, 0xd0, 0xd7, 0xe0, 0x2c, 0x20, 0x4c, 0xc7, 0x60,
	0x03, 0xd2, 0xd4, 0x21, 0x2e, 0xd7, 0x49, 0x95, 0x1b, 0xf1, 0x34, 0x06, 0xa4, 0xcb, 0x08, 0x57,
	0x91, 0x50, 0x3b, 0x71, 0xb5, 0xa3, 0x26, 0x54, 0x6c, 0xdf, 0xc3, 0xd6, 0x6b, 0xc2, 0x6d, 0xd2,
	0x63, 0x24, 0x18, 0x44, 0x3d, 0xfd, 0x30, 0xeb, 0xe8, 0x01, 0xac, 0xb0, 0xd0, 0x88, 0xfa, 0x84,
	0xf0, 0x66, 0x17, 0x58, 0xc4, 0xb2, 0xf5, 0x4b, 0x02, 0xb6, 0xe7, 0x4e, 0x9b, 0x04, 0x75, 0xc6,
	0x80, 0x31, 0x6f, 0x40, 0x44, 0x09, 0x8f, 0xb9, 0x46, 0x84, 0x1e, 0x98, 0x42, 0x12, 0xaa, 0x77,
	0x41, 0x0c, 0x25, 0x4a, 0x1b, 0x3e, 0xf4, 0x39, 0x9f, 0x5e, 0x85, 0xca, 0xcf, 0x61, 0x4d, 0x1f,
	0xe0, 0xe3, 0xeb, 0xa1, 0x87, 0x1d, 0xe9, 0x63, 0xbe, 0xbe, 0x19, 0xc9, 0x60, 0xeb, 0x92, 0xb7,
	0x42, 0xa5, 0xbd, 0xaa, 0xd0, 0x6a, 0x8f, 0x9e, 0x41, 0x41, 0x9e, 0xad, 0xc9, 0xa9, 0x9b, 0xc8,
	0x79, 0x01, 0xd5, 0xcc, 0xa7, 0x90, 0x17, 0x6e, 0x69, 0x62, 0xe6, 0x26, 0x22, 0xf8, 0xf4, 0x4a,
	0xad, 0xad, 0x1f, 0x0d, 0x30, 0x5b, 0x97, 0xfc, 0x98, 0x60, 0x87, 0x30, 0xf4, 0x1c, 0x4c, 0x3c,
	0xec, 0x7b, 0x8c, 0xf2, 0xc1, 0x48, 0x06, 0x67, 0xb5, 0x7e, 0x37, 0x6e, 0x23, 0x04, 0xd6, 0x0e,
	0x34, 0xca, 0x9e, 0x12, 0xa2, 0x4f, 0x67, 0x42, 0xce, 0x18, 0x93, 0xa7, 0xb3, 0x06, 0xe6, 0x84,
	0x81, 0xf2, 0x90, 0x3d, 0x6b, 0x7e, 0xdd, 0x3c, 0xf9, 0xa6, 0x59, 0xfc, 0x07, 0x32, 0x21, 0x7d,
	0x7c, 0x5a, 0xdf, 0x7f, 0x5a, 0x34, 0xc4, 0xd2, 0x96, 0xcb, 0x84, 0xf5, 0x53, 0x12, 0x60, 0xea,
	0xb0, 0x30, 0x1c, 0x8c, 0xcf, 0xbf, 0x27, 0x5d, 0x5d, 0x28, 0x7a, 0x8b, 0x3e, 0x01, 0x70, 0x3d,
	0xde, 0x39, 0x27, 0x3d, 0x8f, 0x85, 0xd7, 0xf2, 0xe6, 0x6e, 0x60, 0xba, 0x1e, 0x7f, 0x21, 0xc1,
	0xe8, 0x63, 0x10, 0x9b, 0x0e, 0xee, 0x71, 0xc2, 0x54, 0x96, 0x6e, 0x62, 0xe6, 0x5c, 0x8f, 0x1f,
	0x08, 0x2c, 0xfa, 0x12, 0x56, 0x03, 0xaf, 0xc7, 0x3b, 0x53, 0x76, 0xe6, 0xf6, 0x2e, 0x24, 0x18,
	0x4d, 0x6d, 0x61, 0x0b, 0x32, 0x34, 0x08, 0xc6, 0x84, 0xc9, 0x04, 0x9b, 0xb6, 0xda, 0x89, 0xb6,
	0x21, 0xeb, 0x4a, 0x34, 0xd1, 0x74, 0x18, 0x41, 0xb9, 0x6f, 0x38, 0xe8, 0x3f, 0xb0, 0x16, 0xaa,
	0x7c, 0xcc, 0x88, 0xcb, 0x05, 0x22, 0x27, 0x11, 0x2b, 0x52, 0xdc, 0x92, 0xd2, 0x86, 0x83, 0x6a,
	0x6a, 0x04, 0xcd, 0xca, 0xe4, 0x95, 0x17, 0x16, 0x40, 0xad, 0x7d, 0xed, 0x93, 0x70, 0x3c, 0xb5,
	0x1e, 0x43, 0x4a, 0xec, 0xe2, 0x49, 0xc9, 0x43, 0xd6, 0x3e, 0x7a, 0x65, 0x1f, 0x9d, 0x1e, 0x17,
	0x0d, 0x94, 0x83, 0xd4, 0xc1, 0x59, 0xfb, 0xb8, 0x98, 0x40, 0x59, 0x48, 0xb6, 0x1a, 0xef, 0x8a,
	0xc9, 0xff, 0x7d, 0x0a, 0xab, 0xf1, 0x06, 0x1d, 0xa7, 0xe7, 0x20, 0xd5, 0x3c, 0x69, 0x1e, 0x15,
	0x0d, 0x84, 0x60, 0xd5, 0x3e, 0x7b, 0x73, 0xd4, 0x79, 0xdb, 0x38, 0x79, 0x73, 0xd0, 0x6e, 0x9c,
	0x34, 0x8b, 0x89, 0xfa, 0xcf, 0x39, 0x28, 0xb4, 0x84, 0x57, 0xa7, 0x84, 0x5d, 0xd0, 0x2e, 0x41,
	0x0d, 0x80, 0x69, 0xd3, 0x40, 0x7b, 0x11, 0x97, 0xe7, 0xda, 0x51, 0xf9, 0xce, 0x12, 0xad, 0xba,
	0xe0, 0x27, 0x50, 0x88, 0x76, 0x3e, 0x14, 0x2d, 0xde, 0x05, 0x9d, 0xb2, 0x7c, 0x6f, 0xa9, 0x5e,
	0x19, 0x3c, 0x06, 0x73, 0xd2, 0x9b, 0xd1, 0x6e, 0x04, 0x3d, 0xfb, 0x58, 0x94, 0xf7, 0x16, 0x2b,
	0x43, 0x3b, 0x55, 0x03, 0xd9, 0xb0, 0x12, 0x7b, 0x79, 0x50, 0xf4, 0xec, 0x45, 0x0f, 0x6a, 0xb9,
	0xb2, 0x1c, 0xa0, 0xbc, 0x7b, 0x09, 0x39, 0xfd, 0x3e, 0xa3, 0x58, 0xaa, 0xe3, 0xcf, 0x7b, 0x79,
	0x77, 0xa1, 0x4e, 0x19, 0xc1, 0x80, 0xe6, 0x87, 0x05, 0xf4, 0xaf, 0x08, 0x65, 0xe9, 0xfc, 0x51,
	0xfe, 0xf7, 0x2d, 0x28, 0x75, 0xc4, 0x3b, 0x58, 0x9b, 0x99, 0x7a, 0xd0, 0xfd, 0x08, 0x73, 0xf1,
	0x08, 0x55, 0xb6, 0x6e, 0x82, 0x4c, 0xf3, 0x33, 0x19, 0xdd, 0x62, 0xf9, 0x99, 0x9d, 0x25, 0xcb,
	0x7b, 0x8b, 0x95, 0x93, 0xfc, 0x34, 0x00, 0xa6, 0x53, 0x69, 0xac, 0x0a, 0xe7, 0x46, 0xdd, 0x58,
	0x15, 0xce, 0x8f, 0xb2, 0xe8, 0x15, 0xac, 0x0b, 0x7f, 0x9b, 0xe4, 0x8a, 0xcb, 0xe1, 0x5a, 0x7e,
	0xf0, 0x76, 0x2c, 0x54, 0xd3, 0x21, 0xbd, 0x5c, 0x9a, 0x57, 0xc4, 0xed, 0xb4, 0x18, 0xb9, 0xf8,
	0x5b, 0x76, 0xbe, 0x83, 0xe2, 0xec, 0x6f, 0x08, 0x8a, 0x06, 0x77, 0xc9, 0x5f, 0x4d, 0xf9, 0xc1,
	0x8d, 0x98, 0x69, 0x6e, 0x67, 0x9e, 0xdb, 0x58, 0x6e, 0x17, 0x3f, 0xfc, 0xb1, 0xdc, 0x2e, 0x79,
	0xad, 0x5f, 0x54, 0x60, 0xc5, 0x63, 0xfd, 0x29, 0xf0, 0x38, 0xd1, 0x32, 0xbe, 0xcd, 0x0d, 0xb0,
	0xeb, 0x0c, 0x09, 0x0b, 0x7e, 0x37, 0x8c, 0xf3, 0x8c, 0xec, 0xb4, 0x8f, 0xff, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0x49, 0x1a, 0x55, 0x55, 0x32, 0x10, 0x00, 0x00,
}
