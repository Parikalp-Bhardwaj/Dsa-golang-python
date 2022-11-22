logs = [[0,5],[1,2],[0,2],[0,5],[1,3]]
k = 5

def findingUsersActiveMinutes(logs,k):
    res = [0] * k
    users = {}
    for id,time in logs:
        if id not in users:
            users[id] = set()
            users[id].add(time)
        else:
            users[id].add(time)


    for user in users:
        res[len(users[user])-1]+=1

    return res


print(findingUsersActiveMinutes(logs,k))