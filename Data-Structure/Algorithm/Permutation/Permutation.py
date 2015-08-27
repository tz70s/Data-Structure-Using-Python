#!/usr/bin/env python3

def Permutation(input_list,start,end):
    if start == end :
        print (input_list)
    else:
        for j in range(start,end):
            temp = input_list[start]
            input_list[start] = input_list[j]
            input_list[j] = temp
            Permutation(input_list,start+1,end)
            temp = input_list[start]
            input_list[start] = input_list[j]
            input_list[j] = temp

newList = list("abcd")
Permutation(newList,0,4)

