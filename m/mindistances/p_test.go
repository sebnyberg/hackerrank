package mindistances

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumDistances(t *testing.T) {
	for _, tc := range []struct {
		a    []int32
		want int32
	}{
		{[]int32{7, 1, 3, 4, 1, 7}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.a), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDistances(tc.a))
		})
	}
}

func abs(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// Complete the minimumDistances function below.
func minimumDistances(a []int32) int32 {
	nlocs := make(map[int32]int32)
	minDist := int32(math.MaxInt32)
	for i, aa := range a {
		if _, exists := nlocs[aa]; !exists {
			nlocs[aa] = int32(i)
			continue
		}
		minDist = min(minDist, abs(int32(i)-nlocs[aa]))
	}
	if minDist == math.MaxInt32 {
		return -1
	}
	return minDist
}
