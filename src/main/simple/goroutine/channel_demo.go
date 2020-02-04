package main

import (
	"fmt"
)

/*
启动8个协程, 计算1000内数字, 数列求和
*/

var (
	numberChan chan int
	resChan    chan map[int]uint64
	exitChan   chan bool

	routineCount = 8
)

func init() {
	numberChan = make(chan int, 1000)
	resChan = make(chan map[int]uint64, 1000)
	exitChan = make(chan bool, routineCount)

	for i := 1; i <= 1000; i++ {
		numberChan <- i
	}
	close(numberChan)

	fmt.Println(".........init finish.............")
}

func main() {

	for i := 0; i < routineCount; i++ {
		go routineSum(i)
	}

	go func() {
		for i := 0; i < routineCount; i++ {
			<-exitChan
		}
		close(resChan)
	}()

	for {
		res, resOk := <-resChan
		if !resOk {
			break
		}
		fmt.Printf("res:%v", res)
	}
	fmt.Println()
}

func routineSum(index int) {
	fmt.Println(fmt.Sprintf("第 %d 个协程正在执行...", index))
	for {
		number, numberOk := <-numberChan
		if !numberOk {
			break
		}
		// sum
		temp := number
		for i := 1; i < number; i++ {
			temp += i
		}
		resChan <- map[int]uint64{number: uint64(temp)}
	}
	fmt.Println(fmt.Sprintf("第 %d 个协程处理完成", index))

	exitChan <- true
}
