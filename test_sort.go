package main

import (
	"fmt"
	"sort"
)

type asd []int

func (ss asd) Len() int {
	return len(ss)
}
func (ss asd) Less(a int, b int) bool {
	return ss[a] > ss[b]
}
func (ss asd) Swap(a int, b int) {
	ss[a], ss[b] = ss[b], ss[a]
}

func sssss(as []int) {
	var ss asd
	ss = as
	sort.Sort(ss)
}

//TestSort test
func TestSort() {
	ddd := []int{1, 2}
	fmt.Println("1111111", ddd)
	sssss(ddd)
	fmt.Println("ddddddddddddd", ddd)
}
