ransomNote = "a"
magazine = "b"
from collections import Counter
def canConstruct(ransomNote,magazine):
    if ransomNote in magazine:
        return True


    r = Counter(ransomNote)
    for i in r:
        if magazine.count(i) < r[i]:
            return False
    return True

print(canConstruct(ransomNote,magazine))