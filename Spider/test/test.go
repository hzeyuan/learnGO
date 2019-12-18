package main

import (
	"fmt"
)

func orDone(done chan chan  interface{}, c chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case w:=<-done:
				fmt.Println("w",w)
			case t := <-c:
				fmt.Println("从channel中取出数据",t)
			case c <-valStream:
				fmt.Println("放入数据到channel中",valStream)
			}
		}
	}()
	return valStream
}

func main() {


	//done :=make(chan chan interface{})
	//ch := make(chan interface{})
	//orDone(done,ch)
	//for i := 0; i < 10; i++ {
	//	ch <- i
	//}
	//for i := 0; i < 10; i++ {
	//	<-ch
	//}


}