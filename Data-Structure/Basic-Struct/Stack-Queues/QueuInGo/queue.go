package QueuInGo

import (
	"errors"
	"fmt"
)

type Queue struct {
	que []interface{}
	len int
}

func NewQue() *Queue {
	queue := &Queue{}
	queue.que = make([]interface{}, 0)
	queue.len = 0
	return queue
}

func (queue *Queue) IsEmpty() bool {
	return queue.len == 0
}

func (queue *Queue) Pop() (err error) {
	if queue.IsEmpty() != true {
		queue.que = queue.que[:queue.len-1]
		queue.len -= 1
	} else {
		err = errors.New("Attempting Pop From Empty Queue")
	}

	return
}

func (queue *Queue) Push(item interface{}) {
	if queue.IsEmpty() == true {
		queue.que = append(queue.que, item)
		queue.len += 1
	} else {

		var tmpque = make([]interface{}, 0)
		tmpque = append(tmpque, item)
		queue.que = append(tmpque, queue.que...)
		queue.len += 1
	}
}

func (queue *Queue) Appear() {
	for _, v := range queue.que {
		fmt.Println(v, " ")
	}
}
