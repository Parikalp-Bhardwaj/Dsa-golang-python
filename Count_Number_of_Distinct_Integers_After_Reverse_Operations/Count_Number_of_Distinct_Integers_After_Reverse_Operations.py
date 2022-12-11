
nums = [1,13,10,12,31]


def countDistinctIntegers(nums):
    s2 = set()
    s1 = set(nums)
    for i in nums:
        str1 = str(i)
        s1.add(int(str1[::-1]))

    return len(s1)

print(countDistinctIntegers(nums))