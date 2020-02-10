package process

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/utils/log_utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (current *Processor) Process() error {
	for {
		tf := utils.Transfer{Conn: current.Conn}
		// 读取客户端消息
		msg, err := tf.Read()
		if err != nil {
			return err
		}

		// 处理客户端消息
		err = current.processMsg(&msg)
		if err != nil {
			return err
		}
	}
}

func (current *Processor) processMsg(msg *message.Message) (err error) {
	switch msg.Type {

	case message.LoginRequestType:
		log_utils.Debug.Println("login request", msg.Data)
		up := UserProcess{Conn: current.Conn}
		return up.processLogin(msg)

	default:
		log_utils.Warning.Println("message.Type undefined", msg)
		return nil
	}
}
