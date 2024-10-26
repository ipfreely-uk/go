package ip

// IP address family.
// Obtain via [V4] or [V6] functions.
type Family[A any] interface {
	// Structs that conform to this interface must be produced by this package
	sealed()
	// IP address version
	Version() Version
	// Address width in bits - 32 or 128
	Width() int
	// Create address from bytes.
	// Returns error if slice is not [Width]/8 bytes.
	FromBytes(...byte) (address A, err error)
	// As FromBytes but panics on error
	MustFromBytes(...byte) (address A)
	// Create address from unsigned integer.
	// All values are valid.
	// For V4 operand 1 returns "0.0.0.1".
	// For V6 operand 1 return "::1".
	FromInt(i uint32) (address A)
	// Informational
	String() string
}
