package main

import "fmt"

func main() {

	intChan := make(chan int, 50)

	intChan = writeChan(intChan)
	fmt.Printf("intChan len=%d, cap=%d\n", len(intChan), cap(intChan))
	close(intChan)

	readChan(intChan)

}

func writeChan(intChan chan int) chan int {
	for i := 0; i < 50; i++ {
		intChan <- i
	}
	return intChan
}

func readChan(intChan chan int) {
	for item := range intChan {
		fmt.Println(item)
	}
}
