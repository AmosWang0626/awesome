package main

import "fmt"

func main() {

	var name string
	var age uint8
	var phone uint64
	var address string

	fmt.Println("请输入用户名，年龄，手机号，住址：")
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Scanln(&phone)
	fmt.Scanln(&address)
	fmt.Printf("用户名:%v, 年龄:%v, 手机号:%v, 住址:%v", name, age, phone, address)

}
