package DoublyLinklist

import (
	"fmt"
	"log"
	"testing"
)

func Test_Doub(t *testing.T) {

	list := NewList()
	fmt.Println(list.IsEmpty())

	list.Insert(5)
	list.Insert(6)
	list.Insert("John")

	fmt.Println()
	log.Println("Testing for Normal Display")
	list.Appeal()
	fmt.Println()
	log.Println("Testing for Reverse Display")
	list.AppealReverse()
	fmt.Println()

	log.Println("Testing Delete item")
	list.Delete(6)
	list.Appeal()
	fmt.Println()

}
