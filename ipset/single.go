package ipset

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"fmt"
	"iter"
	"math/big"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipmask"
)

type single[A ip.Int[A]] struct {
	address A
}

func (b *single[A]) MaskSize() int {
	return b.address.Family().Width()
}

func (b *single[A]) Contains(address A) bool {
	return ip.Eq(b.address, address)
}

func (b *single[A]) Size() *big.Int {
	return big.NewInt(1)
}

func (b *single[A]) Empty() bool {
	return false
}

func (b *single[A]) First() A {
	return b.address
}

func (b *single[A]) Last() A {
	return b.address
}

func (b *single[A]) Addresses() iter.Seq[A] {
	return singleSeq(b.address)
}

func (b *single[A]) Intervals() iter.Seq[Interval[A]] {
	var i Interval[A] = b
	return singleSeq(i)
}

func (b *single[A]) String() string {
	return fmt.Sprintf("%s/%d", b.address.String(), b.MaskSize())
}

func (b *single[A]) Mask() A {
	return ipmask.For(b.address.Family(), b.MaskSize())
}

func (b *single[A]) CidrNotation() string {
	return b.String()
}

// Creates [Block] set from a single address
func NewSingle[A ip.Int[A]](address A) Block[A] {
	return &single[A]{address}
}
