package main

import (
	"fmt"
	"net"
)

var IPServer = "localhost"
var PortServer = "5566"

func main() {
	/*
		Client客户端：
		net.Dial()函数：
			net.Dial(network,address string)(Conn,error){
				var d Dialer
				return d.Dial(network,address)
		}
	*/
	//01、指定服务器IP+Port, 创建通信套接字Socket
	addrStr := fmt.Sprint(IPServer, ":", PortServer)
	conn, err := net.Dial("tcp", addrStr)
	if err != nil {
		fmt.Println("net.Dial()拨号有误！！", err.Error())
		return
	}
	defer conn.Close()
	//02、客户端写数据给服务器
	conn.Write([]byte("我是LJK客户端......"))
	//03、接收服务器回发的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read()读取有误！！", err.Error())
	}
	fmt.Println("拿到的服务器端的回写数据为：", string(buf[:n]))
}
