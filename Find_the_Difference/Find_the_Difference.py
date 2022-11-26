# s="abcd"
# t="abcde"
# s = ""
# t = "y"
s="a"
t="aa"
def findTheDifference(s,t):

    for i in range(len(s)):
        if s[i] != t[i]:
            return t[i]

        return t[len(t)-1]


    # t1 = list(t)
    # for i in s:
    #     if i in t:
    #         t1.remove(i)

    # return "".join(t1)

print(findTheDifference(s,t))