package dao

import (
	"amos.wang/awesome/src/main/application/im/common/imerror"
	"amos.wang/awesome/src/main/application/im/common/module"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

func InitUserDao() {
	MyUserDao = &UserDao{Pool: utils.ImRedisPool()}
}

type UserDao struct {
	Pool *redis.Pool
}

func (current *UserDao) Save(user *module.User) (err error) {
	_, err = current.Pool.Get().Do("hSet", "users", user.Account, string(user.Encode()))
	//log_utils.Debug.Println("Save", result)
	return
}

func (current *UserDao) SelectUserByAccount(conn redis.Conn, account uint64) (user *module.User, err error) {
	result, err := redis.String(conn.Do("hGet", "users", account))
	//log_utils.Debug.Println("Select", result)
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

func (current *UserDao) SelectAll(conn redis.Conn) (userMap *map[string]module.User, err error) {
	result, err := redis.StringMap(conn.Do("hGetAll", "users"))
	//log_utils.Debug.Println("SelectAll", result)
	userMapTemp := make(map[string]module.User)
	for key, val := range result {
		userTemp := &module.User{}
		userMapTemp[key] = *userTemp.Decode([]byte(val))
	}
	return &userMapTemp, err
}

func (current *UserDao) Register(user *module.User) (err error) {
	conn := current.Pool.Get()
	defer conn.Close()

	_, err = current.SelectUserByAccount(conn, user.Account)
	if err != imerror.UserUndefined {
		return imerror.UserAccountExist
	}

	return current.Save(user)
}

func (current *UserDao) Login(account uint64, userPwd string) (user *module.User, err error) {
	conn := current.Pool.Get()
	defer conn.Close()
	user, err = current.SelectUserByAccount(conn, account)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = imerror.UserPwdError
	}
	return
}

func (current *UserDao) UserAll() (userMap *map[string]module.User, err error) {
	conn := current.Pool.Get()
	defer conn.Close()

	return current.SelectAll(conn)
}
