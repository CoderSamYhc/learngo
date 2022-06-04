package day_one

import "fmt"

/**
1、空切片实际是nil；zero value for slice is nil；
 */

type SliceOps struct {

}

func (s *SliceOps) Add() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] // [2 3 4 5]
	s2 := s1[3:5] // [5 6]
	s3 := append(s2, 10) // [5 6 10]
	s4 := append(s3, 11) // [5 6 10 11]
	s5 := append(s4, 12) // [5 6 10 11 12]
	// arr [0 1 2 3 4 5 6 10]
	// s4 和 s5 view了一个新的数组，并且长度比arr长
	fmt.Println(s3, s4, s5, arr)
}

func (s *SliceOps) Create() {
	var slice []int // zero value for slice is nil；
	for i := 0; i < 100; i++ {
		fmt.Printf("len = %v, cap = %v \n", len(slice), cap(slice))
		slice = append(slice, i)
	}
	fmt.Println(slice)
}
