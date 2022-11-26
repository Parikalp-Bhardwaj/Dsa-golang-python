
nums = [3,1,-2,-5,2,-4]

def rearrangeArra(nums):
    ls = [0]*len(nums)
    p=0
    n=1
    for i in nums:
        if i > 0 :
            ls[p]=i
            p+=2
        else:
            ls[n]=i
            n+=2
    return ls

print(rearrangeArra(nums))