package main

import "fmt"

//解耦了数据和程序
//使用了修饰器模式
//做出了 Pipeline 的模式
type VisitorFunc02 func(*Info02, error) error

type Visitor02 interface {
	Visit(VisitorFunc02) error
}

type Info02 struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info02) Visit(fn VisitorFunc02) error {
	return fn(info, nil)
}

type NameVisitor02 struct {
	visitor Visitor02
}

func (v NameVisitor02) Visit(fn VisitorFunc02) error {
	return v.visitor.Visit(func(info *Info02, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

type OtherThingsVisitor02 struct {
	visitor Visitor02
}

func (v OtherThingsVisitor02) Visit(fn VisitorFunc02) error {
	return v.visitor.Visit(func(info *Info02, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

type LogVisitor02 struct {
	visitor Visitor02
}

func (v LogVisitor02) Visit(fn VisitorFunc02) error {
	return v.visitor.Visit(func(info *Info02, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

func main() {
	info := Info02{}
	var v Visitor02 = &info
	v = LogVisitor02{v}
	v = NameVisitor02{v}
	v = OtherThingsVisitor02{v}

	loadFile := func(info *Info02, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)
}
