package main

import "fmt"

func main() {

	expression1()

	strs := [...]string{"123", "456"}
	for m := 0; m < len(strs); m++ {
		fmt.Print(strs[m] + "\t")
	}
	fmt.Println()

	nums := [...]int{11, 12, 13, 14}
	fmt.Println(nums)
	nums[2] = 12
	fmt.Println(nums)

	fmt.Println("..........................")

	nums2 := [...]int{11, 12, 13, 14}
	fmt.Println(nums2)
	expression2(&nums2)
	fmt.Println(nums2)

	nums3 := [...]int{11, 12, 13, 14}
	// 含头不含尾
	fmt.Println(nums3[1:2])
	fmt.Println(len(nums3[1:2]))
	fmt.Println(cap(nums3[1:2]))

}

func expression2(arr *[4]int) {
	(*arr)[2] = 88
}

func expression1() {
	if i := 1; i >= 1 {
		fmt.Println("[i := 1; i >= 1] always true")
	}
	fmt.Println()

	age := 18
	fmt.Printf("请输入你的年龄:")
	//fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("大学生")
	} else if age < 6 {
		fmt.Println("小娃娃")
	} else if age < 12 {
		fmt.Println("小学生")
	} else {
		fmt.Println("中学生")
	}
	fmt.Println()

	score := 60
	fmt.Printf("请输入成绩分界[60, 90]:")
	//fmt.Scanln(&score)
	switch score {
	case 60:
		fmt.Println("及格分数")
	case 90:
		fmt.Println("优秀分数")
	default:
		fmt.Println("输入无效")
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Printf("base for-%v \t", i)
	}
	fmt.Println()

	j := 0
	for j < 5 {
		fmt.Printf("simple for-%v \t", j)
		j++
	}
	fmt.Println()

	k := 0
	for {
		if k < 2 {
			fmt.Printf("haha for-%v \t", k)
		} else {
			break
		}
		k++
	}
	fmt.Println()

	fmt.Println("*********遍历字符串*********")
	str := "家和万事兴"
	for l := 0; l < len(str); l++ {
		fmt.Printf("str >>> %v \t", str[l])
	}
	fmt.Println()

	fmt.Println("*********遍历字符串*********")
	str1 := "家和万事兴"
	str2 := []rune(str1)
	for l := 0; l < len(str2); l++ {
		fmt.Printf("str >>> %c \t", str2[l])
	}
	fmt.Println()

	fmt.Println("*********遍历字符串*********")
	str3 := "家和万事兴"
	for index, val := range str3 {
		fmt.Printf("[%d=%c] \t", index, val)
	}
	fmt.Println()
}
