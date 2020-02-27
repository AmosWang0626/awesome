package main

import "fmt"

type singleLink struct {
	name string
	next *singleLink
}

// 队列(先入先出) --- 单向链表实现
// 缺点: 1.查找的方向只能是一个方向; 2.不能自我删除
func main() {
	link1 := &singleLink{name: "111111"}
	link2 := &singleLink{name: "222222"}
	link3 := &singleLink{name: "333333"}
	link4 := &singleLink{name: "444444"}
	link5 := &singleLink{name: "555555"}
	link6 := &singleLink{name: "666666"}
	link7 := &singleLink{name: "777777"}

	// 1.简单赋值
	link1.next = link3
	link3.next = link4
	link4.next = link6
	// 此时为 1 3 4 6
	link1.rangeLink()
	fmt.Println()

	// 2.将 2 插入链表，形成 1 2 3 4 6
	link1.next = link2
	link2.next = link3
	// 此时为 1 2 3 4 6
	link1.rangeLink()
	fmt.Println()

	// 3.将 5 插入链表，形成 1 2 3 4 5 6
	link4.next = link5
	link5.next = link6
	// 此时为 1 2 3 4 5 6
	link1.rangeLink()
	fmt.Println()

	// 4.将 7 追加入链表，形成 1 2 3 4 5 6 7
	link6.next = link7
	// 此时为 1 2 3 4 5 6 7
	link1.rangeLink()
	fmt.Println()

	// 5.将 3 从链表里删除，形成 1 2 4 5 6
	link2.next = link4
	// 此时为 1 2 4 5 6
	link1.rangeLink()
	fmt.Println()

}

func (current *singleLink) rangeLink() {
	fmt.Print(current.name, "\t")
	if current.next != nil {
		current.next.rangeLink()
	}
}
