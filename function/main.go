package main

import (
	"context"
	"fmt"
	_case "function/case"
	"os"
	"os/signal"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//调用函数
	fmt.Println(_case.Sum(1, 2))
	//函数赋值给变量
	var f1 _case.SumFunc = _case.Sum
	//f1 := _case.Sum
	fmt.Println(f1(1, 2))

	//	将函数作为输入输出中间件
	var f2 _case.SumFunc = _case.LogMiddleware(_case.Sum)
	//f2 := _case.LogMiddleware(_case.Sum)
	//再次申请中间件
	fmt.Println(f2(1, 2))
	f3 := _case.SumFunc(f2)
	fmt.Println(f3.Accumulation(1, 2, 3, 4))
	fmt.Println(f2.Accumulation(1, 2, 3, 4, 5))

	fmt.Println(_case.Fib(10))

	//闭包的陷阱
	_case.ClosureTrap()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	<-ctx.Done()
}
