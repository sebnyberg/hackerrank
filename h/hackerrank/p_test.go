package hackerrank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hackerrankInString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"hereiamstackerrank", "YES"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, hackerrankInString(tc.s))
		})
	}
}

// Complete the hackerrankInString function below.
func hackerrankInString(s string) string {
	hackerrank := "hackerrank"
	pos := 0
	for _, r := range s {
		if pos == len(hackerrank) {
			break
		}
		if r == rune(hackerrank[pos]) {
			pos++
		}
	}
	if pos == len(hackerrank) {
		return "YES"
	}
	return "NO"
}
