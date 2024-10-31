package ipset

import (
	"iter"
	"math"

	"github.com/ipfreely-uk/go/ip"
)

// math.Log(2.0)
var log_2 = 0.6931471805599453

// Subdivides [Interval] into CIDR [Block] sets
func Blocks[A ip.Number[A]](set Interval[A]) iter.Seq[Block[A]] {
	first := set.First()
	last := set.Last()
	mask := ip.SubnetMaskSize(first, last)
	if mask >= 0 {
		block := NewBlock(first, mask)
		return singleSeq(block)
	}
	return blockIterator(set.First(), set.Last())
}

func blockIterator[A ip.Number[A]](start, end A) iter.Seq[Block[A]] {
	// implementation breaks on entire internet but guarded above
	return func(yield func(Block[A]) bool) {
		walkBlocks(start, end, yield)
	}
}

func walkBlocks[A ip.Number[A]](start, end A, yield func(Block[A]) bool) {
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
		last := block.Last()
		if !more || ip.Eq(last, end) {
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
