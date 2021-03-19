package main

import "fmt"

type VisitorFunc03 func(*Info03, error) error

type Visitor03 interface {
	Visit(VisitorFunc03) error
}

type Info03 struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info03) Visit(fn VisitorFunc03) error {
	return fn(info, nil)
}

type DecoratedVisitor struct {
	visitor    Visitor03
	decorators []VisitorFunc03
}

func NewDecoratedVisitor(v Visitor03, fn ...VisitorFunc03) Visitor03 {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc03) error {
	return v.visitor.Visit(func(info *Info03, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func NameVisitorFun(info *Info03, err error) error {
	fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
	return err
}

func OtherThingsVisitorFun(info *Info03, err error) error {
	fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
	return err
}

func main(){
	info := Info03{}
	var v Visitor03 = &info
	v = NewDecoratedVisitor(v, NameVisitorFun, OtherThingsVisitorFun)

	loadFile := func(info *Info03, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)
}
