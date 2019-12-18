package main

import (
	"fmt"
	"time"
)

func Fib(n int) int64 {
	if n < 1 {
		return -1
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}

}

func main() {
	t1 := time.Now()
	sum := Fib(38)
	t3 := time.Since(t1)
	fmt.Println("sum=", sum)
	fmt.Println("use time:", t3)

}
