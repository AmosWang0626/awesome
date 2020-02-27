package main

import (
	"amos.wang/awesome/src/main/utils/date_utils"
	"errors"
	"fmt"
	"sync"
	"time"
)

type bankQueue struct {
	maxSize int
	front   int
	rear    int
	array   [30]int
}

var num = 0

// 银行排号
// 创建一个数组队列，每隔一段时间(随机)，为数组加一个数；
// 创建三个协程，每隔一段时间(随机)，到队列里取出数据。
// x 号协程 服务 >>> x 号客户
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	bankQ := &bankQueue{maxSize: 30, front: -1, rear: -1, array: [30]int{}}

	ticker := time.NewTicker(2 * time.Second)
	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			current := <-t.C
			num++
			fmt.Printf("[%s] >>>>>>>>>>>> [%d]号客户来了~~\n", date_utils.Format(current, date_utils.Hour2Second), num)
			err := bankQ.push(num)
			if err != nil {
				return
			}
			if num%5 == 0 {
				bankQ.rangeArr()
			}
		}
	}(ticker)

	go bankQ.process(1)
	go bankQ.process(2)
	go bankQ.process(3)

	wg.Wait()
}

func (que *bankQueue) process(no int) {
	ticker := time.NewTicker(3 * time.Second)
	go func(t *time.Ticker) {
		for {
			current := <-t.C
			popUser := que.pop()
			if popUser != -1 {
				fmt.Printf("[%s] %d 号协程 服务 [%d]号客户\n",
					date_utils.Format(current, date_utils.Hour2Second), no, popUser)
			}
		}
	}(ticker)
}

func (que *bankQueue) push(item int) (err error) {
	if que.rear == que.maxSize-1 {
		err = errors.New("queue full")
		return
	}
	que.rear++
	que.array[que.rear] = item
	return
}

func (que *bankQueue) pop() int {
	if que.front == que.rear {
		return -1
	}
	que.front++
	num := que.array[que.front]
	que.array[que.front] = -1

	return num
}

func (que *bankQueue) rangeArr() {
	fmt.Println(que.array)
}
