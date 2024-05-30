package network

import (
	"fmt"
	"math/big"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

type single[A ip.Address[A]] struct {
	address A
}

func (b *single[A]) MaskSize() int {
	return b.address.Family().Width()
}

func (b *single[A]) Contains(address A) bool {
	return compare.Eq(b.address, address)
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

func (b *single[A]) Ranges() Iterator[AddressRange[A]] {
	slice := []AddressRange[A]{b}
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
