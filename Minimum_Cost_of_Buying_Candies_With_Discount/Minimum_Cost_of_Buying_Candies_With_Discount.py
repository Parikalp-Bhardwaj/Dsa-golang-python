# cost = [1,2,3]
cost = [6,5,7,9,2,2]

def minimumCost(cost):
    cost.sort(reverse=True)
    ls = []
    j = 0
    c = 0
    for i in cost:
        if j == 2:
            j=0
            continue

        c+=i
        j+=1

    return c

print(minimumCost(cost))