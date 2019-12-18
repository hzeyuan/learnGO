package main

import "io"

// 原始数据
type genericProducer interface {
	Next()(interface{},error)
}
type genericConsumer interface {
	Send(interface{}) error
}

//原始数据加工的方法
type genericMapper func(interface{}) (interface{},error)

func ConcurrentMapPoorErrorHandling(p genericProducer, c genericConsumer, mapper genericMapper) error {
	count := 0
	// empty struct{} is a type in Go. You can also redefine the type as:
	// type DoneSignal struct{}
	done := make(chan struct{})
	for {
		next, err := p.Next()
		if err != nil {
			if err == io.EOF {
				break // There is no more elements in the producer.
			}
			return err // There is an error in the producer. Shut down the mapping.
		}
		count++
		go func(next interface{}) {
			ele, err := mapper(next)
			if err != nil {
				panic(err)
			}
			err = c.Send(ele)
			if err != nil {
				panic(err)
			}
			done <- struct{}{}
		}(next)
	}
	for i := 0; i < count; i++ {
		<-done
	}
	return nil
}

