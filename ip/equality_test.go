package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	{
		v4 := ip.V4()
		one := v4.FromInt(1)
		two := v4.FromInt(2)
		assert.True(t, ip.Eq(one, v4.FromInt(1)))
		assert.True(t, ip.Eq(one, one))
		assert.False(t, ip.Eq(one, two))
		assert.False(t, ip.Eq(two, one))
	}
	{
		v6 := ip.V6()
		one := v6.FromInt(1)
		complement := one.Not()
		complement2 := one.Not()
		assert.True(t, ip.Eq(one, v6.FromInt(1)))
		assert.True(t, ip.Eq(one, one))
		assert.True(t, ip.Eq(complement, complement))
		assert.True(t, ip.Eq(complement, complement2))
		assert.False(t, ip.Eq(one, complement))
		assert.False(t, ip.Eq(complement, one))
	}
}
