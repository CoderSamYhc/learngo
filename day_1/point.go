package day_1

import "fmt"

type Point struct {

}

func (p *Point) Point() {

	var a int = 10
	// 指针类型获取到变量a的地址
	var pa *int = &a

	fmt.Println(a)
	// 通过地址修改了变量a的值
	*pa = 3
	fmt.Println(a)
}

/**
值传递：copy 值，函数内部修改不会影响外部
 */
func (p *Point) PassByVal(v int) {
	v += 1
	fmt.Printf("pass by val %v \n", v)
}

/**
引用传递：传递值的地址，函数内修改会影响外部
 */
func (p *Point)  PassByRef(i *int) {
	*i += 1
	fmt.Printf("pass by val %v \n", *i)
}

/**
交互两个变量的值
 */
func (p *Point) Swap(a, b *int) {
	*a, *b = *b, *a
}