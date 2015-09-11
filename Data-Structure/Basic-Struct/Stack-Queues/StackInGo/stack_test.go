package StackInGo

import (
	//"errors"
	//"fmt"
	"log"
	"testing"
)

func TestStack(t *testing.T) {
	stack := New()
	if stack.len != 0 {
		t.Error("Wrong Initialized!")
	} else {
		log.Println("Initialized Success")
	}

	stack.Push(5)

	if stack.Pop() != nil {
		t.Error(stack.Pop())
	} else {
		log.Println("Successful pop")
	}

	stack.Push(6)
	stack.Push("John")

	stack.Push("A")
	stack.Push(7)

	stack.Appear()

}
