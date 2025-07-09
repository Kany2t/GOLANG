package _case

import "fmt"

// 泛型结构体
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

// 泛型receiver
// 泛型方法示例
func (myStruct MyStruct[T]) GetData() T {
	//var i interface{} = 10
	//a, ok := i.(int)
	//b, ok := t.(int)
	return myStruct.Data
}

type Top interface {
	//any
	GetData() any
}

func ReceiverCase() {
	data := 18
	myStruct := MyStruct[*int]{
		Name: "jack",
		Data: &data,
	}
	data1 := myStruct.GetData()
	fmt.Println("泛型receiver", data1)

	str := "hello"
	myStruct1 := MyStruct[*string]{
		Name: "jack",
		Data: &str,
	}
	str1 := myStruct1.GetData()
	fmt.Println("泛型receiver", str1)
}
