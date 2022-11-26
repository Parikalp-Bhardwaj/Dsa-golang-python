s = "leetcode"
t = "coats"

def minSteps(s,t):
    m = {}
    for i in s:
        if i in m:
            m[i] +=1
        else:
            m[i] = 1

    for t1 in t:
        if t1 in m:
            m[t1] -= 1
        else:
            m[t1] = -1

    res=0
    for i,j in m.items():
        res+=abs(j)

    return res


print(minSteps(s,t))