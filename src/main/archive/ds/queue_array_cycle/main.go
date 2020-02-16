package main

import "fmt"

// 队列-数组实现
// 先入先出
func main() {
	const maxSize = 4
	que := &queue{tail: 0, head: 0, maxSize: maxSize, array: [maxSize]int{}}
	que.push(222222)
	que.push(111111)
	que.push(333333)
	que.push(444444)
	que.push(555555)
	que.queryAll()
	que.pop()
	que.pop()
	que.queryAll()
	que.push(666666)
	que.push(777777)
	que.queryAll()
	que.pop()
	que.pop()
	que.pop()
	que.pop()
	que.queryAll()
	que.push(888888)
	que.push(999999)
	que.queryAll()
}

type queue struct {
	maxSize int
	tail    int
	head    int
	array   [4]int
}

func (que *queue) push(item int) {
	if que.isFull() {
		fmt.Println("\tpush ERROR :: queue full")
		return
	}
	que.tail = que.tail % que.maxSize
	que.array[que.tail] = item
	fmt.Printf("\tpush[%d], value >>> %d\n", que.tail, item)
	que.tail++
}

func (que *queue) pop() int {
	if que.isEmpty() {
		fmt.Println("\tpop ERROR :: queue empty")
		return -1
	}
	que.head = que.head % que.maxSize
	num := que.array[que.head]
	que.array[que.head] = -1
	fmt.Printf("\tpop[%d], value >>> %d\n", que.head, num)
	que.head++
	return num
}

func (que *queue) isEmpty() bool {
	return que.head == que.tail
}

func (que *queue) isFull() bool {
	return (que.tail+1)%que.maxSize == que.head
}

func (que *queue) size() int {
	return (que.tail + que.maxSize - que.head) % que.maxSize
}

func (que *queue) queryAll() {
	fmt.Println("队列数组为：", que.array, "，队列Size为：", que.size())
	tempHead := que.head
	fmt.Print("队列内容为：")
	for i := 0; i < que.size(); i++ {
		fmt.Print(que.array[tempHead], "\t")
		tempHead = (tempHead + 1) % que.maxSize
	}
	fmt.Println()
}
