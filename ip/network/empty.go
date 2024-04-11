package network

import (
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

type empty[A ip.Address[A]] struct{}

func (e *empty[A]) Contains(A) bool {
	return false
}

func (e *empty[A]) Size() *big.Int {
	return big.NewInt(0)
}

func (e *empty[A]) Addresses() Iterator[A] {
	return emptyIterator[A]()
}

func (e *empty[A]) Ranges() Iterator[AddressRange[A]] {
	return emptyIterator[AddressRange[A]]()
}

func (e *empty[A]) String() string {
	return "{}"
}

func emptySet[A ip.Address[A]]() AddressSet[A] {
	return &empty[A]{}
}
