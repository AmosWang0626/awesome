package main

import "fmt"

type doubleLink struct {
	name string
	pre  *doubleLink
	next *doubleLink
}

// 队列(先入先出) --- 双向链表实现
func main() {
	link1 := &doubleLink{name: "111111"}
	link2 := &doubleLink{name: "222222"}
	link3 := &doubleLink{name: "333333"}
	link4 := &doubleLink{name: "444444"}
	link5 := &doubleLink{name: "555555"}
	link6 := &doubleLink{name: "666666"}
	link7 := &doubleLink{name: "777777"}

	link1.lastInsert(link2)
	// 1 2
	link1.rangeForward()
	fmt.Println()

	link2.lastInsert(link4)
	// 1 2 4
	link1.rangeForward()
	fmt.Println()

	link2.centerInsert(link3)
	// 1 2 3 4
	link1.rangeForward()
	fmt.Println()

	link4.lastInsert(link6)
	// 1 2 3 4 6
	link1.rangeForward()
	fmt.Println()

	link4.centerInsert(link5)
	// 1 2 3 4 5 6
	link1.rangeForward()
	fmt.Println()

	link6.lastInsert(link7)
	// 1 2 3 4 5 6 7
	link1.rangeForward()
	fmt.Println()

	link4.deleteLink()
	// 1 2 3 5 6 7
	link1.rangeForward()
	fmt.Println()

	link7.deleteLink()
	// 1 2 3 5 6
	link1.rangeForward()
	fmt.Println()

	fmt.Print("link6 正向遍历\t")
	link6.rangeForward()
	fmt.Println("\nlink6 获取第一个节点后再正向遍历")
	link6.getFirst().rangeForward()
	fmt.Println()

	fmt.Print("link1 逆向遍历\t")
	link1.rangeReverse()
	fmt.Println("\nlink1 获取最后节点后再逆向遍历")
	link1.getLast().rangeReverse()
	fmt.Println()

}

func (current *doubleLink) lastInsert(target *doubleLink) {
	current.next = target
	target.pre = current
}

// 1(src) >> 2(target) << 3
func (current *doubleLink) centerInsert(target *doubleLink) {
	target.next = current.next
	target.pre = current

	current.next.pre = target
	current.next = target
}

// 1 >> 2(src) >> 3
func (current *doubleLink) deleteLink() {
	if current.next == nil {
		current.pre.next = nil
		return
	}
	current.pre.next = current.next
	current.next.pre = current.pre
}

// 获取第一个节点
func (current *doubleLink) getFirst() (first *doubleLink) {
	first = current
	if first.pre != nil {
		return first.pre.getFirst()
	}
	return first
}

// 获取最后一个节点
func (current *doubleLink) getLast() (last *doubleLink) {
	last = current
	if last.next != nil {
		return last.next.getLast()
	}
	return last
}

// 正向遍历
func (current *doubleLink) rangeForward() {
	fmt.Print(current.name, "\t")
	if current.next != nil {
		current.next.rangeForward()
	}
}

// 反向遍历
func (current *doubleLink) rangeReverse() {
	fmt.Print(current.name, "\t")
	if current.pre != nil {
		current.pre.rangeReverse()
	}
}
