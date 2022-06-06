package day_1

import "fmt"

/**
1、key不存在时，获取value类型的初始值
2、除了map、slice、function，其他内建类型都可以作为key，struct 不包含map、slice、function也可以作为key
 */
type Map struct {

}

func (m *Map) Create() {
	vm := map[string]string{
		"name" : "sam",
		"nickname" : "codersam",
	}

	for i, v := range vm {
		fmt.Println(i, v)
	}
}


func (m *Map) FindLongNoRepeatSubStr(s string) int {
	lastOcrrred := make(map[rune]int)
	start, maxLenth := 0, 0
	for i, v := range []rune(s) {
		if lastI, ok := lastOcrrred[v]; ok && lastI >= start {
			start = lastI + 1
		}

		if i - start + 1 > maxLenth {
			maxLenth = 	i - start + 1
		}
		lastOcrrred[v] = i
	}
	return maxLenth
}
