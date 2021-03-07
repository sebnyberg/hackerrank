package thetimeinwords

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_timeInWords(t *testing.T) {
	for _, tc := range []struct {
		h    int32
		m    int32
		want string
	}{
		{5, 0, "five o' clock"},
		{5, 1, "one minute past five"},
		{5, 10, "ten minutes past five"},
		{5, 15, "quarter past five"},
		{5, 30, "half past five"},
		{5, 40, "twenty minutes to six"},
		{5, 45, "quarter to six"},
		{5, 47, "thirteen minutes to six"},
		{5, 28, "twenty eight minutes past five"},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.h, tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, timeInWords(tc.h, tc.m))
		})
	}
}

var hours = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
}

var minutes = map[int]string{
	1:  "one minute",
	2:  "two minutes",
	3:  "three minutes",
	4:  "four minutes",
	5:  "five minutes",
	6:  "six minutes",
	7:  "seven minutes",
	8:  "eight minutes",
	9:  "nine minutes",
	10: "ten minutes",
	11: "eleven minutes",
	12: "twelve minutes",
	13: "thirteen minutes",
	14: "fourteen minutes",
	15: "quarter",
	16: "sixteen minutes",
	17: "seventeen minutes",
	18: "eighteen minutes",
	19: "nineteen minutes",
	20: "twenty minutes",
	21: "twenty one minutes",
	22: "twenty two minutes",
	23: "twenty three minutes",
	24: "twenty four minutes",
	25: "twenty five minutes",
	26: "twenty six minutes",
	27: "twenty seven minutes",
	28: "twenty eight minutes",
	29: "twenty nine minutes",
	30: "half",
}

// Complete the timeInWords function below.
func timeInWords(h int32, m int32) string {
	hour := hours[int(h)]
	switch {
	case m == 0:
		return hour + " o' clock"
	case m <= 30:
		return minutes[int(m)] + " past " + hour
	default:
		m = 60 - m
		h += 1
		if h == 12 {
			h = 1
		}
		return minutes[int(m)] + " to " + hours[int(h)]
	}
}
