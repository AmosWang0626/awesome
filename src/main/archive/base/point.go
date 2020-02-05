package main

import (
	"amos.wang/awesome/src/main/utils"
	"fmt"
)

func main() {

	i := 100
	fmt.Printf("i value %v\n", i)
	fmt.Printf("i memory address %v\n", &i)
	poi := &i
	fmt.Printf("i point %v\n", poi)
	fmt.Printf("i point value %v\n", *poi)

	hi(&i)
	fmt.Println("hi(&i)", i)

	num1 := 100
	fmt.Printf("[num1 %T], [value %v], [addr %v]\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("[num2 %T], [value %v], [addr %v]\n", num2, *num2, &num2)

	fmt.Println(utils.ProjectName, utils.DefaultPassword, utils.Version)
	fmt.Println("ZERO", utils.ZERO, "ONE", utils.ONE, "TWO", utils.TWO, "THREE", utils.THREE, "FOUR", utils.FOUR)
}

func hi(i *int) {
	*i++
	*i++
	fmt.Println("hi(i *int)", *i)
}
