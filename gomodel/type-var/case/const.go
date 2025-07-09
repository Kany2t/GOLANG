package _case

import "fmt"

const (
	B = 1 << (10 * iota)
	KB
	MB
	_ //用下划线表示不接受这个变量，可以表示不连续的
	TB
)

type Gender uint

const (
	//iota枚举符号默认为1，可以累加
	FEMALE Gender = iota
	MALE
	THIRD
)

func ConstAndEnumCase() {
	const (
		A = 10
		B = 20
	)
	size := MB
	var gender = MALE
	fmt.Println(A, B, gender, size)
}
