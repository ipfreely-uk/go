package cidr

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/ipfreely-uk/go/ip/subnet"
)

func Notation[A ip.Address[A]](b network.Block[A]) string {
	return fmt.Sprintf("%s/%d", b.First(), b.MaskSize())
}

func Parse[A ip.Address[A]](f ip.Family[A], notation string) (network.Block[A], error) {
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
	m := subnet.Mask(address.Family(), mask)
	if !compare.Eq(address, m.And(address)) {
		msg := fmt.Sprintf("%s has invalid mask", notation)
		return nil, errors.New(msg)
	}
	return network.NewBlock(address, mask), nil
}
