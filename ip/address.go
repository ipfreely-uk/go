package ip

// Internet protocol version
type Version uint8

const (
	// Internet protocol version 4
	Version4 Version = 4
	// Internet protocol version 6
	Version6 Version = 6
)

const (
	// Number of bits in IPv4 address (32)
	Width4 int = 32
	// Number of bits in IPv6 address (128)
	Width6 int = 128
)

// The parts of [Int] without generic typing.
// [Addr4] and [Addr6] are the only types that can conform to this interface.
type Address interface {
	Spec
	// Address as byte slice
	Bytes() (slice []byte)
	// Normalized string form
	String() (address string)
	// Equivalent to math/bits.LeadingZeros
	LeadingZeros() (count int)
	// Equivalent to math/bits.TrailingZeros
	TrailingZeros() (count int)
	// Approximation to float64
	Float64() (approximation float64)
}

// IP address as generic unsigned integer type.
// [Addr4] and [Addr6] are the only types that can conform to this interface.
type Int[A Address] interface {
	Address
	// IP address family - [V4] or [V6]
	Family() Family[A]
	// Addition with overflow
	Add(addend A) (sum A)
	// Subtraction with overflow
	Subtract(subtrahend A) (difference A)
	// Multiplication with overflow
	Multiply(factor A) (product A)
	// Division
	Divide(denominator A) (quotient A)
	// Modulus
	Mod(denominator A) (remainder A)
	// Bitwise NOT
	Not() (complement A)
	// Bitwise AND
	And(operand A) (address A)
	// Bitwise OR
	Or(operand A) (address A)
	// Bitwise XOR
	Xor(operand A) (address A)
	// Bit shift. Use negative int for left shift; use positive int for right shift.
	Shift(bits int) (address A)
	// Returns 1 if operand is less than this.
	// Returns -1 if operand is more than this.
	// Returns 0 if operand is equal.
	Compare(address A) int
}
