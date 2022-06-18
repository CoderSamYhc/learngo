package day_3

import (
	"time"
)

func TestFunc2() {
	done := make(chan bool)
	c := make(chan string)

	go func() {
		defer close(done)
		s := <-c
		println(s)
	}()

	c<-"hi"
	<-done
}


func producer(x int, ch chan<- int) {
	for i := 0; i < x ; i++ {
		ch <- i * x
	}
}

func consumer(ch <-chan int) {
	for v := range ch {
		println(v)
	}
}

func TestFunc3() {

	ch := make(chan int, 64)
	x := 3

	go producer(x, ch)
	go producer(x, ch)
	go consumer(ch)

	time.Sleep(time.Second * 3)
}

