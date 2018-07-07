package test

import (
	"fmt"
	"strconv"
)

type Foo interface {
	Bar()
}
type impl int

func (i impl) Bar() {
	fmt.Println("dd")
}

type Fbb struct {
	AB int
}

func (i Fbb) Bard() {
	fmt.Println("dd" + strconv.Itoa(i.AB))
}
