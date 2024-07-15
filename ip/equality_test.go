package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	one := ip.V4().FromInt(1)
	two := ip.V4().FromInt(2)
	assert.True(t, ip.Eq(one, one))
	assert.False(t, ip.Eq(one, two))
	assert.False(t, ip.Eq(two, one))
}
