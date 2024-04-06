package network

import (
	"math/big"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
	"github.com/ipfreely-uk/go/ip/subnet"
)

type addressRange[A ip.Address[A]] struct {
	first A
	last  A
}

func (a *addressRange[A]) Contains(address A) bool {
	return a.first.Compare(address) <= 0 && a.last.Compare(address) >= 0
}

func (a *addressRange[A]) Size() *big.Int {
	first := ip.ToBigInt(a.first)
	last := ip.ToBigInt(a.last)
	one := big.NewInt(1)
	sub := last.Sub(last, first)
	return sub.Add(sub, one)
}

func (a *addressRange[A]) First() A {
	return a.first
}

func (a *addressRange[A]) Last() A {
	return a.last
}

func (a *addressRange[A]) Addresses() Iterator[A] {
	return addressIterator(a.first, a.last)
}

func (a *addressRange[A]) Ranges() Iterator[AddressRange[A]] {
	slice := []AddressRange[A]{a}
	return sliceIterator(slice)
}

// Creates new AddressRange.
// Return value conforms to Block if possible.
func NewRange[A ip.Address[A]](first, last A) AddressRange[A] {
	f := compare.Min(first, last)
	l := compare.Max(first, last)
	mask := subnet.MaskSize(f, l)
	if mask >= 0 {
		b := NewBlock(f, mask)
		return b
	}
	return &addressRange[A]{
		f,
		l,
	}
}
