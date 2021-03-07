package manasaandstones

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_stones(t *testing.T) {
	for _, tc := range []struct {
		n    int32
		a    int32
		b    int32
		want []int32
	}{
		{3, 1, 2, []int32{2, 3, 4}},
		{2, 2, 3, []int32{2, 3}},
	} {
		t.Run(fmt.Sprintf("%v/%v/%v", tc.n, tc.a, tc.b), func(t *testing.T) {
			require.Equal(t, tc.want, stones(tc.n, tc.a, tc.b))
		})
	}
}

// Complete the stones function below.
func stones(n int32, a int32, b int32) (res []int32) {
	n--
	results := make(map[int32]struct{})
	for i := int32(0); i <= n; i++ {
		v := (i * a) + ((n - i) * b)
		results[v] = struct{}{}
	}
	for v := range results {
		res = append(res, v)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}
