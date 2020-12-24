package twothree

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	number := 137826495
	ll := newCupsFromNum(number)

	curr := ll.root
	for i := 0; i < 100; i++ {
		curr = ll.simul(curr)
	}

	cup1 := ll.findByLabel(ll.root, 1)
	for i := 0; i < ll.count-1; i++ {
		cup1 = cup1.next
		fmt.Printf("%d", cup1.label)
	}
	fmt.Println()
}

func TestP2(t *testing.T) {
	number := 137826495
	ll := newCupsFromNum(number)

	for i := ll.maxLabel + 1; i <= 1000000; i++ {
		ll.addRight(i)
	}

	curr := ll.root
	for i := 0; i < 10000000; i++ {
		curr = ll.simul(curr)
	}

	cup1 := ll.findByLabel(ll.root, 1)
	fmt.Println(cup1.next.label * cup1.next.next.label)
}

func newCupsFromNum(number int) *cups {
	c := &cups{}
	for ; number > 0; number /= 10 {
		digit := number % 10
		c.addLeft(digit)
	}
	return c
}
