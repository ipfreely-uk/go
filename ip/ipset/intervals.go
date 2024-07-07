package ipset

import (
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

// Tests if IP address ranges have common elements
func Intersect[A ip.Number[A]](i0, i1 Interval[A]) bool {
	return i0.Contains(i1.First()) || i0.Contains(i1.Last()) || i1.Contains(i0.First()) || i1.Contains(i0.Last())
}

// Tests if IP address ranges are one element from overlap
func Adjacent[A ip.Number[A]](i0, i1 Interval[A]) bool {
	return lastNextToFirst(i0.Last(), i1.First()) || lastNextToFirst(i1.Last(), i0.First())
}

func lastNextToFirst[A ip.Number[A]](last, first A) bool {
	return compare.Eq(last, ip.Prev(first)) && isNotZero(first)
}

func isNotZero[A ip.Number[A]](address A) bool {
	zero := address.Family().FromInt(0)
	return zero.Compare(address) != 0
}

// Tests if IP address ranges either [Intersect] or are [Adjacent]
func Contiguous[A ip.Number[A]](i0, i1 Interval[A]) bool {
	return Intersect(i0, i1) || Adjacent(i0, i1)
}

// Joins IP address ranges using least and greatest elements from both.
// Intervals do not have to be [Contiguous].
func Join[A ip.Number[A]](i0, i1 Interval[A]) Interval[A] {
	first := compare.Min(i0.First(), i1.First())
	last := compare.Max(i0.Last(), i1.Last())
	if compare.Eq(i0.First(), first) && compare.Eq(i0.Last(), last) {
		return i0
	}
	if compare.Eq(i1.First(), first) && compare.Eq(i1.Last(), last) {
		return i1
	}
	return NewInterval(first, last)
}
