package imerror

import "errors"

var (
	UserUndefined          = errors.New("用户未注册")
	UserPwdError           = errors.New("用户密码错误")
	UserAccountExist       = errors.New("用户名已占用")
	UserLoginInfoNotNil    = errors.New("用户账号、密码不能为空")
	UserRegisterInfoNotNil = errors.New("用户账号、密码、用户名不能为空")
)
