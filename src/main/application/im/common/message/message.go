package message

import (
	"amos.wang/awesome/src/main/application/im/common/module"
	"encoding/json"
)

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

// login
type LoginRequest struct {
	Account uint64 `json:"account"`
	UserPwd string `json:"user_pwd"`
}

func (m *LoginRequest) Decode(buf []byte) *LoginRequest {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *LoginRequest) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}

type LoginResponse struct {
	Code       int                         // 状态码
	Error      string                      // 错误信息
	Body       string                      // 当前用户信息
	OnlineUser map[uint64]*module.UserInfo // 在线用户
}

func (m *LoginResponse) Decode(buf []byte) *LoginResponse {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *LoginResponse) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}

// register
type RegisterRequest struct {
	Account  uint64 `json:"account"`
	UserPwd  string `json:"user_pwd"`
	Username string `json:"username"`
}

func (m *RegisterRequest) Decode(buf []byte) *RegisterRequest {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *RegisterRequest) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}

type RegisterResponse struct {
	Code       int                         // 状态码
	Error      string                      // 错误信息
	Body       string                      // 当前用户信息
	OnlineUser map[uint64]*module.UserInfo // 在线用户
}

func (m *RegisterResponse) Decode(buf []byte) *RegisterResponse {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *RegisterResponse) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}
