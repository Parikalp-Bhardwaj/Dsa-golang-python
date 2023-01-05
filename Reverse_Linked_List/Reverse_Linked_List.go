package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{6, nil}}}}}}}}}}
	head = (reverseList(head))
	current := head
	l := make([]int, 0)
	for current != nil {
		l = append(l, current.Val)
		current = current.Next
	}

	fmt.Println(l)
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode = nil
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
