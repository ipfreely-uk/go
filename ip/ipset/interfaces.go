package ipset

import (
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

// Discrete ordered set of IP addresses
type Discrete[A ip.Number[A]] interface {
	// Tests if address in set
	Contains(address A) bool
	// Number of unique addresses.
	// The cardinality of the set.
	Size() *big.Int
	// Unique addresses from least to greatest
	Addresses() Iterator[A]
	// Contents as distinct [Interval] sets.
	// Intervals do not [Intersect] and are not [Adjacent].
	// Intervals are returned from least address to greatest.
	Intervals() Iterator[Interval[A]]
	// Informational only
	String() string
}

// Immutable set of IP addresses between first and last inclusive.
//
// A range of one or more IP addresses.
// The name interval was chosen because range is a keyword in Go.
// Interval is a term from mathematical set theory.
type Interval[A ip.Number[A]] interface {
	Discrete[A]
	// Least address
	First() (address A)
	// Greatest address
	Last() (address A)
}

// Immutable RFC-4632 CIDR block.
// Roughly equivalent to the [netip.Prefix] type.
type Block[A ip.Number[A]] interface {
	Interval[A]
	// Mask size in bits
	MaskSize() (bits int)
	// Mask as IP address
	Mask() (address A)
	// The block in CIDR notation.
	CidrNotation() string
}
