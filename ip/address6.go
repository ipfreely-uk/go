package ip

import (
	"math/bits"
	"net/netip"
)

type Address6 struct {
	high uint64
	low  uint64
}

func (a Address6) sealed() {}

func (a Address6) Family() Family[Address6] {
	a.sealed()
	return V6()
}

func (a Address6) Bytes() []byte {
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

func (a Address6) Not() Address6 {
	return Address6{
		^a.high,
		^a.low,
	}
}

func (a Address6) Add(addend Address6) Address6 {
	var low = a.low + addend.low
	var high = a.high + addend.high
	if low < addend.low || low < a.low {
		high = high + 1
	}
	return Address6{
		high,
		low,
	}
}

func (a Address6) Subtract(subtrahend Address6) Address6 {
	var low = a.low - subtrahend.low
	var high = a.high - subtrahend.high
	if a.low < subtrahend.low {
		high = high - 1
	}
	return Address6{
		high,
		low,
	}
}

func (a Address6) Multiply(multiplicand Address6) Address6 {
	this := ToBigInt(a)
	that := ToBigInt(multiplicand)
	max := ToBigInt(MaxAddress(a.Family()))
	result := this.Mul(this, that).Mod(this, max)
	address, _ := FromBigInt(a.Family(), result)
	return address
}

func (a Address6) Divide(denominator Address6) Address6 {
	this := ToBigInt(a)
	that := ToBigInt(denominator)
	max := ToBigInt(MaxAddress(a.Family()))
	result, _ := this.DivMod(this, that, max)
	address, _ := FromBigInt(a.Family(), result)
	return address
}

func (a Address6) Mod(denominator Address6) Address6 {
	this := ToBigInt(a)
	that := ToBigInt(denominator)
	result := this.Mod(this, that)
	address, _ := FromBigInt(a.Family(), result)
	return address
}

func (a Address6) And(operand Address6) Address6 {
	return Address6{
		a.high & operand.high,
		a.low & operand.low,
	}
}

func (a Address6) Or(operand Address6) Address6 {
	return Address6{
		a.high | operand.high,
		a.low | operand.low,
	}
}

func (a Address6) Xor(operand Address6) Address6 {
	return Address6{
		a.high ^ operand.high,
		a.low ^ operand.low,
	}
}

func (a Address6) Shift(bits int) Address6 {
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
	return Address6{
		high,
		low,
	}
}

func (a Address6) Compare(other Address6) int {
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

func (a Address6) LeadingZeros() int {
	high0 := bits.LeadingZeros64(a.high)
	if high0 == 64 {
		return bits.LeadingZeros64(a.low) + 64
	}
	return high0
}

func (a Address6) TrailingZeros() int {
	low0 := bits.TrailingZeros64(a.low)
	if low0 == 64 {
		return bits.TrailingZeros64(a.high) + 64
	}
	return low0
}

func (a Address6) String() string {
	b := a.Bytes()
	addr, _ := netip.AddrFromSlice(b)
	return addr.String()
}
