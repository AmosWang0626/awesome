package main

import "fmt"

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

func Working(usb Usb) {
	usb.Start()
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
