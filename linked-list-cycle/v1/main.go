package main

import "fmt"

type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

// to avoid mistakes when using pointer vs struct for new node creation
func NewListNode(val int) *ListNode {
	n := &ListNode{}
	n.Val = val
	n.Next = nil
	n.Prev = nil
	return n
}

type CircularLinkedList struct {
	head *ListNode
	tail *ListNode
}

func (ll *CircularLinkedList) AddAtBeg(val int) *ListNode {
	n := NewListNode(val)
	n.Next = ll.head
	ll.head = n
	if ll.tail == nil {
		ll.tail = n
	}
	n.Prev = ll.tail
	ll.tail.Next = n

	return n
}

func (ll *CircularLinkedList) DelAtBeg() int {
	if ll.head == nil {
		return -1
	}
	if ll.head == ll.tail {
		value := ll.head
		ll.head = nil
		ll.tail = nil
		return value.Val
	}
	cur := ll.head
	ll.head = cur.Next
	if ll.head != nil {
		ll.head.Prev = ll.tail
		ll.tail.Next = ll.head
	}
	return cur.Val
}

func (ll *CircularLinkedList) AddAtEnd(val int) *ListNode {
	n := NewListNode(val)
	if ll.head == nil {
		ll.head = n
		ll.tail = n
		return n
	}
	if ll.head == ll.tail {
		ll.head.Next = n
		ll.head.Prev = n
		n.Next = ll.head
		n.Prev = ll.head
		ll.tail = n
		return n
	}
	cur := ll.head
	for ; cur.Next != ll.tail; cur = cur.Next {
	}
	cur.Next.Next = n
	n.Prev = cur.Next
	n.Next = ll.head
	ll.tail = n
	return n
}

func (ll *CircularLinkedList) DelAtEnd() int {
	// no item
	if ll.head == nil {
		return -1
	}
	// only one item
	if ll.head == ll.tail {
		return ll.DelAtBeg()
	}
	// more than one, go to second last
	cur := ll.head
	for ; cur.Next != ll.tail; cur = cur.Next {
	}
	retval := cur.Next.Val
	cur.Next = ll.head
	ll.tail = cur
	ll.head.Prev = cur
	return retval
}

func (ll *CircularLinkedList) Count() int {
	var c int = 0
	for cur := ll.head; cur != ll.tail; cur = cur.Next {
		c++
	}
	return c
}

func (ll *CircularLinkedList) Reverse() {
	var prev, next *ListNode
	if ll.head == ll.tail {
		return
	}
	cur := ll.head.Next
	ll.tail = ll.head
	if ll.head.Next != nil {
		for cur != ll.head {
			next = cur.Next
			cur.Next = prev
			cur.Prev = next
			prev = cur
			cur = next
		}
	}
	ll.head = prev
	cur.Next = ll.tail.Next
	ll.tail.Next.Next = ll.tail
	ll.tail.Next = ll.head
	ll.head.Prev = ll.tail
}

func (ll *CircularLinkedList) Display() {
	if ll.head != nil {
		fmt.Println(ll.head.Val, " ")
		if ll.head.Next != nil {
			for cur := ll.head.Next; cur != ll.head; cur = cur.Next {
				fmt.Println(cur.Val, " ")
			}
		}
	}
	fmt.Println("")
}

func (ll *CircularLinkedList) DisplayReverse() {
	if ll.head == nil {
		return
	}
	var cur *ListNode
	for cur = ll.head; cur.Next != ll.tail; cur = cur.Next {
	}
	for ; cur != ll.tail; cur = cur.Prev {
		fmt.Println(cur.Val, " ")
	}
	fmt.Println("")
}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	current := head
	next := head.Next

	for current != next {
		if next == nil || next.Next == nil {
			return false
		}
		current = current.Next
		next = next.Next.Next
	}

	return true
}

func main() {
	ll := CircularLinkedList{}
	ll.AddAtBeg(10)
	ll.AddAtEnd(20)
	ll.AddAtEnd(30)
	fmt.Println(HasCycle(ll.head))
	ll.Display()

	ll = CircularLinkedList{}
	ll.AddAtBeg(10)
	ll.AddAtEnd(20)
	ll.AddAtEnd(10)
	fmt.Println(HasCycle(ll.head))
	ll.Display()

	// ll.Display()
	// ll.Reverse()
	// ll.Display()
	// fmt.Println("Deleting an element at the beginning : ", ll.DelAtBeg())
	// fmt.Println("Deleting an element at the end : ", ll.DelAtEnd())
	// ll.Display()
}
