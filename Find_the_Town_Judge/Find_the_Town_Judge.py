from collections import defaultdict
def findJudge(n,trust):
    res = ([0] * (n+1))
    print(res)
    for i in trust:
        res[i[0]]-=1
        res[i[1]]+=1
    print(res)
    
    for i in range(1,n+1):
        x = res[i]
        print(x)
        if (x == n-1):
            return i

    return -1



# n = 2
# trust = [[1,2]]
n = 3 
trust = [[1,3],[2,3]]
print(findJudge(n,trust))