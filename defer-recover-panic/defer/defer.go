package _case

import (
	"fmt"
	"io"
	"log"
	"os"
)

//defer关键词用来声明延迟调用函数
//该函数可以是匿名函数也可以是具名函数  延迟执行
//defer延迟函数的执行顺序是后进先出和栈一样

func DeferCase1() {
	fmt.Println("开始执行DeferCase1")
	defer func() {
		fmt.Println("调用了匿名函数1")
	}()
	defer f1()

	defer func() {
		fmt.Println("调用了匿名函数2")
	}()

	fmt.Println("DeferCase1结束执行")

}

// 参数预计算
func DeferCase2() {
	i := 1
	//传参
	defer func(j int) {
		fmt.Println("defer j=", j)
	}(i + 1) //参数预计算

	//闭包
	defer func() {
		fmt.Println("defer j=", i)
	}()
	i = 99
	fmt.Println("i=", i)
	//i= 99
	//defer j= 99
	//defer j= 2
}

// 返回值
// defer函数执行在return之后
func DeferCase3() {
	i, j := f2()
	fmt.Println("i=", i, "j=", *j, "g=", g)
	//f2 g: 100
	//i= 100 j= 200 g= 200
}

var g = 100

func f2() (int, *int) {
	defer func() {
		g = 200
	}()
	fmt.Println("f2 g:", g)
	return g, &g
}

func ExceptionCase() {
	defer func() {
		//回复协程捕获异常
		err := recover()
		if err != nil {
			fmt.Println("异常处理的defer recover", err)
			//异常处理
		}

	}()
	fmt.Println("开始执行ExceptionCase")
	panic("抛出异常")
	fmt.Println("ExceptionCase结束执行")

}

func FileReadCase() {
	file, err := os.Open("readme.md")
	if err != nil {
		log.Fatal(err)
	}
	//通过defer对资源释放
	defer func() {
		file.Close()
		fmt.Println("释放文件资源")
	}()
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
	}

}
func f1() {
	fmt.Println("调用了具名函数f1")
}

//打印结果
//开始执行DeferCase1
//DeferCase1结束执行
//调用了匿名函数2
//调用了具名函数f1
//调用了匿名函数1
