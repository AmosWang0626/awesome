package process

import (
	"amos.wang/awesome/src/main/application/im/common/imconstant"
	"amos.wang/awesome/src/main/application/im/common/message"
	"amos.wang/awesome/src/main/application/im/common/module"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/application/im/server/dao"
	"encoding/json"
	"net"
)

type UserProcess struct {
	Conn    net.Conn
	Account uint64
	User    *module.User
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
		registerResp.Body = string(user.Encode())
		current.User = user
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
		loginResp.Body = string(user.Encode())
		current.User = user
		current.Account = user.Account
		MyUserMgr.Save(current)
	}

	// 返回登录结果
	loginRespMsg := message.Message{Type: message.LoginResponseType, Data: string(loginResp.Encode())}

	tf := &utils.Transfer{Conn: current.Conn}
	return tf.Write(loginRespMsg.Encode())
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
	userPressMap := MyUserMgr.Select()
	var userMap map[uint64]*module.User
	userMap = make(map[uint64]*module.User, len(userPressMap))
	for key, process := range userPressMap {
		userMap[key] = process.User
	}
	bytes, _ := json.Marshal(userMap)

	// 返回登录结果
	userMapMsg := message.Message{Type: message.UserOnlineResponseType, Data: string(bytes)}

	tf := &utils.Transfer{Conn: current.Conn}
	return tf.Write(userMapMsg.Encode())
}
