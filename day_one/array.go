package day_one

import "fmt"

/**
1、数组是值类型，调用函数是值传递；可以传指针数组，但一般不会这样使用，会使用切片
2、[]int 这是切片 [1]int 这是数组
3、函数传参定义[5]int 外部定义[...]int的值也必须是5个
4、[5]int 和 [10]int 是不同的类型
 */
type Array struct {

}

func (a *Array) MakeArr() {
	// 指定长度5 初始值都为0
	var arr1 [5]int
	// 指定长度3 自定义初始值
	arr2 := [3]int{1,3,5}
	// 编译器自行判断长度
	arr3 := [...]int{2, 4, 6, 8, 10}
	// 二维数组
	grid := [4][5]bool{}
	fmt.Println(arr1, arr2, arr3, grid)

}

func (a *Array) Each(arr []int) {

	arr[0] = 100
	// i := range 默认返回索引
	// i, v := range 索引 和 值都会返回
	for _, v := range arr {
		fmt.Printf("each %v \n", v)
	}

}
