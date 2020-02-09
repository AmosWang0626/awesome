package message

import "encoding/json"

const (
	LoginRequestType  = "LoginRequest"
	LoginResponseType = "LoginResponse"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (m Message) Decode(buf []byte) Message {
	_ = json.Unmarshal(buf, &m)
	return m
}

func (m Message) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}

// login
type LoginRequest struct {
	UserId   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	Username string `json:"username"`
}

func (m LoginRequest) Decode(buf []byte) LoginRequest {
	_ = json.Unmarshal(buf, &m)
	return m
}

func (m LoginRequest) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}

type LoginResponse struct {
	Code  int    // 状态码
	Error string // 错误信息
}

func (m LoginResponse) Decode(buf []byte) LoginResponse {
	_ = json.Unmarshal(buf, &m)
	return m
}

func (m LoginResponse) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}
