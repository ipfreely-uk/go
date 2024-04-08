package ip

// Internet protocol version
type Version uint8

const (
	// Internet protocol version 4
	Version4 Version = 4
	// Internet protocol version 6
	Version6 Version = 6
)

// Generic IP address type.
type Address[A any] interface {
	// Structs that conform to this interface must be produced by this package
	sealed()
	// IP address family
	Family() Family[A]
	// Address as bytes
	Bytes() []byte
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
	// Bit shift. Use negative int for left shift; use positive in for right shift.
	Shift(int) A
	// Returns 1 if operand is less than this.
	// Returns -1 if operand is more than this.
	// Returns 0 if operand is equal.
	Compare(A) int
	// Similar to math/bits.LeadingZeros*
	LeadingZeros() int
	// Similar to math/bits.TrailingZeros*
	TrailingZeros() int
	// Canonical string form
	String() string
}
