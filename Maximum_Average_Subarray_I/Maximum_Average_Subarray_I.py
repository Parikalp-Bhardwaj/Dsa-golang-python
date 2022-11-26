
nums = [1,12,-5,-6,50,3]
k = 4

def findMaxAverage(nums,k):
    num = sum(nums[0:k])
    total = num
    for i in range(len(nums)-k):
        num -= nums[i]
        num += nums[i+k]
        if num > total:
            total = num
    return total/k

print(findMaxAverage(nums,k))


