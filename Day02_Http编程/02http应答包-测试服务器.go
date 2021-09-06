package main

import "net/http"

func main() {
	/*
		1、http应答包
			http响应报文格式
		2、注册回调行数：
			该回调函数会在服务器被访问时，自动被调用。
	*/
	//创建web服务器
	//01、注册回调行数
	http.HandleFunc("/LJK3D", LJK_Handler)
	//02、绑定服务器监听地址
	http.ListenAndServe("192.168.1.90:5566", nil) //调用自带的handler
}

// LJK_Handler 在服务器被访问时，自动被调用。
func LJK_Handler(w http.ResponseWriter, r *http.Request) {
	//w: 响应数据（回写给客户端的数据）
	//r:请求数据（客户端发送来的请求）
	w.Write([]byte("LJK的第一条goWeb响应数据：Hello World！！！"))
}
