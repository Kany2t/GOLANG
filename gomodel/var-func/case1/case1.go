package case1

import "fmt"

// 算术运算符
func ArithmeticCase() {
	var a = 21
	var b = 10
	var c int
	c = a + b
	fmt.Println("a+b的值为 %d \n", c)
	c = a - b
	fmt.Println("a-b的值为 %d \n", c)
	c = a * b
	fmt.Println("a*b的值为 %d \n", c)
	c = a / b
	fmt.Println("a/b的值为 %d \n", c)
	c = a % b
	fmt.Println("a%%b的值为 %d \n", c)
	a++
	fmt.Println("a++的值为 %d \n", c)
	a--
	fmt.Println("a--的值为 %d \n", c)

}

// 关系运算
func RelationalCase() {
	var a = 21
	var b = 10
	if a == b {
		fmt.Println("a 等于 b")
	} else {
		fmt.Println("a 不等于 b")
	}
	if a < b {
		fmt.Println("a 小于 b")
	} else {
		fmt.Println("a 大于等于 b")
	}
	if a > b {
		fmt.Println("a 大于 b")
	} else {
		fmt.Println("a 小于等于 b")
	}
}

// 逻辑运算
func LogicalCase() {
	var a = true
	var b = false
	if a && b {
		fmt.Println("a 和 b 都为 true")
	}
	if a || b {
		fmt.Println("a 和 b 至少有一个为 true")
	}
	if !a && !b {
		fmt.Println("a 和 b 都为 false")
	}
}

// 位运算
func BitwiseCase() {
	var a uint = 60
	var b uint = 13
	var c uint = 0
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	c = a & b //同1则1，其余为零
	fmt.Printf("%08b\n", c)
	c = a | b //同0则0，其余为1
	fmt.Printf("%08b\n", c)
	c = a ^ b //相同为0，不同为1
	fmt.Printf("%08b\n", c)
	c = a << 2 //左移2位
	fmt.Printf("%08b\n", c)
	c = a >> 2 //右移2位
	fmt.Printf("%08b\n", c)
}

//赋值运算
