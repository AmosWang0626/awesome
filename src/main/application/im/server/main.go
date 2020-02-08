package main

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/utils/log_utils"
	"encoding/binary"
	"fmt"
	"io"
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
	defer conn.Close()
	log_utils.Debug.Println("协程开始处理~~", conn.RemoteAddr())

	for {
		msg, err := readMessage(conn)
		if err != nil {
			log_utils.Error.Println("readMessage(conn)", err)
			if err == io.EOF {
				log_utils.Info.Println("客户端关闭,协程正常退出")
				return
			}
		}
		fmt.Println(msg)
	}

}

func readMessage(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 4096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		log_utils.Error.Println("len conn.Read", err)
		return
	}

	bufLen := binary.BigEndian.Uint32(buf[:4])
	n, err := conn.Read(buf[:bufLen])
	if uint32(n) != bufLen || err != nil {
		log_utils.Error.Println("msg conn.Read", err)
		return
	}

	msg = msg.Decode(buf[:bufLen])

	return
}
