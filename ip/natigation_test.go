package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestMinAddress(t *testing.T) {
	expected := ip.V6().FromInt(0)
	actual := ip.MinAddress(ip.V6())
	assert.Equal(t, expected, actual)
}

func TestMaxAddress(t *testing.T) {
	expected := ip.V4().FromInt(0xFFFFFFFF)
	actual := ip.MaxAddress(ip.V4())
	assert.Equal(t, expected, actual)
}

func TestNext(t *testing.T) {
	expected := ip.V4().FromInt(0)
	max := ip.V4().FromInt(0xFFFFFFFF)
	actual := ip.Next(max)
	assert.Equal(t, expected, actual)
}

func TestPrev(t *testing.T) {
	expected := ip.V6().FromInt(0)
	one := ip.V6().FromInt(1)
	actual := ip.Prev(one)
	assert.Equal(t, expected, actual)
}
