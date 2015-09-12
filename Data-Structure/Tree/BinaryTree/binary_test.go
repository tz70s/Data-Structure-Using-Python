package BinaryTree

import (
	"log"
	"testing"
)

func comp(x, y interface{}) bool {

	if x.(int) < y.(int) {
		return true
	} else {
		return false
	}
}

func Test_Binar(t *testing.T) {
	tree := New(comp)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	findTree := tree.Search(2)

	if findTree.node != 2 {
		t.Error("[Error] Search error")
	} else {
		log.Println("Find tree node : ", findTree.node)
	}

}
