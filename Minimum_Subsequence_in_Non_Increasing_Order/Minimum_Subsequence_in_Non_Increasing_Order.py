nums = [4,3,10,9,8]
# # nums = [4,4,7,6,7]


def minSubsequence(nums):
    nums.sort()
    sumOfVal = sum(nums)
    ls= []
    for i in range(len(nums)-1,-1,-1):
        ls.append(nums[i])
        if sum(ls) > sumOfVal-sum(ls):
            return ls
    return ls
    
   
print(minSubsequence(nums))
