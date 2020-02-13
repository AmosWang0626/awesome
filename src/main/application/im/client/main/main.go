package main

import (
	"amos.wang/awesome/src/main/application/im/client/cprocess"
	"amos.wang/awesome/src/main/utils/input_utils"
	"fmt"
	"os"
	"strconv"
)

var (
	account  uint64
	userPwd  string
	username string
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[ERROR] ::", err)
		}
	}()

	for {
		fmt.Println("\t\t\t欢迎登录 IM 系统")
		fmt.Println("========================================")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Print("请选择(1-3)：")

		var key int
		_, _ = fmt.Scanln(&key)

		fmt.Println()

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			inputAccountAndPwd()

			up := cprocess.UserProcess{}
			err := up.Login(account, userPwd)
			if err != nil {
				fmt.Println("登录失败", err)
			} else {
				fmt.Println("登录成功")
				showLoginSuccessMenu()
			}
		case 2:
			fmt.Println("注册用户")
			inputAccountAndPwd()
			fmt.Print("请输入用户名：")
			_, _ = fmt.Scanf("%s\n", &username)

			up := cprocess.UserProcess{}
			err := up.Register(account, userPwd, username)
			if err != nil {
				fmt.Println("注册失败", err)
			} else {
				fmt.Println("注册成功")
				showLoginSuccessMenu()
			}
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误,请重新输入!")
		}
	}

}

/*
展示登录成功/注册成功服务菜单
*/
func showLoginSuccessMenu() {
	for {
		// 展示服务菜单
		if username == "" {
			username = fmt.Sprint(account)
		}
		cprocess.ShowMenu(username)
	}
}

/*
输入账号密码共用方法
*/
func inputAccountAndPwd() {
	loop := true
	for loop {
		fmt.Print("请输入用户账号(纯数字)：")
		var accountStr string
		_, _ = fmt.Scanf("%s\n", &accountStr)
		isNum := input_utils.IsNum(accountStr)
		if !isNum {
			fmt.Println("用户账号只能是数字")
		} else {
			loop = false
			account, _ = strconv.ParseUint(accountStr, 10, 64)
		}
	}
	fmt.Print("请输入用户密码：")
	_, _ = fmt.Scanln(&userPwd)
}
