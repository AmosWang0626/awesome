package message

import (
	"encoding/json"
)

// user
const (
	LoginRequestType       = "LoginRequest"
	LoginResponseType      = "LoginResponse"
	RegisterRequestType    = "RegisterRequest"
	RegisterResponseType   = "RegisterResponse"
	UserAllRequestType     = "UserAllRequest"
	UserAllResponseType    = "UserAllResponse"
	UserOnlineRequestType  = "UserOnlineRequest"
	UserOnlineResponseType = "UserOnlineResponse"
	OnlineNoticeType       = "OnlineNotice"
	LogoutNoticeType       = "LogoutNotice"
)

const (
	SmsRequestType  = "SmsRequest"
	SmsResponseType = "SmsResponse"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (m *Message) Decode(buf []byte) *Message {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *Message) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}
