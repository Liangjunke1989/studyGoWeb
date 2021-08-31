package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//01、创建监听套接字
	listener, err := net.Listen("tcp", "192.168.1.90:5566")
	if err != nil {
		fmt.Println("net.Listen()出错了！", err.Error())
		return
	}
	defer listener.Close()
	//02、阻塞监听客户端连接请求.
	for {
		fmt.Println("服务器等待客户端连接......")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept()出错了！", err.Error())
			return
		}
		//03、具体完成服务器和客户端的数据通信
		go HandlerConnect(conn)
	}

}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	//获取客户端的网络地址
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接！！！")
	//循环读取客户端发送数据
	buf := make([]byte, 4096)
	for {
		//读取数据
		n, err := conn.Read(buf)
		if "exit\n" == string(buf[:n]) {
			fmt.Println("服务器接收到的mac客户端请求，可以正常关闭")
			return
		}
		if "exit\r\n" == string(buf[:n]) {
			fmt.Println("服务器接收到的win客户端请求，可以正常关闭")
			return
		}
		if n == 0 { //从已经关闭的channel读去数据，是可以读取数据的，值为0
			fmt.Println("服务器检测到客户端已关闭，断开连接")
			return
		}
		if err != nil {
			fmt.Println("conn.Read()错误！", err.Error())
			return
		}
		fmt.Println("服务器读到的数据为", string(buf[:n]))
		//回写数据
		//小写转大写，回发
		upperStr := strings.ToUpper(string(buf[:n]))
		conn.Write([]byte(upperStr))
	}
}
