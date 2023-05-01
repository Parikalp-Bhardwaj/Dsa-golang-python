from collections import defaultdict
def findCenter(edges):
    matrix = defaultdict(list)
    for edge in edges:
        u = edge[0]
        v = edge[1]
        matrix[u].append(v)
        matrix[v].append(u)

    l = len(matrix)-1
    print(matrix,l)
    for key in matrix.keys():
        if len(matrix[key]) == l:
            return key




# edges = [[1,2],[2,3],[4,2]]
edges = [[1,2],[5,1],[1,3],[1,4]]
print(findCenter(edges))