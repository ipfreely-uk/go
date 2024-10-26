package ipset

import (
	"iter"

	"github.com/ipfreely-uk/go/ip"
)

func emptySeq[E any]() iter.Seq[E] {
	return func(yield func(E) bool) {}
}

// TODO: single Seq

func sliceSeq[E any](slice []E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, e := range slice {
			more := yield(e)
			if !more {
				return
			}
		}
	}
}

func addressSeq[A ip.Number[A]](first, last A) iter.Seq[A] {
	return func(yield func(A) bool) {
		current := first
		for {
			more := yield(current)
			if !more || ip.Eq(current, last) {
				return
			}
			current = ip.Next(current)
		}
	}
}

func ranges2AddressSeq[A ip.Number[A]](slice []Interval[A]) iter.Seq[A] {
	return func(yield func(A) bool) {
		for _, i := range slice {
			for a := range i.Addresses() {
				more := yield(a)
				if !more {
					return
				}
			}
		}
	}
}