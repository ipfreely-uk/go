package network

import (
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

// Immutable IP address set.
type AddressSet[A ip.Address[A]] interface {
	// Tests if address in set
	Contains(address A) bool
	// Number of addresses
	Size() *big.Int
	// Consituent addresses
	Addresses() Iterator[A]
	// Consituent ranges
	Ranges() Iterator[Range[A]]
}

// Immutable contiguous range of one or more IP addresses.
// TODO: rename this type to avoid confusion with range keyword.
type Range[A ip.Address[A]] interface {
	AddressSet[A]
	First() A
	Last() A
	Blocks() Iterator[Block[A]]
}

// Immutable RFC-4632 CIDR block.
type Block[A ip.Address[A]] interface {
	Range[A]
	MaskSize() int
}
