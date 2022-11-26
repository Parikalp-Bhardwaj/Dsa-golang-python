num = "2932"
def minimumSum(num):
    n1=sorted(list(map(int,str(num))))
    s1,s2 ="",""
    for i in range(len(n1)):
        if i%2 ==0 :
            s1 += str(n1[i])
        else:
            s2 += str(n1[i])

    return (int(s1)+int(s2))
    


print(minimumSum(num))