package cprocess

import (
	"amos.wang/awesome/src/main/application/im/common/imconstant"
	"amos.wang/awesome/src/main/application/im/common/imerror"
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/utils/log_utils"
	"errors"
	"net"
)

type UserProcess struct {
}

func (current *UserProcess) Register(account uint64, userPwd string, username string) (err error) {
	if account == 0 || userPwd == "" || username == "" {
		log_utils.Error.Println("用户账号、密码、用户名不能为空", account, userPwd)
		return imerror.UserRegisterInfoNotNil
	}
	//log_utils.Debug.Printf("用户账号: %d, 用户密码: %s\n", account, userPwd)

	return nil
}

func (current *UserProcess) Login(account uint64, userPwd string) (err error) {
	if account == 0 || userPwd == "" {
		log_utils.Error.Println("用户账号、密码不能为空", account, userPwd)
		return imerror.UserLoginInfoNotNil
	}
	//log_utils.Debug.Printf("用户账号: %d, 用户密码: %s\n", account, userPwd)

	conn, err := net.Dial(imconstant.Network, imconstant.Address)
	if err != nil {
		log_utils.Error.Println("net.Dial", err)
		return errors.New("客户端连接服务端异常")
	}
	// defer conn.Close()

	loginReq := message.LoginRequest{Account: account, UserPwd: userPwd}
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

	return processMsg(loginRespMsg, conn)
}

func processMsg(msg *message.Message, conn net.Conn) (err error) {
	switch msg.Type {

	case message.LoginResponseType:
		loginResp := &message.LoginResponse{}
		loginResp = loginResp.Decode([]byte(msg.Data))
		if loginResp.Code != 200 {
			return errors.New(loginResp.Error)
		}

		// 登录成功, 开启一个协程, 读取服务器推送信息
		go pushProcess(conn)

		return nil

	default:
		log_utils.Warning.Println("message.Type undefined", msg)
		return nil
	}
}
