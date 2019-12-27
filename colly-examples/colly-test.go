package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
	"io"
	"os"
	"regexp"
	"time"
)

func main() {
	t := time.Now()
	c := colly.NewCollector(func(c *colly.Collector) {
		extensions.RandomUserAgent(c) // 设置随机头
		//c.Async=true
		c.MaxDepth=1
		c.Async=true
	},
	)
	c2:=c.Clone()
	q, _ := queue.New(
		8, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//var needLink = regexp.MustCompile("")
		if ret,_:=regexp.MatchString("https://www\\.tektorg\\.ru/rosnefttkp/procedures/[0-9].*",e.Request.AbsoluteURL(link));ret==true{
			fmt.Println(e.Request.AbsoluteURL(link))
				c.Visit(e.Request.AbsoluteURL(link))
		}
		if ret,_:=regexp.MatchString("https://www\\.tektorg\\.ru/document\\.php\\?id=[0-9].*",e.Request.AbsoluteURL(link));ret==true{
			fmt.Println("下载,加入队列")
			//c.Visit(e.Request.AbsoluteURL(link))
			q.AddURL(e.Request.AbsoluteURL(link))
		}
	})
	c2.OnResponse(func(response *colly.Response) {

			fileName := response.Headers.Get("Content-Disposition")
			fmt.Println("下载 -->%s \n",fileName)
			length :=len(fileName)-20
			f, err := os.Create("./download/"+fileName[length:])
			if err != nil {
				panic(err)
			}
			io.Copy(f, bytes.NewReader(response.Body))

	})

	c.Visit("https://www.tektorg.ru/procedures?q=%D0%A1%D1%82%D1%80%D0%BE%D0%B8%D1%82%D0%B5%D0%BB%D1%8C%D1%81%D1%82%D0%B2%D0%BE&dpfrom=01.01.2019&dpto=31.12.2019&startpricefrom=50000000&limit=500")
	//c.Wait()
	c.Wait()
	q.Run(c2)
	c2.Wait()
	fmt.Printf("花费时间:%s",time.Since(t))
}

