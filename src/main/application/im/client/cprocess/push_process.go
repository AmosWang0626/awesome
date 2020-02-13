package cprocess

import (
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/utils/log_utils"
	"net"
)

/*
处理服务端推送
使用须知: 创建协程, 执行该方法
*/
func pushProcess(conn net.Conn) {
	for {
		tf := utils.Transfer{Conn: conn}
		message, err := tf.Read()
		if err != nil {
			log_utils.Error.Println("push process error", err)
		}
		log_utils.Info.Println(message)
	}
}
