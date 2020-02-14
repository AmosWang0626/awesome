package process

import "net"

var (
	MyUserMgr *UserMgr
)

type UserMgr struct {
	OnlineUsers map[uint64]*UserProcess
}

func init() {
	MyUserMgr = &UserMgr{OnlineUsers: make(map[uint64]*UserProcess, 256)}
}

func (current *UserMgr) Save(up *UserProcess) {
	current.OnlineUsers[up.Account] = up
}

func (current *UserMgr) Delete(Account uint64) {
	delete(current.OnlineUsers, Account)
}

func (current *UserMgr) DeleteByConn(conn net.Conn) {
	for account, process := range current.OnlineUsers {
		if process.Conn == conn {
			current.Delete(account)
		}
	}
}

func (current *UserMgr) Select() map[uint64]*UserProcess {
	return current.OnlineUsers
}

func (current *UserMgr) SelectById(Account uint64) *UserProcess {
	return current.OnlineUsers[Account]
}
