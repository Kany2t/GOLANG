package main

import (
	"fmt"
	_case "function/case"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	调用函数
	fmt.Println(_case.Sum(1, 2))
	//函数赋值给变量
	f1 := _case.Sum
	fmt.Println(f1(1, 2))

	//	将函数作为输入输出中间件

}
