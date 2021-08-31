package main

import (
	"fmt"
	"net"
)

func main() {
	//组织一个 "udp地址结构"
	serverAddr, _ := net.ResolveUDPAddr("udp", "192.168.1.90:5566")
	//创建用于通信的Socket
	udpConn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("服务器创建监听地址出错了！", err.Error())
		return
	}
	defer udpConn.Close()
	buf := make([]byte, 4096)
	//返回3个值，分别是（读取到的字节数、客户端的地址、err）
	n, clientAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("udpConn读取数据有误！", err.Error())
		return
	}
	//模拟处理数据
	fmt.Println("服务器从", clientAddr, "读取到的数据为：", string(buf[:n]))
}
