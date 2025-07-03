package ipmask

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ipfreely-uk/go/ip"
)

var ipv4Masks []ip.Addr4 = allMasks(ip.V4())
var ipv6Masks []ip.Addr6 = allMasks(ip.V6())

// Subnet mask of given bit size as an address.
//
// For IPv4 `0` returns `0.0.0.0` and `32` returns `255.255.255.255`.
// For IPv6 `0` returns `::` and `128` returns `ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`.
// Panics if mask bits exceeds [Family].Width or is less than zero.
func For[A ip.Int[A]](f ip.Family[A], maskBits int) (mask A) {
	validateBits(f, maskBits)

	var r ip.Address
	if f.Version() == ip.Version4 {
		r = ipv4Masks[maskBits]
	} else {
		r = ipv6Masks[maskBits]
	}
	return reflect.ValueOf(r).Interface().(A)
}

func validateBits[A ip.Int[A]](family ip.Family[A], bits int) {
	width := family.Width()
	if bits < 0 || bits > width {
		msg := fmt.Sprintf("wanted 0-%d for IPv%d width; got %d", family.Width(), family.Version(), bits)
		panic(msg)
	}
}

// Number of addresses in subnet with given bit mask size.
// Panics if mask bits exceeds width of family or is less than zero.
func Size[A ip.Int[A]](family ip.Family[A], maskBits int) (count *big.Int) {
	validateBits(family, maskBits)
	size := big.NewInt(int64(family.Width() - maskBits))
	two := big.NewInt(2)
	return two.Exp(two, size, nil)
}

// Mask size in bits for CIDR block.
//
// Returns between 0 and the family width inclusive if first and last form valid CIDR block.
// Returns -1 if first and last do not form valid CIDR block.
func Test[A ip.Int[A]](first, last A) (maskBits int) {
	fam := first.Family()
	xor := first.Xor(last)
	zero := fam.FromInt(0)
	if !ip.Eq(xor.And(first), zero) || !ip.Eq(xor.And(last), xor) {
		return -1
	}
	bits := fam.Width() - xor.Not().TrailingZeros()
	mask := For(fam, bits)
	if ip.Eq(xor.And(mask), zero) {
		return bits
	}
	return -1
}

func allMasks[A ip.Int[A]](family ip.Family[A]) []A {
	masks := []A{}
	for i := 0; i <= family.Width(); i++ {
		masks = append(masks, makeMask(family, i))
	}
	return masks
}

func makeMask[A ip.Int[A]](family ip.Family[A], bits int) A {
	validateBits(family, bits)

	bytes := family.Width() / 8
	arr := make([]byte, bytes)
	fullyMasked := bits / 8
	for i := 0; i < fullyMasked; i++ {
		arr[i] = 0b11111111
	}
	if fullyMasked != bytes {
		mod := bits % 8
		var end byte = 0
		switch mod {
		case 1:
			end = 0b10000000
		case 2:
			end = 0b11000000
		case 3:
			end = 0b11100000
		case 4:
			end = 0b11110000
		case 5:
			end = 0b11111000
		case 6:
			end = 0b11111100
		case 7:
			end = 0b11111110
		}
		arr[fullyMasked] = end
	}
	mask, _ := family.FromBytes(arr...)
	return mask
}

// Tests mask bits cover network address.
// Panics on invalid mask bits.
func Covers[A ip.Int[A]](maskBits int, address A) (maskBitsDoCover bool) {
	fam := address.Family()
	if maskBits < 0 || maskBits > fam.Width() {
		msg := fmt.Sprintf("%d is not a valid mask size for %v", maskBits, address)
		panic(msg)
	}
	zero := fam.FromInt(0)
	iMask := For(fam, maskBits).Not()
	return iMask.And(address).Compare(zero) == 0
}

// Tests mask bits are within width & [Covers] network address.
func IsValid(maskBits int, address ip.Address) (isValid bool) {
	var valid = false
	switch a := address.(type) {
	case ip.Addr4:
		valid = checkCIDRMask(a, maskBits)
	case ip.Addr6:
		valid = checkCIDRMask(a, maskBits)
	}
	return valid
}

func checkCIDRMask[A ip.Int[A]](addr A, mask int) bool {
	return mask >= 0 && mask <= addr.Family().Width() && Covers(mask, addr)
}

// Tests if the address is a valid subnet mask.
//
// Returns count of most significant bits set to true.
// For example, IPv4 mask `255.255.255.0` returns 24.
// If `address` is not a mask returns -1.
func Bits[A ip.Int[A]](address A) (maskBits int) {
	f := address.Family()
	bits := f.Width() - address.TrailingZeros()
	mask := For(f, bits)
	match := ip.Eq(mask, address)
	if match {
		return bits
	}
	return -1
}
