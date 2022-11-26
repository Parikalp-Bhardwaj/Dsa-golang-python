stones = [2,7,4,1,8,1]
def lastStoneWeight(stones):
    while len(stones) > 1:
        stones.sort(reverse=True)
        if stones[0] == stones[1]:
            stones.pop(0)
            stones.pop(0)

        else:
            stones[0] = stones[0] - stones[1]
            stones.pop(1)

    if len(stones) > 0:
        return (stones[0])
        
    return 0


print(lastStoneWeight(stones))