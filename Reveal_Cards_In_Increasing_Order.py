from collections import deque

deck = [17,13,11,2,3,5,7]

def deckRevealedIncreasing(deck):
    dq = deque()
    deck.sort()
    while deck:
        dq.rotate()
        dq.appendleft(deck.pop())
    return dq

print(deckRevealedIncreasing(deck))