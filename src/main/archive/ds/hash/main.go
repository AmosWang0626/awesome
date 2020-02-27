package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Google 面试题
// 一个公司，新员工来报道时，要求将该员工的信息存储起来（id, gender, age, address...）
// 当输入该员工id时，要求查出该员工的所有信息。
// 1.不能用数据库
// 2.添加时，保证按照员工id从低到高插入

const HashSize = 7

type HashTable struct {
	LinkArr [HashSize]HashLink
	Len     int
}

type HashLink struct {
	Head *User
}

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender bool   `json:"gender"`
	Next   *User  `json:"next"`
}

func main() {
	// link 0 [7, 14, 21, 28]	link 1 [8, 15, 22, 29]	link 2 [9, 16, 23, 30]	...
	hashTable := &HashTable{[HashSize]HashLink{}, HashSize}
	hashTable.push(initUser(7))
	hashTable.push(initUser(14))
	hashTable.push(initUser(9))
	hashTable.push(initUser(16))
	hashTable.push(initUser(13))
	fmt.Println("[INIT] 初始化HashTable完成")

	// 遍历 HashTable
	hashTable.rangeHashTable()

	hashTable.getUserById(7)
	hashTable.getUserById(8)

	hashTable.delUserById(7)
	hashTable.delUserById(8)
	hashTable.delUserById(13)

}

// 增
func (current *HashTable) push(user *User) {
	userHash := user.Id % current.Len
	if current.LinkArr[userHash].Head == nil {
		current.LinkArr[userHash] = HashLink{Head: user}
		return
	}
	temp := current.LinkArr[userHash].Head
	for {
		if temp.Next == nil {
			temp.Next = user
			break
		}
		temp = temp.Next
	}
}

// 查
func (current *HashTable) get(id int) (user *User, err error) {
	userHash := id % current.Len
	if current.LinkArr[userHash].Head == nil {
		err = errors.New("hash table not found")
		return
	}
	user = current.LinkArr[userHash].Head
	for {
		if user.Id == id {
			return
		}
		if user.Next == nil {
			user.Next = user
			break
		}
		user = user.Next
	}
	err = errors.New("hash table not found")
	return
}

// 删
func (current *HashTable) delete(id int) (err error) {
	userHash := id % current.Len
	if current.LinkArr[userHash].Head == nil {
		err = errors.New("id error, undefined (head is nil)")
		return
	}
	user := current.LinkArr[userHash].Head
	// 删除的为头结点
	if user.Id == id {
		current.LinkArr[userHash].Head = user.Next
		return
	}
	// 删除的非头结点
	for {
		if user.Next == nil {
			break
		}
		if user.Next.Id == id {
			user.Next = user.Next.Next
			return
		}
		user = user.Next
	}
	err = errors.New("id error, undefined")
	return
}

// 删除用户
func (current *HashTable) delUserById(id int) {
	err := current.delete(id)
	if err != nil {
		fmt.Printf("[DEL] 删除ID为 %d 的用户失败, 用户不存在\n", id)
	} else {
		fmt.Println("[DEL] 删除ID为", id, "的用户成功")
		current.rangeHashTable()
	}
}

// 根据 ID 获取用户
func (current *HashTable) getUserById(id int) {
	user, err := current.get(id)
	if err != nil {
		fmt.Printf("[GET] ID为 %d 的用户不存在\n", id)
		return
	}
	fmt.Println("[GET] ID为", id, "的用户是", user)
}

// 遍历 HashTable
func (current *HashTable) rangeHashTable() {
	fmt.Println("\t============================")
	for i, link := range current.LinkArr {
		temp := link.Head
		fmt.Print("\t", i, "\t")
		for {
			fmt.Print(temp, "\t")
			if temp == nil {
				break
			}
			temp = temp.Next
		}
		fmt.Println()
	}
	fmt.Println("\t============================")
}

// 生成一个用户
func initUser(id int) *User {
	return &User{
		Id:     id,
		Name:   "AMOS_" + strconv.Itoa(id),
		Gender: true}
}

// User 结构体默认 toString 方法
func (u *User) String() string {
	return u.Name
}
