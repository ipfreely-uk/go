package ip

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import "errors"

type family4 struct{}

var f4 = family4{}

func (f *family4) sealed() {}

func (f *family4) Version() Version {
	f.sealed()
	return Version4
}

func (f *family4) Width() int {
	return Width4
}

func (f *family4) FromBytes(b ...byte) (Addr4, error) {
	if len(b) != 4 {
		return Addr4{}, errors.New("IPv4 addresses are 4 bytes")
	}
	var addr = uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	return Addr4{
		addr,
	}, nil
}

func (f *family4) MustFromBytes(b ...byte) Addr4 {
	a, err := f.FromBytes(b...)
	if err != nil {
		panic(err)
	}
	return a
}

func (f *family4) FromInt(i uint32) Addr4 {
	return Addr4{
		i,
	}
}

func (f *family4) String() string {
	return "IPv4"
}

// IPv4 family of addresses
func V4() Family[Addr4] {
	return &f4
}
