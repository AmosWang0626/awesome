package cprocess

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/utils"
)

type SmsProcess struct {
}

func (current *SmsProcess) SendGroupSms(content string) (err error) {
	smsReq := message.SmsRequest{Content: content, UserInfo: MyCurrentUser.UserInfo}
	msg := message.Message{Type: message.SmsRequestType, Data: string(smsReq.Encode())}

	// 请求服务器
	tf := utils.Transfer{Conn: MyCurrentUser.Conn}
	err = tf.Write(msg.Encode())
	return
}
