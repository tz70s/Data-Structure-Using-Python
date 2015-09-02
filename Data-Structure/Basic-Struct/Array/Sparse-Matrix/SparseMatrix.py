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
        return self.getvalue


if "__name__" == "__main__":
    #init Matrix
    matlist = []
    for row , col in range(10),range(10):
        matlist.append(Matrix(row,col,row*col))

    for row in range(10):
        line = []
        for col in range(10):
            line.append(matlist)


