package LinklistInGo

import (
	"fmt"
	"log"
	"testing"
)

func TestLink(t *testing.T) {
	fmt.Println()
	log.Println("Sleepy now ... , Im lazy to accomplish the test file, bye")
	fmt.Println()
	list := New()
	list.Append("1")
	list.Append(2)
	list.Append("John")
	list.Append("Quincy")
	list.Append(true)

	list.Appear()
	list.Search(2)
	list.Remove(2)
	list.Appear()
}
