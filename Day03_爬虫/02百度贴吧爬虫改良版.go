package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	/*
		横向爬取：
			在爬取的网站页面中，以"页"为单位，找寻该网站分页器规律。
			一页一页的爬取网站数据信息。
			大多数网站，采用分页管理模式。----针对这类网站，首先要确立横向爬取方法。
	*/
	//01、指定爬取起始页、终止页
	var start, end int
	fmt.Print("请输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>=start）:")
	fmt.Scan(&end)
	//02、爬取页面操作方法
	DataCrawlerWorking2(start, end)
	//03、提示操作结束
	fmt.Println("爬取成功了！主go程结束...")
}

func DataCrawlerWorking2(start int, end int) {
	fmt.Printf("正在爬取第%d到%d页...\n", start, end)
	//pageChan:=make(chan int)

	sumThread := end - start + 1
	wg.Add(sumThread)
	//循环爬取每一页数据
	for i := start; i <= end; i++ {
		go CrawlingSinglePage2(i)
	}
	wg.Wait()
}
func LJKErrFunc2(err error, info string) {
	if err != nil {
		fmt.Println(info, err.Error())
		//os.Exit(1)       //将当前进程结束，0值代表正常结束。非零值代表非正常结束。
		return
	}
}

// CrawlingSinglePage2 爬取单个页面的函数
func CrawlingSinglePage2(i int) {
	url := "https://tieba.baidu.com/f?kw=unity&ie=utf-8&pn=" + strconv.Itoa((i-1)*50) // Itoa() 将整型转成字符串
	result, err := HttpGet2(url)
	//if err!=nil {
	//	fmt.Println("httpGet有误！",err.Error())
	//	continue
	//}
	LJKErrFunc2(err, "httpGet有误！")
	//fmt.Println("result=",result) //测试

	//将读取到的整网页数据，保存成一个文件
	//os.Create("第 " + strconv.Itoa(i) + " 页" + ".html") //弃用
	fName := "第 " + strconv.Itoa(i) + " 页" + ".html"
	fPath := "./src/Day03_爬虫/DirTest/"
	fPathName := fPath + fName
	f2, err := os.OpenFile(fPathName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	LJKErrFunc2(err, "打开文件有误！")
	f2.WriteString(result)
	f2.Close() // 保存好一个文件，关闭一个文件。
	wg.Done()
}

func HttpGet2(url string) (result string, err error) {
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
