package ipset

import (
	"fmt"
	"math/big"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

type interval[A ip.Number[A]] struct {
	first A
	last  A
}

func (a *interval[A]) Contains(address A) bool {
	return a.first.Compare(address) <= 0 && a.last.Compare(address) >= 0
}

func (a *interval[A]) Size() *big.Int {
	first := ip.ToBigInt(a.first)
	last := ip.ToBigInt(a.last)
	one := big.NewInt(1)
	sub := last.Sub(last, first)
	return sub.Add(sub, one)
}

func (a *interval[A]) First() A {
	return a.first
}

func (a *interval[A]) Last() A {
	return a.last
}

func (a *interval[A]) Addresses() Iterator[A] {
	return addressIterator(a.first, a.last)
}

func (a *interval[A]) Intervals() Iterator[Interval[A]] {
	slice := []Interval[A]{a}
	return sliceIterator(slice)
}

func (e *interval[A]) String() string {
	return fmt.Sprintf("%s-%s", e.first.String(), e.last.String())
}

// Creates new [Interval] instance.
// Return value conforms to [Block] if possible.
func NewInterval[A ip.Number[A]](first, last A) Interval[A] {
	f := compare.Min(first, last)
	l := compare.Max(first, last)
	mask := ip.SubnetMaskSize(f, l)
	if mask >= 0 {
		b := NewBlock(f, mask)
		return b
	}
	return &interval[A]{
		f,
		l,
	}
}