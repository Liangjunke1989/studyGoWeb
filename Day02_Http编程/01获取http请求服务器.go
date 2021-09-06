package main

import (
	"fmt"
	"net"
	"os"
)

//出错函数封装
func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err.Error())
		//return                        //返回当前函数调用
		//runtime.Goexit()              //结束当前go程
		//break                         //跳出当前循环
		os.Exit(1) //将当前进程结束，0值代表正常结束。非零值代表非正常结束。
	}
}
func main() {
	/*
		封装出错函数errFunc(err error,info string)
		http请求服务器
	*/
	listener, err := net.Listen("tcp", "192.168.1.90:5566")
	errFunc(err, "服务器创建监听Socket有误！")
	defer listener.Close()
	conn, err := listener.Accept()
	errFunc(err, "连接服务器有误！")
	defer conn.Close()
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc(err, "服务器读取数据有误！")
	fmt.Printf("|%s|\n", string(buf[:n]))
}
