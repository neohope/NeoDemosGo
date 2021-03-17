package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 泛型支持
// 目前的 Go 语言的泛型只能用 interface{} + reflect来完成
func MapGeneric(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}
	return result
}

func main(){
	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}
	squaredNums:= MapGeneric(nums,square)
	fmt.Println(squaredNums)
	//[1 4 9 16]

	
	upcase := func(s string) string {
		return strings.ToUpper(s)
	}
	strs := []string{"Hao", "Chen", "MegaEase"}
	upStrs := MapGeneric(strs, upcase);
	fmt.Println(upStrs)
	//[HAO CHEN MEGAEASE]

	// 这里会引起panic
	// 因为5没有Len()方法
	x := MapGeneric(5, 5)
	fmt.Println(x)
}
