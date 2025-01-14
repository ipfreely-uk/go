// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset

import (
	"fmt"
	"iter"

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
