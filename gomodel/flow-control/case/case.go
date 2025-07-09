package _case

import "fmt"

func FlowControlCase() {
	ifElseCase(0)
	ifElseCase(1)
	ifElseCase(2)
	forCase()
	switchCase("A", "ABC")

}

func ifElseCase(a int) {
	if a == 0 {
		fmt.Println("执行if语句块")
	} else if a == 1 {
		fmt.Println("执行else if语句块")
	} else {
		fmt.Println("执行else语句块")
	}
}

func forCase() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("位置一执行for循环语句块", i)
	}

	var i int
	var b = true
	for b {
		fmt.Println("位置二执行for循环语句块", i)
		i++
		if i >= 10 {
			b = false
		}
	}
	//遍历切片
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, data := range list {
		fmt.Printf("位置三执行for循环语句块 index: %d ,data: %d  \n", index, data)
	} //遍历切片，第一个值是索引值，第二个值是数据
	mp := map[string]int{"A": 1, "B": 2, "C": 3}
	for key, value := range mp {
		fmt.Printf("位置四执行for循环语句块 key: %v value: %v \n ", key, value)
	}

	str := "今天天气真好"
	for index, char := range str {
		fmt.Printf("位置五执行for循环语句块 index: %v char: %v \n", index, string(char))
	}
}
func switchCase(alpha string, in interface{}) {
	switch alpha {
	case "A":
		fmt.Println("执行A语句块")
	case "B":
		fmt.Println("执行B语句块")
		fallthrough //如果不要break就要添加fallthrough
	case "C", "D":
		fmt.Println("执行C或D语句块")
	default:
		fmt.Println("执行default语句块")
		//实现一个case默认break
	}
	switch in.(type) {
	case int:
		fmt.Println("in输入为int类型")
	case string:
		fmt.Println("in输入为string类型")
	default:
		fmt.Println("in未识别")
	}
}
func gotoCase() {
	var a = 0
lab1:
	fmt.Println("位置1执行goto语句块", a)
	for i := 0; i < 10; i++ {
		a += 1
		if a == 0 {
			a += 1
			goto lab1
		}

	}
}
