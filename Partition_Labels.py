s = "ababcbacadefegdehijhklij"

def partitionLabels(s):
    d={}
    for i in range(len(s)):
        d[s[i]] = i

    res = 0
    out = 0
    ls = []
    for i,char in enumerate(s):
        res = max(d[char],res)
        if i == res:
            ls.append(len(s[out:res+1]))
            out = res+1

    return ls



print(partitionLabels(s))