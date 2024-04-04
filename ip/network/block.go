package network

import (
	"fmt"
	"math/big"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
	"github.com/ipfreely-uk/go/ip/subnet"
)

type block[A ip.Address[A]] struct {
	first A
	last  A
}

func (b *block[A]) MaskSize() int {
	return subnet.MaskSize(b.first, b.last)
}

func (a *block[A]) Contains(address A) bool {
	return a.first.Compare(address) <= 0 && a.last.Compare(address) >= 0
}

func (a *block[A]) Size() *big.Int {
	first := ip.ToBigInt(a.first)
	last := ip.ToBigInt(a.last)
	one := big.NewInt(1)
	sub := last.Sub(last, first)
	return sub.Add(sub, one)
}

func (a *block[A]) First() A {
	return a.first
}

func (a *block[A]) Last() A {
	return a.last
}

func (a *block[A]) Addresses() Iterator[A] {
	return addressIterator(a.first, a.last)
}

func (a *block[A]) Ranges() Iterator[Range[A]] {
	slice := []Range[A]{a}
	return sliceIterator(slice)
}

func (b *block[A]) Blocks() Iterator[Block[A]] {
	slice := []Block[A]{b}
	return sliceIterator(slice)
}

func NewBlock[A ip.Address[A]](network A, mask int) Block[A] {
	fam := network.Family()
	m := subnet.Mask(fam, mask)
	if !compare.Eq(network, m.And(network)) {
		msg := fmt.Sprintf("mask %s does not cover %s", m.String(), network.String())
		panic(msg)
	}
	last := network.Or(m.Not())
	return &block[A]{network, last}
}
