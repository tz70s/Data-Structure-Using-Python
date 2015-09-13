package DoublyLinklist

import (
	"errors"
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
	pre  *Node
}

type List struct {
	head   *Node
	length int
}

func NewList() *List {
	list := &List{}
	list.head = new(Node)
	list.length = 0
	return list
}

func (list *List) IsEmpty() bool {
	return list.length == 0
}

func (list *List) Insert(item interface{}) {

	tmp := &Node{}
	tmp.data = item

	if list.length == 0 {

		list.head = tmp

	} else if list.length == 1 {
		tmp.pre = list.head
		tmp.next = list.head
		list.head.next = tmp
		list.head.pre = tmp
	} else {

		current := new(Node)
		current = list.head
		for current.next != list.head {
			current = current.next
		}
		tmp.next = list.head
		list.head.pre = tmp
		tmp.pre = current
		current.next = tmp
	}

	list.length += 1
}

func (list *List) Delete(item interface{}) (err error) {
	current := new(Node)
	current = list.head

	for current.next != list.head {
		if current.data == item {
			current.pre.next = current.next
			current.next.pre = current.pre
			err = nil
			return err
		} else {
			current = current.next
		}
	}
	err = errors.New("Delete an item doesn't exist")
	return err
}

func (list *List) Appeal() {
	current := new(Node)
	current = list.head
	for current.next != list.head {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println(current.data)
}

func (list *List) AppealReverse() {
	current := new(Node)
	current = list.head.pre
	for current != list.head {
		fmt.Println(current.data)
		current = current.pre
	}
	fmt.Println(current.data)
}
