package ip

import "errors"

type family4 struct{}

var f4 = family4{}

func (f *family4) Version() Version {
	return Version4
}

func (f *family4) Width() int {
	return 32
}

func (f *family4) FromBytes(b ...byte) (Address4, error) {
	if len(b) != 4 {
		return Address4{}, errors.New("IPv4 addresses are 4 bytes")
	}
	var addr = uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	return Address4{
		addr,
	}, nil
}

func (f *family4) FromInt(i uint32) Address4 {
	return Address4{
		i,
	}
}

func V4() Family[Address4] {
	return &f4
}
