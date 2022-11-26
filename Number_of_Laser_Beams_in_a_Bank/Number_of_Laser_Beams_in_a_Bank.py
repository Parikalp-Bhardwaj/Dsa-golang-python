# bank = ["011001","000000","010100","001000"]
bank = ["000","111","000"]

# def numberOfBeams(bank):
#     res = []
#     for b1 in range(len(bank)):
#         count = 0
#         for b2 in range(len(bank[b1])):
#             if bank[b1][b2] == "1":
#                 count+=1
#         if count > 0:
#             res.append(count)

#     if len(res) == 0 or len(res) ==1:
#         return 0
#     return (sum([res[i]*res[i+1] for i in range(len(res)-1)]))



# print(numberOfBeams(bank))



def numberOfBeams(bank):
    def recu(bank,i,prev,res):
        if i == len(bank):
            return res

        
        count = bank[i].count("1")
        if count:
            res += count*prev
            prev = count

        return recu(bank,i+1,prev,res)

    return recu(bank,0,0,0)

print(numberOfBeams(bank))