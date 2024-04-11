package ip

import "errors"

type family4 struct{}

var f4 = family4{}

func (f *family4) sealed() {}

func (f *family4) Version() Version {
	f.sealed()
	return Version4
}

func (f *family4) Width() int {
	return 32
}

func (f *family4) FromBytes(b ...byte) (A4, error) {
	if len(b) != 4 {
		return A4{}, errors.New("IPv4 addresses are 4 bytes")
	}
	var addr = uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	return A4{
		addr,
	}, nil
}

func (f *family4) MustFromBytes(b ...byte) A4 {
	a, err := f.FromBytes(b...)
	if err != nil {
		panic(err)
	}
	return a
}

func (f *family4) FromInt(i uint32) A4 {
	return A4{
		i,
	}
}

// IPv4 family of addresses
func V4() Family[A4] {
	return &f4
}
