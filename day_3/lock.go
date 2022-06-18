package day_3

import (
	"fmt"
	"sync"
	"time"
)

func LockGroup() {

	wg := sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		wg.Add(1) // 不要放到goroutine里

		go func(x int) {
			time.Sleep(time.Second)
			println("goroutine:", x)
			defer wg.Done()
		}(i)

	}
	println("main..")
	wg.Wait()
	println("exit..")
}

func TestFunc() {
	wg := sync.WaitGroup{}

	type demoMap struct {
		Data map[int]struct{
			Id int
			Result int
		}
		Lock sync.RWMutex
	}

	m := demoMap{
		Data: make(map[int]struct{Id int
			Result int}),
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			m.Lock.Lock() // 防止map并发读写
			defer m.Lock.Unlock() // 后入先出
			m.Data[index] = struct {
				Id     int
				Result int
			}{Id: index, Result: (index + 1) * 100}
		}(i)
	}

	wg.Wait()
	fmt.Printf("m: %v", m.Data)
}
