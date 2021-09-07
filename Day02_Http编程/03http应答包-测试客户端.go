package main

import (
	"fmt"
	"net"
	"os"
)

//出错函数封装
func errFunc2(err error, info string) {
	if err != nil {
		fmt.Println(info, err.Error())
		os.Exit(1) //将当前进程结束，0值代表正常结束。非零值代表非正常结束。
	}
}

//伪装 浏览器
func main() {
	conn, err := net.Dial("tcp", "192.168.1.90:5566")
	errFunc2(err, "连接服务器错误！请查看服务器是否开启！")
	defer conn.Close()
	httpRequest := "GET /LJKD HTTP/1.1\r\n" +
		"Host:192.168.3.129:10001\r\n" +
		"\r\n"
	conn.Write([]byte(httpRequest))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc2(err, "客户端读取数据有误！")
	fmt.Println(string(buf[:n]))
}
