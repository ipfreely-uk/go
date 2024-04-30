package compare

// Generic comparison interface
type Comparable[C any] interface {
	// Returns -1 if operand greater; 0 if operand equal; 1 if operand less
	Compare(operand C) int
}

// Convenience function for testing equality
func Eq[C Comparable[C]](a0, a1 C) bool {
	return a0.Compare(a1) == 0
}

// Returns greatest element
func Max[C Comparable[C]](this, that C) C {
	if this.Compare(that) >= 0 {
		return this
	}
	return that
}

// Returns least element
func Min[C Comparable[C]](this, that C) C {
	if this.Compare(that) <= 0 {
		return this
	}
	return that
}
