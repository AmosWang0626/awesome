package main

import (
	"fmt"
	"strconv"
	"strings"
)

type subject struct {
	id   int
	name string
}

func main() {
	ascll := '@'
	fmt.Println("'@'+1:", string(ascll+1))
	fmt.Println("'@'+3:", string(ascll+2))
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

	fmt.Println(len("123456"))
	// 字符串遍历
	fmt.Println(string([]rune("你好")))
	// 字符串转整数
	fmt.Println(strconv.Atoi("123456"))
	// 整数转字符串
	fmt.Println(strconv.Itoa(123) + "1433233")

	// str 2 byte array
	var byt = []byte("hello.go")
	fmt.Println(byt)
	// byte array 2 str
	var str = string(byt)
	fmt.Println(str)

	// 10进制 转 2,8,16
	fmt.Println(strconv.FormatInt(10, 2))
	fmt.Println(strconv.FormatInt(10, 8))
	fmt.Println(strconv.FormatInt(10, 16))

	fmt.Println(strings.Contains("100", "2"))
	fmt.Println(strings.Count("100", "0"))
	fmt.Println(strings.EqualFold("Abc", "abC"))
	fmt.Println(strings.Index("Abc", "bc"))
	fmt.Println(strings.ReplaceAll("abc", "b", "c"))
	fmt.Println(strings.ToLower("Abc"))
	fmt.Println(strings.ToUpper("Abc"))
	fmt.Println(strings.Split("Abc,Def,Gfi", ","))
	// hello world
	fmt.Println(strings.Trim("hello world \r    \n \r \n", " \r\n"))
}
