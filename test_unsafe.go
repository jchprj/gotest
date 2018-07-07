package main

import (
	"fmt"
	"unsafe"
)

//TestUnsafe test
func TestUnsafe() {
	ddd := []int{1, 2}
	// dd := &ddd
	// ba := unsafe.Alignof(dd)
	p := unsafe.Pointer(&ddd)
	// ii := int.s
	pp := *(*int)(p)
	// c := unsafe.Pointer(p)
	var ppp interface{}
	ppp = unsafe.Pointer(uintptr(pp + 8))
	fmt.Println(ppp)
}
