package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type SortHero struct {
	Name string
	Age  int
}

type HeroArr []SortHero

func (s HeroArr) Len() int {
	return len(s)
}

func (s HeroArr) Less(i int, j int) bool {
	return s[i].Age < s[j].Age
}

func (s HeroArr) Swap(i int, j int) {
	//temp := s[i]
	//s[i] = s[j]
	//s[j] = temp
	s[i], s[j] = s[j], s[i]
}

// 结构体排序
func main() {
	var list HeroArr
	// 不加下边这句, 每次生成的随机数都是一样的
	rand.Seed(time.Now().Unix())
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	list = append(list, SortHero{"梁山好汉", rand.Intn(20) + 20})
	for i := range list {
		fmt.Print(list[i].Name, ":", list[i].Age, "\t")
	}

	fmt.Println()

	sort.Sort(list)
	for i := range list {
		fmt.Print(list[i].Name, ":", list[i].Age, "\t")
	}
}
