package _case

import (
	"fmt"
	"math"
)

func MathCase() {
	fmt.Println("2的10次方", math.Pow(2, 10))
	fmt.Println("3的平方根", math.Sqrt(9))
	fmt.Println("绝对值", math.Abs(-1))
	fmt.Println("向上取整", math.Ceil(1.2))
	fmt.Println("向下取整", math.Floor(1.2))
	fmt.Println("四舍五入", math.Round(1.2))
	fmt.Println("返回两数中较大的值：", math.Max(1, 2))
	fmt.Println("返回两数中较小的值：", math.Min(1, 2))
	fmt.Println("常量", math.Pi, math.E)
	fmt.Println("90角的正弦值", math.Sin(math.Pi/2))
	fmt.Println("0角的余弦值", math.Cos(0))
	fmt.Println("0角度的正切值", math.Tan(0))
	//fmt.Println("随机数", math/rand.Intn(10))
	fmt.Println("1反正弦值", math.Asin(1))

}
