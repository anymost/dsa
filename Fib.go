package main

import "fmt"

type Fib struct {
	index int
	value int
}

func NewFib(n int) *Fib {
	g, f := 1, 0
	for 0 < n-1 {
		g = g + f
		f = g - f
		n--
	}
	return &Fib{
		value: g,
		index: n,
	}
}

func (fib *Fib) get() int {
	return fib.value
}

func (fib *Fib) prev() {
	newFib := NewFib(fib.index - 1)
	fib.value = newFib.value
	fib.index = newFib.index
}

func main()  {
	fib := NewFib(7)
	fib.prev()
	fmt.Println(fib.get())
}