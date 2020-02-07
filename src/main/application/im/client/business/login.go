package business

import (
	"amos.wang/awesome/src/main/utils/log_utils"
	"errors"
)

func Login(userId int, userPwd string) (err error) {
	if userId == 0 || userPwd == "" {
		log_utils.Error.Println("用户ID、密码不能为空", userId, userPwd)
		return errors.New("用户ID、密码不能为空")
	}
	log_utils.Info.Printf("用户ID: %d, 用户密码: %s\n", userId, userPwd)
	return nil
}
