package test

import (
	"amos.wang/awesome/src/main/application/im/common/module"
	"amos.wang/awesome/src/main/application/im/server/dao"
	"amos.wang/awesome/src/main/utils/log_utils"
	"testing"
)

func TestSaveUser(t *testing.T) {
	user := &module.User{Account: 1000, Username: "amos.wang", UserPwd: "666666"}
	userDao := &dao.UserDao{Pool: dao.ImRedisPool()}
	userDao.Save(user)

	user, err := userDao.Select(1000)
	if err != nil {
		log_utils.Error.Println(err)
	} else {
		log_utils.Info.Println(*user)
	}

	user, err = userDao.Select(1001)
	if err != nil {
		log_utils.Error.Println(err)
	} else {
		log_utils.Info.Println(*user)
	}
}
