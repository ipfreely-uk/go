package ipset

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ipfreely-uk/go/ip"
)

// Parses RFC 4632 CIDR notation.
// See also the [NewBlock] function.
//
// For the IPv4 family and expression "192.168.0.0/128",
// returns 192.168.0.0 as the network address part and
// 128 as the mask bits.
//
// Returns error if second argument is invalid CIDR notation.
func ParseCIDRNotation[A ip.Number[A]](f ip.Family[A], notation string) (netAddress A, maskBits int, err error) {
	var address A
	var mask int

	addressPart, mask, err := splitCidr(notation)
	if err != nil {
		return address, mask, err
	}
	address, err = ip.Parse(f, addressPart)
	if err != nil {
		return address, mask, err
	}
	if !ip.SubnetMaskCovers(mask, address) {
		msg := fmt.Sprintf("%s has invalid mask", notation)
		return address, mask, errors.New(msg)
	}
	return address, mask, nil
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
// Returns error if argument is invalid CIDR notation.
func ParseUnknownCIDRNotation(notation string) (netAddress ip.Address, maskBits int, err error) {
	var addr ip.Address
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
