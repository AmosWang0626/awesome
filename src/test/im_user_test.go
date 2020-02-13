package test

import (
	"amos.wang/awesome/src/main/application/im/common/module"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/application/im/server/dao"
	"amos.wang/awesome/src/main/utils/log_utils"
	"testing"
)

func TestSaveUser(t *testing.T) {
	pool := utils.ImRedisPool()
	conn := pool.Get()
	defer conn.Close()
	defer pool.Close()

	user := &module.User{Account: 1000, Username: "amos.wang", UserPwd: "666666"}
	userDao := &dao.UserDao{Pool: utils.ImRedisPool()}
	err := userDao.Save(user)
	if err != nil {
		log_utils.Error.Println(err)
	}

	user, err = userDao.SelectUserByAccount(conn, 1000)
	if err != nil {
		log_utils.Error.Println(err)
	} else {
		log_utils.Info.Println(*user)
	}

	user, err = userDao.SelectUserByAccount(conn, 1001)
	if err != nil {
		log_utils.Error.Println(err)
	} else {
		log_utils.Info.Println(*user)
	}
}
