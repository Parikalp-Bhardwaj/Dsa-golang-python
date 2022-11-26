

words = ["cat","bt","hat","tree"]
chars = "atach"

def countCharacters(words,chars):
    s=0
    d= {}
    for i in chars:
        if i not in d:
            d[i]=1
        else:
            d[i]+=1

    for i in words:
        y = len(i)
        d1 = {}
        for j in i:
            if j not in d1:
                d1[j] = 1
            else:
                d1[j]+=1
        str1 = ""
        for l in d1:
            
            if l in d:
                if d[l] < d1[l]:
                    y=0
            
            else:
                y=0
                break

        s+=y

    return s





print(countCharacters(words,chars))