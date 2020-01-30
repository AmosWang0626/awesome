package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := ""
	var num uint64
	fmt.Println("类型转换实例demo")
	var i int64 = 100
	str = strconv.FormatInt(i, 10)
	fmt.Printf("FormatInt %v\n", str)

	num, _ = strconv.ParseUint(str, 10, 10)
	fmt.Printf("ParseUint %v\n", num)
}
