package ipset

import (
	"fmt"
	"iter"

	"github.com/ipfreely-uk/go/ip"
)

// Splits [Block] into subnets of a given size.
// Panics on illegal mask bits.
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
