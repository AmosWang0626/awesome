package cprocess

import (
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/module"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/utils/log_utils"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

/*
处理服务端推送
使用须知: 创建协程, 执行该方法
*/
func pushProcess(conn net.Conn) {
	for {
		tf := utils.Transfer{Conn: conn}

		// 处理服务器返回结果
		msg, err := tf.Read()
		if err != nil {
			log_utils.Error.Println("push process", err)
		}

		err = pushProcessMsg(msg)
		if err != nil {
			log_utils.Error.Println("pushProcessMsg", err)
		}
	}
}

func pushProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {

	case message.UserAllResponseType:
		userMap := make(map[string]module.User)
		err = json.Unmarshal([]byte(msg.Data), &userMap)
		if err != nil {
			return
		}
		for key, user := range userMap {
			fmt.Printf("%v : %v\n", key, string(user.Encode()))
		}

	case message.UserOnlineResponseType:
		userInfoMap := make(map[uint64]module.UserInfo)
		err = json.Unmarshal([]byte(msg.Data), &userInfoMap)
		if err != nil {
			return
		}
		for key, user := range userInfoMap {
			fmt.Printf("%v : %v\n", key, string(user.Encode()))
		}

	case message.OnlineNoticeType:
		var userInfo *module.UserInfo
		err = json.Unmarshal([]byte(msg.Data), &userInfo)
		if err != nil {
			return
		}
		MyClientUserInfoMgr.Save(userInfo)
		// 用户登录成功
		fmt.Println()
		if userInfo.Status == module.NEW {
			log_utils.Info.Println("新用户登录啦", string(userInfo.Encode()))
		} else if userInfo.Status == module.ONLINE {
			log_utils.Info.Println("用户登录啦", string(userInfo.Encode()))
		} else {
			log_utils.Info.Println("用户登录啦", string(userInfo.Encode()))
		}

	case message.LogoutNoticeType:
		// 用户退出成功
		fmt.Println()
		account, err := strconv.ParseUint(msg.Data, 10, 64)
		log_utils.Info.Println("用户退出啦", account)
		if err != nil {
			log_utils.Error.Println("删除本地在线用户失败")
		} else {
			MyClientUserInfoMgr.Delete(account)
		}

	case message.SmsResponseType:
		// 用户群发消息
		smsResp := &message.SmsResponse{}
		smsResp = smsResp.Decode([]byte(msg.Data))
		fmt.Printf("[群发消息] %v: %v\n", smsResp.Username, smsResp.Content)

	default:
		log_utils.Warning.Println("PushProcessMsg message.Type undefined", msg)
	}
	return
}
