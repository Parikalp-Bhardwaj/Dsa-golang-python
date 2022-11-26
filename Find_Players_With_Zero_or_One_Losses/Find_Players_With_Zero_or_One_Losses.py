matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]

def findWinners(matches):
    w=set()
    l={}
    for i in matches:
        if i[1] in l:
            l[i[1]] +=1
        else:
            l[i[1]] = 1

        w.add(i[0])
 

    w1 = [i for i in w if i not in l.keys()]
    

    l1 = []
    for i,j in l.items():
        if j==1:
            l1.append(i)

    return [sorted(w1),sorted(l1)]

print(findWinners(matches))