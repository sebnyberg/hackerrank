package libraryfine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_libraryFine(t *testing.T) {
	for _, tc := range []struct {
		in   []int32
		want int32
	}{
		{[]int32{9, 6, 2015, 6, 6, 2015}, 45},
		{[]int32{1, 1, 2001, 1, 1, 2000}, 10000},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, libraryFine(tc.in[0], tc.in[1], tc.in[2], tc.in[3], tc.in[4], tc.in[5]))
		})
	}
}

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	switch {
	case y1 > y2:
		return 10000
	case y1 < y2:
		return 0
	case m1 > m2:
		return (m1 - m2) * 500
	case m2 > m1:
		return 0
	case d1 > d2:
		return (d1 - d2) * 15
	default:
		return 0
	}
}
