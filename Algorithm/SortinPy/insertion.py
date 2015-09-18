#!/usr/bin/env/python3


def insertion(li):
    print("Hello World")
    for j in range(len(li)-1):
        i = j
        val = li[i+1]
        while i>=0 :
            if val >= li[i]:
                li[i+1] = val
                break
            else:
                li[i+1] = li[i]
                i-=1

    return li

if __name__ == "__main__":
    li = [1,2,4,5,6,7,3,8]
    print(insertion(li))
    print(len(li))
