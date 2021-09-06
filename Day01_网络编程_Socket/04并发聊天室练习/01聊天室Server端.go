package main

import (
	"fmt"
	"net"
)

//创建服务器的监听IP和端口号
var serIp = "192.168.1.90"
var serPort = "5566"

//创建用户结构体
type Client struct {
	C    chan string
	Name string
	Addr string
}

//创建全局map,存储在线用户
var onlineMap map[string]Client

//创建全局channel，传递用户消息
var broadcastMessage = make(chan string)

func main() {
	/*
		1、主go程（服务器）
			负责监听、接收用户（客户端）连接请求，建立通信关系。同时启动相应的go程处理任务。
		2、处理用户连接的go程1：HandleConnect
			负责新上线用户的存储，用户消息读取、发送，用户改名、下线处理及超时处理。
			为了提高效率，同时给一个用户维护多个协程来并发处理上述任务。
		3、用户消息广播的go程2：Manager
			负责在线用户遍历，用户消息广播发送。需要与HandleConnect go程及用户子go程协作完成。
		4、go程间应用数据及通信
			Map:存储所有登陆聊天室的用户信息，key:用户的ip+port.   Value:Client结构体。
			Client结构体：包含成员：用户名Name，网络地址Addr(ip+port),发送消息的通道C（channel）
			通道message：协调并发go程间消息的传递。
	*/
	//获取服务器端的网络地址（ip+port）
	addr := fmt.Sprint(serIp, ":", serPort)
	//1、创建监听Socket
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("创建监听Socket有误！", err.Error())
		return
	}
	defer listener.Close()
	//创建Manager go程，管理map和全局channel
	go Manager()
	//2、循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("监听客户端连接请求有误！", err.Error())
			return
		}
		//处理数据请求
		//启动go程，来处理客户端数据请求
		go HandlerConnect(conn)
	}
}

func Manager() {
	//初始化 onlineMap
	onlineMap = make(map[string]Client)

	//监听全局channel中是否有数据,有数据存储至msg，无数据阻塞。
	for {
		msg := <-broadcastMessage
		//循环发送消息给 所有在线用户
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func HandlerConnect(conn net.Conn) { //处理客户端数据请求
	defer conn.Close()
	//获取用户网络地址
	clientAddr := conn.RemoteAddr().String()
	//创建新连接用户的结构体信息,默认用户名为ip+port
	client := Client{
		C:    make(chan string),
		Name: clientAddr,
		Addr: clientAddr,
	}
	//将新连接用户添加到用户map中
	onlineMap[clientAddr] = client
	//创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(client, conn)
	//发送 用户上线消息 到全局Channel中
	msg := fmt.Sprint("[", clientAddr, "]"+client.Name, "login,上线了！！！")
	broadcastMessage <- msg
	//保证 不退出
	for {

	}
}

//通过通道阻塞的方式，实现go程之间的通信和执行顺序的控制
func WriteMsgToClient(client Client, conn net.Conn) {
	//监听 用户自带channel上是否有消息
	for msg := range client.C {
		conn.Write([]byte(msg + "\n"))
	}
}
