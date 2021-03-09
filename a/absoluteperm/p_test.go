package absoluteperm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_absolutePermutation(t *testing.T) {
	in, err := os.Open("testdata/input")
	check(err)
	out, err := os.Open("testdata/output")
	check(err)
	inScanner := bufio.NewScanner(in)
	outScanner := bufio.NewScanner(out)
	outScanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	for inScanner.Scan() {
		outScanner.Scan()
		var n, k int32
		fmt.Sscanf(inScanner.Text(), "%d %d", &n, &k)
		s := outScanner.Text()
		wantStrs := strings.Split(s, " ")
		want := make([]int32, len(wantStrs))
		for i := range want {
			tmp, _ := strconv.Atoi(wantStrs[i])
			want[i] = int32(tmp)
		}
		got := absolutePermutation(n, k)
		require.Equal(t, len(want), len(got))
		for i := range want {
			if want[i] != got[i] {
				fmt.Println("lol!!")
			}
			require.Equal(t, want[i], got[i])
		}
		require.EqualValues(t, want, got)
		fmt.Println("done with first test!")
	}
}

// Complete the absolutePermutation function below.
func absolutePermutation(n int32, k int32) []int32 {
	if k == 0 {
		res := make([]int32, n)
		for i := int32(0); i < n; i++ {
			res[i] = i + 1
		}
		return res
	}
	if n%(k*2) != 0 {
		return []int32{-1}
	}
	res := make([]int32, n)
	var i int32
	for i < n {
		for j := int32(0); j < k; j++ {
			res[i] = i + 1 + k
			i++
		}
		for j := int32(0); j < k; j++ {
			res[i] = (i + 1) - k
			i++
		}
	}
	return res
}
