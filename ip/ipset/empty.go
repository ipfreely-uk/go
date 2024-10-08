package ipset

import (
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

type empty[A ip.Number[A]] struct{}

func (e *empty[A]) Contains(A) bool {
	return false
}

func (e *empty[A]) Size() *big.Int {
	return big.NewInt(0)
}

func (e *empty[A]) Addresses() Iterator[A] {
	return emptyIterator[A]()
}

func (e *empty[A]) Intervals() Iterator[Interval[A]] {
	return emptyIterator[Interval[A]]()
}

func (e *empty[A]) String() string {
	return "{}"
}

func emptySet[A ip.Number[A]]() Discrete[A] {
	return &empty[A]{}
}
