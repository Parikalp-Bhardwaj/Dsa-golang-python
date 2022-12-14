grid=[[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]


def maxIncreaseKeepingSkyline(grid):
    row = [0]*len(grid)
    col = [0]*len(grid)
    for i in range(len(grid)):
        for j in range(len(grid)):
            row[i] = max(row[i],grid[i][j])
            col[j] = max(col[j],grid[i][j])
            
    out=0
    for i in range(len(grid)):
        for j in range(len(grid)):
            out+=min(col[j],row[i])-grid[i][j]
    return out


print(maxIncreaseKeepingSkyline(grid))