package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Id   int
	Name string
}

// 创建结构体设置字段值
func main() {
	var (
		model *person
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)
	fmt.Println(st.Kind().String())

	st = st.Elem()
	fmt.Println(st.Kind().String())

	elem = reflect.New(st)
	fmt.Println(elem.Kind().String())
	fmt.Println(elem.Elem().Kind().String())

	model = elem.Interface().(*person)
	elem = elem.Elem()

	elem.FieldByName("Id").SetInt(123)
	elem.FieldByName("Name").SetString("amos")

	fmt.Println(*model, model.Name)

}
