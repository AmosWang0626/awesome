package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	funcMemoryAddress()
	structConvert()
	structTag()
	fmt.Println("\n......................")
	cycleArea()
	integerPrint()
}

func (c integer) print() {
	fmt.Println("(c integer) print()", c)
}

func (c *integer) change() integer {
	return (*c) * 2
}

// String() == java toString()
func (c *integer) String() string {
	return fmt.Sprintf("this value is %d", c)
}

func integerPrint() {
	i := integer(2)
	i.print()
	fmt.Println("i.change()", i.change())
	fmt.Println("String()", &i)
}

// 4. 计算圆面积 方法为值的方式
type Cycle struct {
	radius float64
}

func (c *Cycle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func cycleArea() {
	c := Cycle{8}
	fmt.Printf("cycle radius=%g, area=%g\n", (&c).radius, (&c).area())
	fmt.Printf("cycle radius=%g, area=%g\n", c.radius, c.area())

}

// 3. 结构体标签
// name type tag
type Monitor struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

func (m Monitor) getName() string {
	return m.Name
}

func structTag() {
	m1 := Monitor{"123", 18, "this is desc"}
	fmt.Println("m1.getName()", m1.getName())
	str, err := json.Marshal(m1)
	if err != nil {
		fmt.Println("ERROR :: ", err)
	} else {
		fmt.Printf("JSON STR [%v]", string(str))
	}

}

// 2. 两个对象强转

type A struct {
	Num int
}

type B struct {
	Num  int
	Name string
}

type integer int

func structConvert() {
	var a A
	var b B
	// 必须一模一样
	//b = B(a)
	fmt.Println("\n", a, b)

	// 取别名也不能强转
	var i int
	var integer integer = 2
	i = int(integer)

	fmt.Println(i, integer)

}

// 1. 比较内存地址
type Point struct {
	x, y, z int
}

type Rect struct {
	left, right Point
}

type Rect2 struct {
	left, right *Point
}

func funcMemoryAddress() {
	// 一个结构体中, 字段的内存地址为连续分配的
	r1 := Rect{Point{1, 2, 3}, Point{3, 4, 5}}
	fmt.Printf("r1 left x=%p, y=%p, z=%p; right x=%p, y=%p, z=%p",
		&r1.left.x, &r1.left.y, &r1.left.z, &r1.right.x, &r1.right.y, &r1.right.z)

	// 结构体里内容为指针, 指针的地址是连续的, 但不能保证指针的内容地址是连续的
	r2 := Rect2{&Point{1, 2, 3}, &Point{3, 4, 5}}
	fmt.Printf("\nr2 point left %p; right %p", &r2.left, &r2.right)
	fmt.Printf("\nr2 content left %p; right %p", r2.left, r2.right)
}
