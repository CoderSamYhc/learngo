package main

import "fmt"

//func main() {
//
//	s := []int{2, 3, 5, 7}
//	t := 8
//	for i := 0; i < len(s); i++ {
//		for j := i + 1; j < len(s); j++ {
//			if s[i] + s[j] == t {
//				fmt.Println(i, j)
//				break
//			}
//		}
//	}
//}


func main() {
	fmt.Println(test(0))
}

func test(x int) int {
	for i := 0; i < 1000000000; i++  {
		x += i
	}

	return x
}
