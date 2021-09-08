package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//01、指定爬取起始页、终止页
	var start, end int
	fmt.Print("请输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>=start）:")
	fmt.Scan(&end)
	//02、爬取页面操作方法
	DataCrawlerWorking(start, end)
}

func DataCrawlerWorking(start int, end int) {
	fmt.Printf("正在爬取第%d到%d页...\n", start, end)
	//循环爬取每一页数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=unity&ie=utf-8&pn=" + strconv.Itoa((i-1)*50) // Itoa() 将整型转成字符串
		result, err := HttpGet(url)
		//if err!=nil {
		//	fmt.Println("httpGet有误！",err.Error())
		//	continue
		//}
		LJKErrFunc(err, "httpGet有误！")
		//fmt.Println("result=",result) //测试

		//将读取到的整网页数据，保存成一个文件
		//os.Create("第 " + strconv.Itoa(i) + " 页" + ".html") //弃用
		fName := "第 " + strconv.Itoa(i) + " 页" + ".html"
		fPath := "./src/Day03_爬虫/DirTest/"
		fPathName := fPath + fName
		f2, err := os.OpenFile(fPathName, os.O_RDWR|os.O_CREATE, os.ModePerm)
		LJKErrFunc(err, "打开文件有误！")
		f2.WriteString(result)
		f2.Close() // 保存好一个文件，关闭一个文件。
	}
}
func LJKErrFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err.Error())
		//os.Exit(1)       //将当前进程结束，0值代表正常结束。非零值代表非正常结束。
		return
	}
}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //将封装函数内部的错误，传给调用者
		return
	}
	defer resp.Body.Close()
	//循环读取网页数据，传给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("文件读取完毕！")
			break
		}
		if err2 != nil && err2 != io.EOF {
			fmt.Println("文件读取出错了！", err2.Error())
			err = err2
			return
		}
		//累加每一次循环读到的buf数据，存入result
		result += string(buf[:n])
	}
	return result, err
}
