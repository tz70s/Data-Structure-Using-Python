package StackInGo

import (
	"errors"
	"fmt"
)

type Stack struct {
	st  []interface{} // use interface type to accept every type to append in stack
	len int
}

func New() *Stack {
	stack := &Stack{}
	stack.st = make([]interface{}, 0)
	stack.len = 0
	return stack
}

func (stack *Stack) Length() int {
	return stack.len
}

func (stack *Stack) Pop() (err error) {
	if stack.IsEmpty() != true {
		stack.st = stack.st[:stack.len-1]
		stack.len -= 1
	} else {
		err = errors.New("Attempting to Pop an empty stack")
		fmt.Println(err)
	}
	return err

}

func (stack *Stack) IsEmpty() bool {
	return (stack.len == 0)
}

func (stack *Stack) Push(item interface{}) {
	stack.st = append(stack.st, item)
	stack.len += 1
}

func (stack *Stack) Appear() {
	for _, v := range stack.st {
		fmt.Print(v, " ")
	}
}
