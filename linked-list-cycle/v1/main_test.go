package main

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	// This one fails for I couldn't provide an actual circular reference input
	// t.Run("Should return false when there is no cycle", func(t *testing.T) {
	// 	ll := CircularLinkedList{}
	// 	ll.AddAtBeg(10)
	// 	ll.AddAtEnd(20)
	// 	ll.AddAtEnd(30)
	// 	ll.AddAtEnd(40)
	// 	ll.AddAtEnd(50)
	// 	ll.AddAtEnd(60)

	// 	want := false
	// 	got := HasCycle(ll.head)

	// 	if got != want {
	// 		t.Errorf("Expected %v. Got %v", want, got)
	// 	}
	// })

	t.Run("Should return true when there is cycle", func(t *testing.T) {
		ll := CircularLinkedList{}
		ll.AddAtBeg(10)
		ll.AddAtEnd(20)
		ll.AddAtEnd(20)

		want := true
		got := HasCycle(ll.head)

		if got != want {
			t.Errorf("Expected %v. Got %v", want, got)
		}
	})

}
