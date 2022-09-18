from collections import Counter
arr = [3,3,3,3,5,5,5,2,2,7]

def minSetSize(arr):
    d = Counter(arr)
    val = sorted(d.values())[::-1]
    out = len(arr)//2
    i = 0
    res=0
    for i in val:
        out -= i
        res +=1
        if out <= 0:
            break
                
    return res