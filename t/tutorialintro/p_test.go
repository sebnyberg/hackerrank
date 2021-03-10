package tutorialintro

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_introTutorial(t *testing.T) {
	for _, tc := range []struct {
		V    int32
		arr  []int32
		want int32
	}{
		// {4, []int32{1, 4, 5, 7, 9, 12}, 1},
		// {9, []int32{1, 4, 5, 7, 9, 12}, 4},
		{10, []int32{1, 4, 5, 7, 9, 12}, 4},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.V, tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, introTutorial(tc.V, tc.arr))
		})
	}
}

// Complete the introTutorial function below.
func introTutorial(V int32, arr []int32) int32 {
	var l, r int32 = 0, int32(len(arr))
	for l < r {
		m := (l + r) / 2
		if V > arr[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
