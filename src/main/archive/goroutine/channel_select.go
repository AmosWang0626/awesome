package main

import "fmt"

func main() {

	intChannel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChannel <- i
	}
	strChannel := make(chan string, 10)
	for i := 0; i < 10; i++ {
		strChannel <- fmt.Sprintf("hello-%d", i)
	}

label:
	for {
		select {
		case v := <-intChannel:
			fmt.Printf("intChannel %d\n", v)
		case v := <-strChannel:
			fmt.Printf("strChannel %s\n", v)
		default:
			fmt.Println("............default select finish..................")
			// 退出方式1
			//return
			// 退出方式2
			break label
		}
	}

}
