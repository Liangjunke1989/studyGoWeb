package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动发起连接请求
	conn, err := net.Dial("tcp", "192.168.1.90:5566")
	if err != nil {
		fmt.Println("客户端拨号失败，请检查服务器是否开启！", err.Error())
		return
	}
	fmt.Println("客户端与服务器已经建立连接！")
	defer conn.Close()
	//获取用户的键盘数据stdin，将用户的键盘数据发送给服务器   //stdin和scan的区别？  scan不能有空格和会车，遇到空格完成结束了！
	go func() {
		for {
			str := make([]byte, 4096)
			n, err2 := os.Stdin.Read(str)
			if err2 != nil {
				fmt.Println("os.Stdin.Read(),从键盘读取错误！", err2.Error())
				continue
			}
			//写给服务器（读多少，发送多少）
			conn.Write(str[:n])
		}
	}()
	//回显服务器回发的数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("服务器已经主动关闭了！！！，客户端也关闭")
			return
		}
		if err != nil {
			fmt.Println("从服务器读取数据有误！", err.Error())
			return
		}
		fmt.Println("客户端从服务器读取到的回写数据为：", string(buf[:n]))
	}
}
