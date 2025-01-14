// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset

import (
	"iter"

	"github.com/ipfreely-uk/go/ip"
)

// Tests if two discrete sets are equal.
// Iterates each [Discrete] set's [Interval] sets comparing first and last elements.
func Eq[A ip.Int[A]](set0, set1 Discrete[A]) (equal bool) {
	if set0 == set1 {
		return true
	}
	left, lstop := iter.Pull(set0.Intervals())
	defer lstop()
	right, rstop := iter.Pull(set1.Intervals())
	defer rstop()
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
