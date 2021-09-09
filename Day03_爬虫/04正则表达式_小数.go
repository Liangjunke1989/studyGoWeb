package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "3.14 123.123 .68 haha 1.0 abc 7. ab.3 66.6 123."
	//01、解析编译正则表达式
	ret := regexp.MustCompile(`[0-9]+\.[0-9]+`)

	//02、提取需要的信息
	alls := ret.FindAllStringSubmatch(str, -1)
	fmt.Println(alls)
}
