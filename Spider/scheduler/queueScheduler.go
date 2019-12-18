package scheduler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"log"
)

type QueueScheduler struct {
	requestChan     chan interface{}             //待执行任务列表
	workerChan     chan chan interface{}        // 执行任务的工人
	WorkerCount int //工人数量
}

func NewQueueScheduler() *QueueScheduler {
	q := &QueueScheduler{
		make(chan interface{}),
		make(chan chan interface{}),
		10,
	}
	return q
}
// 添加任务
func (q *QueueScheduler) AddTask(task  interface{}) {
	q.requestChan <- task
}

func (q *QueueScheduler) WorkerChann() chan interface{} {
	return make(chan interface{})
}


// 工人需要做的事情
func (q *QueueScheduler)WorkerNeedToDo(r interface{}) string {
	log.Printf("fetching url:%s", r)
	res,_,_:= gorequest.New().Get(r.(string)).End()
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	ele :=doc.Find("h2.post-title ").First()
	//fmt.Println(ele.Text())
	return ele.Text()
}


func (q *QueueScheduler)CreateWorker(out chan interface{},parseFunc func(interface{}) interface{})  {
	in :=make(chan interface{})
	go func() {

		for{
			q.workerChan <- in //创建工人
			//fmt.Println("工人开始工作")
			//r := <-in //减少一个工人并且分配工作
			result := q.WorkerNeedToDo(<-in)
			//if err != nil {
			//	continue
			//}
			out <-result // 工作结果
		}

	}()
}


// 启动调度器
func (q *QueueScheduler) Run() {
	// 初始化 requestChann
	q.requestChan = make(chan interface{})
	// 初始化 workerChan
	q.workerChan = make(chan chan interface{})
	go func() {
		var requestQ  [] interface{}
		var workerQ  []chan  interface{}
		for {
			var activeRequest interface{}
			var activeWorker chan interface{}
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			//随机执行某个case
			select {
			case r := <-q.requestChan: //当请求过来时
				requestQ = append(requestQ, r)
				//fmt.Println("任务来了")
				//fmt.Println("任务数量",len(requestQ ))
			case w := <-q.workerChan: //当工人过来时
				workerQ = append(workerQ, w)
				//fmt.Printf("工人来啦\n",)
				//fmt.Println("工人数量",len(workerQ ))
			case activeWorker <- activeRequest: //把请求分配给工人，然后去掉任务
				//fmt.Println("工人开始工作")
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
