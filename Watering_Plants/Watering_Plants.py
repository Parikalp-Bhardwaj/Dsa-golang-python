plants = [2,2,3,3]
capacity = 5

def wateringPlants(plants,capacity):
    fill = capacity
    count=0
    for i in range(len(plants)):
        if fill < plants[i]:
            fill=capacity
            count+=2*i


        fill -= plants[i]
        count+=1

    return count

print(wateringPlants(plants,capacity))


