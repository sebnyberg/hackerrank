package encryption

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_encryption(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"chillout", "clu hlt io"},
		{"if man was meant to stay on the ground god would have given us roots", "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, encryption(tc.s))
		})
	}
}

func encryption(s string) string {
	sb := make([]byte, 0, len(s))
	for i := range s {
		if s[i] != ' ' {
			sb = append(sb, s[i])
		}
	}
	n := len(sb)
	res := make([]byte, 0, n)
	a := string(sb)
	_ = a
	nrows, ncols := findDimensions(n)
	for i := 0; i < ncols; i++ {
		for j := 0; j < nrows; j++ {
			pos := (j * ncols) + i
			if pos >= n {
				continue
			}
			res = append(res, sb[pos])
		}
		if i < ncols-1 {
			res = append(res, ' ')
		}
	}

	ss := string(res)
	_ = ss
	return string(res)
}

func findDimensions(n int) (nrows, ncols int) {
	i := 1
	for i*i < n {
		i++
	}
	// We know that i-1*i-1 < n < i*i
	// Now we need to test whether i-1*i < n
	if i-1*i >= n {
		return i - 1, i
	}
	return i, i
}
