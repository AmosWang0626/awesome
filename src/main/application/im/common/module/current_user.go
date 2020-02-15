package module

import "net"

type CurrentUser struct {
	Conn net.Conn
	*UserInfo
}
