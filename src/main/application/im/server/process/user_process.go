package process

import (
	"amos.wang/awesome/src/main/application/im/common/imconstant"
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (current *UserProcess) processLogin(msg *message.Message) (err error) {
	var loginReq message.LoginRequest
	loginReq = loginReq.Decode([]byte(msg.Data))

	// 登录相关业务逻辑
	loginResp := message.LoginResponse{Code: imconstant.Success, Error: imconstant.SuccessMsg}
	if loginReq.UserPwd != imconstant.DefaultPassword {
		loginResp.Code = imconstant.Unauthorized
		loginResp.Error = imconstant.UnauthorizedMsg
	}

	// 返回登录结果
	loginRespMsg := message.Message{Type: message.LoginResponseType, Data: string(loginResp.Encode())}

	tf := &utils.Transfer{Conn: current.Conn}
	return tf.Write(loginRespMsg.Encode())
}
