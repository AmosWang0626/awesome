package module

import (
	"encoding/json"
)

const (
	ONLINE = iota
	NEW
)

type User struct {
	Account  uint64 `json:"account"`
	UserPwd  string `json:"user_pwd"`
	Username string `json:"username"`
}

func (m *User) Decode(buf []byte) *User {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *User) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}

/*
脱敏 user_info
*/
type UserInfo struct {
	Account  uint64 `json:"account"`
	Username string `json:"username"`
	Status   byte   `json:"status"`
}

func (m *UserInfo) Decode(buf []byte) *UserInfo {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *UserInfo) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}
