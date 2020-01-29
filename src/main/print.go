package main

import (
	"fmt"
	"main/utils"
)

type subject struct {
	id   int
	name string
}

func main() {
	var aaa = 10
	const bbb = 10
	fmt.Println("Hello, World!", aaa, bbb)

	fmt.Println("天龙八部雪山飞狐\r张飞")
	fmt.Println("姓名\t年龄\t籍贯\t住址")
	fmt.Println("AAA\t18\t河南\t江苏")

	i := 100
	fmt.Println("i =", i)

	subj := subject{1, "world"}
	fmt.Printf("subj.name=%s\n", subj.name)

	fmt.Println(utils.DefaultPassword, utils.ProjectName, utils.Version)
}
