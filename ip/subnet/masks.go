package subnet

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

var ipv4Masks []ip.A4 = allMasks(ip.V4())
var ipv6Masks []ip.A6 = allMasks(ip.V6())

// Mask of given bit size.
// Panics if mask bits exceeds [ip.Family.Width] or is less than zero.
func Mask[A ip.Address[A]](family ip.Family[A], bits int) A {
	validateBits(family, bits)

	var r any
	if family.Version() == ip.Version4 {
		r = ipv4Masks[bits]
	} else {
		r = ipv6Masks[bits]
	}
	return reflect.ValueOf(r).Interface().(A)
}

func validateBits[A ip.Address[A]](family ip.Family[A], bits int) {
	width := family.Width()
	if bits < 0 || bits > width {
		msg := fmt.Sprintf("wanted 0-%d for IPv%d; got %d", family.Width(), family.Version(), bits)
		panic(msg)
	}
}

// Number of addresses in subnet with given bit mask size.
// Panics if mask bits exceeds width of family or is less than zero.
func AddressCount[A ip.Address[A]](family ip.Family[A], bits int) *big.Int {
	validateBits(family, bits)
	size := big.NewInt(int64(family.Width() - bits))
	two := big.NewInt(2)
	return two.Exp(two, size, nil)
}

// Mask size in bits.
// Returns between 0 and [ip.Family] bit width inclusive if first and last form valid CIDR block.
// Returns -1 if first and last do not form valid CIDR block.
func MaskSize[A ip.Address[A]](first, last A) int {
	fam := first.Family()
	xor := first.Xor(last)
	zero := fam.FromInt(0)
	if !compare.Eq(xor.And(first), zero) || !compare.Eq(xor.And(last), xor) {
		return -1
	}
	bits := fam.Width() - xor.Not().TrailingZeros()
	mask := Mask(fam, bits)
	if compare.Eq(xor.And(mask), zero) {
		return bits
	}
	return -1
}

func allMasks[A ip.Address[A]](family ip.Family[A]) []A {
	masks := []A{}
	for i := 0; i <= family.Width(); i++ {
		masks = append(masks, makeMask(family, i))
	}
	return masks
}

func makeMask[A ip.Address[A]](family ip.Family[A], bits int) A {
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
