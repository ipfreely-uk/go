package ipset

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"iter"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/ipfreely-uk/go/ip"
)

type discrete[A ip.Int[A]] struct {
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

func (s *discrete[A]) Empty() bool {
	return false
}

func (s *discrete[A]) Addresses() iter.Seq[A] {
	return intervalToAddressSeq(s.intervals)
}

func (s *discrete[A]) Intervals() iter.Seq[Interval[A]] {
	return sliceSeq(s.intervals)
}

func (s *discrete[A]) String() string {
	LIMIT := 5
	buf := strings.Builder{}
	buf.WriteString("{")
	delim := ""
	for i, set := range s.intervals {
		buf.WriteString(delim)
		if i > LIMIT {
			l := strconv.Itoa(len(s.intervals) - LIMIT)
			buf.WriteRune('[')
			buf.WriteString(l)
			buf.WriteString("...]")
			break
		}
		delim = ", "
		buf.WriteString(set.String())
	}
	buf.WriteString("}")
	return buf.String()
}

// Creates [Discrete] set as a union of addresses from the operand elements.
//
// If set is contiguous range returns result of [NewInterval] function.
// If set is CIDR range returns result of [NewBlock] function.
// Zero-length slice returns the empty set.
func NewDiscrete[A ip.Int[A]](sets ...Discrete[A]) (set Discrete[A]) {
	intervals := rationalize(sets)
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

func rationalize[A ip.Int[A]](sets []Discrete[A]) []Interval[A] {
	set := map[Interval[A]]bool{}
	for _, s := range sets {
		for i := range s.Intervals() {
			merge(set, i)
		}
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

func merge[A ip.Int[A]](intervals map[Interval[A]]bool, i Interval[A]) {
	candidate := i
	for k := range intervals {
		if Contiguous(k, candidate) {
			candidate = Extremes(k, candidate)
			delete(intervals, k)
		}
	}
	intervals[candidate] = true
}
