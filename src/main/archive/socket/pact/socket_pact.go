package pact

import (
	"encoding/json"
	"log"
)

/*
ChatMsg: 聊天对象
*/
type ChatMsg struct {
	ToUser   string `json:"to_user"`
	FromUser string `json:"from_user"`
	Message  string `json:"message"`
	SendTime string `json:"send_time"`
}

func init() {
	log.SetPrefix("[ERROR :: socket_pact] ")
}

func Decode(buf []byte) ChatMsg {
	var cm ChatMsg
	err := json.Unmarshal(buf, &cm)
	if err != nil {
		log.Fatalln("Decode Error", err)
	}
	return cm
}

func Encode(cm *ChatMsg) []byte {
	buf, err := json.Marshal(cm)
	if err != nil {
		log.Fatalln("Encode Error", err)
	}
	return buf
}
