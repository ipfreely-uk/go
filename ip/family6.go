package ip

import "errors"

type family6 struct{}

var f6 = family6{}

func (f *family6) Version() Version {
	return Version6
}

func (f *family6) Width() int {
	return 128
}

func (f *family6) FromBytes(b ...byte) (Address6, error) {
	if len(b) != 16 {
		return Address6{}, errors.New("IPv6 addresses are 16 bytes")
	}
	var high = uint64(b[0])<<24 | uint64(b[1])<<16 | uint64(b[2])<<8 | uint64(b[3])
	high = high | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
	var low = uint64(b[8])<<24 | uint64(b[9])<<16 | uint64(b[10])<<8 | uint64(b[11])
	low = low | uint64(b[12])<<24 | uint64(b[13])<<16 | uint64(b[14])<<8 | uint64(b[15])
	return Address6{
		high,
		low,
	}, nil
}

func (f *family6) FromInt(i uint32) Address6 {
	return Address6{
		0,
		uint64(i),
	}
}

func V6() Family[Address6] {
	return &f6
}
