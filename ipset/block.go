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

type block[A ip.Int[A]] struct {
	first A
	mask  byte
}

func (b *block[A]) MaskSize() int {
	return int(b.mask)
}

func (b *block[A]) Contains(address A) bool {
	return b.first.Compare(address) <= 0 && b.Last().Compare(address) >= 0
}

var one = big.NewInt(1)

func (b *block[A]) Size() *big.Int {
	diff := b.Last().Subtract(b.first)
	bi := ip.ToBigInt(diff)
	return bi.Add(bi, one)
}

func (b *block[A]) Empty() bool {
	return false
}

func (b *block[A]) First() A {
	return b.first
}

func (b *block[A]) Last() A {
	f := b.first.Family()
	maskComplement := ipmask.SubnetMask(f, int(b.mask)).Not()
	return b.first.Or(maskComplement)
}

func (b *block[A]) Addresses() iter.Seq[A] {
	return ip.Inclusive(b.first, b.Last())
}

func (b *block[A]) Intervals() iter.Seq[Interval[A]] {
	var i Interval[A] = b
	return singleSeq(i)
}

func (b *block[A]) String() string {
	return fmt.Sprintf("%s/%d", b.first.String(), b.MaskSize())
}

func (b *block[A]) Mask() A {
	return ipmask.SubnetMask(b.first.Family(), b.MaskSize())
}

func (b *block[A]) CidrNotation() string {
	return b.String()
}

// Creates [Block] set.
//
// Panics if mask does not cover network address or is out of range for address family.
func NewBlock[A ip.Int[A]](network A, mask int) Block[A] {
	fam := network.Family()
	if fam.Width() == mask {
		return NewSingle(network)
	}
	m := ipmask.SubnetMask(fam, mask)
	if !ip.Eq(network, m.And(network)) {
		msg := fmt.Sprintf("mask %s does not cover %s", m.String(), network.String())
		panic(msg)
	}
	return &block[A]{network, byte(mask)}
}
