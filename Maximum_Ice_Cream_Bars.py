costs = [1,3,2,4,1]
coins = 7

def maxIceCream(costs,coins):
    costs.sort()

    res = 0
    for i in range(len(costs)):
        if coins >= costs[i]:
            coins -= costs[i]
            res +=1
    
    return res

print(maxIceCream(costs,coins))