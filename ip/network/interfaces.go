package network

import (
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

// IP address set.
type AddressSet[A ip.Address[A]] interface {
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
type AddressRange[A ip.Address[A]] interface {
	AddressSet[A]
	// Least address
	First() A
	// Greatest address
	Last() A
}

// Immutable RFC-4632 CIDR block.
type Block[A ip.Address[A]] interface {
	AddressRange[A]
	// Mask size in bits
	MaskSize() int
	// Mask as [ip.Address]
	Mask() A
}
