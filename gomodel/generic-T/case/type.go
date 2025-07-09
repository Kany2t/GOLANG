package _case

import "fmt"

type user struct {
	ID   int64
	Name string
	Age  uint8
}
type address struct {
	ID       int
	Province string
	City     string
}

// map（集合）转slice（切片）
func mapTolist[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}

func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TTypeCase() {
	// 创建一个map，键为int64类型，值为user类型
	userMp := make(map[int64]user)
	// 向map中添加两个user类型的值
	userMp[1] = user{ID: 1, Name: "张三", Age: 18}
	userMp[2] = user{ID: 2, Name: "李四", Age: 19}
	// 将map转换为list
	userList := mapTolist[int64, user](userMp)

	ch := make(chan user)
	// 创建一个Chan类型的变量
	//var ch Chan[user]
	// 启动一个goroutine，用于打印Chan中的值
	go myPrintln(ch)
	// 遍历userList，将每个user类型的值发送到Chan中
	for _, u := range userList {
		ch <- u
	}
	// 创建一个map，键为int类型，值为address类型
	addrMp := make(map[int]address)
	// 向map中添加两个address类型的值
	addrMp[1] = address{ID: 1, Province: "北京", City: "北京"}
	addrMp[2] = address{ID: 2, Province: "上海", City: "上海"}
	// 将map转换为list
	addrList := mapTolist[int, address](addrMp)

	// 创建一个Chan类型的变量
	ch1 := make(chan address)
	// 启动一个goroutine，用于打印Chan中的值
	go myPrintln(ch1)
	// 遍历addrList，将每个address类型的值发送到Chan中
	for _, addr := range addrList {
		ch1 <- addr
	}
}

// 泛型切片的定义
type List[T any] []T

// 泛型定义map
type MapT[k comparable, v any] map[k]v

// 泛型定义channel通道
type Chan[T any] chan T
