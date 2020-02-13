package module

import (
	"encoding/json"
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
