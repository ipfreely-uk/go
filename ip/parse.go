package ip

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"errors"
	"strconv"
	"strings"
)

var errInvalid = errors.New("invalid address")
var errAmbiquous = errors.New("leading zeroes not supported")

// Parses address string
func Parse[A Int[A]](family Family[A], candidate string) (address A, err error) {
	var a A
	if family.Version() == Version4 {
		v, err := parse4(candidate)
		if err != nil {
			return a, err
		}
		return family.FromInt(v), nil
	}
	b, err := parse6(candidate)
	if err != nil {
		return a, err

	}
	return family.FromBytes(b...)
}

// As [Parse] but panics on error
func MustParse[A Int[A]](family Family[A], candidate string) (address A) {
	a, err := Parse(family, candidate)
	if err != nil {
		panic(err)
	}
	return a
}

func parse4(candidate string) (uint32, error) {
	last := '-'
	dots := 0
	digits := 0
	var value uint32 = 0
	var quad uint32 = 0
	invalid := false

	for i, c := range candidate {
		if c == '.' {
			dots++
			if last == '.' {
				invalid = true
				break
			}
			value <<= 8
			value |= quad
			quad = 0
			digits = 0
		} else if digits > 0 && quad == 0 && c != '0' {
			return value, errAmbiquous
		} else if c >= '0' && c <= '9' {
			digits++
			quad = quad*10 + uint32(c-'0')
		} else {
			invalid = true
			break
		}
		if digits > 3 || quad > 255 || (i == 0 && c == '.') {
			invalid = true
			break
		}
		last = c
	}
	value <<= 8
	value |= quad
	if dots != 3 || last == '.' {
		invalid = true
	}
	if invalid {
		return value, errInvalid
	}
	return value, nil
}

func parse6(candidate string) ([]byte, error) {
	bytes := make([]byte, 16)
	invalid := false

	if strings.Contains(candidate, ":::") {
		invalid = true
	}
	head, tail, err := headTail6(candidate)
	if err != nil {
		invalid = true
	} else {
		for i, segment := range head {
			n, err := strconv.ParseUint(segment, 16, 16)
			if err != nil {
				invalid = true
				break
			}
			bytes[i*2] = byte(n >> 8)
			bytes[i*2+1] = byte(n)
		}
		offset := len(bytes) - (len(tail) * 2)
		for i, segment := range tail {
			n, err := strconv.ParseUint(segment, 16, 16)
			if err != nil {
				invalid = true
				break
			}
			bytes[offset+(i*2)] = byte(n >> 8)
			bytes[offset+(i*2)+1] = byte(n)
		}
	}

	if invalid {
		return bytes, errInvalid
	}
	return bytes, nil
}

func headTail6(candidate string) ([]string, []string, error) {
	before, after, ok := strings.Cut(candidate, "::")
	if !ok {
		c := split6(candidate)
		if len(c) != 8 {
			return nil, nil, errInvalid
		}
		return c, []string{}, nil
	}
	head, tail := before, after
	h, t := split6(head), split6(tail)
	if len(h)+len(t) > 7 {
		return nil, nil, errInvalid
	}
	return h, t, nil
}

func split6(candidate string) []string {
	if len(candidate) == 0 {
		return []string{}
	}
	return strings.Split(candidate, ":")
}

// Parse IP address string from unknown family
func ParseUnknown(candidate string) (Address, error) {
	value, err := parse4(candidate)
	if err == nil {
		return Addr4{value}, err
	}
	bytes, err := parse6(candidate)
	if err == nil {
		return MustFromBytes(bytes...), nil
	}
	return nil, err
}

// As [ParseUnknown] but panics on error
func MustParseUnknown(candidate string) Address {
	a, err := ParseUnknown(candidate)
	if err != nil {
		panic(err)
	}
	return a
}

// Parse IP address bytes from unknown family.
// Slice length must be 4 (IPv4) or 16 (IPv6).
func FromBytes(address ...byte) (Address, error) {
	length := len(address)
	if length == 4 {
		return V4().FromBytes(address...)
	}
	if length == 16 {
		return V6().FromBytes(address...)
	}
	return nil, errInvalid
}

// As [FromBytes] but panics on error
func MustFromBytes(address ...byte) Address {
	a, err := FromBytes(address...)
	if err != nil {
		panic(err)
	}
	return a
}
