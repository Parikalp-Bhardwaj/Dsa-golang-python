nums = [4,3,2,7,8,2,3,1]
def findDuplicates(nums):
    ls = []
    s1 = set()
    for num in nums:
        if num in s1:
            ls.append(num)
        else:
            s1.add(num)
    return ls

print(findDuplicates(nums))