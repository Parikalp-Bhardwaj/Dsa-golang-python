nums = [3,5]
def specialArray(nums):
    n = max(nums)
    for i in range(n+1):
        count = 0
        for j in nums:
            if j >= i:
                count +=1

        if (count == i):
            return count

    return -1


print(specialArray(nums))