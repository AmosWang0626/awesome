package main

import (
	"fmt"
	"sort"
)

type user struct {
	id    int
	name  string
	age   byte
	score float32
}

func main() {
	//mapSave()
	//mapDeleteQuery()
	//mapSlice()
	mapSort()
}

func mapSave() {
	var mmp map[int]string
	mmp = make(map[int]string, 10)
	mmp[1] = "啦啦啦"
	mmp[2] = "哈哈哈"
	mmp[3] = "咻咻咻"
	fmt.Println(mmp)

	var userMap = make(map[int]user, 3)
	userMap[1] = user{111, "amos", 18, 88.50}
	userMap[2] = user{222, "hello", 18, 99.00}
	userMap[2] = user{222, "hei", 20, 88.00}
	fmt.Println(userMap)

	var userMap2 = map[int]user{
		1: {111, "amos", 18, 88.50},
	}
	userMap2[2] = user{222, "hello", 18, 99.00}
	fmt.Println(userMap2)

	var simpleMap = make(map[int]map[string]string, 2)
	simpleMap[1] = make(map[string]string, 2)
	simpleMap[1]["name"] = "amos"
	simpleMap[1]["age"] = "18"
	simpleMap[2] = make(map[string]string, 2)
	simpleMap[2]["name"] = "amos"
	simpleMap[2]["age"] = "20"
	fmt.Println(simpleMap)
}

func mapDeleteQuery() {
	fmt.Println("......................")
	var userMap = make(map[int]user, 3)
	userMap[1] = user{111, "amos", 18, 88.50}
	userMap[2] = user{222, "hello", 18, 99.00}
	userMap[3] = user{222, "hei", 20, 88.00}
	userMap[4] = user{222, "hei", 20, 88.00}
	fmt.Println(userMap)
	userMap3, ok := userMap[3]
	fmt.Printf("存在? %t, 值为: %v", ok, userMap3)

	delete(userMap, 3)
	fmt.Println(userMap)
	userMap3, ok = userMap[3]
	fmt.Printf("存在? %t, 值为:%v\n", ok, userMap3)

	for k, v := range userMap {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func mapSlice() {
	fmt.Println("......................")
	var userMap = make([]map[int]user, 3)
	userMap[0] = make(map[int]user, 1)
	userMap[0][1] = user{111, "amos", 18, 88.50}
	userMap[1] = make(map[int]user, 1)
	userMap[1][1] = user{222, "grace", 20, 88.50}
	fmt.Println(userMap)

	appendMap := map[int]user{
		1: {333, "grace", 20, 88.50},
	}
	userMap = append(userMap, appendMap)

	for k, v := range userMap {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func mapSort() {

	// 注意下边: size=1, 下边程序仍然会执行
	var intMap = make(map[int]int, 1)
	intMap[7] = 54
	intMap[3] = 21
	intMap[1] = 65
	intMap[5] = 32
	intMap[6] = 11
	fmt.Println(intMap)

	for k, v := range intMap {
		fmt.Printf("key: %v, value: %v\t", k, v)
	}
	fmt.Println()

	var keys []int
	for key, _ := range intMap {
		keys = append(keys, key)
	}
	fmt.Println("Before sorting", keys)
	sort.Ints(keys)
	fmt.Println("After sorting", keys)

	// 将排序后结果注入
	var sortedMap = make(map[int]int, len(keys))
	for _, key := range keys {
		fmt.Printf("key: %v, value: %v\t", key, intMap[key])
		sortedMap[key] = intMap[key]
	}
	fmt.Println()
	for k, v := range sortedMap {
		fmt.Printf("key: %v, value: %v\t", k, v)
	}
	fmt.Println()

}
