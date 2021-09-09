package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
		正则表达式：
			字符
			数量限定符
			转义字符
	*/
	str := "abc a7c mfc cat 8ca azc cba"
	//01、解析编译正则表达式
	ret := regexp.MustCompile(`a.c`) //参数为正则表达式，返回值为结构体
	//可以使用``:表示使用原生字符串

	//提取需要信息
	alls := ret.FindAllStringSubmatch(str, -1) //返回二维数组类型
	fmt.Println(alls)
}
