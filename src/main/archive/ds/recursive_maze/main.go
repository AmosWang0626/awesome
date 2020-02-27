package main

import "fmt"

// 递归走迷宫
// 二维数组
// 1.value 0 没走过的点
// 2.value 1 墙
// 3.value 2 通路
// 4.value 3 死路
func main() {

	var mazeMap [8][7]int

	initMaze(&mazeMap)

	findWay(&mazeMap, 1, 1)

}

// 初始化墙
func initMaze(mazeMap *[8][7]int) {
	for i := 0; i < 7; i++ {
		mazeMap[0][i] = 1
		mazeMap[7][i] = 1
	}

	for i := 0; i < 8; i++ {
		mazeMap[i][0] = 1
		mazeMap[i][6] = 1
	}

	// 加两面墙
	mazeMap[3][1] = 1
	mazeMap[3][2] = 1

	//[1 1 1 1 1 1 1]
	//[1 0 0 0 0 0 1]
	//[1 0 0 0 0 0 1]
	//[1 1 1 0 0 0 1]
	//[1 0 0 0 0 0 1]
	//[1 0 0 0 0 0 1]
	//[1 0 0 0 0 0 1]
	//[1 1 1 1 1 1 1]

	mazeRange(mazeMap)
}

// 老鼠走迷宫
func findWay(mazeMap *[8][7]int, i, j int) bool {
	if mazeMap[6][5] == 2 {
		fmt.Println("找到出路啦")
		mazeRange(mazeMap)
		return true
	}
	// 是墙 return
	if mazeMap[i][j] == 1 {
		return false
	}
	// 上下左右
	if mazeMap[i][j] == 0 {
		mazeMap[i][j] = 2
		if findWay(mazeMap, i+1, j) { // 下
			return true
		}
		if findWay(mazeMap, i, j+1) { // 右
			return true
		}
		if findWay(mazeMap, i-1, j) { // 上
			return true
		}
		if findWay(mazeMap, i, j-1) { // 左
			return true
		}
		mazeMap[i][j] = 3
	}
	return false
}

// 遍历迷宫
func mazeRange(mazeMap *[8][7]int) {
	for i, length := 0, len(mazeMap); i < length; i++ {
		fmt.Println(mazeMap[i])
	}
	fmt.Println("=====分割线=====")
}
