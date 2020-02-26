package main

import (
	"errors"
	"fmt"
)

type stack struct {
	MaxSize int
	Top     int
	Array   [3]string
}

func main() {
	s := stack{MaxSize: 3, Top: -1, Array: [3]string{}}
	s.push("111")
	s.push("222")
	s.push("333")
	s.push("444")
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	s.push("111")
	s.push("222")
	s.push("333")
	fmt.Println(s.pop())
	s.push("444")
	s.push("555")
	fmt.Println(s)
}

func (current *stack) push(str string) {
	if current.Top == current.MaxSize-1 {
		fmt.Println("stack full")
		return
	}
	current.Top++
	current.Array[current.Top] = str
}

func (current *stack) pop() (str string, err error) {
	if current.Top < 0 {
		err = errors.New("stack empty")
		return
	}
	str = current.Array[current.Top]
	current.Top--
	return
}
