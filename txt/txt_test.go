package txt_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/txt"
	"github.com/stretchr/testify/assert"
)

func TestCommaDelim(t *testing.T) {
	{
		zero := big.NewInt(0)
		expected := "0"
		actual := txt.CommaDelim(zero)
		assert.Equal(t, expected, actual)
	}
	{
		thousand := big.NewInt(1000)
		expected := "1,000"
		actual := txt.CommaDelim(thousand)
		assert.Equal(t, expected, actual)
	}
	{
		million := big.NewInt(1000000)
		expected := "1,000,000"
		actual := txt.CommaDelim(million)
		assert.Equal(t, expected, actual)
	}
	{
		notmutated := big.NewInt(1000000)
		expected := txt.CommaDelim(notmutated)
		actual := txt.CommaDelim(notmutated)
		assert.Equal(t, expected, actual)
	}
}
