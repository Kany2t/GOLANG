package _case

import (
	"fmt"
	"os"
)

func FmtCase() {
	//打印到标准输出
	fmt.Println("今天天气很好")
	//格式化，并打印到标准输出
	fmt.Printf("%s天气真好\n", "今天")

	//格式化
	str := fmt.Sprintf("%s今天天气真好", "今天，明天")
	//输出到io.writer
	fmt.Fprint(os.Stderr, str)
}

func FmtCase1() {
	type simple struct {
		value int
	}
	a := simple{
		value: 10,
	}
	//通用占用符
	fmt.Printf("默认格式的值：%v\n", a)
	fmt.Printf("包含字段名称的值：%+v\n", a)
	fmt.Printf("JSON格式的值，go语法表示的值：%#v\n", a)
	fmt.Printf("值的类型：%T\n", a)
	fmt.Printf("输出字面上的百分号：%%10")

	//整数占位符
	v1 := 10
	v2 := 20170
	fmt.Printf("二进制：%b \n", v1)
	fmt.Printf("Unicode码点转字符： %c \n", v2)
	fmt.Printf("八进制：%o \n", v1)
	fmt.Printf("0o八进制：%O  \n", v1)
	fmt.Printf("十进制：%d \n", v1)
	fmt.Printf("十六进制：%x \n", v1)
	fmt.Printf("0x十六进制：%X \n", v1)
	fmt.Printf("十六进制，带0x前缀：%X \n", v1)
	fmt.Printf("用单引号将字符的值包起来： %q \n", v1)
	fmt.Printf("用双引号将字符串的值包起来：%Q \n", v1)

	//宽度设置
	fmt.Printf("宽度为10，默认右对齐，用0填充：%010d \n", v1)
	fmt.Printf("%v的十六进制：%x ; go的语法表示为： %#x ; 指定进制宽度为8，不足8位补0：%08b \n", v1, v1, v1, v1)

}
