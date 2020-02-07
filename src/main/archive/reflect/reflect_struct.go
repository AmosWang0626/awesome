package main

import (
	"fmt"
	"reflect"
)

// 修改结构体字段值
func main() {
	var (
		model *person
		sv    reflect.Value
	)
	model = &person{}
	sv = reflect.ValueOf(model)

	fmt.Println(sv.Kind().String())

	sv = sv.Elem()
	sv.FieldByName("Id").SetInt(123)
	sv.FieldByName("Name").SetString("amos")

	fmt.Println(*model)

}
