nums = [3,6,1,2,5]
k = 2

def partitionArray(nums, k):
    nums = sorted(nums,reverse=True)
    res = 1
    j = 0
    for i in range(1,len(nums)):
        if nums[j] - nums[i] > k:
            res+=1
            j=i

    return res

print(partitionArray(nums,k))