package _case

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

func ConvertCase() {
	//同一类数据转换（数字和数字，字符串和字符串）
	//不同类型数据转换（字符串和数字）
	//接口类型转其他类型
	//数字类型转换
	var num1 int = 100
	fmt.Println(int64(num1))
	//大转小要当心精度丢失
	var num2 int64 = 101
	fmt.Println(int(num2))

	//字符串与数字转换
	var num3 = 100
	fmt.Println(strconv.Itoa(num3) + "abc")
	var str1 = "100"
	fmt.Println(strconv.Atoi(str1))
	//因此，如果你需要将整数转换为字符串，并且需要正确处理负数，那么你应该使用 strconv.Itoa() 函数。如果你确定整数是非负的，那么你可以使用类型转换，但需要注意负数的情况。

	var num4 int64 = 1010
	fmt.Println(strconv.FormatInt(num4, 10))

	var str2 = "1010"
	fmt.Println(strconv.ParseInt(str2, 10, 64))

	//字符串与[]byte转换
	var str3 = "今天天气很好"
	//fmt.Println([]byte(str3))

	bytes1 := []byte(str3)
	fmt.Println(string(bytes1))
	//byte可以得到字节数

	//rune可以得到字符数

	//字符串与[]rune转换，rune是int32的别名
	//将字符串转化为rune切片，实际上rune切片中存储了字符串的Unicode码点
	var rune1 = []rune(str3)
	fmt.Println(rune1)
	fmt.Println(string(rune1[3])) //可以得到字符数

	//接口类型转具体类型
	//断言
	var inf interface{} = 100
	var infStruct interface{} = user{Name: "nick"}
	i, ok := inf.(int)
	fmt.Println(i, ok)
	u, ok := infStruct.(user)
	fmt.Println(u, ok)

	//时间类型字符串
	var t time.Time
	t = time.Now()
	timeStr := t.Format("2006-01-02 15:04:05Z07:00")
	fmt.Println(timeStr)

	//字符串转时间

	t2, _ := time.Parse("2006-01-02 15:04:05Z07:00", timeStr) //忽略后面时区信息
	fmt.Println(t2)

	u1 := user{}
	uPtr := unsafe.Pointer(&u1)
	namePtr := unsafe.Pointer(uintptr(uPtr) + unsafe.Offsetof(u1.Name))
	*(*string)(namePtr) = "nick"
	fmt.Println(u1)
	//使用unsafe包进行指针操作时，需要非常小心，因为这种操作绕过了Go语言的类型系统和内存安全检查。如果操作不当，可能会导致程序崩溃或产生不可预测的行为。
	//在实际编程中，应尽量避免使用unsafe包，除非确实需要绕过Go的类型系统进行一些底层操作。

}
