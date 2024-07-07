package ipset

import (
	"math"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

// TODO: can replace with constant
var log_2 = math.Log(2.0)

// Subdivides [Interval] into CIDR [Block] sets
func Blocks[A ip.Number[A]](set Interval[A]) Iterator[Block[A]] {
	first := set.First()
	last := set.Last()
	mask := ip.SubnetMaskSize(first, last)
	if mask >= 0 {
		block := NewBlock(first, mask)
		slice := []Block[A]{block}
		return sliceIterator(slice)
	}
	return blockIterator(set.First(), set.Last())
}

func blockIterator[A ip.Number[A]](start, end A) Iterator[Block[A]] {
	// implementation breaks on entire internet but guarded above
	current := start
	done := false
	width := start.Family().Width()

	return func() (Block[A], bool) {
		if done {
			return nil, false
		}
		maxSize := width - current.TrailingZeros()
		size := ip.Next(end.Subtract(current))
		l := math.Log(size.Float64())
		x := l / log_2
		maxDiff := int(width - int(math.Floor(x)))
		mask := max(maxSize, maxDiff)
		block := NewBlock(current, mask)
		last := block.Last()
		if compare.Eq(last, end) {
			done = true
		} else {
			current = ip.Next(last)
		}
		return block, true
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
