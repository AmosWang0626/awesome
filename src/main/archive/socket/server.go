package main

import (
	"amos.wang/awesome/src/main/archive/socket/pact"
	"fmt"
	"io"
	"log"
	"net"
)

var (
	users = make(map[string]net.Conn, 0)
)

func main() {
	serverAddress := "0.0.0.0:8888"
	// 0.0.0.0 同时支持IPV4和IPV6地址; 127.0.0.1 仅支持IPV4
	listen, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalln("[服务端] 监听端口异常", err)
	}
	defer listen.Close()
	fmt.Printf("[服务端] 开启监听[%v] LISTEN: %v\n", serverAddress, listen)

	// ASCLL [@:64] [A:65] [B:66] [C:67]
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("[服务端] 客户端连接失败", err)
			continue
		}
		fmt.Printf("[服务端] 客户端连接成功 %v, 客户端IP: %v, 开启协程独立处理\n", conn, conn.RemoteAddr().String())
		go process(conn)
	}
}

/*
1、等待客户端响应
2、如果客户端说exit, 则退出
*/
func process(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	fmt.Printf("[服务端协程] 等待客户端消息: %v\n", conn.RemoteAddr().String())
	for {
		// 如果客户端很长时间不发消息, 该协程会一直阻塞在这里
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("[服务端协程] 接收客户端消息异常", err)
			if err == io.EOF {
				return
			}
			continue
		}

		message := string(buf[:n])
		fmt.Println("[服务端协程] 收到客户端消息:", message)
		if message == "exit" {
			break
		}

		chatMsg := pact.Decode(buf[:n])
		fmt.Println("[服务端协程]", chatMsg.FromUser, ":", chatMsg.Message)
		users[chatMsg.FromUser] = conn
		fmt.Println("[服务端协程] 最新用户列表", users)

		toUserConn := users[chatMsg.ToUser]
		if toUserConn != nil {
			n, err = toUserConn.Write(buf[:n])
			if err == nil {
				fmt.Println("[服务端协程] 消息转发成功", message)
			} else {
				fmt.Println("[服务端协程] 该用户未登录", chatMsg.ToUser)
			}
		}
	}
	// 协程关闭,注销该CONN
	for key, value := range users {
		if value == conn {
			delete(users, key)
		}
	}
	fmt.Println("[服务端协程] 关闭当前协程:", conn)
}
