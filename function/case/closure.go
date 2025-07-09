package _case

import "log"

func Fib(n int) int {
	if n <= 2 {
		log.Fatal("n must be greater than 2")
	}
	t := tool()
	var res int
	for i := 0; i < n-2; i++ {
		res = t()
	}
	return res
}

// 斐波那契数列，X0+X1+X2+...+Xn-1 = Xn
func tool() func() int {
	var x0 = 0
	var x1 = 1
	var x2 = 0
	return func() int {
		x2 = x0 + x1
		x0 = x1
		x1 = x2
		return x2
	}
}

func ClosureTrap() {
	/*错误的方式
	for i := 0; i < 10; i++ {
		go func() {
			log.Println(i)
		}()
	}
	*/
	for i := 0; i < 10; i++ {
		go func(j int) {
			log.Println(j)
		}(i)
	}
}
