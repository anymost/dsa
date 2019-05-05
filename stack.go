package main

// 卡特兰数： 栈混洗的总可能数为 (2n)! / [(n+1)!*n!]
// 非法的栈混洗：如 1 < 2 < 3，如果出现 3, 1, 2

// 中缀表达式：分操作符和操作栈两个部分
// 列一张二维表来区分不同运算符之间的优先级

// 逆波兰表达式 RPN：不适用括号来表示优先级的原酸关系
// 逆波兰表达式的转换：如果是数字，直接续接；如果是操作符，如果栈顶操作符优先级大于当前操作符，则将栈顶操作符续接

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

func match(array []string) bool {
	stack := NewStack()
	for i := 0; i < len(array); i++ {
		val := array[i]
		switch val {
		case "(", "[", "{":
			stack.Push(val)
		case ")", "]", "}":
			if stack.Empty() {
				return false
			} else if (stack.Top() == "(" && val == ")") ||
				(stack.Top() == "[" && val == "]") ||
				(stack.Top() == "{" && val == "}") {
				stack.Pop()
			} else {
				return false
			}
		default:
			continue
		}
	}
	return stack.Empty()
}





