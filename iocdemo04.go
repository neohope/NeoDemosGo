package main

import "errors"

// 比上一个demo增加控制反转功能
// 控制反转，不是由控制逻辑 Undo 来依赖业务逻辑 IntSet，而是由业务逻辑 IntSet 依赖 Undo 。
// 这里依赖的是其实是一个协议，这个协议是一个没有参数的函数数组。
// 从而从业务逻辑中，抽离了Undo的控制逻辑
// 这样一来，我们 Undo 的代码就可以复用了。
type Undo []func()

func (undo *Undo) AddAction(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) UndoAction() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("no functions to undo")
	}

	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

type IntSet4 struct {
	data map[int]bool
	undo Undo
}

func NewIntSet4() IntSet4 {
	return IntSet4{data: make(map[int]bool)}
}

func (set *IntSet4) Undo() error {
	return set.undo.UndoAction()
}

func (set *IntSet4) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet4) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.AddAction(func() { set.Delete(x) })
	} else {
		set.undo.AddAction(nil)
	}
}

func (set *IntSet4) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.AddAction(func() { set.Add(x) })
	} else {
		set.undo.AddAction(nil)
	}
}
