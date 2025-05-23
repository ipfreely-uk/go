package ipset

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"iter"
	"math/big"
	"strings"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipmask"
)

type interval[A ip.Int[A]] struct {
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

func (a *interval[A]) Empty() bool {
	return false
}

func (a *interval[A]) First() A {
	return a.first
}

func (a *interval[A]) Last() A {
	return a.last
}

func (a *interval[A]) Addresses() iter.Seq[A] {
	return ip.Inclusive(a.first, a.last)
}

func (a *interval[A]) Intervals() iter.Seq[Interval[A]] {
	var i Interval[A] = a
	return singleSeq(i)
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
func NewInterval[A ip.Int[A]](first, last A) Interval[A] {
	least, greatest := order(first, last)
	mask := ipmask.Test(least, greatest)
	if mask >= 0 {
		b := NewBlock(least, mask)
		return b
	}
	return &interval[A]{
		least,
		greatest,
	}
}

func order[C ip.Int[C]](this, that C) (least, greatest C) {
	if this.Compare(that) <= 0 {
		return this, that
	}
	return that, this
}
