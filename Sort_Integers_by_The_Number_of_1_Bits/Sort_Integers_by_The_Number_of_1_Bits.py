# You are given an integer array arr. Sort the integers in the array in ascending 
# order by the number of 1's in their binary representation and in case 
# of two or more integers have the same number of 1's you have to sort them in ascending order.
# Return the array after sorting it.




# arr = [0,1,2,3,4,5,6,7,8]
# [0,1,2,4,8,3,5,6,7]

arr = [1024,512,256,128,64,32,16,8,4,2,1]

# arr=[2,3,5,7,11,13,17,19]

# arr = [10000,10000]


def sortByBits(arr):
    arr.sort()
    out = []
    for i in arr:
        out.append(bin(i).count("1")) 

    ls=[]
    for i,j in zip(arr,out):
        ls.append((i,j))
    

    a = sorted(ls, key=lambda x: x[1])   
    ls2 = [i[0] for i in a]
    return ls2

print(sortByBits(arr))