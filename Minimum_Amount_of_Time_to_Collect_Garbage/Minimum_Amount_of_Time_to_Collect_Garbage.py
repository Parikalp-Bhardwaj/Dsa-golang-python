garbage = ["MMM","PGM","GP"]
travel = [3,10]

def garbageCollection(garbage,travel):
    travel = [0]+travel
    track = ["G","P","M"]
    count=0
    for i in range(len(track)):
        k = 0
        res = 0
        for j in range(len(garbage)):
            if track[i] in garbage[j]:
                k=j
                res += garbage[j].count(track[i])

        count += res+sum(travel[:k+1])
    return count


print(garbageCollection(garbage,travel))