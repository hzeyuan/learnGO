package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"go-DIYSpider/baseSpider"
	"testing"
)



func BenchmarkBaseSpider(b *testing.B) {
	b.N=1
	s :=baseSpider.NewBaseSpider()
	for i:=0;i<1;i++{
		s.Url("http://127.0.0.1:5000/goodsTypes/?_limit=1")
	}
	s.Run()

}


func TestBaseSpider(t *testing.T){
	s :=baseSpider.NewBaseSpider()
	for i:=0;i<10;i++{
		s.Url("http://127.0.0.1:5000/goodsTypes/?_limit=1")
	}
	s.Run()
}


func TestFetchPage(t *testing.T){
	resp, _, _ :=gorequest.New().Get("https://hzeyuan.cn/").End()
	doc,_:= goquery.NewDocumentFromReader(resp.Body)
	ele :=doc.Find("h2.post-title ").First()
	fmt.Println(ele.Text())
}