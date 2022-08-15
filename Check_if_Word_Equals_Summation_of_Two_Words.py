# firstWord = "acb"
# secondWord = "cba"
# targetWord = "cdb"

# firstWord = "aaa"
# secondWord = "a"
# targetWord = "aab"

firstWord = "aaa"
secondWord = "a"
targetWord = "aaaa"

def isSumEqual(firstWord,secondWord,targetWord):
    d = {chr(i):j for j,i in enumerate(range(97,123))}
    print(d)
    t1 = [str(d[v]) for v in targetWord]
    f1 = [str(d[v]) for v in firstWord]
    s1 = [str(d[v]) for v in secondWord]

    total = int("".join(s1)) + int("".join(f1))
    if total == int("".join(t1)):
        return True
    return False

print(isSumEqual(firstWord,secondWord,targetWord))