package pmagicsquareforming

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_formingMagicSquare(t *testing.T) {
	for _, tc := range []struct {
		s    [][]int32
		want int32
	}{
		{[][]int32{
			{5, 3, 4},
			{1, 5, 8},
			{6, 4, 2},
		}, 7},
		{[][]int32{
			{4, 8, 2},
			{4, 5, 7},
			{6, 1, 6},
		}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, formingMagicSquare(tc.s))
		})
	}
}

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {
	squares := genMagicSquares()
	minDist := int32(math.MaxInt32)
	// For each magic square
	for _, square := range squares {
		// Try to adjust s so that it matches the magic square
		d := editDistance(s, square, 0, 0)
		if d < minDist {
			minDist = d
		}
	}

	return minDist
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func editDistance(s [][]int32, square [3][3]int, pos int, curDist int) int32 {
	if pos == 9 {
		return int32(curDist)
	}
	pre := s[pos/3][pos%3]
	post := square[pos/3][pos%3]
	dist := abs(int(pre) - post)
	s[pos/3][pos%3] = int32(post)
	d := editDistance(s, square, pos+1, dist+curDist)
	s[pos/3][pos%3] = pre
	return d
}

type MagicSquareFinder struct {
	magicSquares [][3][3]int
	square       [3][3]int
}

func genMagicSquares() [][3][3]int {
	f := &MagicSquareFinder{}
	f.findMagicSquares(0, 0)
	return f.magicSquares
}

func (f *MagicSquareFinder) findMagicSquares(pos int, taken int) {
	if pos == 9 {
		if validateMagicSquare(f.square) {
			f.magicSquares = append(f.magicSquares, f.square)
		}
	}
	bit := 1
	for n := 1; n <= 9; n++ {
		bit <<= 1
		if taken&bit > 0 {
			continue
		}
		f.square[pos/3][pos%3] = n
		taken |= bit
		f.findMagicSquares(pos+1, taken)
		taken &^= bit
		f.square[pos/3][pos%3] = 0
	}
}

func validateMagicSquare(square [3][3]int) bool {
	var rowSums [3]int
	var colSums [3]int
	var diagSums [2]int
	for i := range square {
		diagSums[0] += square[i][i]
		diagSums[1] += square[i][3-i-1]
		for j := range square[i] {
			rowSums[i] += square[i][j]
			colSums[i] += square[j][i]
		}
	}
	first := rowSums[0]
	for i := range rowSums {
		if rowSums[i] != first || colSums[i] != first {
			return false
		}
	}
	if diagSums[0] != first {
		return false
	}
	return diagSums[1] == first
}
