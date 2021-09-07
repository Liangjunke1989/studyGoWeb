package main

import (
	"fmt"
	"net/http"
)

func main() {
	//01、注册回调函数
	http.HandleFunc("/3DTest", ljkHandler)
	//02、绑定服务器监听地址
	err := http.ListenAndServe("192.168.1.90:5566", nil) //调用默认的回调函数
	if err != nil {
		fmt.Println("web服务器创建监听地址有误！", err.Error())
	}
}

//回调函数
func ljkHandler(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter 响应数据
	n, err := w.Write([]byte("目前写给所有客户端的内容，即响应数据为：这个LJK的第一个goWeb服务器"))
	if err != nil {
		fmt.Println("会写数据有误！", err.Error())
	} else {
		fmt.Println("回写的数据个数为：", n)
	}
	//r *http.Request 请求数据
	fmt.Println("Header：", r.Header)
	fmt.Println("URL：", r.URL)
	fmt.Println("Method：", r.Method)
	fmt.Println("Host：", r.Host)
	fmt.Println("RemoteAddr：", r.RemoteAddr)
	fmt.Println("Body：", r.Body) //获取请求体数据

}
