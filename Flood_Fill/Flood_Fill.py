def floodFill(image,sr,sc,color):
    n = len(image)
    m = len(image[0])
    c = image[sr][sc] = color

    def floodFillUtils(image,n,m,sr,sc,color):
        if sr < 0 or sr >= n:
            return
        if sc < 0 or sc >= m:
            return

        if image[sr][sc] == color:
            return
        
        if image[sr][sc] != c:
            return

        image[sr][sc]=color
        
        floodFillUtils(image,n,m,sr-1,sc,color)
        floodFillUtils(image,n,m,sr+1,sc,color)
        floodFillUtils(image,n,m,sr,sc-1,color)
        floodFillUtils(image,n,m,sr,sc+1,color)

    floodFillUtils(image,n,m,sr,sc,color)
    return image

 





image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1
sc = 1
color = 2
print(floodFill(image,sr,sc,color))