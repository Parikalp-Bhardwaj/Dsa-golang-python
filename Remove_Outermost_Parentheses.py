s = "(()())(())(()(()))"

def removeOuterParentheses(self, s)
    res = ""
    c = 0
    for i in range(len(s)):
        if s[i] == "(":
            c+=1
            if c > 1:
                res +="("
        else:
            if c > 1:
                res +=")"
            c -= 1
                
    return res


print(removeOuterParentheses(s))