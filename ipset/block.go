package ipset

import (
	"fmt"
	"iter"
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

type block[A ip.Number[A]] struct {
	first A
	last  A
}

func (b *block[A]) MaskSize() int {
	return ip.SubnetMaskSize(b.first, b.last)
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

func (b *block[A]) Addresses() iter.Seq[A] {
	return addressSeq(b.first, b.last)
}

func (b *block[A]) Intervals() iter.Seq[Interval[A]] {
	var i Interval[A] = b
	return singleSeq(i)
}

func (b *block[A]) String() string {
	return fmt.Sprintf("%s/%d", b.first.String(), b.MaskSize())
}

func (b *block[A]) Mask() A {
	return ip.SubnetMask(b.first.Family(), b.MaskSize())
}

func (b *block[A]) CidrNotation() string {
	return b.String()
}

// Creates [Block] set.
//
// Panics if mask does not cover network address or is out of range for address family.
func NewBlock[A ip.Number[A]](network A, mask int) Block[A] {
	fam := network.Family()
	if fam.Width() == mask {
		return NewSingle(network)
	}
	m := ip.SubnetMask(fam, mask)
	if !ip.Eq(network, m.And(network)) {
		msg := fmt.Sprintf("mask %s does not cover %s", m.String(), network.String())
		panic(msg)
	}
	last := network.Or(m.Not())
	return &block[A]{network, last}
}
