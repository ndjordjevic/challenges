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

	printLinkedList2(&n3)
}

func printLinkedList(head *SinglyLinkedListNode) {
	for head.next != nil {
		fmt.Println(head.data)
		head = head.next
	}

	fmt.Println(head.data)
}

func printLinkedList2(head *SinglyLinkedListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}

}
