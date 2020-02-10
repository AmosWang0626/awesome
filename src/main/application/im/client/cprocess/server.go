package cprocess

import (
	"fmt"
	"os"
)

func ShowMenu(username string) {
	fmt.Println("\t\t\t恭喜 " + username + " 登录成功")
	fmt.Println("========================================")
	fmt.Println("\t\t\t 1 显示在线用户列表")
	fmt.Println("\t\t\t 2 发送消息")
	fmt.Println("\t\t\t 3 信息列表")
	fmt.Println("\t\t\t 4 退出系统")
	fmt.Print("请选择(1-4)：")

	var key int
	_, _ = fmt.Scanln(&key)

	fmt.Println()
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("您的输入有误,请重新输入!")
	}

}
