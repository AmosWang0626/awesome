package message

import (
	"amos.wang/awesome/src/main/application/im/common/module"
	"encoding/json"
)

// sms
type SmsRequest struct {
	*module.UserInfo
	Content string `json:"content"`
}

func (m *SmsRequest) Decode(buf []byte) *SmsRequest {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *SmsRequest) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}

type SmsResponse struct {
	*module.UserInfo
	Content string `json:"content"`
}

func (m *SmsResponse) Decode(buf []byte) *SmsResponse {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *SmsResponse) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}
