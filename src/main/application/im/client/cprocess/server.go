package cprocess

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/module"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/utils/log_utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu(conn net.Conn, username string) {
	defer fmt.Println("-------------------------------------------")

	fmt.Println("\t\t\t恭喜 " + username + " 登录成功")
	fmt.Println("===========================================")
	fmt.Println("\t\t\t 1 显示在线用户列表")
	fmt.Println("\t\t\t 2 发送消息")
	fmt.Println("\t\t\t 3 信息列表")
	fmt.Println("\t\t\t 4 服务端用户管理")
	fmt.Println("\t\t\t 5 服务端在线用户")
	fmt.Println("\t\t\t 6 退出系统")
	fmt.Print("\t请选择(1-6)：")

	var key int
	_, _ = fmt.Scanln(&key)

	fmt.Println("-------------------------------------------")
	switch key {
	case 1:
		fmt.Println("\t 显示在线用户列表")
		userAll(conn)
	case 2:
		fmt.Println("\t 发送消息")
	case 3:
		fmt.Println("\t 信息列表")
	case 4:
		fmt.Println("\t 服务端用户管理")
		userAll(conn)
	case 5:
		fmt.Println("\t 服务端在线用户")
		userOnline(conn)
	case 6:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("\t 您的输入有误,请重新输入!")
	}

}

func userAll(conn net.Conn) {
	msg := message.Message{Type: message.UserAllRequestType}

	// 请求服务器
	tf := utils.Transfer{Conn: conn}
	err := tf.Write(msg.Encode())
	if err != nil {
		log_utils.Error.Println("user all request", err)
	}

	// 处理服务器返回结果
	userAllMsg, err := tf.Read()
	if err != nil {
		log_utils.Error.Println("user all response", err)
	}

	err = serverProcessMsg(userAllMsg, conn)
	if err != nil {
		log_utils.Error.Println("serverProcessMsg", err)
	}
}

func userOnline(conn net.Conn) {
	msg := message.Message{Type: message.UserOnlineRequestType}

	// 请求服务器
	tf := utils.Transfer{Conn: conn}
	err := tf.Write(msg.Encode())
	if err != nil {
		log_utils.Error.Println("user online request", err)
	}

	// 处理服务器返回结果
	userOnlineMsg, err := tf.Read()
	if err != nil {
		log_utils.Error.Println("user online response", err)
	}

	err = serverProcessMsg(userOnlineMsg, conn)
	if err != nil {
		log_utils.Error.Println("serverProcessMsg", err)
	}
}

func serverProcessMsg(msg *message.Message, conn net.Conn) (err error) {
	switch msg.Type {

	case message.UserAllResponseType:
		userMap := make(map[string]module.User)
		err = json.Unmarshal([]byte(msg.Data), &userMap)
		if err != nil {
			return
		}
		for key, user := range userMap {
			fmt.Printf("%v : %v\n", key, string(user.Encode()))
		}
		return

	case message.UserOnlineResponseType:
		userMap := make(map[uint64]module.User)
		err = json.Unmarshal([]byte(msg.Data), &userMap)
		if err != nil {
			return
		}
		for key, user := range userMap {
			fmt.Printf("%v : %v\n", key, string(user.Encode()))
		}
		return

	default:
		log_utils.Warning.Println("message.Type undefined", msg)
		return
	}
}
