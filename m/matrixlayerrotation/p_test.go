package matrixlayerrotation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_matrixRotation(t *testing.T) {
	for _, tc := range []struct {
		matrix    [][]int32
		rotations int32
		want      [][]int32
	}{
		// {
		// 	[][]int32{
		// 		{1, 2, 3, 4},
		// 		{12, 1, 2, 5},
		// 		{11, 4, 3, 6},
		// 		{10, 9, 8, 7},
		// 	}, 2,
		// 	[][]int32{
		// 		{3, 4, 5, 6},
		// 		{2, 3, 4, 7},
		// 		{1, 2, 1, 8},
		// 		{12, 11, 10, 9},
		// 	},
		// },
		{
			[][]int32{
				{1, 2, 3, 4},
				{7, 8, 9, 10},
				{13, 14, 15, 16},
				{19, 20, 21, 22},
				{25, 26, 27, 28},
			}, 7,
			[][]int32{
				{28, 27, 26, 25},
				{22, 9, 15, 19},
				{16, 8, 21, 13},
				{10, 14, 20, 7},
				{4, 3, 2, 1},
			},
		},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.matrix, tc.rotations), func(t *testing.T) {
			require.Equal(t, tc.want, matrixRotation2(tc.matrix, tc.rotations))
		})
	}
}

func matrixRotation(matrix [][]int32, r int32) {
	res := matrixRotation2(matrix, r)
	for i := range res {
		if i > 0 {
			fmt.Print("\n")
		}
		for j := range res[i] {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(res[i][j])
		}
	}
}

func matrixRotation2(matrix [][]int32, r int32) [][]int32 {
	// Figure out the "inner loop" in the matrix
	m, n := len(matrix), len(matrix[0])
	innerM, innerN, nLoops := 2, 2, 0
	if m >= n {
		innerM += m - n
		nLoops = n / 2
	} else {
		innerN += n - m
		nLoops = m / 2
	}
	for offset := nLoops - 1; offset >= 0; offset-- {
		innerLen := 2*innerN + 2*innerM - 4
		ns := make([]int32, innerLen)
		loop := Loop{n: innerN, m: innerM}
		// put elements into array
		for i := 0; i < innerLen; i++ {
			ns[i] = matrix[offset+loop.i][offset+loop.j]
			loop.next()
		}
		// shift loop to alter insert position
		for i := 0; i < int(r)%innerLen; i++ {
			loop.next()
		}
		// insert again
		for i := 0; i < innerLen; i++ {
			matrix[offset+loop.i][offset+loop.j] = ns[i]
			loop.next()
		}
		// expand loop
		innerN += 2
		innerM += 2
	}
	return matrix
}

type Loop struct {
	i, j int
	n, m int
}

// next() moves to the next position in the loop
func (l *Loop) next() {
	switch {
	case l.i == 0 && l.j > 0:
		l.j--
	case l.i == l.m-1 && l.j < l.n-1:
		l.j++
	case l.j == 0 && l.i < l.m-1:
		l.i++
	case l.j == l.n-1 && l.i > 0:
		l.i--
	}
}
