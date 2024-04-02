package network

import (
	"fmt"
	"math/big"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/subnet"
)

type ipRange[A ip.Address[A]] struct {
	first A
	last  A
}

func (a *ipRange[A]) Contains(address A) bool {
	return a.first.Compare(address) <= 0 && a.last.Compare(address) >= 0
}

func (a *ipRange[A]) Size() *big.Int {
	first := ip.ToBigInt(a.first)
	last := ip.ToBigInt(a.last)
	one := big.NewInt(1)
	sub := last.Sub(last, first)
	return sub.Add(sub, one)
}

func (a *ipRange[A]) First() A {
	return a.first
}

func (a *ipRange[A]) Last() A {
	return a.last
}

func (a *ipRange[A]) Addresses() Iterator[A] {
	return addressIterator(a.first, a.last)
}

func (a *ipRange[A]) Ranges() Iterator[Range[A]] {
	slice := []Range[A]{a}
	return sliceIterator(slice)
}

func (a *ipRange[A]) Blocks() Iterator[Block[A]] {
	return blockIterator(a.first, a.last)
}

func NewRange[A ip.Address[A]](first, last A) Range[A] {
	if first.Compare(last) > 0 {
		msg := fmt.Sprintf("first element %s must be less than last %s", first.String(), last.String())
		panic(msg)
	}
	mask := subnet.MaskSize(first, last)
	if mask >= 0 {
		return NewBlock(first, mask)
	}
	return &ipRange[A]{
		first,
		last,
	}
}
