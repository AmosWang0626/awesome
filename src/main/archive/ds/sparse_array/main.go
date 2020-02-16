package main

import "fmt"

type node struct {
	row int
	col int
	val int
}

// 稀疏数组(棋盘演示)
func main() {
	black := 1 // 黑子
	white := 2 // 白子

	// 1. 创建一个数组
	var chessArray [7][8]int
	chessArray[1][2] = black
	chessArray[2][3] = white
	chessArray[3][4] = black
	chessArray[4][5] = white

	fmt.Println("↓↓↓↓↓↓↓↓↓↓ 原始数组 ↓↓↓↓↓↓↓↓↓↓")
	for i, v := range chessArray {
		fmt.Println(i, "\t", v)
	}
	//for _, col := range chessArray {
	//	for _, row := range col {
	//		fmt.Print(row, "\t")
	//	}
	//	fmt.Println()
	//}

	// 2. 创建一个稀疏数组
	var sparseArray []node
	// 二维数组行数/二维数组列数/二维数组默认值
	sparseArray = append(sparseArray, node{row: len(chessArray), col: len(chessArray[0]), val: 0})
	for i, row := range chessArray {
		for j, col := range row {
			if col != 0 {
				sparseArray = append(sparseArray, node{row: i, col: j, val: col})
			}
		}
	}
	fmt.Println("↓↓↓↓↓↓↓↓↓↓ 稀疏数组 ↓↓↓↓↓↓↓↓↓↓")
	fmt.Println(sparseArray)

	fmt.Println("↓↓↓↓↓↓↓↓↓↓ 还原数组(数组还原) ↓↓↓↓↓↓↓↓↓↓")
	var parseArray [7][8]int
	for i, node := range sparseArray {
		if i == 0 {
			continue
		}
		parseArray[node.row][node.col] = node.val
	}
	for i, v := range parseArray {
		fmt.Println(i, "\t", v)
	}

	fmt.Println("↓↓↓↓↓↓↓↓↓↓ 还原数组(切片还原) ↓↓↓↓↓↓↓↓↓↓")
	finalRow := sparseArray[0].row
	finalCol := sparseArray[0].col
	finalDefaultVal := sparseArray[0].val
	var parseSlice [][]int
	parseSlice = make([][]int, 0)
	// 创建一个空的二维切片数组
	//for i := 0; i < finalRow; i++ {
	//	parseSlice = append(parseSlice, make([]int, finalCol))
	//}

	// 遍历行, 如果行在稀疏数组中存在, 则该行特殊处理
	// 否则, 该行为全0
	for i := 0; i < finalRow; i++ {
		var tempRow []int
		for _, node := range sparseArray[1:] {
			if i == node.row {
				for j := 0; j < finalCol; j++ {
					if j == node.col {
						tempRow = append(tempRow, node.val)
					} else {
						tempRow = append(tempRow, finalDefaultVal)
					}
				}
			}
		}
		if len(tempRow) == 0 {
			tempRow = make([]int, finalCol)
		}
		parseSlice = append(parseSlice, tempRow)
	}

	for i, v := range parseSlice {
		fmt.Println(i, "\t", v)
	}
}
