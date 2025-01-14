// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset

import (
	"github.com/ipfreely-uk/go/ip"
)

// Tests if IP address ranges have common elements
func Intersect[A ip.Int[A]](i0, i1 Interval[A]) bool {
	return i0.Contains(i1.First()) || i0.Contains(i1.Last()) || i1.Contains(i0.First()) || i1.Contains(i0.Last())
}

// Tests if IP address ranges are one element from overlap
func Adjacent[A ip.Int[A]](i0, i1 Interval[A]) bool {
	return lastNextToFirst(i0.Last(), i1.First()) || lastNextToFirst(i1.Last(), i0.First())
}

func lastNextToFirst[A ip.Int[A]](last, first A) bool {
	return ip.Eq(last, ip.Prev(first)) && isNotZero(first)
}

func isNotZero[A ip.Int[A]](address A) bool {
	zero := address.Family().FromInt(0)
	return zero.Compare(address) != 0
}

// Tests if IP address ranges either [Intersect] or are [Adjacent]
func Contiguous[A ip.Int[A]](i0, i1 Interval[A]) bool {
	return Intersect(i0, i1) || Adjacent(i0, i1)
}

// Creates [Interval] using least and greatest values from each
func Extremes[A ip.Int[A]](i0, i1 Interval[A]) Interval[A] {
	a := i0.First()
	b := i0.Last()
	x := i1.First()
	y := i1.Last()

	first, _ := order(a, x)
	_, last := order(b, y)
	if ip.Eq(a, first) && ip.Eq(b, last) {
		return i0
	}
	if ip.Eq(x, first) && ip.Eq(y, last) {
		return i1
	}
	return NewInterval(first, last)
}
