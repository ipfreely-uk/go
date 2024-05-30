package ip

import (
	"errors"
	"net/netip"
)

// Parses address string
func Parse[A Address[A]](family Family[A], candidate string) (A, error) {
	parsed, err := netip.ParseAddr(candidate)
	if err != nil {
		return family.FromInt(0), err
	}
	return family.FromBytes(parsed.AsSlice()...)
}

// As [Parse] but panics on error
func MustParse[A Address[A]](family Family[A], candidate string) A {
	a, err := Parse(family, candidate)
	if err != nil {
		panic(err)
	}
	return a
}

// Parse IP address string from unknown family
func ParseUnknown(candidate string) (Unknown, error) {
	parsed, err := netip.ParseAddr(candidate)
	if err != nil {
		return nil, err
	}
	return FromBytes(parsed.AsSlice()...)
}

// As [ParseUnknown] but panics on error
func MustParseUnknown(candidate string) Unknown {
	a, err := ParseUnknown(candidate)
	if err != nil {
		panic(err)
	}
	return a
}

// Parse IP address bytes from unknown family
func FromBytes(address ...byte) (Unknown, error) {
	length := len(address)
	if length == 4 {
		return V4().FromBytes(address...)
	}
	if length == 16 {
		return V6().FromBytes(address...)
	}
	return nil, errors.New("slice must be 4 or 16 bytes")
}

// As [FromBytes] but panics on error
func MustFromBytes(address ...byte) Unknown {
	a, err := FromBytes(address...)
	if err != nil {
		panic(err)
	}
	return a
}
