groupSizes = [3,3,3,3,3,1,3]
# groupSizes = [8,8,8,8,8,8,8,8]

def groupThePeople(groupSizes):
    res = []
    for i in range(len(groupSizes)+1):
        ls = []
        for j in range(len(groupSizes)):
            if groupSizes[j]==i:
                if len(ls) >= i:
                    res.append(ls)
                    ls = []
                ls.append(j)
        if len(ls) >= 1:
            res.append(ls)

    return res

print(groupThePeople(groupSizes))