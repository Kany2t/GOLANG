package _case

import "fmt"

// 声明Animal接口
// 定义Animal行为
type AnimalI interface {
	//吃
	Each()
	//喝
	Drink()
	//睡觉
	Sleep()
	//奔跑
	Run()
}
type animal struct{}

func (a animal) Each() {
	fmt.Println("Animal Each接口实现")
}
func (a *animal) Drink() {
	fmt.Println("Animal Drink接口实现")
}

func (a animal) Sleep() {
	fmt.Println("Animal Sleep接口实现")
}

func (a *animal) Run() {
	fmt.Println("Animal Run接口实现")
}

func Init() {
	a := animal{}
	a.Drink()
	a.Each()
	//println(&animal{})
	//println(&a)
	fmt.Printf("%p\n", &animal{}) // 显式打印指针
	fmt.Printf("%p\n", &a)
}
