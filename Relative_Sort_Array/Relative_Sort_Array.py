# Given two arrays arr1 and arr2, the elements of arr2 are distinct, 
# and all elements in arr2 are also in arr1.
# Sort the elements of arr1 such that the relative ordering of items in arr1 are 
# the same as in arr2. Elements that do not appear in arr2 should be placed at the end 
# of arr1 in ascending order.





# arr1 = [2,3,1,3,2,4,6,7,9,2,19]
# arr2 = [2,1,4,3,9,6]


# arr1 = [28,6,22,8,44,17]
# arr2 = [22,28,8,6]


def relativeSortArray(arr1,arr2):
    ls = []
    for i in range(len(arr2)):
        for j in range(len(arr1)):
            if arr2[i] == arr1[j]:
                ls.append(arr1[j])


    ls2 = sorted([i for i in arr1 if i not in ls])
    return ls + ls2

    

    ls = []
    for i in arr2:
        while i in arr1:
            ls.append(i)
            arr1.remove(i)

    return ls+sorted(arr1)


print(relativeSortArray(arr1,arr2))