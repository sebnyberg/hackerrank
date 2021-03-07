package nondivisiblesubset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nonDivisibleSubset(t *testing.T) {
	for _, tc := range []struct {
		k    int32
		s    []int32
		want int32
	}{
		{9, []int32{422346306, 940894801, 696810740, 862741861, 85835055, 313720373}, 5},
		{4, []int32{19, 10, 12, 10, 24, 25, 22}, 3},
		{3, []int32{1, 7, 2, 4}, 3},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.k, tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, nonDivisibleSubset(tc.k, tc.s))
		})
	}
}

func nonDivisibleSubset(k int32, s []int32) (res int32) {
	valCount := make(map[int32]int32, k)
	for _, ss := range s {
		valCount[ss%k]++
	}
	if valCount[0] > 0 {
		res++
	}
	for i := int32(1); i <= k/2; i++ {
		if i == k-i {
			if valCount[int32(i)] > 0 {
				res++
				break
			}
		}
		if valCount[i] > valCount[k-i] {
			res += valCount[i]
		} else {
			res += valCount[k-i]
		}
	}
	return res
}
