s = "0110"
n = 3
# s = "0110"
# n = 4

# s="1111000101"
# n=5
# s="10010111100001110010"
# n=10
# s = "00110011"
# n=0

def queryString(s,n):

    # print(len(s)-1,n)
    # if len(s)-1 < n:
    #     return False

    
    # else:
    #     ls = [True if s[i]=="0" or s[i]=="1" else False for i in range(n)]
    #     return all(ls)

    for i in range(1,n+1):
        b=str(bin(i))
        b1=b.split("b")[1]
        if b1 not in s:
            return False
    return True


print(queryString(s,n))