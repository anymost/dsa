package main

type Fib struct {
	index int
	value int
}

func NewFib(n int) *Fib {
	g, f, originN := 1, 0, n
	for 0 < n-1 {
		g = g + f
		f = g - f
		n--
	}
	if n > 0 {
		return &Fib{
			value: g,
			index: originN,
		}
	} else {
		return &Fib{
			value: 0,
			index: 0,
		}
	}
}

func (fib *Fib) Get() int {
	return fib.value
}

func (fib *Fib) Prev() {
	newFib := NewFib(fib.index - 1)
	fib.value = newFib.value
	fib.index = newFib.index
}
