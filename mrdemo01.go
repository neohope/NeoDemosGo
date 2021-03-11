package main

import (
	"fmt"
	"strings"
)

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

func Filter(arr []int, fn func(n int) bool) []int {
	var newArray = []int{}
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

func main() {
	var l1 = []string{"Hao", "Chen", "MegaEase"}
	x := MapStrToStr(l1, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("upper string: %v\n", x)
	//["HAO", "CHEN", "MEGAEASE"]


	y := MapStrToInt(l1, func(s string) int {
		return len(s)
	})
	fmt.Printf("string len: %v\n", y)
	//[3, 4, 8]


	var l2 = []string{"Hao", "Chen", "MegaEase"}
	z := Reduce(l2, func(s string) int {
		return len(s)
	})
	fmt.Printf("total len: %v\n", z)
	// 15


	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(intset, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("odd num: %v\n", out)

	out = Filter(intset, func(n int) bool {
		return n > 5
	})
	fmt.Printf("num>5 :%v\n", out)
}
