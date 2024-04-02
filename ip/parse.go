package ip

import (
	"errors"
	"fmt"
	"net"
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

// Parse IP address string from unknown family
func ParseUnknown(candidate string) (any, error) {
	parsed := net.ParseIP(candidate)
	if parsed == nil {
		msg := fmt.Sprintf("%s is not an IP address", candidate)
		return nil, errors.New(msg)
	}
	return FromBytes(parsed...)
}

// Parse IP address bytes from unknown family
func FromBytes(address ...byte) (any, error) {
	length := len(address)
	if length == 4 {
		return V4().FromBytes(address...)
	}
	if length == 16 {
		return V6().FromBytes(address...)
	}
	return nil, errors.New("slice must be 4 or 16 bytes")
}
