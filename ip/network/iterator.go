package network

import (
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

// Iterator function that returns whether element returned and element
type Iterator[E any] func() (E, bool)

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

func addressIterator[A ip.Address[A]](first, last A) Iterator[A] {
	var current A = first
	var done bool = false

	return func() (A, bool) {
		var element A
		if done {
			return element, false
		}
		r := current
		done = compare.Eq(current, last)
		current = ip.Next(current)
		return r, true
	}
}

func ranges2AddressIterator[A ip.Address[A]](slice []AddressRange[A]) Iterator[A] {
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
