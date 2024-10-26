package ip

import "iter"

// Zero address for family.
func MinAddress[A Number[A]](fam Family[A]) (firstInFamily A) {
	return fam.FromInt(0)
}

// Maximum address for family
func MaxAddress[A Number[A]](fam Family[A]) (lastInFamily A) {
	return MinAddress(fam).Not()
}

// Increments argument by one with overflow
func Next[A Number[A]](address A) (incremented A) {
	one := address.Family().FromInt(1)
	return address.Add(one)
}

// Decrements argument by one with underflow
func Prev[A Number[A]](address A) (decremented A) {
	one := address.Family().FromInt(1)
	return address.Subtract(one)
}

// Address iteration for range loops.
// Iteration uses [Next] or [Prev] depending on relative values of first and last.
func Inclusive[A Number[A]](first, last A) iter.Seq[A] {
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
// If first and last are equal there are no results.
func Exclusive[A Number[A]](first, last A) iter.Seq[A] {
	comp := first.Compare(last)
	if comp == 0 {
		return func(yield func(A) bool) {}
	}
	if comp > 0 {
		return Inclusive(first, Next(last))
	}
	return Inclusive(first, Prev(last))
}
