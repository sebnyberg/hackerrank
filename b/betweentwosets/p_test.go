package betweentwosets

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getTotalX(t *testing.T) {
	for _, tc := range []struct {
		a    []int32
		b    []int32
		want int32
	}{
		{[]int32{2, 4}, []int32{16, 32, 96}, 3},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.a, tc.b), func(t *testing.T) {
			require.Equal(t, tc.want, getTotalX(tc.a, tc.b))
		})
	}
}

func getTotalX(a []int32, b []int32) (res int32) {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	var candidates []int32
	for d := a[0]; d <= b[0]; d += a[0] {
		for _, aa := range a {
			if d < aa || d%aa != 0 {
				goto ContinueLoop
			}
		}
		candidates = append(candidates, d)
	ContinueLoop:
	}
	// elements in a are factors of the numbers in candidates
	// check which of candidates also evenly divide elements of b
	for _, d := range candidates {
		for _, bb := range b {
			if bb%d != 0 {
				goto ContinueLoop2
			}
		}
		res++
	ContinueLoop2:
	}

	return res
}
