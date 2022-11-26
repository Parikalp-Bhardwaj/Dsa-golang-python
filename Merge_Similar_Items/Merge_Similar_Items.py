items1 = [[1,1],[3,2],[2,3]]
items2 = [[2,1],[3,2],[1,3]]

def mergeSimilarItems(items1,items2):
    d1 = {i:j for i,j in items1}
    d2 = {i:j for i,j in items2}

    for i,j in items1:
        if i in d2:
            d2[i]=d2[i]+j 
        else:
            d2[i]=j 

    ls = [[i,j] for i,j in d2.items()]
    return sorted(ls)

print(mergeSimilarItems(items1,items2))