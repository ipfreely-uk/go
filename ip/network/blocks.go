package network

import (
	"math"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
	"github.com/ipfreely-uk/go/ip/subnet"
)

// TODO: can replace with constant
var LOG_2 = math.Log2(2.0)

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

	return func() (bool, Block[A]) {
		if done {
			return false, nil
		}
		max := maxMask(current)
		size := ip.Next(end.Subtract(start))
		x := log(size) / LOG_2
		width := start.Family().Width()
		maxDiff := int(width - int(math.Floor(x)))
		if max > maxDiff {
			max = maxDiff
		}
		block := NewBlock(current, max)
		last := block.Last()
		if compare.Eq(last, end) {
			done = true
		} else {
			current = ip.Next(last)
		}
		return true, block
	}
}

func log[A ip.Address[A]](address A) float64 {
	MAX_DIGITS_2 := 977

	bi := ip.ToBigInt(address)
	blex := bi.BitLen() - MAX_DIGITS_2
	if blex > 0 {
		a := address.Shift(blex)
		bi = ip.ToBigInt(a)
	}
	double, _ := bi.Float64()
	res := math.Log2(double)
	if res > 0 {
		res = res + float64(blex)*LOG_2
	}
	return res
}

func maxMask[A ip.Address[A]](address A) int {
	bytes := address.Bytes()
	bits := address.Family().Width()
	for i := len(bytes) - 1; i >= 0; i-- {
		b := bytes[i]
		if b == 0 {
			bits = bits - 8
		} else {
			n := 0
			var m byte = 1
			for j := 0; j <= 8; j++ {
				if b&m != 0 {
					n = j
					break
				}
				m = m << 1
			}
			bits = bits - n
			break
		}
	}
	return bits
}
