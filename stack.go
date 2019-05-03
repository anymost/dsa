package main

import "strings"

type Stack struct {
	val []interface{}
}

func NewStack() *Stack {
	return &Stack{
		val: make([]interface{}, 0),
	}
}

func (stack *Stack) Top() interface{} {
	return stack.val[len(stack.val) - 1]
}

func (stack *Stack) Empty() bool {
	return len(stack.val) == 0
}

func (stack *Stack) Pop() interface{} {
	val := stack.Top()
	stack.val = stack.val[0: len(stack.val) - 1]
	return val
}

func (stack *Stack) Push(val interface{}) {
	stack.val = append(stack.val, val)
}

func Convert(n int, base int) string {
	digit := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	stack := NewStack()
	for n > 0 {
		temp := n%base
		stack.Push(digit[temp])
		n = (n -temp ) / base
	}
	s := ""
	for !stack.Empty() {
		if v, ok := stack.Pop().(string); ok {
			s += v
		}
	}
	return s
}

func match(words []string) bool {
	stack := NewStack()
	for len(words) != 0 {
		tail := words[len(words) - 1]
		words = words[0: len(words) - 1]
		switch tail {
		case "(", "[", "{":
			stack.Push(tail)
		case ")", "]", "}":
			top := stack.Pop()
			if (top == "(" && tail == ")") ||
				( top == "[" && tail == "]") ||
				( top == "{" && tail == "}") {
				continue
			} else {
				return false
			}
		default:
			continue
		}
	}
	return true
}

func main() {
	words := "1+2+3+{[(1+1)+2]+3}"
	v := match(strings.Split(words, ","))
	println(v)
}

