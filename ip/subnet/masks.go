package subnet

import (
	"fmt"
	"math/big"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

// Creates mask of given bit size
func Mask[A ip.Address[A]](family ip.Family[A], bits int) A {
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

func validateBits[A ip.Address[A]](family ip.Family[A], bits int) {
	width := family.Width()
	if bits < 0 || bits > width {
		msg := fmt.Sprintf("wanted 0-%d for IPv%d; got %d", family.Width(), family.Version(), bits)
		panic(msg)
	}
}

// Number of addresses in subnet with given bit mask size
func AddressCount[A ip.Address[A]](family ip.Family[A], bits int) *big.Int {
	validateBits(family, bits)
	size := big.NewInt(int64(family.Width() - bits))
	two := big.NewInt(2)
	return two.Exp(two, size, nil)
}

// Mask size in bits
func MaskSize[A ip.Address[A]](first, last A) int {
	fam := first.Family()
	width := fam.Width()
	xor := first.Xor(last)
	for i := width; i >= 0; i-- {
		mask := Mask(fam, i)
		imask := mask.Not()
		if compare.Eq(imask, xor) {
			return i
		}
	}
	return -1
}
