
boxes = "110"
def minOperations(boxes):
    ls = []

    for i in range(len(boxes)):
        count = 0
        for j in range(len(boxes)):
            if boxes[j] == "1":
                count+=abs(i-j)

        ls.append(count)

    return ls


print(minOperations(boxes))