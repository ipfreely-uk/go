package ip

// Internet protocol version
type Version uint8

const (
	// Internet protocol version 4
	Version4 Version = 4
	// Internet protocol version 6
	Version6 Version = 6
)

// The parts of [Number] without generic typing.
// [Addr4] and [Addr6] are the only types that can conform to this interface.
type Address interface {
	// Structs that conform to this interface must be produced by this package
	sealed()
	// IP address version
	Version() Version
	// Address as bytes
	Bytes() []byte
	// Canonical string form
	String() string
	// Equivalent to math/bits.LeadingZeros
	LeadingZeros() int
	// Equivalent to math/bits.TrailingZeros
	TrailingZeros() int
	// Approximation to float64
	Float64() (approximation float64)
}

// IP address as generic numeric type.
// [Addr4] and [Addr6] are the only types that can conform to this interface.
type Number[A any] interface {
	Address
	// IP address family - [V4] or [V6]
	Family() Family[A]
	// Addition with overflow
	Add(A) A
	// Subtraction with overflow
	Subtract(A) A
	// Multiplication with overflow
	Multiply(A) A
	// Division
	Divide(A) A
	// Modulus
	Mod(A) A
	// Bitwise NOT
	Not() A
	// Bitwise AND
	And(A) A
	// Bitwise OR
	Or(A) A
	// Bitwise XOR
	Xor(A) A
	// Bit shift. Use negative int for left shift; use positive int for right shift.
	Shift(int) A
	// Returns 1 if operand is less than this.
	// Returns -1 if operand is more than this.
	// Returns 0 if operand is equal.
	Compare(A) int
}
