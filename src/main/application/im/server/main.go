package main

import (
	"amos.wang/awesome/src/main/utils/log_utils"
	"fmt"
	"net"
)

func main() {
	port := 8889
	fmt.Println("服务器开始监听, 监听端口:", port)
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
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
	log_utils.Debug.Println("协程开始处理~~", conn.RemoteAddr())
}
