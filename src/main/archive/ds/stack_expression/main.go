package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MathStack struct {
	MaxSize int
	Top     int
	Array   [20]string
}

func main() {
	expression := "-10 + 8 * 5 + 19 - 20 + 10 * 100 - 50 * 10"
	fmt.Println("表达式:", expression)

	numberStack := MathStack{MaxSize: 100, Top: -1, Array: [20]string{}}
	operationStack := MathStack{MaxSize: 100, Top: -1, Array: [20]string{}}

	arr := strings.Split(expression, " ")

	// 先计算优先级高的乘除运算
	for i, length := 0, len(arr); i < length; i++ {
		switch arr[i] {
		case "+":
			_ = operationStack.push(arr[i])
		case "-":
			_ = operationStack.push(arr[i])
		case "*":
			num1Str, _ := numberStack.pop()
			num1, _ := strconv.ParseInt(num1Str, 10, 64)
			i++
			num2, _ := strconv.ParseInt(arr[i], 10, 64)
			_ = numberStack.push(strconv.FormatInt(num1*num2, 10))
		case "/":
			num1Str, _ := numberStack.pop()
			num1, _ := strconv.ParseInt(num1Str, 10, 64)
			i++
			num2, _ := strconv.ParseInt(arr[i], 10, 64)
			_ = numberStack.push(string(num1 * num2))
		default:
			_ = numberStack.push(arr[i])
		}
	}

	var sum int64
	for {
		num1Str, err := numberStack.pop()
		if err != nil {
			break
		}
		num, _ := strconv.ParseInt(num1Str, 10, 64)
		ope, err := operationStack.pop()
		if err != nil {
			ope = "+"
		}
		switch ope {
		case "-":
			sum = sum - num
		case "+":
			sum = sum + num
		}
	}

	fmt.Println("运算结果:", sum)

}

func (current *MathStack) push(str string) (err error) {
	if current.Top == current.MaxSize-1 {
		err = errors.New("stack full")
		return
	}
	current.Top++
	current.Array[current.Top] = str
	return
}

func (current *MathStack) pop() (str string, err error) {
	if current.Top < 0 {
		err = errors.New("stack empty")
		return
	}
	str = current.Array[current.Top]
	current.Top--
	return
}
