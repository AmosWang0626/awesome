package process

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/utils/log_utils"
)

type SmsProcess struct {
}

func (current *SmsProcess) processMsg(msg *message.Message) (err error) {
	smsReq := &message.SmsRequest{}
	smsReq = smsReq.Decode([]byte(msg.Data))
	//log_utils.Info.Printf("USER: %v, SEND_MSG: %v\n", smsReq.Username, smsReq.Content)

	// 服务端转发消息至在线用户
	smsResp := message.SmsResponse{Content: smsReq.Content, UserInfo: smsReq.UserInfo}
	smsRespMsg := message.Message{Type: message.SmsResponseType, Data: string(smsResp.Encode())}
	buf := smsRespMsg.Encode()

	var userProcessMap = MyUserMgr.Select()
	for _, process := range userProcessMap {
		// 消息不发给自己
		if process.Account == smsReq.Account {
			continue
		}
		tf := &utils.Transfer{Conn: process.Conn}
		err := tf.Write(buf)
		if err != nil {
			log_utils.Error.Println("服务端转发群发消息异常", err)
			continue
		}
	}
	return
}
