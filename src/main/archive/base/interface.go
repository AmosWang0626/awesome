package main

import "fmt"

/*
1. golang实现接口 不用显示地实现某个接口
2. 完全实现该接口的方法才算实现了该接口
*/

type Usb interface {
	Start()
	Stop()
}

type Device struct {
	Name string
	Life int
}

type Camera struct {
	Device
}

func (d Camera) Start() {
	fmt.Printf("Camera [%v] Start\n", d.Name)
}
func (d Camera) Stop() {
	fmt.Printf("Camera [%v] Stop\n", d.Name)
}

type Phone struct {
	Device
}

func (d Phone) Start() {
	fmt.Printf("Phone [%v] Start\n", d.Name)
}
func (d Phone) Stop() {
	fmt.Printf("Phone [%v] Stop\n", d.Name)
}
func (d Phone) Call() {
	fmt.Printf("Phone [%v] Call\n", d.Name)
}

func Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	c := Camera{Device{"Sony", 10}}
	p := Phone{Device{"MI", 5}}
	fmt.Println("Camera >>>>>>>>>>>>>>>", c)
	fmt.Println(" Phone >>>>>>>>>>>>>>>", p)
	Working(c)
	Working(p)
}

// 错误例子
type i1 interface {
	test1()
	test2()
}
type i2 interface {
	test1()
	test3()
}

// i3编译错误,应为i1和i2都有test1()
//type i3 interface {
//	i1
//	i2
//}
