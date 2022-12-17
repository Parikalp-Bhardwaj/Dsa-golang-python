from collections import defaultdict
mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]

def diagonalSort(mat):
    res = mat
    d = defaultdict(list)
    for i in range(len(mat)):

        for j in range(len(mat[0])):
            d[i-j].append(mat[i][j])
    for i in d:
        d[i].sort()
    for i in range(len(mat)):
        for j in range(len(mat[0])):
            res[i][j] = d[i-j].pop(0)
    
    return res



print(diagonalSort(mat))