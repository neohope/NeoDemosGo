package main

import "fmt"

// 装饰器例子，类似于AOP
func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Begin")
		f(s)
		fmt.Println("End")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func main() {
	//decorator(Hello)("Hello, World!")
	hello := decorator(Hello)
	hello("Hello")
}
