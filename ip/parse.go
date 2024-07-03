package ip

import (
	"errors"
	"net/netip"
)

// Parses address string
func Parse[A Number[A]](family Family[A], candidate string) (address A, err error) {
	var a A
	parsed, err := netip.ParseAddr(candidate)
	if err != nil {
		return a, err
	}
	return family.FromBytes(parsed.AsSlice()...)
}

// As [Parse] but panics on error
func MustParse[A Number[A]](family Family[A], candidate string) (address A) {
	a, err := Parse(family, candidate)
	if err != nil {
		panic(err)
	}
	return a
}

// Parse IP address string from unknown family
func ParseUnknown(candidate string) (Address, error) {
	parsed, err := netip.ParseAddr(candidate)
	if err != nil {
		return nil, err
	}
	return FromBytes(parsed.AsSlice()...)
}

// As [ParseUnknown] but panics on error
func MustParseUnknown(candidate string) Address {
	a, err := ParseUnknown(candidate)
	if err != nil {
		panic(err)
	}
	return a
}

// Parse IP address bytes from unknown family
func FromBytes(address ...byte) (Address, error) {
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
func MustFromBytes(address ...byte) Address {
	a, err := FromBytes(address...)
	if err != nil {
		panic(err)
	}
	return a
}
