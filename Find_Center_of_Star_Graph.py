
edges = [[1,2],[2,3],[4,2]]
def findCenter(edges):
    s1 = set()
    for i,j in edges:
        if i in s1:
            return i
        if j in s1:
            return j
        s1.add(i)
        s1.add(j)
        


print(findCenter(edges))