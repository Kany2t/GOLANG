package _case

import "fmt"

func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	} // return a > b ? a : b
}

type CusNumT interface {
	int32 | uint8 | float64 | ~int64
}
type MyInt64 int64 //衍生类型，是具有基础类型int64的新类型，与int64不是同一类型

type MyInt32 = int32 //别名类型，与int32是同一类型

func getMaxCusNum[T CusNumT](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	} // return a > b ? a : b
}

func CusNumTCase() {
	var a, b MyInt32 = 3, 4
	var a1, b1 int32 = 5, 6
	fmt.Println("自定义泛型，数字比较", getMaxCusNum(a, b))
	fmt.Println("自定义泛型，数字比较", getMaxCusNum(a1, b1))
}

// 内置类型
func BuiltInCase() {
	var a, b string = "abc", "def"
	fmt.Println("内置comparable类型约束", getBuiltInComparable(a, b))
	fmt.Println("内置any类型约束", printBuiltInAny(a, b))
	var c, d float64 = 1.3, 1.7
	fmt.Println("内置comparable类型约束", getBuiltInComparable(c, d))
}
func getBuiltInComparable[T comparable](a, b T) bool {
	//comparable是内置接口，表示可比较的类型(支持==  ！=的类型)，如int、float、string等
	if a == b {
		return true
	}
	return false
}

func printBuiltInAny[T any](a, b T) bool {
	//any是内置接口，表示任意类型
	fmt.Println(a, b)
	return true
}
