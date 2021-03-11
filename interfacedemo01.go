package main

import (
	"fmt"
)

type person struct {
	Name string
	Sexual string
	Age int
	}

func PrintPerson(p *person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func (p *person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func main() {
	var p = person{
		Name: "Hao Chen",
		Sexual: "Male",
		Age: 44,
	}

	PrintPerson(&p)
	p.Print()
}