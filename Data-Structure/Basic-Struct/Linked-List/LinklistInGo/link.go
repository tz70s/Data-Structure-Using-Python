package LinklistInGo

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
	tail *Node
	len  int
}

func New() *List {
	list := &List{}
	list.len = 0
	list.head = new(Node)
	list.tail = new(Node)
	return list
}

func (list *List) Append(item interface{}) {
	if list.len == 0 {
		tmp := &Node{item, nil}
		list.head = tmp
		list.tail = tmp
		list.len += 1
	} else {
		tmp := &Node{item, nil}
		list.tail.next = tmp
		list.tail = tmp
		list.len += 1
	}
}

func (list *List) Remove(item interface{}) {
	pre := new(Node)
	pre = list.head
	if list.len == 0 {
		fmt.Println("Nothing to be removed")
	}

	if pre.data != item && list.len == 1 {
		fmt.Println("Not Found")
		return
	}

	if pre.data == item {
		list.head = pre.next
		list.len -= 1
	} else {
		tmp := new(Node)
		tmp = pre.next
		for {

			if tmp.data == item {
				pre.next = tmp.next
				list.len -= 1
				return
			} else {
				if tmp == list.tail {
					fmt.Println("Not Found")
					break
				}
				pre = pre.next
				tmp = tmp.next
			}
		}

	}

}

func (list *List) Search(item interface{}) int {
	tmp := new(Node)
	tmp = list.head
	pos := 0
	for {
		if tmp.data == item {
			fmt.Println("Got the item")
			fmt.Println("pos : ", pos)
			return pos
		} else {
			if tmp != list.tail {
				tmp = tmp.next
				pos += 1
			} else {
				fmt.Println("The item is not in the list!")
				break
			}
		}
	}

	return 0
}

func (list *List) Appear() {
	tmp := new(Node)
	if list.len <= 1 {
		fmt.Println(list.head.data)
	} else {
		tmp = list.head
		for {

			fmt.Println(tmp.data)
			tmp = tmp.next
			if tmp == list.tail {
				fmt.Println(tmp.data)
				break
			}

		}
	}
}
