package main

import (
	"fmt"
	"reflect"
)

// 通过反射获取结构体字段,方法,标签
func main() {

	user := user{Id: int32(111), Name: "AMOS", PhoneNo: "18937128861", Age: 18, Description: "this is description"}
	structReflect(user)

}

func structReflect(i interface{}) {

	typ := reflect.TypeOf(i)
	val := reflect.ValueOf(i)
	kd := val.Kind()
	if kd != reflect.Struct {
		return
	}

	numField := val.NumField()
	fmt.Printf("结构体类型: %v, 字段个数: %d\n", typ, numField)

	for i := 0; i < numField; i++ {
		fmt.Printf("\t参数完整类型: %v, Value: %v\n", typ.Field(i), val.Field(i))
		fmt.Printf("\t\tTag['json']: %v\n", typ.Field(i).Tag.Get("json"))
	}

	numMethod := val.NumMethod()
	fmt.Printf("方法数量: %d\n", numMethod)
	// 方法0 方法的 index 是以方法名 ASCLL 排序的
	fmt.Printf("\t返回值: %v\n", val.Method(0).Call(nil))

	params := []reflect.Value{
		reflect.ValueOf("amos"), reflect.ValueOf("18937128861"),
	}
	fmt.Printf("\t返回值: %v\n", val.Method(1).Call(params)[0])

}
