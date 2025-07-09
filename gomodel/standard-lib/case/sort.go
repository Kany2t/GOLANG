package _case

import (
	"fmt"
	"sort"
)

type sortUser struct {
	ID   int64
	Name string
	Age  uint8
}

type ByID []sortUser // 定义一个结构体，实现sort.Interface接口

// 获取长度
func (a ByID) Len() int {
	return len(a)
}

// 交换位置
func (a ByID) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 比较大小
func (a ByID) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}

func SortCase() {
	list := []sortUser{
		{1, "Tom", 18},
		{4, "Jerry", 20},
		{3, "Alice", 19},
		{2, "JamYang", 20},
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Age < list[j].Age
	})
	fmt.Println(list)

	//实现sort.Interface接口
	sort.Sort(ByID(list))
	fmt.Println(list)

}
