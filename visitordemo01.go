package main

import (
"encoding/json"
"encoding/xml"
"fmt"
)

// 访问者模式
type Visitor01 func(shape VShape)

type VShape interface {
	accept(Visitor01)
}

type Circle struct {
	Radius int
}

func (c Circle) accept(v Visitor01) {
	v(c)
}

type Rectangle struct {
	Width, Heigh int
}

func (r Rectangle) accept(v Visitor01) {
	v(r)
}

func JsonVisitor01(shape VShape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor01(shape VShape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	c := Circle{10}
	r :=  Rectangle{100, 200}
	shapes := []VShape{c, r}

	for _, s := range shapes {
		s.accept(JsonVisitor01)
		s.accept(XmlVisitor01)
	}
}
