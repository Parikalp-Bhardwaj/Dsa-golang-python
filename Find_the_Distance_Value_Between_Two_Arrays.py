arr1 = [4,5,8]
arr2 = [10,9,1,8]
d = 2

# arr1 = [1,4,2,3]
# arr2 = [-4,-3,6,10,20,30]
# d = 3

# arr1 = [2,1,100,3]
# arr2 = [-5,-2,10,-3,7]
# d = 6

# arr1 = [-3,-3,4,-1,-10]
# arr2 = [7,10]
# d = 12

def findTheDistanceValue(arr1,arr2,d):
    count = 0
    for i in arr1:
        for j in arr2:
            sm = abs(i - j)
            if sm <= d:
                count+=1
                break
    print(count)
    print(len(arr1)-count)

print(findTheDistanceValue(arr1,arr2,d))