package biggerisgreater

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_biggerIsGreater(t *testing.T) {
	for _, tc := range []struct {
		w    string
		want string
	}{
		{"dkhc", "hcdk"},
		{"abdc", "acbd"},
		{"lmno", "lmon"},
		{"dcba", "no answer"},
		{"dcbb", "no answer"},
		{"abcd", "abdc"},
		{"fedcbabcd", "fedcbabdc"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.w), func(t *testing.T) {
			require.Equal(t, tc.want, biggerIsGreater(tc.w))
		})
	}

	t.Run("using file", func(t *testing.T) {
		f, err := os.Open("testdata/input")
		if err != nil {
			log.Fatalln(err)
		}
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			biggerIsGreater(sc.Text())
		}
	})
}

// Complete the biggerIsGreater function below.
func biggerIsGreater(w string) string {
	n := len(w)
	if n == 1 {
		return "no answer"
	}
	sb := []byte(w)
	i := n - 1
	for {
		if sb[i-1] < sb[i] {
			break
		}
		if i == 1 {
			return "no answer"
		}
		i--
	}

	// Find character closest to sb[i-1] and swap
	minIdx := i
	minVal := sb[i] - sb[i-1]
	for j := i + 1; j < n; j++ {
		if sb[j] > sb[i-1] && sb[j]-sb[i-1] < minVal {
			minVal = sb[j] - sb[i-1]
			minIdx = j
		}
	}

	sb[i-1], sb[minIdx] = sb[minIdx], sb[i-1]

	// Sort i in ascending order
	sort.Slice(sb[i:], func(j, k int) bool { return sb[i+j] < sb[i+k] })
	return string(sb)
}
