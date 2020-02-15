package process

import (
	"amos.wang/awesome/src/main/application/im/common/imconstant"
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/module"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/application/im/server/dao"
	"amos.wang/awesome/src/main/utils/log_utils"
	"encoding/json"
	"net"
	"strconv"
)

type UserProcess struct {
	Conn     net.Conn
	Account  uint64
	UserInfo *module.UserInfo
}

func (current *UserProcess) processRegister(msg *message.Message) (err error) {
	registerReq := &message.RegisterRequest{}
	registerReq = registerReq.Decode([]byte(msg.Data))

	user := &module.User{Account: registerReq.Account, Username: registerReq.Username, UserPwd: registerReq.UserPwd}
	err = dao.MyUserDao.Register(user)

	// 注册相关业务逻辑
	registerResp := message.RegisterResponse{Code: imconstant.Success, Error: imconstant.SuccessMsg}
	if err != nil {
		registerResp.Code = imconstant.PreconditionFailed
		registerResp.Error = err.Error()
	} else {
		// 注册成功
		userInfo := &module.UserInfo{Account: user.Account, Username: user.Username, Status: module.NEW}
		// 返回用户信息给客户端
		registerResp.Body = string(userInfo.Encode())
		registerResp.OnlineUser = onlineUserInfoMap()
		// 用户上线通知
		onlineNotice(userInfo)
		// 保存用户在线状态
		current.UserInfo = userInfo
		current.Account = user.Account
		MyUserMgr.Save(current)
	}

	// 返回登录结果
	loginRespMsg := message.Message{Type: message.LoginResponseType, Data: string(registerResp.Encode())}

	tf := &utils.Transfer{Conn: current.Conn}
	return tf.Write(loginRespMsg.Encode())
}

func (current *UserProcess) processLogin(msg *message.Message) (err error) {
	loginReq := &message.LoginRequest{}
	loginReq = loginReq.Decode([]byte(msg.Data))

	user, err := dao.MyUserDao.Login(loginReq.Account, loginReq.UserPwd)

	// 登录相关业务逻辑
	loginResp := message.LoginResponse{Code: imconstant.Success, Error: imconstant.SuccessMsg}
	if err != nil {
		loginResp.Code = imconstant.PreconditionFailed
		loginResp.Error = err.Error()
	} else {
		// 登录成功
		userInfo := &module.UserInfo{Account: user.Account, Username: user.Username, Status: module.ONLINE}
		// 返回用户信息给客户端
		loginResp.Body = string(userInfo.Encode())
		loginResp.OnlineUser = onlineUserInfoMap()
		// 用户上线通知
		onlineNotice(userInfo)
		// 保存用户在线状态
		current.UserInfo = userInfo
		current.Account = user.Account
		MyUserMgr.Save(current)
	}

	// 返回登录结果
	loginRespMsg := message.Message{Type: message.LoginResponseType, Data: string(loginResp.Encode())}

	tf := &utils.Transfer{Conn: current.Conn}
	return tf.Write(loginRespMsg.Encode())
}

/*
用户退出通知
*/
func LogoutNotice(account uint64) {
	accountMsg := message.Message{Type: message.LogoutNoticeType, Data: strconv.FormatUint(account, 10)}
	userProcessMap := MyUserMgr.Select()
	for _, process := range userProcessMap {
		tf := &utils.Transfer{Conn: process.Conn}
		err := tf.Write(accountMsg.Encode())
		if err != nil {
			log_utils.Error.Println("OnlineNotice", err)
			continue
		}
	}
}

/*
用户上线通知
*/
func onlineNotice(userInfo *module.UserInfo) {
	userInfoMsg := message.Message{Type: message.OnlineNoticeType, Data: string(userInfo.Encode())}
	var userProcessMap = MyUserMgr.Select()
	for _, process := range userProcessMap {
		tf := &utils.Transfer{Conn: process.Conn}
		err := tf.Write(userInfoMsg.Encode())
		if err != nil {
			log_utils.Error.Println("OnlineNotice", err)
			continue
		}
	}
}

func (current *UserProcess) userAll() (err error) {
	userMap, err := dao.MyUserDao.UserAll()
	bytes, _ := json.Marshal(userMap)

	// 返回登录结果
	userMapMsg := message.Message{Type: message.UserAllResponseType, Data: string(bytes)}

	tf := &utils.Transfer{Conn: current.Conn}
	return tf.Write(userMapMsg.Encode())
}

func (current *UserProcess) userOnline() (err error) {
	bytes, _ := json.Marshal(onlineUserInfoMap())

	// 返回登录结果
	userMapMsg := message.Message{Type: message.UserOnlineResponseType, Data: string(bytes)}

	tf := &utils.Transfer{Conn: current.Conn}
	return tf.Write(userMapMsg.Encode())
}

/*
Online Users UserInfo Map
*/
func onlineUserInfoMap() map[uint64]*module.UserInfo {
	userProcessMap := MyUserMgr.Select()

	var userInfoMap map[uint64]*module.UserInfo
	userInfoMap = make(map[uint64]*module.UserInfo, len(userProcessMap))

	for key, process := range userProcessMap {
		userInfoMap[key] = process.UserInfo
	}

	return userInfoMap
}
