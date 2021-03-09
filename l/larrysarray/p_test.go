package larrysarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_larrysArray(t *testing.T) {
	for _, tc := range []struct {
		A    []int32
		want string
	}{
		{[]int32{1, 6, 5, 2, 3, 4}, "NO"},
		{[]int32{3, 1, 2, 4}, "YES"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.A), func(t *testing.T) {
			require.Equal(t, tc.want, larrysArray(tc.A))
		})
	}
}

// Complete the larrysArray function below.
func larrysArray(A []int32) string {
	// brute force
	for i := 0; i < len(A)-2; i++ {
		if A[i] == int32(i+1) { // nothing to do
			continue
		}
		// j marks the target number
		var j int
		for j = i; j < len(A); j++ {
			if A[j] == int32(i+1) {
				break
			}
		}
		// k is just a swap number
		var k int
		for k = i + 1; k == j; k++ {
		}
		if k < j {
			A[i], A[k], A[j] = A[j], A[i], A[k]
		} else {
			A[i], A[j], A[k] = A[j], A[k], A[i]
		}
	}
	for i := 1; i < len(A); i++ {
		if A[i-1] > A[i] {
			return "NO"
		}
	}
	return "YES"
}
