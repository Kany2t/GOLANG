package _case

import "fmt"

// 基本接口，可用于变量定义
type ToString interface {
	String() string
}

//var s ToString

func (u user) String() string {
	return fmt.Sprintf("ID:%d, Name:%s ,Age: %d", u.ID, u.Name, u.Age)
}
func (addr address) String() string {
	return fmt.Sprintf("ID:%d, Province:%s ,City: %s", addr.ID, addr.Province, addr.City)
}

// 泛型接口
type GetKey[T comparable] interface {
	any
	GetID() T
}

// 定义一个user结构体的GetID方法，返回user结构体的ID字段
func (u user) GetID() int64 {
	// 返回user结构体的ID字段
	return u.ID
}
func (addr address) GetID() int {
	return addr.ID
}

//var s GetKey[int]

// slice 转map
func listToMap[k comparable, T GetKey[k]](list []T) map[k]T {
	mp := make(MapT[k, T], len(list))
	for _, data := range list {
		mp[data.GetID()] = data
	}
	return mp
}
func InterfaceCase() {
	userList := []GetKey[int64]{
		user{ID: 1, Name: "张三", Age: 18},
		user{ID: 2, Name: "李四", Age: 19},
	}
	addrList := []GetKey[int]{
		address{ID: 1, Province: "北京", City: "北京"},
		address{ID: 2, Province: "上海", City: "上海"},
	}
	userMp := listToMap[int64, GetKey[int64]](userList)
	fmt.Println(userMp)
	addrMp := listToMap[int, GetKey[int]](addrList)
	fmt.Println(addrMp)
}
