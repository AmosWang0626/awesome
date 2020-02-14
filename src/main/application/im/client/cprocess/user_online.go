package cprocess

import (
	"amos.wang/awesome/src/main/application/im/common/module"
)

var (
	MyClientUserInfoMgr *ClientUserInfoMgr
)

type ClientUserInfoMgr struct {
	OnlineUsers map[uint64]*module.UserInfo
}

func init() {
	MyClientUserInfoMgr = &ClientUserInfoMgr{OnlineUsers: make(map[uint64]*module.UserInfo, 256)}
}

func (current *ClientUserInfoMgr) Save(up *module.UserInfo) {
	current.OnlineUsers[up.Account] = up
}

func (current *ClientUserInfoMgr) SaveAll(userInfoMap map[uint64]*module.UserInfo) {
	for key, userInfo := range userInfoMap {
		current.OnlineUsers[key] = userInfo
	}
}

func (current *ClientUserInfoMgr) Delete(Account uint64) {
	delete(current.OnlineUsers, Account)
}

func (current *ClientUserInfoMgr) Select() map[uint64]*module.UserInfo {
	return current.OnlineUsers
}

func (current *ClientUserInfoMgr) SelectById(Account uint64) *module.UserInfo {
	return current.OnlineUsers[Account]
}
