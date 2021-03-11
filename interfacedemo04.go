package main

import (
	"fmt"
)

type withName struct {
	Name string
}

type countryA struct {
	withName
}

type cityA struct {
	withName
}

type printableA interface {
	PrintStr()
}

func (w withName) PrintStr() {
	fmt.Println(w.Name)
}

func main(){
	c1 := countryA {withName{ "China"}}
	c2 := cityA { withName{"Beijing"}}
	c1.PrintStr()
	c2.PrintStr()
}