package cutthesticks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_cutTheSticks(t *testing.T) {
	for _, tc := range []struct {
		arr  []int32
		want []int32
	}{
		{[]int32{5, 4, 4, 2, 2, 8}, []int32{6, 4, 2, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, cutTheSticks(tc.arr))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int32) (res []int32) {
	counts := make(map[int32]int)
	minVal := 1001
	for _, n := range arr {
		minVal = min(minVal, int(n))
		counts[n]++
	}

	n := len(arr)
	for {
		if n == 0 {
			return res
		}
		res = append(res, int32(n))
		newCounts := make(map[int32]int)
		newMinVal := 1001
		for v := range counts {
			newVal := int(v) - minVal
			if newVal <= 0 {
				n -= counts[v]
			} else {
				newCounts[int32(newVal)] += counts[v]
				newMinVal = min(newMinVal, newVal)
			}
		}
		counts = newCounts
		minVal = newMinVal
	}
}
