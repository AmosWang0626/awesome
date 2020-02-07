package main

import "fmt"

/*
1.常量定义时必须初始化
2.常量不能修改
3.常量只能修饰bool/int/float/string
*/
const (
	DefaultPassword string = "888888"
	ProjectName     string = "awesome"
	Version         string = "1.0.0"
)
const (
	ZERO = iota
	ONE
	TWO
	THREE
	FOUR
)

func main() {
	fmt.Println(ProjectName, DefaultPassword, Version)
	fmt.Println("ZERO", ZERO, "ONE", ONE, "TWO", TWO, "THREE", THREE, "FOUR", FOUR)
}
