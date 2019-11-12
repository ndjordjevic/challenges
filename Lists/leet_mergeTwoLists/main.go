package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	head1 := &ListNode{}

	createList(head, 1, 2, 4)
	createList(head1, 1, 3, 4)

	printList(mergeTwoLists(head, head1))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// base case
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// dummyNode method
	dummy := &ListNode{0, nil}
	var dummyP *ListNode = dummy
	var p1 *ListNode = l1
	var p2 *ListNode = l2

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			dummyP.Next = p1
			p1 = p1.Next
		} else {
			dummyP.Next = p2
			p2 = p2.Next
		}
		dummyP = dummyP.Next
	}

	// ternary operator
	dummyP.Next = map[bool]*ListNode{true: p1, false: p2}[p1 != nil]

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}

}

func createList(head *ListNode, nums ...int) {
	var prev *ListNode

	for _, val := range nums {
		if prev == nil {
			head.Val = val
			prev = head
			continue
		}

		newNode := &ListNode{val, nil}
		prev.Next = newNode
		prev = newNode
	}
}
