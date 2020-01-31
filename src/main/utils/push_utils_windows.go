package utils

import (
	"gopkg.in/toast.v1"
	"log"
	"time"
)

/*
不含标题的通知
*/
func SimpleNotice(message string) {
	Notice("", message)
}

func Notice(title, message string) {
	baseNotice(title, message)
}

// https://gopkg.in/toast.v1
func baseNotice(title, message string) {
	notice := Utf8ToGbk(title)
	message = Utf8ToGbk(message)
	log.SetPrefix(Format(time.Now(), Year2Second))

	notification := toast.Notification{
		// AppID 必填, 可随便拿一个现有程序的 AppID
		// PowerShell 里输入 Get-StartApps 可看到很多程序的 AppID
		// 示例: 运行的AppID >>> Microsoft.Windows.Shell.RunDialog
		AppID:   "Microsoft.Messaging_8wekyb3d8bbwe!x27e26f40ye031y48a6yb130yd1f20388991ax",
		Title:   notice,  // 标题
		Message: message, // 内容
		Icon:    "E:\\project\\golang\\awesome\\src\\main\\res\\talks.png",
		//Actions: []toast.Action{
		//	{"protocol", Utf8ToGbk("打开"), ""},
		//	{"protocol", Utf8ToGbk("关闭"), ""},
		//},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
