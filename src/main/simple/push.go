package main

import (
	"amos.wang/awesome/src/main/utils"
	"fmt"
	"gopkg.in/toast.v1"
	"log"
)

// https://gopkg.in/toast.v1
func main() {
	notice := utils.Utf8ToGbk("通知")
	message := utils.Utf8ToGbk("该休息一下了")

	fmt.Println(utils.GbkToUtf8(notice), ":", utils.GbkToUtf8(message))

	notification := toast.Notification{
		// 必填,可随便拿一个 `Get-StartApps` 运行:Microsoft.Windows.Shell.RunDialog
		AppID:   "Microsoft.Messaging_8wekyb3d8bbwe!x27e26f40ye031y48a6yb130yd1f20388991ax",
		Title:   notice,  // 通知
		Message: message, // 该休息一下了
		Icon:    "E:\\project\\golang\\awesome\\src\\main\\res\\talks.png",
		Actions: []toast.Action{
			{"protocol", utils.Utf8ToGbk("打开"), ""},
			{"protocol", utils.Utf8ToGbk("关闭"), ""},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("send notification success!")
	}
}
