package QueuInGo

import (
	"fmt"
	"log"
	"testing"
)

func TestQue(t *testing.T) {
	queue := NewQue()

	log.Println("Create the queue : ", queue.IsEmpty())
	fmt.Println()

	/*
		if queue.Pop() != nil {
			t.Error(queue.Pop())
		} else {
			log.Println("Queue Pop Success")
		}
	*/

	queue.Push(2)
	queue.Push(3)
	queue.Push(5)
	queue.Push("John")

	log.Println("Queue Push Success")
	queue.Appear()
	fmt.Println()

	if queue.Pop() != nil {
		t.Error(queue.Pop())
	} else {
		log.Println("Queue Pop Success")
	}

	queue.Appear()
	fmt.Println()

}
