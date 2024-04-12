package ip

import (
	"math/bits"
	"net/netip"
)

// Immutable 128bit unsigned integer IP [Address] representation
type A6 struct {
	high uint64
	low  uint64
}

func (a A6) sealed() {}

func (a A6) Family() Family[A6] {
	a.sealed()
	return V6()
}

func (a A6) Bytes() []byte {
	return []byte{
		byte(a.high >> 56),
		byte(a.high >> 48),
		byte(a.high >> 40),
		byte(a.high >> 32),
		byte(a.high >> 24),
		byte(a.high >> 16),
		byte(a.high >> 8),
		byte(a.high),
		byte(a.low >> 56),
		byte(a.low >> 48),
		byte(a.low >> 40),
		byte(a.low >> 32),
		byte(a.low >> 24),
		byte(a.low >> 16),
		byte(a.low >> 8),
		byte(a.low),
	}
}

func (a A6) Not() A6 {
	return A6{
		^a.high,
		^a.low,
	}
}

func (a A6) Add(addend A6) A6 {
	var low = a.low + addend.low
	var high = a.high + addend.high
	if low < addend.low || low < a.low {
		high = high + 1
	}
	return A6{
		high,
		low,
	}
}

func (a A6) Subtract(subtrahend A6) A6 {
	var low = a.low - subtrahend.low
	var high = a.high - subtrahend.high
	if a.low < subtrahend.low {
		high = high - 1
	}
	return A6{
		high,
		low,
	}
}

func (a A6) Multiply(multiplicand A6) A6 {
	this := ToBigInt(a)
	that := ToBigInt(multiplicand)
	max := ToBigInt(MaxAddress(a.Family()))
	result := this.Mul(this, that).Mod(this, max)
	address, _ := FromBigInt(a.Family(), result)
	return address
}

func (a A6) Divide(denominator A6) A6 {
	this := ToBigInt(a)
	that := ToBigInt(denominator)
	max := ToBigInt(MaxAddress(a.Family()))
	result, _ := this.DivMod(this, that, max)
	address, _ := FromBigInt(a.Family(), result)
	return address
}

func (a A6) Mod(denominator A6) A6 {
	this := ToBigInt(a)
	that := ToBigInt(denominator)
	result := this.Mod(this, that)
	address, _ := FromBigInt(a.Family(), result)
	return address
}

func (a A6) And(operand A6) A6 {
	return A6{
		a.high & operand.high,
		a.low & operand.low,
	}
}

func (a A6) Or(operand A6) A6 {
	return A6{
		a.high | operand.high,
		a.low | operand.low,
	}
}

func (a A6) Xor(operand A6) A6 {
	return A6{
		a.high ^ operand.high,
		a.low ^ operand.low,
	}
}

func (a A6) Shift(bits int) A6 {
	var high uint64
	var low uint64
	if bits > 0 {
		n := bits % 64
		x := a.high << (64 - n)
		high = a.high >> n
		low = a.low>>n | x
	} else {
		n := (bits * -1) % 64
		x := a.low >> (64 - n)
		high = a.high<<n | x
		low = a.low << n
	}
	return A6{
		high,
		low,
	}
}

func (a A6) Compare(other A6) int {
	if a.high < other.high {
		return -1
	}
	if a.high > other.high {
		return 1
	}
	if a.low < other.low {
		return -1
	}
	if a.low > other.low {
		return 1
	}
	return 0
}

func (a A6) LeadingZeros() int {
	high0 := bits.LeadingZeros64(a.high)
	if high0 == 64 {
		return bits.LeadingZeros64(a.low) + 64
	}
	return high0
}

func (a A6) TrailingZeros() int {
	low0 := bits.TrailingZeros64(a.low)
	if low0 == 64 {
		return bits.TrailingZeros64(a.high) + 64
	}
	return low0
}

func (a A6) String() string {
	b := a.Bytes()
	addr, _ := netip.AddrFromSlice(b)
	return addr.String()
}

func (a A6) Float64() float64 {
	if a.high == 0 {
		return float64(a.low)
	}
	// TODO: something better
	bi := ToBigInt(a)
	f, _ := bi.Float64()
	return f
}
