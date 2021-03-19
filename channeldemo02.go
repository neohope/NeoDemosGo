package main

import (
	"bytes"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"sync"
)

// 数据生成
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func echo02(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 判断是否为素数
func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// 排查非素数
func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func ()  {
		for n := range in {
			if is_prime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// 求和
func sum02(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		var xid = GetGID()
		for n := range in {
			sum += n
			fmt.Println(xid, n)
		}
		out <- sum
		close(out)
	}()
	return out
}

// 归并
func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 等待每一个协程返回结果
	wg.Add(len(cs))
	for t, c := range cs {
		go func(t int, c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(t, c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	nums := makeRange(1, 10000)
	in := echo02(nums)

	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = sum02(prime(in))
	}

	//也可以这样等待任务完成
	//var c int;
	//fmt.Scan(&c)

	// 等待计算任务完成
	m := merge(chans[:])

	for n := range sum02(m) {
		fmt.Println(n)
	}
}
