# You are given a string s, where every two consecutive vertical bars '|' are grouped into a pair. In other words, the 1st and 2nd '|' make a pair, the 3rd and 4th '|' make a pair, and so forth.

# Return the number of '*' in s, excluding the '*' between each pair of '|'.

# Note that each '|' will belong to exactly one pair.



s = "l|*e*et|c**o|*de|"

def countAsterisks(s):
    
    out = 0
    bars = 0

    
    # for i in s:
    #     if i == "|":
    #         bars +=1

    #     if i == "*" and bars%2==0:
    #         out+=1

    
    i = 0
    while i < len(s):
        if s[i] =="|":
            bars+=1
        if s[i] == "*" and bars%2==0:
            out+=1
        i+=1
    print(out)

print(countAsterisks(s))

