
n = 5
k = 2
def findTheWinner(n,k):
    ls = [i for i in range(1,n+1)]
    s=0
    while len(ls) != 1:
        rem = (s + k-1) % len(ls)
        ls.pop(rem)
        s = rem
        
    return ls[0]

print(findTheWinner(n,k))