package _case

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stderr)
	//设置log格式
}

func LogCase() {
	var a, b = -1, -2
	_, err := sum(a, b)
	if err != nil {
		log.Println(err)
	}
	log.Printf("a:%d,b:%d,两数求和出现错误:sum: %s \n", a, b, err.Error())
	//log.Fatalf("a:%d,b:%d,两数求和出现错误:sum: %s \n", a, b, err.Error()
	//fatalf函数会输出日志信息并终止程序
}
