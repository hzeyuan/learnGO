package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"io"
	"net/url"
	"os"
	"strings"
	"time"
)

//todo:使用colly爬取https://www.quanjing.com中的卡通图片

/*
1.先F12,观察https://www.quanjing.com/search.aspx?q=%E5%8D%A1%E9%80%9A#%E5%8D%A1%E9%80%9A||1|1000|3|2||||||这个网站
2.会发现其实这是一个通过json来加载数据的！通过这个URL:https://www.quanjing.com/Handler/SearchUrl.ashx
3. 我们在看看都传递了哪些参数，照着来就好啦。
*/

func main(){
	t :=time.Now()
	c :=colly.NewCollector(func(collector *colly.Collector) {
		collector.Async=true
		extensions.RandomUserAgent(collector)
	})
	imageC :=c.Clone()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie","BIGipServerPools_Web_ssl=2135533760.47873.0000; Hm_lvt_c01558ab05fd344e898880e9fc1b65c4=1577432018; qimo_seosource_578c8dc0-6fab-11e8-ab7a-fda8d0606763=%E7%BB%94%E6%AC%8F%E5%94%B4; qimo_seokeywords_578c8dc0-6fab-11e8-ab7a-fda8d0606763=; accessId=578c8dc0-6fab-11e8-ab7a-fda8d0606763; pageViewNum=3; Hm_lpvt_c01558ab05fd344e898880e9fc1b65c4=1577432866")
		r.Headers.Add("referer", "https://www.quanjing.com/search.aspx?q=%E5%8D%A1%E9%80%9A")
		r.Headers.Add("sec-fetch-mode", "cors")
		r.Headers.Add("sec-fetch-site", "same-origin")
		r.Headers.Add("accept", "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01")
		r.Headers.Add("accept-encoding", "gzip, deflate, br")
		r.Headers.Add("accept-language", "en,zh-CN;q=0.9,zh;q=0.8")
		r.Headers.Add("X-Requested-With", "XMLHttpRequest")
	})
	c.OnResponse(func(r *colly.Response) {
		var f interface{}
		if err := json.Unmarshal(r.Body[13:len(r.Body)-1], &f);err!=nil{
			panic(err)
		}
		imgList := f.(map[string]interface{})["imglist"]
		for k,img :=range imgList.([]interface{}){
			url :=img.(map[string]interface{})["imgurl"].(string)
			url = url +"#"+img.(map[string]interface{})["caption"].(string)
			fmt.Printf("find -->%d:%s\n",k,url)
			imageC.Visit(url)
		}
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})
	imageC.OnResponse(func(r *colly.Response) {
		fileName :=""
		caption :=strings.Split(r.Request.URL.String(),"#") // 获得刚刚#后面的信息
		if len(caption)>=2{ //这里要判断下没有信息的情况，要不然切片会越界
			fileName =caption[1] +".jpg"
		}else{
			fileName = "未知"
		}
		res, err := url.QueryUnescape(fileName) // 对url格式进行转换，要不然看不懂
		fileName = strings.Replace(res,",","_",-1)// 把信息中的逗号全部换成下换线，逗号文件命名会报错。
		fmt.Printf("下载 -->%s \n",fileName)
		f, err := os.Create("./download/"+fileName)
		if err != nil {
			panic(err)
		}
		io.Copy(f, bytes.NewReader(r.Body))
	})
	//构造URL
	pageSize:= 200 //需要下载图片的数量,
	pageNum :=10
	for i:=0;i<pageNum;i++{
		url :=fmt.Sprintf("https://www.quanjing.com/Handler/SearchUrl.ashx?t=1952&callback=searchresult&q=卡通A&stype=1&pagesize=%d&pagenum=%d&imageType=2&imageColor=&brand=&imageSType=&fr=1&sortFlag=1&imageUType=&btype=&authid=&_=1577435470818",pageSize,i)
		_ = c.Visit(url)
	}

	c.Wait()
	imageC.Wait()
	fmt.Printf("done,cost:%s\n",time.Since(t))
}