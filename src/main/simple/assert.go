package main

import "fmt"

// 断言
func main() {

	var x interface{}
	var b float64 = 12
	x = b
	if y, ok := x.(float64); ok {
		fmt.Println("convert success!", y)
	} else {
		fmt.Println("convert fail!")
	}

	checkType(100, 100.25, true)
}

func checkType(arr ...interface{}) {
	for _, item := range arr {
		switch item.(type) {
		case int:
			fmt.Printf("TYPE:%T \t VALUE:%v", item, item)
		case float32, float64:
			fmt.Printf("TYPE:%T \t VALUE:%v", item, item)
		default:
			fmt.Printf("TYPE:%T \t VALUE:%v", item, item)
		}
		fmt.Println()
	}
}
