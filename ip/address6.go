package ip

import (
	"math/bits"
	"net/netip"
)

// Immutable 128bit unsigned integer IP [Int] representation.
// Use [V6] to create values.
type Addr6 struct {
	high uint64
	low  uint64
}

var v6ZERO = Addr6{}
var v6ONE = Addr6{0, 1}

func (a Addr6) sealed() {}

// Returns [Version6]
func (a Addr6) Version() Version {
	return Version6
}

// Returns [Width6]
func (a Addr6) Width() int {
	return Width6
}

// Returns [V6]
func (a Addr6) Family() Family[Addr6] {
	return V6()
}

// Returns 16 byte slice
func (a Addr6) Bytes() []byte {
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

// See [Int]
func (a Addr6) Not() Addr6 {
	return Addr6{
		^a.high,
		^a.low,
	}
}

// See [Int]
func (a Addr6) Add(addend Addr6) Addr6 {
	var low = a.low + addend.low
	var high = a.high + addend.high
	if low < addend.low || low < a.low {
		high = high + 1
	}
	return Addr6{
		high,
		low,
	}
}

// See [Int]
func (a Addr6) Subtract(subtrahend Addr6) Addr6 {
	var low = a.low - subtrahend.low
	var high = a.high - subtrahend.high
	if a.low < subtrahend.low {
		high = high - 1
	}
	return Addr6{
		high,
		low,
	}
}

// See [Int]
func (a Addr6) Multiply(multiplicand Addr6) Addr6 {
	hi, lo := bits.Mul64(a.low, multiplicand.low)
	hi += (a.high * multiplicand.low) + (a.low * multiplicand.high)
	return Addr6{
		high: hi,
		low:  lo,
	}
}

// See [Int]
func (a Addr6) Divide(denominator Addr6) Addr6 {
	if denominator == v6ZERO {
		panic("divide by zero")
	}
	if denominator == v6ONE {
		return a
	}
	compared := a.Compare(denominator)
	if compared == 0 {
		return v6ONE
	}
	if compared < 0 {
		return v6ZERO
	}
	if a.high == 0 && denominator.high == 0 {
		return Addr6{
			low: a.low / denominator.low,
		}
	}
	this := ToBigInt(a)
	that := ToBigInt(denominator)
	result := this.Div(this, that)
	address, _ := FromBigInt(a.Family(), result)
	return address
}

// See [Int]
func (a Addr6) Mod(denominator Addr6) Addr6 {
	if denominator == v6ONE {
		return v6ZERO
	}
	if denominator != v6ZERO {
		comp := a.Compare(denominator)
		if comp == 0 {
			return v6ZERO
		}
		if comp < 0 {
			return a
		}
	}
	quotient := a.Divide(denominator)
	return a.Subtract(quotient.Multiply(denominator))
}

// See [Int]
func (a Addr6) And(operand Addr6) Addr6 {
	return Addr6{
		a.high & operand.high,
		a.low & operand.low,
	}
}

// See [Int]
func (a Addr6) Or(operand Addr6) Addr6 {
	return Addr6{
		a.high | operand.high,
		a.low | operand.low,
	}
}

// See [Int]
func (a Addr6) Xor(operand Addr6) Addr6 {
	return Addr6{
		a.high ^ operand.high,
		a.low ^ operand.low,
	}
}

// See [Int]
func (a Addr6) Shift(bits int) Addr6 {
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
	return Addr6{
		high,
		low,
	}
}

// See [Int]
func (a Addr6) Compare(other Addr6) int {
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

// See [Int]
func (a Addr6) LeadingZeros() int {
	high0 := bits.LeadingZeros64(a.high)
	if high0 == 64 {
		return bits.LeadingZeros64(a.low) + 64
	}
	return high0
}

// See [Int]
func (a Addr6) TrailingZeros() int {
	low0 := bits.TrailingZeros64(a.low)
	if low0 == 64 {
		return bits.TrailingZeros64(a.high) + 64
	}
	return low0
}

// See [Int]
func (a Addr6) String() string {
	b := a.Bytes()
	addr, _ := netip.AddrFromSlice(b)
	return addr.String()
}

// See [Int]
func (a Addr6) Float64() float64 {
	if a.high == 0 {
		return float64(a.low)
	}
	// TODO: something better
	bi := ToBigInt(a)
	f, _ := bi.Float64()
	return f
}
