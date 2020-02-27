package main

import (
	"errors"
	"fmt"
)

type cycleLink struct {
	id   int
	name string
	next *cycleLink
}

// 约瑟夫问题(类似丢手绢)
// n 个人围成一圈，编号为 k (1 <= k <= n) 的人从 1 开始数数，数到 m 那个人出列；
// 接着它的下一位从 1 开始报数，数到 m 那个人出列，直至所有人都出列，产生一个出队编号序列。
// 真正的约瑟夫问题, n=41, k=1, m=3, 约瑟夫和它的朋友的编号为16, 31。这个游戏最后会剩下他们两个人, 生死权在他们自己手里。
func main() {
	n := 10 // 只会被用一次,初始化环状链表
	k := 1  // 1 <= k <= n 同样只会被用一次,找到开始报数的人
	m := 2  // m = 1 时要特殊处理
	fmt.Println("===============josephus================")
	fmt.Printf("前置条件:\t(n=%d, k=%d, m=%d)\n", n, k, m)

	head := &cycleLink{}
	initCycleLink(head, n)
	fmt.Print("初始化完成:\t")
	rangeCycleLink(head)

	kNode, err := getK(k, head)
	fmt.Print("编号为k的Node:\t")
	if err != nil {
		fmt.Printf("k=%d, ERROR: %v\n", k, err)
		return
	}
	fmt.Printf("%+v\n", kNode)

	kNode.kill(m, n)
}

// (kill 3) 1 2 3 4
func (current *cycleLink) kill(m int, n int) {
	temp := current
	length := m
	if m == 1 {
		length = n + 1
	}
	for i := 0; i < length-2; i++ {
		temp = temp.next
	}
	fmt.Printf("删除节点: id=%d, name=%s\n", temp.next.id, temp.next.name)
	temp.next = temp.next.next
	if temp.next == temp.next.next {
		fmt.Printf("删除自己: id=%d, name=%s\n", temp.next.id, temp.next.name)
		return
	}
	temp.next.kill(m, n-1)
}

// 获取编号为K的节点
func getK(k int, head *cycleLink) (temp *cycleLink, err error) {
	temp = head
	for {
		if temp.id == k {
			return
		}
		// 遍历一圈后结束
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	err = errors.New("not found")
	return
}

// 遍历环状链表
func rangeCycleLink(head *cycleLink) {
	temp := head
	for {
		fmt.Printf("%d:%s\t", temp.id, temp.name)
		// 遍历一圈后结束
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

// 初始化环状链表
func initCycleLink(head *cycleLink, n int) {
	// @ 下一个字符为 A
	initPreName := '@'
	current := head
	for i := 1; i <= n; i++ {
		// 第一个节点,自己指向自己,形成环状单向链表
		if current.next == nil {
			current.id = i
			current.name = string(initPreName + int32(i))
			current.next = current
			continue
		}
		// 第二个及以后
		temp := &cycleLink{id: i, name: string(initPreName + int32(i))}
		temp.next = current.next
		current.next = temp

		current = temp
	}
}
