package _case

import "fmt"

// 指针类型访问之前必须初始化，值类型访问之前可以不用赋值，他有默认值
type user struct {
	Name string
	Age  int
	Addr Address
}
type Address struct {
	Province string
	City     string
}

func StructCase() {
	u := user{
		Name: "nick",
		Age:  22,
	}
	f2(&u)
	fmt.Println(u)
	u1 := &user{
		Name: "nick5",
		Age:  16,
	}
	fmt.Println(u1)
	fmt.Println(*u1)
	//指针类型
	var u2 *user
	u2 = new(user)
	//u2 := new(user)
	u2.Name = "jack"
	u2.Age = 29
	fmt.Println(u2)

}

func f2(u *user) {
	u.Age = 18
}

//值传递
