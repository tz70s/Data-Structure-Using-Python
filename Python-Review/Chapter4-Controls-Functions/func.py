#!/usr/bin/env python3

def func (*li , power=2):
    result = 0
    for numb in li:
        result += numb**power
    return result

print (func(1,2,3,4,5))
print (func(1,2,3,4,5,power=3))

fu = lambda pw : print ("2") if pw is 2 else print("not")
fu(2)
