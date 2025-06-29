// This package is used in example code and does not form part of the library.
// It may be removed without warning when https://pkg.go.dev/golang.org/x/text Decimal supports [big.Int].
package txt

import "math/big"

// Format using comma-delimited thousand separators.
// This function does not support negative numbers.
func CommaDelim(bi *big.Int) string {
	s := bi.String()
	r := ""
	for len(s) > 3 {
		n := len(s) - 3
		segment := s[n:]
		s = s[:n]
		if len(r) > 0 {
			r = segment + "," + r
		} else {
			r = segment
		}
	}
	if len(s) > 0 && len(r) > 0 {
		r = s + "," + r
	} else {
		r = s
	}
	return r
}
