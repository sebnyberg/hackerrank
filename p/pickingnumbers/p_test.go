package pickingnumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pickingNumbers(t *testing.T) {
	for _, tc := range []struct {
		a    []int32
		want int32
	}{
		{[]int32{4, 6, 5, 3, 3, 1}, 3},
		{[]int32{1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.a), func(t *testing.T) {
			require.Equal(t, tc.want, pickingNumbers(tc.a))
		})
	}
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func pickingNumbers(a []int32) (maxLen int32) {
	seen := make(map[int32]struct{})
	for _, n := range a {
		if _, exists := seen[n]; exists {
			continue
		}
		var aboveCount int32
		var belowCount int32
		for _, m := range a {
			d := m - n
			if d == 0 || d == 1 {
				aboveCount++
			}
			if d == 0 || d == -1 {
				belowCount++
			}
		}
		maxLen = max(maxLen, max(aboveCount, belowCount))
		seen[n] = struct{}{}
	}

	return maxLen
}
