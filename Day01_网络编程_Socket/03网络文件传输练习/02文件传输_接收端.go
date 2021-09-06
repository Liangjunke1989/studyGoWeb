package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//创建Socket
	listener, err := net.Listen("tcp", "192.168.1.90:5566")
	if err != nil {
		fmt.Println("net.Listen()有误！", err.Error())
		return
	}
	defer listener.Close()

	//阻塞监听
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept()有误！", err.Error())
		return
	}

	//读取客户端数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read()有误！", err.Error())
		return
	}
	//在服务端显示读取到的数据：方便查看
	fmt.Println("从客户端获取的数据为：", string(buf[:n]))
	fileName := string(buf[:n])

	//会写数据
	str := "我是LJKHomeServer服务端，已经获取到了你发送的数据，这是我的回写数据！"
	_, err = conn.Write([]byte(str))
	if err != nil {
		fmt.Println("conn.Write()回写有误", err.Error())
		return
	}
	fmt.Println("ok,服务器回写成功！！！")
	//获取文件内容
	recvFile(conn, fileName)
}

func recvFile(conn net.Conn, fileName string) {
	//按照文件名创建新文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("文件创建有误！", err.Error())
		return
	}
	defer f.Close()

	//从网络中读数据，写入本地文件
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("接收文件完成！")
			return
		}
		//写入本地文件，读多少，写多少
		f.Write(buf[:n])
	}

}
