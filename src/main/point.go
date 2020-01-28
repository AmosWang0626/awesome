package main

import "fmt"

func main() {

	i := 100
	fmt.Printf("i value %v\n", i)
	fmt.Printf("i memory address %v\n", &i)
	poi := &i
	fmt.Printf("i point %v\n", poi)
	fmt.Printf("i point value %v\n", *poi)

}
