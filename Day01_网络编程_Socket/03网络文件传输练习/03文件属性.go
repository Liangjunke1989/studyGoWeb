package main

import (
	"fmt"
	"os"
)

func main() {
	/*
			获取命令行参数：
				在main函数启动时，向整个程序传参。
			语法：
				go run xxx.go argv1 argv2 argv3 argv4...
					xxx.go:第0个参数
		             argv1：第1个参数
			 		 argv2：第2个参数
						...
			获取文件属性
				fileInfo, err :=os.Stat(文件访问的绝对路径)

				fileInfo.Name(), 文件目录项信息中的文件名，可以把文件名单独提取出来
		    list:=os.Args
		    fmt.Println(list)
	*/
	list := os.Args //获取命令行参数
	if len(list) != 2 {
		fmt.Println("格式为：go run xxx.go 文件名") //文件名为第二个参数 进行传递
		return
	}
	fmt.Println(list)
	//提取文件名
	filePath := list[1]
	//获取文件属性
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("文件路径获取有误", err.Error())
		return
	}
	fmt.Println("文件名：", fileInfo.Name()) //文件目录项信息中的文件名，可以把文件名单独提取出来
	fmt.Println("文件大写：", fileInfo.Size())
}
