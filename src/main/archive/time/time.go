package main

import (
	"amos.wang/awesome/src/main/utils"
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Println(now.Year(), now.Month().String(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
	fmt.Println(now.AddDate(0, 0, 1).Format("2006-01-02"))
	fmt.Println(now.Day())

	fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now().Format("2006/01/02 15:04:05"))

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	fmt.Println(utils.Format(time.Now(), utils.Year2Day))
	fmt.Println(utils.Format(time.Now(), utils.Year2Second))

}
