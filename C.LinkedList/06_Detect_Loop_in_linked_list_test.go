// Given a linked list, check if the linked list has loop or not

// Floyd’s Cycle-Finding Algorithm (This is the fastest method)
// - Traverse linked list using two pointers.
// - Move one pointer(slow_p) by one and another pointer(fast_p) by two.
// - If these pointers meet at the same node then there is a loop.
//   If pointers do not meet then linked list doesn’t have a loop

package clinkedlist

import (
	"testing"
)

func TestDetectLoopInLinkedList(t *testing.T) {
	head := push(nil, 5)
	head = push(head, 4)
	head = push(head, 3)
	head = push(head, 2)
	head = push(head, 1)

	if n := detectLoopInLinkedList(head); n != nil {
		t.Errorf("Wrong, the list hasn't loop")
	}

	// Create loop point
	// 1    2    3    4    5  nil                3
	head.next.next.next.next.next = head.next.next
	if n := detectLoopInLinkedList(head); n == nil {
		t.Errorf("Wrong, the list has loop")
	}
}
