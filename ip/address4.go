package ip

import (
	"math/bits"
	"net/netip"
)

// Immutable 32bit unsigned integer IP [Int] representation.
// Use [V4] to create values.
type Addr4 struct {
	value uint32
}

func (a Addr4) sealed() {}

// Returns [Version4]
func (a Addr4) Version() Version {
	return Version4
}

// Returns [Width4]
func (a Addr4) Width() int {
	return Width4
}

// Returns [V4]
func (a Addr4) Family() Family[Addr4] {
	return V4()
}

// Returns 4 byte slice
func (a Addr4) Bytes() []byte {
	return []byte{
		byte(a.value >> 24),
		byte(a.value >> 16),
		byte(a.value >> 8),
		byte(a.value),
	}
}

// See [Int]
func (a Addr4) Not() Addr4 {
	return Addr4{
		^a.value,
	}
}

// See [Int]
func (a Addr4) Add(addend Addr4) Addr4 {
	return Addr4{
		a.value + addend.value,
	}
}

// See [Int]
func (a Addr4) Subtract(subtrahend Addr4) Addr4 {
	return Addr4{
		a.value - subtrahend.value,
	}
}

// See [Int]
func (a Addr4) Multiply(multiplicand Addr4) Addr4 {
	return Addr4{
		a.value * multiplicand.value,
	}
}

// See [Int]
func (a Addr4) Divide(denominator Addr4) Addr4 {
	return Addr4{
		a.value / denominator.value,
	}
}

// See [Int]
func (a Addr4) Mod(denominator Addr4) Addr4 {
	return Addr4{
		a.value % denominator.value,
	}
}

// See [Int]
func (a Addr4) And(operand Addr4) Addr4 {
	return Addr4{
		a.value & operand.value,
	}
}

// See [Int]
func (a Addr4) Or(operand Addr4) Addr4 {
	return Addr4{
		a.value | operand.value,
	}
}

// See [Int]
func (a Addr4) Xor(operand Addr4) Addr4 {
	return Addr4{
		a.value ^ operand.value,
	}
}

// See [Int]
func (a Addr4) Shift(bits int) Addr4 {
	bits = bits % a.Family().Width()
	var v uint32
	if bits > 0 {
		v = a.value >> bits
	} else {
		v = a.value << (-1 * bits)
	}
	return Addr4{
		v,
	}
}

// See [Int]
func (a Addr4) Compare(other Addr4) int {
	if a.value < other.value {
		return -1
	}
	if a.value > other.value {
		return 1
	}
	return 0
}

// See [Int]
func (a Addr4) LeadingZeros() int {
	return bits.LeadingZeros32(a.value)
}

// See [Int]
func (a Addr4) TrailingZeros() int {
	return bits.TrailingZeros32(a.value)
}

// See [Int]
func (a Addr4) String() string {
	b := a.Bytes()
	addr, _ := netip.AddrFromSlice(b)
	return addr.String()
}

// See [Int]
func (a Addr4) Float64() float64 {
	return float64(a.value)
}
