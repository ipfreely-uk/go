package ip

import (
	"errors"
	"math/big"
)

// Converts any address to big integer
func ToBigInt[A Address[A]](address A) *big.Int {
	return big.NewInt(0).SetBytes(address.Bytes())
}

// Converts big integer to address
func FromBigInt[A Address[A]](family Family[A], i *big.Int) (A, error) {
	b := i.Bytes()
	maxlen := family.Width() / 8
	if i.Sign() < 0 || len(b) > maxlen {
		return family.FromInt(0), errors.New("out of range")
	}
	return family.FromBytes(leftPad(b, maxlen)...)
}

func leftPad(b []byte, l int) []byte {
	if len(b) == l {
		return b
	}
	value := make([]byte, l)
	diff := l - len(b)
	copy(value[diff:], b)
	return value
}
