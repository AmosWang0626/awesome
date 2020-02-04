package main

import (
	"fmt"
	"reflect"
)

// 通过反射获取结构体字段,方法,标签
func main() {

	call1 := func(v1 int, v2 int) {
		fmt.Println(v1, v2)
	}

	call2 := func(v1 int, v2 int, str string) {
		fmt.Println(v1, v2, str)
	}

	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)

	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}

	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")

}
