from collections import defaultdict
def validPath(n,edges,source,destination):
        ls = defaultdict(list)
        for edge in edges:
            u = edge[0]
            v = edge[1]
            ls[u].append(v)
            ls[v].append(u)
        
        queue = []
        visited = [False] * (n+1)
        queue.append(source)
        visited[source] = True
        while queue:
            val = queue[0]
            queue.pop(0)
            temp = ls[val]
            for x in range(len(temp)):
                v = temp[x]
                if visited[v]==False:
                    queue.append(v)
                    visited[v]=True
        return visited[destination]

n=10
source = 5
destination = 9
edges = [[4,3],[1,4],[4,8],[1,7],[6,4],[4,2],[7,4],[4,0],[0,9],[5,4]]
print(validPath(n,edges,source,destination))