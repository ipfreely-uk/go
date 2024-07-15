package ipset

import "github.com/ipfreely-uk/go/ip"

// Tests if two discrete sets are equal.
// Iterates each [Discrete] set's [Interval] sets comparing first and last elements.
func Eq[A ip.Number[A]](set0, set1 Discrete[A]) (equal bool) {
	if set0 == set1 {
		return true
	}
	left := set0.Intervals()
	right := set1.Intervals()
	for {
		this, rok := left()
		that, lok := right()
		if !rok && !lok {
			return true
		}
		if rok != lok {
			return false
		}
		if this == that {
			continue
		}
		if this.First().Compare(that.First()) != 0 {
			return false
		}
		if this.Last().Compare(that.Last()) != 0 {
			return false
		}
	}
}
