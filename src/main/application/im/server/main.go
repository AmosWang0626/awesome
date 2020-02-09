package main

import (
	"amos.wang/awesome/src/main/application/im/common/imconstant"
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/read_write"
	"amos.wang/awesome/src/main/utils/log_utils"
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("服务器开始监听, 监听端口:", imconstant.Address)
	listen, err := net.Listen(imconstant.Network, imconstant.Address)
	defer listen.Close()
	if err != nil {
		log_utils.Error.Println("net.Listen", err)
		return
	}

	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			log_utils.Warning.Println("listen.Accept", err)
			continue
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()
	log_utils.Debug.Println("协程开始处理~~", conn.RemoteAddr())

	for {
		msg, err := read_write.Read(conn)
		if err != nil {
			if err == io.EOF {
				log_utils.Debug.Println("客户端关闭,协程正常退出")
				return
			}
		}
		err = processMsg(conn, &msg)
		if err != nil {
			fmt.Println("处理客户端消息异常", err)
			return
		}
	}

}

func processMsg(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {

	case message.LoginRequestType:
		log_utils.Debug.Println("login request", msg.Data)
		return processLogin(conn, msg)

	default:
		log_utils.Warning.Println("message.Type undefined", msg)
		return nil
	}
}

func processLogin(conn net.Conn, msg *message.Message) (err error) {
	var loginRequest message.LoginRequest
	loginRequest = loginRequest.Decode([]byte(msg.Data))

	loginResp := message.LoginResponse{Code: imconstant.Success, Error: imconstant.SuccessMsg}
	if loginRequest.UserPwd != imconstant.DefaultPassword {
		loginResp.Code = imconstant.Forbidden
		loginResp.Error = imconstant.UnauthorizedMsg
	}
	loginRespMsg := message.Message{Type: message.LoginResponseType, Data: string(loginResp.Encode())}
	return read_write.Write(&loginRespMsg, conn)
}
