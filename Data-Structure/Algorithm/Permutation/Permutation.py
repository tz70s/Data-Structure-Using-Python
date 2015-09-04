#!/usr/bin/env python3
def Permutation(input_list,start,end):
    if start == end :
        print (input_list)
    else:
        for j in range(start,end):
            input_list[start],input_list[j] = input_list[j],input_list[start]
            Permutation(input_list,start+1,end)
            input_list[start],input_list[j] = input_list[j],input_list[start]

newList = list("abcd")
Permutation(newList,0,4)


