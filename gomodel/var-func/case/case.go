package _case

import "errors"

// 形参
// 局部变量
// 全局变量
// 多个返回值哟先判断error
// 函数可以有多个形参和多个返回值，返回值也可以和形参一样拥有参数名称
func SumCase(a, b int) (sum int, err error) {
	if a < 0 && b < 0 {
		err = errors.New("两数相加不能同时小于零")
		return
	}
	sum = a + b
	return

}

// 方法getname
func (u *User) GetName() string {
	return u.Name
}

// 方法（method）是一种与特定类型关联的函数。
func (u *User) GetAge() uint {
	return u.Age
}
func NewUser(name string, age uint) *User {
	return &User{Name: name, Age: age}
}

type User struct {
	Name string
	Age  uint
}
