package main

import "fmt"

// 队列-数组实现
// 先入先出
func main() {
	const maxSize = 5
	que := &queue{front: -1, rear: -1, maxSize: maxSize, array: [maxSize]int{}}
	que.push(222222)
	que.push(111111)
	que.push(333333)
	que.push(555555)
	que.push(444444)
	que.push(666666)
	que.queryAll()
	que.pop()
	que.pop()
	que.pop()
	que.pop()
	que.pop()
	que.pop()
	que.pop()
	que.queryAll()
	que.push(777777)
	que.queryAll()
}

type queue struct {
	maxSize int
	front   int
	rear    int
	array   [5]int
}

func (que *queue) push(item int) {
	if que.rear == que.maxSize-1 {
		fmt.Println("ERROR :: queue full")
		return
	}
	que.rear++
	fmt.Printf("push[%d], value >>> %d\n", que.rear, item)
	que.array[que.rear] = item
}

func (que *queue) pop() int {
	if que.front == que.rear {
		fmt.Println("ERROR :: pop undefined")
		return -1
	}
	que.front++
	num := que.array[que.front]
	que.array[que.front] = -1
	fmt.Printf("pop[%d], value >>> %d\n", que.front, num)

	return num
}

func (que *queue) queryAll() {
	fmt.Println(que.array)
}
