package main

import "fmt"

func main() {
	var head = new(ListNode)
	head.Val = 1
	head.Next = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	reversed := reverseList3(head)

	for reversed != nil {
		fmt.Printf("%v\n", reversed)
		reversed = reversed.Next
	}
}

// ListNode list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var list *ListNode

	for head != nil {
		list = &ListNode{Val: head.Val, Next: list}
		head = head.Next
	}

	return list
}

func reverseList2(head *ListNode) *ListNode {
	var reverse func(head, tail *ListNode) *ListNode
	reverse = func(head, tail *ListNode) *ListNode {
		if head == nil {
			return tail
		}
		next := head.Next
		head.Next = tail
		return reverse(next, head)
	}
	return reverse(head, nil)
}

func reverseList3(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	var next *ListNode

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	head = prev

	return head
}
