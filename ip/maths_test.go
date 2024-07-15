package ip_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

var size4 = func() *big.Int {
	last := ip.ToBigInt(ip.MaxAddress(ip.V4()))
	return last.Add(last, big.NewInt(1))
}()
var size6 = func() *big.Int {
	last := ip.ToBigInt(ip.MaxAddress(ip.V6()))
	return last.Add(last, big.NewInt(1))
}()

type op func(x, y *big.Int) *big.Int

var add op = func(x, y *big.Int) *big.Int { return big.NewInt(0).Add(x, y) }
var mul op = func(x, y *big.Int) *big.Int { return big.NewInt(0).Mul(x, y) }
var sub op = func(x, y *big.Int) *big.Int { return big.NewInt(0).Sub(x, y) }

func testMaths[A ip.Number[A]](t *testing.T, addresses []A) {
	for _, a := range addresses {
		expected, _ := ip.ToBigInt(a).Float64()
		actual := a.Float64()
		assert.Equal(t, expected, actual, a.String())

		assert.Equal(t, a.Version(), a.Family().Version())
	}

	for _, a := range addresses {
		bigA := ip.ToBigInt(a)
		for _, b := range addresses {
			bigB := ip.ToBigInt(b)
			testAdd(t, a, b, bigA, bigB)
			testMultiply(t, a, b, bigA, bigB)
			testSubtract(t, a, b, bigA, bigB)
			testDivide(t, a, b, bigA, bigB)
		}
	}
}

func testAdd[A ip.Number[A]](t *testing.T, a, b A, bigA, bigB *big.Int) {
	r := perform(a.Family(), add, bigA, bigB)

	expected, _ := ip.FromBigInt(a.Family(), r)
	actual := a.Add(b)
	msg := fmt.Sprintf("%s + %s = %s    expected %s", a.String(), b.String(), actual.String(), expected.String())
	assert.Equal(t, expected, actual, msg)
}

func testMultiply[A ip.Number[A]](t *testing.T, a, b A, bigA, bigB *big.Int) {
	r := perform(a.Family(), mul, bigA, bigB)

	expected, _ := ip.FromBigInt(a.Family(), r)
	actual := a.Multiply(b)
	msg := fmt.Sprintf("%s * %s = %s    expected %s", a.String(), b.String(), actual.String(), expected.String())
	assert.Equal(t, expected, actual, msg)
}

func testSubtract[A ip.Number[A]](t *testing.T, a, b A, bigA, bigB *big.Int) {
	r := perform(a.Family(), sub, bigA, bigB)

	expected, _ := ip.FromBigInt(a.Family(), r)
	actual := a.Subtract(b)
	msg := fmt.Sprintf("%s - %s = %s    expected %s", a.String(), b.String(), actual.String(), expected.String())
	assert.Equal(t, expected, actual, msg)
}

func testDivide[A ip.Number[A]](t *testing.T, a, b A, bigA, bigB *big.Int) {
	zero := ip.MinAddress(a.Family())

	if ip.Eq(b, zero) {
		assert.Panics(t, func() {
			a.Divide(b)
		})
		assert.Panics(t, func() {
			a.Mod(b)
		})
	} else {
		r := big.NewInt(0).Div(bigA, bigB)
		expected, _ := ip.FromBigInt(a.Family(), r)
		actual := a.Divide(b)
		msg := fmt.Sprintf("%s / %s = %s    expected %s", a.String(), b.String(), actual.String(), expected.String())
		assert.Equal(t, expected, actual, msg)

		r = big.NewInt(0).Mod(bigA, bigB)
		expected, _ = ip.FromBigInt(a.Family(), r)
		actual = a.Mod(b)
		msg = fmt.Sprintf("%s mod %s = %s    expected %s", a.String(), b.String(), actual.String(), expected.String())
		assert.Equal(t, expected, actual, msg)
	}
}

func perform[A ip.Number[A]](family ip.Family[A], f op, x, y *big.Int) *big.Int {
	z := f(x, y)
	var size *big.Int
	if family.Version() == ip.Version4 {
		size = size4
	} else {
		size = size6
	}
	return z.Mod(z, size)
}
