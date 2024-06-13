package network

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ipfreely-uk/go/ip"
)

// Parses CIDR notation
func ParseCIDRNotation[A ip.Address[A]](f ip.Family[A], notation string) (Block[A], error) {
	addressPart, mask, err := splitCidr(notation)
	if err != nil {
		return nil, err
	}
	address, err := ip.Parse(f, addressPart)
	if err != nil {
		return nil, err
	}
	if !ip.SubnetMaskCovers(mask, address) {
		msg := fmt.Sprintf("%s has invalid mask", notation)
		return nil, errors.New(msg)
	}
	return NewBlock(address, mask), nil
}

func splitCidr(notation string) (addr string, maskBits int, err error) {
	var address string
	var mask int
	split := strings.LastIndex(notation, "/")
	if split < 0 {
		msg := fmt.Sprintf("%s not CIDR notation", notation)
		return address, mask, errors.New(msg)
	}
	address = notation[:split]
	maskPart := notation[split+1:]
	mask, err = strconv.Atoi(maskPart)
	return address, mask, err
}

// Parses CIDR notation where IP address family is unknown.
// Returns error if operand is not valid CIDR notation.
func ParseUnknownCIDRNotation(notation string) (netAddress ip.Untyped, maskBits int, err error) {
	var addr ip.Untyped
	var mask int

	addressPart, mask, err := splitCidr(notation)
	if err != nil {
		return addr, mask, err
	}
	address, err := ip.ParseUnknown(addressPart)
	if err != nil {
		return addr, mask, err
	}
	cover := false
	switch a := address.(type) {
	case ip.Addr4:
		cover = ip.SubnetMaskCovers(mask, a)
	case ip.Addr6:
		cover = ip.SubnetMaskCovers(mask, a)
	}
	if !cover {
		msg := fmt.Sprintf("%s has invalid mask", notation)
		return address, mask, errors.New(msg)
	}
	return address, mask, err
}
