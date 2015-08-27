#!/usr/bin/env python3
def BinarySearchIterative(input_list,target):
    high = len(input_list)-1
    low = 0
    while(low!=high):
        mid = ( high + low )//2
        if target > input_list[mid]:
            low = mid + 1
        elif target == input_list[mid]:
            print ("BinarySearch Success!")
            return input_list[mid]
            break
        else:
            high = mid-1
    return 0
newList = [1,2,3,4,5,6,7,8,9]
print (BinarySearchIterative(newList,7))

