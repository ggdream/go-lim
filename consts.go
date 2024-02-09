package lim

import "encoding/json"

const (
	URLPathAuthSign          = "/auth/sign"
	URLPathUserCreate        = "/user/create"
	URLPathUserDelete        = "/user/delete"
	URLPathUserProfile       = "/user/profile"
	URLPathUserProfileUpdate = "/user/profile/update"
	URLPathMessageSend       = "/message/send"
)

type AuthSignReqModel struct {
	UserID string `json:"user_id"`
}

type AuthSignResModel struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expire_at"`
}

type MessageSendReqModel struct {
	Audio      *Audio  `json:"audio"`
	Custom     *Custom `json:"custom"`
	File       *File   `json:"file"`
	GroupID    *string `json:"group_id"`
	Image      *Image  `json:"image"`
	ReceiverID *string `json:"receiver_id"`
	Record     *Record `json:"record"`
	SenderID   *string `json:"sender_id"`
	Text       *Text   `json:"text"`
	Timestamp  int64   `json:"timestamp"`
	Type       int64   `json:"type"`
	Video      *Video  `json:"video"`
}

type Audio struct {
	ContentType string `json:"content_type"`
	Duration    int64  `json:"duration"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	URL         string `json:"url"`
}

type Custom struct {
	Content string `json:"content"`
}

type File struct {
	ContentType string `json:"content_type"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	URL         string `json:"url"`
}

type Image struct {
	ContentType  string `json:"content_type"`
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	ThumbnailURL string `json:"thumbnail_url"`
	URL          string `json:"url"`
}

type Record struct {
	ContentType string `json:"content_type"`
	Duration    int64  `json:"duration"`
	Size        int64  `json:"size"`
	URL         string `json:"url"`
}

type Text struct {
	Text string `json:"text"`
}

type Video struct {
	ContentType  string  `json:"content_type"`
	Duration     int64   `json:"duration"`
	Name         *string `json:"name"`
	Size         int64   `json:"size"`
	ThumbnailURL string  `json:"thumbnail_url"`
	URL          string  `json:"url"`
}

type MessageSendResModel struct {
	Audio      *Audio  `json:"audio"`
	Custom     *Custom `json:"custom"`
	File       *File   `json:"file"`
	GroupID    *string `json:"group_id"`
	Image      *Image  `json:"image"`
	ReceiverID *string `json:"receiver_id"`
	Record     *Record `json:"record"`
	SenderID   *string `json:"sender_id"`
	Text       *Text   `json:"text"`
	Timestamp  int64   `json:"timestamp"`
	Type       int64   `json:"type"`
	Video      *Video  `json:"video"`
}

type UserCreateReqModel struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type UserCreateResModel struct {
}

type UserDeleteReqModel struct {
	UserID string `json:"user_id"`
}

type UserDeleteResModel struct {
}

type UserProfileReqModel struct {
	UserID string `json:"user_id"`
}

type UserProfileResModel struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type UserUpdateReqModel struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type UserUpdateResModel struct {
}

type ResponseModel struct {
	Code int             `json:"code"`
	Desc string          `json:"desc"`
	Data json.RawMessage `json:"data"`
}
