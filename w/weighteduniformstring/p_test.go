package weighteduniformstring

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

func Test_weightedUniformStrings(t *testing.T) {
	inPath := "testdata/input2"
	wantPath := "testdata/want2"
	f, _ := os.Open(inPath)
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Scan()
	s := sc.Text()
	sc.Scan() // n
	n := MustParseInt(sc.Text())
	queries := make([]int32, n)
	for i := 0; sc.Scan(); i++ {
		queries[i] = int32(MustParseInt(sc.Text()))
	}
	got := weightedUniformStrings(s, queries)
	_ = got
	f2, _ := os.Open(wantPath)
	sc = bufio.NewScanner(f2)
	want := make([]string, n)
	for i := 0; sc.Scan(); i++ {
		want[i] = sc.Text()
	}

	require.Equal(t, len(want), len(got))
	for i := range got {
		if got[i] != want[i] {
			t.Log(queries[i])
			t.Log(got[i])
			t.Log(want[i])
			t.Fail()
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Complete the weightedUniformStrings function below.
func weightedUniformStrings(s string, queries []int32) []string {
	var maxCounts [26]int
	count := 0
	for i := range s {
		count++
		if i > 1 && s[i] != s[i-1] {
			count = 1
		}
		k := int(s[i] - 'a')
		maxCounts[k] = max(maxCounts[k], count)
	}
	res := make([]string, len(queries))
	for i, n := range queries {
		for j := 1; j <= 26; j++ {
			if int(n)%j == 0 && maxCounts[j-1] >= int(n)/j {
				res[i] = "Yes"
				goto ContinueLoop
			}
		}
		res[i] = "No"
	ContinueLoop:
	}
	return res
}
