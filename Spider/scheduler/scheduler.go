package scheduler

type Scheduler interface {
	//WorkerChan() chan interface{}
	Run()
	CreateWorker(out chan interface{},parseFunc func(interface{}) interface{})
	AddTask(task  interface{})
}