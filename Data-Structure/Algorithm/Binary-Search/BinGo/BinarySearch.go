package BinGo

import (
    "fmt"
    "errors"
)

func Iterate(slice [] int, target int) (int, error) {
    high := len(slice)
    low := 0

    for high != low {
        mid := (high + low)/2
        if slice[mid] == target {
            fmt.Println("Matched!", target)
            return target, nil
        } else if slice[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    err := errors.New("Not found anything")
    return 0, err
}

func Recursive(slice [] int, target, high, low int)(int, error) {
    mid := (high + low) / 2
    //err := errors.New("Not found anything")
    if high == low && high != mid {
        panic("Not found anything")
    }

    if slice[mid] == target {
        fmt.Println("Matched!", target)
        return target, nil
    } else if slice[mid] < target {
        low = mid + 1
        return Recursive(slice, target, high, low)
    } else {
        high = mid - 1
        return Recursive(slice, target, high, low)
    }
}
