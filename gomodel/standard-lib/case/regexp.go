package _case

import (
	"fmt"
	"regexp"
)

func RegexpCase() {
	//构建正则表达式对象
	reg := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	//判断字符串是否匹配正则表达式
	fmt.Println(reg.MatchString("abc[1234]"))
	//从给定的字符串查找符合条件的字符串
	bytes := reg.FindAll([]byte("efg[456]"), -1) //-1：这是查找的次数限制。-1表示查找所有匹配项，而不仅仅是第一个匹配项。
	fmt.Println(string(bytes[0]))

}
