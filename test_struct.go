package main

import (
	"fmt"
	"gotest/test"
)

//TestStruct test
func TestStruct() {
	var a test.Fbb
	fmt.Println(a == test.Fbb{})
	a.AB = 1
	fmt.Println(a == test.Fbb{})
	b := &a
	b.Bard()
	a.Bard()
}
