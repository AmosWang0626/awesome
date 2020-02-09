package read_write

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/utils/log_utils"
	"encoding/binary"
	"errors"
	"net"
)

/*
encoded message, conn write len and buffer
*/
func Write(msg *message.Message, conn net.Conn) (err error) {
	msgBuf := msg.Encode()

	var msgBufLen [4]byte
	binary.BigEndian.PutUint32(msgBufLen[:4], uint32(len(msgBuf)))
	_, err = conn.Write(msgBufLen[:4])
	if err != nil {
		log_utils.Error.Println("conn.Write", err)
		return errors.New("发送数据包长度,请求服务端异常")
	}
	_, err = conn.Write(msgBuf)
	if err != nil {
		log_utils.Error.Println("conn.Write", err)
		return errors.New("发送数据包,请求服务端异常")
	}

	return nil
}

/*
conn read len and buffer, decode to message
*/
func Read(conn net.Conn) (msg message.Message, err error) {
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
