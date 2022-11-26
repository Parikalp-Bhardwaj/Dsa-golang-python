from collections import Counter


order = "cba"
s = "abcd"
def customSortString(order,s):
    str1=""
    for i in s:
        if i not in order:
            str1+=i

    res=""
    d = Counter(s)
    for i in order:
        if i in d:
            res+=i

    return res+str1

print(customSortString(order,s))