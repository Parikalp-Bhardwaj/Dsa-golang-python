queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"]
pattern = "FoBa"

def camelMatch(queries,pattern):
    out = []
    for query in queries:
        count = 0
        flag = 0
        ans = ""
        for j in range(len(query)):
            
            if count < len(pattern):
                if query[j] == pattern[count]:
                    ans+=query[j]
                    count +=1
                elif query[j] == query[j].upper():
                    flag = 1
                    break

            elif query[j:] != query[j:].lower():
                flag=1
                break

            else:
                break
        
        if flag == 0 and ans == pattern:
            out.append(True)
        else:
            out.append(False)
    
    return out

print(camelMatch(queries,pattern))