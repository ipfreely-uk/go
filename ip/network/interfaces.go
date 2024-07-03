package network

import (
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

// IP address set.
type AddressSet[A ip.Number[A]] interface {
	// Tests if address in set
	Contains(address A) bool
	// Number of unique addresses
	Size() *big.Int
	// Unique addresses from least to greatest
	Addresses() Iterator[A]
	// Non-contiguous ranges from least to greatest
	Ranges() Iterator[AddressRange[A]]
	// Informational only
	String() string
}

// Immutable contiguous range of one or more IP addresses.
type AddressRange[A ip.Number[A]] interface {
	AddressSet[A]
	// Least address
	First() (address A)
	// Greatest address
	Last() (address A)
}

// Immutable RFC-4632 CIDR block.
// Roughly equivalent to the [netip.Prefix] type.
type Block[A ip.Number[A]] interface {
	AddressRange[A]
	// Mask size in bits
	MaskSize() (bits int)
	// Mask as IP address
	Mask() (address A)
	// The block in CIDR notation.
	CidrNotation() string
}
