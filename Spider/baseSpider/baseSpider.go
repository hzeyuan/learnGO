package baseSpider

import (
	"container/list"
	"go-DIYSpider/fetch"
	"go-DIYSpider/scheduler"
	"os"
)


func judgeType(v interface{}) (string){
	switch i := v.(type) {
	case int:
		return "int"
	case int64:
		return "int64"
	case int32:
		return "int32"
	case string:
		return "string"
	case float64:
		return "float64"
	case []string:
		return "[]string"
	case []interface{}:
		return "[]interface"
	default:
		_ = i
		return "unknown"
	}
}

type spiderRet struct {
	html	string
}

type BaseSpider struct {
	url            *list.List //双向链表
	Fetch          *fetch.Fetch
	spiderRet		map[string]*spiderRet // 结构url-html
	SuccessUrlList []string
	FailUrlList    []string
	Scheduler      scheduler.Scheduler
	WorkerCount    int //开启的worker数量
	HowToSave		func(string)
}

func NewBaseSpider() *BaseSpider {
	baseSpider := &BaseSpider{
		Fetch:fetch.NewFetch(),
		Scheduler:scheduler.NewQueueScheduler(),
		SuccessUrlList:nil,
		spiderRet: make(map[string]*spiderRet),
		WorkerCount:10,
		FailUrlList:nil,
	}
	return baseSpider
}

// 设置URL,支持字符串,切片
func (b *BaseSpider) Url(url ...interface{}) *BaseSpider {
		l :=list.New()
		urlList :=list.New()
		l.PushFront(url)
		for l.Len()!=0 {
			urlT := l.Remove(l.Front())
			if judgeType(urlT) == "string"{
				urlList.PushFront(urlT.(string))
			}else if judgeType(urlT)=="[]string"{
				temp := list.New()
				for _,u :=range urlT.([]string){
					//l.PushFront(u)
					temp.PushFront(u)
				}
				l.PushFrontList(temp)
			}else if judgeType(urlT) == "[]interface"{
				for _,u :=range urlT.([]interface{}){
					l.PushFront(u)
				}


			}
		}
		if b.url ==nil{
			b.url = urlList
		}else{
			b.url.PushBackList(urlList)
		}
	return b
}

// 获取所有的URL
func(b *BaseSpider) GetUrl() (urlList []string){
		for e:=b.url.Front();e!=nil;e=e.Next(){
			urlList = append(urlList,e.Value.(string))
		}
		return urlList
	}

// 判断Fetch对象成功或者失败
func (b *BaseSpider) addUrlToList() *BaseSpider {

		url :=b.Fetch.SuperAgent.Url
		if b.Fetch.StatusCode==200{
		b.SuccessUrlList = append(b.SuccessUrlList,url)
		}else{
		//errMsg := func() (errMsg string) {
		//	for _,e :=range b.Fetch.ErrMsg{
		//		errMsg =strings.Join([]string{error.Error(e)}," ")
		//	}
		//	return errMsg
		//}
		b.FailUrlList = append(b.FailUrlList,url)
	}
		return b
}

func (b *BaseSpider) Run(url ...interface{}) *BaseSpider {
	//var wg = new(sync.WaitGroup)
	b.Url(url...) //添加跟过的url
	urlList :=b.GetUrl()
	workRet :=make(chan interface{})
	go func() { // 添加任务
		for _,w:=range urlList{
			//wg.Add(1)
			b.Scheduler.AddTask(w)
		}
	}()

	//初始化scheduler
	b.Scheduler.Run()

	b.HowToSave = func(name string) {
			f,_ :=os.Create(name)
			defer f.Close()
			txt := <-workRet
			go f.WriteString(txt.(string))

	}


	for i:=0;i<len(urlList);i++{ //待处理的任务
		b.Scheduler.CreateWorker(workRet, func(i interface{}) interface{} {
			return ""
		})
		//_ := <- workRet
		//fmt.Println(<- workRet)
		b.HowToSave("test.txt")
	}
	//wg.Wait()
	return b
}



