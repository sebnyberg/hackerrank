package twopluses

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoPluses(t *testing.T) {
	for _, tc := range []struct {
		grid []string
		want int32
	}{
		// {[]string{
		// 	"GGGGGG",
		// 	"GBBBGB",
		// 	"GGGGGG",
		// 	"GGBBGB",
		// 	"GGGGGG",
		// }, 5,
		// },
		// {[]string{
		// 	"BGBBGB",
		// 	"GGGGGG",
		// 	"BGBBGB",
		// 	"GGGGGG",
		// 	"BGBBGB",
		// 	"BGBBGB",
		// }, 25,
		{[]string{
			"GGGGGGGG",
			"GBGBGGBG",
			"GBGBGGBG",
			"GGGGGGGG",
			"GBGBGGBG",
			"GGGGGGGG",
			"GBGBGGBG",
			"GGGGGGGG",
		}, 81,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, twoPluses(tc.grid))
		})
	}
}

func twoPluses(grid []string) int32 {
	byteGrid := make([][]byte, len(grid))
	for i := range grid {
		byteGrid[i] = []byte(grid[i])
	}

	// This is a brutal brute-force
	// First check if there are at least two good cells in the grid,
	// this is the baseline and eliminates a lot of tries
	gCount := 0
	for i := range grid {
		gCount += strings.Count(grid[i], "G")
	}
	if gCount == 2 {
		return 1
	} else if gCount == 0 {
		return 0
	}

	// At this point we know that a solution must include at least one plus
	// of size 4 or greater, this eliminates a lot of branches
	maxSize := 2
	for _, p := range findPluses(byteGrid, 2) {
		// Mark the plus
		byteGrid[p.i][p.j] = 'F'
		for i := 1; i < p.size; i++ {
			byteGrid[p.i+i][p.j] = 'F'
			byteGrid[p.i-i][p.j] = 'F'
			byteGrid[p.i][p.j+i] = 'F'
			byteGrid[p.i][p.j-i] = 'F'
		}
		// For every other plus, check max size
		for _, p2 := range findPluses(byteGrid, 1) {
			size := (1 + (p.size-1)*4) * (1 + (p2.size-1)*4)
			if size > maxSize {
				maxSize = size
			}
		}
		// Unmark the plus
		byteGrid[p.i][p.j] = 'G'
		for i := 1; i < p.size; i++ {
			byteGrid[p.i+i][p.j] = 'G'
			byteGrid[p.i-i][p.j] = 'G'
			byteGrid[p.i][p.j+i] = 'G'
			byteGrid[p.i][p.j-i] = 'G'
		}
	}
	return int32(maxSize)
}

type plus struct {
	i    int
	j    int
	size int
}

func findPluses(byteGrid [][]byte, minSize int) (res []plus) {
	n, m := len(byteGrid[0]), len(byteGrid)
	for i := range byteGrid {
		for j := range byteGrid[i] {
			if byteGrid[i][j] == 'G' {
				for size := 1; ; size++ {
					if size >= minSize {
						res = append(res, plus{i, j, size})
					}
					if i-size < 0 || byteGrid[i-size][j] != 'G' {
						break
					}
					if i+size >= m || byteGrid[i+size][j] != 'G' {
						break
					}
					if j-size < 0 || byteGrid[i][j-size] != 'G' {
						break
					}
					if j+size >= n || byteGrid[i][j+size] != 'G' {
						break
					}
				}
			}
		}
	}
	return res
}
