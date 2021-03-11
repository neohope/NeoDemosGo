package main

type IntSet2 struct {
	data map[int]bool
}
func NewIntSet2() IntSet2 {
	return IntSet2{make(map[int]bool)}
}
func (set *IntSet2) Add(x int) {
	set.data[x] = true
}
func (set *IntSet2) Delete(x int) {
	delete(set.data, x)
}
func (set *IntSet2) Contains(x int) bool {
	return set.data[x]
}
