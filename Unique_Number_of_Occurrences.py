


arr = [1,2,2,1,1,3]
def uniqueOccurrences(arr):
    d={}
    for i in arr:
        if i in d:
            d[i] += 1
        else:
            d[i] = 1

    s1 = list(set(d.values()))
    if len(d) == len(s1):
        return True
    return False
    


print(uniqueOccurrences(arr))
