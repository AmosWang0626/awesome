package main

import (
	"amos.wang/awesome/src/main/application/im/client/cprocess"
	"fmt"
	"os"
)

var (
	userId  int
	userPwd string
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
			fmt.Print("请输入用户ID(纯数字)：")
			_, _ = fmt.Scanf("%d\n", &userId)
			fmt.Print("请输入用户密码：")
			_, _ = fmt.Scanln(&userPwd)

			up := cprocess.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("登录失败", err)
			} else {
				fmt.Println("登录成功")
				for {
					// 展示服务菜单
					cprocess.ShowMenu(fmt.Sprint(userId))
				}
			}
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误,请重新输入!")
		}
	}

}
