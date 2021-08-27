package main

import (
	"fmt"
	"net"
)

var IP = "localhost"
var Port = "5566"

func main() {
	/*
		Server端：
			Listen函数：
				func Listen(network,address string)(Listener,error)
					network:选用的协议TCP/UDP
					address:IP地址+端口号
			Listener接口:
				type Listener interface{
					Accept()(Conn,error)
					Close() error
					Addr() Addr
				}
			Conn接口：
				type Conn interface{
					Read(b []byte)(n int, err error)
					Write(b []byte)(n int, err error)
					Close() error
					LocalAddr() Addr
					RemoteAddr() Addr
					SetDeadline(t time.Time) error
					SetReadDeadline(t time.Time) error
					SetWriteDeadline(t time.Time) error
			    }
	*/
	addrStr := fmt.Sprint(IP, ":", Port)
	//指定服务器的通信协议，IP地址，端口号.    (创建一个用于监听的socket)     -----创建的是服务器的监听IP和端口
	listener, err := net.Listen("tcp", addrStr)
	if err != nil {
		fmt.Println("net.Listen()出错了！", err.Error())
		return
	}
	defer listener.Close()

	fmt.Printf("服务器等待客户端建立连接...")
	//阻塞监听客户端连接请求.   成功建立连接，返回用于通信的Socket
	conn, err := listener.Accept()

	if err != nil {
		fmt.Println("listener.Accept()出错了！", err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端成功建立连接了.......")
	//读取客户端发送来的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read()发生错误！", err.Error())
		return
	}
	//将读取的数据进行处理
	//conn.Write(buf[:n])
	fmt.Println("从客户端获取到的数据为：", string(buf[:n]))

	//服务器回写数据
	str1 := "来自服务端的一条回写数据！"
	write, _ := conn.Write([]byte(str1))
	fmt.Println(write)
}
