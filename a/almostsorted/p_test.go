package almostsorted

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_almostSorted(t *testing.T) {
	for _, tc := range []struct {
		arr  []int32
		want string
	}{
		// {[]int32{4, 2}, "yes\nswap 1 2"},
		// {[]int32{3, 1, 2}, "no"},
		{[]int32{1, 5, 4, 3, 2, 6}, "yes\nreverse 2 5"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, almostSorted2(tc.arr))
		})
	}
}

func almostSorted(arr []int32) {
	fmt.Println(almostSorted2(arr))
}

// Complete the almostSorted function below.
func almostSorted2(arr []int32) string {
	arr = append(arr, 0, 0)
	copy(arr[1:], arr)
	arr[0] = 0
	arr[len(arr)-1] = math.MaxInt32

	var l int32
	n := int32(len(arr))
	for l = 0; l < n-1; l++ {
		if arr[l] > arr[l+1] {
			break
		}
	}
	if l == n-1 {
		return "yes"
	}

	var r int32
	for r = n - 1; r > l; r-- {
		if arr[r] < arr[r-1] {
			break
		}
	}
	if r == l {
		return "no"
	}
	arr[r], arr[l] = arr[l], arr[r]
	for i := int32(1); i < n; i++ {
		if arr[i] < arr[i-1] {
			goto TryReverse
		}
	}
	return fmt.Sprintf("yes\nswap %d %d", l, r)

TryReverse:
	arr[r], arr[l] = arr[l], arr[r] // undo swap
	// Check if reversed order is OK
	if arr[l-1] > arr[r] || arr[r+1] < arr[l] {
		return "no"
	}
	for i := l; i < r; i++ {
		if arr[i] < arr[i+1] {
			return "no"
		}
	}
	return fmt.Sprintf("yes\nreverse %d %d", l, r)
}
