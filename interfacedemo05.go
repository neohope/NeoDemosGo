package main

import (
	"fmt"
)

//这种编程模式在 Go 的标准库有很多的示例，最著名的就是 io.Read 和 ioutil.ReadAll 的玩法，
//其中 io.Read 是一个接口，你需要实现它的一个 Read(p []byte) (n int, err error) 接口方法，
//只要满足这个规则，就可以被 ioutil.ReadAll这个方法所使用。
//这就是面向对象编程方法的黄金法则——“Program to an interface not an implementation”。

type countryB struct {
	Name string
}

type cityB struct {
	Name string
}

type stringableB interface {
	ToString() string
}
func (c countryB) ToString() string {
	return "Country = " + c.Name
}
func (c cityB) ToString() string{
	return "City = " + c.Name
}

func PrintStr(p stringableB) {
	fmt.Println(p.ToString())
}


func main(){
	d1 := countryB {"USA"}
	d2 := cityB{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}