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
	// Addresses from least to greatest
	Addresses() Iterator[A]
	// Non-contiguous ranges from least to greatest
	Ranges() Iterator[AddressRange[A]]
}

// Immutable contiguous range of one or more IP addresses.
// TODO: rename this type to avoid confusion with range keyword.
type AddressRange[A ip.Address[A]] interface {
	AddressSet[A]
	First() A
	Last() A
}

// Immutable RFC-4632 CIDR block.
type Block[A ip.Address[A]] interface {
	AddressRange[A]
	MaskSize() int
}
