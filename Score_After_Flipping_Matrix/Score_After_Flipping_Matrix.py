grid = [[0,0,1,1],[1,0,1,0],[1,1,0,0]]

def matrixScore(grid):
    m = len(grid)
    n = len(grid[0])
    for i in range(m):
        if grid[i][0] == 0:
            for j in range(n):
                grid[i][j] = 1 - grid[i][j]
    


    for i in range(n):
        count = 0
        for j in range(m):
            if grid[j][i] == 1:
                count+=1
        if count < m - count:
            for j in range(m):
                grid[j][i] = 1 - grid[j][i]
    
    res = 0
    for i in range(m):
        for j in range(n):
            res += grid[i][j] * 2 ** (n-j-1)
        
    return res
    

print(matrixScore(grid))