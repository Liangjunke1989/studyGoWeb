package main

import (
	"fmt"
	"net"
	"strings"
	"time"
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
var broadcastMsgChan = make(chan string)

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
		msg := <-broadcastMsgChan
		//循环发送消息给 所有在线用户
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func HandlerConnect(conn net.Conn) { //处理客户端数据请求
	defer conn.Close()
	//创建channel，判断用户是否活跃
	hasData := make(chan bool)
	//获取用户网络地址
	clientAddr := conn.RemoteAddr().String()
	//创建新连接用户的结构体信息,默认用户名为ip+port
	_client := Client{
		C:    make(chan string),
		Name: clientAddr,
		Addr: clientAddr,
	}
	//将新连接用户添加到用户map中
	onlineMap[clientAddr] = _client
	//创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(_client, conn)
	//发送 用户上线消息 到全局Channel中
	broadcastMsgChan <- makeBroadcastMsg(_client, "login!!!客户端已经成功登陆！！！")
	//创建一个channel，用户判断退出状态
	isQuit := make(chan bool)
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("%s客户端已退出！", _client.Name)
				return
			}
			if err != nil {
				fmt.Println("从客户端读取数据有误！", err.Error())
				return
			}
			clientMsg := string(buf[:n-1])
			//提取在线用户列表：判断是否为who命令，如果是，组织显示信息，写入到socket中
			if clientMsg == "who" && len(clientMsg) == 3 {
				conn.Write([]byte("在线用户列表：\n"))
				//通过遍历map获取在线用户
				for _, clientUser := range onlineMap {
					clientUserInfo := fmt.Sprint(clientUser.Addr + ":" + clientUser.Name)
					conn.Write([]byte(clientUserInfo))
				}
				//判断用户发送了改名命令，返回更新成功的提示信息
			} else if len(clientMsg) >= 8 && clientMsg[:6] == "rename" { //至少有一个字符
				newName := strings.Split(clientMsg, "|")[1]
				_client.Name = newName
				onlineMap[clientAddr] = _client //更新 onlineMap
				conn.Write([]byte("用户重命名，更新成功！！！"))
			} else {
				//将读到的用户消息，广播给所有用户。将读取的数据写入到broadcastMessage
				broadcastMsgChan <- makeBroadcastMsg(_client, clientMsg)
			}
			hasData <- true //如果用户执行了上面的操作，说明用户活跃中，可以向标识位写入操作，让通道畅通。
		}
	}()
	//保证 不退出
	for {
		//监听channel上的数据流动
		select {
		case <-isQuit:
			delete(onlineMap, clientAddr)                                   //将用户从online中移除
			broadcastMsgChan <- makeBroadcastMsg(_client, "logout,客户端退出了！") //写入用户退出消息到广播channel
			return
		case <-time.After(time.Second * 15):
			delete(onlineMap, clientAddr)                                      //将用户从online中移除
			broadcastMsgChan <- makeBroadcastMsg(_client, "logout,客户端超时，退出了！") //写入用户退出消息到广播channel
			return
		case <-hasData: //什么都不做，目的是,说明用户活跃中，让程序不退出
		}
	}
}

//处理创建广播消息
func makeBroadcastMsg(client Client, msgContent string) string {
	return fmt.Sprint("[", client.Addr, "]"+client.Name, ": \n", msgContent)
}

//通过通道阻塞的方式，实现go程之间的通信和执行顺序的控制
func WriteMsgToClient(client Client, conn net.Conn) {
	//监听 用户自带channel上是否有消息
	for msg := range client.C {
		conn.Write([]byte(msg + "\n"))
	}
}
