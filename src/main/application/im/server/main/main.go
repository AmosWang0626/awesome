package main

import (
	"amos.wang/awesome/src/main/application/im/common/imconstant"
	"amos.wang/awesome/src/main/application/im/server/dao"
	"amos.wang/awesome/src/main/application/im/server/process"
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

	// 初始化 Dao
	initDao()

	for {
		conn, err := listen.Accept()
		fmt.Println("等待客户端发送消息...")
		if err != nil {
			log_utils.Warning.Println("listen.Accept", err)
			continue
		}

		go mainProcess(conn)
	}

}

func mainProcess(conn net.Conn) {
	defer conn.Close()
	defer logout(conn)

	ps := process.Processor{Conn: conn}
	err := ps.Process()
	if err != nil {
		if err == io.EOF {
			log_utils.Debug.Println("客户端关闭,协程正常退出")
			return
		}

		fmt.Println("处理客户端消息异常", err)
		return
	}
}

func initDao() {
	dao.InitUserDao()
}

/*
客户端退出,删除对应在线列表中的人
*/
func logout(conn net.Conn) {
	process.MyUserMgr.DeleteByConn(conn)
}
