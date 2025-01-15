package ip

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import "iter"

// Zero address for family.
func MinAddress[A Int[A]](fam Family[A]) (firstInFamily A) {
	return fam.FromInt(0)
}

// Maximum address for family
func MaxAddress[A Int[A]](fam Family[A]) (lastInFamily A) {
	return MinAddress(fam).Not()
}

// Increments argument by one with overflow
func Next[A Int[A]](address A) (incremented A) {
	one := address.Family().FromInt(1)
	return address.Add(one)
}

// Decrements argument by one with underflow
func Prev[A Int[A]](address A) (decremented A) {
	one := address.Family().FromInt(1)
	return address.Subtract(one)
}

// Address iteration for range loops.
// Iteration uses [Next] or [Prev] depending on relative values of first and last.
// The returned sequence is reusable.
func Inclusive[A Int[A]](first, last A) (inclusiveRange iter.Seq[A]) {
	var step func(n A) A
	if first.Compare(last) < 0 {
		step = Next
	} else {
		step = Prev
	}
	current := first
	return func(yield func(A) bool) {
		for {
			more := yield(current)
			if !more {
				return
			}
			if Eq(current, last) {
				return
			}
			current = step(current)
		}
	}
}

// The exclusive version of [Inclusive].
// If start and end are equal the sequence is empty.
func Exclusive[A Int[A]](start, excludedEnd A) (exclusiveRange iter.Seq[A]) {
	comp := start.Compare(excludedEnd)
	if comp == 0 {
		return func(yield func(A) bool) {}
	}
	if comp > 0 {
		return Inclusive(start, Next(excludedEnd))
	}
	return Inclusive(start, Prev(excludedEnd))
}
