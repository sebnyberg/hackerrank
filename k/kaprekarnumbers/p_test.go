package kaprekarnumbers

import (
	"fmt"
	"testing"
)

func Test_kaprekarNumbers(t *testing.T) {
	for _, tc := range []struct {
		p int32
		q int32
	}{
		// {1, 100},
		{1, 10000},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.p, tc.q), func(t *testing.T) {
			kaprekarNumbers(tc.p, tc.q)
			_ = 0
		})
	}
}

func kaprekarNumbers(p int32, q int32) {
	var printed bool
	for n := int(p); n <= int(q); n++ {
		sq := n * n
		width := 0
		for m := sq; m > 0; m /= 10 {
			width++
		}
		rlen := (width + 1) / 2
		r := 0
		multiplier := 1
		for i := 0; i < rlen; i++ {
			r += (sq % 10) * multiplier
			sq /= 10
			multiplier *= 10
		}
		if sq+r == n {
			if printed {
				fmt.Print(" ")
			}
			fmt.Print(n)
			printed = true
		}
	}
	if !printed {
		fmt.Print("INVALID RANGE")
	}
}
