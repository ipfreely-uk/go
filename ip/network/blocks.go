package network

import (
	"math"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
	"github.com/ipfreely-uk/go/ip/subnet"
)

// TODO: can replace with constant
var LOG_2 = math.Log(2.0)

// Subdivides range into valid CIDR blocks
func Blocks[A ip.Address[A]](r Range[A]) Iterator[Block[A]] {
	first := r.First()
	last := r.Last()
	mask := subnet.MaskSize(first, last)
	if mask >= 0 {
		block := NewBlock(first, mask)
		slice := []Block[A]{block}
		return sliceIterator(slice)
	}
	return blockIterator(r.First(), r.Last())
}

func blockIterator[A ip.Address[A]](start, end A) Iterator[Block[A]] {
	// implementation breaks on entire internet but guarded elsewhere
	current := start
	done := false
	width := start.Family().Width()

	return func() (bool, Block[A]) {
		if done {
			return false, nil
		}
		maxSize := width - current.TrailingZeros()
		size := ip.Next(end.Subtract(current))
		x := log(size) / LOG_2
		maxDiff := int(width - int(math.Floor(x)))
		mask := max(maxSize, maxDiff)
		block := NewBlock(current, mask)
		last := block.Last()
		if compare.Eq(last, end) {
			done = true
		} else {
			current = ip.Next(last)
		}
		return true, block
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func log[A ip.Address[A]](address A) float64 {
	MAX_DIGITS_2 := 977

	bitlen := address.Family().Width() - address.LeadingZeros()
	blex := bitlen - MAX_DIGITS_2
	a := address
	if blex > 0 {
		a = address.Shift(blex)
	}
	double := toFloat64(a)
	res := math.Log(double)
	if blex > 0 {
		res = res + float64(blex)*LOG_2
	}
	return res
}

func toFloat64[A ip.Address[A]](address A) float64 {
	bi := ip.ToBigInt(address)
	f, _ := bi.Float64()
	return f
}
