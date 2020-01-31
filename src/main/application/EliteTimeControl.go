package main

import (
	"amos.wang/awesome/src/main/utils"
	"fmt"
	"sync"
	"time"
)

/*
精英都是时间控
> 15/45/90
1.两小时120分钟,45分钟休息一次,按时提醒休息
2.45分钟休息15分钟
3.45分钟,分为三个15分钟. 前15分钟,首次计时开始; 后15分钟,即将到截止时间
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	keep := 0

	// 时间周期
	cycle := 15 * time.Minute

	// 按设置时间周期循环执行
	ticker := time.NewTicker(cycle)

	// 结束时间
	finish := utils.Format(time.Now().Add(cycle*4*2), utils.Year2Second)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			current := <-t.C
			keep += 15
			temp := utils.Format(current, utils.Year2Second)

			fmt.Println(temp, "来了,老弟", keep)
			if temp == finish {
				return
			}

			// 45分钟 休息15分钟
			message := ""
			if keep%60 == 45 {
				message = fmt.Sprintf("该休息一下啦, 休息时间: %d, 截止时间: %v", 15, utils.Format(current.Add(15*time.Minute), utils.Hour2Second))
				fmt.Println(message)
				utils.SimpleNotice(message)
			} else if keep%60 == 30 {
				message = fmt.Sprintf("清醒一下哦, 还有15分钟就可以休息啦")
				utils.SimpleNotice(message)
				fmt.Println(message)
			} else if keep%60 == 0 {
				message = fmt.Sprintf("休息结束, 开始工作吧")
				utils.SimpleNotice(message)
				fmt.Println(message)
			}
		}
	}(ticker)

	wg.Wait()

}
