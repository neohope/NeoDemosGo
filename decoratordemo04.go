package main

import (
	"fmt"
	"reflect"
)

// 第一个参数是装饰器函数，第二个参数是业务函数
// reflect.MakeFunc相当于做了一个代理函数
func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value
	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("begin")
			out = targetFunc.Call(in)
			fmt.Println("end")
			return
		})

	decoratedFunc.Set(v)
	return
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}

func main(){
	type MyFoo func(int, int, int) int
	var myfoo MyFoo
	Decorator(&myfoo, foo)
	myfoo(1, 2, 3)

	mybar := bar
	Decorator(&mybar, bar)
	mybar("hello,", "world!")
}
