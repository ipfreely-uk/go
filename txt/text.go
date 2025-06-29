// This package is used in example code and does not form part of the library.
// It may be removed without warning.
package txt

import "math/big"

// TODO: remove this file when https://pkg.go.dev/golang.org/x/text Decimal supports big.Int

// Format using comma-delimited thousand separators
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
