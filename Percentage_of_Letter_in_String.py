s = "foobar"
letter = "o"


def percentageLetter(s,letter):
    count = 0
    for i in s:
        if i == letter:
            count+=1


    if count > 0:
        return int(count / len(s) *100)
    return count


print(percentageLetter(s,letter))