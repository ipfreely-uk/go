package ipset

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"fmt"
	"iter"
	"math"

	"github.com/ipfreely-uk/go/ip"
)

// Splits [Block] into subnets of a given size.
//
// Panics on illegal mask bits.
// maskBits must be greater or equal to [Block].MaskSize and less than or equal to the bit width of the address family.
func Subnets[A ip.Int[A]](b Block[A], maskBits int) iter.Seq[Block[A]] {
	first := b.First()
	f := first.Family()

	if maskBits < b.MaskSize() || maskBits > f.Width() {
		msg := fmt.Sprintf("%s cannot be split with mask bits %d", b.String(), maskBits)
		panic(msg)
	}

	return func(yield func(Block[A]) bool) {
		one := f.FromInt(1)
		current := first
		for {
			sub := NewBlock(current, maskBits)
			more := yield(sub)
			if !more || ip.Eq(sub.Last(), b.Last()) {
				return
			}
			current = sub.Last().Add(one)
		}
	}
}

// math.Log(2.0)
var log_2 = 0.6931471805599453

// Subdivides [Interval] into CIDR [Block] sets
func Blocks[A ip.Int[A]](set Interval[A]) iter.Seq[Block[A]] {
	first := set.First()
	last := set.Last()
	mask := ip.SubnetMaskSize(first, last)
	if mask >= 0 {
		block := NewBlock(first, mask)
		return singleSeq(block)
	}
	return blockSequence(set.First(), set.Last())
}

func blockSequence[A ip.Int[A]](start, end A) iter.Seq[Block[A]] {
	// implementation breaks on entire internet but guarded above
	return func(yield func(Block[A]) bool) {
		walkBlocks(start, end, yield)
	}
}

func walkBlocks[A ip.Int[A]](start, end A, yield func(Block[A]) bool) {
	current := start
	width := start.Family().Width()
	for {
		maxSize := width - current.TrailingZeros()
		size := ip.Next(end.Subtract(current))
		l := math.Log(size.Float64())
		x := l / log_2
		maxDiff := width - int(math.Floor(x))
		mask := max(maxSize, maxDiff)
		block := NewBlock(current, mask)
		more := yield(block)
		if !more {
			return
		}
		last := block.Last()
		if ip.Eq(last, end) {
			return
		}
		current = ip.Next(last)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
