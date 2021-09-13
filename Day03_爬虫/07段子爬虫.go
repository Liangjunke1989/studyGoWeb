package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var wg001 = sync.WaitGroup{}

func main() {
	/*
		段子爬虫：
			1、先提取URL
				获取网页中每一个段子的URL
	*/

	//01、指定爬取起始页、终止页
	var start, end int
	fmt.Print("请输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>=start）:")
	fmt.Scan(&end)
	//02、爬取页面操作方法
	LJK_DataCrawlerWorking(start, end)
	//03、提示操作结束
	fmt.Println("爬取成功了！主go程结束...")
}

func LJK_DataCrawlerWorking(start int, end int) {
	fmt.Printf("正在爬取第%d到%d页...\n", start, end)
	wg001.Add(end - start + 1)
	//循环爬取每一页数据
	for i := start; i <= end; i++ {
		go LJK_CrawlingSinglePage(i)
	}
	wg001.Wait()
}

//LJK_CrawlingSinglePage 抓取一个网页，带有10个段子---10个url
func LJK_CrawlingSinglePage(i int) {
	//拼接url
	url := "http://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	//封装函数，获取段子的url
	result, err := LJK_HttpGet(url)
	LJK_ErrFunc(err, "HttpGet错误！")

	//解析编译正则表达式
	ret := regexp.MustCompile(`<h1 class="dp-b'><a href="(?s:(.*?))"`)
	//提取需要信息
	alls := ret.FindAllStringSubmatch(result, -1)

	//创建用户存储title、content的切片，初始容量为0
	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)
	for _, jokeURL := range alls {
		fmt.Println("jokeURL:", jokeURL[1]) //此时获取的是URL链接
		t, cont, err := LJK_CrawlingJokeURLPage(jokeURL[1])
		LJK_ErrFunc(err, "爬取jokeURL错误！")
		//往本地文件 写入保存
		fileTitle = append(fileTitle, t)        //将t追加到切片末尾
		fileContent = append(fileContent, cont) //将t追加到切片末尾
	}
	wg001.Done()
	//LJK_Savejoke2File(fileTitle[],fileContent)
}

func LJK_Savejoke2File(index int, title, cont []string) {
	fName := "第 " + strconv.Itoa(index) + " 页" + ".html"
	fPath := "./src/Day03_爬虫/DirTest03/"
	fPathName := fPath + fName
	f, err := os.OpenFile(fPathName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	LJKErrFunc2(err, "打开文件有误！")
	for i := 0; i < len(title); i++ {
		f.WriteString(title[i] + "\n" + cont[i] + "\n")
		f.WriteString("---------------------\n")
	}
	f.Close() // 保存好一个文件，关闭一个文件。

}

// LJK_CrawlingJokeURLPage 爬取一个笑话的title和content
func LJK_CrawlingJokeURLPage(url string) (title, content string, err error) {
	result, err1 := LJK_HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}
	//编译解析正则表达式--title
	ret1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	//提取title
	alls01 := ret1.FindAllStringSubmatch(result, 1) //有2处，取第一个
	for _, tempTitle := range alls01 {
		title = tempTitle[1]
		title = strings.Replace(title, " ", "", -1)
		break
	}

	//编译解析正则表达式--content
	ret2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	//提取title
	alls02 := ret2.FindAllStringSubmatch(result, 1) //有2处，取第一个
	for _, tempTitle := range alls02 {
		content = tempTitle[1]
		content = strings.Replace(content, " ", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "&nbsp;", "", -1)
		break
	}
	return title, content, err
}

//LJK_HttpGet 获取一个网页所有的内容
func LJK_HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return result, err
}
func LJK_ErrFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err.Error())
		//os.Exit(1)       //将当前进程结束，0值代表正常结束。非零值代表非正常结束。
		return
	}
}
