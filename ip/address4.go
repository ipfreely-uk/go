package ip

import (
	"math/bits"
	"net/netip"
)

type Address4 struct {
	value uint32
}

func (a Address4) sealed() {}

func (a Address4) Family() Family[Address4] {
	a.sealed()
	return V4()
}

func (a Address4) Bytes() []byte {
	return []byte{
		byte(a.value >> 24),
		byte(a.value >> 16),
		byte(a.value >> 8),
		byte(a.value),
	}
}

func (a Address4) Not() Address4 {
	return Address4{
		^a.value,
	}
}

func (a Address4) Add(addend Address4) Address4 {
	return Address4{
		a.value + addend.value,
	}
}

func (a Address4) Subtract(addend Address4) Address4 {
	return Address4{
		a.value - addend.value,
	}
}

func (a Address4) Multiply(multiplicand Address4) Address4 {
	return Address4{
		a.value * multiplicand.value,
	}
}

func (a Address4) Divide(denominator Address4) Address4 {
	return Address4{
		a.value / denominator.value,
	}
}

func (a Address4) Mod(denominator Address4) Address4 {
	return Address4{
		a.value % denominator.value,
	}
}

func (a Address4) And(operand Address4) Address4 {
	return Address4{
		a.value & operand.value,
	}
}

func (a Address4) Or(operand Address4) Address4 {
	return Address4{
		a.value | operand.value,
	}
}

func (a Address4) Xor(operand Address4) Address4 {
	return Address4{
		a.value ^ operand.value,
	}
}

func (a Address4) Shift(bits int) Address4 {
	bits = bits % a.Family().Width()
	var v uint32
	if bits > 0 {
		v = a.value >> bits
	} else {
		v = a.value << (-1 * bits)
	}
	return Address4{
		v,
	}
}

func (a Address4) Compare(other Address4) int {
	if a.value < other.value {
		return -1
	}
	if a.value > other.value {
		return 1
	}
	return 0
}

func (a Address4) LeadingZeros() int {
	return bits.LeadingZeros32(a.value)
}

func (a Address4) TrailingZeros() int {
	return bits.TrailingZeros32(a.value)
}

func (a Address4) String() string {
	b := a.Bytes()
	addr, _ := netip.AddrFromSlice(b)
	return addr.String()
}
