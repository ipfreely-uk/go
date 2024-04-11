package ip

// IP address family.
// Obtain via [V4] or [V6].
type Family[A any] interface {
	// Structs that conform to this interface must be produced by this package
	sealed()
	// IP address version
	Version() Version
	// Address width in bits - 32 or 128
	Width() int
	// Create address from bytes.
	// Returns error if slice is not [Width]/8 bytes.
	FromBytes(...byte) (A, error)
	// As FromBytes but panics on error
	MustFromBytes(...byte) A
	// Create address from unsigned integer.
	// All values are valid.
	FromInt(i uint32) A
}
