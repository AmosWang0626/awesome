package test

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {
	loginResp := message.LoginResponse{Code: 200, Error: "success"}
	fmt.Println(string(loginResp.Encode()))

	loginRespMsg := message.Message{Type: message.LoginResponseType, Data: string(loginResp.Encode())}

	fmt.Println(string(loginRespMsg.Encode()))

	var parseLoginRespMsg message.Message
	// !!!!!!!!!!!!!!!! must receive Decode response
	parseLoginRespMsg = parseLoginRespMsg.Decode(loginRespMsg.Encode())

	fmt.Println(parseLoginRespMsg)
}
