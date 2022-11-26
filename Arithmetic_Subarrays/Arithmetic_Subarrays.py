nums = [-12,-9,-3,-12,-6,15,20,-25,-20,-15,-10]
l = [0,1,6,4,8,7]
r = [4,4,9,7,9,10]


def checkArithmeticSubarrays(nums,l,r):
    out = []
    for i in range(len(r)):
        ls = []
        
        for j in range(l[i],r[i]+1):
            ls.append(nums[j])

        ls.sort()
        
        res = []
        for v in range(len(ls)-1):
            val = ls[v]-ls[v+1]
            res.append(val)

        if len(set(res)) == 1:
            out.append(True)
        else:
            out.append(False)
    return out


print(checkArithmeticSubarrays(nums,l,r))