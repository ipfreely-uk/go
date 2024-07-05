package ipset

import (
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

// Discrete set of IP addresses
type Discrete[A ip.Number[A]] interface {
	// Tests if address in set
	Contains(address A) bool
	// Number of unique addresses.
	// The cardinality of the set.
	Size() *big.Int
	// Unique addresses from least to greatest
	Addresses() Iterator[A]
	// Contents as distinct [Interval] types
	Intervals() Iterator[Interval[A]]
	// Informational only
	String() string
}

// Immutable set of IP addresses between first and last inclusive.
//
// A range of one or more IP addresses.
// The name interval was chosen because range is a keyword in Go
// and it is a term in mathematical set theory.
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
