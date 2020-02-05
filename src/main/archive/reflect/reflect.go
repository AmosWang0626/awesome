package main

import (
	"amos.wang/awesome/src/main/application/entity"
	"fmt"
	"reflect"
)

// 反射是运行时的

func main() {

	num := 100

	// 通过反射修改变量的值
	fn := reflect.ValueOf(&num)
	fn.Elem().SetInt(200)
	fmt.Println(num)

	fmt.Println(".........................................")
	hello(num)
	fmt.Println(".........................................")
	hello(entity.User{Id: int32(111), Name: "AMOS", PhoneNo: "18937128861", Age: 18, Description: "this is description"})

}

func hello(i interface{}) {
	refType := reflect.TypeOf(i)
	fmt.Printf("TypeOf \t\t TYPE:[%T]\t\tVALUE:[%v]\n", refType, refType)

	refValue := reflect.ValueOf(i)
	fmt.Printf("ValueOf \t\t TYPE:[%T]\t\tVALUE:[%v]\n", refValue, refValue)

	refValInterface := refValue.Interface()
	fmt.Printf("Interface \t\t TYPE:[%T]\t\tVALUE:[%v]\n", refValInterface, refValInterface)

	// TYPE是类型; KIND是类别
	refValKind := refValue.Kind()
	fmt.Printf("Kind \t\t TYPE:[%T]\t\tVALUE:[%v]\n", refValKind, refValKind)

}
