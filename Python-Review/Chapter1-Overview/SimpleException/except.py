#! usr/bin/env python3

s = input("Enter an integer\n")
try:
    i = int(s)
    print("Valid integer!")
except ValueError as err:
    print(err)
