package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
)

var wg2 = sync.WaitGroup{}

func main() {
	/*
		纵向爬取：
			在一个页面内，按不同的"条目"为单位。找寻各条目之间的规律。
			一条一条的爬取一个网页中的数据信息。也就是同时爬取一饿页面内不同类别的数据。
	*/

	//01、指定爬取起始页、终止页
	var start, end int
	fmt.Print("请输入爬取豆瓣电影Top250的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取豆瓣电影Top250的终止页（>=start）:")
	fmt.Scan(&end)
	//02、爬取页面操作方法
	ToCrawlDataWork(start, end)
	fmt.Println("main...over...")
}

func ToCrawlDataWork(start int, end int) {
	fmt.Printf("正在爬取第%d到%d页...\n", start, end)
	wg2.Add(end - start + 1)
	for i := start; i <= end; i++ {
		go CrawlingSinglePageData(i)
	}
	wg2.Wait()
}
func CrawlingSinglePageData(index int) {
	//获取url
	//url :="https://tieba.baidu.com/f?kw=unity&ie=utf-8&pn=" + strconv.Itoa((index-1)*50)
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((index-1)*25) //+"&filter="
	//爬取url对应页面，封装httpGet3()函数
	result, err := HttpGetDB(url)
	LJKErrFunc3(err, "网页文件数据获取有误！HttpGet3()出错")
	//测试查看result的结果
	//fmt.Println(result)
	//01、解析编译正则表达式--
	ret := regexp.MustCompile(`<img width=“100” alt=“(.*?)"`)
	//02、提取需要信息
	filmNameArr := ret.FindAllStringSubmatch(result, -1)
	//for _, name := range filmNameArr {
	//	fmt.Println("name:",name[1])
	//}
	//将提取的有用信息封装到文件中
	Save2File(index, filmNameArr)
	wg2.Done()
}

func Save2File(index int, filmNameArr [][]string) {
	fName := "第 " + strconv.Itoa(index) + " 页" + ".html"
	fPath := "./src/Day03_爬虫/DirTest/"
	fPathName := fPath + fName
	f2, err := os.OpenFile(fPathName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	LJKErrFunc3(err, "打开文件有误！")
	//往当前文件中写数据
	f2.WriteString("电影名称") //先打印抬头
	n := len(filmNameArr)  // 得到条目数
	for i := 0; i < n; i++ {
		f2.WriteString(filmNameArr[i][1])
	}
	f2.Close() // 保存好一个文件，关闭一个文件。
}

func HttpGetDB(url string) (result string, err error) {
	resp2, err2 := http.Get(url)
	if err2 != nil {
		if err2 != nil {
			err = err2
			return
		}
	}
	defer resp2.Body.Close()
	buf := make([]byte, 4096)
	//循环爬取网页数据
	for {
		n, err3 := resp2.Body.Read(buf)
		if n == 0 {
			fmt.Println("文件读取完毕！")
			return
		}
		if err3 != nil && err3 != io.EOF {
			fmt.Println("文件读取出错了！", err3.Error())
			err = err3
			return
		}
		result += string(buf[:n])
	}
	return result, err
}
func LJKErrFunc3(err error, info string) {
	if err != nil {
		fmt.Println(info, err.Error())
		//os.Exit(1)       //将当前进程结束，0值代表正常结束。非零值代表非正常结束。
		return
	}
}
