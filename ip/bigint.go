// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip

import (
	"errors"
	"math/big"
)

// Converts any [Int] to big integer
func ToBigInt[A Int[A]](address A) (i *big.Int) {
	return big.NewInt(0).SetBytes(address.Bytes())
}

// Converts big integer to [Int].
// Returns error if value out of range for address family.
func FromBigInt[A Int[A]](family Family[A], i *big.Int) (address A, err error) {
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
