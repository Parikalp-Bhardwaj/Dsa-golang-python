# Balanced strings are those that have an equal quantity of 'L' and 'R' characters.

# Given a balanced string s, split it into some number of substrings such that:

# Each substring is balanced.
# Return the maximum number of balanced strings you can obtain.



s = "RLRRLLRLRL"
def balancedStringSplit(s):
    
    count = 0
    j = 0
    for i in s:
        if i == "L":
            j+=1
        else:
            j-=1
        
        if j == 0:
            count+=1

    print(count)

print(balancedStringSplit(s))









