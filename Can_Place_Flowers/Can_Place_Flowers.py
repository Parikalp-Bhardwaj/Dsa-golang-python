# flowerbed = [1,0,0,0,1]
# n = 1
# flowerbed = [1,0,0,0,1]
# n = 2

flowerbed = [1,0,0,0,0,1]
n=2

def canPlaceFlowers(flowerbed,n):
    flowerbed = [0] + flowerbed + [0]
    # flowerbed.append(0)
    print(flowerbed)
    totalL = len(flowerbed) - 1
    i = 1

    while i <= totalL -1:
        if flowerbed[i-1] == 0 and flowerbed[i] ==0 and flowerbed[i+1] == 0:
            flowerbed[i] = 1
            n -=1
        
        if n <= 0:
            return True

        i+=1
    
    print(flowerbed)
    return False
print(canPlaceFlowers(flowerbed,n))