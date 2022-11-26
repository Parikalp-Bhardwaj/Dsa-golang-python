
tickets = [5,1,1,1]
k = 0
def timeRequiredToBuy(ickets, k):
    t=0
    while tickets[k] != 0:
        for i in range(len(tickets)):
            if tickets[i] !=0 and tickets[k] !=0:
                tickets[i] -= 1
                t+=1
    return t

print(timeRequiredToBuy(tickets,k))