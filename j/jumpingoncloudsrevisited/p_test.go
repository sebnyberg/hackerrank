package jumpingoncloudsrevisited

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jumpingOnClouds(t *testing.T) {
	for _, tc := range []struct {
		c    []int32
		k    int32
		want int32
	}{
		{[]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2, 92},
		{[]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3, 80},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.c, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, jumpingOnClouds(tc.c, tc.k))
		})
	}
}

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32, k int32) (e int32) {
	n := len(c)
	pos := 0
	e = 100
	for {
		// Jump
		pos = (pos + int(k)) % n
		e--
		if c[pos] == 1 {
			e -= 2
		}
		// if pos == 0, we are done
		if pos == 0 {
			return e
		}
	}
}
