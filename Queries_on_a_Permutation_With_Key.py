queries = [3,1,2,1]
m = 5

def processQueries(queries,m):
    p = [i for i in range(1,m+1)]
    ls = []
    for query in queries:
        ls.append(p.index(query))
        val = p.pop(p.index(query))
        p.insert(0,val)
    return ls 

print(processQueries(queries,m))

