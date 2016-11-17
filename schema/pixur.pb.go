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
	PicComment
	PicVote
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

type PicVote_Vote int32

const (
	PicVote_UNKNOWN PicVote_Vote = 0
	PicVote_UP      PicVote_Vote = 1
	PicVote_DOWN    PicVote_Vote = 2
	PicVote_NEUTRAL PicVote_Vote = 3
)

var PicVote_Vote_name = map[int32]string{
	0: "UNKNOWN",
	1: "UP",
	2: "DOWN",
	3: "NEUTRAL",
}
var PicVote_Vote_value = map[string]int32{
	"UNKNOWN": 0,
	"UP":      1,
	"DOWN":    2,
	"NEUTRAL": 3,
}

func (x PicVote_Vote) String() string {
	return proto.EnumName(PicVote_Vote_name, int32(x))
}
func (PicVote_Vote) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

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
	// Can this user create comments?
	User_PIC_COMMENT_CREATE User_Capability = 10
	// Can this user vote?
	User_PIC_VOTE_CREATE User_Capability = 11
	// Can this user create other users?
	User_USER_CREATE User_Capability = 4
)

var User_Capability_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "PIC_CREATE",
	2:  "PIC_READ",
	3:  "PIC_INDEX",
	5:  "PIC_SOFT_DELETE",
	6:  "PIC_HARD_DELETE",
	7:  "PIC_PURGE",
	8:  "PIC_UPDATE_VIEW_COUNTER",
	9:  "PIC_TAG_CREATE",
	10: "PIC_COMMENT_CREATE",
	11: "PIC_VOTE_CREATE",
	4:  "USER_CREATE",
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
	"PIC_COMMENT_CREATE":      10,
	"PIC_VOTE_CREATE":         11,
	"USER_CREATE":             4,
}

func (x User_Capability) String() string {
	return proto.EnumName(User_Capability_name, int32(x))
}
func (User_Capability) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

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
	// The denormalized sum of upvotes for this pic
	VoteUp int64 `protobuf:"varint,17,opt,name=vote_up,json=voteUp" json:"vote_up,omitempty"`
	// The denormalized sum of downvotes for this pic
	VoteDown int64 `protobuf:"varint,18,opt,name=vote_down,json=voteDown" json:"vote_down,omitempty"`
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

type PicComment struct {
	PicId     int64 `protobuf:"varint,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
	CommentId int64 `protobuf:"varint,2,opt,name=comment_id,json=commentId" json:"comment_id,omitempty"`
	// parent id of this comment.  0 if root.
	CommentParentId int64 `protobuf:"varint,3,opt,name=comment_parent_id,json=commentParentId" json:"comment_parent_id,omitempty"`
	// author
	UserId     int64                       `protobuf:"varint,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Text       string                      `protobuf:"bytes,7,opt,name=text" json:"text,omitempty"`
	CreatedTs  *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	ModifiedTs *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
}

func (m *PicComment) Reset()                    { *m = PicComment{} }
func (m *PicComment) String() string            { return proto.CompactTextString(m) }
func (*PicComment) ProtoMessage()               {}
func (*PicComment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PicComment) GetCreatedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedTs
	}
	return nil
}

func (m *PicComment) GetModifiedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ModifiedTs
	}
	return nil
}

type PicVote struct {
	PicId      int64                       `protobuf:"varint,1,opt,name=pic_id,json=picId" json:"pic_id,omitempty"`
	UserId     int64                       `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Vote       PicVote_Vote                `protobuf:"varint,3,opt,name=vote,enum=pixur.PicVote_Vote" json:"vote,omitempty"`
	CreatedTs  *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	ModifiedTs *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
}

func (m *PicVote) Reset()                    { *m = PicVote{} }
func (m *PicVote) String() string            { return proto.CompactTextString(m) }
func (*PicVote) ProtoMessage()               {}
func (*PicVote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PicVote) GetCreatedTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedTs
	}
	return nil
}

func (m *PicVote) GetModifiedTs() *google_protobuf1.Timestamp {
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
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (*UserToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
	proto.RegisterType((*PicComment)(nil), "pixur.PicComment")
	proto.RegisterType((*PicVote)(nil), "pixur.PicVote")
	proto.RegisterType((*User)(nil), "pixur.User")
	proto.RegisterType((*UserToken)(nil), "pixur.UserToken")
	proto.RegisterEnum("pixur.Pic_Mime", Pic_Mime_name, Pic_Mime_value)
	proto.RegisterEnum("pixur.Pic_DeletionStatus_Reason", Pic_DeletionStatus_Reason_name, Pic_DeletionStatus_Reason_value)
	proto.RegisterEnum("pixur.PicIdent_Type", PicIdent_Type_name, PicIdent_Type_value)
	proto.RegisterEnum("pixur.PicVote_Vote", PicVote_Vote_name, PicVote_Vote_value)
	proto.RegisterEnum("pixur.User_Capability", User_Capability_name, User_Capability_value)
}

var fileDescriptor0 = []byte{
	// 1295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x17, 0xfd, 0xf8, 0x23, 0x4a, 0xbc, 0xb2, 0x65, 0x7a, 0x92, 0x38, 0x8c, 0xf3, 0xa5, 0x11, 0x94,
	0x45, 0x85, 0x02, 0x55, 0x52, 0x05, 0x09, 0x1a, 0xa4, 0x5d, 0xc8, 0x12, 0x6d, 0xab, 0xb5, 0x25,
	0x61, 0x44, 0x39, 0x45, 0x37, 0x04, 0x43, 0x8e, 0xe5, 0x41, 0x44, 0x52, 0x20, 0x87, 0x71, 0x9c,
	0x5d, 0x5f, 0xa0, 0x40, 0xd1, 0x07, 0xe8, 0xd3, 0x74, 0xd1, 0x67, 0xe8, 0x2b, 0x14, 0x7d, 0x86,
	0x62, 0x86, 0xa4, 0x7e, 0x9c, 0xa4, 0xae, 0x91, 0x45, 0x36, 0xc6, 0xdc, 0x73, 0xcf, 0x5c, 0xcf,
	0x39, 0x77, 0xe6, 0x52, 0x50, 0x9d, 0xd3, 0x37, 0x69, 0xdc, 0x9a, 0xc7, 0x11, 0x8b, 0x50, 0x49,
	0x04, 0xbb, 0x9f, 0x4d, 0xa3, 0x68, 0x3a, 0x23, 0x0f, 0x05, 0xf8, 0x32, 0x3d, 0x7d, 0xe8, 0xa7,
	0xb1, 0xcb, 0x68, 0x14, 0x66, 0xb4, 0xdd, 0xfb, 0x97, 0xf3, 0x8c, 0x06, 0x24, 0x61, 0x6e, 0x30,
	0xcf, 0x08, 0x8d, 0xbf, 0x2a, 0xa0, 0x8c, 0xa8, 0x87, 0x6e, 0x81, 0x36, 0xa7, 0x9e, 0x43, 0x7d,
	0x53, 0xaa, 0x4b, 0x4d, 0x05, 0x97, 0xe6, 0xd4, 0xeb, 0xfb, 0xe8, 0x2e, 0xe8, 0xa7, 0x74, 0x46,
	0x9c, 0x84, 0xbe, 0x25, 0xa6, 0x2c, 0x32, 0x15, 0x0e, 0x8c, 0xe9, 0x5b, 0x82, 0x1e, 0x80, 0x1a,
	0xd0, 0x80, 0x98, 0x4a, 0x5d, 0x6a, 0xd6, 0xda, 0x5b, 0xad, 0xec, 0x7c, 0x23, 0xea, 0xb5, 0x8e,
	0x69, 0x40, 0xb0, 0x48, 0xa2, 0x9b, 0x50, 0x3a, 0xa7, 0x3e, 0x3b, 0x33, 0xd5, 0xac, 0xae, 0x08,
	0xd0, 0x0e, 0x68, 0x67, 0x84, 0x4e, 0xcf, 0x98, 0x59, 0x12, 0x70, 0x1e, 0xa1, 0x67, 0x00, 0x5e,
	0x4c, 0x5c, 0x46, 0x7c, 0x87, 0x25, 0x26, 0xd4, 0xa5, 0x66, 0xb5, 0xbd, 0xdb, 0xca, 0x44, 0xb4,
	0x0a, 0x11, 0x2d, 0xbb, 0x10, 0x81, 0xf5, 0x9c, 0x6d, 0x27, 0xe8, 0x39, 0x54, 0x83, 0xc8, 0xa7,
	0xa7, 0x34, 0xdb, 0x5b, 0xbd, 0x72, 0x2f, 0x14, 0x74, 0x3b, 0x41, 0x7b, 0xb0, 0xe5, 0x93, 0x19,
	0xe1, 0xce, 0x39, 0x09, 0x73, 0x59, 0x9a, 0x98, 0x1b, 0xa2, 0xc0, 0x9d, 0x15, 0x55, 0xbd, 0x9c,
	0x31, 0x16, 0x04, 0x5c, 0xf3, 0xd7, 0x62, 0xf4, 0x1c, 0x6a, 0x6e, 0x48, 0x03, 0x61, 0xbf, 0x43,
	0xc3, 0xd3, 0xc8, 0xdc, 0x14, 0x25, 0x6e, 0xe6, 0x25, 0x3a, 0x45, 0xb2, 0x1f, 0x9e, 0x46, 0x78,
	0xd3, 0x5d, 0x0d, 0xd1, 0x3d, 0x80, 0xd7, 0x94, 0x9c, 0x3b, 0x5e, 0x94, 0x86, 0xcc, 0xac, 0x09,
	0x53, 0x74, 0x8e, 0x74, 0x39, 0x80, 0xbe, 0x04, 0x2d, 0x89, 0xd2, 0xd8, 0x23, 0xe6, 0x56, 0x5d,
	0x69, 0x56, 0xdb, 0xb7, 0x56, 0x8e, 0xb5, 0xcf, 0xfb, 0x21, 0x92, 0x38, 0x27, 0x2d, 0xda, 0x16,
	0xba, 0x01, 0x31, 0x8d, 0xba, 0xd2, 0xd4, 0xb3, 0xb6, 0x0d, 0xdc, 0x80, 0xa0, 0xdb, 0x50, 0x7e,
	0x1d, 0x31, 0xe2, 0xa4, 0x73, 0x73, 0x3b, 0x33, 0x9f, 0x87, 0x93, 0x39, 0xdf, 0x25, 0x12, 0x7e,
	0x74, 0x1e, 0x9a, 0x28, 0x6b, 0x36, 0x07, 0x7a, 0xd1, 0x79, 0xb8, 0xfb, 0x8b, 0x02, 0xb5, 0x75,
	0x03, 0xd0, 0x3e, 0x6c, 0x07, 0x6e, 0xfc, 0x8a, 0xf8, 0x8e, 0x70, 0x22, 0xf3, 0x5d, 0xba, 0xd2,
	0xf7, 0xad, 0x6c, 0x53, 0x2f, 0xdb, 0x63, 0x27, 0xe8, 0x10, 0xd0, 0x9c, 0x84, 0x3e, 0x0d, 0xa7,
	0xab, 0x85, 0xe4, 0x2b, 0x0b, 0x19, 0xf9, 0xae, 0x65, 0xa5, 0x7d, 0xd8, 0x76, 0x3d, 0x96, 0xba,
	0xb3, 0xd5, 0x42, 0xca, 0xd5, 0x27, 0xca, 0x36, 0x2d, 0xeb, 0x98, 0x50, 0xf6, 0x09, 0x73, 0xe9,
	0x2c, 0x11, 0xd7, 0x56, 0xc7, 0x45, 0x88, 0xbe, 0x06, 0x2d, 0x26, 0x6e, 0x12, 0x85, 0xe2, 0xe2,
	0xd6, 0xda, 0xf5, 0x0f, 0xde, 0x8f, 0x16, 0x16, 0x3c, 0x9c, 0xf3, 0xd1, 0xff, 0x41, 0x67, 0x24,
	0x98, 0x47, 0xb1, 0x1b, 0x5f, 0x98, 0x5a, 0x5d, 0x6a, 0x56, 0xf0, 0x12, 0x68, 0x3c, 0x06, 0x2d,
	0xe3, 0xa3, 0x2a, 0x94, 0x27, 0x83, 0xef, 0x07, 0xc3, 0x17, 0x03, 0xe3, 0x7f, 0xa8, 0x02, 0xea,
	0x60, 0x38, 0xb0, 0x0c, 0x09, 0x21, 0xa8, 0xe1, 0xc9, 0x91, 0xe5, 0x9c, 0xf4, 0x87, 0x47, 0x1d,
	0xbb, 0x3f, 0x1c, 0x18, 0xf2, 0x6e, 0x0a, 0xb0, 0x6c, 0x3e, 0x32, 0x40, 0x49, 0xe3, 0x99, 0x68,
	0x80, 0x8e, 0xf9, 0x12, 0xed, 0x42, 0x25, 0x26, 0xa7, 0x24, 0x8e, 0x49, 0x2c, 0xec, 0xd4, 0xf1,
	0x22, 0xbe, 0xf4, 0xd2, 0x94, 0x6b, 0xbc, 0xb4, 0xc6, 0x33, 0x50, 0xf9, 0x03, 0x7f, 0xe7, 0xa4,
	0xdf, 0x8d, 0xac, 0x03, 0x43, 0x42, 0x65, 0x50, 0x0e, 0xfa, 0xfb, 0x86, 0xcc, 0x17, 0xa3, 0xc1,
	0x81, 0xa1, 0xf0, 0xdc, 0x0b, 0x6b, 0xef, 0xd8, 0x50, 0x1b, 0xbf, 0x4b, 0x50, 0x19, 0xf1, 0xc9,
	0x42, 0x42, 0xf6, 0xa1, 0x99, 0xd3, 0x04, 0x95, 0x5d, 0xcc, 0xb3, 0x71, 0x53, 0x5b, 0xbc, 0x9e,
	0x62, 0x57, 0xcb, 0xbe, 0x98, 0x13, 0x2c, 0x18, 0x7c, 0xb6, 0xbc, 0x76, 0x67, 0x69, 0x36, 0x81,
	0x36, 0x70, 0x16, 0xa0, 0x07, 0x50, 0xf5, 0x3d, 0xf6, 0xc8, 0x11, 0x11, 0x6f, 0xa0, 0xd2, 0x94,
	0xf7, 0x64, 0x43, 0xc2, 0xc0, 0xe1, 0x13, 0x81, 0x36, 0xbe, 0x05, 0x95, 0x17, 0x5a, 0xd7, 0x00,
	0xa0, 0x8d, 0x0f, 0x3b, 0xed, 0x27, 0x4f, 0x0d, 0x89, 0x9f, 0x79, 0x7c, 0xd8, 0xf9, 0x2a, 0x93,
	0x71, 0xdc, 0x7b, 0x62, 0x28, 0x48, 0x87, 0x52, 0xaf, 0x6b, 0x3b, 0x8f, 0x0c, 0xb5, 0xb1, 0x0f,
	0x9b, 0x6b, 0xcf, 0x19, 0x3d, 0x81, 0x4a, 0x31, 0x7a, 0xf3, 0x27, 0x70, 0xe7, 0x1d, 0x33, 0x7b,
	0x39, 0x01, 0x2f, 0xa8, 0x8d, 0x3f, 0x24, 0x50, 0x6c, 0x77, 0xca, 0xad, 0x60, 0xee, 0x74, 0xc5,
	0x0a, 0xe6, 0x4e, 0xfb, 0x3e, 0x42, 0xa0, 0x8a, 0x27, 0x9c, 0x35, 0x4f, 0xac, 0xd1, 0x7d, 0xa8,
	0xa6, 0x89, 0x3b, 0x25, 0xf9, 0xa8, 0x50, 0x04, 0x1f, 0x04, 0x94, 0xcd, 0x8a, 0xf5, 0xce, 0x6a,
	0x1f, 0x31, 0x43, 0xcb, 0xd7, 0x99, 0xa1, 0x8d, 0x3f, 0x25, 0xd0, 0x46, 0xd4, 0xcb, 0xe5, 0xbc,
	0xaf, 0xb3, 0x4b, 0x95, 0xf2, 0xfb, 0x54, 0x2a, 0x2b, 0x2a, 0x6f, 0x43, 0x39, 0x4d, 0x48, 0xcc,
	0xb9, 0x95, 0x6c, 0x48, 0xf1, 0xb0, 0xef, 0x7f, 0x32, 0x75, 0xbf, 0xca, 0x00, 0x23, 0xea, 0x75,
	0xa3, 0x20, 0xf8, 0x97, 0xbb, 0x7b, 0x0f, 0xc0, 0xcb, 0x18, 0x4b, 0x95, 0x7a, 0x8e, 0xf4, 0x7d,
	0xf4, 0x05, 0x6c, 0x17, 0xe9, 0xb9, 0x1b, 0xe7, 0xac, 0xac, 0x83, 0x5b, 0x79, 0x62, 0x24, 0xf0,
	0xbe, 0xbf, 0xea, 0x80, 0xba, 0xe6, 0x00, 0x02, 0x95, 0x91, 0x37, 0x4c, 0x9c, 0x5f, 0xc7, 0x62,
	0x7d, 0xc9, 0x95, 0xd2, 0x47, 0xb8, 0xa2, 0x5d, 0xcb, 0x95, 0x9f, 0x65, 0x28, 0x8f, 0xa8, 0x77,
	0x12, 0x31, 0xf2, 0x21, 0x4b, 0x56, 0x74, 0xc8, 0x6b, 0x3a, 0x3e, 0x07, 0x95, 0x7f, 0x5d, 0xf2,
	0x9f, 0x0f, 0x37, 0x96, 0xef, 0x9c, 0x57, 0x6b, 0xf1, 0x3f, 0x58, 0x10, 0x2e, 0x89, 0x53, 0x3f,
	0x42, 0x5c, 0xe9, 0x5a, 0xe2, 0xda, 0xa0, 0x0a, 0x61, 0x6b, 0x33, 0x42, 0x03, 0x79, 0x32, 0xca,
	0xe6, 0x43, 0x8f, 0x23, 0x32, 0x4f, 0x0f, 0xac, 0x89, 0x8d, 0x3b, 0x47, 0x86, 0xd2, 0xf8, 0xa9,
	0x04, 0xea, 0x24, 0x21, 0xf1, 0xaa, 0x6c, 0x69, 0x4d, 0xf6, 0x0e, 0x68, 0x09, 0xf1, 0x62, 0xc2,
	0x84, 0x1d, 0x1b, 0x38, 0x8f, 0xf8, 0x30, 0xa3, 0x7c, 0xc0, 0xe5, 0xcf, 0x20, 0x0b, 0x3e, 0x95,
	0x76, 0xf4, 0x0d, 0x6c, 0xcc, 0xdc, 0x84, 0x39, 0x09, 0x21, 0xe1, 0x7f, 0xbc, 0x16, 0x9c, 0x3f,
	0x26, 0x24, 0xb4, 0x13, 0xf4, 0x14, 0xc0, 0x73, 0xe7, 0xee, 0x4b, 0x3a, 0xa3, 0xec, 0xc2, 0x2c,
	0xd7, 0x95, 0x66, 0xad, 0xbd, 0x93, 0x37, 0x98, 0xbb, 0xd3, 0xea, 0x2e, 0xb2, 0x78, 0x85, 0x89,
	0x1a, 0xb0, 0x19, 0x92, 0x37, 0xcc, 0x61, 0xd1, 0x2b, 0x12, 0x2e, 0xdf, 0x7e, 0x95, 0x83, 0x36,
	0xc7, 0xfa, 0x3e, 0x7a, 0x08, 0x20, 0x8c, 0x15, 0x1c, 0x53, 0x17, 0x3f, 0x87, 0x8c, 0x95, 0xda,
	0x82, 0x87, 0xf5, 0xb4, 0x58, 0x36, 0xfe, 0x96, 0x00, 0x96, 0xff, 0x6f, 0xbd, 0x9b, 0x35, 0x80,
	0x51, 0xbf, 0xeb, 0x74, 0xb1, 0xd5, 0xb1, 0xf9, 0x57, 0x76, 0x03, 0x2a, 0x3c, 0xc6, 0x56, 0xa7,
	0x67, 0xc8, 0x68, 0x13, 0x74, 0x1e, 0xf5, 0x07, 0x3d, 0xeb, 0x07, 0x43, 0x41, 0x37, 0x60, 0x8b,
	0x87, 0xe3, 0xe1, 0xbe, 0xed, 0xf4, 0xac, 0x23, 0xcb, 0xb6, 0x8c, 0x52, 0x01, 0x1e, 0x76, 0x70,
	0xaf, 0x00, 0xb5, 0x62, 0xe3, 0x68, 0x82, 0x0f, 0x2c, 0xa3, 0x8c, 0xee, 0xc2, 0x6d, 0x1e, 0x4e,
	0x46, 0xbd, 0x8e, 0xcd, 0xbf, 0xe0, 0xd6, 0x0b, 0xa7, 0x3b, 0x9c, 0x0c, 0x6c, 0x0b, 0x1b, 0x15,
	0xfe, 0x61, 0xe7, 0x49, 0xbb, 0x73, 0x50, 0x1c, 0x43, 0x47, 0x3b, 0x80, 0xc4, 0xb1, 0x86, 0xc7,
	0xc7, 0xd6, 0xc0, 0x2e, 0x70, 0x28, 0xfe, 0xd9, 0xc9, 0xd0, 0xb6, 0x0a, 0xb0, 0x8a, 0xb6, 0xa0,
	0x3a, 0x19, 0x5b, 0xb8, 0x00, 0xd4, 0xc6, 0x6f, 0x12, 0xe8, 0x0b, 0x27, 0xd0, 0x1d, 0xa8, 0x2c,
	0xec, 0xcc, 0x6e, 0x62, 0x99, 0xe5, 0x56, 0xae, 0x5f, 0x2e, 0xf9, 0x3a, 0x97, 0xeb, 0xf2, 0xfd,
	0x50, 0xae, 0x73, 0x3f, 0xf6, 0x2a, 0x3f, 0x6a, 0x89, 0x77, 0x46, 0x02, 0xf7, 0xa5, 0x26, 0x98,
	0x8f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x40, 0xf1, 0xb6, 0x85, 0xdd, 0x0c, 0x00, 0x00,
}
