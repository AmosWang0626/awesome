package business

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/utils/log_utils"
	"encoding/binary"
	"errors"
	"net"
)

func Login(userId int, userPwd string) (err error) {
	if userId == 0 || userPwd == "" {
		log_utils.Error.Println("用户ID、密码不能为空", userId, userPwd)
		return errors.New("用户ID、密码不能为空")
	}
	log_utils.Info.Printf("用户ID: %d, 用户密码: %s\n", userId, userPwd)

	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	defer conn.Close()
	if err != nil {
		log_utils.Error.Println("net.Dial", err)
		return errors.New("客户端连接服务端异常")
	}
	loginRequest := message.LoginRequest{UserId: userId, UserPwd: userPwd, Username: ""}
	loginBuf := loginRequest.Encode()
	msg := message.Message{Type: message.LoginRequestType, Data: string(loginBuf)}
	msgBuf := msg.Encode()

	var bufLen [4]byte
	binary.BigEndian.PutUint32(bufLen[:4], uint32(len(msgBuf)))
	_, err = conn.Write(bufLen[:4])
	if err != nil {
		log_utils.Error.Println("conn.Write", err)
		return errors.New("发送数据包长度,请求服务端异常")
	}
	_, err = conn.Write(msgBuf)
	if err != nil {
		log_utils.Error.Println("conn.Write", err)
		return errors.New("发送数据包,请求服务端异常")
	}

	return nil
}
