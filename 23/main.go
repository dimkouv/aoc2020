package twothree

type cups struct {
	root               *cup
	final              *cup
	count              int
	minLabel, maxLabel int
}

type cup struct {
	label int
	next  *cup
	prev  *cup
}

func (ll *cups) simul(c *cup) *cup {
	targetLabel := c.label - 1
	if targetLabel < ll.minLabel {
		targetLabel = ll.maxLabel
	}

	c1 := c.next
	c2 := c1.next
	c3 := c2.next

	for c1.label == targetLabel || c2.label == targetLabel || c3.label == targetLabel {
		targetLabel -= 1
		if targetLabel < ll.minLabel {
			targetLabel = ll.maxLabel
		}
	}

	c.next = c3.next

	place := ll.findByLabel(c, targetLabel)
	place.next, c3.next = c1, place.next
	return c.next
}

func (ll *cups) findByLabel(startFrom *cup, label int) *cup {
	bwSearch, fwSearch := startFrom, startFrom
	for {
		bwSearch = bwSearch.prev
		if bwSearch.label == label {
			return bwSearch
		}
		fwSearch = fwSearch.next
		if fwSearch.label == label {
			return fwSearch
		}
	}
}

func (ll *cups) addLeft(num int) {
	c := &cup{label: num}

	ll.updateStats(c)
	ll.root = c
	ll.final.next, c.prev = ll.root, ll.final
	ll.count++
}

func (ll *cups) addRight(num int) {
	c := &cup{label: num}

	ll.updateStats(c)
	ll.final.next, c.prev = c, ll.final
	ll.final = c
	ll.count++
}

func (ll *cups) updateStats(c *cup) {
	if ll.count == 0 {
		ll.final = c
		ll.minLabel = c.label
		ll.maxLabel = c.label
	}
	if c.label < ll.minLabel {
		ll.minLabel = c.label
	}
	if c.label > ll.maxLabel {
		ll.maxLabel = c.label
	}
	if ll.root != nil {
		c.next = ll.root
	}
}
