package appendanddelete

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_appendAndDelete(t *testing.T) {
	for _, tc := range []struct {
		s    string
		t    string
		k    int32
		want string
	}{
		{"hackerhappy", "hackerrank", 9, "Yes"},
		{"abcd", "abcdert", 10, "No"},
	} {
		t.Run(fmt.Sprintf("%v/%v/%v", tc.s, tc.t, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, appendAndDelete(tc.s, tc.t, tc.k))
		})
	}
}

func appendAndDelete(s string, t string, k int32) string {
	// Remove whole string and add it again
	// After removing all of s, we can add as many as we want to get k
	if len(s)+len(t) <= int(k) {
		return "Yes"
	}

	var minCost int
	if len(t) > len(s) {
		minCost = len(t) - len(s)
		t = t[:len(s)]
	} else if len(s) > len(t) {
		minCost = len(s) - len(t)
		s = s[:len(t)]
	}
	for i := range s {
		if s[i] != t[i] {
			minCost += (len(s) - i) * 2
			break
		}
	}

	if minCost > int(k) {
		return "No"
	}

	if (int(k)-minCost)%2 == 0 {
		return "Yes"
	}
	return "No"
}
