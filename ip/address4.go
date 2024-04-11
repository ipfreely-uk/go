package ip

import (
	"math/bits"
	"net/netip"
)

// Immutable 32bit unsigned integer IP [Address] representation
type A4 struct {
	value uint32
}

func (a A4) sealed() {}

// Returns [V4]
func (a A4) Family() Family[A4] {
	a.sealed()
	return V4()
}

// Returns 4 byte slice
func (a A4) Bytes() []byte {
	return []byte{
		byte(a.value >> 24),
		byte(a.value >> 16),
		byte(a.value >> 8),
		byte(a.value),
	}
}

// Bitwise NOT
func (a A4) Not() A4 {
	return A4{
		^a.value,
	}
}

// Addition with overflow
func (a A4) Add(addend A4) A4 {
	return A4{
		a.value + addend.value,
	}
}

// Subtraction with underflow
func (a A4) Subtract(addend A4) A4 {
	return A4{
		a.value - addend.value,
	}
}

// Multiplication with overflow
func (a A4) Multiply(multiplicand A4) A4 {
	return A4{
		a.value * multiplicand.value,
	}
}

// Division
func (a A4) Divide(denominator A4) A4 {
	return A4{
		a.value / denominator.value,
	}
}

// Modulus
func (a A4) Mod(denominator A4) A4 {
	return A4{
		a.value % denominator.value,
	}
}

// Bitwise AND
func (a A4) And(operand A4) A4 {
	return A4{
		a.value & operand.value,
	}
}

// Bitwise OR
func (a A4) Or(operand A4) A4 {
	return A4{
		a.value | operand.value,
	}
}

// Bitwise XOR
func (a A4) Xor(operand A4) A4 {
	return A4{
		a.value ^ operand.value,
	}
}

func (a A4) Shift(bits int) A4 {
	bits = bits % a.Family().Width()
	var v uint32
	if bits > 0 {
		v = a.value >> bits
	} else {
		v = a.value << (-1 * bits)
	}
	return A4{
		v,
	}
}

func (a A4) Compare(other A4) int {
	if a.value < other.value {
		return -1
	}
	if a.value > other.value {
		return 1
	}
	return 0
}

func (a A4) LeadingZeros() int {
	return bits.LeadingZeros32(a.value)
}

func (a A4) TrailingZeros() int {
	return bits.TrailingZeros32(a.value)
}

func (a A4) String() string {
	b := a.Bytes()
	addr, _ := netip.AddrFromSlice(b)
	return addr.String()
}
