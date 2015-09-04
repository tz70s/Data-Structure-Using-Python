#!/usr/bin/env python3

class NotFoundError(Exception):pass

class Node:
    def __init__(self, initdata):
        self.data = initdata
        self.next = None

    def getData(self):
        return self.data

    def getNext(self):
        return self.next

    def setData(self, newdata):
        self.data = newdata

    def setNext(self, newnext):
        self.next = newnext

class UnorderedList:
    def __init__(self):
        self.head = None
        self.end = None

    def isEmpty(self):
        return self.head == None

    def fistNode(self, item):
        temp = Node(item)
        self.head = temp
        self.end = temp

    def insert(self, item):
        if self.head == None:
            return self.fistNode(item)
        else:
            temp = Node(item)
            self.end.setNext(temp)
            self.end = temp

    def reveal(self):
        temp = self.head
        st = ''
        while temp != None:
            st += str(temp.getData()) + ' -> '
            temp = temp.getNext()
        print (st)

    def search(self, item):
        pos = 1
        temp = self.head
        while temp != None:
            if temp.getData() == item:
                return item , pos
            else:
                temp = temp.getNext()
                pos += 1
        raise FullError()

    def delete(self):
        temp = self.head
        while temp.getNext() != self.end:
            temp = temp.getNext()
        end = temp
        temp.setNext(None)
        
    def delNumber(self,item):
        pre = self.head
        now = pre.getNext()
        nex = now.getNext()
        while now is not None:
            if item == now.getData():
                pre.setNext(nex)
                return 0
            else:
                pre = pre.getNext()
                now = pre.getNext()
                nex = now.getNext()
        raise FullError()

