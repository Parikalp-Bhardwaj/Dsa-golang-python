
words = ["mass","as","hero","superhero"]
def stringMatching(words):
    ls = []
    for i in range(len(words)):
        for j in range(len(words)):
            if words[i] == words[j]:
                continue
            if words[i] in words[j]:
                ls.append(words[i])
    return ls
print(stringMatching(words))