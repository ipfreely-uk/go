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

func (b *block[A]) Contains(address A) bool {
	return b.first.Compare(address) <= 0 && b.last.Compare(address) >= 0
}

func (b *block[A]) Size() *big.Int {
	diff := b.last.Subtract(b.first)
	bi := ip.ToBigInt(diff)
	one := big.NewInt(1)
	return bi.Add(bi, one)
}

func (b *block[A]) First() A {
	return b.first
}

func (b *block[A]) Last() A {
	return b.last
}

func (b *block[A]) Addresses() Iterator[A] {
	return addressIterator(b.first, b.last)
}

func (b *block[A]) Ranges() Iterator[AddressRange[A]] {
	slice := []AddressRange[A]{b}
	return sliceIterator(slice)
}

func (b *block[A]) String() string {
	return fmt.Sprintf("%s/%d", b.first.String(), b.MaskSize())
}

func (b *block[A]) Mask() A {
	return subnet.Mask(b.first.Family(), b.MaskSize())
}

func (b *block[A]) CidrNotation() string {
	return b.String()
}

// Creates [Block].
// Panics if mask does not cover network address or is out of range for address family.
func NewBlock[A ip.Address[A]](network A, mask int) Block[A] {
	fam := network.Family()
	if fam.Width() == mask {
		return &single[A]{network}
	}
	m := subnet.Mask(fam, mask)
	if !compare.Eq(network, m.And(network)) {
		msg := fmt.Sprintf("mask %s does not cover %s", m.String(), network.String())
		panic(msg)
	}
	last := network.Or(m.Not())
	return &block[A]{network, last}
}
