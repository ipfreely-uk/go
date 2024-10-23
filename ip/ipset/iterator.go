package ipset

import (
	"iter"

	"github.com/ipfreely-uk/go/ip"
)

// Iterator function that returns element and whether element returned.
//
// Mutates internal state until exhausted.
// When exists is false the iterator is exhausted and element should not be used.
type Iterator[E any] func() (element E, exists bool)

func emptyIterator[E any]() Iterator[E] {
	return func() (E, bool) {
		var nothing E
		return nothing, false
	}
}

func sliceIterator[E any](slice []E) Iterator[E] {
	var index = 0

	return func() (E, bool) {
		var present bool = false
		var element E
		if len(slice) != index {
			present = true
			element = slice[index]
			index++
		}
		return element, present
	}
}

func addressIterator[A ip.Number[A]](first, last A) Iterator[A] {
	var current A = first
	var done bool = false

	return func() (A, bool) {
		var element A
		if done {
			return element, false
		}
		r := current
		done = ip.Eq(current, last)
		current = ip.Next(current)
		return r, true
	}
}

func ranges2AddressIterator[A ip.Number[A]](slice []Interval[A]) Iterator[A] {
	ranges := sliceIterator(slice)
	rnge, rok := ranges()
	addresses := rnge.Addresses()

	return func() (A, bool) {
		var result A
		for rok {
			result, aok := addresses()
			if aok {
				return result, true
			}
			rnge, rok = ranges()
			if !rok {
				break
			}
			addresses = rnge.Addresses()
		}
		return result, false
	}
}

// Adapts [Iterator] for use in range loops
func Seq[A any](i Iterator[A]) iter.Seq[A] {
	return func(yield func(A) bool) {
		for {
			a, ok := i()
			if !ok {
				return
			}
			if !yield(a) {
				return
			}
		}
	}
}
