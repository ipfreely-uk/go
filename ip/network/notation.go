package network

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
)

// Parses CIDR notation
func ParseCIDRNotation[A ip.Address[A]](f ip.Family[A], notation string) (Block[A], error) {
	split := strings.LastIndex(notation, "/")
	if split < 0 {
		msg := fmt.Sprintf("%s not CIDR notation", notation)
		return nil, errors.New(msg)
	}
	addressPart := notation[:split]
	address, err := ip.Parse(f, addressPart)
	if err != nil {
		return nil, err
	}
	maskPart := notation[split+1:]
	mask, err := strconv.Atoi(maskPart)
	if err != nil {
		return nil, err
	}
	// TODO: too many repeated checks; make Block API better
	if mask < 0 || mask > address.Family().Width() {
		msg := fmt.Sprintf("%s has invalid mask", notation)
		return nil, errors.New(msg)
	}
	m := ip.SubnetMask(address.Family(), mask)
	if !compare.Eq(address, m.And(address)) {
		msg := fmt.Sprintf("%s has invalid mask", notation)
		return nil, errors.New(msg)
	}
	return NewBlock(address, mask), nil
}

// Parses CIDR notation where IP address family is unknown.
// Returns error if operand is not valid CIDR notation.
// Returns [Block] when oprand is valid.
func ParseUnknownCIDRNotation(notation string) (any, error) {
	b, err := ParseCIDRNotation(ip.V4(), notation)
	if err == nil {
		return b, err
	}
	return ParseCIDRNotation(ip.V6(), notation)
}
