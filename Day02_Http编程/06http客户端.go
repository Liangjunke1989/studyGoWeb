package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//获取服务器的响应包内容
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("请求错误！", err.Error())
		return
	}
	defer resp.Body.Close()

	//简单查看响应包
	fmt.Println("Header：", resp.Header)
	fmt.Println("Status：", resp.Status)
	fmt.Println("StatusCode：", resp.StatusCode)
	fmt.Println("Proto：", resp.Proto)
	fmt.Println("Body：", resp.Body) //获取响应体数据
	//fmt.Printf("Body格式%T", resp.Body) //获取响应体数据
	buf := make([]byte, 256)
	var result string
	//循环读取数据
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			println("读取完成！")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("数据读取有误！", err.Error())
			return
		}
		result += string(buf[:n])
	}
	fmt.Println(result)
}
