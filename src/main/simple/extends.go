package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string
	Age  int
}

type B struct {
	PhoneNo string
	Desc    string
}

// 继承
type C struct {
	A
	B
	int
}

func main() {
	var c C
	c.Name = "amos"
	c.Age = 18
	c.PhoneNo = "18937128861"
	jsonStr, _ := json.Marshal(c)
	fmt.Println(string(jsonStr))

	c2 := C{A{"amos", 18}, B{"18937128861", "this is description"}, 100}
	jsonStr2, _ := json.Marshal(c2)
	fmt.Println(string(jsonStr2))
	fmt.Println(c2.int)
}
