package main

import (
	"fmt"
	"reflect"
)

// 补充类型校验，提升泛型支持
// 但反射性能较差

// map
func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

// 在原有内存map
func TransformInPlace(slice, function interface{}) {
	transform(slice, function, true)
}

// 先校验slice类型、function签名是否正确
func transform(slice, function interface{}, inPlace bool) interface{} {

	//检查slice为slice类型
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	//检查出参入参符合函数签名
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	//输出是否重用输入内存
	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}

	//反射调用函数
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}

	return sliceOutType.Interface()
}

// 先校验fn为函数
// 然后确保fn的入参和出差类型，与types一致
func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {

	//Check it is a funciton
	if fn.Kind() != reflect.Func {
		return false
	}
	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

// Reduce 健壮版本
func SReduce(slice, pairFunc, zero interface{}) interface{} {

	//检查slice为slice类型
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}

	len := sliceInType.Len()
	if len == 0 {
		return zero
	} else if len == 1 {
		return sliceInType.Index(0)
	}

	//检查出参入参符合函数签名
	elemType := sliceInType.Type().Elem()
	fn := reflect.ValueOf(pairFunc)
	if !verifyFuncSignature(fn, elemType, elemType, elemType) {
		t := elemType.String()
		panic("reduce: function must be of type func(" + t + ", " + t + ") " + t)
	}

	//反射调用函数
	var ins [2]reflect.Value
	ins[0] = sliceInType.Index(0)
	ins[1] = sliceInType.Index(1)
	out := fn.Call(ins[:])[0]

	for i := 2; i < len; i++ {
		ins[0] = out
		ins[1] = sliceInType.Index(i)
		out = fn.Call(ins[:])[0]
	}
	return out.Interface()
}

// Reduce 健壮版本
func SReduceE(slice, ans, pairFunc, zero interface{}) interface{} {

	//检查slice为slice类型
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}

	len := sliceInType.Len()
	if len == 0 {
		return zero
	} else if len == 1 {
		return sliceInType.Index(0)
	}

	//省略了函数参数校验相关代码
	fn := reflect.ValueOf(pairFunc)

	//反射调用函数
	var ins [2]reflect.Value
	ins[0] = reflect.ValueOf(ans)
	ins[1] = sliceInType.Index(0)
	out := fn.Call(ins[:])[0]

	//反射调用函数
	for i := 2; i < len; i++ {
		ins[0] = out
		ins[1] = sliceInType.Index(i)
		out = fn.Call(ins[:])[0]
	}

	return out.Interface()
}

// Filter 健壮版本
func SFilter(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}

// 需要修改slice长度，所以要slicePtr
func SFilterInPlace(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: wrong type, not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	in.Elem().SetLen(n)
}

// 即使这样，由于slice是传值，所以也无法直接修改slice
func SFilterInPlace1(slice, function interface{}) {
	op := make([]int, 0)
	op = slice.([]int)
	_, n := filter(op, function, true)
	in2:=reflect.ValueOf(&op)
	in2.Elem().SetLen(n)
	fmt.Println(in2)
}

var boolType = reflect.ValueOf(true).Type()

func filter(slice, function interface{}, inPlace bool) (interface{}, int) {
	//检查slice为slice类型
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	//检查出参入参符合函数签名
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, boolType) {
		panic("filter: function must be of type func(" + elemType.String() + ") bool")
	}

	//反射调用函数
	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}

	//输出是否重用输入内存
	out := sliceInType
	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i]))
	}

	return out.Interface(), len(which)
}

type TEmployee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

func main(){
	slist := []string{"1", "2", "3", "4", "5", "6"}
	sresult := Transform(slist, func(a string) string{
		return a +a +a
	})
	fmt.Println(sresult)
	//{"111","222","333","444","555","666"}


	ilist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	TransformInPlace(ilist, func (a int) int {
		return a*3
	})
	fmt.Println(ilist)
	//{3, 6, 9, 12, 15, 18, 21, 24, 27}


	var elist = []TEmployee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
	}
	TransformInPlace(elist, func(e TEmployee) TEmployee {
		e.Salary += 1000
		e.Age += 1
		return e
	})
	fmt.Println(elist)
	//[{Hao 45 0 9000} {Bob 35 10 6000} {Alice 24 5 10000} {Jack 27 0 5000} {Tom 49 9 8500}]

	mresult := SReduce(slist, func(a string, b string) string{
		return a + b
	}, nil)
	fmt.Println(mresult)

	mresultE := SReduceE(elist, 0 ,func(a int, e TEmployee) int{
		return a + e.Salary
	}, nil)
	fmt.Println(mresultE)

	fresult := SFilter(ilist, func(a int) bool{
		return a>10
	})
	fmt.Println(fresult)

	SFilterInPlace(&ilist, func(a int) bool{
		return a>15
	})
	fmt.Println(ilist)

	// 这个是错误的
	SFilterInPlace1(ilist, func(a int) bool{
		return a>20
	})
	fmt.Println(ilist)
}