package ip

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ipfreely-uk/go/ip/compare"
)

var ipv4Masks []Addr4 = allMasks(V4())
var ipv6Masks []Addr6 = allMasks(V6())

// Subnet mask of given bit size.
// Panics if mask bits exceeds [Family_Width] or is less than zero.
func SubnetMask[A Address[A]](family Family[A], bits int) (mask A) {
	validateBits(family, bits)

	var r any
	if family.Version() == Version4 {
		r = ipv4Masks[bits]
	} else {
		r = ipv6Masks[bits]
	}
	return reflect.ValueOf(r).Interface().(A)
}

func validateBits[A Address[A]](family Family[A], bits int) {
	width := family.Width()
	if bits < 0 || bits > width {
		msg := fmt.Sprintf("wanted 0-%d for IPv%d; got %d", family.Width(), family.Version(), bits)
		panic(msg)
	}
}

// Number of addresses in subnet with given bit mask size.
// Panics if mask bits exceeds width of family or is less than zero.
func SubnetAddressCount[A Address[A]](family Family[A], bits int) (count *big.Int) {
	validateBits(family, bits)
	size := big.NewInt(int64(family.Width() - bits))
	two := big.NewInt(2)
	return two.Exp(two, size, nil)
}

// Mask size in bits.
// Returns between 0 and [Family] bit width inclusive if first and last form valid CIDR block.
// Returns -1 if first and last do not form valid CIDR block.
func SubnetMaskSize[A Address[A]](first, last A) (bits int) {
	fam := first.Family()
	xor := first.Xor(last)
	zero := fam.FromInt(0)
	if !compare.Eq(xor.And(first), zero) || !compare.Eq(xor.And(last), xor) {
		return -1
	}
	bitz := fam.Width() - xor.Not().TrailingZeros()
	mask := SubnetMask(fam, bitz)
	if compare.Eq(xor.And(mask), zero) {
		return bitz
	}
	return -1
}

func allMasks[A Address[A]](family Family[A]) []A {
	masks := []A{}
	for i := 0; i <= family.Width(); i++ {
		masks = append(masks, makeMask(family, i))
	}
	return masks
}

func makeMask[A Address[A]](family Family[A], bits int) A {
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
func SubnetMaskCovers[A Address[A]](maskBits int, address A) (maskBitsDoCover bool) {
	if maskBits < 0 {
		return false
	}
	fam := address.Family()
	if maskBits > fam.Width() {
		return false
	}
	zero := fam.FromInt(0)
	iMask := SubnetMask(fam, maskBits).Not()
	return iMask.And(address).Compare(zero) == 0
}
