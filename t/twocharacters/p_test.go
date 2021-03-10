package twocharacters

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_alternate(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int32
	}{
		{"beabeefeab", 5},
		{"asvkugfiugsalddlasguifgukvsa", 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, alternate(tc.s))
		})
	}
}

// Complete the alternate function below.
func alternate(s string) int32 {
	var counts [26]int32
	for _, r := range s {
		counts[r-'a']++
	}
	var maxAlternating int32
	for i, n1 := range counts {
		for j := i + 1; j < 26; j++ {
			if n1 == 0 || counts[j] == 0 || n1+counts[j] < maxAlternating {
				continue
			}
			cur, next := rune(i+'a'), rune(j+'a')
			i := strings.IndexAny(s, string([]byte{byte(next), byte(cur)}))
			if rune(s[i]) != cur {
				cur, next = next, cur
			}
			count := int32(1)
			for i := i + 1; i < len(s); i++ {
				switch rune(s[i]) {
				case next:
					count++
					cur, next = next, cur
				case cur:
					goto ContinueSearch
				}
			}
			if count > maxAlternating {
				maxAlternating = count
			}
		ContinueSearch:
		}
	}
	return maxAlternating
}
