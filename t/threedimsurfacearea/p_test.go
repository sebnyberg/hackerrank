package threedimsurfacearea

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_surfaceArea(t *testing.T) {
	for _, tc := range []struct {
		A    [][]int32
		want int32
	}{
		// {[][]int32{{1}}, 6},
		{[][]int32{
			{1, 3, 4},
			{2, 2, 3},
			{1, 2, 4},
		}, 60},
	} {
		t.Run(fmt.Sprintf("%+v", tc.A), func(t *testing.T) {
			require.Equal(t, tc.want, surfaceArea(tc.A))
		})
	}
}

func surfaceArea(A [][]int32) (res int32) {
	// Seeing as this is a histogram, the surfare area of the top/bottom
	// will always be the number of 2*(non-zero entries in A)
	for i := range A {
		for j := range A[i] {
			if A[i][j] != 0 {
				res += 2
			}
		}
	}

	// Sides are more difficult
	// From left-to-right, the total surfare area is the number of height
	// differences + 1
	// The same approach works for the cube shifted 90 degrees
	for i := range A {
		var prevHeight int32
		for j := range A[i] {
			if A[i][j] > prevHeight {
				res += 2 * (A[i][j] - prevHeight)
			}
			prevHeight = A[i][j]
		}
	}

	for i := range A[0] {
		var prevHeight int32
		for j := range A {
			if A[j][i] > prevHeight {
				res += 2 * (A[j][i] - prevHeight)
			}
			prevHeight = A[j][i]
		}
	}

	return res
}
