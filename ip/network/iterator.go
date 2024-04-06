package network

import (
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

// Iterator function that returns whether element returned and element
type Iterator[E any] func() (bool, E)

func emptyIterator[E any]() Iterator[E] {
	return func() (bool, E) {
		var nothing E
		return false, nothing
	}
}

func sliceIterator[E any](slice []E) Iterator[E] {
	var index = 0

	return func() (bool, E) {
		var present bool = false
		var element E
		if len(slice) != index {
			present = true
			element = slice[index]
			index++
		}
		return present, element
	}
}

func addressIterator[A ip.Address[A]](first, last A) Iterator[A] {
	var current A = first
	var done bool = false

	return func() (bool, A) {
		var element A
		if done {
			return false, element
		}
		r := current
		done = compare.Eq(current, last)
		current = ip.Next(current)
		return true, r
	}
}

func ranges2AddressIterator[A ip.Address[A]](slice []AddressRange[A]) Iterator[A] {
	ranges := sliceIterator(slice)
	rok, rnge := ranges()
	addresses := rnge.Addresses()

	return func() (bool, A) {
		var result A
		for rok {
			aok, result := addresses()
			if aok {
				return true, result
			}
			rok, rnge = ranges()
			if !rok {
				break
			}
			addresses = rnge.Addresses()
		}
		return false, result
	}
}
