package day_one

import "fmt"

/**
1、左开右闭 指定的最大索引不包含在结果集内
2、切片是数组的一个视图，即切片是引用传递，函数内对其某个索引的值进行修改，会影响外部的值
3、切片包含 ptr、len、cap；len包含ptr，cap包含len和ptr
4、切片可以向后扩展，不可以向前扩展；s[i]不可以超越len(s)，向后扩展不可以超越底层数组的cap(s)
5、添加元素超过底层数组的cap(s)时，底层会开辟一个新的cap更大的新数组
 */
type Slice struct {

}

func (s *Slice) MakeSlice() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// 左开右闭 索引为6的不包含在内
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
}

func (s *Slice) UpdateSlice(slice []int) {
	slice[0] = 100

	fmt.Printf("slice %v \n", slice)
}

func (s *Slice) ReSlice() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] // [2 3 4 5]
	s2 := s1[3:5] // [5 6]
	fmt.Printf("s1 = %v, s1 len = %v, cap = %v \n",s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, s2 len = %v, cap = %v \n", s2, len(s2), cap(s2))
	fmt.Printf("s1 index 6 = %v \n", s1[:cap(s1)])

}
