package beautifultriplets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_beautifulTriplets(t *testing.T) {
	for _, tc := range []struct {
		d    int32
		arr  []int32
		want int32
	}{
		{1, []int32{2, 2, 3, 4, 5}, 3},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.d, tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, beautifulTriplets(tc.d, tc.arr))
		})
	}
}

// Complete the beautifulTriplets function below.
func beautifulTriplets(d int32, arr []int32) (res int32) {
	ncounts := make(map[int32]int32)
	for _, a := range arr {
		ncounts[a]++
	}
	for k, counts := range ncounts {
		res += counts * ncounts[k+d] * ncounts[k+2*d]
	}
	return res
}
