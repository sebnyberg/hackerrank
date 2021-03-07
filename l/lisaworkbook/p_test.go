package lisaworkbook

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_workbook(t *testing.T) {
	for _, tc := range []struct {
		n    int32
		k    int32
		arr  []int32
		want int32
	}{
		{5, 3, []int32{4, 2, 6, 1, 10}, 4},
	} {
		t.Run(fmt.Sprintf("%v/%v/%+v", tc.n, tc.k, tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, workbook(tc.n, tc.k, tc.arr))
		})
	}
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// Complete the workbook function below.
func workbook(n int32, k int32, arr []int32) (nspecial int32) {
	page := int32(1)
	for chapter := int32(1); chapter <= n; chapter++ {
		firstProblem, lastProblem := int32(1), min(arr[chapter-1], int32(1)+k-1)
		if page >= firstProblem && page <= lastProblem {
			nspecial++
		}
		page++
		for lastProblem != arr[chapter-1] {
			firstProblem += k
			lastProblem = min(lastProblem+k, arr[chapter-1])
			if page >= firstProblem && page <= lastProblem {
				nspecial++
			}
			page++
		}
	}
	return nspecial
}
