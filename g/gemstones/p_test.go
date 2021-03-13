package gemstones

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGemstones(t *testing.T) {
	for _, tc := range []struct {
		arr  []string
		want int
	}{
		{[]string{"abcdde", "baccd", "eeabg"}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, gemstones(tc.arr))
		})
	}
}

// Complete the gemstones function below.
func gemstones(arr []string) int32 {
	var gems int
	for i, s := range arr {
		var newGems int
		for _, ch := range s {
			newGems |= (1 << (ch - 'a'))
		}
		if i == 0 {
			gems |= newGems
		} else {
			gems &= newGems
		}
	}
	val := 1
	res := 0
	for i := 0; i < 27; i++ {
		if gems&val > 0 {
			res++
		}
		val <<= 1
	}
	return int32(res)
}
