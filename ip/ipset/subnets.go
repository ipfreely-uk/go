package ipset

import (
	"fmt"

	"github.com/ipfreely-uk/go/ip"
)

// Splits [Block] into subnets of a given size.
// Panics on illegal mask bits.
func Subnets[A ip.Number[A]](b Block[A], maskBits int) Iterator[Block[A]] {
	first := b.First()
	f := first.Family()

	if maskBits < b.MaskSize() || maskBits > f.Width() {
		msg := fmt.Sprintf("%s cannot be split with mask bits %d", b.String(), maskBits)
		panic(msg)
	}

	one := f.FromInt(1)
	current := first
	done := false

	return func() (element Block[A], exists bool) {
		var sub Block[A]
		if !done {
			sub = NewBlock(current, maskBits)
			done = ip.Eq(sub.Last(), b.Last())
			current = sub.Last().Add(one)
			return sub, true
		}
		return sub, false
	}
}
