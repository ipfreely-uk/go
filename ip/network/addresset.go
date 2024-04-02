package network

import (
	"math/big"
	"sort"

	"github.com/ipfreely-uk/go/ip"
)

type addressset[A ip.Address[A]] struct {
	ranges []Range[A]
}

func (s *addressset[A]) Contains(address A) bool {
	for _, r := range s.ranges {
		if r.Contains(address) {
			return true
		}
	}
	return false
}

func (s *addressset[A]) Size() *big.Int {
	sum := big.NewInt(0)
	for _, r := range s.ranges {
		sum = sum.Add(sum, r.Size())
	}
	return sum
}

func (s *addressset[A]) Addresses() Iterator[A] {
	return ranges2AddressIterator(s.ranges)
}

func (s *addressset[A]) Ranges() Iterator[Range[A]] {
	return sliceIterator(s.ranges)
}

func NewSet[A ip.Address[A]](spans ...Range[A]) AddressSet[A] {
	if len(spans) == 0 {
		return emptySet[A]()
	}
	spans = rationalize(spans)
	if len(spans) == 1 {
		return spans[0]
	}
	return &addressset[A]{
		spans,
	}
}

func rationalize[A ip.Address[A]](spans []Range[A]) []Range[A] {
	result := make([]Range[A], len(spans))
	count := len(result)
outer:
	for i := count - 1; i >= 0; i-- {
		a := result[i]
		for j := 0; j < i; j++ {
			b := result[j]
			if Contiguous(a, b) {
				c := Join(a, b)
				result[i] = c
				count--
				continue outer
			}
		}
	}
	result = result[:count]
	sort.Slice(result, func(i, j int) bool {
		fi := result[i].First()
		fj := result[j].First()
		return fi.Compare(fj) < 0
	})
	return result
}
