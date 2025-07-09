package _case

import (
	"errors"
	"log"
)

func Sum(a, b int) (sum int, err error) {
	if a < 0 && b < 0 {
		err = errors.New("a and b must be positive")
		return 0, err
	}
	sum = a + b
	return
}

type SumFunc func(a, b int) (int, error)

// 将函数作为输入输出的实现中间件
func LogMiddleware(in SumFunc) SumFunc {
	//返回的函数为闭包函数
	return func(a, b int) (int, error) {
		log.Printf("日志中间件，记录操作数： a:%d b:%d", a, b)
		return in(a, b)
	}
}

//关键点
//函数作为参数和返回值（高阶函数特性）：
//Go 中函数是一等公民，可以像变量一样传递。
//LogMiddleware 接收一个函数 in，并返回一个增强后的函数。
//闭包（Closure）：
//返回的函数是一个闭包，捕获了外部函数的 in，因此可以调用原始函数。
//装饰器模式：
//不修改 in 的逻辑，而是通过包装添加额外功能（日志记录）。
