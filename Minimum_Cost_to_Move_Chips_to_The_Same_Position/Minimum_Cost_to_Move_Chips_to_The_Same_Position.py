position = [1,2,3]

def minCostToMoveChips(position):
    even,odd = 0,0
    for pos in position:
        if pos%2==0:
            even+=1
        else:
            odd+=1
    
    return min(even,odd)

print(minCostToMoveChips(position))