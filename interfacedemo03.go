package main

import (
	"fmt"
)

type country struct {
	Name string
}

type city struct {
	Name string
}

type printable interface {
	PrintStr()
}

func (c country) PrintStr() {
	fmt.Println(c.Name)
}

func (c city) PrintStr() {
	fmt.Println(c.Name)
}

func main(){
	c1 := country{"China"}
	c2 := city {"Beijing"}
	c1.PrintStr()
	c2.PrintStr()
}