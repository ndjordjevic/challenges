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

	n4 := SinglyLinkedListNode{
		4,
		&n3,
	}

	h := deleteNode(&n4, 0)

	printLinkedList2(h)
}

func deleteNode(head *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	if position == 0 {
		return head.next
	}

	cur := head

	for i := int32(0); i < position-1; i++ {
		if cur == nil {
			return head
		}
		cur = cur.next
	}

	cur.next = cur.next.next

	return head
}

func printLinkedList2(head *SinglyLinkedListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}

}
