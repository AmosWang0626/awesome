package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [2]float64
	ptr    *int
	slice  []int
	strMap map[string]string
}

// 面向对象 哈哈哈
func main() {
	var person Person
	// 指针/切片slice/map的零值都是nil,即还没有分配空间
	// print default value
	fmt.Println(person)

	person.Name = "amos"
	person.Age = 18
	person.Scores = [2]float64{12, 21}
	var ptr = &person.Age
	person.ptr = ptr
	person.slice = []int{1, 2}
	person.strMap = map[string]string{"key": "value"}
	fmt.Println(person)

	// 通过指针进行为对象字段赋值操作
	person2 := new(Person)
	// 标准的为指针对象赋值的方式如下
	// !!!!!!!!!! *person2.Name .的优先级要比*高,所以要这样写: (*person2).Name !!!!!!!!!!!
	(*person2).Name = "smith"
	// go 设计者为了方便起见, 也可以用下边的方式进行赋值操作
	person2.Name = "grace"
	person2.Age = 18
	fmt.Println(*person2)
	fmt.Println((*person2).Name)
	fmt.Println((*person2).Age)
	// go 设计者为了方便起见, 也可以用下边的方式进行取值操作
	fmt.Println(person2.Name)
	fmt.Println(person2.Age)

	person3 := &Person{"micro", 50, [2]float64{12, 21}, nil, nil, nil}
	person3.Name = "smith"
	person3.Age = 20
	fmt.Println(*person3)

}
