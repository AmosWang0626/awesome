package cprocess

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/utils/log_utils"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func ShowMenu(conn net.Conn, username string) {
	defer fmt.Println("↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑")
	// 休眠100毫秒,等待服务端返回结果
	defer time.Sleep(100 * time.Millisecond)

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

	fmt.Println("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓")

	sms := &SmsProcess{}

	switch key {
	case 1:
		fmt.Println("\t 显示在线用户列表")
		clientOnlineUser()
	case 2:
		fmt.Println("\t 发送消息")
		sendGroupMessage(sms)
	case 3:
		fmt.Println("\t 信息列表")
	case 4:
		fmt.Println("\t 服务端用户管理")
		serverUserAll(conn)
	case 5:
		fmt.Println("\t 服务端在线用户")
		serverUserOnline(conn)
	case 6:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("\t 您的输入有误,请重新输入!")
	}

}

func sendGroupMessage(sms *SmsProcess) {
	fmt.Println("你想对大家说点什么呢？[输入exit则退出]")
	for {
		fmt.Print("请输入吧：")
		content, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		content = strings.Trim(content, " \r\n")
		if "exit" == content {
			return
		}
		err := sms.SendGroupSms(content)
		if err != nil {
			log_utils.Error.Println("发送群聊消息失败,失败原因", err)
		}
	}
}

func clientOnlineUser() {
	for key, user := range MyClientUserInfoMgr.Select() {
		fmt.Printf("%v : %v\n", key, string(user.Encode()))
	}
}

func serverUserAll(conn net.Conn) {
	msg := message.Message{Type: message.UserAllRequestType}

	// 请求服务器
	tf := utils.Transfer{Conn: conn}
	err := tf.Write(msg.Encode())
	if err != nil {
		log_utils.Error.Println("user all request", err)
	}
}

func serverUserOnline(conn net.Conn) {
	msg := message.Message{Type: message.UserOnlineRequestType}

	// 请求服务器
	tf := utils.Transfer{Conn: conn}
	err := tf.Write(msg.Encode())
	if err != nil {
		log_utils.Error.Println("user online request", err)
	}
}
