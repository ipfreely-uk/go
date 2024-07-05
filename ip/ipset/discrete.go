package ipset

import (
	"math/big"
	"sort"
	"strings"

	"github.com/ipfreely-uk/go/ip"
)

type discrete[A ip.Number[A]] struct {
	ranges []Interval[A]
}

func (s *discrete[A]) Contains(address A) bool {
	for _, r := range s.ranges {
		if r.Contains(address) {
			return true
		}
	}
	return false
}

func (s *discrete[A]) Size() *big.Int {
	sum := big.NewInt(0)
	for _, r := range s.ranges {
		sum = sum.Add(sum, r.Size())
		println(sum.String())
	}
	return sum
}

func (s *discrete[A]) Addresses() Iterator[A] {
	return ranges2AddressIterator(s.ranges)
}

func (s *discrete[A]) Intervals() Iterator[Interval[A]] {
	return sliceIterator(s.ranges)
}

func (s *discrete[A]) String() string {
	buf := strings.Builder{}
	buf.WriteString("{")
	delim := ""
	next := s.Intervals()
	for r, exists := next(); exists; r, exists = next() {
		buf.WriteString(delim)
		delim = ", "
		buf.WriteString(r.String())
	}
	buf.WriteString("}")
	return buf.String()
}

// Creates [Discrete] from given IP address ranges.
// Ranges may overlap.
// If set reduces to contiguous range returns type that conforms to [Interval].
func NewDiscrete[A ip.Number[A]](ranges ...Interval[A]) Discrete[A] {
	if len(ranges) == 0 {
		return emptySet[A]()
	}
	ranges = rationalize(ranges)
	if len(ranges) == 1 {
		return ranges[0]
	}
	return &discrete[A]{
		ranges,
	}
}

func rationalize[A ip.Number[A]](spans []Interval[A]) []Interval[A] {
	set := map[Interval[A]]bool{}
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
	result := []Interval[A]{}
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
