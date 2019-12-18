package main

import (
	"fmt"
	"go-DIYSpider/baseSpider"
	"time"
)

func main() {
	urlList := []string{
		"https://hzeyuan.cn/",
		"https://hzeyuan.cn/",
	}
	for i:=0;i<2;i++{
		urlList = append(urlList,"https://hzeyuan.cn/")
	}
	t:=time.Now()
	baseSpider.NewBaseSpider().Run(urlList)
	fmt.Println(time.Since(t))
}
