package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Id   int
	Name string
}

// 修改结构体字段值
func main() {
	var (
		model *user
		sv    reflect.Value
	)
	model = &user{}
	sv = reflect.ValueOf(model)

	fmt.Println(sv.Kind().String())

	sv = sv.Elem()
	sv.FieldByName("Id").SetInt(123)
	sv.FieldByName("Name").SetString("amos")

	fmt.Println(*model)

}
