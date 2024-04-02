package compare_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	one := ip.V4().FromInt(1)
	two := ip.V4().FromInt(2)
	assert.True(t, compare.Eq(one, one))
	assert.False(t, compare.Eq(one, two))
	assert.False(t, compare.Eq(two, one))
}

func TestMin(t *testing.T) {
	one := ip.V4().FromInt(1)
	two := ip.V4().FromInt(2)
	assert.Equal(t, one, compare.Min(one, two))
	assert.Equal(t, one, compare.Min(two, one))
	assert.Equal(t, one, compare.Min(one, one))
}

func TestMax(t *testing.T) {
	one := ip.V4().FromInt(1)
	two := ip.V4().FromInt(2)
	assert.Equal(t, two, compare.Max(one, two))
	assert.Equal(t, two, compare.Max(two, one))
	assert.Equal(t, one, compare.Max(one, one))
}
