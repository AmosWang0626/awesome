package message

const (
	LoginRequestType  = "LoginRequest"
	LoginResponseType = "LoginResponse"
)

type Message struct {
	Type string
	Data string
}

// login
type LoginRequest struct {
	UserId   int
	UserPwd  string
	Username string
}

type LoginResponse struct {
	Code  int    // 状态码 500 未注册 200 登录成功
	Error string // 错误信息
}
