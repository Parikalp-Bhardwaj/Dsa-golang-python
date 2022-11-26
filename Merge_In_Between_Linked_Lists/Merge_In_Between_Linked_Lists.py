# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeInBetween(self, list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
        node,nextNode = None,None
        currentNode,count = list1,0
        while currentNode:
            if count == a-1:
                node = currentNode
            if count == b:
                nextNode = currentNode.next
                break

            count+=1
            currentNode = currentNode.next

        node.next = list2
        while list2.next:
            list2 = list2.next
        list2.next = nextNode
        return list1