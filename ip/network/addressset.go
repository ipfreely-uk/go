package network

import (
	"math/big"
	"sort"

	"github.com/ipfreely-uk/go/ip"
)

type addressset[A ip.Address[A]] struct {
	ranges []AddressRange[A]
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

func (s *addressset[A]) Ranges() Iterator[AddressRange[A]] {
	return sliceIterator(s.ranges)
}

func NewSet[A ip.Address[A]](spans ...AddressRange[A]) AddressSet[A] {
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

func rationalize[A ip.Address[A]](spans []AddressRange[A]) []AddressRange[A] {
	set := map[AddressRange[A]]bool{}
	for _, r := range spans {
		set[r] = true
	}
	for i := range spans {
		a := spans[i]
		for b := range set {
			if Contiguous(a, b) {
				a = Join(a, b)
				delete(set, b)
			}
		}
		set[a] = true
	}
	result := []AddressRange[A]{}
	for r := range set {
		result = append(result, r)
	}
	sort.Slice(result, func(i, j int) bool {
		fi := result[i].First()
		fj := result[j].First()
		return fi.Compare(fj) < 0
	})
	return result
}
