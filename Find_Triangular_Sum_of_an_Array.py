nums = [1,2,3,4,5]
def triangularSum(nums):
    n = len(nums)

    while n > 1:
        ls = []
        for i in range(1,n):
            ls.append((nums[i] + nums[i-1])%10)

        n-=1
        nums=ls
        
    return nums[0]

print(triangularSum(nums))