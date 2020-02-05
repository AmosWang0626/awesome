package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (cal Cal) GetSub(username string) {
	fmt.Printf("%v 完成了减法运算: %d - %d = %d\n",
		username, cal.Num1, cal.Num2, cal.Num1-cal.Num2)
}

func main() {
	var (
		model *Cal
		elem  reflect.Value
	)
	model = &Cal{}
	elem = reflect.ValueOf(model)
	elem.Elem().FieldByName("Num1").SetInt(123)
	elem.Elem().FieldByName("Num2").SetInt(100)
	fmt.Println(elem)

	fmt.Println(elem.NumMethod())
	res := elem.Method(0).Call([]reflect.Value{reflect.ValueOf("amos")})
	fmt.Println(res)

}
