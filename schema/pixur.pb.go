// Code generated by protoc-gen-go.
// source: pixur.proto
// DO NOT EDIT!

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	pixur.proto

It has these top-level messages:
	Pic
	PicIdent
	AnimationInfo
	Tag
	PicTag
	User
	UserToken
*/
package schema

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

type Pic_Mime int32

const (
	Pic_UNKNOWN Pic_Mime = 0
	Pic_JPEG    Pic_Mime = 1
	Pic_GIF     Pic_Mime = 2
	Pic_PNG     Pic_Mime = 3
	Pic_WEBM    Pic_Mime = 4
)

var Pic_Mime_name = map[int32]string{
	0: "UNKNOWN",
	1: "JPEG",
	2: "GIF",
	3: "PNG",
	4: "WEBM",
}
var Pic_Mime_value = map[string]int32{
	"UNKNOWN": 0,
	"JPEG":    1,
	"GIF":     2,
	"PNG":     3,
	"WEBM":    4,
}

func (x Pic_Mime) String() string {
	return proto.EnumName(Pic_Mime_name, int32(x))
}
func (Pic_Mime) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Pic_DeletionStatus_Reason int32

const (
	// The reason is not know, due to limitations of proto
	Pic_DeletionStatus_UNKNOWN Pic_DeletionStatus_Reason = 0
	// No specific reason.  This is a catch-all reason.
	Pic_DeletionStatus_NONE Pic_DeletionStatus_Reason = 1
	// The pic is in violation of the rules.
	Pic_DeletionStatus_RULE_VIOLATION Pic_DeletionStatus_Reason = 2
)

var Pic_DeletionStatus_Reason_name = map[int32]string{
	0: "UNKNOWN",
	1: "NONE",
	2: "RULE_VIOLATION",
}
var Pic_DeletionStatus_Reason_value = map[string]int32{
	"UNKNOWN":        0,
	"NONE":           1,
	"RULE_VIOLATION": 2,
}

func (x Pic_DeletionStatus_Reason) String() string {
	return proto.EnumName(Pic_DeletionStatus_Reason_name, int32(x))
}
func (Pic_DeletionStatus_Reason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

type PicIdent_Type int32

const (
	PicIdent_UNKNOWN PicIdent_Type = 0
	PicIdent_SHA256  PicIdent_Type = 1
	PicIdent_SHA1    PicIdent_Type = 2
	PicIdent_MD5     PicIdent_Type = 3
	PicIdent_DCT_0   PicIdent_Type = 4
)

var PicIdent_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "SHA256",
	2: "SHA1",
	3: "MD5",
	4: "DCT_0",
}
var PicIdent_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"SHA256":  1,
	"SHA1":    2,
	"MD5":     3,
	"DCT_0":   4,
}

func (x PicIdent_Type) String() string {
	return proto.EnumName(PicIdent_Type_name, int32(x))
}
func (PicIdent_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type User_Capability int32

const (
	User_UNKNOWN User_Capability = 0
	// Can this user create and upload pictures
	User_PIC_CREATE User_Capability = 1
	// Can this user view the actual image data (grants pic token)
	User_PIC_READ User_Capability = 2
	// Can this user perform general pic index queries?
	User_PIC_INDEX User_Capability = 3
	// Can this user soft delete a pic?
	User_PIC_SOFT_DELETE User_Capability = 5
	// Can this user hard delete a pic?
	User_PIC_HARD_DELETE User_Capability = 6
	// Can this user purge a pic?
	User_PIC_PURGE User_Capability = 7
	// Can this user increment the pic view counter?
	User_PIC_UPDATE_VIEW_COUNTER User_Capability = 8
	// Can this user add tags and pic tags?
	User_PIC_TAG_CREATE User_Capability = 9
	// Can this user create other users?
	User_USER_CREATE User_Capability = 4
)

var User_Capability_name = map[int32]string{
	0: "UNKNOWN",
	1: "PIC_CREATE",
	2: "PIC_READ",
	3: "PIC_INDEX",
	5: "PIC_SOFT_DELETE",
	6: "PIC_HARD_DELETE",
	7: "PIC_PURGE",
	8: "PIC_UPDATE_VIEW_COUNTER",
	9: "PIC_TAG_CREATE",
	4: "USER_CREATE",
}
var User_Capability_value = map[string]int32{
	"UNKNOWN":                 0,
	"PIC_CREATE":              1,
	"PIC_READ":                2,
	"PIC_INDEX":               3,
	"PIC_SOFT_DELETE":         5,
	"PIC_HARD_DELETE":         6,
	"PIC_PURGE":               7,
	"PIC_UPDATE_VIEW_COUNTER": 8,
	"PIC_TAG_CREATE":          9,
	"USER_CREATE":             4,
}

func (x User_Capability) String() string {
	return proto.EnumName(User_Capability_name, int32(x))
}
func (User_Capability) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Pic struct {
	PicId      int64                       `protobuf:"varint,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
	FileSize   int64                       `protobuf:"varint,2,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	Mime       Pic_Mime                    `protobuf:"varint,3,opt,name=mime,enum=pixur.Pic_Mime" json:"mime,omitempty"`
	Width      int64                       `protobuf:"varint,4,opt,name=width" json:"width,omitempty"`
	Height     int64                       `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	CreatedTs  *google_protobuf1.Timestamp `protobuf:"bytes,10,opt,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	ModifiedTs *google_protobuf1.Timestamp `protobuf:"bytes,11,opt,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
	// If present, the pic is on the path to removal.  When the pic is marked
	// for deletion, it is delisted from normal indexing operations.  When the
	// pic is actually "deleted" only the pic object is removed.
	DeletionStatus *Pic_DeletionStatus `protobuf:"bytes,12,opt,name=deletion_status,json=deletionStatus" json:"deletion_status,omitempty"`
	// Only present on animated images.
	AnimationInfo *AnimationInfo    `protobuf:"bytes,13,opt,name=animation_info,json=animationInfo" json:"animation_info,omitempty"`
	ViewCount     int64             `protobuf:"varint,14,opt,name=view_count,json=viewCount" json:"view_count,omitempty"`
	Source        []*Pic_FileSource `protobuf:"bytes,15,rep,name=source" json:"source,omitempty"`
	FileName      []string          `protobuf:"bytes,16,rep,name=file_name,json=fileName" json:"file_name,omitempty"`
}

func (m *Pic) Reset()                    { *m = Pic{} }
func (m *Pic) String() string            { return proto.CompactTextString(m) }
func (*Pic) ProtoMessage()               {}
func (*Pic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pic) GetCreatedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedTs
	}
	return nil
}

func (m *Pic) GetModifiedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ModifiedTs
	}
	return nil
}

func (m *Pic) GetDeletionStatus() *Pic_DeletionStatus {
	if m != nil {
		return m.DeletionStatus
	}
	return nil
}

func (m *Pic) GetAnimationInfo() *AnimationInfo {
	if m != nil {
		return m.AnimationInfo
	}
	return nil
}

func (m *Pic) GetSource() []*Pic_FileSource {
	if m != nil {
		return m.Source
	}
	return nil
}

type Pic_DeletionStatus struct {
	// Represents when this Pic was marked for deletion
	MarkedDeletedTs *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=marked_deleted_ts,json=markedDeletedTs" json:"marked_deleted_ts,omitempty"`
	// Represents when this picture will be auto deleted.  Note that the Pic
	// may exist for a short period after this time.  (may be absent)
	PendingDeletedTs *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=pending_deleted_ts,json=pendingDeletedTs" json:"pending_deleted_ts,omitempty"`
	// Determines when Pic was actually deleted.  (present after the Pic is
	// hard deleted, a.k.a purging)
	ActualDeletedTs *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=actual_deleted_ts,json=actualDeletedTs" json:"actual_deleted_ts,omitempty"`
	// Gives an explanation for why this pic was removed.
	Details string `protobuf:"bytes,4,opt,name=details" json:"details,omitempty"`
	// The reason the pic was removed.
	Reason Pic_DeletionStatus_Reason `protobuf:"varint,5,opt,name=reason,enum=pixur.Pic_DeletionStatus_Reason" json:"reason,omitempty"`
	// Determines if this pic can be undeleted if re uploaded.  Currently the
	// only reason is due to disk space concerns.
	Temporary bool `protobuf:"varint,6,opt,name=temporary" json:"temporary,omitempty"`
}

func (m *Pic_DeletionStatus) Reset()                    { *m = Pic_DeletionStatus{} }
func (m *Pic_DeletionStatus) String() string            { return proto.CompactTextString(m) }
func (*Pic_DeletionStatus) ProtoMessage()               {}
func (*Pic_DeletionStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Pic_DeletionStatus) GetMarkedDeletedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.MarkedDeletedTs
	}
	return nil
}

func (m *Pic_DeletionStatus) GetPendingDeletedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.PendingDeletedTs
	}
	return nil
}

func (m *Pic_DeletionStatus) GetActualDeletedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ActualDeletedTs
	}
	return nil
}

type Pic_FileSource struct {
	Url       string                      `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Referrer  string                      `protobuf:"bytes,2,opt,name=referrer" json:"referrer,omitempty"`
	CreatedTs *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
}

func (m *Pic_FileSource) Reset()                    { *m = Pic_FileSource{} }
func (m *Pic_FileSource) String() string            { return proto.CompactTextString(m) }
func (*Pic_FileSource) ProtoMessage()               {}
func (*Pic_FileSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *Pic_FileSource) GetCreatedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedTs
	}
	return nil
}

// A picture identifier
type PicIdent struct {
	PicId int64         `protobuf:"varint,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
	Type  PicIdent_Type `protobuf:"varint,2,opt,name=type,enum=pixur.PicIdent_Type" json:"type,omitempty"`
	Value []byte        `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// dct0 are the upper 8x8 corner of the 32x32 dct of the image
	Dct0Values []float32 `protobuf:"fixed32,4,rep,packed,name=dct0_values,json=dct0Values" json:"dct0_values,omitempty"`
}

func (m *PicIdent) Reset()                    { *m = PicIdent{} }
func (m *PicIdent) String() string            { return proto.CompactTextString(m) }
func (*PicIdent) ProtoMessage()               {}
func (*PicIdent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AnimationInfo struct {
	// How long this animated image in time.  There must be more than 1 frame
	// for this value to be set.
	Duration *google_protobuf.Duration `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
}

func (m *AnimationInfo) Reset()                    { *m = AnimationInfo{} }
func (m *AnimationInfo) String() string            { return proto.CompactTextString(m) }
func (*AnimationInfo) ProtoMessage()               {}
func (*AnimationInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AnimationInfo) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type Tag struct {
	TagId      int64                       `protobuf:"varint,1,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
	Name       string                      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	UsageCount int64                       `protobuf:"varint,3,opt,name=usage_count,json=usageCount" json:"usage_count,omitempty"`
	CreatedTs  *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	ModifiedTs *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Tag) GetCreatedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedTs
	}
	return nil
}

func (m *Tag) GetModifiedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ModifiedTs
	}
	return nil
}

type PicTag struct {
	PicId int64  `protobuf:"varint,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
	TagId int64  `protobuf:"varint,2,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// The user who originally created this tag.  optional.
	UserId     int64                       `protobuf:"varint,8,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	CreatedTs  *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	ModifiedTs *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
}

func (m *PicTag) Reset()                    { *m = PicTag{} }
func (m *PicTag) String() string            { return proto.CompactTextString(m) }
func (*PicTag) ProtoMessage()               {}
func (*PicTag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PicTag) GetCreatedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedTs
	}
	return nil
}

func (m *PicTag) GetModifiedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ModifiedTs
	}
	return nil
}

type User struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Hashed secret token
	Secret []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	// Identity of the user, usually an email.
	Ident      string                      `protobuf:"bytes,3,opt,name=ident" json:"ident,omitempty"`
	CreatedTs  *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	ModifiedTs *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
	LastSeenTs *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=last_seen_ts,json=lastSeenTs" json:"last_seen_ts,omitempty"`
	Capability []User_Capability           `protobuf:"varint,7,rep,name=capability,enum=pixur.User_Capability" json:"capability,omitempty"`
	// Always increment-then-get
	NextTokenId int64        `protobuf:"varint,8,opt,name=next_token_id,json=nextTokenId" json:"next_token_id,omitempty"`
	UserToken   []*UserToken `protobuf:"bytes,9,rep,name=user_token,json=userToken" json:"user_token,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *User) GetCreatedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedTs
	}
	return nil
}

func (m *User) GetModifiedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ModifiedTs
	}
	return nil
}

func (m *User) GetLastSeenTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastSeenTs
	}
	return nil
}

func (m *User) GetUserToken() []*UserToken {
	if m != nil {
		return m.UserToken
	}
	return nil
}

// Represent the valid auth tokens.  When a user logs out, these will be
// deleted.
type UserToken struct {
	TokenId    int64                       `protobuf:"varint,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	CreatedTs  *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	LastSeenTs *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=last_seen_ts,json=lastSeenTs" json:"last_seen_ts,omitempty"`
}

func (m *UserToken) Reset()                    { *m = UserToken{} }
func (m *UserToken) String() string            { return proto.CompactTextString(m) }
func (*UserToken) ProtoMessage()               {}
func (*UserToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserToken) GetCreatedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedTs
	}
	return nil
}

func (m *UserToken) GetLastSeenTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastSeenTs
	}
	return nil
}

func init() {
	proto.RegisterType((*Pic)(nil), "pixur.Pic")
	proto.RegisterType((*Pic_DeletionStatus)(nil), "pixur.Pic.DeletionStatus")
	proto.RegisterType((*Pic_FileSource)(nil), "pixur.Pic.FileSource")
	proto.RegisterType((*PicIdent)(nil), "pixur.PicIdent")
	proto.RegisterType((*AnimationInfo)(nil), "pixur.AnimationInfo")
	proto.RegisterType((*Tag)(nil), "pixur.Tag")
	proto.RegisterType((*PicTag)(nil), "pixur.PicTag")
	proto.RegisterType((*User)(nil), "pixur.User")
	proto.RegisterType((*UserToken)(nil), "pixur.UserToken")
	proto.RegisterEnum("pixur.Pic_Mime", Pic_Mime_name, Pic_Mime_value)
	proto.RegisterEnum("pixur.Pic_DeletionStatus_Reason", Pic_DeletionStatus_Reason_name, Pic_DeletionStatus_Reason_value)
	proto.RegisterEnum("pixur.PicIdent_Type", PicIdent_Type_name, PicIdent_Type_value)
	proto.RegisterEnum("pixur.User_Capability", User_Capability_name, User_Capability_value)
}

var fileDescriptor0 = []byte{
	// 1108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0xcf, 0x92, 0xda, 0xc6,
	0x13, 0xfe, 0x09, 0x09, 0x21, 0x35, 0xbb, 0x20, 0xcf, 0xcf, 0x7f, 0x64, 0x9c, 0xc4, 0x14, 0xbe,
	0x70, 0x09, 0x76, 0x70, 0xd9, 0x15, 0x97, 0x93, 0x03, 0x0b, 0x62, 0x97, 0xc4, 0x06, 0x6a, 0x10,
	0xde, 0x54, 0x2e, 0x2a, 0xad, 0x34, 0xb0, 0x53, 0x46, 0x12, 0x25, 0x8d, 0x6c, 0xaf, 0xdf, 0x22,
	0x4f, 0x90, 0xb7, 0xc8, 0x3d, 0x87, 0x1c, 0x72, 0xcf, 0x2d, 0x2f, 0x93, 0x9a, 0x91, 0xc4, 0x9f,
	0xb5, 0x5d, 0x9b, 0x2d, 0x1f, 0x72, 0xd3, 0xd7, 0xfd, 0x75, 0xd3, 0x5f, 0xf7, 0x74, 0x03, 0xd5,
	0x35, 0x7d, 0x97, 0xc6, 0x9d, 0x75, 0x1c, 0xb1, 0x08, 0x95, 0x05, 0x68, 0x7c, 0xb5, 0x8c, 0xa2,
	0xe5, 0x8a, 0x3c, 0x14, 0xc6, 0xb3, 0x74, 0xf1, 0xd0, 0x4f, 0x63, 0x97, 0xd1, 0x28, 0xcc, 0x68,
	0x8d, 0xfb, 0x97, 0xfd, 0x8c, 0x06, 0x24, 0x61, 0x6e, 0xb0, 0xce, 0x08, 0xad, 0xdf, 0x34, 0x90,
	0xa7, 0xd4, 0x43, 0xb7, 0x40, 0x5d, 0x53, 0xcf, 0xa1, 0xbe, 0x29, 0x35, 0xa5, 0xb6, 0x8c, 0xcb,
	0x6b, 0xea, 0x8d, 0x7c, 0x74, 0x0f, 0xf4, 0x05, 0x5d, 0x11, 0x27, 0xa1, 0xef, 0x89, 0x59, 0x12,
	0x1e, 0x8d, 0x1b, 0x66, 0xf4, 0x3d, 0x41, 0x0f, 0x40, 0x09, 0x68, 0x40, 0x4c, 0xb9, 0x29, 0xb5,
	0x6b, 0xdd, 0x7a, 0x27, 0xab, 0x6f, 0x4a, 0xbd, 0xce, 0x4b, 0x1a, 0x10, 0x2c, 0x9c, 0xe8, 0x26,
	0x94, 0xdf, 0x52, 0x9f, 0x9d, 0x9b, 0x4a, 0x96, 0x57, 0x00, 0x74, 0x1b, 0xd4, 0x73, 0x42, 0x97,
	0xe7, 0xcc, 0x2c, 0x0b, 0x73, 0x8e, 0xd0, 0x33, 0x00, 0x2f, 0x26, 0x2e, 0x23, 0xbe, 0xc3, 0x12,
	0x13, 0x9a, 0x52, 0xbb, 0xda, 0x6d, 0x74, 0x32, 0x11, 0x9d, 0x42, 0x44, 0xc7, 0x2e, 0x44, 0x60,
	0x3d, 0x67, 0xdb, 0x09, 0x7a, 0x0e, 0xd5, 0x20, 0xf2, 0xe9, 0x82, 0x66, 0xb1, 0xd5, 0x2b, 0x63,
	0xa1, 0xa0, 0xdb, 0x09, 0x3a, 0x82, 0xba, 0x4f, 0x56, 0x84, 0x77, 0xce, 0x49, 0x98, 0xcb, 0xd2,
	0xc4, 0x3c, 0x10, 0x09, 0xee, 0xee, 0xa8, 0x1a, 0xe4, 0x8c, 0x99, 0x20, 0xe0, 0x9a, 0xbf, 0x87,
	0xd1, 0x73, 0xa8, 0xb9, 0x21, 0x0d, 0x44, 0xfb, 0x1d, 0x1a, 0x2e, 0x22, 0xf3, 0x50, 0xa4, 0xb8,
	0x99, 0xa7, 0xe8, 0x15, 0xce, 0x51, 0xb8, 0x88, 0xf0, 0xa1, 0xbb, 0x0b, 0xd1, 0x97, 0x00, 0x6f,
	0x28, 0x79, 0xeb, 0x78, 0x51, 0x1a, 0x32, 0xb3, 0x26, 0x9a, 0xa2, 0x73, 0x4b, 0x9f, 0x1b, 0xd0,
	0xd7, 0xa0, 0x26, 0x51, 0x1a, 0x7b, 0xc4, 0xac, 0x37, 0xe5, 0x76, 0xb5, 0x7b, 0x6b, 0xa7, 0xac,
	0x21, 0x9f, 0x87, 0x70, 0xe2, 0x9c, 0xb4, 0x19, 0x5b, 0xe8, 0x06, 0xc4, 0x34, 0x9a, 0x72, 0x5b,
	0xcf, 0xc6, 0x36, 0x76, 0x03, 0xd2, 0xf8, 0x45, 0x86, 0xda, 0xbe, 0x14, 0x34, 0x84, 0x1b, 0x81,
	0x1b, 0xbf, 0x26, 0xbe, 0x23, 0x34, 0x65, 0x1d, 0x94, 0xae, 0xec, 0x60, 0x3d, 0x0b, 0x1a, 0x64,
	0x31, 0x76, 0x82, 0x4e, 0x00, 0xad, 0x49, 0xe8, 0xd3, 0x70, 0xb9, 0x9b, 0xa8, 0x74, 0x65, 0x22,
	0x23, 0x8f, 0xda, 0x66, 0x1a, 0xc2, 0x0d, 0xd7, 0x63, 0xa9, 0xbb, 0xda, 0x4d, 0x24, 0x5f, 0x5d,
	0x51, 0x16, 0xb4, 0xcd, 0x63, 0x42, 0xc5, 0x27, 0xcc, 0xa5, 0xab, 0x44, 0x3c, 0x40, 0x1d, 0x17,
	0x10, 0x7d, 0x0b, 0x6a, 0x4c, 0xdc, 0x24, 0x0a, 0xc5, 0x13, 0xac, 0x75, 0x9b, 0x9f, 0x9c, 0x74,
	0x07, 0x0b, 0x1e, 0xce, 0xf9, 0xe8, 0x0b, 0xd0, 0x19, 0x09, 0xd6, 0x51, 0xec, 0xc6, 0x17, 0xa6,
	0xda, 0x94, 0xda, 0x1a, 0xde, 0x1a, 0x5a, 0x8f, 0x41, 0xcd, 0xf8, 0xa8, 0x0a, 0x95, 0xf9, 0xf8,
	0xc7, 0xf1, 0xe4, 0x74, 0x6c, 0xfc, 0x0f, 0x69, 0xa0, 0x8c, 0x27, 0x63, 0xcb, 0x90, 0x10, 0x82,
	0x1a, 0x9e, 0xbf, 0xb0, 0x9c, 0x57, 0xa3, 0xc9, 0x8b, 0x9e, 0x3d, 0x9a, 0x8c, 0x8d, 0x52, 0x23,
	0x05, 0xd8, 0x8e, 0x11, 0x19, 0x20, 0xa7, 0xf1, 0x4a, 0x0c, 0x40, 0xc7, 0xfc, 0x13, 0x35, 0x40,
	0x8b, 0xc9, 0x82, 0xc4, 0x31, 0x89, 0x45, 0x3b, 0x75, 0xbc, 0xc1, 0x97, 0x76, 0x46, 0xbe, 0xc6,
	0xce, 0xb4, 0x9e, 0x81, 0xc2, 0x57, 0xf5, 0x83, 0x4a, 0x7f, 0x98, 0x5a, 0xc7, 0x86, 0x84, 0x2a,
	0x20, 0x1f, 0x8f, 0x86, 0x46, 0x89, 0x7f, 0x4c, 0xc7, 0xc7, 0x86, 0xcc, 0x7d, 0xa7, 0xd6, 0xd1,
	0x4b, 0x43, 0x69, 0xfd, 0x21, 0x81, 0x36, 0xe5, 0x37, 0x82, 0x84, 0xec, 0x53, 0xd7, 0xa3, 0x0d,
	0x0a, 0xbb, 0x58, 0x67, 0x87, 0xa3, 0xb6, 0xd9, 0x83, 0x22, 0xaa, 0x63, 0x5f, 0xac, 0x09, 0x16,
	0x0c, 0x7e, 0x25, 0xde, 0xb8, 0xab, 0x34, 0xbb, 0x25, 0x07, 0x38, 0x03, 0xe8, 0x01, 0x54, 0x7d,
	0x8f, 0x3d, 0x72, 0x04, 0xe2, 0x03, 0x94, 0xdb, 0xa5, 0xa3, 0x92, 0x21, 0x61, 0xe0, 0xe6, 0x57,
	0xc2, 0xda, 0xfa, 0x1e, 0x14, 0x9e, 0x68, 0x5f, 0x03, 0x80, 0x3a, 0x3b, 0xe9, 0x75, 0x9f, 0x3c,
	0x35, 0x24, 0x5e, 0xf3, 0xec, 0xa4, 0xf7, 0x4d, 0x26, 0xe3, 0xe5, 0xe0, 0x89, 0x21, 0x23, 0x1d,
	0xca, 0x83, 0xbe, 0xed, 0x3c, 0x32, 0x94, 0xd6, 0x10, 0x0e, 0xf7, 0x16, 0x13, 0x3d, 0x01, 0xad,
	0x38, 0xa2, 0xf9, 0x0a, 0xdc, 0xfd, 0xa0, 0x99, 0x83, 0x9c, 0x80, 0x37, 0xd4, 0xd6, 0x9f, 0x12,
	0xc8, 0xb6, 0xbb, 0xe4, 0xad, 0x60, 0xee, 0x72, 0xa7, 0x15, 0xcc, 0x5d, 0x8e, 0x7c, 0x84, 0x40,
	0x11, 0xcb, 0x98, 0x0d, 0x4f, 0x7c, 0xa3, 0xfb, 0x50, 0x4d, 0x13, 0x77, 0x49, 0xf2, 0xa5, 0x97,
	0x05, 0x1f, 0x84, 0x29, 0xdb, 0xfa, 0xfd, 0xc9, 0xaa, 0x9f, 0x71, 0x0d, 0x2b, 0xd7, 0xb9, 0x86,
	0xad, 0xbf, 0x25, 0x50, 0xa7, 0xd4, 0xcb, 0xe5, 0x7c, 0x6c, 0xb2, 0x5b, 0x95, 0xa5, 0x8f, 0xa9,
	0x94, 0x77, 0x54, 0xde, 0x81, 0x4a, 0x9a, 0x90, 0x98, 0x73, 0xb5, 0xec, 0xd6, 0x73, 0x38, 0xf2,
	0xff, 0x33, 0x75, 0x7f, 0x29, 0xa0, 0xcc, 0x13, 0x12, 0xef, 0x56, 0x26, 0xed, 0x55, 0x76, 0x1b,
	0xd4, 0x84, 0x78, 0x31, 0x61, 0x42, 0xdd, 0x01, 0xce, 0x11, 0x7f, 0xa5, 0x94, 0xbf, 0xdc, 0x5c,
	0x5f, 0x06, 0x2e, 0xe9, 0x50, 0x3e, 0x43, 0x47, 0xf9, 0x5a, 0xff, 0x59, 0xdf, 0xc1, 0xc1, 0xca,
	0x4d, 0x98, 0x93, 0x10, 0x12, 0xfe, 0xbb, 0x0e, 0x02, 0xe7, 0xcf, 0x08, 0x09, 0xed, 0x04, 0x3d,
	0x05, 0xf0, 0xdc, 0xb5, 0x7b, 0x46, 0x57, 0x94, 0x5d, 0x98, 0x95, 0xa6, 0xdc, 0xae, 0x75, 0x6f,
	0xe7, 0x1b, 0xca, 0xbb, 0xd3, 0xe9, 0x6f, 0xbc, 0x78, 0x87, 0x89, 0x5a, 0x70, 0x18, 0x92, 0x77,
	0xcc, 0x61, 0xd1, 0x6b, 0x12, 0x6e, 0x87, 0x5a, 0xe5, 0x46, 0x9b, 0xdb, 0x46, 0x3e, 0x7a, 0x08,
	0x20, 0x1a, 0x2b, 0x38, 0xa6, 0x2e, 0xfe, 0xb1, 0x8c, 0x9d, 0xdc, 0x82, 0x87, 0xf5, 0xb4, 0xf8,
	0x6c, 0xfd, 0x2e, 0x01, 0x6c, 0x7f, 0x6f, 0x7f, 0x95, 0x6b, 0x00, 0xd3, 0x51, 0xdf, 0xe9, 0x63,
	0xab, 0x67, 0xf3, 0xf3, 0x79, 0x00, 0x1a, 0xc7, 0xd8, 0xea, 0x0d, 0x8c, 0x12, 0x3a, 0x04, 0x9d,
	0xa3, 0xd1, 0x78, 0x60, 0xfd, 0x64, 0xc8, 0xe8, 0xff, 0x50, 0xe7, 0x70, 0x36, 0x19, 0xda, 0xce,
	0xc0, 0x7a, 0x61, 0xd9, 0x96, 0x51, 0x2e, 0x8c, 0x27, 0x3d, 0x3c, 0x28, 0x8c, 0x6a, 0x11, 0x38,
	0x9d, 0xe3, 0x63, 0xcb, 0xa8, 0xa0, 0x7b, 0x70, 0x87, 0xc3, 0xf9, 0x74, 0xd0, 0xb3, 0xf9, 0x69,
	0xb6, 0x4e, 0x9d, 0xfe, 0x64, 0x3e, 0xb6, 0x2d, 0x6c, 0x68, 0xfc, 0x62, 0x73, 0xa7, 0xdd, 0x3b,
	0x2e, 0xca, 0xd0, 0x51, 0x1d, 0xaa, 0xf3, 0x99, 0x85, 0x0b, 0x83, 0xd2, 0xfa, 0x55, 0x02, 0x7d,
	0x23, 0x0e, 0xdd, 0x05, 0x6d, 0xd3, 0xa1, 0xec, 0x71, 0x55, 0x58, 0xde, 0x9d, 0xfd, 0xf7, 0x52,
	0xba, 0xce, 0x7b, 0xb9, 0x3c, 0x72, 0xf9, 0x3a, 0x23, 0x3f, 0xd2, 0x7e, 0x56, 0x13, 0xef, 0x9c,
	0x04, 0xee, 0x99, 0x2a, 0x98, 0x8f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xf0, 0xff, 0xb6,
	0x53, 0x0a, 0x00, 0x00,
}
