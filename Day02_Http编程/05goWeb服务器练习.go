package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	//注册回调函数
	http.HandleFunc("/3d/", myHandler)
	//绑定监听地址
	http.ListenAndServe("192.168.1.90:5566", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("客户端请求：", r.URL)
	OpenSendFile(r.URL.String(), w)
}

func OpenSendFile(fName string, w http.ResponseWriter) {
	//创建带路径的文件名
	pathFileName := "F:/test" + fName
	f, err := os.Open(pathFileName)
	if err != nil {
		fmt.Println("打开文件失败！！", err.Error())
		w.Write([]byte("服务器打开文件失败！！+ html404错误页面！"))
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf) //从服务器本地将文件内容读取
		if n == 0 {
			fmt.Println("文件读取完毕！")
			return
		}
		if err != nil {
			fmt.Println("文件读取有误！", err.Error())
			return
		}
		w.Write(buf[:n]) //写到 客户端（浏览器）上
	}
}
