package dao

import (
	"amos.wang/awesome/src/main/application/im/common/imerror"
	"amos.wang/awesome/src/main/application/im/common/module"
	"amos.wang/awesome/src/main/utils/log_utils"
	"github.com/garyburd/redigo/redis"
)

type UserDao struct {
	Pool *redis.Pool
}

func (current *UserDao) Save(user *module.User) {
	result, _ := current.Pool.Get().Do("hSet", "users", user.Account, string(user.Encode()))
	log_utils.Debug.Println("Save", result)
}

func (current *UserDao) Select(account uint64) (user *module.User, err error) {
	result, err := redis.String(current.Pool.Get().Do("hGet", "users", account))
	log_utils.Debug.Println("Select", result)
	if err != nil {
		if err == redis.ErrNil {
			err = imerror.UserUndefined
		}
		return
	}
	user = &module.User{}
	user = user.Decode([]byte(result))
	return
}
