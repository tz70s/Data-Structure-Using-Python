#!/usr/bin/env python3
'''
while collection data type using copy
if the right side is immutable data type like string , integer etc
the leftside will create an object reference
else if the right side is an object reference
the left side will create an object reference pointing to the same place with the right side
see the example as following
'''
# same object reference place
print ("same object reference place")
a = [1,2,3]
b = a
print ("a:",a)
print ("b:",b)
a[1] = 5
print ("a:",a)
print ("b:",b)

#shallow copy , using slice
print ("shallow copy , using slice ")
c = [1,2,3,4,[6,7,8]]
d = c[:]
print ("c:",c)
print ("d:",d)
c[1] = 5
print ("c:",c)
print ("d:",d)
d[1] = 6
print ("c:",c)
print ("d:",d)
print ("The shallow copy will create like independent ,\n however,\
       while there is a collection type nested in the collection,\n\
       there will pointing to the same reference place")
c[4][2] = 15
print ("c:",c)
print ("d:",d)
#deep copy
import copy
print ("deep copy , use the function copy.deepcopy in module copy")
e = [1,2,3,4,[6,7,8]]
f = copy.deepcopy(e)
print ("e:",e)
print ("f:",f)
e[4][2] = 15
print ("e:",e)
print ("f:",f)

