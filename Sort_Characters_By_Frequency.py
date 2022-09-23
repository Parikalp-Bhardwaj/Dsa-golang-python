from collections import Counter
s="tree"

def frequencySort(s):
    s1 = Counter(sorted(s))
    str1=""
    for i,j in s1.most_common():
        str1+=i*j
    return str1

print(frequencySort(s))