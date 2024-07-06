package ipset

import "github.com/ipfreely-uk/go/ip"

// Tests if two discrete sets are equal
func Eq[A ip.Number[A]](s0, s1 Discrete[A]) (equal bool) {
	if s0 == s1 {
		return true
	}
	left := s0.Intervals()
	right := s1.Intervals()
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
