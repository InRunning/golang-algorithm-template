package stack

type Stack struct {
	Top    *node
	Length int
}

type node struct {
	value interface{}
	prev  *node
}

func NewStack() *Stack {
	newStack := &Stack{nil, 0}
	return newStack
}

func (stack *Stack) Push(value interface{}) {
	n := &node{
		value: value,
		prev:  stack.Top,
	}
	stack.Top = n
}

func (stack *Stack) Pop() interface{} {
	n := *stack.Top
	stack.Top = stack.Top.prev
	return n.value
}
