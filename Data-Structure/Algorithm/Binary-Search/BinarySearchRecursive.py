def BinarySearchRecursive(input_list,target,high,low):
    mid = (high+low)//2
    if target > input_list[mid]:
        return BinarySearchRecursive(input_list,target,high,mid+1)
    elif target == input_list[mid]:
        print ("BinarySearch Success")
        return input_list[mid]
    else:
        return BinarySearchRecursive(input_list,target,mid-1,low)
    if high==low:
        return 0

newList = [1,2,3,4,5,6,8,9,10,14,16]
print (BinarySearchRecursive(newList,10,len(newList)-1,0))

