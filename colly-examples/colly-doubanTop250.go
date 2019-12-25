package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"regexp"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	number := 1
	c := colly.NewCollector(func(c *colly.Collector) {
		extensions.RandomUserAgent(c) // 设置随机头
		c.Async=true
	},
		//过滤url,去除不是https://movie.douban.com/top250?start=0&filter= 的url
		colly.URLFilters(
			regexp.MustCompile("^(https://movie\\.douban\\.com/top250)\\?start=[0-9].*&filter="),
		),
	) // 创建收集器
	// 响应的格式为HTML,提取页面中的链接
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Printf("find link: %s\n", e.Request.AbsoluteURL(link))
		c.Visit(e.Request.AbsoluteURL(link))
	})
	// 获取电影信息
	c.OnHTML("div.info", func(e *colly.HTMLElement) {
		e.DOM.Each(func(i int, selection *goquery.Selection) {
			movies := selection.Find("span.title").First().Text()
			director := strings.Join(strings.Fields(selection.Find("div.bd p").First().Text()), " ")
			quote := selection.Find("p.quote span.inq").Text()
			fmt.Printf("%d --> %s:%s %s\n", number, movies, director, quote)
			number += 1
		})
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})
	c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Wait()
	fmt.Printf("花费时间:%s",time.Since(t))
}
