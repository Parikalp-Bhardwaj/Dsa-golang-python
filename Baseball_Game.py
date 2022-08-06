'''

You are keeping score for a baseball game with strange rules. The game consists of several rounds, where the scores of past rounds may affect future rounds' scores.

At the beginning of the game, you start with an empty record. You are given a list of strings ops, where ops[i] is the ith operation you must apply to the record and is one of the following:

An integer x - Record a new score of x.
"+" - Record a new score that is the sum of the previous two scores. It is guaranteed there will always be two previous scores.
"D" - Record a new score that is double the previous score. It is guaranteed there will always be a previous score.
"C" - Invalidate the previous score, removing it from the record. It is guaranteed there will always be a previous score.
Return the sum of all the scores on the record. The test cases are generated so that the answer fits in a 32-bit integer.



'''






ops = ["5","2","C","D","+"]
def calPoints(ops):
    
    ls = []
    k = 0
    for i in range(len(ops)):
        if ops[i] != "C" and ops[i] != "D" and ops[i] != "+":
            ls.append(int(ops[i]))
        
        elif ops[i] == "C":
            ls.remove(ls[len(ls) -1] )

        elif ops[i] == "D":
            ls.append(2*ls[len(ls) -1])

        elif ops[i] == "+":
            ls.append(ls[len(ls) -1] + ls[len(ls) -2])

    return sum(ls)


print(calPoints(ops))
