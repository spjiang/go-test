package iterator

import "testing"

func TestIterator(t *testing.T) {
	var aggregate Aggregate
	aggregate = NewNumbers(1, 10)

	IteratorPrint(aggregate.Iterator())
}
