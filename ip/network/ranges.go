package network

import (
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

// Test if [Range]s have common elements
func Intersect[A ip.Address[A]](r0, r1 AddressRange[A]) bool {
	return r0.Contains(r1.First()) || r0.Contains(r1.Last()) || r1.Contains(r0.First()) || r1.Contains(r0.Last())
}

// Tests if [Range]s are one element from overlap
func Adjacent[A ip.Address[A]](r0, r1 AddressRange[A]) bool {
	return lastNextToFirst(r0.Last(), r1.First()) || lastNextToFirst(r1.Last(), r0.First())
}

func lastNextToFirst[A ip.Address[A]](last, first A) bool {
	zero := last.Family().FromInt(0)
	return !compare.Eq(zero, first) && compare.Eq(last, ip.Prev(first))
}

// Tests if [Range]s either [Intersect] or are [Adjacent]
func Contiguous[A ip.Address[A]](r0, r1 AddressRange[A]) bool {
	return Intersect(r0, r1) || Adjacent(r0, r1)
}

// Joins ranges using least and greatest elements from both.
// Ranges do not have to be contiguous.
func Join[A ip.Address[A]](r0, r1 AddressRange[A]) AddressRange[A] {
	first := compare.Min(r0.First(), r1.First())
	last := compare.Max(r0.Last(), r1.Last())
	if compare.Eq(r0.First(), first) && compare.Eq(r0.Last(), last) {
		return r0
	}
	if compare.Eq(r1.First(), first) && compare.Eq(r1.Last(), last) {
		return r1
	}
	return NewRange(first, last)
}
