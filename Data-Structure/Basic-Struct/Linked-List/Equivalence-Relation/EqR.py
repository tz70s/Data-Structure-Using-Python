#!/usr/bin/env python3

import linklist
import stack

class CreateError(Exception):pass


def main():
    try:
        myList = linklist.UnorderedList()
        myStack = stack.Stack()
    except CreateError as err:
        print (err)

    myList.insert(5)
    myList.insert(6)
    myList.reveal()
    myStack.push(5)
    myStack.push(6)
    print(myStack.reveal())



if __name__ == "__main__":
    main()
