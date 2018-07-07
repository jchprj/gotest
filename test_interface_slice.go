package main

import (
	"fmt"
	"reflect"
)

type iii []interface{}

func testInterface(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			fmt.Println("s:", s.Index(i))
		}
	}
}

func remove(t interface{}, value interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			fmt.Println("s:", s.Len(), s.Index(i), reflect.ValueOf(value), s.Index(i) == reflect.ValueOf(value))
			if s.Index(i) == reflect.ValueOf(value) {
				if i == 0 {
					s = s.Slice(i+1, s.Len())
				} else {
					s = reflect.AppendSlice(s.Slice(0, i-1), s.Slice(i+1, s.Len()))
				}
				i--
			}
		}
	}
}

//TestInterfaceSlice test
func TestInterfaceSlice() {
	ddd := []int{1, 2}
	d1 := 1
	remove(ddd, d1)
	testInterface(ddd)
}
