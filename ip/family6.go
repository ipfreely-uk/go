package ip

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import "errors"

type family6 struct{}

var f6 = family6{}

func (f *family6) sealed() {}

func (f *family6) Version() Version {
	f.sealed()
	return Version6
}

func (f *family6) Width() int {
	return Width6
}

func (f *family6) FromBytes(b ...byte) (Addr6, error) {
	if len(b) != 16 {
		return Addr6{}, errors.New("IPv6 addresses are 16 bytes")
	}
	var high = uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32
	high |= uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
	var low = uint64(b[8])<<56 | uint64(b[9])<<48 | uint64(b[10])<<40 | uint64(b[11])<<32
	low |= uint64(b[12])<<24 | uint64(b[13])<<16 | uint64(b[14])<<8 | uint64(b[15])
	return Addr6{
		high,
		low,
	}, nil
}

func (f *family6) MustFromBytes(b ...byte) Addr6 {
	a, err := f.FromBytes(b...)
	if err != nil {
		panic(err)
	}
	return a
}

func (f *family6) FromInt(i uint32) Addr6 {
	return Addr6{
		0,
		uint64(i),
	}
}

func (f *family6) String() string {
	return "IPv6"
}

// IPv6 family of addresses
func V6() Family[Addr6] {
	return &f6
}
