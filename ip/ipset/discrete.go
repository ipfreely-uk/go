package ipset

import (
	"math/big"
	"sort"
	"strings"

	"github.com/ipfreely-uk/go/ip"
)

type discrete[A ip.Number[A]] struct {
	intervals []Interval[A]
}

func (s *discrete[A]) Contains(address A) bool {
	for _, r := range s.intervals {
		compare := address.Compare(r.First())
		if compare < 0 {
			break
		}
		if compare == 0 || address.Compare(r.Last()) <= 0 {
			return true
		}
	}
	return false
}

func (s *discrete[A]) Size() *big.Int {
	sum := big.NewInt(0)
	for _, r := range s.intervals {
		sum = sum.Add(sum, r.Size())
		println(sum.String())
	}
	return sum
}

func (s *discrete[A]) Addresses() Iterator[A] {
	return ranges2AddressIterator(s.intervals)
}

func (s *discrete[A]) Intervals() Iterator[Interval[A]] {
	return sliceIterator(s.intervals)
}

func (s *discrete[A]) String() string {
	buf := strings.Builder{}
	buf.WriteString("{")
	delim := ""
	for _, i := range s.intervals {
		buf.WriteString(delim)
		delim = ", "
		buf.WriteString(i.String())
	}
	buf.WriteString("}")
	return buf.String()
}

// Creates [Discrete] set as a union of addresses from the operand elements.
//
// If set is contiguous range returns result of [NewInterval] function.
// If set is CIDR range returns result of [NewBlock] function.
// If zero-length slice returns the empty set.
func NewDiscrete[A ip.Number[A]](sets ...Discrete[A]) (set Discrete[A]) {
	intervals := toIntervals(sets)
	intervals = rationalize(intervals)
	if len(intervals) == 1 {
		return intervals[0]
	}
	if len(intervals) == 0 {
		return emptySet[A]()
	}
	return &discrete[A]{
		intervals,
	}
}

func toIntervals[A ip.Number[A]](sets []Discrete[A]) []Interval[A] {
	result := []Interval[A]{}
	for _, set := range sets {
		next := set.Intervals()
		for s, ok := next(); ok; s, ok = next() {
			result = append(result, s)
		}
	}
	return result
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
