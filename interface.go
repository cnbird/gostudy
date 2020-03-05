package main

import (
	"fmt"
	"sort"
)

// 声明 Hero 结构体
type Hero struct {
	Name string
	Age  int
	//Score float64
}

// 声明一个Hero 结构体切片类型
type HeroSlice []Hero

// 实现Interface 接口 的Len 方法
func (hs HeroSlice) Len() int {
	//fmt.Println("dadada")
	return len(hs)
}

// 实现Interface 接口 的Less 方法
func (hs HeroSlice) Less(i, j int) bool {
	//fmt.Println(i, j)
	return hs[i].Age < hs[j].Age
}

// 实现Interface 接口 的Swap 方法
func (hs HeroSlice) Swap(i, j int) {
	//fmt.Println(i, j)
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main() {
	var heroes HeroSlice
	//for i := 0; i < 10; i++ {
	//	hero := Hero{
	//		Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
	//		Age:  rand.Intn(25),
	//		//Score: rand.Float64(),
	//	}
	//	heroes = append(heroes, hero)
	//}
	heroes = HeroSlice{
		{
			Name: "AAA",
			Age:  55,
		},
		{
			Name: "BBB",
			Age:  22,
		},
		{
			Name: "CCC",
			Age:  0,
		},
		{
			Name: "DDD",
			Age:  22,
		},
		{
			Name: "EEE",
			Age:  11,
		},
	}
	//排序前的信息
	for _, v := range heroes {
		fmt.Println(v)
	}
	//	排序后的效果
	fmt.Printf("-------------------排序后的效果----------------------\n")
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}
}
