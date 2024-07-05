package ipset

import (
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

// Tests if IP address ranges have common elements
func Intersect[A ip.Number[A]](r0, r1 Interval[A]) bool {
	return r0.Contains(r1.First()) || r0.Contains(r1.Last()) || r1.Contains(r0.First()) || r1.Contains(r0.Last())
}

// Tests if IP address ranges are one element from overlap
func Adjacent[A ip.Number[A]](r0, r1 Interval[A]) bool {
	return lastNextToFirst(r0.Last(), r1.First()) || lastNextToFirst(r1.Last(), r0.First())
}

func lastNextToFirst[A ip.Number[A]](last, first A) bool {
	return compare.Eq(last, ip.Prev(first)) && isNotZero(first)
}

func isNotZero[A ip.Number[A]](address A) bool {
	zero := address.Family().FromInt(0)
	return zero.Compare(address) != 0
}

// Tests if IP address ranges either [Intersect] or are [Adjacent]
func Contiguous[A ip.Number[A]](r0, r1 Interval[A]) bool {
	return Intersect(r0, r1) || Adjacent(r0, r1)
}

// Joins IP address ranges using least and greatest elements from both.
// Ranges do not have to be contiguous.
func Join[A ip.Number[A]](r0, r1 Interval[A]) Interval[A] {
	first := compare.Min(r0.First(), r1.First())
	last := compare.Max(r0.Last(), r1.Last())
	if compare.Eq(r0.First(), first) && compare.Eq(r0.Last(), last) {
		return r0
	}
	if compare.Eq(r1.First(), first) && compare.Eq(r1.Last(), last) {
		return r1
	}
	return NewInterval(first, last)
}
