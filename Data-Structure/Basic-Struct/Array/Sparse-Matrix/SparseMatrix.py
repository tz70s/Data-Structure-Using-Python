#!/usr/bin/env python3
from operator import attrgetter

class Matrix:
    def __init__(self, row=0, col = 0, value = 0):
        self.row = row
        self.col = col
        self.value = value

    def getrow(self):
        return self.row

    def getcol(self):
        return self.col

    def getvalue(self):
        return self.value


def matInit():
    matlist = []
    spmatlist = []
    for line in open("mat.txt"):
        matlist.append(line.split())
    for row in range(len(matlist)):
        for col in range(len(matlist)):
            if int(matlist[row][col]) != 0:
                spmatlist.append(Matrix(row,col,int(matlist[row][col])))
            else:
                continue
    for n in range(len(spmatlist)):
        print ("row: {0}, col: {1}, value: {2}".format(spmatlist[n].row,spmatlist[n].col,spmatlist[n].value))

    del matlist
    Transpose(spmatlist)

def Transpose(spmat):
    # row -> col , col -> row
    for n in spmat:
        n.row , n.col = n.col , n.row
    spmat.sort(key = attrgetter('row'))
    for n in spmat:
        print ("row: {0}, col: {1}, value: {2}".format(n.row,n.col,n.value))


if __name__ == "__main__":
    matInit()
    
