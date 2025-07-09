package _case

import "fmt"

/*func NewCase() {
	//通过new函数可以创建任意类型并返回指针
	mpPtr := new(map[string]*user)
	//(*mpPtr)["A"] = &user{}
	if mpPtr == nil {
		fmt.Println("map地址为空")
	}
	if *mpPtr == nil {
		fmt.Println("map值为空")
	}
	slicePtr := new([]user)
	if *slicePtr == nil {
		fmt.Println("slice内容为空")
	}
	*slicePtr = append(*slicePtr, user{Name: "jack", Age: 18})
	userPtr := new(user)
	strPtr := new(string)
	fmt.Println(mpPtr, slicePtr, userPtr, strPtr)
}
*/

/*func MakeCase() {
	//初始化切片，并且设置长度和容量
	slice := make([]int, 10, 20)
	slice[0] = 10
	//初始化集合，并设置集合初始大小
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	//初始化通道，设置通道的读写方向和缓冲大小
	ch := make(chan int, 10)    //可读可写
	ch1 := make(chan<- int, 10) //只写
	ch2 := make(<-chan int, 10) //只读
	fmt.Println(slice, mp, ch, ch1, ch2)

}
*/

func SliceCase() {
	/*var slice []int
	slice = []int{1, 2, 3, 4, 5, 6}
	slice1 := make([]int, 10)
	slice1[1] = 10
	fmt.Println(slice, slice1)
	//切片截取
	slice2 := slice1[1:3]
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))
	//切片的附加
	slice2 = append(slice2, 10)
	//如果添加数据的长度超过容量，则会扩容，采用新数组，已经不是原来的数组了
	*/

	//集合是无序的
	//map存储的是一个一个键值对
	//Go语言中的映射是无序的，遍历映射时键值对的顺序是不确定的。
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	mp["B"] = "b"
	mp["C"] = "c"
	mp["D"] = "d"

	fmt.Println(mp)
	for k, v := range mp {
		fmt.Println(k, v)
	}
	//删除集合中的元素
	delete(mp, "A")
	fmt.Println(mp)

}

//make只能用于slice map channel三个，不能用于其他类型，返回的是原始类型，不是指针
/*为什么针对slice, map和chan类型专门定义一个make函数？
答案：这是因为slice, map和chan的底层结构上要求在使用slice，map和chan的时候必须初始化，如果不初始化，那slice，map和chan的值就是零值，也就是nil。
 我们知道：map如果是nil，是不能往map插入元素的，插入元素会引发panic chan如果是nil，往chan发送数据或者从chan接收数据都会阻塞 slice会有点特殊，
理论上slice如果是nil，也是没法用的。但是append函数处理了nil slice的情况，可以调用append函数对nil slice做扩容。但是我们使用slice，总是会希望可以自定义长度或者容量，这个时候就需要用到make。*/
