# piles =[2,4,1,2,7,8]
# piles = [2,4,5]
piles = [9,8,7,6,5,1,2,3,4]

def maxCoins(piles):
    piles.sort(reverse=True)
    ls=[]
    j=0
    for i in range(1,len(piles)//3+1):
        ls.append(piles[i+j])
        j+=1
    return sum(ls)

print(maxCoins(piles))