
# arr = [40,10,20,30]
arr = [100,100,100]
# arr = [37,12,28,9,100,56,80,5,12]
def arrayRankTransform(arr):
    # arr.sort()
    s1 = sorted(list(set(arr)))
    d = {i:j+1 for j,i in enumerate(s1)}
    return [d[e] for e in arr]


print(arrayRankTransform(arr))