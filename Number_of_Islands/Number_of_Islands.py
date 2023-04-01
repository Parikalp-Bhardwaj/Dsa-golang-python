def numIslands(grid):
    n = len(grid)
    m = len(grid[0])
    visited = [[False for i in range(n)] for j in range(m)]

    def numIslandsUtils(i,j,grid,visited):
        if grid[i][j] == "0" or visited[i][j] == True:
            return

        visited[i][j] = True
        if i+1 < len(grid):
            numIslandsUtils(i+1,j,grid,visited)
        if i-1 < len(grid):
            numIslandsUtils(i-1,j,grid,visited)
        if j+1 < len(grid):
            numIslandsUtils(i,j+1,grid,visited)
        if j-1 < len(grid):
            numIslandsUtils(i,j-1,grid,visited)
        
    ans = 0
    for i in range(n):
        for j in range(m):
            if grid[i][j] == '1' and visited[i][j]==False:
                ans+=1
                numIslandsUtils(i,j,grid,visited)


    

    return ans




# grid = [
#   ["1","1","1","1","0"],
#   ["1","1","0","1","0"],
#   ["1","1","0","0","0"],
#   ["0","0","0","0","0"]
# ]
grid =[["1","1","1","1","1","0","1","1","1","1"],["1","0","1","0","1","1","1","1","1","1"],["0","1","1","1","0","1","1","1","1","1"],["1","1","0","1","1","0","0","0","0","1"],["1","0","1","0","1","0","0","1","0","1"],["1","0","0","1","1","1","0","1","0","0"],["0","0","1","0","0","1","1","1","1","0"],["1","0","1","1","1","0","0","1","1","1"],["1","1","1","1","1","1","1","1","0","1"],["1","0","1","1","1","1","1","1","1","0"]]
print(numIslands(grid))
