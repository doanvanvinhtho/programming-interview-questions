// Given a linked list, check if the linked list has loop or not

// Floyd’s Cycle-Finding Algorithm (This is the fastest method)
// - Traverse linked list using two pointers.
// - Move one pointer(slow_p) by one and another pointer(fast_p) by two.
// - If these pointers meet at the same node then there is a loop.
//   If pointers do not meet then linked list doesn’t have a loop

package clinkedlist

type node struct {
	value int
	next  *node
}

func push(n *node, value int) *node {
	newNode := &node{
		value: value,
		next:  n,
	}
	return newNode
}

func detectLoopInLinkedList(n *node) *node {
	slowPointer := n
	fastPointer := n

	for slowPointer != nil && fastPointer != nil && fastPointer.next != nil {
		slowPointer = slowPointer.next
		fastPointer = fastPointer.next.next
		if slowPointer == fastPointer {
			return slowPointer
		}
	}

	return nil
}
