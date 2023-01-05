package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, nil}}}}}}}}
	val := 6
	head = (removeElements(head, val))
	l := make([]int, 0)
	current := head
	for current != nil {
		l = append(l, current.Val)
		current = current.Next
	}
	fmt.Println(l)

}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} else if head.Val == val {
		return removeElements(head.Next, val)
	} else {
		head.Next = removeElements(head.Next, val)
		return head
	}
}
