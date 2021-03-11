package main

import (
	"bytes"
	"fmt"
)

// foo与bar共享内存，bar修改了，foo也就修改了
func test01() {
	var foo = make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99

	fmt.Println("foo[2] =>",foo[2])
}

// append()这个函数在 cap 不够用的时候，就会重新分配内存以扩大容量，如果够用，就不会重新分配内存了！
// append后，由于a重新申请分配了内存，a与b就不再共享内存了，修改a对b就没有影响了
func test02() {
	var a = make([]int, 32)
	b := a[1:16]

	a[2] = 3
	fmt.Println("b[1] =>",b[1])

	a = append(a, 1)
	a[2] = 7
	fmt.Println("b[1] =>",b[1])
	fmt.Println("a =>",a)
	fmt.Println("b =>",b)
}

// 这里有些坑，dir1做append时，由于空间够用，把dir2和path的数据反而都修改了
func test03() {
	var path = []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1,"suffix"...)
	fmt.Println("path =>",string(path)) //prints: path => AAAAsuffixBBBB
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB

	// dir1扩容了
	dir1 = append(dir1,"suffix"...)
	fmt.Println("path =>",string(path)) //prints: path => AAAAsuffixBBBB
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffixsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB
}

// 解决方法
// 使用Full Slice Expression，最后一个参数叫“Limited Capacity”，
func test04() {
	var path = []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')

	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1,"suffix"...)
	fmt.Println("path =>",string(path)) //prints: path => AAAAsuffixBBBB
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
}

func main() {
	fmt.Println(">>>test01")
	test01()
	fmt.Println(">>>test02")
	test02()
	fmt.Println(">>>test03")
	test03()
	fmt.Println(">>>test04")
	test04()
}
