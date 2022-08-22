
nums = [3,2,4]
target = 6


def twoSum(nums,target):
    hashMap = {}

    for i in range(len(nums)):
        hashMap[nums[i]] = i

    ls= []
    for i in range(len(nums)):
        rem =  target -nums[i] 
        if rem in hashMap and hashMap[rem] != i:
            return [i,hashMap[rem]]


print(twoSum(nums,target))