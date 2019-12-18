package fetch

import (
	"github.com/parnurzeal/gorequest"
)


//type fetch interface {
//	Get(url string) // 实现get方法
//}

type HTML string

//抓取结构
type Fetch struct {
	Url        string
	SuperAgent *gorequest.SuperAgent //使用goRequest
	html       HTML
	StatusCode int
	ErrMsg     []error
}


// 设置默认的抓取方式
func  NewFetch() *Fetch{
	fetch := &Fetch{
		Url:        "",
		SuperAgent: gorequest.New(),
		html:       "",
	}
	return fetch
}

// 返回html
func (f *Fetch) GetHTML() string{
	return string(f.html)
}

func (f *Fetch) Get(url string) (gorequest.Response, string, []error) {
	resp, html, errs:=f.SuperAgent.Get(url).End()
	f.html = HTML(html)
	f.ErrMsg = errs
	return resp, html, errs
}

func (f *Fetch) Post() *Fetch{
	return f
}