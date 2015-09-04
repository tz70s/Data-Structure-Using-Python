#!/usr/bin/env python3

import linklist
import stack

class CreateError(Exception):pass
class IsFullError(Exception):pass


def equivalence():
    listnode = []
    for i in range(12):
        listnode.append(linklist.UnorderedList())
        listnode[i].fistNode(i)
    try:
        fh = open("eq.txt")
    except EnvironmentError as err:
        print(err)
    lines = []
    #iterate the fh
    for line in fh:
        line = line.strip("\n")
        lines.append(line.split())

    #read the first element , puts into list head
    for fir , sec in lines:
        fir = int(fir)
        sec = int(sec)
        listnode[fir].insert(sec)

    #reveal
    for i in range(12):
        listnode[i].reveal()


def main():
    equivalence()

def another():
    print ("Hello World")


if __name__ == "__main__":
    another()
    main()
