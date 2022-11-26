
words = ["abc","deq","mee","aqq","dkd","ccc"]
pattern = "abb"

# words = ["a","b","c"]
# pattern = "a"


def nextGreaterElement(words,pattern):
    def get_pattern(w):
        res = []
        d = {}
        i=0
        for w1 in w:
            if w1 in d:
                res.append(d[w1])
            else:
                i+=1
                d[w1] = i
                res.append(i)

        return res

    p = get_pattern(pattern)

    out = []
    for w2 in words:
        if get_pattern(w2) == p:
            out.append(w2)

    return out


print(nextGreaterElement(words,pattern))