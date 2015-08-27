#!/usr/bin/env python3
def selectionSort(input_list):
    for i in range(len(input_list)-1):
        minn = i
        for j in range(i,len(input_list)):
            if input_list[j]<input_list[minn]:
                minn = j
        temp = input_list[i]
        input_list[i] = input_list[minn]
        input_list[minn] = temp
    return input_list

newList = [1,24,3,2,7,5,8,3]
selectionSort(newList)
print(newList)
