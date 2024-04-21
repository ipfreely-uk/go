package ip

import (
	"math/bits"
	"net/netip"
)

// Immutable 32bit unsigned integer IP [Address] representation
type Addr4 struct {
	value uint32
}

func (a Addr4) sealed() {}

// Returns [V4]
func (a Addr4) Family() Family[Addr4] {
	a.sealed()
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

// Bitwise NOT
func (a Addr4) Not() Addr4 {
	return Addr4{
		^a.value,
	}
}

// Addition with overflow
func (a Addr4) Add(addend Addr4) Addr4 {
	return Addr4{
		a.value + addend.value,
	}
}

// Subtraction with underflow
func (a Addr4) Subtract(addend Addr4) Addr4 {
	return Addr4{
		a.value - addend.value,
	}
}

// Multiplication with overflow
func (a Addr4) Multiply(multiplicand Addr4) Addr4 {
	return Addr4{
		a.value * multiplicand.value,
	}
}

// Division
func (a Addr4) Divide(denominator Addr4) Addr4 {
	return Addr4{
		a.value / denominator.value,
	}
}

// Modulus
func (a Addr4) Mod(denominator Addr4) Addr4 {
	return Addr4{
		a.value % denominator.value,
	}
}

// Bitwise AND
func (a Addr4) And(operand Addr4) Addr4 {
	return Addr4{
		a.value & operand.value,
	}
}

// Bitwise OR
func (a Addr4) Or(operand Addr4) Addr4 {
	return Addr4{
		a.value | operand.value,
	}
}

// Bitwise XOR
func (a Addr4) Xor(operand Addr4) Addr4 {
	return Addr4{
		a.value ^ operand.value,
	}
}

// Bitwise shift
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

func (a Addr4) Compare(other Addr4) int {
	if a.value < other.value {
		return -1
	}
	if a.value > other.value {
		return 1
	}
	return 0
}

func (a Addr4) LeadingZeros() int {
	return bits.LeadingZeros32(a.value)
}

func (a Addr4) TrailingZeros() int {
	return bits.TrailingZeros32(a.value)
}

func (a Addr4) String() string {
	b := a.Bytes()
	addr, _ := netip.AddrFromSlice(b)
	return addr.String()
}

func (a Addr4) Float64() float64 {
	return float64(a.value)
}
