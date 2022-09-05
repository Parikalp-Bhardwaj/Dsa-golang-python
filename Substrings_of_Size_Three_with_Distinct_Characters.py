s = "xyzzaz"
def countGoodSubstrings(s):
    count=0
    for i in range(len(s)):
        str1=""
        for j in range(i,len(s)):
            str1+=s[j]
            s1=list(set(str1))
            if len(s1)== 3 and len(str1) == 3:
                count+=1

    return count

print(countGoodSubstrings(s))