package utils

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/utils/log_utils"
	"encoding/binary"
	"errors"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

/*
encoded message, conn write len and buffer
*/
func (current *Transfer) Write(buf []byte) (err error) {
	var bufLen [4]byte
	binary.BigEndian.PutUint32(bufLen[:4], uint32(len(buf)))
	_, err = current.Conn.Write(bufLen[:4])
	if err != nil {
		log_utils.Error.Println("conn.Write", err)
		return errors.New("发送数据包长度,请求服务端异常")
	}
	_, err = current.Conn.Write(buf)
	if err != nil {
		log_utils.Error.Println("conn.Write", err)
		return errors.New("发送数据包,请求服务端异常")
	}

	return nil
}

/*
conn read len and buffer, decode to message
*/
func (current *Transfer) Read() (msg message.Message, err error) {
	_, err = current.Conn.Read(current.Buf[:4])
	if err != nil {
		log_utils.Error.Println("len conn.Read", err)
		return
	}

	bufLen := binary.BigEndian.Uint32(current.Buf[:4])
	n, err := current.Conn.Read(current.Buf[:bufLen])
	if uint32(n) != bufLen || err != nil {
		log_utils.Error.Println("msg conn.Read", err)
		return
	}

	msg = msg.Decode(current.Buf[:bufLen])

	return
}
