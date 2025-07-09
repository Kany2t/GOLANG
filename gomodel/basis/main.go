package main

//init函数是一个特殊的函数，被称为初始化函数，包引用时优先调用
import (
	"fmt"
	p1 "gomodel/basis/package1"
	//可以取别名，写在路径前面
	_ "gomodel/basis/package2"
	//依赖引用
)

// 先调用方法
func init() {
	fmt.Println("hello world2")
}

// main函数是整个函数的入口
func main() {
	fmt.Println("hello world")

	p1.F1()
	//包名可以和引用目录不一致（包名和文件夹名可以不一致）
	//golang公有成员与私有成员通过成员标识符的首字母大小写来区分
	//首字母大写表示公有，首字母小写表示私有成员，只有两种情况
	//init函数先于入口函数执行

}
