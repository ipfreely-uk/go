package ipset

import (
	"fmt"
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

type single[A ip.Number[A]] struct {
	address A
}

func (b *single[A]) MaskSize() int {
	return b.address.Family().Width()
}

func (b *single[A]) Contains(address A) bool {
	return ip.Eq(b.address, address)
}

func (b *single[A]) Size() *big.Int {
	return big.NewInt(1)
}

func (b *single[A]) First() A {
	return b.address
}

func (b *single[A]) Last() A {
	return b.address
}

func (b *single[A]) Addresses() Iterator[A] {
	return addressIterator(b.address, b.address)
}

func (b *single[A]) Intervals() Iterator[Interval[A]] {
	slice := []Interval[A]{b}
	return sliceIterator(slice)
}

func (b *single[A]) String() string {
	return fmt.Sprintf("%s/%d", b.address.String(), b.MaskSize())
}

func (b *single[A]) Mask() A {
	return ip.SubnetMask(b.address.Family(), b.MaskSize())
}

func (b *single[A]) CidrNotation() string {
	return b.String()
}

// Creates [Block] set from a single address
func NewSingle[A ip.Number[A]](address A) Block[A] {
	return &single[A]{address}
}
