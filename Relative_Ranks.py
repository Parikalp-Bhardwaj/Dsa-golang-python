score = [10,3,8,9,4]

def findRelativeRanks(score):
    s = score.copy()
    s.sort(reverse=True)
    out = []
    for i in s:
        if s.index(i) == 0:
            out.append("Gold Medal")
        elif s.index(i) == 1:
            out.append("Silver Medal")
        elif s.index(i) == 2:
            out.append("Bronze Medal")
        else:
            out.append(str(s.index(i)+1))
    return out


print(findRelativeRanks(score))