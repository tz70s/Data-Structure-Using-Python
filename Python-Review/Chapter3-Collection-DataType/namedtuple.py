#!/usr/bin/env python3

from collections import namedtuple

# User define tuples' index 
Named = namedtuple("Named",["First","Second","Third"])

tup = Named("hey","I'm","tuple")

print (tup.Third)

#format using
# using key-value pair mapping

print ("{First} !{Second} {Third}".format(**tup._asdict()))

hello = ['a','b','c','d','e']

print (hello[2:4])

hello.pop(3)

print (hello)

del hello[3]

print (hello)

hello = hello[::-1]

print (hello)