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

var dk_wg = sync.WaitGroup{}

func main() {
	//爬取的目标网址
	url := "https://www.douyu.com/directory"
	//爬取整个网页，将整个网页的全部信息，保存在result中
	result, err := DK_HttpGet(url)
	DK_ErrFunc(err, "获取网页数据有误！")
	//解析正则表达式
	ret := regexp.MustCompile(`src="(?s:(.*?))">`)
	//提取每一张图片的url
	alls := ret.FindAllStringSubmatch(result, -1)
	fmt.Println(len(alls))
	dk_wg.Add(len(alls))
	for index, logoURL := range alls {
		go DK_SaveImage(index, logoURL[1])
	}
	dk_wg.Wait()
	fmt.Println("从斗鱼获取logo完成！")
}

func DK_SaveImage(index int, url string) {
	fName := strconv.Itoa(index) + ".png"
	fPath := "./src/Day03_爬虫/DirTest03/"
	fPathName := fPath + fName
	f, err := os.OpenFile(fPathName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	DK_ErrFunc(err, "打开文件有误！")
	defer f.Close() // 保存好一个文件，关闭一个文件。

	resp123, err := http.Get(url)
	DK_ErrFunc(err, "url打开有误！")
	fmt.Println(resp123)
	defer resp123.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp123.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		f.Write(buf[:n])
	}
	dk_wg.Done()
}

//DK_HttpGet 获取该网页中所有的内容，result返回
func DK_HttpGet(url string) (result string, err error) {
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
func DK_ErrFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err.Error())
		//os.Exit(1)       //将当前进程结束，0值代表正常结束。非零值代表非正常结束。
		return
	}
}
