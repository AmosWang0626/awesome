package main

import (
	"errors"
	"fmt"
)

func main() {
	exception()

	fmt.Println("..........................")

	err := selfException()
	if err != nil {
		panic(err)
	}
	fmt.Println("this line cannot run")
}

func exception() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("[ERROR] ::", err)
		}
	}()

	n1 := 100
	n2 := 0
	n3 := n1 / n2
	fmt.Println("1231213131", n3)
}

func selfException() (error error) {
	error = errors.New("@#$%^&*()")
	return
}
