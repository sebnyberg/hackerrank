package fairrations

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fairRations(t *testing.T) {
	// for _, tc := range []struct {
	// 	B    []int32
	// 	want int32
	// }{
	// 	{[]int32{2, 3, 4, 5, 6}, 4},
	// } {
	// 	t.Run(fmt.Sprintf("%+v", tc.B), func(t *testing.T) {
	// 		require.Equal(t, tc.want, fairRations(tc.B))
	// 	})
	// }

	t.Run("testdata", func(t *testing.T) {
		f, _ := os.Open("testdata/input")
		content, _ := ioutil.ReadAll(f)
		nstrs := strings.Split(string(content), " ")
		ns := make([]int32, len(nstrs))
		for i, nstr := range nstrs {
			n, _ := strconv.Atoi(nstr)
			ns[i] = int32(n)
		}
		require.Equal(t, -1, fairRations(ns))
	})
}

// Complete the fairRations function below.
func fairRations(B []int32) (loaves int32) {
	n := len(B)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return B[0] % 2
	}
	for i := 1; i < n-1; i++ {
		if B[i-1]%2 == 0 {
			continue
		}
		loaves += 2
		B[i-1]++
		B[i]++
	}
	if B[n-2]%2 == 1 {
		if B[n-1]%2 == 0 {
			return -1
		}
		return loaves + 2
	}
	if B[n-1]%2 == 1 {
		return -1
	}
	return loaves
}
