nums = [1,3,5,2,4,8,2,2]

def minMaxGame(nums):
    while len(nums) != 1:
        ls = []
        k=0
        for i in range(0,len(nums),2):
            if k%2==0:
                ls.append(min(nums[i],nums[i+1]))
            else:
                ls.append(max(nums[i],nums[i+1]))
            k+=1
        nums=ls
    return nums[0]


print(minMaxGame(nums))