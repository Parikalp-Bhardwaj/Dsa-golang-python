nums = [2,3,-1,8,4]

def findMiddleInde(nums):
    total = sum(nums)
    left_total = 0
    for i in range(len(nums)):
        total-=nums[i]
        if total == left_total:
            return i

        left_total += nums[i]

    print("left",left_total)
    return -1

print(findMiddleInde(nums))