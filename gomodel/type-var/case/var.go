package _case

// b:=true是错误的
import "fmt"

func VarDeclareCase() {
	var i, x, l = 1, 2, 3
	//通过var关键词声明变量
	var j = 100
	var f float32 = 100.123
	//可以自动检测数据类型，如果写错了会标灰

	//通过:=来赋值，只能用于局部变量
	b := true
	//数组，基本不用，因为是定长
	var arr = [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{1, 2, 3, 4, 5}
	var arr2 [5]int
	arr2[1] = 4
	arr2[3] = 5
	//没有赋值的为默认值为0
	fmt.Println(i, x, l, j, f, b, arr, arr1, arr2)
	//指针类型，用于表示变量地址的数据类型
	//var intPtr *int
	var i1 = 100
	var floatPtr *float32
	floatPtr = &f
	f1(&i1)
	fmt.Println(i1)
	fmt.Println(floatPtr)
	//var intPtr *int
	var inter interface{}
	fmt.Println(inter)

}

func f1(i *int) {
	*i = *i + 1
}
