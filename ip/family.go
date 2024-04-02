package ip

// IP address family.
// Obtain implementations via V4() or V6().
type Family[A any] interface {
	// IP address version
	Version() Version
	// Address width in bits - 32 or 128
	Width() int
	// Create address from bytes
	FromBytes(...byte) (A, error)
	// Create address from unsigned integer
	FromInt(i uint32) A
}
