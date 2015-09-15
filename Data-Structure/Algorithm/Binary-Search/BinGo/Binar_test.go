package BinGo

import (
    "testing"
    "log"
)

func Test_Bi(t *testing.T) {
    sli := []int{1,2,3,4,5,6,7,8,11}
    if _, err := Iterate(sli,5) ; err != nil {
        t.Error(err)
    } else {
        log.Println("End of iterate Binary Search")
    }

    if _, err := Recursive(sli, 3, len(sli), 0); err != nil {
        t.Error(err)
    } else {
        log.Println("End of Recursive Binary Search")
    }

}
