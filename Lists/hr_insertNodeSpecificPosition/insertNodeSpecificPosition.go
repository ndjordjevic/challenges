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

	h := insertNodeAtPosition(&n3, 7, 1)

	printLinkedList2(h)
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	cur := llist
	var prev *SinglyLinkedListNode

	for i := 0; cur != nil; i++ {
		if int32(i) == position {
			n := SinglyLinkedListNode{
				data,
				cur,
			}

			if prev == nil {
				llist = &n
			} else {
				prev.next = &n
			}

			break
		}

		prev = cur
		cur = cur.next
	}

	return llist
}

func printLinkedList2(head *SinglyLinkedListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}

}
