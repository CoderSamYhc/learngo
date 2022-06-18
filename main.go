package main

import (
	"github.com/CoderSamYhc/learngo/day_3/ops"
	"strings"
	"time"
)

func main() {

	/* 指针
	var p = day_1.Point{}
	p.Point()
	var val = 2
	p.PassByVal(val)
	fmt.Printf("global val %v \n", val)
	p.PassByRef(&val)
	fmt.Printf("global ref %v \n", val)

	var a, b = 2, 10

	fmt.Println(a, b)
	p.Swap(&a, &b)
	fmt.Println(a, b)
	*/

	/* 数组
		var arr = day_1.Array{}
		arr.MakeArr()

		eachArr := [...]int{1, 3, 4, 5, 9}
		arr.Each(eachArr[:])
		fmt.Println(eachArr)
	*/

	/* 切片
	var s = day_1.Slice{}
	s.MakeSlice()

	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s.UpdateSlice(arr[2:])
	fmt.Printf("array %v \n", arr)

	s.ReSlice()
	*/

	/* 切片例题
		var s = day_1.SliceOps{}
		s.Add()
		s.Create()
		s.Copy()
		s.Delete()
	*/
	/* map
		var m  = day_1.Map{}
		m.Create()
		fmt.Println(m.FindLongNoRepeatSubStr("bcbac"))
		fmt.Println(m.FindLongNoRepeatSubStr("我我妮妮"))
	*/

	/*
	var duck day_2.Bird
	duck = &day_2.Duck{}
	duck.SetName("大黄鸭")
	duck.Call()

	var tiger day_2.Catamount
	tiger = &day_2.Tiger{}
	tiger.SetName("花老虎")
	tiger.Eating(duck.GetName())
	*/

	//var wg sync.WaitGroup
	//
	//t := day_3.Total{}
	//wg.Add(2)
	//go t.Worker(&wg)
	//go t.Worker(&wg)
	//
	//wg.Wait()
	//fmt.Println(t.Value)

	//var m1 map[string]string
	//fmt.Println(m1)

	//day_3.Do()

	//day_3.TestFunc3()

	p := ops.NewPublisher(100*time.Millisecond, 10)

	defer p.Close()

	all := p.Subscriber()
	golang := p.SubscriberTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello, world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range all {
			println("all", msg.(string))
		}
	}()

	go func() {
		for msg := range golang {
			println("golang", msg.(string))
		}
	}()

	time.Sleep(time.Second*3)
}
