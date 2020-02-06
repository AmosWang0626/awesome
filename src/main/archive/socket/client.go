package main

import (
	"amos.wang/awesome/src/main/archive/socket/pact"
	"amos.wang/awesome/src/main/utils"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

/*
1、连接服务端
2、读取键盘输入
3、发送给服务端
4、如果发送给服务端的内容是exit，则客户端退出
*/

func main() {
	// Dial 拨号
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatalln("[客户端] 连接服务端异常", err)
	}
	defer conn.Close()

	fmt.Printf("[客户端] 连接服务端成功, CONN: %v\n", conn)

	fmt.Print("请输入用户名:")
	username, err := bufio.NewReader(os.Stdin).ReadString('\n')
	username = strings.Trim(username, " \r\n")

	go processChat(conn)

	for {
		message := sendMessage(conn, username)
		if "exit" == message {
			break
		}
	}
	fmt.Println("[客户端] 关闭当前客户端")
}

/*
读取键盘输入,并将消息发送给服务端
*/
func sendMessage(conn net.Conn, username string) (request string) {
	fmt.Print("请输入消息[格式:username|input]:")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalln("[客户端] 读取键盘输入异常", err)
	}
	input = strings.Trim(input, " \r\n")

	if input == "exit" {
		request = "exit"
	} else {
		// {"to_user":"A","from_user":"amos","message":"hello world","send_time":"2020-02-06 16:56:14"}
		infoArr := strings.Split(input, "|")
		chatMsg := pact.ChatMsg{FromUser: username, ToUser: strings.Trim(infoArr[0], " "),
			Message: strings.Trim(infoArr[1], " "), SendTime: utils.NowStr()}
		request = string(pact.Encode(&chatMsg))
	}

	n, err := conn.Write([]byte(request))
	if err != nil {
		fmt.Println("[客户端] 发送消息异常", err)
		return "exit"
	}
	fmt.Printf("[客户端] 发送消息成功, 消息长度 %d(字节), 消息内容为: %v\n", n, input)
	return
}

/*
处理服务端消息
*/
func processChat(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		// 如果服务端很长时间不发消息, 该协程会一直阻塞在这里
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("[客户端] 接收服务端消息异常", err)
		} else {
			chatMsg := pact.Decode(buf[:n])
			fmt.Printf("\n[客户端] %v : %v\n请输入消息[格式:username|input]:", chatMsg.FromUser, chatMsg.Message)
		}
	}
}
