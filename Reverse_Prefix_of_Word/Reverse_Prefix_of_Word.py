word = "abcdefd"
ch = "d"
# word = "abcd"
# ch = "z"
def reversePrefix(word,ch):
    w1 = list(word)
   
    for i in w1:
        if i == ch:
            idx = w1.index(ch)+1
            return "".join(w1[:idx][::-1]+w1[idx:])
        
    return "".join(w1)


print(reversePrefix(word,ch))