package queensattack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_queensAttack(t *testing.T) {
	for _, tc := range []struct {
		in        []int32
		obstacles [][]int32
		want      int32
	}{
		{[]int32{8, 1, 4, 4}, [][]int32{{3, 5}}, 24},
		{[]int32{4, 0, 4, 4}, nil, 9},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.in, tc.obstacles), func(t *testing.T) {
			require.Equal(t, tc.want, queensAttack(tc.in[0], tc.in[1], tc.in[2], tc.in[3], tc.obstacles))
		})
	}
}

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) (res int32) {
	obstaclesMap := make(map[[2]int32]bool)
	for _, obstacle := range obstacles {
		obstaclesMap[[2]int32{obstacle[0], obstacle[1]}] = true
	}

	ok := func(row, col int32) bool {
		return row > 0 && col > 0 && row <= n && col <= n && !obstaclesMap[[2]int32{row, col}]
	}
	for _, direction := range [][2]int32{
		{1, -1}, {1, 0}, {1, 1}, {0, 1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
	} {
		pos := [2]int32{r_q + direction[0], c_q + direction[1]}
		for ok(pos[0], pos[1]) {
			res++
			pos[0] += direction[0]
			pos[1] += direction[1]
		}
	}

	return res
}
