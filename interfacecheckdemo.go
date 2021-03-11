package main

import "fmt"

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s* Square) Sides() int {
	return 4
}

func main() {
	s := Square{len: 5}
	fmt.Printf("%d\n",s.Sides())

	//可以用这个方法做校验
	//从而判断接口是否有全部实现
	//var _ Shape = (*Square)(nil)
}