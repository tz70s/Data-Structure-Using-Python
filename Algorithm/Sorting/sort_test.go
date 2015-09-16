package Sorting

import (
	"fmt"
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	list := []int{1, 3, 6, 2, 8, 4, 7, 5}
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tmplist := Newslice(list)

	log.Println("Testing for Insertion-sort")
	fmt.Println(InsertionSort(tmplist))

	if !Compare(InsertionSort(tmplist), sorted) {
		t.Error("Insertion Sort failed")
	} else {
		log.Println("original : ", list)
		log.Println("Insertion Sort Success")
	}

}
