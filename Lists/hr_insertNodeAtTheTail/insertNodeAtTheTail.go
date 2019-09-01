package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func main() {
	n1 := SinglyLinkedListNode{
		1,
		nil,
	}

	n2 := SinglyLinkedListNode{
		2,
		&n1,
	}

	n3 := SinglyLinkedListNode{
		3,
		&n2,
	}

	h := insertNodeAtTail(&n3, 7)

	printLinkedList2(h)
}

func insertNodeAtTail(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	curr := head
	for curr.next != nil {
		curr = curr.next
	}

	n := SinglyLinkedListNode{
		data,
		nil,
	}

	curr.next = &n

	return head
}

func printLinkedList2(head *SinglyLinkedListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}

}
