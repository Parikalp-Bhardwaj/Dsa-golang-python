# s = "aaabbb"
s = "abab"

def checkString(s):
    f = -1
    for i in range(len(s)-1):
        if s[i] == "b":
            f = i
            break
    print(f)
    if f == -1:
        return True
    
    for i in range(len(s)-1,-1,-1):
        if s[i] == "a":
            return f > i
    
    return True

print(checkString(s))