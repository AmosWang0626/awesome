package main

import "fmt"

type cycleLink struct {
	id   int
	name string
	next *cycleLink
}

// 环状单链表，形成环状相对简单
// 注意：遍历一个被删除的节点，怎么办
func main() {
	link1 := &cycleLink{id: 1, name: "111111"}
	link2 := &cycleLink{id: 2, name: "222222"}
	link3 := &cycleLink{id: 3, name: "333333"}

	// init cycle
	link1.next = link2
	link2.next = link1
	fmt.Print("1.初始化\t")
	link2.rangeCycleLink()

	fmt.Print("2.插入\t")
	link2.insert(link3)
	link2.rangeCycleLink()

	fmt.Println("3.删除")
	link2.deleteById(2)

	fmt.Print("4.遍历\t")
	link2.rangeCycleLink()

	fmt.Print("5.遍历\t")
	link1.rangeCycleLink()

}

// 1(current) >> 2(target) >> 3
func (current *cycleLink) insert(target *cycleLink) {
	target.next = current.next
	current.next = target
}

// 1 >> 2(current) >> 3
func (current *cycleLink) deleteById(id int) {
	temp := current
	for {
		if temp.next.id == id {
			del := temp.next
			temp.next = temp.next.next
			del.next = nil
			break
		}
		if temp.next == current {
			break
		}
		temp = temp.next
	}
}

// cycle link range
func (current *cycleLink) rangeCycleLink() {
	if current.next == nil {
		fmt.Println("不要遍历我呦，我已经被删除了~~")
		return
	}
	temp := current
	for {
		fmt.Print(temp.name, "\t")
		if temp.next == current {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}
