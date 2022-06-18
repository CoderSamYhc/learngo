package day_3

import "time"

var c int
func test() int {
	c++
	return c
}

// 延时执行
func Do() {
	a := 100

	go func(x, y int) {
		time.Sleep(time.Second)
		println(x, y)
	}(a, test())

	a += 100
	println("do:", a, test())
	time.Sleep(time.Second * 3)

}