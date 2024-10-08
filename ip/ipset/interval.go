package ipset

import (
	"math/big"
	"strings"

	"github.com/ipfreely-uk/go/ip"
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
	f := e.first.String()
	l := e.last.String()
	len := len(f) + len(l) + 1
	buf := strings.Builder{}
	buf.Grow(len)
	buf.WriteString(f)
	buf.WriteRune('-')
	buf.WriteString(l)
	return buf.String()
}

// Creates [Interval] set.
//
// If range is valid CIDR block returns value from [NewBlock] instead.
func NewInterval[A ip.Number[A]](first, last A) Interval[A] {
	f := least(first, last)
	l := greatest(first, last)
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
