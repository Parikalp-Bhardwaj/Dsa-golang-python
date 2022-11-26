nums = [9,12,5,10,14,3,10]
pivot = 10
def pivotArray(nums,pivot):
    ls = []
    ls2 = []
    mid = []
    for i in nums:
        if i < pivot:
            ls.append(i)
        elif i > pivot:
            ls2.append(i)
        else:
            mid.append(i)
    return ls + mid +ls2

print(pivotArray(nums,pivot))