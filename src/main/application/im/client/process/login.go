package process

import (
	"amos.wang/awesome/src/main/application/im/common/imconstant"
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/utils/log_utils"
	"errors"
	"net"
)

func Login(userId int, userPwd string) (err error) {
	if userId == 0 || userPwd == "" {
		log_utils.Error.Println("用户ID、密码不能为空", userId, userPwd)
		return errors.New("用户ID、密码不能为空")
	}
	//log_utils.Debug.Printf("用户ID: %d, 用户密码: %s\n", userId, userPwd)

	conn, err := net.Dial(imconstant.Network, imconstant.Address)
	if err != nil {
		log_utils.Error.Println("net.Dial", err)
		return errors.New("客户端连接服务端异常")
	}
	defer conn.Close()

	loginReq := message.LoginRequest{UserId: userId, UserPwd: userPwd, Username: ""}
	msg := message.Message{Type: message.LoginRequestType, Data: string(loginReq.Encode())}

	// 请求服务器
	tf := utils.Transfer{Conn: conn}
	err = tf.Write(msg.Encode())
	if err != nil {
		return
	}

	// 处理服务器返回结果
	loginRespMsg, err := tf.Read()
	if err != nil {
		return
	}

	return processMsg(&loginRespMsg)
}

func processMsg(msg *message.Message) (err error) {
	switch msg.Type {

	case message.LoginResponseType:
		var loginResp message.LoginResponse
		loginResp = loginResp.Decode([]byte(msg.Data))
		if loginResp.Code != 200 {
			return errors.New(loginResp.Error)
		}

		return nil

	default:
		log_utils.Warning.Println("message.Type undefined", msg)
		return nil
	}
}
