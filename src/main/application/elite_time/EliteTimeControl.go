package main

import (
	"amos.wang/awesome/src/main/utils/date_utils"
	"amos.wang/awesome/src/main/utils/push_utils"
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
	finish := date_utils.Format(time.Now().Add(cycle*4*2), date_utils.Year2Second)

	fmt.Println("开始执行定时提醒任务 :: 一个周期60分钟, 每30分钟提醒一次注意力集中, 每45分钟休息一次, 休息15分钟")

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			current := <-t.C
			keep += 15
			temp := date_utils.Format(current, date_utils.Year2Second)

			fmt.Println(temp, "来了,老弟", keep)
			if temp == finish {
				return
			}

			// 45分钟 休息15分钟
			message := ""
			if keep%60 == 45 {
				message = fmt.Sprintf("该休息一下啦, 休息时间: %d, 截止时间: %v", 15, date_utils.Format(current.Add(15*time.Minute), date_utils.Hour2Second))
				fmt.Println(message)
				push_utils.SimpleNotice(message)
			} else if keep%60 == 30 {
				message = fmt.Sprintf("清醒一下哦, 还有15分钟就可以休息啦")
				push_utils.SimpleNotice(message)
				fmt.Println(message)
			} else if keep%60 == 0 {
				message = fmt.Sprintf("休息结束, 开始工作吧")
				push_utils.SimpleNotice(message)
				fmt.Println(message)
			}
		}
	}(ticker)

	wg.Wait()

}
